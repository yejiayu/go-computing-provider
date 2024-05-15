package account

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

type TaskStub struct {
	client          *ethclient.Client
	task            *Task
	privateK        string
	publicK         string
	ContractAddress string
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
	taskClient, err := NewTask(cpAccountAddress, client)
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
