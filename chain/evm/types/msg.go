package types

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	sdkmath "cosmossdk.io/math"
	"google.golang.org/protobuf/proto"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/ethereum/go-ethereum/common"
	cmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	_ sdk.Msg    = &MsgEthereumTx{}
	_ sdk.Tx     = &MsgEthereumTx{}
	_ ante.GasTx = &MsgEthereumTx{}
	_ sdk.Msg    = &MsgUpdateParams{}

	_ codectypes.UnpackInterfacesMessage = MsgEthereumTx{}
)

// message type and route constants
const (
	// TypeMsgEthereumTx defines the type string of an Ethereum transaction
	TypeMsgEthereumTx = "ethereum_tx"
)

func NewTxWithData(txData ethtypes.TxData) *MsgEthereumTx {
	return &MsgEthereumTx{Raw: NewEthereumTx(txData)}
}

// NewTx returns a reference to a new Ethereum transaction message.
func NewTx(
	chainID *big.Int, nonce uint64, to *common.Address, amount *big.Int,
	gasLimit uint64, gasPrice, gasFeeCap, gasTipCap *big.Int, input []byte, accesses *ethtypes.AccessList,
) *MsgEthereumTx {
	return newMsgEthereumTx(chainID, nonce, to, amount, gasLimit, gasPrice, gasFeeCap, gasTipCap, input, accesses)
}

// NewTxContract returns a reference to a new Ethereum transaction
// message designated for contract creation.
func NewTxContract(
	chainID *big.Int,
	nonce uint64,
	amount *big.Int,
	gasLimit uint64,
	gasPrice, gasFeeCap, gasTipCap *big.Int,
	input []byte,
	accesses *ethtypes.AccessList,
) *MsgEthereumTx {
	return newMsgEthereumTx(chainID, nonce, nil, amount, gasLimit, gasPrice, gasFeeCap, gasTipCap, input, accesses)
}

func newMsgEthereumTx(
	chainID *big.Int, nonce uint64, to *common.Address, amount *big.Int,
	gasLimit uint64, gasPrice, gasFeeCap, gasTipCap *big.Int, input []byte, accesses *ethtypes.AccessList,
) *MsgEthereumTx {
	var txData ethtypes.TxData

	switch {
	case gasFeeCap != nil:
		var accessList ethtypes.AccessList
		if accesses != nil {
			accessList = *accesses
		}
		txData = &ethtypes.DynamicFeeTx{
			ChainID:    chainID,
			Nonce:      nonce,
			To:         to,
			Value:      amount,
			Gas:        gasLimit,
			GasTipCap:  gasTipCap,
			GasFeeCap:  gasFeeCap,
			Data:       input,
			AccessList: accessList,
		}
	case accesses != nil:
		txData = &ethtypes.AccessListTx{
			ChainID:    chainID,
			Nonce:      nonce,
			To:         to,
			Value:      amount,
			Gas:        gasLimit,
			GasPrice:   gasPrice,
			Data:       input,
			AccessList: *accesses,
		}
	default:
		txData = &ethtypes.LegacyTx{
			Nonce:    nonce,
			To:       to,
			Value:    amount,
			Gas:      gasLimit,
			GasPrice: gasPrice,
			Data:     input,
		}
	}

	return NewTxWithData(txData)
}

func (msg *MsgEthereumTx) FromEthereumTx(tx *ethtypes.Transaction) {
	msg.Raw.Transaction = tx
}

// FromSignedEthereumTx populates the message fields from the given signed ethereum transaction, and set From field.
func (msg *MsgEthereumTx) FromSignedEthereumTx(tx *ethtypes.Transaction, signer ethtypes.Signer) error {
	msg.Raw.Transaction = tx

	from, err := ethtypes.Sender(signer, tx)
	if err != nil {
		return err
	}

	msg.From = from.Bytes()
	return nil
}

// Route returns the route value of an MsgEthereumTx.
func (msg MsgEthereumTx) Route() string { return RouterKey }

// Type returns the type value of an MsgEthereumTx.
func (msg MsgEthereumTx) Type() string { return TypeMsgEthereumTx }

