package conf

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var config *ComputeNode

const (
	TaskUbiContract    = "ubi_task"
	TokenContract      = "swan_token"
	CollateralContract = "swan_collateral"
	DefaultRpc         = "swan"
	BaseRpc            = "goerli"
)

// ComputeNode is a compute node config
type ComputeNode struct {
	API      API
	LOG      LOG
	HUB      HUB
	MCS      MCS
	Registry Registry
	RPC      RPC
	CONTRACT CONTRACT
}

type API struct {
	Port          int
	MultiAddress  string
	RedisUrl      string
	RedisPassword string
	Domain        string
	NodeName      string
}

type LOG struct {
	CrtFile string
	KeyFile string
}

type HUB struct {
	WalletAddress    string
	ServerUrl        string
	AccessToken      string
	BalanceThreshold float64
}

type MCS struct {
	ApiKey        string
	AccessToken   string
	BucketName    string
	Network       string
	FileCachePath string
}

type Registry struct {
	ServerAddress string
	UserName      string
	Password      string
}

type RPC struct {
	GoerliUrl   string `toml:"GOERLI_URL"`
	SwanTestnet string `toml:"SWAN_TESTNET"`
	SwanMainnet string `toml:"SWAN_MAINNET"`
}

type CONTRACT struct {
	SwanToken  string `toml:"SWAN_CONTRACT"`
	Collateral string `toml:"SWAN_COLLATERAL_CONTRACT"`
	UbiTask    string `toml:"UBI_TASK_CONTRACT"`
}

func GetRpcByName(rpcName string) (string, error) {
	var rpc string
	switch rpcName {
	case BaseRpc:
		rpc = GetConfig().RPC.GoerliUrl
		break
	case DefaultRpc:
		rpc = GetConfig().RPC.SwanTestnet
		break
	}
	return rpc, nil
}

func InitConfig(cpRepoPath string) error {
	configFile := filepath.Join(cpRepoPath, "config.toml")

	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		return fmt.Errorf("failed load config file, path: %s, error: %w", configFile, err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("Required fields not given")
		}
	}
	return nil
}

func GetConfig() *ComputeNode {
	return config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"API"},
		{"LOG"},
		{"HUB"},
		{"MCS"},
		{"Registry"},

		{"API", "MultiAddress"},
		{"API", "Domain"},
		{"API", "RedisUrl"},

		{"LOG", "CrtFile"},
		{"LOG", "KeyFile"},

		{"HUB", "ServerUrl"},
		{"HUB", "AccessToken"},
		{"HUB", "WalletAddress"},

		{"MCS", "ApiKey"},
		{"MCS", "BucketName"},
		{"MCS", "Network"},
		{"MCS", "FileCachePath"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("Required fields ", v)
		}
	}

	return true
}
