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

// MixinCoordinatorApprovalCoordinatorApproval is an auto generated low-level Go binding around an user-defined struct.
type MixinCoordinatorApprovalCoordinatorApproval struct {
	TxOrigin             common.Address
	TransactionHash      [32]byte
	TransactionSignature []byte
}

// CoordinatorABI is the input ABI used to generate the binding from.
const CoordinatorABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"EIP712_COORDINATOR_APPROVAL_SCHEMA_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_COORDINATOR_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_COORDINATOR_DOMAIN_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_COORDINATOR_DOMAIN_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_EXCHANGE_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structLibZeroExTransaction.ZeroExTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"transactionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"approvalSignatures\",\"type\":\"bytes[]\"}],\"name\":\"assertValidCoordinatorApprovals\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"decodeOrdersFromFillData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structLibZeroExTransaction.ZeroExTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"transactionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"approvalSignatures\",\"type\":\"bytes[]\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"transactionSignature\",\"type\":\"bytes\"}],\"internalType\":\"structMixinCoordinatorApproval.CoordinatorApproval\",\"name\":\"approval\",\"type\":\"tuple\"}],\"name\":\"getCoordinatorApprovalHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"approvalHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"getSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeBank\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContractAddressIfExists\",\"type\":\"address\"}],\"name\":\"initializeCoordinatorDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContractAddressIfExists\",\"type\":\"address\"}],\"name\":\"initializeExchangeDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Coordinator is an auto generated Go binding around an Ethereum contract.
type Coordinator struct {
	CoordinatorCaller     // Read-only binding to the contract
	CoordinatorTransactor // Write-only binding to the contract
	CoordinatorFilterer   // Log filterer for contract events
}

// CoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoordinatorSession struct {
	Contract     *Coordinator         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoordinatorCallerSession struct {
	Contract *CoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoordinatorTransactorSession struct {
	Contract     *CoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoordinatorRaw struct {
	Contract *Coordinator // Generic contract binding to access the raw methods on
}

// CoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoordinatorCallerRaw struct {
	Contract *CoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// CoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoordinatorTransactorRaw struct {
	Contract *CoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoordinator creates a new instance of Coordinator, bound to a specific deployed contract.
func NewCoordinator(address common.Address, backend bind.ContractBackend) (*Coordinator, error) {
	contract, err := bindCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coordinator{CoordinatorCaller: CoordinatorCaller{contract: contract}, CoordinatorTransactor: CoordinatorTransactor{contract: contract}, CoordinatorFilterer: CoordinatorFilterer{contract: contract}}, nil
}

// NewCoordinatorCaller creates a new read-only instance of Coordinator, bound to a specific deployed contract.
func NewCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*CoordinatorCaller, error) {
	contract, err := bindCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoordinatorCaller{contract: contract}, nil
}

// NewCoordinatorTransactor creates a new write-only instance of Coordinator, bound to a specific deployed contract.
func NewCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CoordinatorTransactor, error) {
	contract, err := bindCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoordinatorTransactor{contract: contract}, nil
}

// NewCoordinatorFilterer creates a new log filterer instance of Coordinator, bound to a specific deployed contract.
func NewCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CoordinatorFilterer, error) {
	contract, err := bindCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoordinatorFilterer{contract: contract}, nil
}

// bindCoordinator binds a generic wrapper to an already deployed contract.
func bindCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coordinator *CoordinatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coordinator.Contract.CoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coordinator *CoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coordinator.Contract.CoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coordinator *CoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coordinator.Contract.CoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coordinator *CoordinatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Coordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coordinator *CoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coordinator *CoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coordinator.Contract.contract.Transact(opts, method, params...)
}

// EIP712COORDINATORAPPROVALSCHEMAHASH is a free data retrieval call binding the contract method 0xe1c71578.
//
// Solidity: function EIP712_COORDINATOR_APPROVAL_SCHEMA_HASH() constant returns(bytes32)
func (_Coordinator *CoordinatorCaller) EIP712COORDINATORAPPROVALSCHEMAHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Coordinator.contract.Call(opts, out, "EIP712_COORDINATOR_APPROVAL_SCHEMA_HASH")
	return *ret0, err
}

