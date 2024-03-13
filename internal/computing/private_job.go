package computing

import (
	"encoding/json"
	"fmt"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func ReceivePrivateJob(c *gin.Context) {
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

	if len(strings.TrimSpace(jobData.Config.Description)) == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.description field"))
		return
	}

	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	nodeID := GetNodeId(cpRepoPath)

	signature, err := verifySignatureForHub(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%s", nodeID, jobData.SourceURI), jobData.Signature)
	if err != nil {
		logs.GetLogger().Errorf("verifySignature for private job failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, "verify sign data failed"))
		return
	}

	logs.GetLogger().Infof("private job sign verifing, task_id: %s,  verify: %v", jobData.UUID, signature)
	if !signature {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SpaceSignatureError, "signature verify failed"))
		return
	}

	available, gpuProductName, err := checkResourceAvailableForSpace(jobData.Config.Description)
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

	var hostName string
	var logHost string
	prefixStr := generateString(10)
	if strings.HasPrefix(conf.GetConfig().API.Domain, ".") {
		hostName = prefixStr + conf.GetConfig().API.Domain
		logHost = "log" + conf.GetConfig().API.Domain
	} else {
		hostName = strings.Join([]string{prefixStr, conf.GetConfig().API.Domain}, ".")
		logHost = "log." + conf.GetConfig().API.Domain
	}

	if _, err = celeryService.DelayTask(constants.TASK_DEPLOY, jobData.SourceURI, hostName, jobData.Duration, jobData.UUID, jobData.UUID, gpuProductName, jobData.User, true); err != nil {
		logs.GetLogger().Errorf("Failed sync delpoy task, error: %v", err)
		return
	}

	var privateJob models.PrivateJobResp
	privateJob.UUID = jobData.UUID
	privateJob.RealURI = fmt.Sprintf("https://%s", hostName)

	multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
	wsUrl := fmt.Sprintf("wss://%s:%s/api/v1/computing/lagrange/spaces/log?space_id=%s", logHost, multiAddressSplit[4], jobData.UUID)
	privateJob.BuildLog = wsUrl + "&type=build"
	privateJob.ContainerLog = wsUrl + "&type=container"

	if err = submitPrivateJob(&privateJob); err != nil {
		privateJob.ResultURI = ""
	}
	logs.GetLogger().Infof("submit private job detail: %+v", jobData)
	c.JSON(http.StatusOK, privateJob)
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

	jobData.Status = 1
	jobData.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	bytes, err := json.Marshal(jobData)
	if err != nil {
		logs.GetLogger().Errorf("Failed Marshal JobData, error: %v", err)
		return err
	}
	if err = os.WriteFile(taskDetailFilePath, bytes, os.ModePerm); err != nil {
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
