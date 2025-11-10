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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052606460065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604051620021933803806200219383398181016040528101906200006a9190620004ca565b82828260405180602001604052805f81525060405180602001604052805f815250885f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000ff575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000f6919062000588565b60405180910390fd5b6200011081620001e860201b60201c565b508160049081620001229190620007da565b508060059081620001349190620007da565b50505060065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337d2c2f48484846040518463ffffffff1660e01b815260040162000197939291906200091f565b6020604051808303815f875af1158015620001b4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620001da9190620009a2565b5050505050505050620009d2565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620002e582620002ba565b9050919050565b620002f781620002d9565b811462000302575f80fd5b50565b5f815190506200031581620002ec565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200036b8262000323565b810181811067ffffffffffffffff821117156200038d576200038c62000333565b5b80604052505050565b5f620003a1620002a9565b9050620003af828262000360565b919050565b5f67ffffffffffffffff821115620003d157620003d062000333565b5b620003dc8262000323565b9050602081019050919050565b5f5b8381101562000408578082015181840152602081019050620003eb565b5f8484015250505050565b5f620004296200042384620003b4565b62000396565b9050828152602081018484840111156200044857620004476200031f565b5b62000455848285620003e9565b509392505050565b5f82601f8301126200047457620004736200031b565b5b81516200048684826020860162000413565b91505092915050565b5f60ff82169050919050565b620004a6816200048f565b8114620004b1575f80fd5b50565b5f81519050620004c4816200049b565b92915050565b5f805f8060808587031215620004e557620004e4620002b2565b5b5f620004f48782880162000305565b945050602085015167ffffffffffffffff811115620005185762000517620002b6565b5b62000526878288016200045d565b935050604085015167ffffffffffffffff8111156200054a5762000549620002b6565b5b62000558878288016200045d565b92505060606200056b87828801620004b4565b91505092959194509250565b6200058281620002d9565b82525050565b5f6020820190506200059d5f83018462000577565b92915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620005f257607f821691505b602082108103620006085762000607620005ad565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026200066c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200062f565b6200067886836200062f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620006c2620006bc620006b68462000690565b62000699565b62000690565b9050919050565b5f819050919050565b620006dd83620006a2565b620006f5620006ec82620006c9565b8484546200063b565b825550505050565b5f90565b6200070b620006fd565b62000718818484620006d2565b505050565b5b818110156200073f57620007335f8262000701565b6001810190506200071e565b5050565b601f8211156200078e5762000758816200060e565b620007638462000620565b8101602085101562000773578190505b6200078b620007828562000620565b8301826200071d565b50505b505050565b5f82821c905092915050565b5f620007b05f198460080262000793565b1980831691505092915050565b5f620007ca83836200079f565b9150826002028217905092915050565b620007e582620005a3565b67ffffffffffffffff81111562000801576200080062000333565b5b6200080d8254620005da565b6200081a82828562000743565b5f60209050601f83116001811462000850575f84156200083b578287015190505b620008478582620007bd565b865550620008b6565b601f19841662000860866200060e565b5f5b82811015620008895784890151825560018201915060208501945060208101905062000862565b86831015620008a95784890151620008a5601f8916826200079f565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f620008da82620005a3565b620008e68185620008be565b9350620008f8818560208601620003e9565b620009038162000323565b840191505092915050565b62000919816200048f565b82525050565b5f6060820190508181035f830152620009398186620008ce565b905081810360208301526200094f8185620008ce565b90506200096060408301846200090e565b949350505050565b5f8115159050919050565b6200097e8162000968565b811462000989575f80fd5b50565b5f815190506200099c8162000973565b92915050565b5f60208284031215620009ba57620009b9620002b2565b5b5f620009c9848285016200098c565b91505092915050565b6117b380620009e05f395ff3fe6080604052600436106100e7575f3560e01c806370a082311161008957806395d89b411161005857806395d89b41146102c9578063a9059cbb146102f3578063dd62ed3e1461032f578063f2fde38b1461036b576100e7565b806370a0823114610225578063715018a61461026157806379cc6790146102775780638da5cb5b1461029f576100e7565b806323b872dd116100c557806323b872dd1461017b578063313ce567146101b757806340c10f19146101e157806342966c68146101fd576100e7565b806306fdde03146100eb578063095ea7b31461011557806318160ddd14610151575b5f80fd5b3480156100f6575f80fd5b506100ff610393565b60405161010c91906111ca565b60405180910390f35b348015610120575f80fd5b5061013b60048036038101906101369190611288565b610440565b60405161014891906112e0565b60405180910390f35b34801561015c575f80fd5b50610165610462565b6040516101729190611308565b60405180910390f35b348015610186575f80fd5b506101a1600480360381019061019c9190611321565b610501565b6040516101ae91906112e0565b60405180910390f35b3480156101c2575f80fd5b506101cb61052f565b6040516101d8919061138c565b60405180910390f35b6101fb60048036038101906101f69190611288565b6105dd565b005b348015610208575f80fd5b50610223600480360381019061021e91906113a5565b6105f3565b005b348015610230575f80fd5b5061024b600480360381019061024691906113d0565b610607565b6040516102589190611308565b60405180910390f35b34801561026c575f80fd5b506102756106aa565b005b348015610282575f80fd5b5061029d60048036038101906102989190611288565b6106bd565b005b3480156102aa575f80fd5b506102b36106dd565b6040516102c0919061140a565b60405180910390f35b3480156102d4575f80fd5b506102dd610704565b6040516102ea91906111ca565b60405180910390f35b3480156102fe575f80fd5b5061031960048036038101906103149190611288565b6107b2565b60405161032691906112e0565b60405180910390f35b34801561033a575f80fd5b5061035560048036038101906103509190611423565b6107d4565b6040516103629190611308565b60405180910390f35b348015610376575f80fd5b50610391600480360381019061038c91906113d0565b610856565b005b60608060065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b81526004016103f0919061140a565b5f60405180830381865afa15801561040a573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061043291906115a9565b905050809150508091505090565b5f8061044a6108da565b90506104578185856108e1565b600191505092915050565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4dc2aa4306040518263ffffffff1660e01b81526004016104bd919061140a565b602060405180830381865afa1580156104d8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104fc9190611645565b905090565b5f8061050b6108da565b90506105188582856108f3565b610523858585610986565b60019150509392505050565b5f8060065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b815260040161058b919061140a565b5f60405180830381865afa1580156105a5573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906105cd91906115a9565b9091509050809150508091505090565b6105e5610a76565b6105ef8282610afd565b5050565b6106046105fe6108da565b82610b7c565b50565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f7888aec30846040518363ffffffff1660e01b8152600401610664929190611670565b602060405180830381865afa15801561067f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106a39190611645565b9050919050565b6106b2610a76565b6106bb5f610bfb565b565b6106cf826106c96108da565b836108f3565b6106d98282610b7c565b5050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60608060065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b8152600401610761919061140a565b5f60405180830381865afa15801561077b573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906107a391906115a9565b90915050809150508091505090565b5f806107bc6108da565b90506107c9818585610986565b600191505092915050565b5f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b61085e610a76565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108ce575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016108c5919061140a565b60405180910390fd5b6108d781610bfb565b50565b5f33905090565b6108ee8383836001610cbc565b505050565b5f6108fe84846107d4565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156109805781811015610971578281836040517ffb8f41b200000000000000000000000000000000000000000000000000000000815260040161096893929190611697565b60405180910390fd5b61097f84848484035f610cbc565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109f6575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016109ed919061140a565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a66575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610a5d919061140a565b60405180910390fd5b610a71838383610e8b565b505050565b610a7e6108da565b73ffffffffffffffffffffffffffffffffffffffff16610a9c6106dd565b73ffffffffffffffffffffffffffffffffffffffff1614610afb57610abf6108da565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610af2919061140a565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b6d575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610b64919061140a565b60405180910390fd5b610b785f8383610e8b565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610bec575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610be3919061140a565b60405180910390fd5b610bf7825f83610e8b565b5050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610d2c575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610d23919061140a565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610d9c575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610d93919061140a565b60405180910390fd5b8160025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610e85578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610e7c9190611308565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610f605760065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b8152600401610f1a9291906116cc565b6020604051808303815f875af1158015610f36573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f5a919061171d565b506110d6565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036110355760065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac84836040518363ffffffff1660e01b8152600401610fef9291906116cc565b6020604051808303815f875af115801561100b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061102f919061171d565b506110d5565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663beabacc88484846040518463ffffffff1660e01b815260040161109393929190611748565b6020604051808303815f875af11580156110af573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110d3919061171d565b505b5b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516111339190611308565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561117757808201518184015260208101905061115c565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61119c82611140565b6111a6818561114a565b93506111b681856020860161115a565b6111bf81611182565b840191505092915050565b5f6020820190508181035f8301526111e28184611192565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611224826111fb565b9050919050565b6112348161121a565b811461123e575f80fd5b50565b5f8135905061124f8161122b565b92915050565b5f819050919050565b61126781611255565b8114611271575f80fd5b50565b5f813590506112828161125e565b92915050565b5f806040838503121561129e5761129d6111f3565b5b5f6112ab85828601611241565b92505060206112bc85828601611274565b9150509250929050565b5f8115159050919050565b6112da816112c6565b82525050565b5f6020820190506112f35f8301846112d1565b92915050565b61130281611255565b82525050565b5f60208201905061131b5f8301846112f9565b92915050565b5f805f60608486031215611338576113376111f3565b5b5f61134586828701611241565b935050602061135686828701611241565b925050604061136786828701611274565b9150509250925092565b5f60ff82169050919050565b61138681611371565b82525050565b5f60208201905061139f5f83018461137d565b92915050565b5f602082840312156113ba576113b96111f3565b5b5f6113c784828501611274565b91505092915050565b5f602082840312156113e5576113e46111f3565b5b5f6113f284828501611241565b91505092915050565b6114048161121a565b82525050565b5f60208201905061141d5f8301846113fb565b92915050565b5f8060408385031215611439576114386111f3565b5b5f61144685828601611241565b925050602061145785828601611241565b9150509250929050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61149f82611182565b810181811067ffffffffffffffff821117156114be576114bd611469565b5b80604052505050565b5f6114d06111ea565b90506114dc8282611496565b919050565b5f67ffffffffffffffff8211156114fb576114fa611469565b5b61150482611182565b9050602081019050919050565b5f61152361151e846114e1565b6114c7565b90508281526020810184848401111561153f5761153e611465565b5b61154a84828561115a565b509392505050565b5f82601f83011261156657611565611461565b5b8151611576848260208601611511565b91505092915050565b61158881611371565b8114611592575f80fd5b50565b5f815190506115a38161157f565b92915050565b5f805f606084860312156115c0576115bf6111f3565b5b5f84015167ffffffffffffffff8111156115dd576115dc6111f7565b5b6115e986828701611552565b935050602084015167ffffffffffffffff81111561160a576116096111f7565b5b61161686828701611552565b925050604061162786828701611595565b9150509250925092565b5f8151905061163f8161125e565b92915050565b5f6020828403121561165a576116596111f3565b5b5f61166784828501611631565b91505092915050565b5f6040820190506116835f8301856113fb565b61169060208301846113fb565b9392505050565b5f6060820190506116aa5f8301866113fb565b6116b760208301856112f9565b6116c460408301846112f9565b949350505050565b5f6040820190506116df5f8301856113fb565b6116ec60208301846112f9565b9392505050565b6116fc816112c6565b8114611706575f80fd5b50565b5f81519050611717816116f3565b92915050565b5f60208284031215611732576117316111f3565b5b5f61173f84828501611709565b91505092915050565b5f60608201905061175b5f8301866113fb565b61176860208301856113fb565b61177560408301846112f9565b94935050505056fea264697066735822122019ec31921e777a6c5f5d27b3afb27acc7727f7637b7555591c0b50efc95e6bc964736f6c63430008180033",
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
