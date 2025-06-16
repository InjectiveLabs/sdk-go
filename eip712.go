package sdk

import (
	"encoding/json"
	"fmt"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ethsecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"

	injcodectypes "github.com/InjectiveLabs/sdk-go/chain/codec/types"
	secp256k1 "github.com/InjectiveLabs/sdk-go/chain/crypto/ethsecp256k1"
	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"
	"github.com/InjectiveLabs/sdk-go/typeddata"
)

var (
	chainTypesCodec codec.ProtoCodecMarshaler
	GlobalCdc       = codec.NewProtoCodec(injcodectypes.NewInterfaceRegistry())
)

func init() {
	registry := injcodectypes.NewInterfaceRegistry()
	chaintypes.RegisterInterfaces(registry)
	chainTypesCodec = codec.NewProtoCodec(registry)
}

// Verify all signatures for a tx and return an error if any are invalid. Note,
// the Eip712SigVerificationDecorator decorator will not get executed on ReCheck.
//
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type Eip712SigVerificationDecorator struct {
	ak authante.AccountKeeper
}

func NewEip712SigVerificationDecorator(ak authante.AccountKeeper) Eip712SigVerificationDecorator {
	return Eip712SigVerificationDecorator{
		ak: ak,
	}
}

func (svd Eip712SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// no need to verify signatures on recheck tx
	if ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, errors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// stdSigs contains the sequence number, account number, and signatures.
	// When simulating, this would just be a 0-length slice.
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return ctx, err
	}

	signerAddrs, _ := sigTx.GetSigners()

	// check that signer length and signature length are the same
	if len(sigs) != len(signerAddrs) {
		return ctx, errors.Wrapf(sdkerrors.ErrUnauthorized, "invalid number of signer;  expected: %d, got %d", len(signerAddrs), len(sigs))
	}

	for i, sig := range sigs {
		acc, err := authante.GetSignerAcc(ctx, svd.ak, signerAddrs[i])
		if err != nil {
			return ctx, err
		}

		// retrieve pubkey
		pubKey := acc.GetPubKey()
		if !simulate && pubKey == nil {
			return ctx, errors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
		}

		// Check account sequence number.
		if sig.Sequence != acc.GetSequence() {
			return ctx, errors.Wrapf(
				sdkerrors.ErrWrongSequence,
				"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
			)
		}

		// retrieve signer data
		genesis := ctx.BlockHeight() == 0
		chainID := ctx.ChainID()
		var accNum uint64
		if !genesis {
			accNum = acc.GetAccountNumber()
		}
		signerData := authsigning.SignerData{
			Address:       acc.GetAddress().String(),
			ChainID:       chainID,
			AccountNumber: accNum,
			Sequence:      acc.GetSequence(),
		}

		if !simulate {
			typedData, err := GenerateTypedDataAndVerifySignatureEIP712(pubKey, signerData, sig.Data, tx.(authsigning.Tx))
			if err != nil {
				ctx.Logger().Error("Eip712SigVerificationDecorator failed to verify signature", "error", err)
				errMsg := fmt.Sprintf("signature verification failed: %s; please verify account number (%d) and chain-id (%s)", err.Error(), accNum, chainID)
				if typedData != nil {
					bz, _ := json.Marshal(typedData)
					errMsg = fmt.Sprintf("%s, eip712: %s", errMsg, string(bz))
				}

				return ctx, errors.Wrap(sdkerrors.ErrUnauthorized, errMsg)
			}
		}
	}

	return next(ctx, tx, simulate)
}

