package computing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gomodule/redigo/redis"
	"github.com/robfig/cron/v3"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/models"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var deployingChan = make(chan models.Job)
var TaskMap sync.Map

type CronTask struct {
	nodeId string
}

func NewCronTask(nodeId string) *CronTask {
	return &CronTask{nodeId: nodeId}
}

func (task *CronTask) RunTask() {
	addNodeLabel()
	checkJobStatus()
	task.checkCollateralBalance()
	task.cleanAbnormalDeployment()
	task.setFailedUbiTaskStatus()
	task.reportClusterResourceToHub()
	task.watchNameSpaceForDeleted()
	task.updateUbiTaskReward()
}

func checkJobStatus() {
	go func() {
		for {
			select {
			case job := <-deployingChan:
				TaskMap.Store(job.Uuid, &job)
			case <-time.After(3 * time.Second):
				TaskMap.Range(func(key, value any) bool {
					jobUuid := key.(string)
					job := value.(*models.Job)
					if reportJobStatus(jobUuid, job.Status) && job.Status == models.JobDeployToK8s {
						TaskMap.Delete(jobUuid)
					}
					return true
				})
			}
		}
	}()
}

func (task *CronTask) reportClusterResourceToHub() {
	ownerAddress, err := GetOwnerAddress()
	if err != nil {
		return
	}

	location, err := getLocation()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	c := cron.New(cron.WithSeconds())
	c.AddFunc("0/10 * * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("Failed report cp resource's summary, error: %+v", err)
			}
		}()

		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			reportClusterResource(location, task.nodeId, ownerAddress)
			checkClusterProviderStatus()
		}
	})
	c.Start()
}

func (task *CronTask) watchNameSpaceForDeleted() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/50 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("watchNameSpaceForDeleted catch panic error: %+v", err)
			}
		}()
		service := NewK8sService()
		namespaces, err := service.ListNamespace(context.TODO())
		if err != nil {
			logs.GetLogger().Errorf("Failed get all namespace, error: %+v", err)
			return
		}

		for _, namespace := range namespaces {
			getPods, err := service.GetPods(namespace, "")
			if err != nil {
				logs.GetLogger().Errorf("Failed get pods form namespace,namepace: %s, error: %+v", namespace, err)
				continue
			}
			if !getPods && (strings.HasPrefix(namespace, constants.K8S_NAMESPACE_NAME_PREFIX) || strings.HasPrefix(namespace, "ubi-task")) {
				if err = service.DeleteNameSpace(context.TODO(), namespace); err != nil {
					logs.GetLogger().Errorf("Failed delete namespace, namepace: %s, error: %+v", namespace, err)
				}
			}
		}
		NewDockerService().CleanResource()
	})
	c.Start()
}

func (task *CronTask) watchExpiredTask() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/5 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("watchExpiredTask catch panic error: %+v", err)
			}
		}()
		conn := redisPool.Get()
		prefix := constants.REDIS_SPACE_PREFIX + "*"
		keys, err := redis.Strings(conn.Do("KEYS", prefix))
		if err != nil {
			logs.GetLogger().Errorf("Failed get redis %s prefix, error: %+v", prefix, err)
			return
		}

		var deleteKey []string
		for _, key := range keys {
			jobMetadata, err := RetrieveJobMetadata(key)
			if err != nil {
				logs.GetLogger().Errorf("Failed get redis key data, key: %s, error: %+v", key, err)
				return
			}

			namespace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(jobMetadata.WalletAddress)
			if len(strings.TrimSpace(jobMetadata.TaskUuid)) == 0 {
				taskStatus, err := checkTaskStatusByHub(jobMetadata.TaskUuid, task.nodeId)
				if err != nil {
					logs.GetLogger().Errorf("Failed check task status by Orchestrator service, error: %+v", err)
					return
				}
				if strings.Contains(taskStatus, "Task not found") {
					logs.GetLogger().Infof("task_uuid: %s, task not found on the orchestrator service, starting to delete it.", jobMetadata.TaskUuid)
					deleteJob(namespace, jobMetadata.SpaceUuid)
					deleteKey = append(deleteKey, key)
					continue
				}
				if strings.Contains(taskStatus, "Terminated") || strings.Contains(taskStatus, "Terminated") ||
					strings.Contains(taskStatus, "Cancelled") || strings.Contains(taskStatus, "Failed") {
					logs.GetLogger().Infof("task_uuid: %s, current status is %s, starting to delete it.", jobMetadata.TaskUuid, taskStatus)
					if err = deleteJob(namespace, jobMetadata.SpaceUuid); err == nil {
						deleteKey = append(deleteKey, key)
						continue
					}
				}
			}

			if time.Now().Unix() > jobMetadata.ExpireTime {
				expireTimeStr := time.Unix(jobMetadata.ExpireTime, 0).Format("2006-01-02 15:04:05")
				logs.GetLogger().Infof("<timer-task> redis-key: %s,expireTime: %s. the job starting terminated", key, expireTimeStr)
				if err = deleteJob(namespace, jobMetadata.SpaceUuid); err == nil {
					deleteKey = append(deleteKey, key)
					continue
				}
			}

			k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(jobMetadata.WalletAddress)
			deployName := constants.K8S_DEPLOY_NAME_PREFIX + jobMetadata.SpaceUuid
			service := NewK8sService()
			if _, err = service.k8sClient.AppsV1().Deployments(k8sNameSpace).Get(context.TODO(), deployName, metav1.GetOptions{}); err != nil && errors.IsNotFound(err) {
				deleteKey = append(deleteKey, key)
				continue
			}
		}
		conn.Do("DEL", redis.Args{}.AddFlat(deleteKey)...)
		if len(deleteKey) > 0 {
			logs.GetLogger().Infof("Delete redis keys finished, keys: %+v", deleteKey)
			deleteKey = nil
		}
	})
	c.Start()
}

