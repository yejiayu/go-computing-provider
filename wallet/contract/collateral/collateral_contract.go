// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package collateral

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

// CollateralMetaData contains all meta data concerning the Collateral contract.
var CollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receivingWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"LockCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"UnlockCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"lockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"taskBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalFrozenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unlockCollateral\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use CollateralMetaData.ABI instead.
var CollateralABI = CollateralMetaData.ABI

// Collateral is an auto generated Go binding around an Ethereum contract.
type Collateral struct {
	CollateralCaller     // Read-only binding to the contract
	CollateralTransactor // Write-only binding to the contract
	CollateralFilterer   // Log filterer for contract events
}

// CollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollateralSession struct {
	Contract     *Collateral       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollateralCallerSession struct {
	Contract *CollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollateralTransactorSession struct {
	Contract     *CollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollateralRaw struct {
	Contract *Collateral // Generic contract binding to access the raw methods on
}

// CollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollateralCallerRaw struct {
	Contract *CollateralCaller // Generic read-only contract binding to access the raw methods on
}

// CollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollateralTransactorRaw struct {
	Contract *CollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollateral creates a new instance of Collateral, bound to a specific deployed contract.
func NewCollateral(address common.Address, backend bind.ContractBackend) (*Collateral, error) {
	contract, err := bindCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Collateral{CollateralCaller: CollateralCaller{contract: contract}, CollateralTransactor: CollateralTransactor{contract: contract}, CollateralFilterer: CollateralFilterer{contract: contract}}, nil
}

// NewCollateralCaller creates a new read-only instance of Collateral, bound to a specific deployed contract.
func NewCollateralCaller(address common.Address, caller bind.ContractCaller) (*CollateralCaller, error) {
	contract, err := bindCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralCaller{contract: contract}, nil
}

// NewCollateralTransactor creates a new write-only instance of Collateral, bound to a specific deployed contract.
func NewCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*CollateralTransactor, error) {
	contract, err := bindCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralTransactor{contract: contract}, nil
}

// NewCollateralFilterer creates a new log filterer instance of Collateral, bound to a specific deployed contract.
func NewCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*CollateralFilterer, error) {
	contract, err := bindCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollateralFilterer{contract: contract}, nil
}

// bindCollateral binds a generic wrapper to an already deployed contract.
func bindCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CollateralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collateral *CollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Collateral.Contract.CollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collateral *CollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.Contract.CollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collateral *CollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collateral.Contract.CollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collateral *CollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Collateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collateral *CollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collateral *CollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collateral.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Collateral *CollateralCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Collateral *CollateralSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.Balances(&_Collateral.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Collateral *CollateralCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.Balances(&_Collateral.CallOpts, arg0)
}

// FrozenBalance is a free data retrieval call binding the contract method 0x266565a9.
//
// Solidity: function frozenBalance(address ) view returns(uint256)
func (_Collateral *CollateralCaller) FrozenBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "frozenBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenBalance is a free data retrieval call binding the contract method 0x266565a9.
//
// Solidity: function frozenBalance(address ) view returns(uint256)
func (_Collateral *CollateralSession) FrozenBalance(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.FrozenBalance(&_Collateral.CallOpts, arg0)
}

