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

// PriceFeederABI is the input ABI used to generate the binding from.
const PriceFeederABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oraclePriceFeeder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidatorTester\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RegisterFuturesMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RegisterPerpetualMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketExpirationTimestamp\",\"type\":\"uint256\"}],\"name\":\"SetSettlementPrice\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_FEEDER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"currentPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"futuresSettlementPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"futuresSettlementTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isRegisteredMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"marketFundingIntervals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"registerPerpetualMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"registerFuturesMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getFundingInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setExpiryFuturesSettlementPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getExpiryFuturesSettlementPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PriceFeeder is an auto generated Go binding around an Ethereum contract.
type PriceFeeder struct {
	PriceFeederCaller     // Read-only binding to the contract
	PriceFeederTransactor // Write-only binding to the contract
	PriceFeederFilterer   // Log filterer for contract events
}

// PriceFeederCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeederCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeederTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeederTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeederFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeederFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeederSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeederSession struct {
	Contract     *PriceFeeder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeederCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeederCallerSession struct {
	Contract *PriceFeederCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceFeederTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeederTransactorSession struct {
	Contract     *PriceFeederTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceFeederRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeederRaw struct {
	Contract *PriceFeeder // Generic contract binding to access the raw methods on
}

// PriceFeederCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeederCallerRaw struct {
	Contract *PriceFeederCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeederTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeederTransactorRaw struct {
	Contract *PriceFeederTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeeder creates a new instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeeder(address common.Address, backend bind.ContractBackend) (*PriceFeeder, error) {
	contract, err := bindPriceFeeder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeeder{PriceFeederCaller: PriceFeederCaller{contract: contract}, PriceFeederTransactor: PriceFeederTransactor{contract: contract}, PriceFeederFilterer: PriceFeederFilterer{contract: contract}}, nil
}

// NewPriceFeederCaller creates a new read-only instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeederCaller(address common.Address, caller bind.ContractCaller) (*PriceFeederCaller, error) {
	contract, err := bindPriceFeeder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeederCaller{contract: contract}, nil
}

// NewPriceFeederTransactor creates a new write-only instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeederTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeederTransactor, error) {
	contract, err := bindPriceFeeder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeederTransactor{contract: contract}, nil
}

// NewPriceFeederFilterer creates a new log filterer instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeederFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeederFilterer, error) {
	contract, err := bindPriceFeeder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeederFilterer{contract: contract}, nil
}

