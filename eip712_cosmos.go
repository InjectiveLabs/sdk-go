package sdk

import (
	"context"
	"encoding/json"
	"fmt"

	basev1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	"cosmossdk.io/x/tx/signing"
	"cosmossdk.io/x/tx/signing/aminojson"
	"github.com/cosmos/cosmos-proto/anyutil"
	"github.com/cosmos/cosmos-sdk/codec"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"

	"github.com/InjectiveLabs/sdk-go/typeddata"
)

type EIP712Wrapper func(
	cdc codec.ProtoCodecMarshaler,
	tx authsigning.Tx,
	signerData *authsigning.SignerData,
	opts Web3ExtensionOptions,
) (typeddata.TypedData, error)

// WrapTxToEIP712LegacyAmino is the method previously used to generate EIP712 typed data (until sdk v0.50)
func WrapTxToEIP712LegacyAmino(
	cdc codec.ProtoCodecMarshaler,
	tx authsigning.Tx,
	signerData *authsigning.SignerData,
	opts Web3ExtensionOptions,
) (typeddata.TypedData, error) {

	data := legacytx.StdSignBytes(
		signerData.ChainID,
		signerData.AccountNumber,
		signerData.Sequence,
		tx.GetTimeoutHeight(),
		legacytx.StdFee{Amount: tx.GetFee(), Gas: tx.GetGas()},
		tx.GetMsgs(),
		tx.GetMemo(),
	)

	return WrapTxToEIP712WithSignBytes(data, cdc, opts.ChainID, tx.GetMsgs(), opts.FeePayer, tx.GetTimeoutHeight())
}

// WrapTxToEIP712AminoJSON is an ultimate method that wraps aminojson-encoded Cosmos Tx JSON data
// into an EIP712-compatible request. All messages must be of the same type.
func WrapTxToEIP712AminoJSON(
	cdc codec.ProtoCodecMarshaler,
	tx authsigning.Tx,
	authSignerData *authsigning.SignerData,
	opts Web3ExtensionOptions,
) (typeddata.TypedData, error) {
	signerData := signing.SignerData{
		ChainID:       authSignerData.ChainID,
		AccountNumber: authSignerData.AccountNumber,
		Sequence:      authSignerData.Sequence,
		Address:       authSignerData.Address,
	}

	txBody := &txv1beta1.TxBody{
		Memo:          tx.GetMemo(),
		TimeoutHeight: tx.GetTimeoutHeight(),
	}

	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		anyMsg, err := anyutil.New(protoadapt.MessageV2Of(msg))
		if err != nil {
			return typeddata.TypedData{}, fmt.Errorf("can't wrap msg as Any after protoadaptV2Of, msg: %v, err: %w", msg, err)
		}

		txBody.Messages = append(txBody.Messages, anyMsg)
	}

	fee := &txv1beta1.Fee{
		GasLimit: tx.GetGas(),
	}

	for _, feeAmount := range tx.GetFee() {
		coin := basev1beta1.Coin{Denom: feeAmount.Denom, Amount: feeAmount.Amount.String()}
		fee.Amount = append(fee.Amount, &coin)
	}

	bodyBz, err := proto.MarshalOptions{Deterministic: true}.Marshal(txBody)
	if err != nil {
		return typeddata.TypedData{}, fmt.Errorf("can't proto marshal txBody, err: %w", err)
	}

	txData := signing.TxData{
		Body: txBody,
		AuthInfo: &txv1beta1.AuthInfo{
			Fee: fee,
		},
		BodyBytes: bodyBz,
	}

	handler := aminojson.NewSignModeHandler(aminojson.SignModeHandlerOptions{})
	data, err := handler.GetSignBytes(context.Background(), signerData, txData)
	if err != nil {
		return typeddata.TypedData{}, fmt.Errorf("can't GetSignBytes using aminojson, err: %w", err)
	}

	return WrapTxToEIP712WithSignBytes(data, cdc, opts.ChainID, msgs, opts.FeePayer, tx.GetTimeoutHeight())
}

