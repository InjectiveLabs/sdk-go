package zeroex

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
)

type EIP712Domain struct {
	VerifyingContract common.Address `json:"verifyingContract"`
	ChainID           *big.Int       `json:"chainID"`
}

var eip712OrderTypes = gethsigner.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"ZOrder": {
		{
			Name: "makerAddress",
			Type: "address",
		},
		{
			Name: "takerAddress",
			Type: "address",
		},
		{
			Name: "feeRecipientAddress",
			Type: "address",
		},
		{
			Name: "senderAddress",
			Type: "address",
		},
		{
			Name: "makerAssetAmount",
			Type: "uint256",
		},
		{
			Name: "takerAssetAmount",
			Type: "uint256",
		},
		{
			Name: "makerFee",
			Type: "uint256",
		},
		{
			Name: "takerFee",
			Type: "uint256",
		},
		{
			Name: "expirationTimeSeconds",
			Type: "uint256",
		},
		{
			Name: "salt",
			Type: "uint256",
		},
		{
			Name: "makerAssetData",
			Type: "bytes",
		},
		{
			Name: "takerAssetData",
			Type: "bytes",
		},
		{
			Name: "makerFeeAssetData",
			Type: "bytes",
		},
		{
			Name: "takerFeeAssetData",
			Type: "bytes",
		},
	},
}

var eip712TransactionTypes = gethsigner.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"ZeroExTransaction": {
		{
			Name: "salt",
			Type: "uint256",
		},
		{
			Name: "expirationTimeSeconds",
			Type: "uint256",
		},
		{
			Name: "gasPrice",
			Type: "uint256",
		},
		{
			Name: "signerAddress",
			Type: "address",
		},
		{
			Name: "data",
			Type: "bytes",
		},
	},
}

var eip712CoordinatorApprovalTypes = gethsigner.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"CoordinatorApproval": {
		{
			Name: "txOrigin",
			Type: "address",
		},
		{
			Name: "transactionHash",
			Type: "bytes",
		},
		{
			Name: "transactionSignature",
			Type: "bytes",
		},
	},
}
