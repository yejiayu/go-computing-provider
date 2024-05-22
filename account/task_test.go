package account

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestTaskAll(t *testing.T) {
	client, err := ethclient.Dial("https://rpc-atom-internal.swanchain.io")
	if err != nil {
		log.Println(err)
		return
	}
	defer client.Close()
	task, err := NewECPTask(common.HexToAddress("0xfb8b2F56c12F6739E231E51D9F0F2e09d70fA8F2"), client)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(task)
	info, err := task.GetTaskInfo(nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("task info: %#v\n", info)

	// 获取交易收据
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x26014ce734a4c6df32ed81c5b6f5b17a915c7c869cb82f6b1582a16ec74f256f"))
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}

	data, err := os.ReadFile("/Users/sonic/Documents/go_work/go-computing-provider/account/token.json")
	// 解析 ERC-20 合约 ABI
	contractAbi, err := abi.JSON(strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// 遍历交易日志
	for _, l := range receipt.Logs {
		// 尝试解析日志
		event := struct {
			From  common.Address
			To    common.Address
			Value *big.Int
		}{}
		if err := contractAbi.UnpackIntoInterface(&event, "Transfer", l.Data); err != nil {
			println("Failed to unpack log: %v", err)
			continue
		}

		// 检查索引参数
		if len(l.Topics) == 3 && l.Topics[0] == contractAbi.Events["Transfer"].ID {
			event.From = common.HexToAddress(l.Topics[1].Hex())
			event.To = common.HexToAddress(l.Topics[2].Hex())
			fmt.Printf("从 %s 转账到 %s, 数量为 %s 代币\n", event.From.Hex(), event.To.Hex(), event.Value.String())
		}
	}

}