// ValidateBasic implements the sdk.Msg interface. It performs basic validation
// checks of a Transaction. If returns an error if validation fails.
func (msg MsgEthereumTx) ValidateBasic() error {
	// From and Raw are only two fields allowed in new transaction format.
	if len(msg.From) == 0 {
		return errorsmod.Wrapf(errortypes.ErrInvalidRequest, "sender address is missing")
	}
	if msg.Raw.Transaction == nil {
		return errorsmod.Wrapf(errortypes.ErrInvalidRequest, "raw tx is missing")
	}

	// Check removed fields not exists
	if len(msg.DeprecatedFrom) != 0 {
		return errorsmod.Wrapf(errortypes.ErrInvalidRequest, "deprecated From field is not empty")
	}
	if len(msg.DeprecatedHash) != 0 {
		return errorsmod.Wrapf(errortypes.ErrInvalidRequest, "deprecated Hash field is not empty")
	}
	// Validate Size_ field, should be kept empty
	if msg.Size_ != 0 {
		return errorsmod.Wrapf(errortypes.ErrInvalidRequest, "tx size is deprecated")
	}
	if msg.Data != nil {
		return errorsmod.Wrapf(errortypes.ErrInvalidRequest, "tx data is deprecated in favor of Raw")
	}

	if err := msg.Raw.Validate(); err != nil {
		return err
	}

	return nil
}

// GetMsgs returns a single MsgEthereumTx as an sdk.Msg.
func (msg *MsgEthereumTx) GetMsgs() []sdk.Msg {
	return []sdk.Msg{msg}
}

func (msg *MsgEthereumTx) GetMsgsV2() ([]proto.Message, error) {
	return nil, errors.New("not implemented")
}

// GetSender convert the From field to common.Address
// From should always be set, which is validated in ValidateBasic
func (msg *MsgEthereumTx) GetSender() common.Address {
	return common.BytesToAddress(msg.From)
}

// GetSenderLegacy fallbacks to old behavior if From is empty, should be used by json-rpc
func (msg *MsgEthereumTx) GetSenderLegacy(signer ethtypes.Signer) (common.Address, error) {
	if len(msg.From) > 0 {
		return msg.GetSender(), nil
	}
	sender, err := msg.recoverSender(signer)
	if err != nil {
		return common.Address{}, err
	}
	msg.From = sender.Bytes()
	return sender, nil
}

// recoverSender recovers the sender address from the transaction signature.
func (msg *MsgEthereumTx) recoverSender(signer ethtypes.Signer) (common.Address, error) {
	return ethtypes.Sender(signer, msg.AsTransaction())
}

// GetSignBytes returns the Amino bytes of an Ethereum transaction message used
// for signing.
//
// NOTE: This method cannot be used as a chain ID is needed to create valid bytes
// to sign over. Use 'RLPSignBytes' instead.
func (msg MsgEthereumTx) GetSignBytes() []byte {
	panic("must use 'RLPSignBytes' with a chain ID to get the valid bytes to sign")
}

// Sign calculates a secp256k1 ECDSA signature and signs the transaction. It
// takes a keyring signer and the chainID to sign an Ethereum transaction according to
// EIP155 standard.
// This method mutates the transaction as it populates the V, R, S
// fields of the Transaction's Signature.
// The function will fail if the sender address is not defined for the msg or if
// the sender is not registered on the keyring
func (msg *MsgEthereumTx) Sign(ethSigner ethtypes.Signer, keyringSigner keyring.Signer) error {
	from := msg.GetFrom()
	if from.Empty() {
		return fmt.Errorf("sender address not defined for message")
	}

	tx := msg.AsTransaction()
	txHash := ethSigner.Hash(tx)

	sig, _, err := keyringSigner.SignByAddress(from, txHash.Bytes(), signing.SignMode_SIGN_MODE_TEXTUAL)
	if err != nil {
		return err
	}

	tx, err = tx.WithSignature(ethSigner, sig)
	if err != nil {
		return err
	}

	msg.Raw.Transaction = tx
	return nil
}

// GetGas implements the GasTx interface. It returns the GasLimit of the transaction.
func (msg MsgEthereumTx) GetGas() uint64 {
	return msg.AsTransaction().Gas()
}

// GetFee returns the fee for non dynamic fee tx
func (msg MsgEthereumTx) GetFee() *big.Int {
	tx := msg.AsTransaction()
	price := tx.GasPrice()
	return price.Mul(price, new(big.Int).SetUint64(tx.Gas()))
}

// GetEffectiveFee returns the fee for dynamic fee tx
func (msg MsgEthereumTx) GetEffectiveFee(baseFee *big.Int) *big.Int {
	price := msg.GetEffectiveGasPrice(baseFee)
	return price.Mul(price, new(big.Int).SetUint64(msg.GetGas()))
}