// WrapTxToEIP712WithSignBytes wraps provided signBytes to EIP712 TypedData
func WrapTxToEIP712WithSignBytes(
	signBytes []byte,
	cdc codec.ProtoCodecMarshaler,
	chainID int64,
	msgs []cosmtypes.Msg,
	feePayer cosmtypes.AccAddress,
	timeoutHeight uint64,
) (typeddata.TypedData, error) {
	txData := make(map[string]interface{})
	if err := json.Unmarshal(signBytes, &txData); err != nil {
		err = fmt.Errorf("failed to unmarshal data provided into WrapTxToEIP712: %w", err)
		return typeddata.TypedData{}, err
	}

	domain := typeddata.TypedDataDomain{
		Name:              "Injective Web3",
		Version:           "1.0.0",
		ChainId:           ethmath.NewHexOrDecimal256(chainID),
		VerifyingContract: "cosmos",
		Salt:              "0",
	}

	msgTypes, err := typeddata.ExtractMsgTypes(cdc, "MsgValue", msgs[0])
	if err != nil {
		return typeddata.TypedData{}, err
	}

	if feePayer != nil {
		feeInfo := txData["fee"].(map[string]interface{})
		feeInfo["feePayer"] = feePayer.String()

		// also patching msgTypes to include feePayer
		msgTypes["Fee"] = []typeddata.Type{
			{Name: "feePayer", Type: "string"},
			{Name: "amount", Type: "Coin[]"},
			{Name: "gas", Type: "string"},
		}
	}

	// in case timeoutHeight is 0 it will be missing in txData
	// and ComputeTypedDataAndHash will fail after this call
	if timeoutHeight == 0 {
		txData["timeout_height"] = "0"
	}

	td := typeddata.TypedData{
		Types:       msgTypes,
		PrimaryType: "Tx",
		Domain:      domain,
		Message:     txData,
	}

	return td, nil
}

func WrapTxToEIP712V2(
	cdc codec.ProtoCodecMarshaler,
	tx authsigning.Tx,
	signerData *authsigning.SignerData,
	opts Web3ExtensionOptions,
) (typeddata.TypedData, error) {
	domain := typeddata.TypedDataDomain{
		Name:              "Injective Web3",
		Version:           "1.0.0",
		ChainId:           ethmath.NewHexOrDecimal256(opts.ChainID),
		VerifyingContract: "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
		Salt:              "0",
	}

	msgs := tx.GetMsgs()
	msgsJsons := make([]json.RawMessage, len(msgs))
	for idx, m := range msgs {
		bzMsg, err := cdc.MarshalInterfaceJSON(m)
		if err != nil {
			return typeddata.TypedData{}, fmt.Errorf("cannot marshal json at index %d: %w", idx, err)
		}

		msgsJsons[idx] = bzMsg
	}

	bzMsgs, err := json.Marshal(msgsJsons)
	if err != nil {
		return typeddata.TypedData{}, fmt.Errorf("marshal json err: %w", err)
	}

	feeInfo := legacytx.StdFee{
		Amount: tx.GetFee(),
		Gas:    tx.GetGas(),
	}

	if opts.FeePayer != nil {
		feeInfo.Payer = opts.FeePayer.String()
	}

	bzFee, err := json.Marshal(feeInfo)
	if err != nil {
		return typeddata.TypedData{}, fmt.Errorf("marshal fee info failed: %w", err)
	}

	ctx := map[string]interface{}{
		"account_number": signerData.AccountNumber,
		"sequence":       signerData.Sequence,
		"timeout_height": tx.GetTimeoutHeight(),
		"chain_id":       signerData.ChainID,
		"memo":           tx.GetMemo(),
		"fee":            json.RawMessage(bzFee),
	}

	bzTxContext, err := json.Marshal(ctx)
	if err != nil {
		return typeddata.TypedData{}, fmt.Errorf("marshal json err: %w", err)
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

	return td, nil
}
