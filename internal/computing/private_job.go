package computing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"io"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func ReceivePrivateJob(c *gin.Context) {
	if conf.GetConfig().HUB.BidMode == conf.BidMode_Auto || conf.GetConfig().HUB.BidMode == conf.BidMode_None {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, "the provider does not accept private task"))
		return
	}

	var jobData models.PrivateJobReq
	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.GetLogger().Infof("private job received Data: %+v", jobData)

	if len(jobData.Signature) == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.SpaceSignatureError, "missing signature field"))
		return
	}

	if jobData.Config.Vcpu == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.Vcpu field"))
		return
	}

	if jobData.Config.Memory == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.Memory field"))
		return
	}

	if jobData.Config.Storage == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.Storage field"))
		return
	}

	if jobData.Config.GPU != 0 && len(strings.TrimSpace(jobData.Config.GPUModel)) == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.gpu_model field"))
	}

	if jobData.Config.GPU == 0 && len(strings.TrimSpace(jobData.Config.GPUModel)) != 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.gpu field"))
	}

	//cpRepoPath, _ := os.LookupEnv("CP_PATH")
	//nodeID := GetNodeId(cpRepoPath)
	//
	//signature, err := verifySignatureForHub(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%s", nodeID, jobData.UUID), jobData.Signature)
	//if err != nil {
	//	logs.GetLogger().Errorf("verifySignature for private job failed, error: %+v", err)
	//	c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, "verify sign data failed"))
	//	return
	//}
	//
	//logs.GetLogger().Infof("private job sign verifing, task_id: %s,  verify: %v", jobData.UUID, signature)
	//if !signature {
	//	c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SpaceSignatureError, "signature verify failed"))
	//	return
	//}

	available, gpuProductName, err := checkResourceAvailableForPrivate(jobData.Config.Vcpu, jobData.Config.Memory, jobData.Config.Storage, jobData.Config.GPUModel, jobData.Config.GPU)
	if err != nil {
		logs.GetLogger().Errorf("check job resource failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	if !available {
		logs.GetLogger().Warnf(" task id: %s, name: %s, not found a resources available", jobData.UUID, jobData.Name)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckAvailableResources))
		return
	}

	var logHost string
	if strings.HasPrefix(conf.GetConfig().API.Domain, ".") {
		logHost = "log" + conf.GetConfig().API.Domain
	} else {
		logHost = "log." + conf.GetConfig().API.Domain
	}

	if _, err = celeryService.DelayTask(constants.PRIVATE_DEPLOY, jobData.Name, logHost, jobData.Duration, jobData.UUID, gpuProductName, jobData.User,
		jobData.Config.Vcpu, jobData.Config.Memory, jobData.Config.Storage, jobData.Config.GPUModel, jobData.Config.SshKey, jobData.Config.Image, jobData.Config.GPU); err != nil {
		logs.GetLogger().Errorf("Failed sync delpoy task, error: %v", err)
		return
	}

	var privateJob models.PrivateJobResp
	privateJob.UUID = jobData.UUID
	multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
	wsUrl := fmt.Sprintf("wss://%s:%s/api/v1/computing/lagrange/spaces/log?space_id=%s", logHost, multiAddressSplit[4], jobData.UUID)
	privateJob.ContainerLog = wsUrl + "&type=container&order=private"
	privateJob.UpdatedAt = time.Now().Unix()
	logs.GetLogger().Infof("submit private job detail: %+v", jobData)
	c.JSON(http.StatusOK, util.CreateSuccessResponse(privateJob))
}