// EIP712COORDINATORAPPROVALSCHEMAHASH is a free data retrieval call binding the contract method 0xe1c71578.
//
// Solidity: function EIP712_COORDINATOR_APPROVAL_SCHEMA_HASH() constant returns(bytes32)
func (_Coordinator *CoordinatorSession) EIP712COORDINATORAPPROVALSCHEMAHASH() ([32]byte, error) {
	return _Coordinator.Contract.EIP712COORDINATORAPPROVALSCHEMAHASH(&_Coordinator.CallOpts)
}

// EIP712COORDINATORAPPROVALSCHEMAHASH is a free data retrieval call binding the contract method 0xe1c71578.
//
// Solidity: function EIP712_COORDINATOR_APPROVAL_SCHEMA_HASH() constant returns(bytes32)
func (_Coordinator *CoordinatorCallerSession) EIP712COORDINATORAPPROVALSCHEMAHASH() ([32]byte, error) {
	return _Coordinator.Contract.EIP712COORDINATORAPPROVALSCHEMAHASH(&_Coordinator.CallOpts)
}

// EIP712COORDINATORDOMAINHASH is a free data retrieval call binding the contract method 0xfb6961cc.
//
// Solidity: function EIP712_COORDINATOR_DOMAIN_HASH() constant returns(bytes32)
func (_Coordinator *CoordinatorCaller) EIP712COORDINATORDOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Coordinator.contract.Call(opts, out, "EIP712_COORDINATOR_DOMAIN_HASH")
	return *ret0, err
}

// EIP712COORDINATORDOMAINHASH is a free data retrieval call binding the contract method 0xfb6961cc.
//
// Solidity: function EIP712_COORDINATOR_DOMAIN_HASH() constant returns(bytes32)
func (_Coordinator *CoordinatorSession) EIP712COORDINATORDOMAINHASH() ([32]byte, error) {
	return _Coordinator.Contract.EIP712COORDINATORDOMAINHASH(&_Coordinator.CallOpts)
}

// EIP712COORDINATORDOMAINHASH is a free data retrieval call binding the contract method 0xfb6961cc.
//
// Solidity: function EIP712_COORDINATOR_DOMAIN_HASH() constant returns(bytes32)
func (_Coordinator *CoordinatorCallerSession) EIP712COORDINATORDOMAINHASH() ([32]byte, error) {
	return _Coordinator.Contract.EIP712COORDINATORDOMAINHASH(&_Coordinator.CallOpts)
}

// EIP712COORDINATORDOMAINNAME is a free data retrieval call binding the contract method 0x89fab5b7.
//
// Solidity: function EIP712_COORDINATOR_DOMAIN_NAME() constant returns(string)
func (_Coordinator *CoordinatorCaller) EIP712COORDINATORDOMAINNAME(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Coordinator.contract.Call(opts, out, "EIP712_COORDINATOR_DOMAIN_NAME")
	return *ret0, err
}

// EIP712COORDINATORDOMAINNAME is a free data retrieval call binding the contract method 0x89fab5b7.
//
// Solidity: function EIP712_COORDINATOR_DOMAIN_NAME() constant returns(string)
func (_Coordinator *CoordinatorSession) EIP712COORDINATORDOMAINNAME() (string, error) {
	return _Coordinator.Contract.EIP712COORDINATORDOMAINNAME(&_Coordinator.CallOpts)
}

// EIP712COORDINATORDOMAINNAME is a free data retrieval call binding the contract method 0x89fab5b7.
//
// Solidity: function EIP712_COORDINATOR_DOMAIN_NAME() constant returns(string)
func (_Coordinator *CoordinatorCallerSession) EIP712COORDINATORDOMAINNAME() (string, error) {
	return _Coordinator.Contract.EIP712COORDINATORDOMAINNAME(&_Coordinator.CallOpts)
}

// EIP712COORDINATORDOMAINVERSION is a free data retrieval call binding the contract method 0xb2562b7a.
//
// Solidity: function EIP712_COORDINATOR_DOMAIN_VERSION() constant returns(string)
func (_Coordinator *CoordinatorCaller) EIP712COORDINATORDOMAINVERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Coordinator.contract.Call(opts, out, "EIP712_COORDINATOR_DOMAIN_VERSION")
	return *ret0, err
}

// EIP712COORDINATORDOMAINVERSION is a free data retrieval call binding the contract method 0xb2562b7a.
//
// Solidity: function EIP712_COORDINATOR_DOMAIN_VERSION() constant returns(string)
func (_Coordinator *CoordinatorSession) EIP712COORDINATORDOMAINVERSION() (string, error) {
	return _Coordinator.Contract.EIP712COORDINATORDOMAINVERSION(&_Coordinator.CallOpts)
}