// FrozenBalance is a free data retrieval call binding the contract method 0x266565a9.
//
// Solidity: function frozenBalance(address ) view returns(uint256)
func (_Collateral *CollateralCallerSession) FrozenBalance(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.FrozenBalance(&_Collateral.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Collateral *CollateralCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Collateral *CollateralSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Collateral.Contract.IsAdmin(&_Collateral.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Collateral *CollateralCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Collateral.Contract.IsAdmin(&_Collateral.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Collateral *CollateralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Collateral *CollateralSession) Owner() (common.Address, error) {
	return _Collateral.Contract.Owner(&_Collateral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Collateral *CollateralCallerSession) Owner() (common.Address, error) {
	return _Collateral.Contract.Owner(&_Collateral.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Collateral *CollateralCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Collateral *CollateralSession) ProxiableUUID() ([32]byte, error) {
	return _Collateral.Contract.ProxiableUUID(&_Collateral.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Collateral *CollateralCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Collateral.Contract.ProxiableUUID(&_Collateral.CallOpts)
}

// TaskBalance is a free data retrieval call binding the contract method 0x637a570a.
//
// Solidity: function taskBalance(address ) view returns(uint256)
func (_Collateral *CollateralCaller) TaskBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "taskBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskBalance is a free data retrieval call binding the contract method 0x637a570a.
//
// Solidity: function taskBalance(address ) view returns(uint256)
func (_Collateral *CollateralSession) TaskBalance(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.TaskBalance(&_Collateral.CallOpts, arg0)
}

// TaskBalance is a free data retrieval call binding the contract method 0x637a570a.
//
// Solidity: function taskBalance(address ) view returns(uint256)
func (_Collateral *CollateralCallerSession) TaskBalance(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.TaskBalance(&_Collateral.CallOpts, arg0)
}

// TotalFrozenBalance is a free data retrieval call binding the contract method 0x19cf4697.
//
// Solidity: function totalFrozenBalance(address ) view returns(uint256)
func (_Collateral *CollateralCaller) TotalFrozenBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "totalFrozenBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFrozenBalance is a free data retrieval call binding the contract method 0x19cf4697.
//
// Solidity: function totalFrozenBalance(address ) view returns(uint256)
func (_Collateral *CollateralSession) TotalFrozenBalance(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.TotalFrozenBalance(&_Collateral.CallOpts, arg0)
}

// TotalFrozenBalance is a free data retrieval call binding the contract method 0x19cf4697.
//
// Solidity: function totalFrozenBalance(address ) view returns(uint256)
func (_Collateral *CollateralCallerSession) TotalFrozenBalance(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.TotalFrozenBalance(&_Collateral.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Collateral *CollateralCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Collateral *CollateralSession) Version() (*big.Int, error) {
	return _Collateral.Contract.Version(&_Collateral.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Collateral *CollateralCallerSession) Version() (*big.Int, error) {
	return _Collateral.Contract.Version(&_Collateral.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_Collateral *CollateralTransactor) AddAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "addAdmin", newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_Collateral *CollateralSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.AddAdmin(&_Collateral.TransactOpts, newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_Collateral *CollateralTransactorSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.AddAdmin(&_Collateral.TransactOpts, newAdmin)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address recipient) payable returns()
func (_Collateral *CollateralTransactor) Deposit(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "deposit", recipient)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address recipient) payable returns()
func (_Collateral *CollateralSession) Deposit(recipient common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Deposit(&_Collateral.TransactOpts, recipient)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address recipient) payable returns()
func (_Collateral *CollateralTransactorSession) Deposit(recipient common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Deposit(&_Collateral.TransactOpts, recipient)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address recipient, uint256 amount) payable returns()
func (_Collateral *CollateralTransactor) DepositETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "depositETH", recipient, amount)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address recipient, uint256 amount) payable returns()
func (_Collateral *CollateralSession) DepositETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.DepositETH(&_Collateral.TransactOpts, recipient, amount)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address recipient, uint256 amount) payable returns()
func (_Collateral *CollateralTransactorSession) DepositETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.DepositETH(&_Collateral.TransactOpts, recipient, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Collateral *CollateralTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Collateral *CollateralSession) Initialize() (*types.Transaction, error) {
	return _Collateral.Contract.Initialize(&_Collateral.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Collateral *CollateralTransactorSession) Initialize() (*types.Transaction, error) {
	return _Collateral.Contract.Initialize(&_Collateral.TransactOpts)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x22d39e28.
//
// Solidity: function lockCollateral(address taskContract, address[] cpList, uint256 collateral) returns()
func (_Collateral *CollateralTransactor) LockCollateral(opts *bind.TransactOpts, taskContract common.Address, cpList []common.Address, collateral *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "lockCollateral", taskContract, cpList, collateral)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x22d39e28.
//
// Solidity: function lockCollateral(address taskContract, address[] cpList, uint256 collateral) returns()
func (_Collateral *CollateralSession) LockCollateral(taskContract common.Address, cpList []common.Address, collateral *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.LockCollateral(&_Collateral.TransactOpts, taskContract, cpList, collateral)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x22d39e28.
//
// Solidity: function lockCollateral(address taskContract, address[] cpList, uint256 collateral) returns()
func (_Collateral *CollateralTransactorSession) LockCollateral(taskContract common.Address, cpList []common.Address, collateral *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.LockCollateral(&_Collateral.TransactOpts, taskContract, cpList, collateral)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Collateral *CollateralTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Collateral *CollateralSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.RemoveAdmin(&_Collateral.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Collateral *CollateralTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.RemoveAdmin(&_Collateral.TransactOpts, admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Collateral *CollateralTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Collateral *CollateralSession) RenounceOwnership() (*types.Transaction, error) {
	return _Collateral.Contract.RenounceOwnership(&_Collateral.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Collateral *CollateralTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Collateral.Contract.RenounceOwnership(&_Collateral.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Collateral *CollateralTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Collateral *CollateralSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.TransferOwnership(&_Collateral.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Collateral *CollateralTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.TransferOwnership(&_Collateral.TransactOpts, newOwner)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x3f001fd9.
//
// Solidity: function unlockCollateral(address recipient) payable returns()
func (_Collateral *CollateralTransactor) UnlockCollateral(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "unlockCollateral", recipient)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x3f001fd9.
//
// Solidity: function unlockCollateral(address recipient) payable returns()
func (_Collateral *CollateralSession) UnlockCollateral(recipient common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.UnlockCollateral(&_Collateral.TransactOpts, recipient)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x3f001fd9.
//
// Solidity: function unlockCollateral(address recipient) payable returns()
func (_Collateral *CollateralTransactorSession) UnlockCollateral(recipient common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.UnlockCollateral(&_Collateral.TransactOpts, recipient)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Collateral *CollateralTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Collateral *CollateralSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.UpgradeTo(&_Collateral.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Collateral *CollateralTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.UpgradeTo(&_Collateral.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Collateral *CollateralTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Collateral *CollateralSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Collateral.Contract.UpgradeToAndCall(&_Collateral.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Collateral *CollateralTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Collateral.Contract.UpgradeToAndCall(&_Collateral.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Collateral *CollateralTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Collateral *CollateralSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Withdraw(&_Collateral.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Collateral *CollateralTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Withdraw(&_Collateral.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Collateral *CollateralTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Collateral *CollateralSession) Receive() (*types.Transaction, error) {
	return _Collateral.Contract.Receive(&_Collateral.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Collateral *CollateralTransactorSession) Receive() (*types.Transaction, error) {
	return _Collateral.Contract.Receive(&_Collateral.TransactOpts)
}

// CollateralAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Collateral contract.
type CollateralAdminChangedIterator struct {
	Event *CollateralAdminChanged // Event containing the contract specifics and raw log

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
func (it *CollateralAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralAdminChanged)
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
		it.Event = new(CollateralAdminChanged)
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
func (it *CollateralAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralAdminChanged represents a AdminChanged event raised by the Collateral contract.
type CollateralAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Collateral *CollateralFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*CollateralAdminChangedIterator, error) {

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &CollateralAdminChangedIterator{contract: _Collateral.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Collateral *CollateralFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *CollateralAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralAdminChanged)
				if err := _Collateral.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Collateral *CollateralFilterer) ParseAdminChanged(log types.Log) (*CollateralAdminChanged, error) {
	event := new(CollateralAdminChanged)
	if err := _Collateral.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Collateral contract.
type CollateralBeaconUpgradedIterator struct {
	Event *CollateralBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *CollateralBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralBeaconUpgraded)
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
		it.Event = new(CollateralBeaconUpgraded)
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
func (it *CollateralBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralBeaconUpgraded represents a BeaconUpgraded event raised by the Collateral contract.
type CollateralBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Collateral *CollateralFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*CollateralBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &CollateralBeaconUpgradedIterator{contract: _Collateral.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Collateral *CollateralFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *CollateralBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralBeaconUpgraded)
				if err := _Collateral.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Collateral *CollateralFilterer) ParseBeaconUpgraded(log types.Log) (*CollateralBeaconUpgraded, error) {
	event := new(CollateralBeaconUpgraded)
	if err := _Collateral.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Collateral contract.
type CollateralDepositIterator struct {
	Event *CollateralDeposit // Event containing the contract specifics and raw log

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
func (it *CollateralDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralDeposit)
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
		it.Event = new(CollateralDeposit)
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
func (it *CollateralDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralDeposit represents a Deposit event raised by the Collateral contract.
type CollateralDeposit struct {
	FundingWallet   common.Address
	ReceivingWallet common.Address
	DepositAmount   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address fundingWallet, address receivingWallet, uint256 depositAmount)
func (_Collateral *CollateralFilterer) FilterDeposit(opts *bind.FilterOpts) (*CollateralDepositIterator, error) {

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &CollateralDepositIterator{contract: _Collateral.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address fundingWallet, address receivingWallet, uint256 depositAmount)
func (_Collateral *CollateralFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CollateralDeposit) (event.Subscription, error) {

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralDeposit)
				if err := _Collateral.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address fundingWallet, address receivingWallet, uint256 depositAmount)
func (_Collateral *CollateralFilterer) ParseDeposit(log types.Log) (*CollateralDeposit, error) {
	event := new(CollateralDeposit)
	if err := _Collateral.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Collateral contract.
type CollateralInitializedIterator struct {
	Event *CollateralInitialized // Event containing the contract specifics and raw log

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
func (it *CollateralInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralInitialized)
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
		it.Event = new(CollateralInitialized)
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
func (it *CollateralInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralInitialized represents a Initialized event raised by the Collateral contract.
type CollateralInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Collateral *CollateralFilterer) FilterInitialized(opts *bind.FilterOpts) (*CollateralInitializedIterator, error) {

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CollateralInitializedIterator{contract: _Collateral.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Collateral *CollateralFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CollateralInitialized) (event.Subscription, error) {

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralInitialized)
				if err := _Collateral.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Collateral *CollateralFilterer) ParseInitialized(log types.Log) (*CollateralInitialized, error) {
	event := new(CollateralInitialized)
	if err := _Collateral.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralLockCollateralIterator is returned from FilterLockCollateral and is used to iterate over the raw logs and unpacked data for LockCollateral events raised by the Collateral contract.
type CollateralLockCollateralIterator struct {
	Event *CollateralLockCollateral // Event containing the contract specifics and raw log

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
func (it *CollateralLockCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralLockCollateral)
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
		it.Event = new(CollateralLockCollateral)
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
func (it *CollateralLockCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralLockCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralLockCollateral represents a LockCollateral event raised by the Collateral contract.
type CollateralLockCollateral struct {
	TaskContract     common.Address
	CpList           []common.Address
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLockCollateral is a free log retrieval operation binding the contract event 0x725d92f258a928f1444ddfbc373ac2965cd5af9acfd0ec698099db2759db9415.
//
// Solidity: event LockCollateral(address taskContract, address[] cpList, uint256 collateralAmount)
func (_Collateral *CollateralFilterer) FilterLockCollateral(opts *bind.FilterOpts) (*CollateralLockCollateralIterator, error) {

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "LockCollateral")
	if err != nil {
		return nil, err
	}
	return &CollateralLockCollateralIterator{contract: _Collateral.contract, event: "LockCollateral", logs: logs, sub: sub}, nil
}

// WatchLockCollateral is a free log subscription operation binding the contract event 0x725d92f258a928f1444ddfbc373ac2965cd5af9acfd0ec698099db2759db9415.
//
// Solidity: event LockCollateral(address taskContract, address[] cpList, uint256 collateralAmount)
func (_Collateral *CollateralFilterer) WatchLockCollateral(opts *bind.WatchOpts, sink chan<- *CollateralLockCollateral) (event.Subscription, error) {

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "LockCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralLockCollateral)
				if err := _Collateral.contract.UnpackLog(event, "LockCollateral", log); err != nil {
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

// ParseLockCollateral is a log parse operation binding the contract event 0x725d92f258a928f1444ddfbc373ac2965cd5af9acfd0ec698099db2759db9415.
//
// Solidity: event LockCollateral(address taskContract, address[] cpList, uint256 collateralAmount)
func (_Collateral *CollateralFilterer) ParseLockCollateral(log types.Log) (*CollateralLockCollateral, error) {
	event := new(CollateralLockCollateral)
	if err := _Collateral.contract.UnpackLog(event, "LockCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Collateral contract.
type CollateralOwnershipTransferredIterator struct {
	Event *CollateralOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CollateralOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralOwnershipTransferred)
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
		it.Event = new(CollateralOwnershipTransferred)
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
func (it *CollateralOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralOwnershipTransferred represents a OwnershipTransferred event raised by the Collateral contract.
type CollateralOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Collateral *CollateralFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CollateralOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralOwnershipTransferredIterator{contract: _Collateral.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Collateral *CollateralFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CollateralOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralOwnershipTransferred)
				if err := _Collateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Collateral *CollateralFilterer) ParseOwnershipTransferred(log types.Log) (*CollateralOwnershipTransferred, error) {
	event := new(CollateralOwnershipTransferred)
	if err := _Collateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralUnlockCollateralIterator is returned from FilterUnlockCollateral and is used to iterate over the raw logs and unpacked data for UnlockCollateral events raised by the Collateral contract.
type CollateralUnlockCollateralIterator struct {
	Event *CollateralUnlockCollateral // Event containing the contract specifics and raw log

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
func (it *CollateralUnlockCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralUnlockCollateral)
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
		it.Event = new(CollateralUnlockCollateral)
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
func (it *CollateralUnlockCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralUnlockCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralUnlockCollateral represents a UnlockCollateral event raised by the Collateral contract.
type CollateralUnlockCollateral struct {
	TaskContract     common.Address
	Cp               common.Address
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnlockCollateral is a free log retrieval operation binding the contract event 0x0b07bdad0ef069f045a8095f09327d7282a29a5101416833dff00b4b2f888f77.
//
// Solidity: event UnlockCollateral(address taskContract, address cp, uint256 collateralAmount)
func (_Collateral *CollateralFilterer) FilterUnlockCollateral(opts *bind.FilterOpts) (*CollateralUnlockCollateralIterator, error) {

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "UnlockCollateral")
	if err != nil {
		return nil, err
	}
	return &CollateralUnlockCollateralIterator{contract: _Collateral.contract, event: "UnlockCollateral", logs: logs, sub: sub}, nil
}

// WatchUnlockCollateral is a free log subscription operation binding the contract event 0x0b07bdad0ef069f045a8095f09327d7282a29a5101416833dff00b4b2f888f77.
//
// Solidity: event UnlockCollateral(address taskContract, address cp, uint256 collateralAmount)
func (_Collateral *CollateralFilterer) WatchUnlockCollateral(opts *bind.WatchOpts, sink chan<- *CollateralUnlockCollateral) (event.Subscription, error) {

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "UnlockCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralUnlockCollateral)
				if err := _Collateral.contract.UnpackLog(event, "UnlockCollateral", log); err != nil {
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

// ParseUnlockCollateral is a log parse operation binding the contract event 0x0b07bdad0ef069f045a8095f09327d7282a29a5101416833dff00b4b2f888f77.
//
// Solidity: event UnlockCollateral(address taskContract, address cp, uint256 collateralAmount)
func (_Collateral *CollateralFilterer) ParseUnlockCollateral(log types.Log) (*CollateralUnlockCollateral, error) {
	event := new(CollateralUnlockCollateral)
	if err := _Collateral.contract.UnpackLog(event, "UnlockCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Collateral contract.
type CollateralUpgradedIterator struct {
	Event *CollateralUpgraded // Event containing the contract specifics and raw log

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
func (it *CollateralUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralUpgraded)
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
		it.Event = new(CollateralUpgraded)
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
func (it *CollateralUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralUpgraded represents a Upgraded event raised by the Collateral contract.
type CollateralUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Collateral *CollateralFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CollateralUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CollateralUpgradedIterator{contract: _Collateral.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Collateral *CollateralFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CollateralUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralUpgraded)
				if err := _Collateral.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Collateral *CollateralFilterer) ParseUpgraded(log types.Log) (*CollateralUpgraded, error) {
	event := new(CollateralUpgraded)
	if err := _Collateral.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Collateral contract.
type CollateralWithdrawIterator struct {
	Event *CollateralWithdraw // Event containing the contract specifics and raw log

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
func (it *CollateralWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralWithdraw)
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
		it.Event = new(CollateralWithdraw)
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
func (it *CollateralWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralWithdraw represents a Withdraw event raised by the Collateral contract.
type CollateralWithdraw struct {
	FundingWallet  common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address fundingWallet, uint256 withdrawAmount)
func (_Collateral *CollateralFilterer) FilterWithdraw(opts *bind.FilterOpts) (*CollateralWithdrawIterator, error) {

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &CollateralWithdrawIterator{contract: _Collateral.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address fundingWallet, uint256 withdrawAmount)
func (_Collateral *CollateralFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CollateralWithdraw) (event.Subscription, error) {

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralWithdraw)
				if err := _Collateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address fundingWallet, uint256 withdrawAmount)
func (_Collateral *CollateralFilterer) ParseWithdraw(log types.Log) (*CollateralWithdraw, error) {
	event := new(CollateralWithdraw)
	if err := _Collateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
