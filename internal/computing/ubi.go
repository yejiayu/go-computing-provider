package computing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"github.com/swanchain/go-computing-provider/account"
	"github.com/swanchain/go-computing-provider/build"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"github.com/swanchain/go-computing-provider/wallet"
	batchv1 "k8s.io/api/batch/v1"
	coreV1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

func DoUbiTaskForK8s(c *gin.Context) {

	var ubiTask models.UBITaskReq
	if err := c.ShouldBindJSON(&ubiTask); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("receive ubi task received: %+v", ubiTask)

	if ubiTask.ID == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: id"))
		return
	}
	if strings.TrimSpace(ubiTask.Name) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: name"))
		return
	}

	if ubiTask.Type != 0 && ubiTask.Type != 1 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "the value of task_type is 0 or 1"))
		return
	}
	if strings.TrimSpace(ubiTask.ZkType) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: zk_type"))
		return
	}

	if strings.TrimSpace(ubiTask.InputParam) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: input_param"))
		return
	}

	if strings.TrimSpace(ubiTask.Signature) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: signature"))
		return
	}

	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	nodeID := GetNodeId(cpRepoPath)

	signature, err := verifySignature(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%d", nodeID, ubiTask.ID), ubiTask.Signature)
	if err != nil {
		logs.GetLogger().Errorf("verifySignature for ubi task failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskParamError, "sign data failed"))
		return
	}

	logs.GetLogger().Infof("ubi task sign verifing, task_id: %d, type: %s, verify: %v", ubiTask.ID, ubiTask.ZkType, signature)
	if !signature {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskParamError, "signature verify failed"))
		return
	}

	var gpuFlag = "0"
	var ubiTaskToRedis = new(models.CacheUbiTaskDetail)
	ubiTaskToRedis.TaskId = strconv.Itoa(ubiTask.ID)
	ubiTaskToRedis.TaskType = "CPU"
	if ubiTask.Type == 1 {
		ubiTaskToRedis.TaskType = "GPU"
		gpuFlag = "1"
	}
	ubiTaskToRedis.Status = constants.UBI_TASK_RECEIVED_STATUS
	ubiTaskToRedis.ZkType = ubiTask.ZkType
	ubiTaskToRedis.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	SaveUbiTaskMetadata(ubiTaskToRedis)

	var envFilePath string
	envFilePath = filepath.Join(os.Getenv("CP_PATH"), "fil-c2.env")
	envVars, err := godotenv.Read(envFilePath)
	if err != nil {
		logs.GetLogger().Errorf("reading fil-c2-env.env failed, error: %v", err)
		return
	}

	c2GpuConfig := envVars["RUST_GPU_TOOLS_CUSTOM_GPU"]
	c2GpuName := convertGpuName(strings.TrimSpace(c2GpuConfig))
	nodeName, architecture, needCpu, needMemory, needStorage, err := checkResourceAvailableForUbi(ubiTask.Type, c2GpuName, ubiTask.Resource)
	if err != nil {
		ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
		SaveUbiTaskMetadata(ubiTaskToRedis)
		logs.GetLogger().Errorf("check resource failed, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	if nodeName == "" {
		ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
		SaveUbiTaskMetadata(ubiTaskToRedis)
		logs.GetLogger().Warnf("ubi task id: %d, type: %s, not found a resources available", ubiTask.ID, ubiTaskToRedis.TaskType)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckAvailableResources))
		return
	}

	var ubiTaskImage string
	if architecture == constants.CPU_AMD {
		ubiTaskImage = build.UBITaskImageAmdCpu
		if gpuFlag == "1" {
			ubiTaskImage = build.UBITaskImageAmdGpu
		}
	} else if architecture == constants.CPU_INTEL {
		ubiTaskImage = build.UBITaskImageIntelCpu
		if gpuFlag == "1" {
			ubiTaskImage = build.UBITaskImageIntelGpu
		}
	}

	mem := strings.Split(strings.TrimSpace(ubiTask.Resource.Memory), " ")[1]
	memUnit := strings.ReplaceAll(mem, "B", "")
	disk := strings.Split(strings.TrimSpace(ubiTask.Resource.Storage), " ")[1]
	diskUnit := strings.ReplaceAll(disk, "B", "")
	memQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needMemory, memUnit))
	if err != nil {
		ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
		SaveUbiTaskMetadata(ubiTaskToRedis)
		logs.GetLogger().Error("get memory failed, error: %+v", err)
		return
	}

	storageQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needStorage, diskUnit))
	if err != nil {
		ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
		SaveUbiTaskMetadata(ubiTaskToRedis)
		logs.GetLogger().Error("get storage failed, error: %+v", err)
		return
	}

	maxMemQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needMemory*2, memUnit))
	if err != nil {
		ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
		SaveUbiTaskMetadata(ubiTaskToRedis)
		logs.GetLogger().Error("get memory failed, error: %+v", err)
		return
	}

	maxStorageQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needStorage*2, diskUnit))
	if err != nil {
		ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
		SaveUbiTaskMetadata(ubiTaskToRedis)
		logs.GetLogger().Error("get storage failed, error: %+v", err)
		return
	}

	resourceRequirements := coreV1.ResourceRequirements{
		Limits: coreV1.ResourceList{
			coreV1.ResourceCPU:              *resource.NewQuantity(needCpu*2, resource.DecimalSI),
			coreV1.ResourceMemory:           maxMemQuantity,
			coreV1.ResourceEphemeralStorage: maxStorageQuantity,
			"nvidia.com/gpu":                resource.MustParse(gpuFlag),
		},
		Requests: coreV1.ResourceList{
			coreV1.ResourceCPU:              *resource.NewQuantity(needCpu, resource.DecimalSI),
			coreV1.ResourceMemory:           memQuantity,
			coreV1.ResourceEphemeralStorage: storageQuantity,
			"nvidia.com/gpu":                resource.MustParse(gpuFlag),
		},
	}

	go func() {
		var namespace = "ubi-task-" + strconv.Itoa(ubiTask.ID)
		var err error
		defer func() {
			key := constants.REDIS_UBI_C2_PERFIX + strconv.Itoa(ubiTask.ID)
			ubiTaskRun, err := RetrieveUbiTaskMetadata(key)
			if err != nil {
				logs.GetLogger().Errorf("get ubi task detail from db failed, ubiTaskId: %s, error: %+v", key, err)
				return
			}
			if ubiTaskRun.TaskId == "" {
				ubiTaskRun = new(models.CacheUbiTaskDetail)
				ubiTaskRun.TaskId = ubiTaskToRedis.TaskId
				ubiTaskRun.TaskType = ubiTaskToRedis.TaskType
				ubiTaskRun.ZkType = ubiTask.ZkType
				ubiTaskRun.CreateTime = ubiTaskToRedis.CreateTime
				SaveUbiTaskMetadata(ubiTaskRun)
			}
		}()

		k8sService := NewK8sService()
		if _, err = k8sService.GetNameSpace(context.TODO(), namespace, metaV1.GetOptions{}); err != nil {
			if errors.IsNotFound(err) {
				k8sNamespace := &v1.Namespace{
					ObjectMeta: metaV1.ObjectMeta{
						Name: namespace,
					},
				}
				_, err = k8sService.CreateNameSpace(context.TODO(), k8sNamespace, metaV1.CreateOptions{})
				if err != nil {
					logs.GetLogger().Errorf("create namespace failed, error: %v", err)
					return
				}
			}
		}

		receiveUrl := fmt.Sprintf("%s:%d/api/v1/computing/cp/receive/ubi", k8sService.GetAPIServerEndpoint(), conf.GetConfig().API.Port)
		execCommand := []string{"ubi-bench", "c2"}
		JobName := strings.ToLower(ubiTask.ZkType) + "-" + strconv.Itoa(ubiTask.ID)
		filC2Param := envVars["FIL_PROOFS_PARAMETER_CACHE"]
		if gpuFlag == "0" {
			delete(envVars, "RUST_GPU_TOOLS_CUSTOM_GPU")
			envVars["BELLMAN_NO_GPU"] = "1"
		}

		delete(envVars, "FIL_PROOFS_PARAMETER_CACHE")
		var useEnvVars []v1.EnvVar
		for k, v := range envVars {
			useEnvVars = append(useEnvVars, v1.EnvVar{
				Name:  k,
				Value: v,
			})
		}

		useEnvVars = append(useEnvVars, v1.EnvVar{
			Name:  "RECEIVE_PROOF_URL",
			Value: receiveUrl,
		},
			v1.EnvVar{
				Name:  "TASKID",
				Value: strconv.Itoa(ubiTask.ID),
			},
			v1.EnvVar{
				Name:  "TASK_TYPE",
				Value: strconv.Itoa(ubiTask.Type),
			},
			v1.EnvVar{
				Name:  "ZK_TYPE",
				Value: ubiTask.ZkType,
			},
			v1.EnvVar{
				Name:  "NAME_SPACE",
				Value: namespace,
			},
			v1.EnvVar{
				Name:  "PARAM_URL",
				Value: ubiTask.InputParam,
			},
		)

		job := &batchv1.Job{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      JobName,
				Namespace: namespace,
			},
			Spec: batchv1.JobSpec{
				Template: v1.PodTemplateSpec{
					Spec: v1.PodSpec{
						NodeName:     nodeName,
						NodeSelector: generateLabel(strings.ReplaceAll(c2GpuName, " ", "-")),
						Containers: []v1.Container{
							{
								Name:  JobName + generateString(5),
								Image: ubiTaskImage,
								Env:   useEnvVars,
								VolumeMounts: []v1.VolumeMount{
									{
										Name:      "proof-params",
										MountPath: "/var/tmp/filecoin-proof-parameters",
									},
								},
								Command:         execCommand,
								Resources:       resourceRequirements,
								ImagePullPolicy: coreV1.PullIfNotPresent,
							},
						},
						Volumes: []v1.Volume{
							{
								Name: "proof-params",
								VolumeSource: v1.VolumeSource{
									HostPath: &v1.HostPathVolumeSource{
										Path: filC2Param,
									},
								},
							},
						},
						RestartPolicy: "Never",
					},
				},
				BackoffLimit:            new(int32),
				TTLSecondsAfterFinished: new(int32),
			},
		}

		*job.Spec.BackoffLimit = 1
		*job.Spec.TTLSecondsAfterFinished = 120

		if _, err = k8sService.k8sClient.BatchV1().Jobs(namespace).Create(context.TODO(), job, metaV1.CreateOptions{}); err != nil {
			logs.GetLogger().Errorf("Failed creating ubi task job: %v", err)
			return
		}

		// TODO(yejiayu): Change it to the parameters you prefer or make it a configuration file.
		err = wait.PollImmediate(2*time.Second, 60*time.Second, func() (bool, error) {
			pods, err := k8sService.k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{
				LabelSelector: fmt.Sprintf("job-name=%s", JobName),
			})
			if err != nil {
				return false, err
			}

			for _, p := range pods.Items {
				for _, condition := range p.Status.Conditions {
					if condition.Type != corev1.PodReady && condition.Status != corev1.ConditionTrue {
						return false, nil
					}
				}
			}
			return true, nil
		})
		if err != nil {
			logs.GetLogger().Errorf("Failed waiting pods create: %v", err)
			return
		}

		pods, err := k8sService.k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{
			LabelSelector: fmt.Sprintf("job-name=%s", JobName),
		})
		if err != nil {
			logs.GetLogger().Errorf("Failed list ubi pods: %v", err)
			return
		}

		var podName string
		for _, pod := range pods.Items {
			podName = pod.Name
			break
		}

		req := k8sService.k8sClient.CoreV1().Pods(namespace).GetLogs(podName, &v1.PodLogOptions{
			Container:  "",
			Follow:     true,
			Timestamps: true,
		})

		time.Sleep(2 * time.Second)
		podLogs, err := req.Stream(context.Background())
		if err != nil {
			logs.GetLogger().Errorf("Error opening log stream: %v", err)
			return
		}
		defer podLogs.Close()

		ubiLogFileName := filepath.Join(cpRepoPath, "ubi.log")
		logFile, err := os.OpenFile(ubiLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logs.GetLogger().Errorf("opening ubi log file failed, error: %v", err)
			return
		}
		defer logFile.Close()

		if _, err = io.Copy(logFile, podLogs); err != nil {
			logs.GetLogger().Errorf("write ubi log to file failed, error: %v", err)
			return
		}
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func ReceiveUbiProofForK8s(c *gin.Context) {
	var c2Proof struct {
		TaskId    string `json:"task_id"`
		TaskType  string `json:"task_type"`
		Proof     string `json:"proof"`
		ZkType    string `json:"zk_type"`
		NameSpace string `json:"name_space"`
	}

	var submitUBIProofTx string
	var err error
	defer func() {
		key := constants.REDIS_UBI_C2_PERFIX + c2Proof.TaskId
		ubiTask, _ := RetrieveUbiTaskMetadata(key)
		if submitUBIProofTx != "" {
			ubiTask.Status = constants.UBI_TASK_SUCCESS_STATUS
		} else {
			ubiTask.Status = constants.UBI_TASK_FAILED_STATUS
		}
		ubiTask.Tx = submitUBIProofTx
		SaveUbiTaskMetadata(ubiTask)
		if strings.TrimSpace(c2Proof.NameSpace) != "" {
			k8sService := NewK8sService()
			k8sService.k8sClient.CoreV1().Namespaces().Delete(context.TODO(), c2Proof.NameSpace, metaV1.DeleteOptions{})
		}
	}()

	if err := c.ShouldBindJSON(&c2Proof); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("task_id: %s, C2 proof out received: %+v", c2Proof.TaskId, c2Proof)

	chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
	if err != nil {
		logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
		return
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
		return
	}
	defer client.Close()

	cpStub, err := account.NewAccountStub(client)
	if err != nil {
		logs.GetLogger().Errorf("create ubi task client failed, error: %v,", err)
		return
	}
	cpAccount, err := cpStub.GetCpAccountInfo()
	if err != nil {
		logs.GetLogger().Errorf("get account info failed, error: %v,", err)
		return
	}

	localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
	if err != nil {
		logs.GetLogger().Errorf("setup wallet failed, error: %v,", err)
		return
	}

	ki, err := localWallet.FindKey(cpAccount.OwnerAddress)
	if err != nil || ki == nil {
		logs.GetLogger().Errorf("the address: %s, private key %v,", conf.GetConfig().HUB.WalletAddress, wallet.ErrKeyInfoNotFound)
		return
	}

	accountStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
	if err != nil {
		logs.GetLogger().Errorf("create ubi task client failed, error: %v,", err)
		return
	}

	taskType, err := strconv.ParseUint(c2Proof.TaskType, 10, 8)
	if err != nil {
		logs.GetLogger().Errorf("conversion to uint8 error: %v", err)
		return
	}

	submitUBIProofTx, err = accountStub.SubmitUBIProof(c2Proof.TaskId, uint8(taskType), c2Proof.ZkType, c2Proof.Proof)
	if err != nil {
		logs.GetLogger().Errorf("submit ubi proof tx failed, error: %v,", err)
		return
	}

	fmt.Printf("submitUBIProofTx: %s", submitUBIProofTx)
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func DoUbiTaskForDocker(c *gin.Context) {

	var ubiTask models.UBITaskReq
	if err := c.ShouldBindJSON(&ubiTask); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("receive ubi task received: %+v", ubiTask)

	if ubiTask.ID == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: id"))
		return
	}
	if strings.TrimSpace(ubiTask.Name) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: name"))
		return
	}

	if ubiTask.Type != 0 && ubiTask.Type != 1 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "the value of task_type is 0 or 1"))
		return
	}
	if strings.TrimSpace(ubiTask.ZkType) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: zk_type"))
		return
	}

	if strings.TrimSpace(ubiTask.InputParam) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: input_param"))
		return
	}

	if strings.TrimSpace(ubiTask.Signature) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: signature"))
		return
	}

	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	nodeID := GetNodeId(cpRepoPath)

	signature, err := verifySignature(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%d", nodeID, ubiTask.ID), ubiTask.Signature)
	if err != nil {
		logs.GetLogger().Errorf("verifySignature for ubi task failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskParamError, "sign data failed"))
		return
	}

	logs.GetLogger().Infof("ubi task sign verifing, task_id: %d, type: %s, verify: %v", ubiTask.ID, ubiTask.ZkType, signature)
	if !signature {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskParamError, "signature verify failed"))
		return
	}

	var gpuFlag = "0"
	var ubiTaskToRedis = new(models.CacheUbiTaskDetail)
	ubiTaskToRedis.TaskId = strconv.Itoa(ubiTask.ID)
	ubiTaskToRedis.TaskType = "CPU"
	if ubiTask.Type == 1 {
		ubiTaskToRedis.TaskType = "GPU"
		gpuFlag = "1"
	}
	ubiTaskToRedis.Status = constants.UBI_TASK_RECEIVED_STATUS
	ubiTaskToRedis.ZkType = ubiTask.ZkType
	ubiTaskToRedis.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	SaveUbiTaskMetadata(ubiTaskToRedis)

	suffice, architecture, _, needMemory, err := checkResourceForUbi(ubiTask.Resource)
	if err != nil {
		ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
		SaveUbiTaskMetadata(ubiTaskToRedis)
		logs.GetLogger().Errorf("check resource failed, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	if !suffice {
		ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
		SaveUbiTaskMetadata(ubiTaskToRedis)
		logs.GetLogger().Warnf("ubi task id: %d, type: %s, not found a resources available", ubiTask.ID, ubiTaskToRedis.TaskType)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckAvailableResources))
		return
	}

	var ubiTaskImage string
	if architecture == constants.CPU_AMD {
		ubiTaskImage = build.UBITaskImageAmdCpu
		if gpuFlag == "1" {
			ubiTaskImage = build.UBITaskImageAmdGpu
		}
	} else if architecture == constants.CPU_INTEL {
		ubiTaskImage = build.UBITaskImageIntelCpu
		if gpuFlag == "1" {
			ubiTaskImage = build.UBITaskImageIntelGpu
		}
	}

	if err := NewDockerService().PullImage(ubiTaskImage); err != nil {
		logs.GetLogger().Errorf("pull %s image failed, error: %v", ubiTaskImage, err)
		return
	}

	go func() {
		defer func() {
			key := constants.REDIS_UBI_C2_PERFIX + strconv.Itoa(ubiTask.ID)
			ubiTaskRun, err := RetrieveUbiTaskMetadata(key)
			if err != nil {
				logs.GetLogger().Errorf("get ubi task detail from db failed, ubiTaskId: %s, error: %+v", key, err)
				return
			}
			if ubiTaskRun.TaskId == "" {
				ubiTaskRun = new(models.CacheUbiTaskDetail)
				ubiTaskRun.TaskId = ubiTaskToRedis.TaskId
				ubiTaskRun.TaskType = ubiTaskToRedis.TaskType
				ubiTaskRun.ZkType = ubiTask.ZkType
				ubiTaskRun.CreateTime = ubiTaskToRedis.CreateTime
				SaveUbiTaskMetadata(ubiTaskRun)
			}
		}()

		multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
		receiveUrl := fmt.Sprintf("http://%s:%s/api/v1/computing/cp/docker/receive/ubi", multiAddressSplit[2], multiAddressSplit[4])
		execCommand := []string{"ubi-bench", "c2"}
		JobName := strings.ToLower(ubiTask.ZkType) + "-" + strconv.Itoa(ubiTask.ID)

		var env = []string{"RECEIVE_PROOF_URL=" + receiveUrl}
		env = append(env, "TASKID="+strconv.Itoa(ubiTask.ID))
		env = append(env, "TASK_TYPE="+strconv.Itoa(ubiTask.Type))
		env = append(env, "ZK_TYPE="+ubiTask.ZkType)
		env = append(env, "NAME_SPACE=docker-ubi-task")
		env = append(env, "PARAM_URL="+ubiTask.InputParam)

		var needResource container.Resources
		if gpuFlag == "0" {
			env = append(env, "BELLMAN_NO_GPU=1")

			needResource = container.Resources{
				Memory: needMemory * 1024 * 1024 * 1024,
			}

		} else {
			gpuEnv, ok := os.LookupEnv("RUST_GPU_TOOLS_CUSTOM_GPU")
			if ok {
				env = append(env, "RUST_GPU_TOOLS_CUSTOM_GPU="+gpuEnv)
			}
			needResource = container.Resources{
				Memory: needMemory * 1024 * 1024 * 1024,
				DeviceRequests: []container.DeviceRequest{
					{
						Driver:       "nvidia",
						Count:        -1,
						Capabilities: [][]string{{"gpu"}},
						Options:      nil,
					},
				},
			}
		}

		filC2Param, ok := os.LookupEnv("FIL_PROOFS_PARAMETER_CACHE")
		if !ok {
			filC2Param = "/var/tmp/filecoin-proof-parameters"
		}

		hostConfig := &container.HostConfig{
			Binds:     []string{fmt.Sprintf("%s:/var/tmp/filecoin-proof-parameters", filC2Param)},
			Resources: needResource,
		}
		containerConfig := &container.Config{
			Image:        ubiTaskImage,
			Cmd:          execCommand,
			Env:          env,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		}

		containerName := JobName + generateString(5)
		dockerService := NewDockerService()
		if err = dockerService.ContainerCreateAndStart(containerConfig, hostConfig, containerName); err != nil {
			logs.GetLogger().Errorf("create ubi task container failed, error: %v", err)
		}

		containerLogStream, err := dockerService.GetContainerLogStream(containerName)
		if err != nil {
			logs.GetLogger().Errorf("get docker container log stream failed, error: %v", err)
		}
		defer containerLogStream.Close()

		ubiLogFileName := filepath.Join(cpRepoPath, "ubi.log")
		logFile, err := os.OpenFile(ubiLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logs.GetLogger().Errorf("opening ubi log file failed, error: %v", err)
			return
		}
		defer logFile.Close()

		if _, err = io.Copy(logFile, containerLogStream); err != nil {
			logs.GetLogger().Errorf("write ubi log to file failed, error: %v", err)
			return
		}
	}()
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func checkResourceForUbi(resource *models.TaskResource) (bool, string, int64, int64, error) {
	dockerService := NewDockerService()
	containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
	if err != nil {
		return false, "", 0, 0, err
	}

	var nodeResource models.NodeResource
	if err := json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
		logs.GetLogger().Error("collect host hardware resource failed, error: %+v", err)
		return false, "", 0, 0, err
	}

	needCpu, _ := strconv.ParseInt(resource.CPU, 10, 64)
	var needMemory, needStorage float64
	if len(strings.Split(strings.TrimSpace(resource.Memory), " ")) > 0 {
		needMemory, err = strconv.ParseFloat(strings.Split(strings.TrimSpace(resource.Memory), " ")[0], 64)

	}
	if len(strings.Split(strings.TrimSpace(resource.Storage), " ")) > 0 {
		needStorage, err = strconv.ParseFloat(strings.Split(strings.TrimSpace(resource.Storage), " ")[0], 64)
	}

	remainderCpu, _ := strconv.ParseInt(nodeResource.Cpu.Free, 10, 64)

	var remainderMemory, remainderStorage float64
	if len(strings.Split(strings.TrimSpace(nodeResource.Memory.Free), " ")) > 0 {
		remainderMemory, _ = strconv.ParseFloat(strings.Split(strings.TrimSpace(nodeResource.Memory.Free), " ")[0], 64)
	}
	if len(strings.Split(strings.TrimSpace(nodeResource.Storage.Free), " ")) > 0 {
		remainderStorage, err = strconv.ParseFloat(strings.Split(strings.TrimSpace(nodeResource.Storage.Free), " ")[0], 64)
	}

	var gpuMap = make(map[string]int)
	if nodeResource.Gpu.AttachedGpus > 0 {
		for _, detail := range nodeResource.Gpu.Details {
			if detail.Status == models.Available {
				gpuMap[detail.ProductName] += 1
			}
		}
	}

	logs.GetLogger().Infof("checkResourceForUbi: needCpu: %d, needMemory: %.2f, needStorage: %.2f", needCpu, needMemory, needStorage)
	logs.GetLogger().Infof("checkResourceForUbi: remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f, remainingGpu: %+v", remainderCpu, remainderMemory, remainderStorage, gpuMap)
	if needCpu <= remainderCpu && needMemory <= remainderMemory && needStorage <= remainderStorage {
		return true, nodeResource.CpuName, needCpu, int64(needMemory), nil
	}
	return false, nodeResource.CpuName, needCpu, int64(needMemory), nil
}

