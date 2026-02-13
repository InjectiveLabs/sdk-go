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

// MintBurnBankERC20MetaData contains all meta data concerning the MintBurnBankERC20 contract.
var MintBurnBankERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Failure\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60806040819052600680546001600160a01b03191660641790556113c23881900390819083398101604081905261003591610257565b60408051602080820183525f8083528351918201909352918252849184918491886001600160a01b03811661008357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61008c8161014c565b5060046100998382610375565b5060056100a68282610375565b5050505f835111806100b857505f8251115b806100c557505f8160ff16115b1561014057600654604051630df4b0bd60e21b81526001600160a01b03909116906337d2c2f4906100fe9086908690869060040161045a565b6020604051808303815f875af115801561011a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061013e9190610492565b505b505050505050506104b8565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156101c95781810151838201526020016101b1565b50505f910152565b5f82601f8301126101e0575f5ffd5b81516001600160401b038111156101f9576101f961019b565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102275761022761019b565b60405281815283820160200185101561023e575f5ffd5b61024f8260208301602087016101af565b949350505050565b5f5f5f5f6080858703121561026a575f5ffd5b84516001600160a01b0381168114610280575f5ffd5b60208601519094506001600160401b0381111561029b575f5ffd5b6102a7878288016101d1565b604087015190945090506001600160401b038111156102c4575f5ffd5b6102d0878288016101d1565b925050606085015160ff811681146102e6575f5ffd5b939692955090935050565b600181811c9082168061030557607f821691505b60208210810361032357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561037057805f5260205f20601f840160051c8101602085101561034e5750805b601f840160051c820191505b8181101561036d575f815560010161035a565b50505b505050565b81516001600160401b0381111561038e5761038e61019b565b6103a28161039c84546102f1565b84610329565b6020601f8211600181146103d4575f83156103bd5750848201515b5f19600385901b1c1916600184901b17845561036d565b5f84815260208120601f198516915b8281101561040357878501518255602094850194600190920191016103e3565b508482101561042057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f81518084526104468160208601602086016101af565b601f01601f19169290920160200192915050565b606081525f61046c606083018661042f565b828103602084015261047e818661042f565b91505060ff83166040830152949350505050565b5f602082840312156104a2575f5ffd5b815180151581146104b1575f5ffd5b9392505050565b610efd806104c55f395ff3fe6080604052600436106100e4575f3560e01c806370a082311161008757806395d89b411161005757806395d89b4114610254578063a9059cbb14610268578063dd62ed3e14610287578063f2fde38b146102cb575f5ffd5b806370a08231146101dc578063715018a6146101fb57806379cc67901461020f5780638da5cb5b1461022e575f5ffd5b806323b872dd116100c257806323b872dd14610163578063313ce5671461018257806340c10f19146101a857806342966c68146101bd575f5ffd5b806306fdde03146100e8578063095ea7b31461011257806318160ddd14610141575b5f5ffd5b3480156100f3575f5ffd5b506100fc6102ea565b6040516101099190610b47565b60405180910390f35b34801561011d575f5ffd5b5061013161012c366004610b94565b610364565b6040519015158152602001610109565b34801561014c575f5ffd5b5061015561037b565b604051908152602001610109565b34801561016e575f5ffd5b5061013161017d366004610bbc565b6103ea565b34801561018d575f5ffd5b5061019661040d565b60405160ff9091168152602001610109565b6101bb6101b6366004610b94565b610485565b005b3480156101c8575f5ffd5b506101bb6101d7366004610bf6565b61049b565b3480156101e7575f5ffd5b506101556101f6366004610c0d565b6104a8565b348015610206575f5ffd5b506101bb610520565b34801561021a575f5ffd5b506101bb610229366004610b94565b610533565b348015610239575f5ffd5b505f546040516001600160a01b039091168152602001610109565b34801561025f575f5ffd5b506100fc610548565b348015610273575f5ffd5b50610131610282366004610b94565b6105c1565b348015610292575f5ffd5b506101556102a1366004610c2d565b6001600160a01b039182165f90815260026020908152604080832093909416825291909152205490565b3480156102d6575f5ffd5b506101bb6102e5366004610c0d565b6105ce565b6006546040516315d10ab960e11b815230600482015260609182916001600160a01b0390911690632ba21572906024015f60405180830381865afa158015610334573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261035b9190810190610d04565b50909392505050565b5f3361037181858561060d565b5060019392505050565b6006546040516339370aa960e21b81523060048201525f916001600160a01b03169063e4dc2aa490602401602060405180830381865afa1580156103c1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103e59190610d81565b905090565b5f336103f785828561061f565b61040285858561069b565b506001949350505050565b6006546040516315d10ab960e11b81523060048201525f9182916001600160a01b0390911690632ba21572906024015f60405180830381865afa158015610456573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261047d9190810190610d04565b949350505050565b61048d6106f8565b6104978282610724565b5050565b6104a53382610758565b50565b600654604051633de222bb60e21b81523060048201526001600160a01b0383811660248301525f92169063f7888aec90604401602060405180830381865afa1580156104f6573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061051a9190610d81565b92915050565b6105286106f8565b6105315f61078c565b565b61053e82338361061f565b6104978282610758565b6006546040516315d10ab960e11b815230600482015260609182916001600160a01b0390911690632ba21572906024015f60405180830381865afa158015610592573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105b99190810190610d04565b509392505050565b5f3361037181858561069b565b6105d66106f8565b6001600160a01b03811661060457604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6104a58161078c565b61061a83838360016107db565b505050565b6001600160a01b038381165f908152600260209081526040808320938616835292905220545f19811015610695578181101561068757604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016105fb565b61069584848484035f6107db565b50505050565b6001600160a01b0383166106c457604051634b637e8f60e11b81525f60048201526024016105fb565b6001600160a01b0382166106ed5760405163ec442f0560e01b81525f60048201526024016105fb565b61061a8383836108ad565b5f546001600160a01b031633146105315760405163118cdaa760e01b81523360048201526024016105fb565b6001600160a01b03821661074d5760405163ec442f0560e01b81525f60048201526024016105fb565b6104975f83836108ad565b6001600160a01b03821661078157604051634b637e8f60e11b81525f60048201526024016105fb565b610497825f836108ad565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0384166108045760405163e602df0560e01b81525f60048201526024016105fb565b6001600160a01b03831661082d57604051634a1406b160e11b81525f60048201526024016105fb565b6001600160a01b038085165f908152600260209081526040808320938716835292905220829055801561069557826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161089f91815260200190565b60405180910390a350505050565b6001600160a01b0383166109d3576006546040516340c10f1960e01b81526001600160a01b03848116600483015260248201849052909116906340c10f19906044016020604051808303815f875af1925050508015610929575060408051601f3d908101601f1916820190925261092691810190610d98565b60015b6109cd57610935610db7565b806308c379a00361098b5750610949610dd0565b80610954575061098d565b806040516020016109659190610e4c565b60408051601f198184030181529082905262461bcd60e51b82526105fb91600401610b47565b505b3d8080156109b6576040519150601f19603f3d011682016040523d82523d5f602084013e6109bb565b606091505b50806040516020016109659190610e83565b50610ad3565b6001600160a01b038216610a5557600654604051632770a7eb60e21b81526001600160a01b0385811660048301526024820184905290911690639dc29fac906044016020604051808303815f875af1158015610a31573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109cd9190610d98565b6006546040516317d5759960e31b81526001600160a01b0385811660048301528481166024830152604482018490529091169063beabacc8906064016020604051808303815f875af1158015610aad573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ad19190610d98565b505b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610b1891815260200190565b60405180910390a3505050565b5f5b83811015610b3f578181015183820152602001610b27565b50505f910152565b602081525f8251806020840152610b65816040850160208701610b25565b601f01601f19169190910160400192915050565b80356001600160a01b0381168114610b8f575f5ffd5b919050565b5f5f60408385031215610ba5575f5ffd5b610bae83610b79565b946020939093013593505050565b5f5f5f60608486031215610bce575f5ffd5b610bd784610b79565b9250610be560208501610b79565b929592945050506040919091013590565b5f60208284031215610c06575f5ffd5b5035919050565b5f60208284031215610c1d575f5ffd5b610c2682610b79565b9392505050565b5f5f60408385031215610c3e575f5ffd5b610c4783610b79565b9150610c5560208401610b79565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b601f8201601f1916810167ffffffffffffffff81118282101715610c9857610c98610c5e565b6040525050565b5f82601f830112610cae575f5ffd5b815167ffffffffffffffff811115610cc857610cc8610c5e565b604051610cdf601f8301601f191660200182610c72565b818152846020838601011115610cf3575f5ffd5b61047d826020830160208701610b25565b5f5f5f60608486031215610d16575f5ffd5b835167ffffffffffffffff811115610d2c575f5ffd5b610d3886828701610c9f565b935050602084015167ffffffffffffffff811115610d54575f5ffd5b610d6086828701610c9f565b925050604084015160ff81168114610d76575f5ffd5b809150509250925092565b5f60208284031215610d91575f5ffd5b5051919050565b5f60208284031215610da8575f5ffd5b81518015158114610c26575f5ffd5b5f60033d1115610dcd5760045f5f3e505f5160e01c5b90565b5f60443d1015610ddd5790565b6040513d600319016004823e80513d602482011167ffffffffffffffff82111715610e0757505090565b808201805167ffffffffffffffff811115610e23575050505090565b3d8401600319018282016020011115610e3d575050505090565b6105b960208285010185610c72565b6f03330b4b632b2103a379036b4b73a1d160851b81525f8251610e76816010850160208701610b25565b9190910160100192915050565b7f6661696c656420746f206d696e743a20756e6b6e6f776e206572726f723a200081525f8251610eba81601f850160208701610b25565b91909101601f019291505056fea26469706673582212208bc73576e871a975e4d1b36ec8d562772efff0e5efcd79595bd22762d64b42c864736f6c634300081e0033",
}

// MintBurnBankERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use MintBurnBankERC20MetaData.ABI instead.
var MintBurnBankERC20ABI = MintBurnBankERC20MetaData.ABI

// MintBurnBankERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MintBurnBankERC20MetaData.Bin instead.
var MintBurnBankERC20Bin = MintBurnBankERC20MetaData.Bin

// DeployMintBurnBankERC20 deploys a new Ethereum contract, binding an instance of MintBurnBankERC20 to it.
func DeployMintBurnBankERC20(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address, name_ string, symbol_ string, decimals_ uint8) (common.Address, *types.Transaction, *MintBurnBankERC20, error) {
	parsed, err := MintBurnBankERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MintBurnBankERC20Bin), backend, initialOwner, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MintBurnBankERC20{MintBurnBankERC20Caller: MintBurnBankERC20Caller{contract: contract}, MintBurnBankERC20Transactor: MintBurnBankERC20Transactor{contract: contract}, MintBurnBankERC20Filterer: MintBurnBankERC20Filterer{contract: contract}}, nil
}

// MintBurnBankERC20 is an auto generated Go binding around an Ethereum contract.
type MintBurnBankERC20 struct {
	MintBurnBankERC20Caller     // Read-only binding to the contract
	MintBurnBankERC20Transactor // Write-only binding to the contract
	MintBurnBankERC20Filterer   // Log filterer for contract events
}

// MintBurnBankERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type MintBurnBankERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintBurnBankERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MintBurnBankERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintBurnBankERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintBurnBankERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintBurnBankERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintBurnBankERC20Session struct {
	Contract     *MintBurnBankERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MintBurnBankERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintBurnBankERC20CallerSession struct {
	Contract *MintBurnBankERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MintBurnBankERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintBurnBankERC20TransactorSession struct {
	Contract     *MintBurnBankERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MintBurnBankERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type MintBurnBankERC20Raw struct {
	Contract *MintBurnBankERC20 // Generic contract binding to access the raw methods on
}

// MintBurnBankERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintBurnBankERC20CallerRaw struct {
	Contract *MintBurnBankERC20Caller // Generic read-only contract binding to access the raw methods on
}

// MintBurnBankERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintBurnBankERC20TransactorRaw struct {
	Contract *MintBurnBankERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMintBurnBankERC20 creates a new instance of MintBurnBankERC20, bound to a specific deployed contract.
func NewMintBurnBankERC20(address common.Address, backend bind.ContractBackend) (*MintBurnBankERC20, error) {
	contract, err := bindMintBurnBankERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintBurnBankERC20{MintBurnBankERC20Caller: MintBurnBankERC20Caller{contract: contract}, MintBurnBankERC20Transactor: MintBurnBankERC20Transactor{contract: contract}, MintBurnBankERC20Filterer: MintBurnBankERC20Filterer{contract: contract}}, nil
}

// NewMintBurnBankERC20Caller creates a new read-only instance of MintBurnBankERC20, bound to a specific deployed contract.
func NewMintBurnBankERC20Caller(address common.Address, caller bind.ContractCaller) (*MintBurnBankERC20Caller, error) {
	contract, err := bindMintBurnBankERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintBurnBankERC20Caller{contract: contract}, nil
}

// NewMintBurnBankERC20Transactor creates a new write-only instance of MintBurnBankERC20, bound to a specific deployed contract.
func NewMintBurnBankERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*MintBurnBankERC20Transactor, error) {
	contract, err := bindMintBurnBankERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintBurnBankERC20Transactor{contract: contract}, nil
}

