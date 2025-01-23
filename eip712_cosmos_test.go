package sdk

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/typeddata"
)

func TestEIP712Wrapper(t *testing.T) {
	require := require.New(t)

	// Setup test codec
	chainCtx, _ := chainclient.NewClientContext("", "", nil)
	protoCodec := codec.NewProtoCodec(chainCtx.InterfaceRegistry)

	// Test fixture

	var testReq prepareEip712Payload
	err := json.Unmarshal([]byte(testEIP712WrapperJSON), &testReq)
	require.NoError(err)

	legacytx.RegressionTestingAminoCodec = codec.NewLegacyAmino()

	msgs, err := unmarshalTestMsgs(protoCodec, testReq.Msgs)
	require.NoError(err)

	feePriceAmount, ok := math.NewIntFromString(testReq.Fee.Price[0].Amount)
	if !ok {
		require.True(ok, fmt.Sprintf("wrong fee price amount: %s", testReq.Fee.Price[0].Amount))
	}

	// Create wrapper function
	wrapper := func(
		cdc *codec.ProtoCodec,
		chainID uint64,
		signerData *authsigning.SignerData,
		timeoutHeight uint64,
		memo string,
		feeInfo legacytx.StdFee,
		msgs []cosmtypes.Msg,
		feeDelegation *FeeDelegationOptions,
	) (typeddata.TypedData, error) {
		return WrapTxToEIP712(cdc, chainID, signerData, timeoutHeight, memo, feeInfo, msgs, feeDelegation)
	}

	// Test wrapper execution
	typedData, err := wrapper(
		protoCodec,
		uint64(testReq.ChainID),
		&authsigning.SignerData{
			ChainID:       "injective-777",
			AccountNumber: 3,
			Sequence:      0,
		},
		uint64(testReq.TimeoutHeight),
		testReq.Memo,
		legacytx.StdFee{
			Gas: uint64(testReq.Fee.Gas),
			Amount: []cosmtypes.Coin{
				{
					Denom:  testReq.Fee.Price[0].Denom,
					Amount: feePriceAmount,
				},
			},
		},
		msgs,
		&FeeDelegationOptions{},
	)
	require.NoError(err)

	require.NotNil(typedData)
	require.Equal("Injective Web3", typedData.Domain.Name)
	require.Equal("1.0.0", typedData.Domain.Version)
	require.Equal("cosmos", typedData.Domain.VerifyingContract)
	require.Equal("0", typedData.Domain.Salt)
	require.Equal("Tx", typedData.PrimaryType)

	typedDataRaw, err := json.MarshalIndent(typedData, "", "\t")
	require.NoError(err)
	require.Equal(testEIP712WrapperOutJSON, string(typedDataRaw))
}

