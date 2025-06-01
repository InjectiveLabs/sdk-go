// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// CosmosCoin is an auto generated low-level Go binding around an user-defined struct.
type CosmosCoin struct {
	Amount *big.Int
	Denom  string
}

// StakingModuleMetaData contains all meta data concerning the StakingModule contract.
var StakingModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin\",\"name\":\"balance\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorSrcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorDstAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"withdrawDelegatorRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingModuleMetaData.ABI instead.
var StakingModuleABI = StakingModuleMetaData.ABI

// StakingModule is an auto generated Go binding around an Ethereum contract.
type StakingModule struct {
	StakingModuleCaller     // Read-only binding to the contract
	StakingModuleTransactor // Write-only binding to the contract
	StakingModuleFilterer   // Log filterer for contract events
}

// StakingModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingModuleSession struct {
	Contract     *StakingModule    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingModuleCallerSession struct {
	Contract *StakingModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakingModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingModuleTransactorSession struct {
	Contract     *StakingModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingModuleRaw struct {
	Contract *StakingModule // Generic contract binding to access the raw methods on
}

// StakingModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingModuleCallerRaw struct {
	Contract *StakingModuleCaller // Generic read-only contract binding to access the raw methods on
}

// StakingModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingModuleTransactorRaw struct {
	Contract *StakingModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingModule creates a new instance of StakingModule, bound to a specific deployed contract.
func NewStakingModule(address common.Address, backend bind.ContractBackend) (*StakingModule, error) {
	contract, err := bindStakingModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingModule{StakingModuleCaller: StakingModuleCaller{contract: contract}, StakingModuleTransactor: StakingModuleTransactor{contract: contract}, StakingModuleFilterer: StakingModuleFilterer{contract: contract}}, nil
}

// NewStakingModuleCaller creates a new read-only instance of StakingModule, bound to a specific deployed contract.
func NewStakingModuleCaller(address common.Address, caller bind.ContractCaller) (*StakingModuleCaller, error) {
	contract, err := bindStakingModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingModuleCaller{contract: contract}, nil
}

// NewStakingModuleTransactor creates a new write-only instance of StakingModule, bound to a specific deployed contract.
func NewStakingModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingModuleTransactor, error) {
	contract, err := bindStakingModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingModuleTransactor{contract: contract}, nil
}

// NewStakingModuleFilterer creates a new log filterer instance of StakingModule, bound to a specific deployed contract.
func NewStakingModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingModuleFilterer, error) {
	contract, err := bindStakingModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingModuleFilterer{contract: contract}, nil
}

// bindStakingModule binds a generic wrapper to an already deployed contract.
func bindStakingModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingModule *StakingModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingModule.Contract.StakingModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingModule *StakingModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingModule.Contract.StakingModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingModule *StakingModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingModule.Contract.StakingModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingModule *StakingModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingModule *StakingModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingModule *StakingModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingModule.Contract.contract.Transact(opts, method, params...)
}

