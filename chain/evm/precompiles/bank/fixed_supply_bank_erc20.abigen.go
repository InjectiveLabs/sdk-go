// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// FixedSupplyBankERC20MetaData contains all meta data concerning the FixedSupplyBankERC20 contract.
var FixedSupplyBankERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"initial_supply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Failure\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60806040819052600580546001600160a01b031916606417905561143d38819003908190833981016040819052610035916104ac565b83838360405180602001604052805f81525060405180602001604052805f815250816003908161006591906105b3565b50600461007282826105b3565b5050505f8351118061008457505f8251115b8061009157505f8160ff16115b1561010c57600554604051630df4b0bd60e21b81526001600160a01b03909116906337d2c2f4906100ca90869086908690600401610698565b6020604051808303815f875af11580156100e6573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061010a91906106d0565b505b505081159050610120576101203382610129565b5050505061081e565b6001600160a01b0382166101575760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6101625f8383610166565b5050565b6001600160a01b03831661028c576005546040516340c10f1960e01b81526001600160a01b03848116600483015260248201849052909116906340c10f19906044016020604051808303815f875af19250505080156101e2575060408051601f3d908101601f191682019092526101df918101906106d0565b60015b610286576101ee6106f6565b806308c379a003610244575061020261070f565b8061020d5750610246565b8060405160200161021e9190610791565b60408051601f198184030181529082905262461bcd60e51b825261014e916004016107c8565b505b3d80801561026f576040519150601f19603f3d011682016040523d82523d5f602084013e610274565b606091505b508060405160200161021e91906107da565b5061038c565b6001600160a01b03821661030e57600554604051632770a7eb60e21b81526001600160a01b0385811660048301526024820184905290911690639dc29fac906044016020604051808303815f875af11580156102ea573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061028691906106d0565b6005546040516317d5759960e31b81526001600160a01b0385811660048301528481166024830152604482018490529091169063beabacc8906064016020604051808303815f875af1158015610366573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061038a91906106d0565b505b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516103d191815260200190565b60405180910390a3505050565b634e487b7160e01b5f52604160045260245ffd5b601f8201601f191681016001600160401b0381118282101715610417576104176103de565b6040525050565b5f5b83811015610438578181015183820152602001610420565b50505f910152565b5f82601f83011261044f575f5ffd5b81516001600160401b03811115610468576104686103de565b60405161047f601f8301601f1916602001826103f2565b818152846020838601011115610493575f5ffd5b6104a482602083016020870161041e565b949350505050565b5f5f5f5f608085870312156104bf575f5ffd5b84516001600160401b038111156104d4575f5ffd5b6104e087828801610440565b602087015190955090506001600160401b038111156104fd575f5ffd5b61050987828801610440565b935050604085015160ff8116811461051f575f5ffd5b6060959095015193969295505050565b600181811c9082168061054357607f821691505b60208210810361056157634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156105ae57805f5260205f20601f840160051c8101602085101561058c5750805b601f840160051c820191505b818110156105ab575f8155600101610598565b50505b505050565b81516001600160401b038111156105cc576105cc6103de565b6105e0816105da845461052f565b84610567565b6020601f821160018114610612575f83156105fb5750848201515b5f19600385901b1c1916600184901b1784556105ab565b5f84815260208120601f198516915b828110156106415787850151825560209485019460019092019101610621565b508482101561065e57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f815180845261068481602086016020860161041e565b601f01601f19169290920160200192915050565b606081525f6106aa606083018661066d565b82810360208401526106bc818661066d565b91505060ff83166040830152949350505050565b5f602082840312156106e0575f5ffd5b815180151581146106ef575f5ffd5b9392505050565b5f60033d111561070c5760045f5f3e505f5160e01c5b90565b5f60443d101561071c5790565b6040513d600319016004823e80516001600160401b0381113d6024830111171561074557505090565b81810180516001600160401b03811115610760575050505090565b3d840160031901828201602001111561077a575050505090565b610789602082850101856103f2565b509392505050565b6f03330b4b632b2103a379036b4b73a1d160851b81525f82516107bb81601085016020870161041e565b9190910160100192915050565b602081525f6106ef602083018461066d565b7f6661696c656420746f206d696e743a20756e6b6e6f776e206572726f723a200081525f825161081181601f85016020870161041e565b91909101601f0192915050565b610c128061082b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610090575f3560e01c8063313ce56711610063578063313ce567146100fe57806370a082311461011857806395d89b411461012b578063a9059cbb14610133578063dd62ed3e14610146575f5ffd5b806306fdde0314610094578063095ea7b3146100b257806318160ddd146100d557806323b872dd146100eb575b5f5ffd5b61009c61017e565b6040516100a99190610873565b60405180910390f35b6100c56100c03660046108c0565b6101f8565b60405190151581526020016100a9565b6100dd61020f565b6040519081526020016100a9565b6100c56100f93660046108e8565b61027e565b6101066102a1565b60405160ff90911681526020016100a9565b6100dd610126366004610922565b610319565b61009c610391565b6100c56101413660046108c0565b61040a565b6100dd610154366004610942565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6005546040516315d10ab960e11b815230600482015260609182916001600160a01b0390911690632ba21572906024015f60405180830381865afa1580156101c8573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526101ef9190810190610a19565b50909392505050565b5f33610205818585610417565b5060019392505050565b6005546040516339370aa960e21b81523060048201525f916001600160a01b03169063e4dc2aa490602401602060405180830381865afa158015610255573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102799190610a96565b905090565b5f3361028b858285610429565b6102968585856104aa565b506001949350505050565b6005546040516315d10ab960e11b81523060048201525f9182916001600160a01b0390911690632ba21572906024015f60405180830381865afa1580156102ea573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103119190810190610a19565b949350505050565b600554604051633de222bb60e21b81523060048201526001600160a01b0383811660248301525f92169063f7888aec90604401602060405180830381865afa158015610367573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061038b9190610a96565b92915050565b6005546040516315d10ab960e11b815230600482015260609182916001600160a01b0390911690632ba21572906024015f60405180830381865afa1580156103db573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526104029190810190610a19565b509392505050565b5f336102058185856104aa565b6104248383836001610507565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f198110156104a4578181101561049657604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b6104a484848484035f610507565b50505050565b6001600160a01b0383166104d357604051634b637e8f60e11b81525f600482015260240161048d565b6001600160a01b0382166104fc5760405163ec442f0560e01b81525f600482015260240161048d565b6104248383836105d9565b6001600160a01b0384166105305760405163e602df0560e01b81525f600482015260240161048d565b6001600160a01b03831661055957604051634a1406b160e11b81525f600482015260240161048d565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156104a457826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105cb91815260200190565b60405180910390a350505050565b6001600160a01b0383166106ff576005546040516340c10f1960e01b81526001600160a01b03848116600483015260248201849052909116906340c10f19906044016020604051808303815f875af1925050508015610655575060408051601f3d908101601f1916820190925261065291810190610aad565b60015b6106f957610661610acc565b806308c379a0036106b75750610675610ae5565b8061068057506106b9565b806040516020016106919190610b61565b60408051601f198184030181529082905262461bcd60e51b825261048d91600401610873565b505b3d8080156106e2576040519150601f19603f3d011682016040523d82523d5f602084013e6106e7565b606091505b50806040516020016106919190610b98565b506107ff565b6001600160a01b03821661078157600554604051632770a7eb60e21b81526001600160a01b0385811660048301526024820184905290911690639dc29fac906044016020604051808303815f875af115801561075d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106f99190610aad565b6005546040516317d5759960e31b81526001600160a01b0385811660048301528481166024830152604482018490529091169063beabacc8906064016020604051808303815f875af11580156107d9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107fd9190610aad565b505b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161084491815260200190565b60405180910390a3505050565b5f5b8381101561086b578181015183820152602001610853565b50505f910152565b602081525f8251806020840152610891816040850160208701610851565b601f01601f19169190910160400192915050565b80356001600160a01b03811681146108bb575f5ffd5b919050565b5f5f604083850312156108d1575f5ffd5b6108da836108a5565b946020939093013593505050565b5f5f5f606084860312156108fa575f5ffd5b610903846108a5565b9250610911602085016108a5565b929592945050506040919091013590565b5f60208284031215610932575f5ffd5b61093b826108a5565b9392505050565b5f5f60408385031215610953575f5ffd5b61095c836108a5565b915061096a602084016108a5565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b601f8201601f1916810167ffffffffffffffff811182821017156109ad576109ad610973565b6040525050565b5f82601f8301126109c3575f5ffd5b815167ffffffffffffffff8111156109dd576109dd610973565b6040516109f4601f8301601f191660200182610987565b818152846020838601011115610a08575f5ffd5b610311826020830160208701610851565b5f5f5f60608486031215610a2b575f5ffd5b835167ffffffffffffffff811115610a41575f5ffd5b610a4d868287016109b4565b935050602084015167ffffffffffffffff811115610a69575f5ffd5b610a75868287016109b4565b925050604084015160ff81168114610a8b575f5ffd5b809150509250925092565b5f60208284031215610aa6575f5ffd5b5051919050565b5f60208284031215610abd575f5ffd5b8151801515811461093b575f5ffd5b5f60033d1115610ae25760045f5f3e505f5160e01c5b90565b5f60443d1015610af25790565b6040513d600319016004823e80513d602482011167ffffffffffffffff82111715610b1c57505090565b808201805167ffffffffffffffff811115610b38575050505090565b3d8401600319018282016020011115610b52575050505090565b61040260208285010185610987565b6f03330b4b632b2103a379036b4b73a1d160851b81525f8251610b8b816010850160208701610851565b9190910160100192915050565b7f6661696c656420746f206d696e743a20756e6b6e6f776e206572726f723a200081525f8251610bcf81601f850160208701610851565b91909101601f019291505056fea2646970667358221220700a7c6e315d053531301f75f3f482db54f371f66e2d6d558a3b367b17af6b8164736f6c634300081e0033",
}

// FixedSupplyBankERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use FixedSupplyBankERC20MetaData.ABI instead.
var FixedSupplyBankERC20ABI = FixedSupplyBankERC20MetaData.ABI

// FixedSupplyBankERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FixedSupplyBankERC20MetaData.Bin instead.
var FixedSupplyBankERC20Bin = FixedSupplyBankERC20MetaData.Bin

// DeployFixedSupplyBankERC20 deploys a new Ethereum contract, binding an instance of FixedSupplyBankERC20 to it.
func DeployFixedSupplyBankERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, initial_supply_ *big.Int) (common.Address, *types.Transaction, *FixedSupplyBankERC20, error) {
	parsed, err := FixedSupplyBankERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FixedSupplyBankERC20Bin), backend, name_, symbol_, decimals_, initial_supply_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FixedSupplyBankERC20{FixedSupplyBankERC20Caller: FixedSupplyBankERC20Caller{contract: contract}, FixedSupplyBankERC20Transactor: FixedSupplyBankERC20Transactor{contract: contract}, FixedSupplyBankERC20Filterer: FixedSupplyBankERC20Filterer{contract: contract}}, nil
}

// FixedSupplyBankERC20 is an auto generated Go binding around an Ethereum contract.
type FixedSupplyBankERC20 struct {
	FixedSupplyBankERC20Caller     // Read-only binding to the contract
	FixedSupplyBankERC20Transactor // Write-only binding to the contract
	FixedSupplyBankERC20Filterer   // Log filterer for contract events
}

// FixedSupplyBankERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type FixedSupplyBankERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedSupplyBankERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FixedSupplyBankERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedSupplyBankERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FixedSupplyBankERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedSupplyBankERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FixedSupplyBankERC20Session struct {
	Contract     *FixedSupplyBankERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FixedSupplyBankERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FixedSupplyBankERC20CallerSession struct {
	Contract *FixedSupplyBankERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FixedSupplyBankERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FixedSupplyBankERC20TransactorSession struct {
	Contract     *FixedSupplyBankERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FixedSupplyBankERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type FixedSupplyBankERC20Raw struct {
	Contract *FixedSupplyBankERC20 // Generic contract binding to access the raw methods on
}

// FixedSupplyBankERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FixedSupplyBankERC20CallerRaw struct {
	Contract *FixedSupplyBankERC20Caller // Generic read-only contract binding to access the raw methods on
}

// FixedSupplyBankERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FixedSupplyBankERC20TransactorRaw struct {
	Contract *FixedSupplyBankERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFixedSupplyBankERC20 creates a new instance of FixedSupplyBankERC20, bound to a specific deployed contract.
func NewFixedSupplyBankERC20(address common.Address, backend bind.ContractBackend) (*FixedSupplyBankERC20, error) {
	contract, err := bindFixedSupplyBankERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyBankERC20{FixedSupplyBankERC20Caller: FixedSupplyBankERC20Caller{contract: contract}, FixedSupplyBankERC20Transactor: FixedSupplyBankERC20Transactor{contract: contract}, FixedSupplyBankERC20Filterer: FixedSupplyBankERC20Filterer{contract: contract}}, nil
}

// NewFixedSupplyBankERC20Caller creates a new read-only instance of FixedSupplyBankERC20, bound to a specific deployed contract.
func NewFixedSupplyBankERC20Caller(address common.Address, caller bind.ContractCaller) (*FixedSupplyBankERC20Caller, error) {
	contract, err := bindFixedSupplyBankERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyBankERC20Caller{contract: contract}, nil
}

// NewFixedSupplyBankERC20Transactor creates a new write-only instance of FixedSupplyBankERC20, bound to a specific deployed contract.
func NewFixedSupplyBankERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*FixedSupplyBankERC20Transactor, error) {
	contract, err := bindFixedSupplyBankERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyBankERC20Transactor{contract: contract}, nil
}