func ExtendJob(c *gin.Context) {
	var jobData struct {
		Uuid     string `json:"uuid"`
		Duration int    `json:"duration"`
	}

	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.GetLogger().Infof("extend private Job received: %+v", jobData)

	if strings.TrimSpace(jobData.Uuid) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: uuid"})
		return
	}

	if jobData.Duration == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: duration"})
		return
	}

	conn := redisPool.Get()
	prefix := constants.REDIS_SPACE_PREFIX + "*"
	keys, err := redis.Strings(conn.Do("KEYS", prefix))
	if err != nil {
		logs.GetLogger().Errorf("Failed get redis %s prefix, error: %+v", prefix, err)
		return
	}

	var spaceDetail models.CacheSpaceDetail
	for _, key := range keys {
		jobMetadata, err := RetrieveJobMetadata(key)
		if err != nil {
			logs.GetLogger().Errorf("Failed get redis key data, key: %s, error: %+v", key, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "query data failed"})
			return
		}
		if strings.EqualFold(jobMetadata.TaskUuid, jobData.Uuid) {
			spaceDetail = jobMetadata
			break
		}
	}

	redisKey := constants.REDIS_SPACE_PREFIX + spaceDetail.SpaceUuid
	leftTime := spaceDetail.ExpireTime - time.Now().Unix()
	if leftTime < 0 {
		c.JSON(http.StatusOK, map[string]string{
			"status":  "failed",
			"message": "The job was terminated due to its expiration date",
		})
		return
	} else {
		fullArgs := []interface{}{redisKey}
		fields := map[string]string{
			"wallet_address": spaceDetail.WalletAddress,
			"space_name":     spaceDetail.SpaceName,
			"expire_time":    strconv.Itoa(int(time.Now().Unix()) + int(leftTime) + jobData.Duration),
			"space_uuid":     spaceDetail.SpaceUuid,
			"job_uuid":       spaceDetail.JobUuid,
			"task_type":      spaceDetail.TaskType,
			"deploy_name":    spaceDetail.DeployName,
			"hardware":       spaceDetail.Hardware,
			"url":            "",
			"task_uuid":      spaceDetail.TaskUuid,
			"space_type":     spaceDetail.SpaceType,
		}

		for key, val := range fields {
			fullArgs = append(fullArgs, key, val)
		}
		redisConn := redisPool.Get()
		defer redisConn.Close()

		redisConn.Do("HSET", fullArgs...)
		redisConn.Do("SET", spaceDetail.SpaceUuid, "wait-delete", "EX", int(leftTime)+jobData.Duration)
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func TerminateJob(c *gin.Context) {
	taskUuid := c.Param("task_uuid")
	if taskUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task_uuid is required"})
		return
	}

	conn := redisPool.Get()
	prefix := constants.REDIS_SPACE_PREFIX + "*"
	keys, err := redis.Strings(conn.Do("KEYS", prefix))
	if err != nil {
		logs.GetLogger().Errorf("Failed get redis %s prefix, error: %+v", prefix, err)
		return
	}

	var jobDetail models.CacheSpaceDetail
	for _, key := range keys {
		jobMetadata, err := RetrieveJobMetadata(key)
		if err != nil {
			logs.GetLogger().Errorf("Failed get redis key data for , key: %s, error: %+v", key, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "query data failed"})
			return
		}
		if strings.EqualFold(jobMetadata.TaskUuid, taskUuid) {
			jobDetail = jobMetadata
			break
		}
	}

	if jobDetail.WalletAddress == "" {
		c.JSON(http.StatusOK, util.CreateSuccessResponse("deleted success"))
		return
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task_uuid: %s, delete space request failed, error: %+v", taskUuid, err)
				return
			}
		}()
		k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(jobDetail.WalletAddress)
		deleteJob(k8sNameSpace, jobDetail.SpaceUuid)
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse("deleted success"))
}

func submitPrivateJob(jobData *models.PrivateJobResp) error {
	logs.GetLogger().Printf("submitting private job...")
	oldMask := syscall.Umask(0)
	defer syscall.Umask(oldMask)

	fileCachePath := conf.GetConfig().MCS.FileCachePath
	folderPath := "jobs"
	jobDetailFile := filepath.Join(folderPath, uuid.NewString()+".json")
	os.MkdirAll(filepath.Join(fileCachePath, folderPath), os.ModePerm)
	taskDetailFilePath := filepath.Join(fileCachePath, jobDetailFile)

	jobBytes, err := json.Marshal(jobData)
	if err != nil {
		logs.GetLogger().Errorf("Failed Marshal JobData, error: %v", err)
		return err
	}
	if err = os.WriteFile(taskDetailFilePath, jobBytes, os.ModePerm); err != nil {
		logs.GetLogger().Errorf("Failed jobData write to file, error: %v", err)
		return err
	}

	storageService := NewStorageService()
	mcsOssFile, err := storageService.UploadFileToBucket(jobDetailFile, taskDetailFilePath, true)
	if err != nil {
		logs.GetLogger().Errorf("Failed upload file to bucket, error: %v", err)
		return err
	}
	logs.GetLogger().Infof("jobuuid: %s successfully submitted to IPFS", jobData.UUID)

	gatewayUrl, err := storageService.GetGatewayUrl()
	if err != nil {
		logs.GetLogger().Errorf("Failed get mcs ipfs gatewayUrl, error: %v", err)
		return err
	}
	jobData.ResultURI = *gatewayUrl + "/ipfs/" + mcsOssFile.PayloadCid
	return nil
}

func DeployPrivateTask(name string, logHost string, duration int, taskUuid string, gpuProductName string, walletAddress string, cpu, memory, storage int, gpModel, sshKey, image string, gpuNum int) {
	updatePrivateStatus(taskUuid, JobStatusDeploying, "", "")
	var success bool
	var spaceUuid string

	defer func() {
		if !success {
			k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(walletAddress)
			deleteJob(k8sNameSpace, spaceUuid)
		}

		if err := recover(); err != nil {
			logs.GetLogger().Errorf("deploy private task painc, error: %+v", err)
			return
		}
	}()

	spaceUuid = strings.ToLower(taskUuid)

	conn := redisPool.Get()
	fullArgs := []interface{}{constants.REDIS_SPACE_PREFIX + spaceUuid}
	fields := map[string]string{
		"wallet_address": walletAddress,
		"space_name":     name,
		"expire_time":    strconv.Itoa(int(time.Now().Unix()) + duration),
		"task_uuid":      taskUuid,
	}

	for key, val := range fields {
		fullArgs = append(fullArgs, key, val)
	}
	_, _ = conn.Do("HSET", fullArgs...)

	deploy := NewDeploy(spaceUuid, "", walletAddress, "", int64(duration), taskUuid, constants.SPACE_TYPE_PRIVATE)
	deploy.WithSpaceInfo(spaceUuid, name)
	deploy.WithGpuProductName(gpuProductName)
	deploy.WithHardware(cpu, memory, storage, gpModel, gpuNum)
	deploy.WithImage(image)

	sshUrl, err := deploy.WithSshKey(sshKey).DeploySshTaskToK8s()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	var privateJob models.PrivateJobResp
	privateJob.UUID = taskUuid
	privateJob.RealURI = sshUrl

	multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
	privateJob.ContainerLog = fmt.Sprintf("wss://%s:%s/api/v1/computing/lagrange/spaces/log?space_id=%s&type=container&order=private",
		logHost, multiAddressSplit[4], taskUuid)
	privateJob.UpdatedAt = time.Now().Unix()
	privateJob.RealURI = sshUrl
	if err = submitPrivateJob(&privateJob); err != nil {
		privateJob.ResultURI = ""
	}
	updatePrivateStatus(taskUuid, JobStatusRunning, sshUrl, privateJob.ResultURI)

	success = true
}