// Delegation is a free data retrieval call binding the contract method 0x241774e6.
//
// Solidity: function delegation(address delegatorAddress, string validatorAddress) view returns(uint256 shares, (uint256,string) balance)
func (_StakingModule *StakingModuleCaller) Delegation(opts *bind.CallOpts, delegatorAddress common.Address, validatorAddress string) (struct {
	Shares  *big.Int
	Balance CosmosCoin
}, error) {
	var out []interface{}
	err := _StakingModule.contract.Call(opts, &out, "delegation", delegatorAddress, validatorAddress)

	outstruct := new(struct {
		Shares  *big.Int
		Balance CosmosCoin
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[1], new(CosmosCoin)).(*CosmosCoin)

	return *outstruct, err

}

// Delegation is a free data retrieval call binding the contract method 0x241774e6.
//
// Solidity: function delegation(address delegatorAddress, string validatorAddress) view returns(uint256 shares, (uint256,string) balance)
func (_StakingModule *StakingModuleSession) Delegation(delegatorAddress common.Address, validatorAddress string) (struct {
	Shares  *big.Int
	Balance CosmosCoin
}, error) {
	return _StakingModule.Contract.Delegation(&_StakingModule.CallOpts, delegatorAddress, validatorAddress)
}

// Delegation is a free data retrieval call binding the contract method 0x241774e6.
//
// Solidity: function delegation(address delegatorAddress, string validatorAddress) view returns(uint256 shares, (uint256,string) balance)
func (_StakingModule *StakingModuleCallerSession) Delegation(delegatorAddress common.Address, validatorAddress string) (struct {
	Shares  *big.Int
	Balance CosmosCoin
}, error) {
	return _StakingModule.Contract.Delegation(&_StakingModule.CallOpts, delegatorAddress, validatorAddress)
}

// Delegate is a paid mutator transaction binding the contract method 0x03f24de1.
//
// Solidity: function delegate(string validatorAddress, uint256 amount) returns(bool success)
func (_StakingModule *StakingModuleTransactor) Delegate(opts *bind.TransactOpts, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "delegate", validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x03f24de1.
//
// Solidity: function delegate(string validatorAddress, uint256 amount) returns(bool success)
func (_StakingModule *StakingModuleSession) Delegate(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Delegate(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x03f24de1.
//
// Solidity: function delegate(string validatorAddress, uint256 amount) returns(bool success)
func (_StakingModule *StakingModuleTransactorSession) Delegate(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Delegate(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x7dd0209d.
//
// Solidity: function redelegate(string validatorSrcAddress, string validatorDstAddress, uint256 amount) returns(bool success)
func (_StakingModule *StakingModuleTransactor) Redelegate(opts *bind.TransactOpts, validatorSrcAddress string, validatorDstAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "redelegate", validatorSrcAddress, validatorDstAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x7dd0209d.
//
// Solidity: function redelegate(string validatorSrcAddress, string validatorDstAddress, uint256 amount) returns(bool success)
func (_StakingModule *StakingModuleSession) Redelegate(validatorSrcAddress string, validatorDstAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Redelegate(&_StakingModule.TransactOpts, validatorSrcAddress, validatorDstAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x7dd0209d.
//
// Solidity: function redelegate(string validatorSrcAddress, string validatorDstAddress, uint256 amount) returns(bool success)
func (_StakingModule *StakingModuleTransactorSession) Redelegate(validatorSrcAddress string, validatorDstAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Redelegate(&_StakingModule.TransactOpts, validatorSrcAddress, validatorDstAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string validatorAddress, uint256 amount) returns(bool success)
func (_StakingModule *StakingModuleTransactor) Undelegate(opts *bind.TransactOpts, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "undelegate", validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string validatorAddress, uint256 amount) returns(bool success)
func (_StakingModule *StakingModuleSession) Undelegate(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Undelegate(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string validatorAddress, uint256 amount) returns(bool success)
func (_StakingModule *StakingModuleTransactorSession) Undelegate(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Undelegate(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// WithdrawDelegatorRewards is a paid mutator transaction binding the contract method 0x6636125e.
//
// Solidity: function withdrawDelegatorRewards(string validatorAddress) returns((uint256,string)[] amount)
func (_StakingModule *StakingModuleTransactor) WithdrawDelegatorRewards(opts *bind.TransactOpts, validatorAddress string) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "withdrawDelegatorRewards", validatorAddress)
}

// WithdrawDelegatorRewards is a paid mutator transaction binding the contract method 0x6636125e.
//
// Solidity: function withdrawDelegatorRewards(string validatorAddress) returns((uint256,string)[] amount)
func (_StakingModule *StakingModuleSession) WithdrawDelegatorRewards(validatorAddress string) (*types.Transaction, error) {
	return _StakingModule.Contract.WithdrawDelegatorRewards(&_StakingModule.TransactOpts, validatorAddress)
}

// WithdrawDelegatorRewards is a paid mutator transaction binding the contract method 0x6636125e.
//
// Solidity: function withdrawDelegatorRewards(string validatorAddress) returns((uint256,string)[] amount)
func (_StakingModule *StakingModuleTransactorSession) WithdrawDelegatorRewards(validatorAddress string) (*types.Transaction, error) {
	return _StakingModule.Contract.WithdrawDelegatorRewards(&_StakingModule.TransactOpts, validatorAddress)
}
