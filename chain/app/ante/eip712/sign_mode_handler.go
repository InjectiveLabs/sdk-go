package eip712

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	ethmath "github.com/ethereum/go-ethereum/common/math"

	"github.com/InjectiveLabs/sdk-go/typeddata"
)

// todo: eventually, refactor the whole ante pkg

// SignModeHandler is used to provide a way to verify EIP712 signatures from within Cosmos SDK.
// This gives multisig accounts the ability to send transactions with Ledger signatures
type SignModeHandler struct {
	cdc codec.Codec
}

func NewSignModeHandler(cdc codec.Codec) SignModeHandler {
	return SignModeHandler{cdc: cdc}
}

func (SignModeHandler) Mode() signingv1beta1.SignMode {
	return signingv1beta1.SignMode_SIGN_MODE_EIP712_V2
}

func (h SignModeHandler) GetSignBytes(
	_ context.Context,
	signerData signing.SignerData,
	txData signing.TxData,
) ([]byte, error) {
	msgsJSON := make([]json.RawMessage, len(txData.Body.Messages))
	for idx, anyPB := range txData.Body.Messages {
		anyCDC := &codectypes.Any{TypeUrl: anyPB.TypeUrl, Value: anyPB.Value}

		var msg sdk.Msg
		if err := h.cdc.UnpackAny(anyCDC, &msg); err != nil {
			return nil, fmt.Errorf("cannot unpack msg at index %d: %w", idx, err)
		}

		msgJSON, err := h.cdc.MarshalInterfaceJSON(msg)
		if err != nil {
			return nil, fmt.Errorf("cannot marshal json at index %d: %w", idx, err)
		}

		msgsJSON[idx] = msgJSON
	}

	bzMsgs, err := json.Marshal(msgsJSON)
	if err != nil {
		return nil, fmt.Errorf("marshal json err: %w", err)
	}

	feeInfo, err := StdFeeFromTxData(txData)
	if err != nil {
		return nil, err
	}

	bzFee, err := json.Marshal(feeInfo)
	if err != nil {
		return nil, fmt.Errorf("marshal fee info failed: %w", err)
	}

	ctx := map[string]any{
		"account_number": signerData.AccountNumber,
		"sequence":       signerData.Sequence,
		"timeout_height": txData.Body.TimeoutHeight,
		"chain_id":       signerData.ChainID,
		"memo":           txData.Body.Memo,
		"fee":            json.RawMessage(bzFee),
	}

	bzTxContext, err := json.Marshal(ctx)
	if err != nil {
		return nil, fmt.Errorf("marshal json err: %w", err)
	}

	chainID := int64(1) // default: injective-1 == 1 (eth)
	if signerData.ChainID != "injective-1" {
		chainID = 11155111 // injective-777 == 11155111 (sepolia)
	}

	domain := typeddata.TypedDataDomain{
		Name:              "Injective Web3",
		Version:           "1.0.0",
		ChainId:           ethmath.NewHexOrDecimal256(chainID),
		VerifyingContract: "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
		Salt:              "0",
	}

	td := typeddata.TypedData{
		Types:       typeddata.SignableTypes(),
		PrimaryType: "Tx",
		Domain:      domain,
		Message: typeddata.TypedDataMessage{
			"context": string(bzTxContext),
			"msgs":    string(bzMsgs),
		},
	}

	_, raw, err := typeddata.ComputeTypedDataAndHash(td)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func StdFeeFromTxData(txData signing.TxData) (legacytx.StdFee, error) {
	if txData.AuthInfo == nil || txData.AuthInfo.Fee == nil {
		return legacytx.StdFee{}, errors.New("missing auth info fee")
	}

	amount := make(sdk.Coins, 0, len(txData.AuthInfo.Fee.Amount))
	for i, c := range txData.AuthInfo.Fee.Amount {
		amt, ok := sdkmath.NewIntFromString(c.Amount)
		if !ok {
			return legacytx.StdFee{}, fmt.Errorf("invalid fee amount at index %d: %q", i, c.Amount)
		}

		coin := sdk.Coin{Denom: c.Denom, Amount: amt}
		if err := coin.Validate(); err != nil {
			return legacytx.StdFee{}, fmt.Errorf("invalid fee coin at index %d: %w", i, err)
		}

		amount = append(amount, coin)
	}

	if !amount.IsValid() {
		return legacytx.StdFee{}, errors.New("invalid fee coins")
	}

	fee := legacytx.StdFee{
		Amount: amount.Sort(),
		Gas:    txData.AuthInfo.Fee.GasLimit,
		Payer:  txData.AuthInfo.Fee.Payer,
		// Granter: txData.AuthInfo.Fee.Granter, // eip712_cosmos.go does not include Granter
	}

	return fee, nil
}
