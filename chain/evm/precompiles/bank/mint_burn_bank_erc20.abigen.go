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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052606460065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801562000051575f80fd5b5060405162002104380380620021048339818101604052810190620000779190620004d7565b82828260405180602001604052805f81525060405180602001604052805f815250885f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200010c575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040162000103919062000595565b60405180910390fd5b6200011d81620001f560201b60201c565b5081600490816200012f9190620007e7565b508060059081620001419190620007e7565b50505060065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337d2c2f48484846040518463ffffffff1660e01b8152600401620001a4939291906200092c565b6020604051808303815f875af1158015620001c1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620001e79190620009af565b5050505050505050620009df565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620002f282620002c7565b9050919050565b6200030481620002e6565b81146200030f575f80fd5b50565b5f815190506200032281620002f9565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620003788262000330565b810181811067ffffffffffffffff821117156200039a576200039962000340565b5b80604052505050565b5f620003ae620002b6565b9050620003bc82826200036d565b919050565b5f67ffffffffffffffff821115620003de57620003dd62000340565b5b620003e98262000330565b9050602081019050919050565b5f5b8381101562000415578082015181840152602081019050620003f8565b5f8484015250505050565b5f620004366200043084620003c1565b620003a3565b9050828152602081018484840111156200045557620004546200032c565b5b62000462848285620003f6565b509392505050565b5f82601f83011262000481576200048062000328565b5b81516200049384826020860162000420565b91505092915050565b5f60ff82169050919050565b620004b3816200049c565b8114620004be575f80fd5b50565b5f81519050620004d181620004a8565b92915050565b5f805f8060808587031215620004f257620004f1620002bf565b5b5f620005018782880162000312565b945050602085015167ffffffffffffffff811115620005255762000524620002c3565b5b62000533878288016200046a565b935050604085015167ffffffffffffffff811115620005575762000556620002c3565b5b62000565878288016200046a565b92505060606200057887828801620004c1565b91505092959194509250565b6200058f81620002e6565b82525050565b5f602082019050620005aa5f83018462000584565b92915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620005ff57607f821691505b602082108103620006155762000614620005ba565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620006797fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200063c565b6200068586836200063c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620006cf620006c9620006c3846200069d565b620006a6565b6200069d565b9050919050565b5f819050919050565b620006ea83620006af565b62000702620006f982620006d6565b84845462000648565b825550505050565b5f90565b620007186200070a565b62000725818484620006df565b505050565b5b818110156200074c57620007405f826200070e565b6001810190506200072b565b5050565b601f8211156200079b5762000765816200061b565b62000770846200062d565b8101602085101562000780578190505b620007986200078f856200062d565b8301826200072a565b50505b505050565b5f82821c905092915050565b5f620007bd5f1984600802620007a0565b1980831691505092915050565b5f620007d78383620007ac565b9150826002028217905092915050565b620007f282620005b0565b67ffffffffffffffff8111156200080e576200080d62000340565b5b6200081a8254620005e7565b6200082782828562000750565b5f60209050601f8311600181146200085d575f841562000848578287015190505b620008548582620007ca565b865550620008c3565b601f1984166200086d866200061b565b5f5b8281101562000896578489015182556001820191506020850194506020810190506200086f565b86831015620008b65784890151620008b2601f891682620007ac565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f620008e782620005b0565b620008f38185620008cb565b935062000905818560208601620003f6565b620009108162000330565b840191505092915050565b62000926816200049c565b82525050565b5f6060820190508181035f830152620009468186620008db565b905081810360208301526200095c8185620008db565b90506200096d60408301846200091b565b949350505050565b5f8115159050919050565b6200098b8162000975565b811462000996575f80fd5b50565b5f81519050620009a98162000980565b92915050565b5f60208284031215620009c757620009c6620002bf565b5b5f620009d68482850162000999565b91505092915050565b61171780620009ed5f395ff3fe608060405234801561000f575f80fd5b50600436106100f3575f3560e01c806370a082311161009557806395d89b411161006457806395d89b411461025d578063a9059cbb1461027b578063dd62ed3e146102ab578063f2fde38b146102db576100f3565b806370a08231146101e9578063715018a61461021957806379cc6790146102235780638da5cb5b1461023f576100f3565b806323b872dd116100d157806323b872dd14610163578063313ce5671461019357806340c10f19146101b157806342966c68146101cd576100f3565b806306fdde03146100f7578063095ea7b31461011557806318160ddd14610145575b5f80fd5b6100ff6102f7565b60405161010c919061112e565b60405180910390f35b61012f600480360381019061012a91906111ec565b6103a4565b60405161013c9190611244565b60405180910390f35b61014d6103c6565b60405161015a919061126c565b60405180910390f35b61017d60048036038101906101789190611285565b610465565b60405161018a9190611244565b60405180910390f35b61019b610493565b6040516101a891906112f0565b60405180910390f35b6101cb60048036038101906101c691906111ec565b610541565b005b6101e760048036038101906101e29190611309565b610557565b005b61020360048036038101906101fe9190611334565b61056b565b604051610210919061126c565b60405180910390f35b61022161060e565b005b61023d600480360381019061023891906111ec565b610621565b005b610247610641565b604051610254919061136e565b60405180910390f35b610265610668565b604051610272919061112e565b60405180910390f35b610295600480360381019061029091906111ec565b610716565b6040516102a29190611244565b60405180910390f35b6102c560048036038101906102c09190611387565b610738565b6040516102d2919061126c565b60405180910390f35b6102f560048036038101906102f09190611334565b6107ba565b005b60608060065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b8152600401610354919061136e565b5f60405180830381865afa15801561036e573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610396919061150d565b905050809150508091505090565b5f806103ae61083e565b90506103bb818585610845565b600191505092915050565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4dc2aa4306040518263ffffffff1660e01b8152600401610421919061136e565b602060405180830381865afa15801561043c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046091906115a9565b905090565b5f8061046f61083e565b905061047c858285610857565b6104878585856108ea565b60019150509392505050565b5f8060065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b81526004016104ef919061136e565b5f60405180830381865afa158015610509573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610531919061150d565b9091509050809150508091505090565b6105496109da565b6105538282610a61565b5050565b61056861056261083e565b82610ae0565b50565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f7888aec30846040518363ffffffff1660e01b81526004016105c89291906115d4565b602060405180830381865afa1580156105e3573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061060791906115a9565b9050919050565b6106166109da565b61061f5f610b5f565b565b6106338261062d61083e565b83610857565b61063d8282610ae0565b5050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60608060065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba21572306040518263ffffffff1660e01b81526004016106c5919061136e565b5f60405180830381865afa1580156106df573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610707919061150d565b90915050809150508091505090565b5f8061072061083e565b905061072d8185856108ea565b600191505092915050565b5f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6107c26109da565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610832575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610829919061136e565b60405180910390fd5b61083b81610b5f565b50565b5f33905090565b6108528383836001610c20565b505050565b5f6108628484610738565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156108e457818110156108d5578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016108cc939291906115fb565b60405180910390fd5b6108e384848484035f610c20565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361095a575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610951919061136e565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109ca575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016109c1919061136e565b60405180910390fd5b6109d5838383610def565b505050565b6109e261083e565b73ffffffffffffffffffffffffffffffffffffffff16610a00610641565b73ffffffffffffffffffffffffffffffffffffffff1614610a5f57610a2361083e565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610a56919061136e565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ad1575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610ac8919061136e565b60405180910390fd5b610adc5f8383610def565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b50575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610b47919061136e565b60405180910390fd5b610b5b825f83610def565b5050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610c90575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610c87919061136e565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610d00575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610cf7919061136e565b60405180910390fd5b8160025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610de9578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610de0919061126c565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ec45760065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b8152600401610e7e929190611630565b6020604051808303815f875af1158015610e9a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ebe9190611681565b5061103a565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610f995760065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac84836040518363ffffffff1660e01b8152600401610f53929190611630565b6020604051808303815f875af1158015610f6f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f939190611681565b50611039565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663beabacc88484846040518463ffffffff1660e01b8152600401610ff7939291906116ac565b6020604051808303815f875af1158015611013573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110379190611681565b505b5b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051611097919061126c565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156110db5780820151818401526020810190506110c0565b5f8484015250505050565b5f601f19601f8301169050919050565b5f611100826110a4565b61110a81856110ae565b935061111a8185602086016110be565b611123816110e6565b840191505092915050565b5f6020820190508181035f83015261114681846110f6565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6111888261115f565b9050919050565b6111988161117e565b81146111a2575f80fd5b50565b5f813590506111b38161118f565b92915050565b5f819050919050565b6111cb816111b9565b81146111d5575f80fd5b50565b5f813590506111e6816111c2565b92915050565b5f806040838503121561120257611201611157565b5b5f61120f858286016111a5565b9250506020611220858286016111d8565b9150509250929050565b5f8115159050919050565b61123e8161122a565b82525050565b5f6020820190506112575f830184611235565b92915050565b611266816111b9565b82525050565b5f60208201905061127f5f83018461125d565b92915050565b5f805f6060848603121561129c5761129b611157565b5b5f6112a9868287016111a5565b93505060206112ba868287016111a5565b92505060406112cb868287016111d8565b9150509250925092565b5f60ff82169050919050565b6112ea816112d5565b82525050565b5f6020820190506113035f8301846112e1565b92915050565b5f6020828403121561131e5761131d611157565b5b5f61132b848285016111d8565b91505092915050565b5f6020828403121561134957611348611157565b5b5f611356848285016111a5565b91505092915050565b6113688161117e565b82525050565b5f6020820190506113815f83018461135f565b92915050565b5f806040838503121561139d5761139c611157565b5b5f6113aa858286016111a5565b92505060206113bb858286016111a5565b9150509250929050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611403826110e6565b810181811067ffffffffffffffff82111715611422576114216113cd565b5b80604052505050565b5f61143461114e565b905061144082826113fa565b919050565b5f67ffffffffffffffff82111561145f5761145e6113cd565b5b611468826110e6565b9050602081019050919050565b5f61148761148284611445565b61142b565b9050828152602081018484840111156114a3576114a26113c9565b5b6114ae8482856110be565b509392505050565b5f82601f8301126114ca576114c96113c5565b5b81516114da848260208601611475565b91505092915050565b6114ec816112d5565b81146114f6575f80fd5b50565b5f81519050611507816114e3565b92915050565b5f805f6060848603121561152457611523611157565b5b5f84015167ffffffffffffffff8111156115415761154061115b565b5b61154d868287016114b6565b935050602084015167ffffffffffffffff81111561156e5761156d61115b565b5b61157a868287016114b6565b925050604061158b868287016114f9565b9150509250925092565b5f815190506115a3816111c2565b92915050565b5f602082840312156115be576115bd611157565b5b5f6115cb84828501611595565b91505092915050565b5f6040820190506115e75f83018561135f565b6115f4602083018461135f565b9392505050565b5f60608201905061160e5f83018661135f565b61161b602083018561125d565b611628604083018461125d565b949350505050565b5f6040820190506116435f83018561135f565b611650602083018461125d565b9392505050565b6116608161122a565b811461166a575f80fd5b50565b5f8151905061167b81611657565b92915050565b5f6020828403121561169657611695611157565b5b5f6116a38482850161166d565b91505092915050565b5f6060820190506116bf5f83018661135f565b6116cc602083018561135f565b6116d9604083018461125d565b94935050505056fea2646970667358221220361f190f8f643e37acec2cbc4bb3a7e57425705745e3da45f46cb66d482dc39d64736f6c63430008180033",
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
// Solidity: function mint(address to, uint256 amount) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MintBurnBankERC20 *MintBurnBankERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MintBurnBankERC20.Contract.Mint(&_MintBurnBankERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
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
