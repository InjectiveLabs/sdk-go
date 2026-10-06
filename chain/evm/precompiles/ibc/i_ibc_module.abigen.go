// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibc

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

// IIBCModuleTimeoutHeight is an auto generated low-level Go binding around an user-defined struct.
type IIBCModuleTimeoutHeight struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IBCModuleMetaData contains all meta data concerning the IBCModule contract.
var IBCModuleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"receiver\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structIIBCModule.TimeoutHeight\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revisionHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IBCModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use IBCModuleMetaData.ABI instead.
var IBCModuleABI = IBCModuleMetaData.ABI

// IBCModule is an auto generated Go binding around an Ethereum contract.
type IBCModule struct {
	IBCModuleCaller     // Read-only binding to the contract
	IBCModuleTransactor // Write-only binding to the contract
	IBCModuleFilterer   // Log filterer for contract events
}

// IBCModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBCModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBCModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBCModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBCModuleSession struct {
	Contract     *IBCModule        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBCModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBCModuleCallerSession struct {
	Contract *IBCModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IBCModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBCModuleTransactorSession struct {
	Contract     *IBCModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IBCModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBCModuleRaw struct {
	Contract *IBCModule // Generic contract binding to access the raw methods on
}

// IBCModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBCModuleCallerRaw struct {
	Contract *IBCModuleCaller // Generic read-only contract binding to access the raw methods on
}

// IBCModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBCModuleTransactorRaw struct {
	Contract *IBCModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBCModule creates a new instance of IBCModule, bound to a specific deployed contract.
func NewIBCModule(address common.Address, backend bind.ContractBackend) (*IBCModule, error) {
	contract, err := bindIBCModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBCModule{IBCModuleCaller: IBCModuleCaller{contract: contract}, IBCModuleTransactor: IBCModuleTransactor{contract: contract}, IBCModuleFilterer: IBCModuleFilterer{contract: contract}}, nil
}

// NewIBCModuleCaller creates a new read-only instance of IBCModule, bound to a specific deployed contract.
func NewIBCModuleCaller(address common.Address, caller bind.ContractCaller) (*IBCModuleCaller, error) {
	contract, err := bindIBCModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBCModuleCaller{contract: contract}, nil
}

// NewIBCModuleTransactor creates a new write-only instance of IBCModule, bound to a specific deployed contract.
func NewIBCModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*IBCModuleTransactor, error) {
	contract, err := bindIBCModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBCModuleTransactor{contract: contract}, nil
}

// NewIBCModuleFilterer creates a new log filterer instance of IBCModule, bound to a specific deployed contract.
func NewIBCModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*IBCModuleFilterer, error) {
	contract, err := bindIBCModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBCModuleFilterer{contract: contract}, nil
}

// bindIBCModule binds a generic wrapper to an already deployed contract.
func bindIBCModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBCModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBCModule *IBCModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBCModule.Contract.IBCModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBCModule *IBCModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBCModule.Contract.IBCModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBCModule *IBCModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBCModule.Contract.IBCModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBCModule *IBCModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBCModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBCModule *IBCModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBCModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBCModule *IBCModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBCModule.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0xc10876ab.
//
// Solidity: function transfer(string sourceChannel, string receiver, address token, uint256 amount, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp, string memo) returns(uint64 sequence)
func (_IBCModule *IBCModuleTransactor) Transfer(opts *bind.TransactOpts, sourceChannel string, receiver string, token common.Address, amount *big.Int, timeoutHeight IIBCModuleTimeoutHeight, timeoutTimestamp uint64, memo string) (*types.Transaction, error) {
	return _IBCModule.contract.Transact(opts, "transfer", sourceChannel, receiver, token, amount, timeoutHeight, timeoutTimestamp, memo)
}

// Transfer is a paid mutator transaction binding the contract method 0xc10876ab.
//
// Solidity: function transfer(string sourceChannel, string receiver, address token, uint256 amount, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp, string memo) returns(uint64 sequence)
func (_IBCModule *IBCModuleSession) Transfer(sourceChannel string, receiver string, token common.Address, amount *big.Int, timeoutHeight IIBCModuleTimeoutHeight, timeoutTimestamp uint64, memo string) (*types.Transaction, error) {
	return _IBCModule.Contract.Transfer(&_IBCModule.TransactOpts, sourceChannel, receiver, token, amount, timeoutHeight, timeoutTimestamp, memo)
}

// Transfer is a paid mutator transaction binding the contract method 0xc10876ab.
//
// Solidity: function transfer(string sourceChannel, string receiver, address token, uint256 amount, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp, string memo) returns(uint64 sequence)
func (_IBCModule *IBCModuleTransactorSession) Transfer(sourceChannel string, receiver string, token common.Address, amount *big.Int, timeoutHeight IIBCModuleTimeoutHeight, timeoutTimestamp uint64, memo string) (*types.Transaction, error) {
	return _IBCModule.Contract.Transfer(&_IBCModule.TransactOpts, sourceChannel, receiver, token, amount, timeoutHeight, timeoutTimestamp, memo)
}