func ReceiveUbiProofForDocker(c *gin.Context) {
	var submitUBIProofTx string
	var err error
	var c2Proof models.UbiC2Proof
	defer func() {
		key := constants.REDIS_UBI_C2_PERFIX + c2Proof.TaskId
		ubiTask, _ := RetrieveUbiTaskMetadata(key)
		if submitUBIProofTx != "" {
			ubiTask.Status = constants.UBI_TASK_SUCCESS_STATUS
		} else {
			ubiTask.Status = constants.UBI_TASK_FAILED_STATUS
		}
		ubiTask.Tx = submitUBIProofTx
		SaveUbiTaskMetadata(ubiTask)
	}()

	if err := c.ShouldBindJSON(&c2Proof); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("task_id: %s, c2 proof out received: %+v", c2Proof.TaskId, c2Proof)

	retries := 3
	for i := 0; i < retries; i++ {
		submitUBIProofTx, err = submitUBIProof(c2Proof)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func GetCpResource(c *gin.Context) {
	location, err := getLocation()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed get location info"})
		return
	}

	dockerService := NewDockerService()
	containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
	if err != nil {
		logs.GetLogger().Error("collect host hardware resource failed, error: %+v", err)
		return
	}

	var nodeResource models.NodeResource
	if err := json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
		logs.GetLogger().Error("hardware info parse to json failed, error: %+v", err)
		return
	}

	cpRepo, _ := os.LookupEnv("CP_PATH")
	c.JSON(http.StatusOK, models.ClusterResource{
		Region:       location,
		ClusterInfo:  []*models.NodeResource{&nodeResource},
		MultiAddress: conf.GetConfig().API.MultiAddress,
		NodeName:     conf.GetConfig().API.NodeName,
		NodeId:       GetNodeId(cpRepo),
	})
}

