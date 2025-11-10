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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"initial_supply_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052606460055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405162001f3f38038062001f3f83398181016040528101906200006a9190620006ca565b83838360405180602001604052805f81525060405180602001604052805f81525081600390816200009c9190620009a5565b508060049081620000ae9190620009a5565b50505060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337d2c2f48484846040518463ffffffff1660e01b8152600401620001119392919062000aea565b6020604051808303815f875af11580156200012e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000154919062000b6d565b505050505f81111562000174576200017333826200017e60201b60201c565b5b5050505062000c8d565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603620001f1575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401620001e8919062000be0565b60405180910390fd5b620002045f83836200020860201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603620002e45760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b81526004016200029a92919062000c0c565b6020604051808303815f875af1158015620002b7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620002dd919062000b6d565b5062000466565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603620003c05760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac84836040518363ffffffff1660e01b81526004016200037692919062000c0c565b6020604051808303815f875af115801562000393573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620003b9919062000b6d565b5062000465565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663beabacc88484846040518463ffffffff1660e01b8152600401620004209392919062000c37565b6020604051808303815f875af11580156200043d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000463919062000b6d565b505b5b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620004c5919062000c72565b60405180910390a3505050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200053382620004eb565b810181811067ffffffffffffffff82111715620005555762000554620004fb565b5b80604052505050565b5f62000569620004d2565b905062000577828262000528565b919050565b5f67ffffffffffffffff821115620005995762000598620004fb565b5b620005a482620004eb565b9050602081019050919050565b5f5b83811015620005d0578082015181840152602081019050620005b3565b5f8484015250505050565b5f620005f1620005eb846200057c565b6200055e565b90508281526020810184848401111562000610576200060f620004e7565b5b6200061d848285620005b1565b509392505050565b5f82601f8301126200063c576200063b620004e3565b5b81516200064e848260208601620005db565b91505092915050565b5f60ff82169050919050565b6200066e8162000657565b811462000679575f80fd5b50565b5f815190506200068c8162000663565b92915050565b5f819050919050565b620006a68162000692565b8114620006b1575f80fd5b50565b5f81519050620006c4816200069b565b92915050565b5f805f8060808587031215620006e557620006e4620004db565b5b5f85015167ffffffffffffffff811115620007055762000704620004df565b5b620007138782880162000625565b945050602085015167ffffffffffffffff811115620007375762000736620004df565b5b620007458782880162000625565b935050604062000758878288016200067c565b92505060606200076b87828801620006b4565b91505092959194509250565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620007c657607f821691505b602082108103620007dc57620007db62000781565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620008407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000803565b6200084c868362000803565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6200088d62000887620008818462000692565b62000864565b62000692565b9050919050565b5f819050919050565b620008a8836200086d565b620008c0620008b78262000894565b8484546200080f565b825550505050565b5f90565b620008d6620008c8565b620008e38184846200089d565b505050565b5b818110156200090a57620008fe5f82620008cc565b600181019050620008e9565b5050565b601f82111562000959576200092381620007e2565b6200092e84620007f4565b810160208510156200093e578190505b620009566200094d85620007f4565b830182620008e8565b50505b505050565b5f82821c905092915050565b5f6200097b5f19846008026200095e565b1980831691505092915050565b5f6200099583836200096a565b9150826002028217905092915050565b620009b08262000777565b67ffffffffffffffff811115620009cc57620009cb620004fb565b5b620009d88254620007ae565b620009e58282856200090e565b5f60209050601f83116001811462000a1b575f841562000a06578287015190505b62000a12858262000988565b86555062000a81565b601f19841662000a2b86620007e2565b5f5b8281101562000a545784890151825560018201915060208501945060208101905062000a2d565b8683101562000a74578489015162000a70601f8916826200096a565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f62000aa58262000777565b62000ab1818562000a89565b935062000ac3818560208601620005b1565b62000ace81620004eb565b840191505092915050565b62000ae48162000657565b82525050565b5f6060820190508181035f83015262000b04818662000a99565b9050818103602083015262000b1a818562000a99565b905062000b2b604083018462000ad9565b949350505050565b5f8115159050919050565b62000b498162000b33565b811462000b54575f80fd5b50565b5f8151905062000b678162000b3e565b92915050565b5f6020828403121562000b855762000b84620004db565b5b5f62000b948482850162000b57565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000bc88262000b9d565b9050919050565b62000bda8162000bbc565b82525050565b5f60208201905062000bf55f83018462000bcf565b92915050565b62000c068162000692565b82525050565b5f60408201905062000c215f83018562000bcf565b62000c30602083018462000bfb565b9392505050565b5f60608201905062000c4c5f83018662000bcf565b62000c5b602083018562000bcf565b62000c6a604083018462000bfb565b949350505050565b5f60208201905062000c875f83018462000bfb565b92915050565b6112a48062000c9b5f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c8063313ce56711610064578063313ce5671461013157806370a082311461014f57806395d89b411461017f578063a9059cbb1461019d578063dd62ed3e146101cd57610091565b806306fdde0314610095578063095ea7b3146100b357806318160ddd146100e357806323b872dd14610101575b5f80fd5b61009d6101fd565b6040516100aa9190610ce6565b60405180910390f35b6100cd60048036038101906100c89190610da4565b6102aa565b6040516100da9190610dfc565b60405180910390f35b6100eb6102cc565b6040516100f89190610e24565b60405180910390f35b61011b60048036038101906101169190610e3d565b61036b565b6040516101289190610dfc565b60405180910390f35b610139610399565b6040516101469190610ea8565b60405180910390f35b61016960048036038101906101649190610ec1565b610447565b6040516101769190610e24565b60405180910390f35b6101876104ea565b6040516101949190610ce6565b60405180910390f35b6101b760048036038101906101b29190610da4565b610598565b6040516101c49190610dfc565b60405180910390f35b6101e760048036038101906101e29190610eec565b6105ba565b6040516101f49190610e24565b60405180910390f35b60608060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b815260040161025a9190610f39565b5f60405180830381865afa158015610274573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061029c919061109a565b905050809150508091505090565b5f806102b461063c565b90506102c1818585610643565b600191505092915050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4dc2aa4306040518263ffffffff1660e01b81526004016103279190610f39565b602060405180830381865afa158015610342573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103669190611136565b905090565b5f8061037561063c565b9050610382858285610655565b61038d8585856106e8565b60019150509392505050565b5f8060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b81526004016103f59190610f39565b5f60405180830381865afa15801561040f573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610437919061109a565b9091509050809150508091505090565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f7888aec30846040518363ffffffff1660e01b81526004016104a4929190611161565b602060405180830381865afa1580156104bf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104e39190611136565b9050919050565b60608060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b81526004016105479190610f39565b5f60405180830381865afa158015610561573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610589919061109a565b90915050809150508091505090565b5f806105a261063c565b90506105af8185856106e8565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f33905090565b61065083838360016107d8565b505050565b5f61066084846105ba565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156106e257818110156106d3578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016106ca93929190611188565b60405180910390fd5b6106e184848484035f6107d8565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610758575f6040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260040161074f9190610f39565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107c8575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016107bf9190610f39565b60405180910390fd5b6107d38383836109a7565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610848575f6040517fe602df0500000000000000000000000000000000000000000000000000000000815260040161083f9190610f39565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108b8575f6040517f94280d620000000000000000000000000000000000000000000000000000000081526004016108af9190610f39565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555080156109a1578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516109989190610e24565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a7c5760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b8152600401610a369291906111bd565b6020604051808303815f875af1158015610a52573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a76919061120e565b50610bf2565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b515760055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac84836040518363ffffffff1660e01b8152600401610b0b9291906111bd565b6020604051808303815f875af1158015610b27573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b4b919061120e565b50610bf1565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663beabacc88484846040518463ffffffff1660e01b8152600401610baf93929190611239565b6020604051808303815f875af1158015610bcb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bef919061120e565b505b5b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610c4f9190610e24565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610c93578082015181840152602081019050610c78565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610cb882610c5c565b610cc28185610c66565b9350610cd2818560208601610c76565b610cdb81610c9e565b840191505092915050565b5f6020820190508181035f830152610cfe8184610cae565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d4082610d17565b9050919050565b610d5081610d36565b8114610d5a575f80fd5b50565b5f81359050610d6b81610d47565b92915050565b5f819050919050565b610d8381610d71565b8114610d8d575f80fd5b50565b5f81359050610d9e81610d7a565b92915050565b5f8060408385031215610dba57610db9610d0f565b5b5f610dc785828601610d5d565b9250506020610dd885828601610d90565b9150509250929050565b5f8115159050919050565b610df681610de2565b82525050565b5f602082019050610e0f5f830184610ded565b92915050565b610e1e81610d71565b82525050565b5f602082019050610e375f830184610e15565b92915050565b5f805f60608486031215610e5457610e53610d0f565b5b5f610e6186828701610d5d565b9350506020610e7286828701610d5d565b9250506040610e8386828701610d90565b9150509250925092565b5f60ff82169050919050565b610ea281610e8d565b82525050565b5f602082019050610ebb5f830184610e99565b92915050565b5f60208284031215610ed657610ed5610d0f565b5b5f610ee384828501610d5d565b91505092915050565b5f8060408385031215610f0257610f01610d0f565b5b5f610f0f85828601610d5d565b9250506020610f2085828601610d5d565b9150509250929050565b610f3381610d36565b82525050565b5f602082019050610f4c5f830184610f2a565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610f9082610c9e565b810181811067ffffffffffffffff82111715610faf57610fae610f5a565b5b80604052505050565b5f610fc1610d06565b9050610fcd8282610f87565b919050565b5f67ffffffffffffffff821115610fec57610feb610f5a565b5b610ff582610c9e565b9050602081019050919050565b5f61101461100f84610fd2565b610fb8565b9050828152602081018484840111156110305761102f610f56565b5b61103b848285610c76565b509392505050565b5f82601f83011261105757611056610f52565b5b8151611067848260208601611002565b91505092915050565b61107981610e8d565b8114611083575f80fd5b50565b5f8151905061109481611070565b92915050565b5f805f606084860312156110b1576110b0610d0f565b5b5f84015167ffffffffffffffff8111156110ce576110cd610d13565b5b6110da86828701611043565b935050602084015167ffffffffffffffff8111156110fb576110fa610d13565b5b61110786828701611043565b925050604061111886828701611086565b9150509250925092565b5f8151905061113081610d7a565b92915050565b5f6020828403121561114b5761114a610d0f565b5b5f61115884828501611122565b91505092915050565b5f6040820190506111745f830185610f2a565b6111816020830184610f2a565b9392505050565b5f60608201905061119b5f830186610f2a565b6111a86020830185610e15565b6111b56040830184610e15565b949350505050565b5f6040820190506111d05f830185610f2a565b6111dd6020830184610e15565b9392505050565b6111ed81610de2565b81146111f7575f80fd5b50565b5f81519050611208816111e4565b92915050565b5f6020828403121561122357611222610d0f565b5b5f611230848285016111fa565b91505092915050565b5f60608201905061124c5f830186610f2a565b6112596020830185610f2a565b6112666040830184610e15565b94935050505056fea2646970667358221220f46d089a4318bb946632ecd6859aaafd21c001203190b421f595ecd0e45f0f8b64736f6c63430008180033",
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