func unmarshalTestMsgs(cdc codec.ProtoCodecMarshaler, msgsBytes []json.RawMessage) (msgs []cosmtypes.Msg, err error) {
	for _, protoMsg := range msgsBytes {
		var msg cosmtypes.Msg
		if err := cdc.UnmarshalInterfaceJSON(protoMsg, &msg); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal msg")
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

var testEIP712WrapperJSON = `{
    "chainId": "11155111",
    "signerAddress": "0xaf79152ac5df276d9a8e1e2e22822f9713474902",
    "sequence": "0",
    "accountNumber": "3",
    "memo": "",
    "timeoutHeight": "999999999",
    "fee": {
        "price": [
            {
                "denom": "inj",
                "amount": "160000000"
            }
        ],
        "gas": "400000",
        "delegateFee": false
    },
    "msgs": [
        {
            "@type": "/injective.exchange.v1beta1.MsgCancelSpotOrder",
            "sender": "inj14au322k9munkmx5wrchz9q30juf5wjgz2cfqku",
            "market_id": "0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0",
            "subaccount_id": "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000000",
            "order_hash": "0x3b1bcc8ab01e1e0f1f2cf9de5f44267bed6368fabd62adbcf3826a207340194e",
            "cid": "order-123"
        }
    ],
    "eip712Wrapper": "v1"
}`

var testEIP712WrapperOutJSON = `{
	"types": {
		"Coin": [
			{
				"name": "denom",
				"type": "string"
			},
			{
				"name": "amount",
				"type": "string"
			}
		],
		"EIP712Domain": [
			{
				"name": "name",
				"type": "string"
			},
			{
				"name": "version",
				"type": "string"
			},
			{
				"name": "chainId",
				"type": "uint256"
			},
			{
				"name": "verifyingContract",
				"type": "string"
			},
			{
				"name": "salt",
				"type": "string"
			}
		],
		"Fee": [
			{
				"name": "amount",
				"type": "Coin[]"
			},
			{
				"name": "gas",
				"type": "string"
			}
		],
		"Msg": [
			{
				"name": "type",
				"type": "string"
			},
			{
				"name": "value",
				"type": "MsgValue"
			}
		],
		"MsgValue": [
			{
				"name": "sender",
				"type": "string"
			},
			{
				"name": "market_id",
				"type": "string"
			},
			{
				"name": "subaccount_id",
				"type": "string"
			},
			{
				"name": "order_hash",
				"type": "string"
			},
			{
				"name": "cid",
				"type": "string"
			}
		],
		"Tx": [
			{
				"name": "account_number",
				"type": "string"
			},
			{
				"name": "chain_id",
				"type": "string"
			},
			{
				"name": "fee",
				"type": "Fee"
			},
			{
				"name": "memo",
				"type": "string"
			},
			{
				"name": "msgs",
				"type": "Msg[]"
			},
			{
				"name": "sequence",
				"type": "string"
			},
			{
				"name": "timeout_height",
				"type": "string"
			}
		]
	},
	"primaryType": "Tx",
	"domain": {
		"name": "Injective Web3",
		"version": "1.0.0",
		"chainId": "0xaa36a7",
		"verifyingContract": "cosmos",
		"salt": "0"
	},
	"message": {
		"account_number": "3",
		"chain_id": "injective-777",
		"fee": {
			"amount": [
				{
					"amount": "160000000",
					"denom": "inj"
				}
			],
			"gas": "400000"
		},
		"memo": "",
		"msgs": [
			{
				"type": "exchange/MsgCancelSpotOrder",
				"value": {
					"cid": "order-123",
					"market_id": "0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0",
					"order_hash": "0x3b1bcc8ab01e1e0f1f2cf9de5f44267bed6368fabd62adbcf3826a207340194e",
					"sender": "inj14au322k9munkmx5wrchz9q30juf5wjgz2cfqku",
					"subaccount_id": "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000000"
				}
			}
		],
		"sequence": "0",
		"timeout_height": "999999999"
	}
}`

type prepareEip712Payload struct {
	// Specify chainID for the target tx
	ChainID uint64AsString `json:"chainId"`
	// Specify Ethereum address of a signer
	SignerAddress string `json:"signerAddress"`
	// Sequence number of the transaction signer
	Sequence uint64AsString `json:"sequence"`
	// Account number of the transaction signer
	AccountNumber uint64AsString `json:"accountNumber"`
	// Textual memo information to attach with tx
	Memo string `json:"memo"`
	// Block height until which the transaction is valid.
	TimeoutHeight uint64AsString `json:"timeoutHeight"`
	// Transaction fee details.
	Fee *cosmosTxFee `json:"fee"`
	// List of Cosmos proto3-encoded Msgs to include in a single tx
	Msgs []json.RawMessage `json:"msgs"`
	// The wrapper of the EIP712 message, 'v1'/'v2' or 'V1'/'V2'
	Eip712Wrapper string `json:"eip712Wrapper"`
}

type cosmosTxFee struct {
	// Transaction gas price
	Price []*cosmosCoin `json:"price"`
	// Transaction gas limit
	Gas uint64AsString `json:"gas"`
	// Explicitly require fee delegation when set to true. Or don't care = false.
	// Will be replaced by other flag in the next version.
	DelegateFee bool `json:"delegateFee"`
}

type cosmosCoin struct {
	// Coin denominator
	Denom string `json:"denom"`
	// Coin amount (big int)
	Amount string `json:"amount"`
}

type uint64AsString uint64

// MarshalJSON encodes uint64AsString as a string.
func (u uint64AsString) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(u), 10))
}

// UnmarshalJSON decodes a JSON string into uint64AsString.
func (u *uint64AsString) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*u = uint64AsString(v)
	return nil
}

type feeDelegationOptions struct {
	FeePayer cosmtypes.AccAddress
}
