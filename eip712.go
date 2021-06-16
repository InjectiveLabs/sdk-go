package sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/typeddata"
)

type EIP712Domain struct {
	VerifyingContract common.Address `json:"verifyingContract"`
	ChainID           *big.Int       `json:"chainID"`
}

var eip712OrderTypes = typeddata.Types{
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
	"Order": {
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

var eip712SpotOrderTypes = typeddata.Types{
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
			Type: "string",
		},
	},
	"SpotOrder": {
		{
			Name: "subaccountID",
			Type: "bytes32",
		},
		{
			Name: "sender",
			Type: "address",
		},
		{
			Name: "feeRecipient",
			Type: "address",
		},
		{
			Name: "expiry",
			Type: "uint64",
		},
		{
			Name: "marketID",
			Type: "bytes32",
		},
		{
			Name: "supplyAmount",
			Type: "uint128",
		},
		{
			Name: "receiveAmount",
			Type: "uint128",
		},
		{
			Name: "salt",
			Type: "uint64",
		},
		{
			Name: "orderType",
			Type: "uint8",
		},
		{
			Name: "triggerPrice",
			Type: "uint128",
		},
	},
}

var eip712TransactionTypes = typeddata.Types{
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

var eip712CoordinatorApprovalTypes = typeddata.Types{
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
