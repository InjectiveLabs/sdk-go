package types

import (
	"bytes"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	CoinbaseOraclePublicKey = "0xfCEAdAFab14d46e20144F48824d0C09B1a03F2BC"
	preamblePrefix          = "\x19Ethereum Signed Message:\n32"
)

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
	return ValidateEthereumSignature(hash, signature, common.HexToAddress(CoinbaseOraclePublicKey))
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

// ValidateEthereumSignature takes a message, an associated signature and public key and
// returns an error if the signature isn't valid
// TODO: refactor to shared common dir, copy pasted below code from Peggy
func ValidateEthereumSignature(hash common.Hash, signature []byte, ethAddress common.Address) error {

	// convert malformed coinbase sig in oracle response to 65-byte signature
	trimmedSig := signature[:65]
	trimmedSig[64] = signature[95]

	if len(trimmedSig) < 65 {
		return sdkerrors.Wrap(ErrInvalidEthereumSignature, "signature too short")
	}

	// calculate recover id
	if trimmedSig[64] == 27 || signature[64] == 28 {
		trimmedSig[64] -= 27
	}

	// manually build the hash with ethereum prefix
	preamblePrefix := []byte(preamblePrefix)
	preambleMessage := append(preamblePrefix, hash.Bytes()...) //nolint:gocritic
	preambleHash := crypto.Keccak256Hash(preambleMessage)

	// verify signature
	pubkey, err := crypto.SigToPub(preambleHash.Bytes(), trimmedSig)
	if err != nil {
		return sdkerrors.Wrap(err, "signature to public key")
	}
	addr := crypto.PubkeyToAddress(*pubkey)

	if !bytes.Equal(addr.Bytes(), ethAddress.Bytes()) {
		return sdkerrors.Wrapf(ErrInvalidEthereumSignature, "signature not matching, expected %s but got %s", ethAddress.Hex(), addr.Hex())
	}

	return nil
}
