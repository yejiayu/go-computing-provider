// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ubi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"}],\"name\":\"ProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUri\",\"type\":\"string\"}],\"name\":\"TaskAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"addChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"taskUri\",\"type\":\"string\"}],\"name\":\"assignUBITask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"removeChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setChallengeDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"}],\"name\":\"submitUBIProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"taskUri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"challengeDeadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Main *MainCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Main *MainSession) ProxiableUUID() ([32]byte, error) {
	return _Main.Contract.ProxiableUUID(&_Main.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Main *MainCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Main.Contract.ProxiableUUID(&_Main.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(address cp, string nodeId, uint256 taskId, string taskUid, string taskUri, string proof, uint8 taskType, bool isCompleted, uint256 challengeDeadline, uint256 claimableAmount)
func (_Main *MainCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	Cp                common.Address
	NodeId            string
	TaskId            *big.Int
	TaskUid           string
	TaskUri           string
	Proof             string
	TaskType          uint8
	IsCompleted       bool
	ChallengeDeadline *big.Int
	ClaimableAmount   *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Cp                common.Address
		NodeId            string
		TaskId            *big.Int
		TaskUid           string
		TaskUri           string
		Proof             string
		TaskType          uint8
		IsCompleted       bool
		ChallengeDeadline *big.Int
		ClaimableAmount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cp = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NodeId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TaskId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TaskUid = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.TaskUri = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Proof = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.TaskType = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.IsCompleted = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.ChallengeDeadline = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.ClaimableAmount = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(address cp, string nodeId, uint256 taskId, string taskUid, string taskUri, string proof, uint8 taskType, bool isCompleted, uint256 challengeDeadline, uint256 claimableAmount)
func (_Main *MainSession) Tasks(arg0 string) (struct {
	Cp                common.Address
	NodeId            string
	TaskId            *big.Int
	TaskUid           string
	TaskUri           string
	Proof             string
	TaskType          uint8
	IsCompleted       bool
	ChallengeDeadline *big.Int
	ClaimableAmount   *big.Int
}, error) {
	return _Main.Contract.Tasks(&_Main.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(address cp, string nodeId, uint256 taskId, string taskUid, string taskUri, string proof, uint8 taskType, bool isCompleted, uint256 challengeDeadline, uint256 claimableAmount)
func (_Main *MainCallerSession) Tasks(arg0 string) (struct {
	Cp                common.Address
	NodeId            string
	TaskId            *big.Int
	TaskUid           string
	TaskUri           string
	Proof             string
	TaskType          uint8
	IsCompleted       bool
	ChallengeDeadline *big.Int
	ClaimableAmount   *big.Int
}, error) {
	return _Main.Contract.Tasks(&_Main.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Main *MainCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Main *MainSession) Version() (*big.Int, error) {
	return _Main.Contract.Version(&_Main.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Main *MainCallerSession) Version() (*big.Int, error) {
	return _Main.Contract.Version(&_Main.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Main *MainTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Main *MainSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddAdmin(&_Main.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Main *MainTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddAdmin(&_Main.TransactOpts, admin)
}

// AddChallenger is a paid mutator transaction binding the contract method 0x0ceb6780.
//
// Solidity: function addChallenger(address challenger) returns()
func (_Main *MainTransactor) AddChallenger(opts *bind.TransactOpts, challenger common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addChallenger", challenger)
}

// AddChallenger is a paid mutator transaction binding the contract method 0x0ceb6780.
//
// Solidity: function addChallenger(address challenger) returns()
func (_Main *MainSession) AddChallenger(challenger common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddChallenger(&_Main.TransactOpts, challenger)
}

// AddChallenger is a paid mutator transaction binding the contract method 0x0ceb6780.
//
// Solidity: function addChallenger(address challenger) returns()
func (_Main *MainTransactorSession) AddChallenger(challenger common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddChallenger(&_Main.TransactOpts, challenger)
}

// AssignUBITask is a paid mutator transaction binding the contract method 0x4228f5ca.
//
// Solidity: function assignUBITask(string taskUid, address cpAddress, string nodeId, string taskUri) returns()
func (_Main *MainTransactor) AssignUBITask(opts *bind.TransactOpts, taskUid string, cpAddress common.Address, nodeId string, taskUri string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "assignUBITask", taskUid, cpAddress, nodeId, taskUri)
}

// AssignUBITask is a paid mutator transaction binding the contract method 0x4228f5ca.
//
// Solidity: function assignUBITask(string taskUid, address cpAddress, string nodeId, string taskUri) returns()
func (_Main *MainSession) AssignUBITask(taskUid string, cpAddress common.Address, nodeId string, taskUri string) (*types.Transaction, error) {
	return _Main.Contract.AssignUBITask(&_Main.TransactOpts, taskUid, cpAddress, nodeId, taskUri)
}

// AssignUBITask is a paid mutator transaction binding the contract method 0x4228f5ca.
//
// Solidity: function assignUBITask(string taskUid, address cpAddress, string nodeId, string taskUri) returns()
func (_Main *MainTransactorSession) AssignUBITask(taskUid string, cpAddress common.Address, nodeId string, taskUri string) (*types.Transaction, error) {
	return _Main.Contract.AssignUBITask(&_Main.TransactOpts, taskUid, cpAddress, nodeId, taskUri)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xbb8c9797.
//
// Solidity: function claimReward(string taskUid) returns()
func (_Main *MainTransactor) ClaimReward(opts *bind.TransactOpts, taskUid string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "claimReward", taskUid)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xbb8c9797.
//
// Solidity: function claimReward(string taskUid) returns()
func (_Main *MainSession) ClaimReward(taskUid string) (*types.Transaction, error) {
	return _Main.Contract.ClaimReward(&_Main.TransactOpts, taskUid)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xbb8c9797.
//
// Solidity: function claimReward(string taskUid) returns()
func (_Main *MainTransactorSession) ClaimReward(taskUid string) (*types.Transaction, error) {
	return _Main.Contract.ClaimReward(&_Main.TransactOpts, taskUid)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Main *MainTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Main *MainSession) Initialize() (*types.Transaction, error) {
	return _Main.Contract.Initialize(&_Main.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Main *MainTransactorSession) Initialize() (*types.Transaction, error) {
	return _Main.Contract.Initialize(&_Main.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Main *MainTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Main *MainSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveAdmin(&_Main.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Main *MainTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveAdmin(&_Main.TransactOpts, admin)
}

// RemoveChallenger is a paid mutator transaction binding the contract method 0x6c578c1d.
//
// Solidity: function removeChallenger(address challenger) returns()
func (_Main *MainTransactor) RemoveChallenger(opts *bind.TransactOpts, challenger common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeChallenger", challenger)
}

// RemoveChallenger is a paid mutator transaction binding the contract method 0x6c578c1d.
//
// Solidity: function removeChallenger(address challenger) returns()
func (_Main *MainSession) RemoveChallenger(challenger common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveChallenger(&_Main.TransactOpts, challenger)
}

// RemoveChallenger is a paid mutator transaction binding the contract method 0x6c578c1d.
//
// Solidity: function removeChallenger(address challenger) returns()
func (_Main *MainTransactorSession) RemoveChallenger(challenger common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveChallenger(&_Main.TransactOpts, challenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Main.Contract.RenounceOwnership(&_Main.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Main.Contract.RenounceOwnership(&_Main.TransactOpts)
}

// SetChallengeDuration is a paid mutator transaction binding the contract method 0x22c3c1b3.
//
// Solidity: function setChallengeDuration(uint256 duration) returns()
func (_Main *MainTransactor) SetChallengeDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setChallengeDuration", duration)
}

// SetChallengeDuration is a paid mutator transaction binding the contract method 0x22c3c1b3.
//
// Solidity: function setChallengeDuration(uint256 duration) returns()
func (_Main *MainSession) SetChallengeDuration(duration *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetChallengeDuration(&_Main.TransactOpts, duration)
}

// SetChallengeDuration is a paid mutator transaction binding the contract method 0x22c3c1b3.
//
// Solidity: function setChallengeDuration(uint256 duration) returns()
func (_Main *MainTransactorSession) SetChallengeDuration(duration *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetChallengeDuration(&_Main.TransactOpts, duration)
}

// SetReward is a paid mutator transaction binding the contract method 0x1b5dd146.
//
// Solidity: function setReward(string taskUid, uint256 amount) returns()
func (_Main *MainTransactor) SetReward(opts *bind.TransactOpts, taskUid string, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setReward", taskUid, amount)
}

// SetReward is a paid mutator transaction binding the contract method 0x1b5dd146.
//
// Solidity: function setReward(string taskUid, uint256 amount) returns()
func (_Main *MainSession) SetReward(taskUid string, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetReward(&_Main.TransactOpts, taskUid, amount)
}

// SetReward is a paid mutator transaction binding the contract method 0x1b5dd146.
//
// Solidity: function setReward(string taskUid, uint256 amount) returns()
func (_Main *MainTransactorSession) SetReward(taskUid string, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetReward(&_Main.TransactOpts, taskUid, amount)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x16ef9c2a.
//
// Solidity: function submitUBIProof(address cpAddress, string nodeID, string taskUid, uint256 taskId, uint8 taskType, string proof) returns()
func (_Main *MainTransactor) SubmitUBIProof(opts *bind.TransactOpts, cpAddress common.Address, nodeID string, taskUid string, taskId *big.Int, taskType uint8, proof string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "submitUBIProof", cpAddress, nodeID, taskUid, taskId, taskType, proof)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x16ef9c2a.
//
// Solidity: function submitUBIProof(address cpAddress, string nodeID, string taskUid, uint256 taskId, uint8 taskType, string proof) returns()
func (_Main *MainSession) SubmitUBIProof(cpAddress common.Address, nodeID string, taskUid string, taskId *big.Int, taskType uint8, proof string) (*types.Transaction, error) {
	return _Main.Contract.SubmitUBIProof(&_Main.TransactOpts, cpAddress, nodeID, taskUid, taskId, taskType, proof)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x16ef9c2a.
//
// Solidity: function submitUBIProof(address cpAddress, string nodeID, string taskUid, uint256 taskId, uint8 taskType, string proof) returns()
func (_Main *MainTransactorSession) SubmitUBIProof(cpAddress common.Address, nodeID string, taskUid string, taskId *big.Int, taskType uint8, proof string) (*types.Transaction, error) {
	return _Main.Contract.SubmitUBIProof(&_Main.TransactOpts, cpAddress, nodeID, taskUid, taskId, taskType, proof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Main *MainTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Main *MainSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Main.Contract.UpgradeTo(&_Main.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Main *MainTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Main.Contract.UpgradeTo(&_Main.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Main *MainTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Main *MainSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Main.Contract.UpgradeToAndCall(&_Main.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Main *MainTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Main.Contract.UpgradeToAndCall(&_Main.TransactOpts, newImplementation, data)
}

// MainAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Main contract.
type MainAdminChangedIterator struct {
	Event *MainAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAdminChanged represents a AdminChanged event raised by the Main contract.
type MainAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Main *MainFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*MainAdminChangedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &MainAdminChangedIterator{contract: _Main.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Main *MainFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *MainAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAdminChanged)
				if err := _Main.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Main *MainFilterer) ParseAdminChanged(log types.Log) (*MainAdminChanged, error) {
	event := new(MainAdminChanged)
	if err := _Main.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Main contract.
type MainBeaconUpgradedIterator struct {
	Event *MainBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainBeaconUpgraded represents a BeaconUpgraded event raised by the Main contract.
type MainBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Main *MainFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*MainBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &MainBeaconUpgradedIterator{contract: _Main.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Main *MainFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *MainBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainBeaconUpgraded)
				if err := _Main.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Main *MainFilterer) ParseBeaconUpgraded(log types.Log) (*MainBeaconUpgraded, error) {
	event := new(MainBeaconUpgraded)
	if err := _Main.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Main contract.
type MainInitializedIterator struct {
	Event *MainInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainInitialized represents a Initialized event raised by the Main contract.
type MainInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Main *MainFilterer) FilterInitialized(opts *bind.FilterOpts) (*MainInitializedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MainInitializedIterator{contract: _Main.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Main *MainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MainInitialized) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainInitialized)
				if err := _Main.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Main *MainFilterer) ParseInitialized(log types.Log) (*MainInitialized, error) {
	event := new(MainInitialized)
	if err := _Main.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Main contract.
type MainOwnershipTransferredIterator struct {
	Event *MainOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainOwnershipTransferred represents a OwnershipTransferred event raised by the Main contract.
type MainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainOwnershipTransferredIterator{contract: _Main.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainOwnershipTransferred)
				if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) ParseOwnershipTransferred(log types.Log) (*MainOwnershipTransferred, error) {
	event := new(MainOwnershipTransferred)
	if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainProofSubmittedIterator is returned from FilterProofSubmitted and is used to iterate over the raw logs and unpacked data for ProofSubmitted events raised by the Main contract.
type MainProofSubmittedIterator struct {
	Event *MainProofSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainProofSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainProofSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainProofSubmitted represents a ProofSubmitted event raised by the Main contract.
type MainProofSubmitted struct {
	TaskUid string
	Proof   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitted is a free log retrieval operation binding the contract event 0x5bf50ff84da30bf962972249a44be3f7046a585680039b34329fda0374ee6de8.
//
// Solidity: event ProofSubmitted(string taskUid, string proof)
func (_Main *MainFilterer) FilterProofSubmitted(opts *bind.FilterOpts) (*MainProofSubmittedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ProofSubmitted")
	if err != nil {
		return nil, err
	}
	return &MainProofSubmittedIterator{contract: _Main.contract, event: "ProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchProofSubmitted is a free log subscription operation binding the contract event 0x5bf50ff84da30bf962972249a44be3f7046a585680039b34329fda0374ee6de8.
//
// Solidity: event ProofSubmitted(string taskUid, string proof)
func (_Main *MainFilterer) WatchProofSubmitted(opts *bind.WatchOpts, sink chan<- *MainProofSubmitted) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ProofSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainProofSubmitted)
				if err := _Main.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProofSubmitted is a log parse operation binding the contract event 0x5bf50ff84da30bf962972249a44be3f7046a585680039b34329fda0374ee6de8.
//
// Solidity: event ProofSubmitted(string taskUid, string proof)
func (_Main *MainFilterer) ParseProofSubmitted(log types.Log) (*MainProofSubmitted, error) {
	event := new(MainProofSubmitted)
	if err := _Main.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Main contract.
type MainRewardClaimedIterator struct {
	Event *MainRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRewardClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRewardClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRewardClaimed represents a RewardClaimed event raised by the Main contract.
type MainRewardClaimed struct {
	TaskUid      string
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x1f6879383a736478c14e8412bd7541665b9323615c5a087c069490f94efb9eeb.
//
// Solidity: event RewardClaimed(string taskUid, uint256 rewardAmount)
func (_Main *MainFilterer) FilterRewardClaimed(opts *bind.FilterOpts) (*MainRewardClaimedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return &MainRewardClaimedIterator{contract: _Main.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x1f6879383a736478c14e8412bd7541665b9323615c5a087c069490f94efb9eeb.
//
// Solidity: event RewardClaimed(string taskUid, uint256 rewardAmount)
func (_Main *MainFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *MainRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRewardClaimed)
				if err := _Main.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimed is a log parse operation binding the contract event 0x1f6879383a736478c14e8412bd7541665b9323615c5a087c069490f94efb9eeb.
//
// Solidity: event RewardClaimed(string taskUid, uint256 rewardAmount)
func (_Main *MainFilterer) ParseRewardClaimed(log types.Log) (*MainRewardClaimed, error) {
	event := new(MainRewardClaimed)
	if err := _Main.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRewardSetIterator is returned from FilterRewardSet and is used to iterate over the raw logs and unpacked data for RewardSet events raised by the Main contract.
type MainRewardSetIterator struct {
	Event *MainRewardSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRewardSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRewardSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRewardSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRewardSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRewardSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRewardSet represents a RewardSet event raised by the Main contract.
type MainRewardSet struct {
	TaskUid      string
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardSet is a free log retrieval operation binding the contract event 0xb08102b1d1c18789a33a8161fceb0da6c50c71149461848e12eca7fdaa523fc5.
//
// Solidity: event RewardSet(string taskUid, uint256 rewardAmount)
func (_Main *MainFilterer) FilterRewardSet(opts *bind.FilterOpts) (*MainRewardSetIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RewardSet")
	if err != nil {
		return nil, err
	}
	return &MainRewardSetIterator{contract: _Main.contract, event: "RewardSet", logs: logs, sub: sub}, nil
}

// WatchRewardSet is a free log subscription operation binding the contract event 0xb08102b1d1c18789a33a8161fceb0da6c50c71149461848e12eca7fdaa523fc5.
//
// Solidity: event RewardSet(string taskUid, uint256 rewardAmount)
func (_Main *MainFilterer) WatchRewardSet(opts *bind.WatchOpts, sink chan<- *MainRewardSet) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RewardSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRewardSet)
				if err := _Main.contract.UnpackLog(event, "RewardSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardSet is a log parse operation binding the contract event 0xb08102b1d1c18789a33a8161fceb0da6c50c71149461848e12eca7fdaa523fc5.
//
// Solidity: event RewardSet(string taskUid, uint256 rewardAmount)
func (_Main *MainFilterer) ParseRewardSet(log types.Log) (*MainRewardSet, error) {
	event := new(MainRewardSet)
	if err := _Main.contract.UnpackLog(event, "RewardSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTaskAssignedIterator is returned from FilterTaskAssigned and is used to iterate over the raw logs and unpacked data for TaskAssigned events raised by the Main contract.
type MainTaskAssignedIterator struct {
	Event *MainTaskAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainTaskAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTaskAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainTaskAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainTaskAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTaskAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTaskAssigned represents a TaskAssigned event raised by the Main contract.
type MainTaskAssigned struct {
	TaskUid string
	Cp      common.Address
	TaskUri string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTaskAssigned is a free log retrieval operation binding the contract event 0x76f7c7ed164ace167c3203b6fc03bb99e1748aed29011453a5d12c351653fbb9.
//
// Solidity: event TaskAssigned(string taskUid, address cp, string taskUri)
func (_Main *MainFilterer) FilterTaskAssigned(opts *bind.FilterOpts) (*MainTaskAssignedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "TaskAssigned")
	if err != nil {
		return nil, err
	}
	return &MainTaskAssignedIterator{contract: _Main.contract, event: "TaskAssigned", logs: logs, sub: sub}, nil
}

// WatchTaskAssigned is a free log subscription operation binding the contract event 0x76f7c7ed164ace167c3203b6fc03bb99e1748aed29011453a5d12c351653fbb9.
//
// Solidity: event TaskAssigned(string taskUid, address cp, string taskUri)
func (_Main *MainFilterer) WatchTaskAssigned(opts *bind.WatchOpts, sink chan<- *MainTaskAssigned) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "TaskAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTaskAssigned)
				if err := _Main.contract.UnpackLog(event, "TaskAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskAssigned is a log parse operation binding the contract event 0x76f7c7ed164ace167c3203b6fc03bb99e1748aed29011453a5d12c351653fbb9.
//
// Solidity: event TaskAssigned(string taskUid, address cp, string taskUri)
func (_Main *MainFilterer) ParseTaskAssigned(log types.Log) (*MainTaskAssigned, error) {
	event := new(MainTaskAssigned)
	if err := _Main.contract.UnpackLog(event, "TaskAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Main contract.
type MainUpgradedIterator struct {
	Event *MainUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainUpgraded represents a Upgraded event raised by the Main contract.
type MainUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Main *MainFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MainUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MainUpgradedIterator{contract: _Main.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Main *MainFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MainUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainUpgraded)
				if err := _Main.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Main *MainFilterer) ParseUpgraded(log types.Log) (*MainUpgraded, error) {
	event := new(MainUpgraded)
	if err := _Main.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
