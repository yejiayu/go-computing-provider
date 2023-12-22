package computing

import (
	"encoding/json"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/lagrangedao/go-computing-provider/conf"
	"github.com/robfig/cron/v3"
	"io"
	"math"
	"net/http"
)

type CronTask struct {
}

func NewCronTask() *CronTask {
	return &CronTask{}
}

func (task *CronTask) RunTask() {
	task.checkCollateralBalance()
}

func (task *CronTask) checkCollateralBalance() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0/15 * * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [checkCollateralBalance], error: %+v", err)
			}
		}()

		url := fmt.Sprintf("%s/cp/collateral/%s", conf.GetConfig().HUB.ServerUrl, conf.GetConfig().HUB.WalletAddress)
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
			logs.GetLogger().Warnf("The current balance is: %0.3f，insufficient collateral Balance, collateral Operation Required. Please run: computing-provider collateral [fromWalletAddress] [amount]", result)
		}
	})
	c.Start()
}