type PrivateJobStatus int

const (
	JobStatusDeploying PrivateJobStatus = 3
	JobStatusRunning   PrivateJobStatus = 4
)

func updatePrivateStatus(jobUuid string, jobStatus PrivateJobStatus, result, resultUri string) {
	reqParam := make(map[string]interface{})
	reqParam["job_uuid"] = jobUuid
	reqParam["status"] = jobStatus
	if result != "" {
		reqParam["real_uri"] = result
	}
	if resultUri != "" {
		reqParam["result_uri"] = resultUri
	}

	payload, err := json.Marshal(reqParam)
	if err != nil {
		logs.GetLogger().Errorf("Failed convert to json, error: %+v", err)
		return
	}

	client := &http.Client{}
	url := conf.GetConfig().UBI.UbiUrl + "/jobs/" + jobUuid
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		logs.GetLogger().Errorf("Error creating request: %v", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logs.GetLogger().Errorf("Failed send a request, error: %+v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyData, _ := io.ReadAll(resp.Body)
		logs.GetLogger().Errorf("report private job status failed, %s", string(bodyData))
		return
	}
	return
}

func checkResourceAvailableForPrivate(cpu, memory, storage int, gpuModel string, gpuNum int) (bool, string, error) {
	taskType, hardwareDetail := getHardwareDetailForPrivate(cpu, memory, storage, gpuModel, gpuNum)
	k8sService := NewK8sService()

	activePods, err := k8sService.GetAllActivePod(context.TODO())
	if err != nil {
		return false, "", err
	}

	nodes, err := k8sService.k8sClient.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return false, "", err
	}

	nodeGpuSummary, err := k8sService.GetNodeGpuSummary(context.TODO())
	if err != nil {
		logs.GetLogger().Errorf("Failed collect k8s gpu, error: %+v", err)
		return false, "", err
	}

	for _, node := range nodes.Items {
		nodeGpu, remainderResource, _ := GetNodeResource(activePods, &node)
		remainderCpu := remainderResource[ResourceCpu]
		remainderMemory := float64(remainderResource[ResourceMem] / 1024 / 1024 / 1024)
		remainderStorage := float64(remainderResource[ResourceStorage] / 1024 / 1024 / 1024)

		needCpu := hardwareDetail.Cpu.Quantity
		needMemory := float64(hardwareDetail.Memory.Quantity)
		needStorage := float64(hardwareDetail.Storage.Quantity)
		logs.GetLogger().Infof("checkResourceAvailableForPrivate: needCpu: %d, needMemory: %.2f, needStorage: %.2f, gpuName: %s, gpuNum: %d", needCpu, needMemory, needStorage, gpuModel, gpuNum)
		logs.GetLogger().Infof("checkResourceAvailableForPrivate: remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f", remainderCpu, remainderMemory, remainderStorage)
		if needCpu <= remainderCpu && needMemory <= remainderMemory && needStorage <= remainderStorage {
			if taskType == "CPU" {
				return true, "", nil
			} else if taskType == "GPU" {
				var usedCount int64 = 0
				gpuName := strings.ToUpper(strings.ReplaceAll(hardwareDetail.Gpu.Unit, " ", "-"))
				logs.GetLogger().Infof("gpuName: %s, nodeGpu: %+v, nodeGpuSummary: %+v", gpuName, nodeGpu, nodeGpuSummary)
				var gpuProductName = ""
				for name, count := range nodeGpu {
					if strings.Contains(strings.ToUpper(name), gpuName) {
						usedCount = count
						gpuProductName = strings.ReplaceAll(strings.ToUpper(name), " ", "-")
						break
					}
				}

				for gName, gCount := range nodeGpuSummary[node.Name] {
					if strings.Contains(strings.ToUpper(gName), gpuName) {
						gpuProductName = strings.ReplaceAll(strings.ToUpper(gName), " ", "-")
						if usedCount+hardwareDetail.Gpu.Quantity <= gCount {
							return true, gpuProductName, nil
						}
					}
				}
				continue
			}
		}
	}
	return false, "", nil
}