func submitUBIProof(c2Proof models.UbiC2Proof) (string, error) {
	chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
	if err != nil {
		logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
		return "", err
	}
	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
		return "", err
	}
	client.Close()

	cpStub, err := account.NewAccountStub(client)
	if err != nil {
		logs.GetLogger().Errorf("create ubi task client failed, error: %v,", err)
		return "", err
	}
	cpAccount, err := cpStub.GetCpAccountInfo()
	if err != nil {
		logs.GetLogger().Errorf("get account info failed, error: %v,", err)
		return "", err
	}

	localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
	if err != nil {
		logs.GetLogger().Errorf("setup wallet failed, error: %v,", err)
		return "", err
	}

	ki, err := localWallet.FindKey(cpAccount.OwnerAddress)
	if err != nil || ki == nil {
		logs.GetLogger().Errorf("the address: %s, private key %v,", cpAccount.OwnerAddress, wallet.ErrKeyInfoNotFound)
		return "", err
	}

	accountStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
	if err != nil {
		logs.GetLogger().Errorf("create ubi task client failed, error: %v,", err)
		return "", err
	}

	taskType, err := strconv.ParseUint(c2Proof.TaskType, 10, 8)
	if err != nil {
		logs.GetLogger().Errorf("conversion to uint8 error: %v", err)
		return "", err
	}

	submitUBIProofTx, err := accountStub.SubmitUBIProof(c2Proof.TaskId, uint8(taskType), c2Proof.ZkType, c2Proof.Proof)
	if err != nil {
		logs.GetLogger().Errorf("submit ubi proof tx failed, error: %v,", err)
		return "", err
	}
	return submitUBIProofTx, nil
}

