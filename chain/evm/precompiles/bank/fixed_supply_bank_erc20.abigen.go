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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"initial_supply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052606460055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604051611e70380380611e7083398181016040528101906100669190610677565b83838360405180602001604052805f81525060405180602001604052805f8152508160039081610096919061091a565b5080600490816100a6919061091a565b50505060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337d2c2f48484846040518463ffffffff1660e01b815260040161010793929190610a40565b6020604051808303815f875af1158015610123573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101479190610ab8565b505050505f81111561016457610163338261016d60201b60201c565b5b50505050610bbf565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036101dd575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016101d49190610b22565b60405180910390fd5b6101ee5f83836101f260201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036102c75760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b8152600401610281929190610b4a565b6020604051808303815f875af115801561029d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102c19190610ab8565b5061043d565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361039c5760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac84836040518363ffffffff1660e01b8152600401610356929190610b4a565b6020604051808303815f875af1158015610372573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103969190610ab8565b5061043c565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663beabacc88484846040518463ffffffff1660e01b81526004016103fa93929190610b71565b6020604051808303815f875af1158015610416573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061043a9190610ab8565b505b5b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161049a9190610ba6565b60405180910390a3505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610506826104c0565b810181811067ffffffffffffffff82111715610525576105246104d0565b5b80604052505050565b5f6105376104a7565b905061054382826104fd565b919050565b5f67ffffffffffffffff821115610562576105616104d0565b5b61056b826104c0565b9050602081019050919050565b5f5b8381101561059557808201518184015260208101905061057a565b5f8484015250505050565b5f6105b26105ad84610548565b61052e565b9050828152602081018484840111156105ce576105cd6104bc565b5b6105d9848285610578565b509392505050565b5f82601f8301126105f5576105f46104b8565b5b81516106058482602086016105a0565b91505092915050565b5f60ff82169050919050565b6106238161060e565b811461062d575f5ffd5b50565b5f8151905061063e8161061a565b92915050565b5f819050919050565b61065681610644565b8114610660575f5ffd5b50565b5f815190506106718161064d565b92915050565b5f5f5f5f6080858703121561068f5761068e6104b0565b5b5f85015167ffffffffffffffff8111156106ac576106ab6104b4565b5b6106b8878288016105e1565b945050602085015167ffffffffffffffff8111156106d9576106d86104b4565b5b6106e5878288016105e1565b93505060406106f687828801610630565b925050606061070787828801610663565b91505092959194509250565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061076157607f821691505b6020821081036107745761077361071d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026107d67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261079b565b6107e0868361079b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61081b61081661081184610644565b6107f8565b610644565b9050919050565b5f819050919050565b61083483610801565b61084861084082610822565b8484546107a7565b825550505050565b5f5f905090565b61085f610850565b61086a81848461082b565b505050565b5b8181101561088d576108825f82610857565b600181019050610870565b5050565b601f8211156108d2576108a38161077a565b6108ac8461078c565b810160208510156108bb578190505b6108cf6108c78561078c565b83018261086f565b50505b505050565b5f82821c905092915050565b5f6108f25f19846008026108d7565b1980831691505092915050565b5f61090a83836108e3565b9150826002028217905092915050565b61092382610713565b67ffffffffffffffff81111561093c5761093b6104d0565b5b610946825461074a565b610951828285610891565b5f60209050601f831160018114610982575f8415610970578287015190505b61097a85826108ff565b8655506109e1565b601f1984166109908661077a565b5f5b828110156109b757848901518255600182019150602085019450602081019050610992565b868310156109d457848901516109d0601f8916826108e3565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f610a0382610713565b610a0d81856109e9565b9350610a1d818560208601610578565b610a26816104c0565b840191505092915050565b610a3a8161060e565b82525050565b5f6060820190508181035f830152610a5881866109f9565b90508181036020830152610a6c81856109f9565b9050610a7b6040830184610a31565b949350505050565b5f8115159050919050565b610a9781610a83565b8114610aa1575f5ffd5b50565b5f81519050610ab281610a8e565b92915050565b5f60208284031215610acd57610acc6104b0565b5b5f610ada84828501610aa4565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b0c82610ae3565b9050919050565b610b1c81610b02565b82525050565b5f602082019050610b355f830184610b13565b92915050565b610b4481610644565b82525050565b5f604082019050610b5d5f830185610b13565b610b6a6020830184610b3b565b9392505050565b5f606082019050610b845f830186610b13565b610b916020830185610b13565b610b9e6040830184610b3b565b949350505050565b5f602082019050610bb95f830184610b3b565b92915050565b6112a480610bcc5f395ff3fe608060405234801561000f575f5ffd5b5060043610610091575f3560e01c8063313ce56711610064578063313ce5671461013157806370a082311461014f57806395d89b411461017f578063a9059cbb1461019d578063dd62ed3e146101cd57610091565b806306fdde0314610095578063095ea7b3146100b357806318160ddd146100e357806323b872dd14610101575b5f5ffd5b61009d6101fd565b6040516100aa9190610ce6565b60405180910390f35b6100cd60048036038101906100c89190610da4565b6102aa565b6040516100da9190610dfc565b60405180910390f35b6100eb6102cc565b6040516100f89190610e24565b60405180910390f35b61011b60048036038101906101169190610e3d565b61036b565b6040516101289190610dfc565b60405180910390f35b610139610399565b6040516101469190610ea8565b60405180910390f35b61016960048036038101906101649190610ec1565b610447565b6040516101769190610e24565b60405180910390f35b6101876104ea565b6040516101949190610ce6565b60405180910390f35b6101b760048036038101906101b29190610da4565b610598565b6040516101c49190610dfc565b60405180910390f35b6101e760048036038101906101e29190610eec565b6105ba565b6040516101f49190610e24565b60405180910390f35b60608060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b815260040161025a9190610f39565b5f60405180830381865afa158015610274573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061029c919061109a565b905050809150508091505090565b5f5f6102b461063c565b90506102c1818585610643565b600191505092915050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4dc2aa4306040518263ffffffff1660e01b81526004016103279190610f39565b602060405180830381865afa158015610342573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103669190611136565b905090565b5f5f61037561063c565b9050610382858285610655565b61038d8585856106e8565b60019150509392505050565b5f5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b81526004016103f59190610f39565b5f60405180830381865afa15801561040f573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610437919061109a565b9091509050809150508091505090565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f7888aec30846040518363ffffffff1660e01b81526004016104a4929190611161565b602060405180830381865afa1580156104bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104e39190611136565b9050919050565b60608060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b81526004016105479190610f39565b5f60405180830381865afa158015610561573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610589919061109a565b90915050809150508091505090565b5f5f6105a261063c565b90506105af8185856106e8565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f33905090565b61065083838360016107d8565b505050565b5f61066084846105ba565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156106e257818110156106d3578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016106ca93929190611188565b60405180910390fd5b6106e184848484035f6107d8565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610758575f6040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260040161074f9190610f39565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107c8575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016107bf9190610f39565b60405180910390fd5b6107d38383836109a7565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610848575f6040517fe602df0500000000000000000000000000000000000000000000000000000000815260040161083f9190610f39565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108b8575f6040517f94280d620000000000000000000000000000000000000000000000000000000081526004016108af9190610f39565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555080156109a1578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516109989190610e24565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a7c5760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b8152600401610a369291906111bd565b6020604051808303815f875af1158015610a52573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a76919061120e565b50610bf2565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b515760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac84836040518363ffffffff1660e01b8152600401610b0b9291906111bd565b6020604051808303815f875af1158015610b27573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b4b919061120e565b50610bf1565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663beabacc88484846040518463ffffffff1660e01b8152600401610baf93929190611239565b6020604051808303815f875af1158015610bcb573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bef919061120e565b505b5b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610c4f9190610e24565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610c93578082015181840152602081019050610c78565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610cb882610c5c565b610cc28185610c66565b9350610cd2818560208601610c76565b610cdb81610c9e565b840191505092915050565b5f6020820190508181035f830152610cfe8184610cae565b905092915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d4082610d17565b9050919050565b610d5081610d36565b8114610d5a575f5ffd5b50565b5f81359050610d6b81610d47565b92915050565b5f819050919050565b610d8381610d71565b8114610d8d575f5ffd5b50565b5f81359050610d9e81610d7a565b92915050565b5f5f60408385031215610dba57610db9610d0f565b5b5f610dc785828601610d5d565b9250506020610dd885828601610d90565b9150509250929050565b5f8115159050919050565b610df681610de2565b82525050565b5f602082019050610e0f5f830184610ded565b92915050565b610e1e81610d71565b82525050565b5f602082019050610e375f830184610e15565b92915050565b5f5f5f60608486031215610e5457610e53610d0f565b5b5f610e6186828701610d5d565b9350506020610e7286828701610d5d565b9250506040610e8386828701610d90565b9150509250925092565b5f60ff82169050919050565b610ea281610e8d565b82525050565b5f602082019050610ebb5f830184610e99565b92915050565b5f60208284031215610ed657610ed5610d0f565b5b5f610ee384828501610d5d565b91505092915050565b5f5f60408385031215610f0257610f01610d0f565b5b5f610f0f85828601610d5d565b9250506020610f2085828601610d5d565b9150509250929050565b610f3381610d36565b82525050565b5f602082019050610f4c5f830184610f2a565b92915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610f9082610c9e565b810181811067ffffffffffffffff82111715610faf57610fae610f5a565b5b80604052505050565b5f610fc1610d06565b9050610fcd8282610f87565b919050565b5f67ffffffffffffffff821115610fec57610feb610f5a565b5b610ff582610c9e565b9050602081019050919050565b5f61101461100f84610fd2565b610fb8565b9050828152602081018484840111156110305761102f610f56565b5b61103b848285610c76565b509392505050565b5f82601f83011261105757611056610f52565b5b8151611067848260208601611002565b91505092915050565b61107981610e8d565b8114611083575f5ffd5b50565b5f8151905061109481611070565b92915050565b5f5f5f606084860312156110b1576110b0610d0f565b5b5f84015167ffffffffffffffff8111156110ce576110cd610d13565b5b6110da86828701611043565b935050602084015167ffffffffffffffff8111156110fb576110fa610d13565b5b61110786828701611043565b925050604061111886828701611086565b9150509250925092565b5f8151905061113081610d7a565b92915050565b5f6020828403121561114b5761114a610d0f565b5b5f61115884828501611122565b91505092915050565b5f6040820190506111745f830185610f2a565b6111816020830184610f2a565b9392505050565b5f60608201905061119b5f830186610f2a565b6111a86020830185610e15565b6111b56040830184610e15565b949350505050565b5f6040820190506111d05f830185610f2a565b6111dd6020830184610e15565b9392505050565b6111ed81610de2565b81146111f7575f5ffd5b50565b5f81519050611208816111e4565b92915050565b5f6020828403121561122357611222610d0f565b5b5f611230848285016111fa565b91505092915050565b5f60608201905061124c5f830186610f2a565b6112596020830185610f2a565b6112666040830184610e15565b94935050505056fea2646970667358221220e72ce46492ce4539131ab8c7bc5f37356347b1bada6320589d941b313e79bb5c64736f6c634300081b0033",
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
