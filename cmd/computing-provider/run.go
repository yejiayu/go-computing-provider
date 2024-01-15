package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/lagrangedao/go-computing-provider/conf"
	"github.com/lagrangedao/go-computing-provider/internal/computing"
	"github.com/lagrangedao/go-computing-provider/internal/initializer"
	ubi "github.com/lagrangedao/go-computing-provider/ubi/contract"
	"github.com/lagrangedao/go-computing-provider/util"
	"github.com/lagrangedao/go-computing-provider/wallet"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Start a cp process",
	Action: func(cctx *cli.Context) error {
		logs.GetLogger().Info("Start in computing provider mode.")

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		initializer.ProjectInit(cpRepoPath)

		r := gin.Default()
		r.Use(cors.Middleware(cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			ValidateHeaders: false,
		}))
		pprof.Register(r)

		v1 := r.Group("/api/v1")
		cpManager(v1.Group("/computing"))

		shutdownChan := make(chan struct{})
		httpStopper, err := util.ServeHttp(r, "cp-api", ":"+strconv.Itoa(conf.GetConfig().API.Port))
		if err != nil {
			logs.GetLogger().Fatal("failed to start cp-api endpoint: %s", err)
		}

		finishCh := util.MonitorShutdown(shutdownChan,
			util.ShutdownHandler{Component: "cp-api", StopFunc: httpStopper},
		)
		<-finishCh

		return nil
	},
}

func cpManager(router *gin.RouterGroup) {
	router.GET("/host/info", computing.GetServiceProviderInfo)
	router.POST("/lagrange/jobs", computing.ReceiveJob)
	router.POST("/lagrange/jobs/redeploy", computing.RedeployJob)
	router.DELETE("/lagrange/jobs", computing.CancelJob)
	router.GET("/lagrange/cp", computing.StatisticalSources)
	router.POST("/lagrange/jobs/renew", computing.ReNewJob)
	router.GET("/lagrange/spaces/log", computing.GetSpaceLog)
	router.POST("/lagrange/cp/proof", computing.DoProof)
	router.POST("/lagrange/cp/ubi", computing.DoUbiProof)
	router.POST("/lagrange/cp/receive/ubi", computing.ReceiveProof)
	router.POST("/lagrange/cp/test/ubi", computing.SendTask)

	router.GET("/cp", computing.StatisticalSources)
	router.GET("/cp/info", computing.GetCpInfo)
	router.POST("/cp/ubi", computing.DoUbiTask)
	router.POST("/cp/receive/ubi", computing.ReceiveUbiProof)

}

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize a new cp",
	Action: func(cctx *cli.Context) error {
		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		initializer.ProjectInit(cpRepoPath)

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(conf.GetConfig().HUB.WalletAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the address: %s, private key %w,", conf.GetConfig().HUB.WalletAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		privateKey, err := crypto.HexToECDSA(ki.PrivateKey)
		if err != nil {
			return err
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("failed to cast public key to ECDSA")
		}
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		contractSource := ""
		contractContent, err := os.ReadFile(contractSource)
		if err != nil {
			return err
		}

		contractSourceCode, err := abi.JSON(strings.NewReader(string(contractContent)))
		if err != nil {
			return err
		}

		deploymentData, err := contractSourceCode.Pack("", fromAddress, uint8(1), "aaa", []string{"/aa/bb"})
		if err != nil {
			return err
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return err
		}

		msg := ethereum.CallMsg{
			To:   nil,
			Data: deploymentData,
		}

		gas, err := client.EstimateGas(context.Background(), msg)
		if err != nil {
			return err
		}

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return err
		}

		tx := types.NewTransaction(nonce, fromAddress, nil, gas, gasPrice, deploymentData)
		signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
		if err != nil {
			return err
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			return err
		}

		fmt.Printf("deploy contract: %s", signedTx.Hash().String())

		return nil
	},
}

var changeMultiAddressCmd = &cli.Command{
	Name:  "change-multi-addr",
	Usage: "Update CP of MultiAddress",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "multi-addr",
			Usage: "Specify a new MultiAddress",
		},
	},
	Action: func(cctx *cli.Context) error {
		multiAddr := cctx.String("multi-addr")
		if strings.TrimSpace(multiAddr) == "" {
			return fmt.Errorf("multi-addr is not empty")
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		initializer.ProjectInit(cpRepoPath)

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(conf.GetConfig().HUB.WalletAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the address: %s, private key %w,", conf.GetConfig().HUB.WalletAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := ubi.NewCpStub(client, ubi.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			logs.GetLogger().Errorf("create ubi task client failed, error: %v,", err)
			return err
		}

		submitUBIProofTx, err := cpStub.ChangeMultiAddress([]string{multiAddr})
		if err != nil {
			logs.GetLogger().Errorf("change multi-addr tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeMultiAddress: %s", submitUBIProofTx)

		return nil
	},
}

var ChangeOwnerAddressCmd = &cli.Command{
	Name:  "change-owner-addr",
	Usage: "Update CP of OwnerAddress",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "owner-addr",
			Usage: "Specify a new OwnerAddress",
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddr := cctx.String("owner-addr")
		if strings.TrimSpace(ownerAddr) == "" {
			return fmt.Errorf("owner-addr is not empty")
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		initializer.ProjectInit(cpRepoPath)

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(conf.GetConfig().HUB.WalletAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the address: %s, private key %w,", conf.GetConfig().HUB.WalletAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := ubi.NewCpStub(client, ubi.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			logs.GetLogger().Errorf("create cp client failed, error: %v,", err)
			return err
		}

		submitUBIProofTx, err := cpStub.ChangeOwnerAddress(common.HexToAddress(ownerAddr))
		if err != nil {
			logs.GetLogger().Errorf("change owner address tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeOwnerAddress: %s", submitUBIProofTx)

		return nil
	},
}