func (task *CronTask) checkCollateralBalance() {
	ownerAddress, err := GetOwnerAddress()
	if err != nil {
		return
	}

	c := cron.New(cron.WithSeconds())
	c.AddFunc("0/15 * * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [checkCollateralBalance], error: %+v", err)
			}
		}()

		url := fmt.Sprintf("%s/cp/collateral/%s", conf.GetConfig().HUB.ServerUrl, ownerAddress)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			logs.GetLogger().Errorf("create req failed: %+v", err)
			return
		}
		req.Header.Set("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logs.GetLogger().Errorf("send req failed: %+v", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logs.GetLogger().Errorf("read response failed: %+v", err)
			return
		}
		var collateral struct {
			Data struct {
				Balance float64 `json:"balance"`
			} `json:"data"`
			Message string `json:"message"`
			Status  string `json:"status"`
		}
		err = json.Unmarshal(body, &collateral)
		if err != nil {
			logs.GetLogger().Errorf("json conversion failed: %+v", err)
			return
		}

		result := collateral.Data.Balance / 1e18
		result = math.Round(result*1000) / 1000
		if result <= conf.GetConfig().HUB.BalanceThreshold {
			logs.GetLogger().Warnf("No sufficient collateral Balance, the current collateral balance is: %0.3f. Please run: computing-provider collateral [fromWalletAddress] [amount]", result)
		}
	})
	c.Start()
}

func (task *CronTask) cleanAbnormalDeployment() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* 0/30 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [cleanAbnormalDeployment], error: %+v", err)
			}
		}()

		k8sService := NewK8sService()
		namespaces, err := k8sService.ListNamespace(context.TODO())
		if err != nil {
			logs.GetLogger().Errorf("Failed get all namespace, error: %+v", err)
			return
		}

		for _, namespace := range namespaces {
			if strings.HasPrefix(namespace, constants.K8S_NAMESPACE_NAME_PREFIX) {
				deployments, err := k8sService.k8sClient.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
				if err != nil {
					logs.GetLogger().Errorf("Error getting deployments in namespace %s: %v\n", namespace, err)
					continue
				}

				for _, deployment := range deployments.Items {
					creationTimestamp := deployment.ObjectMeta.CreationTimestamp.Time
					currentTime := time.Now()
					age := currentTime.Sub(creationTimestamp)
					if (deployment.Status.AvailableReplicas == 0 && age.Hours() >= 2) || age.Hours() > 24*15 {
						logs.GetLogger().Infof("Cleaning up deployment %s in namespace %s", deployment.Name, namespace)
						err := k8sService.k8sClient.AppsV1().Deployments(namespace).Delete(context.TODO(), deployment.Name, metav1.DeleteOptions{})
						if err != nil {
							if errors.IsNotFound(err) {
								logs.GetLogger().Errorf("Deployment %s not found. Ignoring", deployment.Name)
							} else {
								logs.GetLogger().Errorf("Error deleting deployment %s: %v", deployment.Name, err)
							}
						} else {
							logs.GetLogger().Errorf("Deployment %s deleted successfully.", deployment.Name)
						}
					}
				}
			}
		}
	})
	c.Start()
}

func (task *CronTask) setFailedUbiTaskStatus() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/8 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [cleanAbnormalDeployment], error: %+v", err)
			}
		}()

		conn := redisPool.Get()
		prefix := constants.REDIS_UBI_C2_PERFIX + "*"
		keys, err := redis.Strings(conn.Do("KEYS", prefix))
		if err != nil {
			logs.GetLogger().Errorf("Failed get redis %s prefix, error: %+v", prefix, err)
			return
		}
		for _, key := range keys {
			ubiTask, err := RetrieveUbiTaskMetadata(key)
			if err != nil {
				logs.GetLogger().Errorf("Failed get ubi task from redis, key: %s, error: %+v", key, err)
				return
			}

			JobName := strings.ToLower(ubiTask.ZkType) + "-" + ubiTask.TaskId
			k8sNameSpace := "ubi-task-" + ubiTask.TaskId

			service := NewK8sService()
			if _, err = service.k8sClient.BatchV1().Jobs(k8sNameSpace).Get(context.TODO(), JobName, metav1.GetOptions{}); err != nil && errors.IsNotFound(err) {
				if ubiTask.Status != constants.UBI_TASK_SUCCESS_STATUS && ubiTask.Status != constants.UBI_TASK_FAILED_STATUS {
					ubiTask.Status = constants.UBI_TASK_FAILED_STATUS
				}
			}

			if ubiTask.Tx != "" {
				ubiTask.Status = constants.UBI_TASK_SUCCESS_STATUS
			}
			SaveUbiTaskMetadata(ubiTask)

		}

	})
	c.Start()
}

