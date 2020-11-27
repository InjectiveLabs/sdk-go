// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HubABI is the input ABI used to generate the binding from.
const HubABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iERC20TokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recipientCosmosAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iERC20TokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromCosmosAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iERC20TokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proposedCosmosCoin\",\"type\":\"string\"}],\"name\":\"TemplateCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iERC20TokenContract\",\"type\":\"address\"}],\"name\":\"UnregisterChild\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"iERC20TokenContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"recipientCosmosAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"childMapping\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"iERC20TokenContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"fromCosmosAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposedCosmosCoin\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"spawnTemplate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"child\",\"type\":\"address\"}],\"name\":\"unregisterChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Hub is an auto generated Go binding around an Ethereum contract.
type Hub struct {
	HubCaller     // Read-only binding to the contract
	HubTransactor // Write-only binding to the contract
	HubFilterer   // Log filterer for contract events
}

// HubCaller is an auto generated read-only Go binding around an Ethereum contract.
type HubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HubSession struct {
	Contract     *Hub              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HubCallerSession struct {
	Contract *HubCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HubTransactorSession struct {
	Contract     *HubTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubRaw is an auto generated low-level Go binding around an Ethereum contract.
type HubRaw struct {
	Contract *Hub // Generic contract binding to access the raw methods on
}

// HubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HubCallerRaw struct {
	Contract *HubCaller // Generic read-only contract binding to access the raw methods on
}

// HubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HubTransactorRaw struct {
	Contract *HubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHub creates a new instance of Hub, bound to a specific deployed contract.
func NewHub(address common.Address, backend bind.ContractBackend) (*Hub, error) {
	contract, err := bindHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hub{HubCaller: HubCaller{contract: contract}, HubTransactor: HubTransactor{contract: contract}, HubFilterer: HubFilterer{contract: contract}}, nil
}

// NewHubCaller creates a new read-only instance of Hub, bound to a specific deployed contract.
func NewHubCaller(address common.Address, caller bind.ContractCaller) (*HubCaller, error) {
	contract, err := bindHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HubCaller{contract: contract}, nil
}

// NewHubTransactor creates a new write-only instance of Hub, bound to a specific deployed contract.
func NewHubTransactor(address common.Address, transactor bind.ContractTransactor) (*HubTransactor, error) {
	contract, err := bindHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HubTransactor{contract: contract}, nil
}

// NewHubFilterer creates a new log filterer instance of Hub, bound to a specific deployed contract.
func NewHubFilterer(address common.Address, filterer bind.ContractFilterer) (*HubFilterer, error) {
	contract, err := bindHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HubFilterer{contract: contract}, nil
}

// bindHub binds a generic wrapper to an already deployed contract.
func bindHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (f *HubRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return f.Contract.HubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (f *HubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.Contract.HubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (f *HubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return f.Contract.HubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (f *HubCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return f.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (f *HubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (f *HubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return f.Contract.contract.Transact(opts, method, params...)
}

// ChildMapping is a free data retrieval call binding the contract method 0xea18b417.
//
// Solidity: function childMapping(address ) view returns(bool)
func (f *HubCaller) ChildMapping(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := f.contract.Call(opts, out, "childMapping", arg0)
	return *ret0, err
}

// ChildMapping is a free data retrieval call binding the contract method 0xea18b417.
//
// Solidity: function childMapping(address ) view returns(bool)
func (f *HubSession) ChildMapping(arg0 common.Address) (bool, error) {
	return f.Contract.ChildMapping(&f.CallOpts, arg0)
}

// ChildMapping is a free data retrieval call binding the contract method 0xea18b417.
//
// Solidity: function childMapping(address ) view returns(bool)
func (f *HubCallerSession) ChildMapping(arg0 common.Address) (bool, error) {
	return f.Contract.ChildMapping(&f.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (f *HubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := f.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (f *HubSession) Owner() (common.Address, error) {
	return f.Contract.Owner(&f.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (f *HubCallerSession) Owner() (common.Address, error) {
	return f.Contract.Owner(&f.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address iERC20TokenContract, string recipientCosmosAddress, uint256 amount) returns()
func (f *HubTransactor) Burn(opts *bind.TransactOpts, iERC20TokenContract common.Address, recipientCosmosAddress string, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "burn", iERC20TokenContract, recipientCosmosAddress, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address iERC20TokenContract, string recipientCosmosAddress, uint256 amount) returns()
func (f *HubSession) Burn(iERC20TokenContract common.Address, recipientCosmosAddress string, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Burn(&f.TransactOpts, iERC20TokenContract, recipientCosmosAddress, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address iERC20TokenContract, string recipientCosmosAddress, uint256 amount) returns()
func (f *HubTransactorSession) Burn(iERC20TokenContract common.Address, recipientCosmosAddress string, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Burn(&f.TransactOpts, iERC20TokenContract, recipientCosmosAddress, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xda14cbbc.
//
// Solidity: function mint(address iERC20TokenContract, string fromCosmosAddress, address recipient, uint256 amount) returns()
func (f *HubTransactor) Mint(opts *bind.TransactOpts, iERC20TokenContract common.Address, fromCosmosAddress string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "mint", iERC20TokenContract, fromCosmosAddress, recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xda14cbbc.
//
// Solidity: function mint(address iERC20TokenContract, string fromCosmosAddress, address recipient, uint256 amount) returns()
func (f *HubSession) Mint(iERC20TokenContract common.Address, fromCosmosAddress string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Mint(&f.TransactOpts, iERC20TokenContract, fromCosmosAddress, recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xda14cbbc.
//
// Solidity: function mint(address iERC20TokenContract, string fromCosmosAddress, address recipient, uint256 amount) returns()
func (f *HubTransactorSession) Mint(iERC20TokenContract common.Address, fromCosmosAddress string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Mint(&f.TransactOpts, iERC20TokenContract, fromCosmosAddress, recipient, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (f *HubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (f *HubSession) RenounceOwnership() (*types.Transaction, error) {
	return f.Contract.RenounceOwnership(&f.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (f *HubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return f.Contract.RenounceOwnership(&f.TransactOpts)
}

// SpawnTemplate is a paid mutator transaction binding the contract method 0xef688fb9.
//
// Solidity: function spawnTemplate(string proposedCosmosCoin, string name, string symbol, uint256 decimals) returns()
func (f *HubTransactor) SpawnTemplate(opts *bind.TransactOpts, proposedCosmosCoin string, name string, symbol string, decimals *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "spawnTemplate", proposedCosmosCoin, name, symbol, decimals)
}

// SpawnTemplate is a paid mutator transaction binding the contract method 0xef688fb9.
//
// Solidity: function spawnTemplate(string proposedCosmosCoin, string name, string symbol, uint256 decimals) returns()
func (f *HubSession) SpawnTemplate(proposedCosmosCoin string, name string, symbol string, decimals *big.Int) (*types.Transaction, error) {
	return f.Contract.SpawnTemplate(&f.TransactOpts, proposedCosmosCoin, name, symbol, decimals)
}

// SpawnTemplate is a paid mutator transaction binding the contract method 0xef688fb9.
//
// Solidity: function spawnTemplate(string proposedCosmosCoin, string name, string symbol, uint256 decimals) returns()
func (f *HubTransactorSession) SpawnTemplate(proposedCosmosCoin string, name string, symbol string, decimals *big.Int) (*types.Transaction, error) {
	return f.Contract.SpawnTemplate(&f.TransactOpts, proposedCosmosCoin, name, symbol, decimals)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (f *HubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return f.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (f *HubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return f.Contract.TransferOwnership(&f.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (f *HubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return f.Contract.TransferOwnership(&f.TransactOpts, newOwner)
}

// UnregisterChild is a paid mutator transaction binding the contract method 0xe3db630c.
//
// Solidity: function unregisterChild(address child) returns()
func (f *HubTransactor) UnregisterChild(opts *bind.TransactOpts, child common.Address) (*types.Transaction, error) {
	return f.contract.Transact(opts, "unregisterChild", child)
}

// UnregisterChild is a paid mutator transaction binding the contract method 0xe3db630c.
//
// Solidity: function unregisterChild(address child) returns()
func (f *HubSession) UnregisterChild(child common.Address) (*types.Transaction, error) {
	return f.Contract.UnregisterChild(&f.TransactOpts, child)
}

// UnregisterChild is a paid mutator transaction binding the contract method 0xe3db630c.
//
// Solidity: function unregisterChild(address child) returns()
func (f *HubTransactorSession) UnregisterChild(child common.Address) (*types.Transaction, error) {
	return f.Contract.UnregisterChild(&f.TransactOpts, child)
}

// HubBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Hub contract.
type HubBurnIterator struct {
	Event *HubBurn // Event containing the contract specifics and raw log

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
func (it *HubBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubBurn)
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
		it.Event = new(HubBurn)
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
func (it *HubBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubBurn represents a Burn event raised by the Hub contract.
type HubBurn struct {
	IERC20TokenContract    common.Address
	Target                 common.Address
	RecipientCosmosAddress string
	Amount                 *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x2d1cbee916b310b4f96bc080243ebebf4daa27a66b334773cf5d77d12226f4cd.
//
// Solidity: event Burn(address iERC20TokenContract, address target, string recipientCosmosAddress, uint256 amount)
func (f *HubFilterer) FilterBurn(opts *bind.FilterOpts) (*HubBurnIterator, error) {

	logs, sub, err := f.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &HubBurnIterator{contract: f.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x2d1cbee916b310b4f96bc080243ebebf4daa27a66b334773cf5d77d12226f4cd.
//
// Solidity: event Burn(address iERC20TokenContract, address target, string recipientCosmosAddress, uint256 amount)
func (f *HubFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *HubBurn) (event.Subscription, error) {

	logs, sub, err := f.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubBurn)
				if err := f.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x2d1cbee916b310b4f96bc080243ebebf4daa27a66b334773cf5d77d12226f4cd.
//
// Solidity: event Burn(address iERC20TokenContract, address target, string recipientCosmosAddress, uint256 amount)
func (f *HubFilterer) ParseBurn(log types.Log) (*HubBurn, error) {
	event := new(HubBurn)
	if err := f.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HubMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Hub contract.
type HubMintIterator struct {
	Event *HubMint // Event containing the contract specifics and raw log

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
func (it *HubMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubMint)
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
		it.Event = new(HubMint)
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
func (it *HubMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubMint represents a Mint event raised by the Hub contract.
type HubMint struct {
	IERC20TokenContract common.Address
	FromCosmosAddress   string
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x85e1b778603e7ba009b4b23a5909ddb0ba4e73f01838cb2029e066158fd62f74.
//
// Solidity: event Mint(address iERC20TokenContract, string fromCosmosAddress, address recipient, uint256 amount)
func (f *HubFilterer) FilterMint(opts *bind.FilterOpts) (*HubMintIterator, error) {

	logs, sub, err := f.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &HubMintIterator{contract: f.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x85e1b778603e7ba009b4b23a5909ddb0ba4e73f01838cb2029e066158fd62f74.
//
// Solidity: event Mint(address iERC20TokenContract, string fromCosmosAddress, address recipient, uint256 amount)
func (f *HubFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *HubMint) (event.Subscription, error) {

	logs, sub, err := f.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubMint)
				if err := f.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x85e1b778603e7ba009b4b23a5909ddb0ba4e73f01838cb2029e066158fd62f74.
//
// Solidity: event Mint(address iERC20TokenContract, string fromCosmosAddress, address recipient, uint256 amount)
func (f *HubFilterer) ParseMint(log types.Log) (*HubMint, error) {
	event := new(HubMint)
	if err := f.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hub contract.
type HubOwnershipTransferredIterator struct {
	Event *HubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubOwnershipTransferred)
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
		it.Event = new(HubOwnershipTransferred)
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
func (it *HubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubOwnershipTransferred represents a OwnershipTransferred event raised by the Hub contract.
type HubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (f *HubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HubOwnershipTransferredIterator{contract: f.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (f *HubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubOwnershipTransferred)
				if err := f.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (f *HubFilterer) ParseOwnershipTransferred(log types.Log) (*HubOwnershipTransferred, error) {
	event := new(HubOwnershipTransferred)
	if err := f.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HubTemplateCreationIterator is returned from FilterTemplateCreation and is used to iterate over the raw logs and unpacked data for TemplateCreation events raised by the Hub contract.
type HubTemplateCreationIterator struct {
	Event *HubTemplateCreation // Event containing the contract specifics and raw log

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
func (it *HubTemplateCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubTemplateCreation)
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
		it.Event = new(HubTemplateCreation)
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
func (it *HubTemplateCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubTemplateCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubTemplateCreation represents a TemplateCreation event raised by the Hub contract.
type HubTemplateCreation struct {
	IERC20TokenContract common.Address
	ProposedCosmosCoin  string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTemplateCreation is a free log retrieval operation binding the contract event 0xa2c3e18d8cd0d905a1f440e7daccb30a2c440da0f9bc44f81a69b6f21933f4fc.
//
// Solidity: event TemplateCreation(address iERC20TokenContract, string proposedCosmosCoin)
func (f *HubFilterer) FilterTemplateCreation(opts *bind.FilterOpts) (*HubTemplateCreationIterator, error) {

	logs, sub, err := f.contract.FilterLogs(opts, "TemplateCreation")
	if err != nil {
		return nil, err
	}
	return &HubTemplateCreationIterator{contract: f.contract, event: "TemplateCreation", logs: logs, sub: sub}, nil
}

// WatchTemplateCreation is a free log subscription operation binding the contract event 0xa2c3e18d8cd0d905a1f440e7daccb30a2c440da0f9bc44f81a69b6f21933f4fc.
//
// Solidity: event TemplateCreation(address iERC20TokenContract, string proposedCosmosCoin)
func (f *HubFilterer) WatchTemplateCreation(opts *bind.WatchOpts, sink chan<- *HubTemplateCreation) (event.Subscription, error) {

	logs, sub, err := f.contract.WatchLogs(opts, "TemplateCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubTemplateCreation)
				if err := f.contract.UnpackLog(event, "TemplateCreation", log); err != nil {
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

// ParseTemplateCreation is a log parse operation binding the contract event 0xa2c3e18d8cd0d905a1f440e7daccb30a2c440da0f9bc44f81a69b6f21933f4fc.
//
// Solidity: event TemplateCreation(address iERC20TokenContract, string proposedCosmosCoin)
func (f *HubFilterer) ParseTemplateCreation(log types.Log) (*HubTemplateCreation, error) {
	event := new(HubTemplateCreation)
	if err := f.contract.UnpackLog(event, "TemplateCreation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HubUnregisterChildIterator is returned from FilterUnregisterChild and is used to iterate over the raw logs and unpacked data for UnregisterChild events raised by the Hub contract.
type HubUnregisterChildIterator struct {
	Event *HubUnregisterChild // Event containing the contract specifics and raw log

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
func (it *HubUnregisterChildIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubUnregisterChild)
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
		it.Event = new(HubUnregisterChild)
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
func (it *HubUnregisterChildIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubUnregisterChildIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubUnregisterChild represents a UnregisterChild event raised by the Hub contract.
type HubUnregisterChild struct {
	IERC20TokenContract common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUnregisterChild is a free log retrieval operation binding the contract event 0x9a0579b1a8a2e82095ed4ab2e299c227ab4fcd5fa467d94a576012e39cd4925a.
//
// Solidity: event UnregisterChild(address iERC20TokenContract)
func (f *HubFilterer) FilterUnregisterChild(opts *bind.FilterOpts) (*HubUnregisterChildIterator, error) {

	logs, sub, err := f.contract.FilterLogs(opts, "UnregisterChild")
	if err != nil {
		return nil, err
	}
	return &HubUnregisterChildIterator{contract: f.contract, event: "UnregisterChild", logs: logs, sub: sub}, nil
}

// WatchUnregisterChild is a free log subscription operation binding the contract event 0x9a0579b1a8a2e82095ed4ab2e299c227ab4fcd5fa467d94a576012e39cd4925a.
//
// Solidity: event UnregisterChild(address iERC20TokenContract)
func (f *HubFilterer) WatchUnregisterChild(opts *bind.WatchOpts, sink chan<- *HubUnregisterChild) (event.Subscription, error) {

	logs, sub, err := f.contract.WatchLogs(opts, "UnregisterChild")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubUnregisterChild)
				if err := f.contract.UnpackLog(event, "UnregisterChild", log); err != nil {
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

// ParseUnregisterChild is a log parse operation binding the contract event 0x9a0579b1a8a2e82095ed4ab2e299c227ab4fcd5fa467d94a576012e39cd4925a.
//
// Solidity: event UnregisterChild(address iERC20TokenContract)
func (f *HubFilterer) ParseUnregisterChild(log types.Log) (*HubUnregisterChild, error) {
	event := new(HubUnregisterChild)
	if err := f.contract.UnpackLog(event, "UnregisterChild", log); err != nil {
		return nil, err
	}
	return event, nil
}
