package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lagrangedao/go-computing-provider/conf"
	"math/big"
	"strings"
)

type CpStub struct {
	client   *ethclient.Client
	account  *Account
	privateK string
	publicK  string
}

type CpOption func(*CpStub)

func WithCpPrivateKey(pk string) CpOption {
	return func(obj *CpStub) {
		obj.privateK = pk
	}
}

func NewAccountStub(client *ethclient.Client, options ...CpOption) (*CpStub, error) {
	stub := &CpStub{}
	for _, option := range options {
		option(stub)
	}

	cpAccountAddress := common.HexToAddress(conf.GetConfig().CONTRACT.CpAccount)
	taskClient, err := NewAccount(cpAccountAddress, client)
	if err != nil {
		return nil, fmt.Errorf("create cp account contract client, error: %+v", err)
	}

	stub.account = taskClient
	stub.client = client
	return stub, nil
}

func (s *CpStub) SubmitUBIProof(taskId string, taskType uint8, zkType, proof string) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create transaction, error: %+v", publicAddress, err)
	}

	transaction, err := s.account.SubmitUBIProof(txOptions, taskId, taskType, zkType, proof)
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create SubmitUBIProof tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CpStub) ChangeMultiAddress(newMultiAddress []string) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create transaction, error: %+v", publicAddress, err)
	}

	transaction, err := s.account.ChangeMultiaddrs(txOptions, newMultiAddress)
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create ChangeMultiaddrs tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CpStub) ChangeOwnerAddress(newOwner common.Address) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create transaction, error: %+v", publicAddress, err)
	}

	transaction, err := s.account.ChangeOwnerAddress(txOptions, newOwner)
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create ChangeOwnerAddress tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CpStub) ChangeBeneficiary(newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create transaction, error: %+v", publicAddress, err)
	}

	transaction, err := s.account.ChangeBeneficiary(txOptions, newBeneficiary, newQuota, newExpiration)
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create ChangeBeneficiary tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CpStub) GetOwner() (string, error) {
	ownerAddress, err := s.account.GetOwner(&bind.CallOpts{})
	if err != nil {
		return "", fmt.Errorf("cpAccount client create GetOwner tx error: %+v", err)
	}
	return ownerAddress.Hex(), nil
}

func (s *CpStub) privateKeyToPublicKey() (common.Address, error) {
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

func (s *CpStub) createTransactOpts() (*bind.TransactOpts, error) {
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