func getLocalIp() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}

			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("not found local ip")
}

func reportClusterResourceForDocker() {
	dockerService := NewDockerService()
	containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
	if err != nil {
		logs.GetLogger().Errorf("collect host hardware resource failed, error: %v", err)
		resourceExporterContainerName := "resource-exporter"
		err = dockerService.RemoveImageByName(resourceExporterContainerName)
		if err != nil {
			logs.GetLogger().Errorf("remove %s container failed, error: %v", resourceExporterContainerName, err)
			return
		}

		imageName := "filswan/hardware-exporter:v1.0"
		dockerService.PullImage(imageName)
		if err = dockerService.ContainerCreateAndStart(&container.Config{
			Image:        imageName,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		}, nil, resourceExporterContainerName); err != nil {
			logs.GetLogger().Errorf("create resource-exporter container failed, error: %v", err)
		}
		return
	}

	var nodeResource models.NodeResource
	if err := json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
		logs.GetLogger().Warnf("hardware info parse to json failed, restarting resource-exporter")
		resourceExporterContainerName := "resource-exporter"
		err = NewDockerService().RemoveImageByName(resourceExporterContainerName)
		if err != nil {
			logs.GetLogger().Errorf("remove %s container failed, error: %v", resourceExporterContainerName, err)
			return
		}
		imageName := "filswan/hardware-exporter:v1.0"
		dockerService.PullImage(imageName)
		if err = dockerService.ContainerCreateAndStart(&container.Config{
			Image:        imageName,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		}, nil, resourceExporterContainerName); err != nil {
			logs.GetLogger().Errorf("create resource-exporter container failed, error: %v", err)
		}
		return
	}

	var freeGpuMap = make(map[string]int)
	if nodeResource.Gpu.AttachedGpus > 0 {
		for _, g := range nodeResource.Gpu.Details {
			if g.Status == models.Available {
				freeGpuMap[g.ProductName] += 1
			} else {
				freeGpuMap[g.ProductName] = 0
			}
		}
	}
	logs.GetLogger().Infof("collect hardware resource, freeCpu:%s, freeMemory: %s, freeStorage: %s, freeGpu: %v",
		nodeResource.Cpu.Free, nodeResource.Memory.Free, nodeResource.Storage.Free, freeGpuMap)
}

func CleanDockerResource() {
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for range ticker.C {
			NewDockerService().CleanResource()
		}
	}()

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			reportClusterResourceForDocker()
		}
	}()

	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		for range ticker.C {
			conn := GetRedisClient()
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

				if ubiTask.Status != constants.UBI_TASK_SUCCESS_STATUS && ubiTask.Status != constants.UBI_TASK_FAILED_STATUS {
					ubiTask.Status = constants.UBI_TASK_FAILED_STATUS
				}
				if ubiTask.Tx != "" && ubiTask.Status != constants.UBI_TASK_SUCCESS_STATUS {
					ubiTask.Status = constants.UBI_TASK_SUCCESS_STATUS
				}
				SaveUbiTaskMetadata(ubiTask)
			}
		}
	}()
}
