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
	Bin: "0x6080604052606460055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040516123e63803806123e6833981810160405281019061006691906107a9565b83838360405180602001604052805f81525060405180602001604052805f81525081600390816100969190610a4c565b5080600490816100a69190610a4c565b5050505f835111806100b857505f8251115b806100c557505f8160ff16115b1561016a5760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337d2c2f48484846040518463ffffffff1660e01b815260040161012893929190610b72565b6020604051808303815f875af1158015610144573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101689190610bea565b505b5050505f81111561018657610185338261018f60201b60201c565b5b50505050610e9b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036101ff575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016101f69190610c54565b60405180910390fd5b6102105f838361021460201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036103f95760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b81526004016102a3929190610c7c565b6020604051808303815f875af19250505080156102de57506040513d601f19601f820116820180604052508101906102db9190610bea565b60015b6103f3576102ea610caf565b806308c379a00361036557506102fe610cce565b806103095750610367565b8060405160200161031a9190610dbd565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035c9190610de2565b60405180910390fd5b505b3d805f8114610391576040519150601f19603f3d011682016040523d82523d5f602084013e610396565b606091505b50806040516020016103a89190610e28565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ea9190610de2565b60405180910390fd5b5061056f565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036104ce5760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac84836040518363ffffffff1660e01b8152600401610488929190610c7c565b6020604051808303815f875af11580156104a4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104c89190610bea565b5061056e565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663beabacc88484846040518463ffffffff1660e01b815260040161052c93929190610e4d565b6020604051808303815f875af1158015610548573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061056c9190610bea565b505b5b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516105cc9190610e82565b60405180910390a3505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610638826105f2565b810181811067ffffffffffffffff8211171561065757610656610602565b5b80604052505050565b5f6106696105d9565b9050610675828261062f565b919050565b5f67ffffffffffffffff82111561069457610693610602565b5b61069d826105f2565b9050602081019050919050565b5f5b838110156106c75780820151818401526020810190506106ac565b5f8484015250505050565b5f6106e46106df8461067a565b610660565b905082815260208101848484011115610700576106ff6105ee565b5b61070b8482856106aa565b509392505050565b5f82601f830112610727576107266105ea565b5b81516107378482602086016106d2565b91505092915050565b5f60ff82169050919050565b61075581610740565b811461075f575f5ffd5b50565b5f815190506107708161074c565b92915050565b5f819050919050565b61078881610776565b8114610792575f5ffd5b50565b5f815190506107a38161077f565b92915050565b5f5f5f5f608085870312156107c1576107c06105e2565b5b5f85015167ffffffffffffffff8111156107de576107dd6105e6565b5b6107ea87828801610713565b945050602085015167ffffffffffffffff81111561080b5761080a6105e6565b5b61081787828801610713565b935050604061082887828801610762565b925050606061083987828801610795565b91505092959194509250565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061089357607f821691505b6020821081036108a6576108a561084f565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026109087fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826108cd565b61091286836108cd565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61094d61094861094384610776565b61092a565b610776565b9050919050565b5f819050919050565b61096683610933565b61097a61097282610954565b8484546108d9565b825550505050565b5f5f905090565b610991610982565b61099c81848461095d565b505050565b5b818110156109bf576109b45f82610989565b6001810190506109a2565b5050565b601f821115610a04576109d5816108ac565b6109de846108be565b810160208510156109ed578190505b610a016109f9856108be565b8301826109a1565b50505b505050565b5f82821c905092915050565b5f610a245f1984600802610a09565b1980831691505092915050565b5f610a3c8383610a15565b9150826002028217905092915050565b610a5582610845565b67ffffffffffffffff811115610a6e57610a6d610602565b5b610a78825461087c565b610a838282856109c3565b5f60209050601f831160018114610ab4575f8415610aa2578287015190505b610aac8582610a31565b865550610b13565b601f198416610ac2866108ac565b5f5b82811015610ae957848901518255600182019150602085019450602081019050610ac4565b86831015610b065784890151610b02601f891682610a15565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f610b3582610845565b610b3f8185610b1b565b9350610b4f8185602086016106aa565b610b58816105f2565b840191505092915050565b610b6c81610740565b82525050565b5f6060820190508181035f830152610b8a8186610b2b565b90508181036020830152610b9e8185610b2b565b9050610bad6040830184610b63565b949350505050565b5f8115159050919050565b610bc981610bb5565b8114610bd3575f5ffd5b50565b5f81519050610be481610bc0565b92915050565b5f60208284031215610bff57610bfe6105e2565b5b5f610c0c84828501610bd6565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610c3e82610c15565b9050919050565b610c4e81610c34565b82525050565b5f602082019050610c675f830184610c45565b92915050565b610c7681610776565b82525050565b5f604082019050610c8f5f830185610c45565b610c9c6020830184610c6d565b9392505050565b5f8160e01c9050919050565b5f60033d1115610ccb5760045f5f3e610cc85f51610ca3565b90505b90565b5f60443d10610d5a57610cdf6105d9565b60043d036004823e80513d602482011167ffffffffffffffff82111715610d07575050610d5a565b808201805167ffffffffffffffff811115610d255750505050610d5a565b80602083010160043d038501811115610d42575050505050610d5a565b610d518260200185018661062f565b82955050505050505b90565b7f6661696c656420746f206d696e743a2000000000000000000000000000000000815250565b5f81905092915050565b5f610d9782610845565b610da18185610d83565b9350610db18185602086016106aa565b80840191505092915050565b5f610dc782610d5d565b601082019150610dd78284610d8d565b915081905092915050565b5f6020820190508181035f830152610dfa8184610b2b565b905092915050565b7f6661696c656420746f206d696e743a20756e6b6e6f776e206572726f723a2000815250565b5f610e3282610e02565b601f82019150610e428284610d8d565b915081905092915050565b5f606082019050610e605f830186610c45565b610e6d6020830185610c45565b610e7a6040830184610c6d565b949350505050565b5f602082019050610e955f830184610c6d565b92915050565b61153e80610ea85f395ff3fe608060405234801561000f575f5ffd5b5060043610610091575f3560e01c8063313ce56711610064578063313ce5671461013157806370a082311461014f57806395d89b411461017f578063a9059cbb1461019d578063dd62ed3e146101cd57610091565b806306fdde0314610095578063095ea7b3146100b357806318160ddd146100e357806323b872dd14610101575b5f5ffd5b61009d6101fd565b6040516100aa9190610df6565b60405180910390f35b6100cd60048036038101906100c89190610eb4565b6102aa565b6040516100da9190610f0c565b60405180910390f35b6100eb6102cc565b6040516100f89190610f34565b60405180910390f35b61011b60048036038101906101169190610f4d565b61036b565b6040516101289190610f0c565b60405180910390f35b610139610399565b6040516101469190610fb8565b60405180910390f35b61016960048036038101906101649190610fd1565b610447565b6040516101769190610f34565b60405180910390f35b6101876104ea565b6040516101949190610df6565b60405180910390f35b6101b760048036038101906101b29190610eb4565b610598565b6040516101c49190610f0c565b60405180910390f35b6101e760048036038101906101e29190610ffc565b6105ba565b6040516101f49190610f34565b60405180910390f35b60608060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b815260040161025a9190611049565b5f60405180830381865afa158015610274573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061029c91906111aa565b905050809150508091505090565b5f5f6102b461063c565b90506102c1818585610643565b600191505092915050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4dc2aa4306040518263ffffffff1660e01b81526004016103279190611049565b602060405180830381865afa158015610342573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103669190611246565b905090565b5f5f61037561063c565b9050610382858285610655565b61038d8585856106e8565b60019150509392505050565b5f5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b81526004016103f59190611049565b5f60405180830381865afa15801561040f573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061043791906111aa565b9091509050809150508091505090565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f7888aec30846040518363ffffffff1660e01b81526004016104a4929190611271565b602060405180830381865afa1580156104bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104e39190611246565b9050919050565b60608060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b81526004016105479190611049565b5f60405180830381865afa158015610561573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061058991906111aa565b90915050809150508091505090565b5f5f6105a261063c565b90506105af8185856106e8565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f33905090565b61065083838360016107d8565b505050565b5f61066084846105ba565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156106e257818110156106d3578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016106ca93929190611298565b60405180910390fd5b6106e184848484035f6107d8565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610758575f6040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260040161074f9190611049565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107c8575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016107bf9190611049565b60405180910390fd5b6107d38383836109a7565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610848575f6040517fe602df0500000000000000000000000000000000000000000000000000000000815260040161083f9190611049565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108b8575f6040517f94280d620000000000000000000000000000000000000000000000000000000081526004016108af9190611049565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555080156109a1578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516109989190610f34565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b8c5760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b8152600401610a369291906112cd565b6020604051808303815f875af1925050508015610a7157506040513d601f19601f82011682018060405250810190610a6e919061131e565b60015b610b8657610a7d611355565b806308c379a003610af85750610a91611374565b80610a9c5750610afa565b80604051602001610aad9190611463565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aef9190610df6565b60405180910390fd5b505b3d805f8114610b24576040519150601f19603f3d011682016040523d82523d5f602084013e610b29565b606091505b5080604051602001610b3b91906114ae565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7d9190610df6565b60405180910390fd5b50610d02565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c615760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac84836040518363ffffffff1660e01b8152600401610c1b9291906112cd565b6020604051808303815f875af1158015610c37573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c5b919061131e565b50610d01565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663beabacc88484846040518463ffffffff1660e01b8152600401610cbf939291906114d3565b6020604051808303815f875af1158015610cdb573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cff919061131e565b505b5b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d5f9190610f34565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610da3578082015181840152602081019050610d88565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610dc882610d6c565b610dd28185610d76565b9350610de2818560208601610d86565b610deb81610dae565b840191505092915050565b5f6020820190508181035f830152610e0e8184610dbe565b905092915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e5082610e27565b9050919050565b610e6081610e46565b8114610e6a575f5ffd5b50565b5f81359050610e7b81610e57565b92915050565b5f819050919050565b610e9381610e81565b8114610e9d575f5ffd5b50565b5f81359050610eae81610e8a565b92915050565b5f5f60408385031215610eca57610ec9610e1f565b5b5f610ed785828601610e6d565b9250506020610ee885828601610ea0565b9150509250929050565b5f8115159050919050565b610f0681610ef2565b82525050565b5f602082019050610f1f5f830184610efd565b92915050565b610f2e81610e81565b82525050565b5f602082019050610f475f830184610f25565b92915050565b5f5f5f60608486031215610f6457610f63610e1f565b5b5f610f7186828701610e6d565b9350506020610f8286828701610e6d565b9250506040610f9386828701610ea0565b9150509250925092565b5f60ff82169050919050565b610fb281610f9d565b82525050565b5f602082019050610fcb5f830184610fa9565b92915050565b5f60208284031215610fe657610fe5610e1f565b5b5f610ff384828501610e6d565b91505092915050565b5f5f6040838503121561101257611011610e1f565b5b5f61101f85828601610e6d565b925050602061103085828601610e6d565b9150509250929050565b61104381610e46565b82525050565b5f60208201905061105c5f83018461103a565b92915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6110a082610dae565b810181811067ffffffffffffffff821117156110bf576110be61106a565b5b80604052505050565b5f6110d1610e16565b90506110dd8282611097565b919050565b5f67ffffffffffffffff8211156110fc576110fb61106a565b5b61110582610dae565b9050602081019050919050565b5f61112461111f846110e2565b6110c8565b9050828152602081018484840111156111405761113f611066565b5b61114b848285610d86565b509392505050565b5f82601f83011261116757611166611062565b5b8151611177848260208601611112565b91505092915050565b61118981610f9d565b8114611193575f5ffd5b50565b5f815190506111a481611180565b92915050565b5f5f5f606084860312156111c1576111c0610e1f565b5b5f84015167ffffffffffffffff8111156111de576111dd610e23565b5b6111ea86828701611153565b935050602084015167ffffffffffffffff81111561120b5761120a610e23565b5b61121786828701611153565b925050604061122886828701611196565b9150509250925092565b5f8151905061124081610e8a565b92915050565b5f6020828403121561125b5761125a610e1f565b5b5f61126884828501611232565b91505092915050565b5f6040820190506112845f83018561103a565b611291602083018461103a565b9392505050565b5f6060820190506112ab5f83018661103a565b6112b86020830185610f25565b6112c56040830184610f25565b949350505050565b5f6040820190506112e05f83018561103a565b6112ed6020830184610f25565b9392505050565b6112fd81610ef2565b8114611307575f5ffd5b50565b5f81519050611318816112f4565b92915050565b5f6020828403121561133357611332610e1f565b5b5f6113408482850161130a565b91505092915050565b5f8160e01c9050919050565b5f60033d11156113715760045f5f3e61136e5f51611349565b90505b90565b5f60443d1061140057611385610e16565b60043d036004823e80513d602482011167ffffffffffffffff821117156113ad575050611400565b808201805167ffffffffffffffff8111156113cb5750505050611400565b80602083010160043d0385018111156113e8575050505050611400565b6113f782602001850186611097565b82955050505050505b90565b7f6661696c656420746f206d696e743a2000000000000000000000000000000000815250565b5f81905092915050565b5f61143d82610d6c565b6114478185611429565b9350611457818560208601610d86565b80840191505092915050565b5f61146d82611403565b60108201915061147d8284611433565b915081905092915050565b7f6661696c656420746f206d696e743a20756e6b6e6f776e206572726f723a2000815250565b5f6114b882611488565b601f820191506114c88284611433565b915081905092915050565b5f6060820190506114e65f83018661103a565b6114f3602083018561103a565b6115006040830184610f25565b94935050505056fea2646970667358221220c69c0105de6f868fb2241b8ef2c7ca8a58124df477c32bdc906718a163cf478b64736f6c634300081b0033",
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
