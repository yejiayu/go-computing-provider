package account

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"sync"
)

type TaskStub struct {
	client          *ethclient.Client
	task            *ECPTask
	privateK        string
	publicK         string
	ContractAddress string
	nonceX          uint64
	taskL           sync.Mutex
}

type TaskOption func(*TaskStub)

func WithTaskPrivateKey(pk string) TaskOption {
	return func(obj *TaskStub) {
		obj.privateK = pk
	}
}

func WithTaskContractAddress(contractAddress string) TaskOption {
	return func(obj *TaskStub) {
		obj.ContractAddress = contractAddress
	}
}

func NewTaskStub(client *ethclient.Client, options ...TaskOption) (*TaskStub, error) {
	stub := &TaskStub{}
	for _, option := range options {
		option(stub)
	}

	if stub.ContractAddress == "" {
		return nil, errors.New("missing task contract address")
	}

	cpAccountAddress := common.HexToAddress(stub.ContractAddress)
	taskClient, err := NewECPTask(cpAccountAddress, client)
	if err != nil {
		return nil, fmt.Errorf("create task contract client, error: %+v", err)
	}

	stub.task = taskClient
	stub.client = client
	return stub, nil
}

func (s *TaskStub) SubmitUBIProof(proof string) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, task client submit ubi proof transaction, error: %+v", publicAddress, err)
	}
	transaction, err := s.task.SubmitProof(txOptions, proof)
	if err != nil {
		return "", fmt.Errorf("address: %s, task client submit ubi proof tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *TaskStub) GetTaskInfo() (ECPTaskTaskInfo, error) {
	return s.task.GetTaskInfo(&bind.CallOpts{})
}

func (s *TaskStub) privateKeyToPublicKey() (common.Address, error) {
	if len(strings.TrimSpace(s.privateK)) == 0 {
		return common.Address{}, fmt.Errorf("wallet address private key must be not empty")
	}

	privateKey, err := crypto.HexToECDSA(s.privateK)
	if err != nil {
		return common.Address{}, fmt.Errorf("parses private key error: %+v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func (s *TaskStub) createTransactOpts() (*bind.TransactOpts, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return nil, err
	}

	nonce, err := s.client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		return nil, fmt.Errorf("address: %s, collateral client get nonce error: %+v", publicAddress, err)
	}

	suggestGasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("address: %s, collateral client retrieves the currently suggested gas price, error: %+v", publicAddress, err)
	}

	chainId, err := s.client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("address: %s, collateral client get networkId, error: %+v", publicAddress, err)
	}

	privateKey, err := crypto.HexToECDSA(s.privateK)
	if err != nil {
		return nil, fmt.Errorf("parses private key error: %+v", err)
	}

	txOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("address: %s, collateral client create transaction, error: %+v", publicAddress, err)
	}
	txOptions.Nonce = big.NewInt(int64(nonce))
	suggestGasPrice = suggestGasPrice.Mul(suggestGasPrice, big.NewInt(3))
	suggestGasPrice = suggestGasPrice.Div(suggestGasPrice, big.NewInt(2))
	txOptions.GasFeeCap = suggestGasPrice
	txOptions.Context = context.Background()
	return txOptions, nil
}

func (s *TaskStub) getNonce() error {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return err
	}
	nonce, err := s.client.NonceAt(context.Background(), publicAddress, nil)
	if err != nil {
		return fmt.Errorf("address: %s, collateral client get nonce error: %+v", publicAddress, err)
	}
	s.taskL.Lock()
	defer s.taskL.Unlock()
	s.nonceX = nonce
	return nil
}

// GetReward  status: 1: Challenged  2: Slashed  3: rewarded
func (s *TaskStub) GetReward() (status int, reward string, err error) {
	reward = "0.0"
	taskInfo, err := s.GetTaskInfo()
	if err != nil {
		return 0, reward, err
	}

	if taskInfo.ChallengeTx != "" {
		return 1, reward, nil
	}

	if taskInfo.SlashTx != "" {
		return 2, reward, nil
	}

	if taskInfo.RewardTx != "" {
		receipt, err := s.client.TransactionReceipt(context.Background(), common.HexToHash(taskInfo.RewardTx))
		if err != nil {
			return 0, reward, err
		}
		contractAbi, err := abi.JSON(strings.NewReader(ECPTaskMetaData.ABI))
		if err != nil {
			return 0, reward, err
		}

		for _, l := range receipt.Logs {
			event := struct {
				From  common.Address
				To    common.Address
				Value *big.Int
			}{}
			if err := contractAbi.UnpackIntoInterface(&event, "Transfer", l.Data); err != nil {
				continue
			}

			if len(l.Topics) == 3 && l.Topics[0] == contractAbi.Events["Transfer"].ID {
				reward = balanceToStr(event.Value)
			}
		}
		return 3, reward, nil
	}
	return 0, reward, nil
}
