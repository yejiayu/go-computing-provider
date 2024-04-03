package conf

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var config *ComputeNode

const (
	DefaultRpc = "swan"

	BidMode_All     int = 0
	BidMode_Auto    int = 1
	BidMode_Private int = 2
	BidMode_None    int = 3
)

// ComputeNode is a compute node config
type ComputeNode struct {
	API      API
	UBI      UBI
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
	Domain        string
	NodeName      string
	RedisUrl      string
	RedisPassword string
}
type UBI struct {
	UbiTask     bool
	UbiEnginePk string
	UbiUrl      string
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
	OrchestratorPk   string
	VerifySign       bool
	BidMode          int
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
	SwanTestnet string `toml:"SWAN_TESTNET"`
	SwanMainnet string `toml:"SWAN_MAINNET"`
}

type CONTRACT struct {
	SwanToken  string `toml:"SWAN_CONTRACT"`
	Collateral string `toml:"SWAN_COLLATERAL_CONTRACT"`
}

func GetRpcByName(rpcName string) (string, error) {
	var rpc string
	switch rpcName {
	case DefaultRpc:
		rpc = GetConfig().RPC.SwanTestnet
		break
	}
	return rpc, nil
}

func InitConfig(cpRepoPath string, standalone bool) error {
	configFile := filepath.Join(cpRepoPath, "config.toml")
	metaData, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		return fmt.Errorf("failed load config file, path: %s, error: %w", configFile, err)
	}
	if standalone {
		if !requiredFieldsAreGivenForSeparate(metaData) {
			log.Fatal("Required fields not given")
		}
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
		{"UBI"},
		{"HUB"},
		{"MCS"},
		{"Registry"},
		{"RPC"},
		{"CONTRACT"},

		{"API", "MultiAddress"},
		{"API", "Domain"},
		{"API", "RedisUrl"},

		{"LOG", "CrtFile"},
		{"LOG", "KeyFile"},

		{"UBI", "UbiTask"},
		{"UBI", "UbiEnginePk"},
		{"UBI", "UbiUrl"},

		{"HUB", "ServerUrl"},
		{"HUB", "AccessToken"},
		{"HUB", "WalletAddress"},
		{"HUB", "BidMode"},

		{"MCS", "ApiKey"},
		{"MCS", "BucketName"},
		{"MCS", "Network"},
		{"MCS", "FileCachePath"},

		{"RPC", "SWAN_TESTNET"},

		{"CONTRACT", "SWAN_CONTRACT"},
		{"CONTRACT", "SWAN_COLLATERAL_CONTRACT"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("Required fields ", v)
		}
	}

	return true
}

func requiredFieldsAreGivenForSeparate(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"API"},
		{"UBI"},
		{"HUB"},

		{"API", "MultiAddress"},
		{"API", "RedisUrl"},

		{"UBI", "UbiTask"},
		{"UBI", "UbiEnginePk"},
		{"UBI", "UbiUrl"},

		{"HUB", "ServerUrl"},
		{"HUB", "AccessToken"},
		{"HUB", "BidMode"},

		{"RPC", "SWAN_TESTNET"},

		{"CONTRACT", "SWAN_CONTRACT"},
		{"CONTRACT", "SWAN_COLLATERAL_CONTRACT"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("Required fields ", v)
		}
	}

	return true
}