// EIP712COORDINATORDOMAINVERSION is a free data retrieval call binding the contract method 0xb2562b7a.
//
// Solidity: function EIP712_COORDINATOR_DOMAIN_VERSION() constant returns(string)
func (_Coordinator *CoordinatorCallerSession) EIP712COORDINATORDOMAINVERSION() (string, error) {
	return _Coordinator.Contract.EIP712COORDINATORDOMAINVERSION(&_Coordinator.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() constant returns(bytes32)
func (_Coordinator *CoordinatorCaller) EIP712EXCHANGEDOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Coordinator.contract.Call(opts, out, "EIP712_EXCHANGE_DOMAIN_HASH")
	return *ret0, err
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() constant returns(bytes32)
func (_Coordinator *CoordinatorSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return _Coordinator.Contract.EIP712EXCHANGEDOMAINHASH(&_Coordinator.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() constant returns(bytes32)
func (_Coordinator *CoordinatorCallerSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return _Coordinator.Contract.EIP712EXCHANGEDOMAINHASH(&_Coordinator.CallOpts)
}

// AssertValidCoordinatorApprovals is a free data retrieval call binding the contract method 0x52813679.
//
// Solidity: function assertValidCoordinatorApprovals(ZeroExTransaction transaction, address txOrigin, bytes transactionSignature, bytes[] approvalSignatures) constant returns()
func (_Coordinator *CoordinatorCaller) AssertValidCoordinatorApprovals(opts *bind.CallOpts, transaction ZeroExTransaction, txOrigin common.Address, transactionSignature []byte, approvalSignatures [][]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Coordinator.contract.Call(opts, out, "assertValidCoordinatorApprovals", transaction, txOrigin, transactionSignature, approvalSignatures)
	return err
}

// AssertValidCoordinatorApprovals is a free data retrieval call binding the contract method 0x52813679.
//
// Solidity: function assertValidCoordinatorApprovals(ZeroExTransaction transaction, address txOrigin, bytes transactionSignature, bytes[] approvalSignatures) constant returns()
func (_Coordinator *CoordinatorSession) AssertValidCoordinatorApprovals(transaction ZeroExTransaction, txOrigin common.Address, transactionSignature []byte, approvalSignatures [][]byte) error {
	return _Coordinator.Contract.AssertValidCoordinatorApprovals(&_Coordinator.CallOpts, transaction, txOrigin, transactionSignature, approvalSignatures)
}

// AssertValidCoordinatorApprovals is a free data retrieval call binding the contract method 0x52813679.
//
// Solidity: function assertValidCoordinatorApprovals(ZeroExTransaction transaction, address txOrigin, bytes transactionSignature, bytes[] approvalSignatures) constant returns()
func (_Coordinator *CoordinatorCallerSession) AssertValidCoordinatorApprovals(transaction ZeroExTransaction, txOrigin common.Address, transactionSignature []byte, approvalSignatures [][]byte) error {
	return _Coordinator.Contract.AssertValidCoordinatorApprovals(&_Coordinator.CallOpts, transaction, txOrigin, transactionSignature, approvalSignatures)
}

// DecodeOrdersFromFillData is a free data retrieval call binding the contract method 0xee55b968.
//
// Solidity: function decodeOrdersFromFillData(bytes data) constant returns([]Order orders)
func (_Coordinator *CoordinatorCaller) DecodeOrdersFromFillData(opts *bind.CallOpts, data []byte) ([]Order, error) {
	var (
		ret0 = new([]Order)
	)
	out := ret0
	err := _Coordinator.contract.Call(opts, out, "decodeOrdersFromFillData", data)
	return *ret0, err
}

// DecodeOrdersFromFillData is a free data retrieval call binding the contract method 0xee55b968.
//
// Solidity: function decodeOrdersFromFillData(bytes data) constant returns([]Order orders)
func (_Coordinator *CoordinatorSession) DecodeOrdersFromFillData(data []byte) ([]Order, error) {
	return _Coordinator.Contract.DecodeOrdersFromFillData(&_Coordinator.CallOpts, data)
}

// DecodeOrdersFromFillData is a free data retrieval call binding the contract method 0xee55b968.
//
// Solidity: function decodeOrdersFromFillData(bytes data) constant returns([]Order orders)
func (_Coordinator *CoordinatorCallerSession) DecodeOrdersFromFillData(data []byte) ([]Order, error) {
	return _Coordinator.Contract.DecodeOrdersFromFillData(&_Coordinator.CallOpts, data)
}

// GetCoordinatorApprovalHash is a free data retrieval call binding the contract method 0xfdd059a5.
//
// Solidity: function getCoordinatorApprovalHash(MixinCoordinatorApprovalCoordinatorApproval approval) constant returns(bytes32 approvalHash)
func (_Coordinator *CoordinatorCaller) GetCoordinatorApprovalHash(opts *bind.CallOpts, approval MixinCoordinatorApprovalCoordinatorApproval) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Coordinator.contract.Call(opts, out, "getCoordinatorApprovalHash", approval)
	return *ret0, err
}

// GetCoordinatorApprovalHash is a free data retrieval call binding the contract method 0xfdd059a5.
//
// Solidity: function getCoordinatorApprovalHash(MixinCoordinatorApprovalCoordinatorApproval approval) constant returns(bytes32 approvalHash)
func (_Coordinator *CoordinatorSession) GetCoordinatorApprovalHash(approval MixinCoordinatorApprovalCoordinatorApproval) ([32]byte, error) {
	return _Coordinator.Contract.GetCoordinatorApprovalHash(&_Coordinator.CallOpts, approval)
}

// GetCoordinatorApprovalHash is a free data retrieval call binding the contract method 0xfdd059a5.
//
// Solidity: function getCoordinatorApprovalHash(MixinCoordinatorApprovalCoordinatorApproval approval) constant returns(bytes32 approvalHash)
func (_Coordinator *CoordinatorCallerSession) GetCoordinatorApprovalHash(approval MixinCoordinatorApprovalCoordinatorApproval) ([32]byte, error) {
	return _Coordinator.Contract.GetCoordinatorApprovalHash(&_Coordinator.CallOpts, approval)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x0f7d8e39.
//
// Solidity: function getSignerAddress(bytes32 hash, bytes signature) constant returns(address signerAddress)
func (_Coordinator *CoordinatorCaller) GetSignerAddress(opts *bind.CallOpts, hash [32]byte, signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Coordinator.contract.Call(opts, out, "getSignerAddress", hash, signature)
	return *ret0, err
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x0f7d8e39.
//
// Solidity: function getSignerAddress(bytes32 hash, bytes signature) constant returns(address signerAddress)
func (_Coordinator *CoordinatorSession) GetSignerAddress(hash [32]byte, signature []byte) (common.Address, error) {
	return _Coordinator.Contract.GetSignerAddress(&_Coordinator.CallOpts, hash, signature)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x0f7d8e39.
//
// Solidity: function getSignerAddress(bytes32 hash, bytes signature) constant returns(address signerAddress)
func (_Coordinator *CoordinatorCallerSession) GetSignerAddress(hash [32]byte, signature []byte) (common.Address, error) {
	return _Coordinator.Contract.GetSignerAddress(&_Coordinator.CallOpts, hash, signature)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xda4fe074.
//
// Solidity: function executeTransaction(ZeroExTransaction transaction, address txOrigin, bytes transactionSignature, bytes[] approvalSignatures) returns()
func (_Coordinator *CoordinatorTransactor) ExecuteTransaction(opts *bind.TransactOpts, transaction ZeroExTransaction, txOrigin common.Address, transactionSignature []byte, approvalSignatures [][]byte) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "executeTransaction", transaction, txOrigin, transactionSignature, approvalSignatures)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xda4fe074.
//
// Solidity: function executeTransaction(ZeroExTransaction transaction, address txOrigin, bytes transactionSignature, bytes[] approvalSignatures) returns()
func (_Coordinator *CoordinatorSession) ExecuteTransaction(transaction ZeroExTransaction, txOrigin common.Address, transactionSignature []byte, approvalSignatures [][]byte) (*types.Transaction, error) {
	return _Coordinator.Contract.ExecuteTransaction(&_Coordinator.TransactOpts, transaction, txOrigin, transactionSignature, approvalSignatures)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xda4fe074.
//
// Solidity: function executeTransaction(ZeroExTransaction transaction, address txOrigin, bytes transactionSignature, bytes[] approvalSignatures) returns()
func (_Coordinator *CoordinatorTransactorSession) ExecuteTransaction(transaction ZeroExTransaction, txOrigin common.Address, transactionSignature []byte, approvalSignatures [][]byte) (*types.Transaction, error) {
	return _Coordinator.Contract.ExecuteTransaction(&_Coordinator.TransactOpts, transaction, txOrigin, transactionSignature, approvalSignatures)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _exchange, address _stakeBank, uint256 _chainId) returns()
func (_Coordinator *CoordinatorTransactor) Initialize(opts *bind.TransactOpts, _exchange common.Address, _stakeBank common.Address, _chainId *big.Int) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "initialize", _exchange, _stakeBank, _chainId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _exchange, address _stakeBank, uint256 _chainId) returns()
func (_Coordinator *CoordinatorSession) Initialize(_exchange common.Address, _stakeBank common.Address, _chainId *big.Int) (*types.Transaction, error) {
	return _Coordinator.Contract.Initialize(&_Coordinator.TransactOpts, _exchange, _stakeBank, _chainId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _exchange, address _stakeBank, uint256 _chainId) returns()
func (_Coordinator *CoordinatorTransactorSession) Initialize(_exchange common.Address, _stakeBank common.Address, _chainId *big.Int) (*types.Transaction, error) {
	return _Coordinator.Contract.Initialize(&_Coordinator.TransactOpts, _exchange, _stakeBank, _chainId)
}

// InitializeCoordinatorDomain is a paid mutator transaction binding the contract method 0xf7ce0ab3.
//
// Solidity: function initializeCoordinatorDomain(uint256 chainId, address verifyingContractAddressIfExists) returns()
func (_Coordinator *CoordinatorTransactor) InitializeCoordinatorDomain(opts *bind.TransactOpts, chainId *big.Int, verifyingContractAddressIfExists common.Address) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "initializeCoordinatorDomain", chainId, verifyingContractAddressIfExists)
}

// InitializeCoordinatorDomain is a paid mutator transaction binding the contract method 0xf7ce0ab3.
//
// Solidity: function initializeCoordinatorDomain(uint256 chainId, address verifyingContractAddressIfExists) returns()
func (_Coordinator *CoordinatorSession) InitializeCoordinatorDomain(chainId *big.Int, verifyingContractAddressIfExists common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.InitializeCoordinatorDomain(&_Coordinator.TransactOpts, chainId, verifyingContractAddressIfExists)
}

// InitializeCoordinatorDomain is a paid mutator transaction binding the contract method 0xf7ce0ab3.
//
// Solidity: function initializeCoordinatorDomain(uint256 chainId, address verifyingContractAddressIfExists) returns()
func (_Coordinator *CoordinatorTransactorSession) InitializeCoordinatorDomain(chainId *big.Int, verifyingContractAddressIfExists common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.InitializeCoordinatorDomain(&_Coordinator.TransactOpts, chainId, verifyingContractAddressIfExists)
}

// InitializeExchangeDomain is a paid mutator transaction binding the contract method 0xcca92fb5.
//
// Solidity: function initializeExchangeDomain(uint256 chainId, address verifyingContractAddressIfExists) returns()
func (_Coordinator *CoordinatorTransactor) InitializeExchangeDomain(opts *bind.TransactOpts, chainId *big.Int, verifyingContractAddressIfExists common.Address) (*types.Transaction, error) {
	return _Coordinator.contract.Transact(opts, "initializeExchangeDomain", chainId, verifyingContractAddressIfExists)
}

// InitializeExchangeDomain is a paid mutator transaction binding the contract method 0xcca92fb5.
//
// Solidity: function initializeExchangeDomain(uint256 chainId, address verifyingContractAddressIfExists) returns()
func (_Coordinator *CoordinatorSession) InitializeExchangeDomain(chainId *big.Int, verifyingContractAddressIfExists common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.InitializeExchangeDomain(&_Coordinator.TransactOpts, chainId, verifyingContractAddressIfExists)
}

// InitializeExchangeDomain is a paid mutator transaction binding the contract method 0xcca92fb5.
//
// Solidity: function initializeExchangeDomain(uint256 chainId, address verifyingContractAddressIfExists) returns()
func (_Coordinator *CoordinatorTransactorSession) InitializeExchangeDomain(chainId *big.Int, verifyingContractAddressIfExists common.Address) (*types.Transaction, error) {
	return _Coordinator.Contract.InitializeExchangeDomain(&_Coordinator.TransactOpts, chainId, verifyingContractAddressIfExists)
}