// VerifySignature verifies a transaction signature contained in SignatureData abstracting over different signing modes
// and single vs multi-signatures.
func GenerateTypedDataAndVerifySignatureEIP712(
	pubKey cryptotypes.PubKey,
	signerData authsigning.SignerData,
	sigData signing.SignatureData,
	tx authsigning.Tx,
) (*typeddata.TypedData, error) {
	data, ok := sigData.(*signing.SingleSignatureData)
	if !ok {
		return nil, fmt.Errorf("unexpected SignatureData %T", sigData)
	}

	var eip712Wrapper EIP712Wrapper
	switch data.SignMode {
	case signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON:
		eip712Wrapper = WrapTxToEIP712AminoJSON
	case signing.SignMode_SIGN_MODE_EIP712_V2:
		eip712Wrapper = WrapTxToEIP712V2
	default:
		return nil, fmt.Errorf("unexpected SignatureData %T: wrong SignMode: %v", sigData, data.SignMode)
	}

	opts, err := GetWeb3ExtensionOptions(tx, signerData)
	if err != nil {
		return nil, err
	}

	typedData, err := eip712Wrapper(GlobalCdc, tx, &signerData, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack tx data in EIP712 object")
	}

	sigHash, _, err := typeddata.ComputeTypedDataAndHash(typedData)
	if err != nil {
		return &typedData, err
	}

	if feeDelegated := opts.FeePayer != nil && opts.FeePayerSig != nil; feeDelegated {
		feePayerPubkey, err := ethsecp256k1.RecoverPubkey(sigHash, opts.FeePayerSig)
		if err != nil {
			return &typedData, errors.Wrap(err, "failed to recover delegated fee payer from sig")
		}

		ecPubKey, err := ethcrypto.UnmarshalPubkey(feePayerPubkey)
		if err != nil {
			return &typedData, errors.Wrap(err, "failed to unmarshal recovered fee payer pubkey")
		}

		recoveredFeePayerAcc := sdk.AccAddress((&secp256k1.PubKey{Key: ethcrypto.CompressPubkey(ecPubKey)}).Address().Bytes())
		if !recoveredFeePayerAcc.Equals(opts.FeePayer) {
			return &typedData, fmt.Errorf("failed to verify delegated fee payer sig")
		}
	}

	if len(data.Signature) != 65 {
		return &typedData, fmt.Errorf("signature length doesnt match typical [R||S||V] signature 65 bytes")
	}

	// VerifySignature of secp256k1 accepts 64 byte signature [R||S]
	// WARNING! Under NO CIRCUMSTANCES try to use pubKey.VerifySignature there
	if !ethsecp256k1.VerifySignature(pubKey.Bytes(), sigHash, data.Signature[:64]) {
		return &typedData, fmt.Errorf("unable to verify signer signature of EIP712 typed data")
	}

	return &typedData, nil
}

type Web3ExtensionOptions struct {
	ChainID     int64
	FeePayer    sdk.AccAddress
	FeePayerSig []byte
}

func GetWeb3ExtensionOptions(tx authsigning.Tx, signerData authsigning.SignerData) (Web3ExtensionOptions, error) {
	txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx)
	if !ok {
		return Web3ExtensionOptions{}, nil
	}

	anyOpts := txWithExtensions.GetExtensionOptions()
	if len(anyOpts) == 0 {
		return Web3ExtensionOptions{}, nil
	}

	var optIface txtypes.TxExtensionOptionI
	if err := chainTypesCodec.UnpackAny(anyOpts[0], &optIface); err != nil {
		return Web3ExtensionOptions{}, errors.Wrap(err, "failed to proto-unpack Web3Extension")
	}

	extOpts, ok := optIface.(*chaintypes.ExtensionOptionsWeb3Tx)
	if !ok {
		return Web3ExtensionOptions{}, nil
	}

	// chainID in EIP712 typed data is allowed to not match signerData.ChainID,
	// but limited to certain options: 1 (Ethereum Mainnet), 11155111 (Ethereum Sepolia),
	// 1776 (Injective EVM Mainnet), 1439 (Injective EVM Testnet), thus Metamask will be
	// able to submit signatures without switching networks.
	hasValidMainnetChainID := signerData.ChainID == "injective-1" && (extOpts.TypedDataChainID == 1 || extOpts.TypedDataChainID == 1776)
	hasValidNonMainnetChainID := (signerData.ChainID == "injective-777" || signerData.ChainID == "injective-888") && (extOpts.TypedDataChainID == 11155111 || extOpts.TypedDataChainID == 1439)

	if !hasValidMainnetChainID && !hasValidNonMainnetChainID {
		return Web3ExtensionOptions{}, fmt.Errorf("invalid TypedDataChainID in Web3Extension: %d for %s chain", extOpts.TypedDataChainID, signerData.ChainID)
	}

	chainID := int64(extOpts.TypedDataChainID)

	var (
		feePayer    sdk.AccAddress
		feePayerSig []byte
	)

	if extOpts.FeePayer != "" {
		payer, err := sdk.AccAddressFromBech32(extOpts.FeePayer)
		if err != nil {
			return Web3ExtensionOptions{}, errors.Wrap(err, "failed to parse feePayer in Web3Extension")
		}

		sig := extOpts.FeePayerSig
		if len(sig) == 0 {
			return Web3ExtensionOptions{}, fmt.Errorf("no feePayerSig provided in Web3Extension")
		}

		feePayer = payer
		feePayerSig = sig
	}

	return Web3ExtensionOptions{
		ChainID:     chainID,
		FeePayer:    feePayer,
		FeePayerSig: feePayerSig,
	}, nil
}
