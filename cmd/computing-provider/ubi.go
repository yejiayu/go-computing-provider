package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/lagrangedao/go-computing-provider/conf"
	"github.com/lagrangedao/go-computing-provider/constants"
	"github.com/lagrangedao/go-computing-provider/internal/computing"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"os"
)

var ubiTaskCmd = &cli.Command{
	Name:  "ubi-task",
	Usage: "Manage ubi tasks",
	Subcommands: []*cli.Command{
		ubiTaskList,
	},
}

var ubiTaskList = &cli.Command{
	Name:  "list",
	Usage: "List ubi task",
	Action: func(cctx *cli.Context) error {

		cpPath, exit := os.LookupEnv("CP_PATH")
		if !exit {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=xxx")
		}
		if err := conf.InitConfig(cpPath); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		conn := computing.GetRedisClient()
		prefix := constants.REDIS_FULL_PREFIX + "*"
		keys, err := redis.Strings(conn.Do("KEYS", prefix))
		if err != nil {
			return fmt.Errorf("failed get redis %s prefix, error: %+v", prefix, err)
		}

		var taskData [][]string
		var rowColorList []RowColor
		var number int
		for _, key := range keys {
			ubiTask, err := computing.RetrieveUbiTaskMetadata(key)
			if err != nil {
				return fmt.Errorf("failed get ubi task: %s, error: %+v", key, err)
			}

			var status string

			taskData = append(taskData,
				[]string{ubiTask.TaskId, ubiTask.TaskType, ubiTask.ZkType, ubiTask.Tx, ubiTask.Status, status, ubiTask.Reward, ubiTask.CreateTime})

			var rowColor []tablewriter.Colors
			if ubiTask.Status == "success" {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
			} else {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
			}

			rowColorList = append(rowColorList, RowColor{
				row:    number,
				column: []int{4},
				color:  rowColor,
			})

			number++
		}

		header := []string{"TASK ID", "TASK TYPE", "ZK TYPE", "TRANSACTION HASH", "STATUS", "REWARD", " Create Time"}
		NewVisualTable(header, taskData, rowColorList).Generate()

		return nil

	},
}