func (task *CronTask) updateUbiTaskReward() {
	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	nodeId := GetNodeId(cpRepoPath)

	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/10 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [updateUbiTaskReward], error: %+v", err)
			}
		}()

		taskList, err := NewTaskService().GetTaskListNoReward()
		if err != nil {
			logs.GetLogger().Errorf("Failed get task list, error: %+v", err)
			return
		}

		for _, entity := range taskList {
			ubiTask := entity
			reward, err := getReward(nodeId, strconv.Itoa(int(ubiTask.Id)))
			if err != nil {
				logs.GetLogger().Errorf("get ubi task reward failed, taskId: %d, error: %v", ubiTask.Id, err)
				continue
			}
			ubiTask.Reward = reward
			NewTaskService().SaveTaskEntity(ubiTask)
		}

	})
	c.Start()
}

func addNodeLabel() {
	k8sService := NewK8sService()
	nodes, err := k8sService.k8sClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	nodeGpuInfoMap, err := k8sService.GetResourceExporterPodLog(context.TODO())
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Infof("collect all node: %d", len(nodes.Items))
	for _, node := range nodes.Items {
		cpNode := node
		if collectInfo, ok := nodeGpuInfoMap[cpNode.Name]; ok {
			for _, detail := range collectInfo.Gpu.Details {
				if err = k8sService.AddNodeLabel(cpNode.Name, detail.ProductName); err != nil {
					logs.GetLogger().Errorf("add node label, nodeName %s, gpuName: %s, error: %+v", cpNode.Name, detail.ProductName, err)
					continue
				}
			}
			err := k8sService.AddNodeLabel(cpNode.Name, collectInfo.CpuName)
			if err != nil {
				logs.GetLogger().Errorf("nodeName: %s, error: %v", cpNode.Name, err)
			}
		}
	}
}

func reportClusterResource(location, nodeId, ownerAddress string) {
	k8sService := NewK8sService()
	statisticalSources, err := k8sService.StatisticalSources(context.TODO())
	if err != nil {
		logs.GetLogger().Errorf("Failed k8s statistical sources, error: %+v", err)
		return
	}
	clusterSource := models.ClusterResource{
		NodeId:        nodeId,
		Region:        location,
		ClusterInfo:   statisticalSources,
		PublicAddress: ownerAddress,
	}

	payload, err := json.Marshal(clusterSource)
	if err != nil {
		logs.GetLogger().Errorf("Failed convert to json, error: %+v", err)
		return
	}

	client := &http.Client{}
	url := conf.GetConfig().HUB.ServerUrl + "/cp/summary"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		logs.GetLogger().Errorf("Error creating request: %v", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logs.GetLogger().Errorf("Failed send a request, error: %+v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logs.GetLogger().Errorf("report cluster node resources failed, status code: %d", resp.StatusCode)
		return
	}

}

func checkTaskStatusByHub(taskUuid, nodeId string) (string, error) {
	url := fmt.Sprintf("%s/check_task_status_with_node_id/%s/%s", conf.GetConfig().HUB.ServerUrl, taskUuid, nodeId)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("offset", "0")
	req.Header.Add("limit", "10")
	req.Header.Add("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var taskStatus struct {
		Data struct {
			JobStatus  string `json:"job_status"`
			TaskStatus string `json:"task_status"`
		} `json:"data"`
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	err = json.Unmarshal(respBody, &taskStatus)
	if err != nil {
		logs.GetLogger().Errorf("check_task_status_with_node_id resp: %s", string(respBody))
		return "", err
	}
	if taskStatus.Status == "failed" {
		return taskStatus.Message, nil
	}
	return taskStatus.Status, nil
}

func reportJobStatus(jobUuid string, jobStatus models.JobStatus) bool {
	ownerAddress, err := GetOwnerAddress()
	if err != nil {
		return false
	}

	reqParam := map[string]interface{}{
		"job_uuid":       jobUuid,
		"status":         jobStatus,
		"public_address": ownerAddress,
	}

	payload, err := json.Marshal(reqParam)
	if err != nil {
		logs.GetLogger().Errorf("Failed convert to json, error: %+v", err)
		return false
	}

	client := &http.Client{}
	url := conf.GetConfig().HUB.ServerUrl + "/job/status"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		logs.GetLogger().Errorf("Error creating request: %v", err)
		return false
	}
	req.Header.Set("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logs.GetLogger().Errorf("Failed send a request, error: %+v", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	logs.GetLogger().Debugf("report job status successfully. uuid: %s, status: %s", jobUuid, jobStatus)
	return true
}
