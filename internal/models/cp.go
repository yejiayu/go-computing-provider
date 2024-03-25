package models

import (
	"math/big"
	"time"
)

type BidStatus string

const (
	BidDisabledStatus    BidStatus = "bidding_disabled"
	BidEnabledStatus     BidStatus = "bidding_enabled"
	BidGpuDisabledStatus BidStatus = "bidding_gpu_disabled"

	ActiveStatus   string = "Active"
	InactiveStatus string = "Inactive"
)

type ComputingProvider struct {
	Name          string `json:"name"`
	NodeId        string `json:"node_id"`
	MultiAddress  string `json:"multi_address"`
	Autobid       int    `json:"autobid"`
	Status        string `json:"status"`
	PublicAddress string `json:"public_address"`
}

type JobData struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Duration int    `json:"duration"`
	//Hardware      string `json:"hardware"`
	JobSourceURI                string `json:"job_source_uri"`
	JobResultURI                string `json:"job_result_uri,omitempty"`
	StorageSource               string `json:"storage_source,omitempty"`
	TaskUUID                    string `json:"task_uuid"`
	CreatedAt                   string `json:"created_at"`
	UpdatedAt                   string `json:"updated_at,omitempty"`
	BuildLog                    string `json:"build_log,omitempty"`
	ContainerLog                string `json:"container_log"`
	NodeIdJobSourceUriSignature string `json:"node_id_job_source_uri_signature"`
	JobRealUri                  string `json:"job_real_uri,omitempty"`
}

type Job struct {
	Uuid   string
	Status JobStatus
	Url    string
}

type JobStatus string

const (
	JobDownloadSource JobStatus = "downloadSource" // download file form job_resource_uri
	JobUploadResult   JobStatus = "uploadResult"   // upload task result to mcs
	JobBuildImage     JobStatus = "buildImage"     // build images
	JobPushImage      JobStatus = "pushImage"      // push image to registry
	JobPullImage      JobStatus = "pullImage"      // download file form job_resource_uri
	JobDeployToK8s    JobStatus = "deployToK8s"    // deploy image to k8s
)

type DeleteJobReq struct {
	CreatorWallet string `json:"creator_wallet"`
	SpaceName     string `json:"space_name"`
}

type SpaceJSON struct {
	Data struct {
		Files []SpaceFile `json:"files"`
		Owner struct {
			PublicAddress string `json:"public_address"`
		} `json:"owner"`
		Space struct {
			Uuid        string `json:"uuid"`
			Name        string `json:"name"`
			ActiveOrder struct {
				Config SpaceHardware `json:"config"`
			} `json:"activeOrder"`
		} `json:"space"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type SpaceFile struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type SpaceHardware struct {
	Description  string `json:"description"`
	HardwareType string `json:"hardware_type"`
	Memory       int    `json:"memory"`
	Name         string `json:"name"`
	Vcpu         int    `json:"vcpu"`
}

type Resource struct {
	Cpu     Specification
	Memory  Specification
	Gpu     Specification
	Storage Specification
}

type Specification struct {
	Quantity int64
	Unit     string
}

type CacheSpaceDetail struct {
	WalletAddress string
	SpaceName     string
	SpaceUuid     string
	ExpireTime    int64
	JobUuid       string
	TaskType      string
	DeployName    string
	Hardware      string
	Url           string
	TaskUuid      string
	SpaceType     string
}

type UBITaskReq struct {
	ID         int           `json:"id"`
	Name       string        `json:"name,omitempty"`
	Type       int           `json:"type"`
	ZkType     string        `json:"zk_type"`
	InputParam string        `json:"input_param"`
	Signature  string        `json:"signature"`
	Resource   *TaskResource `json:"resource"`
}

type TaskResource struct {
	CPU     string `json:"cpu"`
	GPU     string `json:"gpu"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
}

type CacheUbiTaskDetail struct {
	TaskId     string `json:"task_id"`
	TaskType   string `json:"task_type"`
	ZkType     string `json:"zk_type"`
	Tx         string `json:"tx"`
	Status     string `json:"status"`
	Reward     string `json:"reward"`
	CreateTime string `json:"create_time"`
}

type Account struct {
	OwnerAddress   string
	NodeId         string
	MultiAddresses []string
	UbiFlag        uint8
	Beneficiary    struct {
		BeneficiaryAddress string
		Quota              *big.Int
		Expiration         *big.Int
	}
}

// ===== private job =====

type PrivateJobReq struct {
	UUID      string  `json:"uuid"`
	Name      string  `json:"name"`
	Duration  int     `json:"duration"`
	SourceURI string  `json:"job_source_uri"`
	Type      int     `json:"type"` // 1 ssh  2 space
	Source    string  `json:"storage_source"`
	Signature string  `json:"signature"`
	User      string  `json:"user"` // wallet
	Config    *Config `json:"config"`
}

type Config struct {
	Vcpu     int    `json:"vcpu"`
	Memory   int    `json:"memory"`
	Storage  int    `json:"storage"`
	GPU      int    `json:"gpu,omitempty"`
	GPUModel string `json:"gpu_model,omitempty"`
	Image    string `json:"image,omitempty"`
	SshKey   string `json:"ssh_key,omitempty"`
}

type PrivateJobResp struct {
	UUID         string `json:"uuid"`
	ResultURI    string `json:"result_uri"`
	RealURI      string `json:"real_uri,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty"`
	ContainerLog string `json:"container_log,omitempty"`
	BuildLog     string `json:"build_log,omitempty"`
	Status       int    `json:"status,omitempty"`
}

type PrivateJobStatusReq struct {
	UUID   string `json:"uuid"`
	Status int    `json:"status"` // 1: deploying; 2: running; 3: Available; 4: unavailable
}

type PrivateJobExtendReq struct {
	UUID     string `json:"uuid"`
	Duration int    `json:"duration"`
}

// ===== private job =====

type TaskList []CacheUbiTaskDetail

func (t TaskList) Len() int {
	return len(t)
}

func (t TaskList) Less(i, j int) bool {
	timeI, _ := time.Parse("2006-01-02 15:04:05", t[i].CreateTime)
	timeJ, _ := time.Parse("2006-01-02 15:04:05", t[j].CreateTime)
	return timeI.Before(timeJ)
}

func (t TaskList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