// bindPriceFeeder binds a generic wrapper to an already deployed contract.
func bindPriceFeeder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeederABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (f *PriceFeederRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return f.Contract.PriceFeederCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (f *PriceFeederRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.Contract.PriceFeederTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (f *PriceFeederRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return f.Contract.PriceFeederTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (f *PriceFeederCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return f.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (f *PriceFeederTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (f *PriceFeederTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return f.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (f *PriceFeederCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (f *PriceFeederSession) DEFAULTADMINROLE() ([32]byte, error) {
	return f.Contract.DEFAULTADMINROLE(&f.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (f *PriceFeederCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return f.Contract.DEFAULTADMINROLE(&f.CallOpts)
}

// PRICEFEEDER is a free data retrieval call binding the contract method 0x882ed671.
//
// Solidity: function PRICE_FEEDER() view returns(bytes32)
func (f *PriceFeederCaller) PRICEFEEDER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "PRICE_FEEDER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRICEFEEDER is a free data retrieval call binding the contract method 0x882ed671.
//
// Solidity: function PRICE_FEEDER() view returns(bytes32)
func (f *PriceFeederSession) PRICEFEEDER() ([32]byte, error) {
	return f.Contract.PRICEFEEDER(&f.CallOpts)
}

// PRICEFEEDER is a free data retrieval call binding the contract method 0x882ed671.
//
// Solidity: function PRICE_FEEDER() view returns(bytes32)
func (f *PriceFeederCallerSession) PRICEFEEDER() ([32]byte, error) {
	return f.Contract.PRICEFEEDER(&f.CallOpts)
}

// CurrentPrices is a free data retrieval call binding the contract method 0xacd0ecaf.
//
// Solidity: function currentPrices(bytes32 ) view returns(uint256)
func (f *PriceFeederCaller) CurrentPrices(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "currentPrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPrices is a free data retrieval call binding the contract method 0xacd0ecaf.
//
// Solidity: function currentPrices(bytes32 ) view returns(uint256)
func (f *PriceFeederSession) CurrentPrices(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.CurrentPrices(&f.CallOpts, arg0)
}

// CurrentPrices is a free data retrieval call binding the contract method 0xacd0ecaf.
//
// Solidity: function currentPrices(bytes32 ) view returns(uint256)
func (f *PriceFeederCallerSession) CurrentPrices(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.CurrentPrices(&f.CallOpts, arg0)
}

// FuturesSettlementPrice is a free data retrieval call binding the contract method 0xc60f6ef5.
//
// Solidity: function futuresSettlementPrice(bytes32 ) view returns(uint256)
func (f *PriceFeederCaller) FuturesSettlementPrice(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "futuresSettlementPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FuturesSettlementPrice is a free data retrieval call binding the contract method 0xc60f6ef5.
//
// Solidity: function futuresSettlementPrice(bytes32 ) view returns(uint256)
func (f *PriceFeederSession) FuturesSettlementPrice(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.FuturesSettlementPrice(&f.CallOpts, arg0)
}

// FuturesSettlementPrice is a free data retrieval call binding the contract method 0xc60f6ef5.
//
// Solidity: function futuresSettlementPrice(bytes32 ) view returns(uint256)
func (f *PriceFeederCallerSession) FuturesSettlementPrice(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.FuturesSettlementPrice(&f.CallOpts, arg0)
}

// FuturesSettlementTimes is a free data retrieval call binding the contract method 0x470b2d0e.
//
// Solidity: function futuresSettlementTimes(bytes32 ) view returns(uint256)
func (f *PriceFeederCaller) FuturesSettlementTimes(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "futuresSettlementTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FuturesSettlementTimes is a free data retrieval call binding the contract method 0x470b2d0e.
//
// Solidity: function futuresSettlementTimes(bytes32 ) view returns(uint256)
func (f *PriceFeederSession) FuturesSettlementTimes(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.FuturesSettlementTimes(&f.CallOpts, arg0)
}

// FuturesSettlementTimes is a free data retrieval call binding the contract method 0x470b2d0e.
//
// Solidity: function futuresSettlementTimes(bytes32 ) view returns(uint256)
func (f *PriceFeederCallerSession) FuturesSettlementTimes(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.FuturesSettlementTimes(&f.CallOpts, arg0)
}

// GetExpiryFuturesSettlementPrice is a free data retrieval call binding the contract method 0x57ab828a.
//
// Solidity: function getExpiryFuturesSettlementPrice(bytes32 marketID) view returns(uint256)
func (f *PriceFeederCaller) GetExpiryFuturesSettlementPrice(opts *bind.CallOpts, marketID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getExpiryFuturesSettlementPrice", marketID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpiryFuturesSettlementPrice is a free data retrieval call binding the contract method 0x57ab828a.
//
// Solidity: function getExpiryFuturesSettlementPrice(bytes32 marketID) view returns(uint256)
func (f *PriceFeederSession) GetExpiryFuturesSettlementPrice(marketID [32]byte) (*big.Int, error) {
	return f.Contract.GetExpiryFuturesSettlementPrice(&f.CallOpts, marketID)
}

// GetExpiryFuturesSettlementPrice is a free data retrieval call binding the contract method 0x57ab828a.
//
// Solidity: function getExpiryFuturesSettlementPrice(bytes32 marketID) view returns(uint256)
func (f *PriceFeederCallerSession) GetExpiryFuturesSettlementPrice(marketID [32]byte) (*big.Int, error) {
	return f.Contract.GetExpiryFuturesSettlementPrice(&f.CallOpts, marketID)
}

// GetFundingInterval is a free data retrieval call binding the contract method 0xaa7566de.
//
// Solidity: function getFundingInterval(bytes32 marketID) view returns(uint256)
func (f *PriceFeederCaller) GetFundingInterval(opts *bind.CallOpts, marketID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getFundingInterval", marketID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingInterval is a free data retrieval call binding the contract method 0xaa7566de.
//
// Solidity: function getFundingInterval(bytes32 marketID) view returns(uint256)
func (f *PriceFeederSession) GetFundingInterval(marketID [32]byte) (*big.Int, error) {
	return f.Contract.GetFundingInterval(&f.CallOpts, marketID)
}

// GetFundingInterval is a free data retrieval call binding the contract method 0xaa7566de.
//
// Solidity: function getFundingInterval(bytes32 marketID) view returns(uint256)
func (f *PriceFeederCallerSession) GetFundingInterval(marketID [32]byte) (*big.Int, error) {
	return f.Contract.GetFundingInterval(&f.CallOpts, marketID)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 marketID) view returns(uint256)
func (f *PriceFeederCaller) GetPrice(opts *bind.CallOpts, marketID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getPrice", marketID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 marketID) view returns(uint256)
func (f *PriceFeederSession) GetPrice(marketID [32]byte) (*big.Int, error) {
	return f.Contract.GetPrice(&f.CallOpts, marketID)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 marketID) view returns(uint256)
func (f *PriceFeederCallerSession) GetPrice(marketID [32]byte) (*big.Int, error) {
	return f.Contract.GetPrice(&f.CallOpts, marketID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (f *PriceFeederCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (f *PriceFeederSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return f.Contract.GetRoleAdmin(&f.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (f *PriceFeederCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return f.Contract.GetRoleAdmin(&f.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (f *PriceFeederCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (f *PriceFeederSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return f.Contract.GetRoleMember(&f.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (f *PriceFeederCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return f.Contract.GetRoleMember(&f.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (f *PriceFeederCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (f *PriceFeederSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return f.Contract.GetRoleMemberCount(&f.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (f *PriceFeederCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return f.Contract.GetRoleMemberCount(&f.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (f *PriceFeederCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (f *PriceFeederSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return f.Contract.HasRole(&f.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (f *PriceFeederCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return f.Contract.HasRole(&f.CallOpts, role, account)
}

// IsRegisteredMarket is a free data retrieval call binding the contract method 0x17eab5d6.
//
// Solidity: function isRegisteredMarket(bytes32 ) view returns(bool)
func (f *PriceFeederCaller) IsRegisteredMarket(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isRegisteredMarket", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredMarket is a free data retrieval call binding the contract method 0x17eab5d6.
//
// Solidity: function isRegisteredMarket(bytes32 ) view returns(bool)
func (f *PriceFeederSession) IsRegisteredMarket(arg0 [32]byte) (bool, error) {
	return f.Contract.IsRegisteredMarket(&f.CallOpts, arg0)
}

// IsRegisteredMarket is a free data retrieval call binding the contract method 0x17eab5d6.
//
// Solidity: function isRegisteredMarket(bytes32 ) view returns(bool)
func (f *PriceFeederCallerSession) IsRegisteredMarket(arg0 [32]byte) (bool, error) {
	return f.Contract.IsRegisteredMarket(&f.CallOpts, arg0)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (f *PriceFeederCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "marketCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (f *PriceFeederSession) MarketCount() (*big.Int, error) {
	return f.Contract.MarketCount(&f.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (f *PriceFeederCallerSession) MarketCount() (*big.Int, error) {
	return f.Contract.MarketCount(&f.CallOpts)
}

// MarketFundingIntervals is a free data retrieval call binding the contract method 0x19cb625c.
//
// Solidity: function marketFundingIntervals(bytes32 ) view returns(uint256)
func (f *PriceFeederCaller) MarketFundingIntervals(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "marketFundingIntervals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketFundingIntervals is a free data retrieval call binding the contract method 0x19cb625c.
//
// Solidity: function marketFundingIntervals(bytes32 ) view returns(uint256)
func (f *PriceFeederSession) MarketFundingIntervals(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MarketFundingIntervals(&f.CallOpts, arg0)
}

// MarketFundingIntervals is a free data retrieval call binding the contract method 0x19cb625c.
//
// Solidity: function marketFundingIntervals(bytes32 ) view returns(uint256)
func (f *PriceFeederCallerSession) MarketFundingIntervals(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MarketFundingIntervals(&f.CallOpts, arg0)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (f *PriceFeederTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return f.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (f *PriceFeederSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return f.Contract.GrantRole(&f.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (f *PriceFeederTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return f.Contract.GrantRole(&f.TransactOpts, role, account)
}

// RegisterFuturesMarket is a paid mutator transaction binding the contract method 0xa4716060.
//
// Solidity: function registerFuturesMarket(bytes32 marketID, uint256 expirationTime, uint256 price) returns()
func (f *PriceFeederTransactor) RegisterFuturesMarket(opts *bind.TransactOpts, marketID [32]byte, expirationTime *big.Int, price *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "registerFuturesMarket", marketID, expirationTime, price)
}

// RegisterFuturesMarket is a paid mutator transaction binding the contract method 0xa4716060.
//
// Solidity: function registerFuturesMarket(bytes32 marketID, uint256 expirationTime, uint256 price) returns()
func (f *PriceFeederSession) RegisterFuturesMarket(marketID [32]byte, expirationTime *big.Int, price *big.Int) (*types.Transaction, error) {
	return f.Contract.RegisterFuturesMarket(&f.TransactOpts, marketID, expirationTime, price)
}

// RegisterFuturesMarket is a paid mutator transaction binding the contract method 0xa4716060.
//
// Solidity: function registerFuturesMarket(bytes32 marketID, uint256 expirationTime, uint256 price) returns()
func (f *PriceFeederTransactorSession) RegisterFuturesMarket(marketID [32]byte, expirationTime *big.Int, price *big.Int) (*types.Transaction, error) {
	return f.Contract.RegisterFuturesMarket(&f.TransactOpts, marketID, expirationTime, price)
}

// RegisterPerpetualMarket is a paid mutator transaction binding the contract method 0x2be3ccd9.
//
// Solidity: function registerPerpetualMarket(bytes32 marketID, uint256 fundingInterval, uint256 price) returns()
func (f *PriceFeederTransactor) RegisterPerpetualMarket(opts *bind.TransactOpts, marketID [32]byte, fundingInterval *big.Int, price *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "registerPerpetualMarket", marketID, fundingInterval, price)
}

// RegisterPerpetualMarket is a paid mutator transaction binding the contract method 0x2be3ccd9.
//
// Solidity: function registerPerpetualMarket(bytes32 marketID, uint256 fundingInterval, uint256 price) returns()
func (f *PriceFeederSession) RegisterPerpetualMarket(marketID [32]byte, fundingInterval *big.Int, price *big.Int) (*types.Transaction, error) {
	return f.Contract.RegisterPerpetualMarket(&f.TransactOpts, marketID, fundingInterval, price)
}

// RegisterPerpetualMarket is a paid mutator transaction binding the contract method 0x2be3ccd9.
//
// Solidity: function registerPerpetualMarket(bytes32 marketID, uint256 fundingInterval, uint256 price) returns()
func (f *PriceFeederTransactorSession) RegisterPerpetualMarket(marketID [32]byte, fundingInterval *big.Int, price *big.Int) (*types.Transaction, error) {
	return f.Contract.RegisterPerpetualMarket(&f.TransactOpts, marketID, fundingInterval, price)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (f *PriceFeederTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return f.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (f *PriceFeederSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return f.Contract.RenounceRole(&f.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (f *PriceFeederTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return f.Contract.RenounceRole(&f.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (f *PriceFeederTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return f.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (f *PriceFeederSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return f.Contract.RevokeRole(&f.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (f *PriceFeederTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return f.Contract.RevokeRole(&f.TransactOpts, role, account)
}

// SetExpiryFuturesSettlementPrice is a paid mutator transaction binding the contract method 0xec6def6a.
//
// Solidity: function setExpiryFuturesSettlementPrice(bytes32 marketID, uint256 price) returns()
func (f *PriceFeederTransactor) SetExpiryFuturesSettlementPrice(opts *bind.TransactOpts, marketID [32]byte, price *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "setExpiryFuturesSettlementPrice", marketID, price)
}

// SetExpiryFuturesSettlementPrice is a paid mutator transaction binding the contract method 0xec6def6a.
//
// Solidity: function setExpiryFuturesSettlementPrice(bytes32 marketID, uint256 price) returns()
func (f *PriceFeederSession) SetExpiryFuturesSettlementPrice(marketID [32]byte, price *big.Int) (*types.Transaction, error) {
	return f.Contract.SetExpiryFuturesSettlementPrice(&f.TransactOpts, marketID, price)
}

// SetExpiryFuturesSettlementPrice is a paid mutator transaction binding the contract method 0xec6def6a.
//
// Solidity: function setExpiryFuturesSettlementPrice(bytes32 marketID, uint256 price) returns()
func (f *PriceFeederTransactorSession) SetExpiryFuturesSettlementPrice(marketID [32]byte, price *big.Int) (*types.Transaction, error) {
	return f.Contract.SetExpiryFuturesSettlementPrice(&f.TransactOpts, marketID, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x10d8d74d.
//
// Solidity: function setPrice(bytes32 marketID, uint256 price) returns()
func (f *PriceFeederTransactor) SetPrice(opts *bind.TransactOpts, marketID [32]byte, price *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "setPrice", marketID, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x10d8d74d.
//
// Solidity: function setPrice(bytes32 marketID, uint256 price) returns()
func (f *PriceFeederSession) SetPrice(marketID [32]byte, price *big.Int) (*types.Transaction, error) {
	return f.Contract.SetPrice(&f.TransactOpts, marketID, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x10d8d74d.
//
// Solidity: function setPrice(bytes32 marketID, uint256 price) returns()
func (f *PriceFeederTransactorSession) SetPrice(marketID [32]byte, price *big.Int) (*types.Transaction, error) {
	return f.Contract.SetPrice(&f.TransactOpts, marketID, price)
}

// PriceFeederRegisterFuturesMarketIterator is returned from FilterRegisterFuturesMarket and is used to iterate over the raw logs and unpacked data for RegisterFuturesMarket events raised by the PriceFeeder contract.
type PriceFeederRegisterFuturesMarketIterator struct {
	Event *PriceFeederRegisterFuturesMarket // Event containing the contract specifics and raw log

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
func (it *PriceFeederRegisterFuturesMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederRegisterFuturesMarket)
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
		it.Event = new(PriceFeederRegisterFuturesMarket)
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
func (it *PriceFeederRegisterFuturesMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederRegisterFuturesMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederRegisterFuturesMarket represents a RegisterFuturesMarket event raised by the PriceFeeder contract.
type PriceFeederRegisterFuturesMarket struct {
	MarketID       [32]byte
	ExpirationTime *big.Int
	InitialPrice   *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRegisterFuturesMarket is a free log retrieval operation binding the contract event 0x658829bdba490af5d973a4e5cd6a98ec9eccc538d005c267a710a7eb62602c5f.
//
// Solidity: event RegisterFuturesMarket(bytes32 indexed marketID, uint256 expirationTime, uint256 initialPrice, uint256 timestamp)
func (f *PriceFeederFilterer) FilterRegisterFuturesMarket(opts *bind.FilterOpts, marketID [][32]byte) (*PriceFeederRegisterFuturesMarketIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "RegisterFuturesMarket", marketIDRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeederRegisterFuturesMarketIterator{contract: f.contract, event: "RegisterFuturesMarket", logs: logs, sub: sub}, nil
}

// WatchRegisterFuturesMarket is a free log subscription operation binding the contract event 0x658829bdba490af5d973a4e5cd6a98ec9eccc538d005c267a710a7eb62602c5f.
//
// Solidity: event RegisterFuturesMarket(bytes32 indexed marketID, uint256 expirationTime, uint256 initialPrice, uint256 timestamp)
func (f *PriceFeederFilterer) WatchRegisterFuturesMarket(opts *bind.WatchOpts, sink chan<- *PriceFeederRegisterFuturesMarket, marketID [][32]byte) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "RegisterFuturesMarket", marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederRegisterFuturesMarket)
				if err := f.contract.UnpackLog(event, "RegisterFuturesMarket", log); err != nil {
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

// ParseRegisterFuturesMarket is a log parse operation binding the contract event 0x658829bdba490af5d973a4e5cd6a98ec9eccc538d005c267a710a7eb62602c5f.
//
// Solidity: event RegisterFuturesMarket(bytes32 indexed marketID, uint256 expirationTime, uint256 initialPrice, uint256 timestamp)
func (f *PriceFeederFilterer) ParseRegisterFuturesMarket(log types.Log) (*PriceFeederRegisterFuturesMarket, error) {
	event := new(PriceFeederRegisterFuturesMarket)
	if err := f.contract.UnpackLog(event, "RegisterFuturesMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeederRegisterPerpetualMarketIterator is returned from FilterRegisterPerpetualMarket and is used to iterate over the raw logs and unpacked data for RegisterPerpetualMarket events raised by the PriceFeeder contract.
type PriceFeederRegisterPerpetualMarketIterator struct {
	Event *PriceFeederRegisterPerpetualMarket // Event containing the contract specifics and raw log

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
func (it *PriceFeederRegisterPerpetualMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederRegisterPerpetualMarket)
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
		it.Event = new(PriceFeederRegisterPerpetualMarket)
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
func (it *PriceFeederRegisterPerpetualMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederRegisterPerpetualMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederRegisterPerpetualMarket represents a RegisterPerpetualMarket event raised by the PriceFeeder contract.
type PriceFeederRegisterPerpetualMarket struct {
	MarketID        [32]byte
	FundingInterval *big.Int
	InitialPrice    *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegisterPerpetualMarket is a free log retrieval operation binding the contract event 0xf4a892d94307acfa4ef3f684cd3bdabf189017b03f42210f4937a4ba2a4edf8c.
//
// Solidity: event RegisterPerpetualMarket(bytes32 indexed marketID, uint256 fundingInterval, uint256 initialPrice, uint256 timestamp)
func (f *PriceFeederFilterer) FilterRegisterPerpetualMarket(opts *bind.FilterOpts, marketID [][32]byte) (*PriceFeederRegisterPerpetualMarketIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "RegisterPerpetualMarket", marketIDRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeederRegisterPerpetualMarketIterator{contract: f.contract, event: "RegisterPerpetualMarket", logs: logs, sub: sub}, nil
}

// WatchRegisterPerpetualMarket is a free log subscription operation binding the contract event 0xf4a892d94307acfa4ef3f684cd3bdabf189017b03f42210f4937a4ba2a4edf8c.
//
// Solidity: event RegisterPerpetualMarket(bytes32 indexed marketID, uint256 fundingInterval, uint256 initialPrice, uint256 timestamp)
func (f *PriceFeederFilterer) WatchRegisterPerpetualMarket(opts *bind.WatchOpts, sink chan<- *PriceFeederRegisterPerpetualMarket, marketID [][32]byte) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "RegisterPerpetualMarket", marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederRegisterPerpetualMarket)
				if err := f.contract.UnpackLog(event, "RegisterPerpetualMarket", log); err != nil {
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

// ParseRegisterPerpetualMarket is a log parse operation binding the contract event 0xf4a892d94307acfa4ef3f684cd3bdabf189017b03f42210f4937a4ba2a4edf8c.
//
// Solidity: event RegisterPerpetualMarket(bytes32 indexed marketID, uint256 fundingInterval, uint256 initialPrice, uint256 timestamp)
func (f *PriceFeederFilterer) ParseRegisterPerpetualMarket(log types.Log) (*PriceFeederRegisterPerpetualMarket, error) {
	event := new(PriceFeederRegisterPerpetualMarket)
	if err := f.contract.UnpackLog(event, "RegisterPerpetualMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeederRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PriceFeeder contract.
type PriceFeederRoleAdminChangedIterator struct {
	Event *PriceFeederRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeederRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederRoleAdminChanged)
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
		it.Event = new(PriceFeederRoleAdminChanged)
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
func (it *PriceFeederRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederRoleAdminChanged represents a RoleAdminChanged event raised by the PriceFeeder contract.
type PriceFeederRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (f *PriceFeederFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PriceFeederRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeederRoleAdminChangedIterator{contract: f.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (f *PriceFeederFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PriceFeederRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederRoleAdminChanged)
				if err := f.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (f *PriceFeederFilterer) ParseRoleAdminChanged(log types.Log) (*PriceFeederRoleAdminChanged, error) {
	event := new(PriceFeederRoleAdminChanged)
	if err := f.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeederRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PriceFeeder contract.
type PriceFeederRoleGrantedIterator struct {
	Event *PriceFeederRoleGranted // Event containing the contract specifics and raw log

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
func (it *PriceFeederRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederRoleGranted)
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
		it.Event = new(PriceFeederRoleGranted)
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
func (it *PriceFeederRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederRoleGranted represents a RoleGranted event raised by the PriceFeeder contract.
type PriceFeederRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (f *PriceFeederFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PriceFeederRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeederRoleGrantedIterator{contract: f.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (f *PriceFeederFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PriceFeederRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederRoleGranted)
				if err := f.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (f *PriceFeederFilterer) ParseRoleGranted(log types.Log) (*PriceFeederRoleGranted, error) {
	event := new(PriceFeederRoleGranted)
	if err := f.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeederRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PriceFeeder contract.
type PriceFeederRoleRevokedIterator struct {
	Event *PriceFeederRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PriceFeederRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederRoleRevoked)
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
		it.Event = new(PriceFeederRoleRevoked)
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
func (it *PriceFeederRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederRoleRevoked represents a RoleRevoked event raised by the PriceFeeder contract.
type PriceFeederRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (f *PriceFeederFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PriceFeederRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeederRoleRevokedIterator{contract: f.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (f *PriceFeederFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PriceFeederRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederRoleRevoked)
				if err := f.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (f *PriceFeederFilterer) ParseRoleRevoked(log types.Log) (*PriceFeederRoleRevoked, error) {
	event := new(PriceFeederRoleRevoked)
	if err := f.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeederSetPriceIterator is returned from FilterSetPrice and is used to iterate over the raw logs and unpacked data for SetPrice events raised by the PriceFeeder contract.
type PriceFeederSetPriceIterator struct {
	Event *PriceFeederSetPrice // Event containing the contract specifics and raw log

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
func (it *PriceFeederSetPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederSetPrice)
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
		it.Event = new(PriceFeederSetPrice)
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
func (it *PriceFeederSetPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederSetPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederSetPrice represents a SetPrice event raised by the PriceFeeder contract.
type PriceFeederSetPrice struct {
	MarketID  [32]byte
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetPrice is a free log retrieval operation binding the contract event 0xabe0d897903d7a41d97b5dfbbb279a7175e68631bb9862bc4b3a1a209a855ac6.
//
// Solidity: event SetPrice(bytes32 indexed marketID, uint256 price, uint256 timestamp)
func (f *PriceFeederFilterer) FilterSetPrice(opts *bind.FilterOpts, marketID [][32]byte) (*PriceFeederSetPriceIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "SetPrice", marketIDRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeederSetPriceIterator{contract: f.contract, event: "SetPrice", logs: logs, sub: sub}, nil
}

// WatchSetPrice is a free log subscription operation binding the contract event 0xabe0d897903d7a41d97b5dfbbb279a7175e68631bb9862bc4b3a1a209a855ac6.
//
// Solidity: event SetPrice(bytes32 indexed marketID, uint256 price, uint256 timestamp)
func (f *PriceFeederFilterer) WatchSetPrice(opts *bind.WatchOpts, sink chan<- *PriceFeederSetPrice, marketID [][32]byte) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "SetPrice", marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederSetPrice)
				if err := f.contract.UnpackLog(event, "SetPrice", log); err != nil {
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

// ParseSetPrice is a log parse operation binding the contract event 0xabe0d897903d7a41d97b5dfbbb279a7175e68631bb9862bc4b3a1a209a855ac6.
//
// Solidity: event SetPrice(bytes32 indexed marketID, uint256 price, uint256 timestamp)
func (f *PriceFeederFilterer) ParseSetPrice(log types.Log) (*PriceFeederSetPrice, error) {
	event := new(PriceFeederSetPrice)
	if err := f.contract.UnpackLog(event, "SetPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeederSetSettlementPriceIterator is returned from FilterSetSettlementPrice and is used to iterate over the raw logs and unpacked data for SetSettlementPrice events raised by the PriceFeeder contract.
type PriceFeederSetSettlementPriceIterator struct {
	Event *PriceFeederSetSettlementPrice // Event containing the contract specifics and raw log

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
func (it *PriceFeederSetSettlementPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederSetSettlementPrice)
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
		it.Event = new(PriceFeederSetSettlementPrice)
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
func (it *PriceFeederSetSettlementPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederSetSettlementPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederSetSettlementPrice represents a SetSettlementPrice event raised by the PriceFeeder contract.
type PriceFeederSetSettlementPrice struct {
	MarketID                  [32]byte
	Price                     *big.Int
	Timestamp                 *big.Int
	MarketExpirationTimestamp *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetSettlementPrice is a free log retrieval operation binding the contract event 0x4d1e7bc440f21abae9f84f1f4b5f1721205dd6fbd01100ffdd0e4b7d1e9be62c.
//
// Solidity: event SetSettlementPrice(bytes32 indexed marketID, uint256 price, uint256 timestamp, uint256 marketExpirationTimestamp)
func (f *PriceFeederFilterer) FilterSetSettlementPrice(opts *bind.FilterOpts, marketID [][32]byte) (*PriceFeederSetSettlementPriceIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "SetSettlementPrice", marketIDRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeederSetSettlementPriceIterator{contract: f.contract, event: "SetSettlementPrice", logs: logs, sub: sub}, nil
}

// WatchSetSettlementPrice is a free log subscription operation binding the contract event 0x4d1e7bc440f21abae9f84f1f4b5f1721205dd6fbd01100ffdd0e4b7d1e9be62c.
//
// Solidity: event SetSettlementPrice(bytes32 indexed marketID, uint256 price, uint256 timestamp, uint256 marketExpirationTimestamp)
func (f *PriceFeederFilterer) WatchSetSettlementPrice(opts *bind.WatchOpts, sink chan<- *PriceFeederSetSettlementPrice, marketID [][32]byte) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "SetSettlementPrice", marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederSetSettlementPrice)
				if err := f.contract.UnpackLog(event, "SetSettlementPrice", log); err != nil {
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

// ParseSetSettlementPrice is a log parse operation binding the contract event 0x4d1e7bc440f21abae9f84f1f4b5f1721205dd6fbd01100ffdd0e4b7d1e9be62c.
//
// Solidity: event SetSettlementPrice(bytes32 indexed marketID, uint256 price, uint256 timestamp, uint256 marketExpirationTimestamp)
func (f *PriceFeederFilterer) ParseSetSettlementPrice(log types.Log) (*PriceFeederSetSettlementPrice, error) {
	event := new(PriceFeederSetSettlementPrice)
	if err := f.contract.UnpackLog(event, "SetSettlementPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