// GetEffectiveGasPrice returns the fee for dynamic fee tx
func (msg MsgEthereumTx) GetEffectiveGasPrice(baseFee *big.Int) *big.Int {
	tx := msg.AsTransaction()
	if baseFee == nil {
		return tx.GasPrice()
	}
	// for legacy tx, both gasTipCap and gasFeeCap are gasPrice, the result is equavalent.
	return cmath.BigMin(new(big.Int).Add(tx.GasTipCap(), baseFee), tx.GasFeeCap())
}

// GetFrom loads the ethereum sender address from the sigcache and returns an
// sdk.AccAddress from its bytes
func (msg *MsgEthereumTx) GetFrom() sdk.AccAddress {
	return sdk.AccAddress(msg.From)
}

// AsTransaction creates an Ethereum Transaction type from the msg fields
func (msg *MsgEthereumTx) AsTransaction() *ethtypes.Transaction {
	tx := msg.Raw.Transaction
	if tx != nil {
		return tx
	}

	// fallback to legacy format
	txData, err := UnpackTxData(msg.Data)
	if err != nil {
		return nil
	}
	msg.Raw = NewEthereumTx(txData.AsEthereumData())
	return msg.Raw.Transaction
}

// AsMessage creates an Ethereum core.Message from the msg fields
func (msg *MsgEthereumTx) AsMessage(baseFee *big.Int) *core.Message {
	tx := msg.AsTransaction()
	ethMsg := &core.Message{
		Nonce:             tx.Nonce(),
		GasLimit:          tx.Gas(),
		GasPrice:          new(big.Int).Set(tx.GasPrice()),
		GasFeeCap:         new(big.Int).Set(tx.GasFeeCap()),
		GasTipCap:         new(big.Int).Set(tx.GasTipCap()),
		To:                tx.To(),
		Value:             tx.Value(),
		Data:              tx.Data(),
		AccessList:        tx.AccessList(),
		SkipAccountChecks: false,

		From: common.BytesToAddress(msg.From),
	}
	// If baseFee provided, set gasPrice to effectiveGasPrice.
	if baseFee != nil {
		ethMsg.GasPrice = cmath.BigMin(ethMsg.GasPrice.Add(ethMsg.GasTipCap, baseFee), ethMsg.GasFeeCap)
	}
	return ethMsg
}

// VerifySender verify the sender address against the signature values using the latest signer for the given chainID.
func (msg *MsgEthereumTx) VerifySender(signer ethtypes.Signer) error {
	from, err := msg.recoverSender(signer)
	if err != nil {
		return err
	}

	if !bytes.Equal(msg.From, from.Bytes()) {
		return fmt.Errorf("sender verification failed. got %s, expected %s", HexAddress(from.Bytes()), HexAddress(msg.From))
	}
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (msg MsgEthereumTx) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return unpacker.UnpackAny(msg.Data, new(TxData))
}

// UnmarshalBinary decodes the canonical encoding of transactions.
func (msg *MsgEthereumTx) UnmarshalBinary(b []byte, signer ethtypes.Signer) error {
	tx := &ethtypes.Transaction{}
	if err := tx.UnmarshalBinary(b); err != nil {
		return err
	}
	return msg.FromSignedEthereumTx(tx, signer)
}

func (msg *MsgEthereumTx) Hash() common.Hash {
	return msg.AsTransaction().Hash()
}

// BuildTx builds the canonical cosmos tx from ethereum msg
func (msg *MsgEthereumTx) BuildTx(b client.TxBuilder, evmDenom string) (authsigning.Tx, error) {
	builder, ok := b.(authtx.ExtensionOptionsTxBuilder)
	if !ok {
		return nil, errors.New("unsupported builder")
	}

	option, err := codectypes.NewAnyWithValue(&ExtensionOptionsEthereumTx{})
	if err != nil {
		return nil, err
	}

	fees := make(sdk.Coins, 0)
	fee := msg.GetFee()
	feeAmt := sdkmath.NewIntFromBigInt(fee)
	if feeAmt.Sign() > 0 {
		fees = append(fees, sdk.NewCoin(evmDenom, feeAmt))
	}

	builder.SetExtensionOptions(option)

	if err := builder.SetMsgs(&MsgEthereumTx{
		From: msg.From,
		Raw:  msg.Raw,
	}); err != nil {
		return nil, err
	}
	builder.SetFeeAmount(fees)
	builder.SetGasLimit(msg.GetGas())
	return builder.GetTx(), nil
}

// ValidateBasic does a sanity check of the provided data
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}

	return m.Params.Validate()
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&m))
}