// NewMintBurnBankERC20Filterer creates a new log filterer instance of MintBurnBankERC20, bound to a specific deployed contract.
func NewMintBurnBankERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*MintBurnBankERC20Filterer, error) {
	contract, err := bindMintBurnBankERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintBurnBankERC20Filterer{contract: contract}, nil
}

// bindMintBurnBankERC20 binds a generic wrapper to an already deployed contract.
func bindMintBurnBankERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MintBurnBankERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintBurnBankERC20 *MintBurnBankERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintBurnBankERC20.Contract.MintBurnBankERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintBurnBankERC20 *MintBurnBankERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.MintBurnBankERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintBurnBankERC20 *MintBurnBankERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.MintBurnBankERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintBurnBankERC20 *MintBurnBankERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintBurnBankERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MintBurnBankERC20 *MintBurnBankERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintBurnBankERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MintBurnBankERC20.Contract.Allowance(&_MintBurnBankERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MintBurnBankERC20 *MintBurnBankERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MintBurnBankERC20.Contract.Allowance(&_MintBurnBankERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MintBurnBankERC20 *MintBurnBankERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintBurnBankERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _MintBurnBankERC20.Contract.BalanceOf(&_MintBurnBankERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MintBurnBankERC20 *MintBurnBankERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MintBurnBankERC20.Contract.BalanceOf(&_MintBurnBankERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MintBurnBankERC20 *MintBurnBankERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MintBurnBankERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Decimals() (uint8, error) {
	return _MintBurnBankERC20.Contract.Decimals(&_MintBurnBankERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MintBurnBankERC20 *MintBurnBankERC20CallerSession) Decimals() (uint8, error) {
	return _MintBurnBankERC20.Contract.Decimals(&_MintBurnBankERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MintBurnBankERC20 *MintBurnBankERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MintBurnBankERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Name() (string, error) {
	return _MintBurnBankERC20.Contract.Name(&_MintBurnBankERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MintBurnBankERC20 *MintBurnBankERC20CallerSession) Name() (string, error) {
	return _MintBurnBankERC20.Contract.Name(&_MintBurnBankERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintBurnBankERC20 *MintBurnBankERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintBurnBankERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Owner() (common.Address, error) {
	return _MintBurnBankERC20.Contract.Owner(&_MintBurnBankERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintBurnBankERC20 *MintBurnBankERC20CallerSession) Owner() (common.Address, error) {
	return _MintBurnBankERC20.Contract.Owner(&_MintBurnBankERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MintBurnBankERC20 *MintBurnBankERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MintBurnBankERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Symbol() (string, error) {
	return _MintBurnBankERC20.Contract.Symbol(&_MintBurnBankERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MintBurnBankERC20 *MintBurnBankERC20CallerSession) Symbol() (string, error) {
	return _MintBurnBankERC20.Contract.Symbol(&_MintBurnBankERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintBurnBankERC20 *MintBurnBankERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintBurnBankERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) TotalSupply() (*big.Int, error) {
	return _MintBurnBankERC20.Contract.TotalSupply(&_MintBurnBankERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintBurnBankERC20 *MintBurnBankERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _MintBurnBankERC20.Contract.TotalSupply(&_MintBurnBankERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MintBurnBankERC20 *MintBurnBankERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.Approve(&_MintBurnBankERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.Approve(&_MintBurnBankERC20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Transactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Burn(value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.Burn(&_MintBurnBankERC20.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.Burn(&_MintBurnBankERC20.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Session) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.BurnFrom(&_MintBurnBankERC20.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.BurnFrom(&_MintBurnBankERC20.TransactOpts, account, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) payable returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) payable returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.Mint(&_MintBurnBankERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) payable returns()
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.Mint(&_MintBurnBankERC20.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintBurnBankERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.RenounceOwnership(&_MintBurnBankERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.RenounceOwnership(&_MintBurnBankERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MintBurnBankERC20 *MintBurnBankERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.Transfer(&_MintBurnBankERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.Transfer(&_MintBurnBankERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MintBurnBankERC20 *MintBurnBankERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MintBurnBankERC20 *MintBurnBankERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.TransferFrom(&_MintBurnBankERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.TransferFrom(&_MintBurnBankERC20.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MintBurnBankERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.TransferOwnership(&_MintBurnBankERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.TransferOwnership(&_MintBurnBankERC20.TransactOpts, newOwner)
}

// MintBurnBankERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MintBurnBankERC20 contract.
type MintBurnBankERC20ApprovalIterator struct {
	Event *MintBurnBankERC20Approval // Event containing the contract specifics and raw log

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
func (it *MintBurnBankERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintBurnBankERC20Approval)
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
		it.Event = new(MintBurnBankERC20Approval)
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
func (it *MintBurnBankERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintBurnBankERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintBurnBankERC20Approval represents a Approval event raised by the MintBurnBankERC20 contract.
type MintBurnBankERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MintBurnBankERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintBurnBankERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MintBurnBankERC20ApprovalIterator{contract: _MintBurnBankERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MintBurnBankERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintBurnBankERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintBurnBankERC20Approval)
				if err := _MintBurnBankERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) ParseApproval(log types.Log) (*MintBurnBankERC20Approval, error) {
	event := new(MintBurnBankERC20Approval)
	if err := _MintBurnBankERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintBurnBankERC20FailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the MintBurnBankERC20 contract.
type MintBurnBankERC20FailureIterator struct {
	Event *MintBurnBankERC20Failure // Event containing the contract specifics and raw log

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
func (it *MintBurnBankERC20FailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintBurnBankERC20Failure)
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
		it.Event = new(MintBurnBankERC20Failure)
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
func (it *MintBurnBankERC20FailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintBurnBankERC20FailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintBurnBankERC20Failure represents a Failure event raised by the MintBurnBankERC20 contract.
type MintBurnBankERC20Failure struct {
	Message string
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x66c9257b5635d9c11609ab746e0972276ff2412ab2085de9630ecb2300a019a6.
//
// Solidity: event Failure(string message, bytes data)
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) FilterFailure(opts *bind.FilterOpts) (*MintBurnBankERC20FailureIterator, error) {

	logs, sub, err := _MintBurnBankERC20.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &MintBurnBankERC20FailureIterator{contract: _MintBurnBankERC20.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x66c9257b5635d9c11609ab746e0972276ff2412ab2085de9630ecb2300a019a6.
//
// Solidity: event Failure(string message, bytes data)
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *MintBurnBankERC20Failure) (event.Subscription, error) {

	logs, sub, err := _MintBurnBankERC20.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintBurnBankERC20Failure)
				if err := _MintBurnBankERC20.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) ParseFailure(log types.Log) (*MintBurnBankERC20Failure, error) {
	event := new(MintBurnBankERC20Failure)
	if err := _MintBurnBankERC20.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintBurnBankERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MintBurnBankERC20 contract.
type MintBurnBankERC20OwnershipTransferredIterator struct {
	Event *MintBurnBankERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MintBurnBankERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintBurnBankERC20OwnershipTransferred)
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
		it.Event = new(MintBurnBankERC20OwnershipTransferred)
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
func (it *MintBurnBankERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintBurnBankERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintBurnBankERC20OwnershipTransferred represents a OwnershipTransferred event raised by the MintBurnBankERC20 contract.
type MintBurnBankERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintBurnBankERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintBurnBankERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintBurnBankERC20OwnershipTransferredIterator{contract: _MintBurnBankERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MintBurnBankERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintBurnBankERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintBurnBankERC20OwnershipTransferred)
				if err := _MintBurnBankERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) ParseOwnershipTransferred(log types.Log) (*MintBurnBankERC20OwnershipTransferred, error) {
	event := new(MintBurnBankERC20OwnershipTransferred)
	if err := _MintBurnBankERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintBurnBankERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MintBurnBankERC20 contract.
type MintBurnBankERC20TransferIterator struct {
	Event *MintBurnBankERC20Transfer // Event containing the contract specifics and raw log

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
func (it *MintBurnBankERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintBurnBankERC20Transfer)
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
		it.Event = new(MintBurnBankERC20Transfer)
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
func (it *MintBurnBankERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintBurnBankERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintBurnBankERC20Transfer represents a Transfer event raised by the MintBurnBankERC20 contract.
type MintBurnBankERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MintBurnBankERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintBurnBankERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MintBurnBankERC20TransferIterator{contract: _MintBurnBankERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MintBurnBankERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintBurnBankERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintBurnBankERC20Transfer)
				if err := _MintBurnBankERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MintBurnBankERC20 *MintBurnBankERC20Filterer) ParseTransfer(log types.Log) (*MintBurnBankERC20Transfer, error) {
	event := new(MintBurnBankERC20Transfer)
	if err := _MintBurnBankERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
