package types

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"
)

const CoinbaseOraclePublicKey = "0xfCEAdAFab14d46e20144F48824d0C09B1a03F2BC"

const CoinbaseABIJSON = `[{
	"name": "coinbase",
	"stateMutability": "pure",
	"type": "function",
	"inputs": [
		{ "internalType": "string",   "name": "kind",      "type": "string"   },
		{ "internalType": "uint64",   "name": "timestamp", "type": "uint64"   },
		{ "internalType": "string",   "name": "key",	   "type": "string"   },
		{ "internalType": "uint64",   "name": "value", 	   "type": "uint64"   }
	],
	"outputs": []
}]`

func ValidateCoinbaseSignature(message, signature []byte) error {
	hash := crypto.Keccak256Hash(message)
	return chaintypes.ValidateEthereumSignature(hash, signature, common.HexToAddress(CoinbaseOraclePublicKey))
}

func ParseCoinbaseMessage(message []byte) (*CoinbasePriceState, error) {
	contractAbi, abiErr := abi.JSON(strings.NewReader(CoinbaseABIJSON))
	if abiErr != nil {
		panic("Bad ABI constant!")
	}

	// computed from web3.sha3('coinbase(string,uint64,string,uint64)').substr(0,10)
	coinbaseMethod := common.FromHex("0x96f180bf")
	method, err := contractAbi.MethodById(coinbaseMethod)
	if err != nil {
		return nil, err
	}

	values, err := method.Inputs.Unpack(message)
	if err != nil {
		return nil, err
	}

	var priceState CoinbasePriceState
	if err := method.Inputs.Copy(&priceState, values); err != nil {
		return nil, err
	}

	if priceState.Kind != "prices" || priceState.Timestamp == 0 || priceState.Key == "" || priceState.Value == 0 {
		return nil, ErrBadCoinbaseMessage
	}

	return &priceState, nil
}
