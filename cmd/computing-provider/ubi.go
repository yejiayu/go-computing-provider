package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	cors "github.com/itsjamie/gin-cors"
	"github.com/olekukonko/tablewriter"
	"github.com/swanchain/go-computing-provider/account"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/initializer"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"github.com/urfave/cli/v2"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var ubiTaskCmd = &cli.Command{
	Name:  "ubi-task",
	Usage: "Manage ubi tasks",
	Subcommands: []*cli.Command{
		ubiTaskListCmd,
		daemonCmd,
	},
}

var ubiTaskListCmd = &cli.Command{
	Name:  "list",
	Usage: "List ubi task",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "show-failed",
			Usage: "show failed/failing ubi tasks",
		},
	},
	Action: func(cctx *cli.Context) error {

		cpPath, exit := os.LookupEnv("CP_PATH")
		if !exit {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=xxx")
		}
		if err := conf.InitConfig(cpPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		showFailed := cctx.Bool("show-failed")

		nodeID := computing.GetNodeId(cpPath)

		conn := computing.GetRedisClient()
		prefix := constants.REDIS_UBI_C2_PERFIX + "*"
		keys, err := redis.Strings(conn.Do("KEYS", prefix))
		if err != nil {
			return fmt.Errorf("failed get redis %s prefix, error: %+v", prefix, err)
		}

		var taskData [][]string
		var rowColorList []RowColor
		var taskList models.TaskList

		if showFailed {
			for _, key := range keys {
				ubiTask, err := computing.RetrieveUbiTaskMetadata(key)
				if err != nil {
					return fmt.Errorf("failed get ubi task: %s, error: %+v", key, err)
				}
				taskList = append(taskList, *ubiTask)
			}
		} else {
			for _, key := range keys {
				ubiTask, err := computing.RetrieveUbiTaskMetadata(key)
				if err != nil {
					return fmt.Errorf("failed get ubi task: %s, error: %+v", key, err)
				}
				if ubiTask.Status == constants.UBI_TASK_FAILED_STATUS {
					continue
				}
				taskList = append(taskList, *ubiTask)
			}
		}

		sort.Sort(taskList)
		for i, task := range taskList {

			reward, err := getReward(nodeID, task.TaskId)
			if err != nil {
				logs.GetLogger().Errorf("get task id: %s, reward failed, error: %v", task.TaskId, err)
			}

			taskData = append(taskData,
				[]string{task.TaskId, task.TaskType, task.ZkType, task.Tx, task.Status, reward, task.CreateTime})

			var rowColor []tablewriter.Colors
			if task.Status == constants.UBI_TASK_RECEIVED_STATUS {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgWhiteColor}}
			} else if task.Status == constants.UBI_TASK_RUNNING_STATUS {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgCyanColor}}
			} else if task.Status == constants.UBI_TASK_SUCCESS_STATUS {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
			} else if task.Status == constants.UBI_TASK_FAILED_STATUS {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
			}

			rowColorList = append(rowColorList, RowColor{
				row:    i,
				column: []int{4},
				color:  rowColor,
			})
		}

		header := []string{"TASK ID", "TASK TYPE", "ZK TYPE", "TRANSACTION HASH", "STATUS", "REWARD", "CREATE TIME"}
		NewVisualTable(header, taskData, rowColorList).Generate(true)

		return nil

	},
}

//go:embed docker-compose.yml
var dockerComposeContent string

var daemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a cp process",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "multi-address",
			Usage: "The multiAddress for libp2p(public ip)",
		},
	},
	Action: func(cctx *cli.Context) error {
		logs.GetLogger().Info("Start in computing provider mode.")

		cpRepoPath, exit := os.LookupEnv("CP_PATH")
		if !exit {
			cpRepoPath = cctx.String(FlagCpRepo)
			if len(strings.TrimSpace(cpRepoPath)) == 0 {
				return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=xxx or add --repo param")
			}
		}
		os.Setenv("CP_PATH", cpRepoPath)

		err := computing.StopPreviousServices(dockerComposeContent, cpRepoPath)
		if err != nil {
			return fmt.Errorf("stop pre-dependency-env failed, error: %v", err)
		}
		computing.NewDockerService().CleanResource()

		err = computing.RunDockerCompose(dockerComposeContent, cpRepoPath)
		if err != nil {
			return fmt.Errorf("start pre-dependency-env failed, error: %v", err)
		}

		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			logs.GetLogger().Fatal(err)
		}

		chainRpc, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			return err
		}
		client, err := ethclient.Dial(chainRpc)
		if err != nil {
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client)
		if err != nil {
			return err
		}
		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount from chain failed, error: %v", err)
		}

		conf.GetConfig().HUB.WalletAddress = cpAccount.OwnerAddress

		nodeId := computing.InitComputingProvider(cpRepoPath)
		go initializer.SendHeartbeats(nodeId)
		computing.ReportHardwareResource(nodeId)

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
		router := v1.Group("/computing")

		router.GET("/cp", computing.GetCpResource)
		router.GET("/cp/info", computing.GetCpInfo)
		router.POST("/cp/ubi", computing.DoUbiTaskForDocker)
		router.POST("/cp/docker/receive/ubi", computing.ReceiveUbiProofForDocker)

		shutdownChan := make(chan struct{})
		httpStopper, err := util.ServeHttp(r, "cp-api", ":"+strconv.Itoa(conf.GetConfig().API.Port), false)
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

func getReward(nodeId, taskId string) (string, error) {
	var taskInfo TaskInfo

	url := fmt.Sprintf("%s/rewards?node_id=%s&task_id=%s", conf.GetConfig().UBI.UbiUrl, nodeId, taskId)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get ubi task reward failed")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(bytes, &taskInfo)
	if err != nil {
		return "", err
	}
	if len(taskInfo.Data.List) > 0 {
		task := taskInfo.Data.List[0]
		floatVal, _ := strconv.ParseFloat(task.Amount, 64)
		return fmt.Sprintf("%.2f", floatVal), nil
	} else {
		return "0.0", nil
	}
}

type TaskInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total int `json:"total"`
		List  []struct {
			TaskId          int    `json:"task_id"`
			BeneficiaryAddr string `json:"beneficiary_addr"`
			Amount          string `json:"amount"`
			From            string `json:"from"`
			TxHash          string `json:"tx_hash"`
			ChainId         int    `json:"chain_id"`
			CreatedAt       int    `json:"created_at"`
		} `json:"list"`
	} `json:"data"`
}