// NewFixedSupplyBankERC20Filterer creates a new log filterer instance of FixedSupplyBankERC20, bound to a specific deployed contract.
func NewFixedSupplyBankERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*FixedSupplyBankERC20Filterer, error) {
	contract, err := bindFixedSupplyBankERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyBankERC20Filterer{contract: contract}, nil
}

// bindFixedSupplyBankERC20 binds a generic wrapper to an already deployed contract.
func bindFixedSupplyBankERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FixedSupplyBankERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FixedSupplyBankERC20.Contract.FixedSupplyBankERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.FixedSupplyBankERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.FixedSupplyBankERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FixedSupplyBankERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FixedSupplyBankERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FixedSupplyBankERC20.Contract.Allowance(&_FixedSupplyBankERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FixedSupplyBankERC20.Contract.Allowance(&_FixedSupplyBankERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FixedSupplyBankERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _FixedSupplyBankERC20.Contract.BalanceOf(&_FixedSupplyBankERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FixedSupplyBankERC20.Contract.BalanceOf(&_FixedSupplyBankERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FixedSupplyBankERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Session) Decimals() (uint8, error) {
	return _FixedSupplyBankERC20.Contract.Decimals(&_FixedSupplyBankERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20CallerSession) Decimals() (uint8, error) {
	return _FixedSupplyBankERC20.Contract.Decimals(&_FixedSupplyBankERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FixedSupplyBankERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Session) Name() (string, error) {
	return _FixedSupplyBankERC20.Contract.Name(&_FixedSupplyBankERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20CallerSession) Name() (string, error) {
	return _FixedSupplyBankERC20.Contract.Name(&_FixedSupplyBankERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FixedSupplyBankERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Session) Symbol() (string, error) {
	return _FixedSupplyBankERC20.Contract.Symbol(&_FixedSupplyBankERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20CallerSession) Symbol() (string, error) {
	return _FixedSupplyBankERC20.Contract.Symbol(&_FixedSupplyBankERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FixedSupplyBankERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Session) TotalSupply() (*big.Int, error) {
	return _FixedSupplyBankERC20.Contract.TotalSupply(&_FixedSupplyBankERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _FixedSupplyBankERC20.Contract.TotalSupply(&_FixedSupplyBankERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.Approve(&_FixedSupplyBankERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.Approve(&_FixedSupplyBankERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.Transfer(&_FixedSupplyBankERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.Transfer(&_FixedSupplyBankERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.TransferFrom(&_FixedSupplyBankERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FixedSupplyBankERC20.Contract.TransferFrom(&_FixedSupplyBankERC20.TransactOpts, from, to, value)
}

// FixedSupplyBankERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FixedSupplyBankERC20 contract.
type FixedSupplyBankERC20ApprovalIterator struct {
	Event *FixedSupplyBankERC20Approval // Event containing the contract specifics and raw log

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
func (it *FixedSupplyBankERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedSupplyBankERC20Approval)
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
		it.Event = new(FixedSupplyBankERC20Approval)
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
func (it *FixedSupplyBankERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedSupplyBankERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedSupplyBankERC20Approval represents a Approval event raised by the FixedSupplyBankERC20 contract.
type FixedSupplyBankERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FixedSupplyBankERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FixedSupplyBankERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyBankERC20ApprovalIterator{contract: _FixedSupplyBankERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FixedSupplyBankERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FixedSupplyBankERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedSupplyBankERC20Approval)
				if err := _FixedSupplyBankERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Filterer) ParseApproval(log types.Log) (*FixedSupplyBankERC20Approval, error) {
	event := new(FixedSupplyBankERC20Approval)
	if err := _FixedSupplyBankERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FixedSupplyBankERC20FailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the FixedSupplyBankERC20 contract.
type FixedSupplyBankERC20FailureIterator struct {
	Event *FixedSupplyBankERC20Failure // Event containing the contract specifics and raw log

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
func (it *FixedSupplyBankERC20FailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedSupplyBankERC20Failure)
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
		it.Event = new(FixedSupplyBankERC20Failure)
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
func (it *FixedSupplyBankERC20FailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedSupplyBankERC20FailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedSupplyBankERC20Failure represents a Failure event raised by the FixedSupplyBankERC20 contract.
type FixedSupplyBankERC20Failure struct {
	Message string
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x66c9257b5635d9c11609ab746e0972276ff2412ab2085de9630ecb2300a019a6.
//
// Solidity: event Failure(string message, bytes data)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Filterer) FilterFailure(opts *bind.FilterOpts) (*FixedSupplyBankERC20FailureIterator, error) {

	logs, sub, err := _FixedSupplyBankERC20.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &FixedSupplyBankERC20FailureIterator{contract: _FixedSupplyBankERC20.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x66c9257b5635d9c11609ab746e0972276ff2412ab2085de9630ecb2300a019a6.
//
// Solidity: event Failure(string message, bytes data)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Filterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *FixedSupplyBankERC20Failure) (event.Subscription, error) {

	logs, sub, err := _FixedSupplyBankERC20.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedSupplyBankERC20Failure)
				if err := _FixedSupplyBankERC20.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x66c9257b5635d9c11609ab746e0972276ff2412ab2085de9630ecb2300a019a6.
//
// Solidity: event Failure(string message, bytes data)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Filterer) ParseFailure(log types.Log) (*FixedSupplyBankERC20Failure, error) {
	event := new(FixedSupplyBankERC20Failure)
	if err := _FixedSupplyBankERC20.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FixedSupplyBankERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FixedSupplyBankERC20 contract.
type FixedSupplyBankERC20TransferIterator struct {
	Event *FixedSupplyBankERC20Transfer // Event containing the contract specifics and raw log

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
func (it *FixedSupplyBankERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedSupplyBankERC20Transfer)
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
		it.Event = new(FixedSupplyBankERC20Transfer)
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
func (it *FixedSupplyBankERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedSupplyBankERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedSupplyBankERC20Transfer represents a Transfer event raised by the FixedSupplyBankERC20 contract.
type FixedSupplyBankERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FixedSupplyBankERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FixedSupplyBankERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyBankERC20TransferIterator{contract: _FixedSupplyBankERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FixedSupplyBankERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FixedSupplyBankERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedSupplyBankERC20Transfer)
				if err := _FixedSupplyBankERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FixedSupplyBankERC20 *FixedSupplyBankERC20Filterer) ParseTransfer(log types.Log) (*FixedSupplyBankERC20Transfer, error) {
	event := new(FixedSupplyBankERC20Transfer)
	if err := _FixedSupplyBankERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
