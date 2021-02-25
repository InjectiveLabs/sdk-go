package types

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	zeroex "github.com/InjectiveLabs/sdk-go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgRegisterDerivativeMarket{}
	_ sdk.Msg = &MsgSuspendDerivativeMarket{}
	_ sdk.Msg = &MsgResumeDerivativeMarket{}
	_ sdk.Msg = &MsgRegisterDerivativeMarket{}
	_ sdk.Msg = &MsgCreateDerivativeOrder{}
	_ sdk.Msg = &MsgSoftCancelDerivativeOrder{}
	_ sdk.Msg = &MsgRegisterSpotMarket{}
	_ sdk.Msg = &MsgSuspendSpotMarket{}
	_ sdk.Msg = &MsgResumeSpotMarket{}
	_ sdk.Msg = &MsgCreateSpotOrder{}
	_ sdk.Msg = &MsgExecuteDerivativeTakeOrder{}
	_ sdk.Msg = &MsgExecuteTECTransaction{}
	_ sdk.Msg = &MsgInitExchange{}
)

func (msg *MsgInitExchange) Route() string {
	return RouterKey
}

func (msg *MsgInitExchange) Type() string {
	return "msgInitExchange"
}

func (msg *MsgInitExchange) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender.String())
	} else if !common.IsHexAddress(msg.ExchangeAddress) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ExchangeAddress)
	}
	return nil
}

func (msg *MsgInitExchange) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgInitExchange) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg *MsgSoftCancelDerivativeOrder) Route() string {
	return RouterKey
}

func (msg *MsgSoftCancelDerivativeOrder) Type() string {
	return "softCancelDerivativeOrder"
}

func (msg *MsgSoftCancelDerivativeOrder) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order == nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no make order specified")
	}

	order := msg.Order.ToSignedOrder()
	quantity := order.TakerAssetAmount
	price := order.MakerAssetAmount
	orderHash, err := order.ComputeOrderHash()
	makerAddress := common.HexToAddress(msg.Order.MakerAddress)

	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("hash check failed: %v", err))
	} else if !isValidSignature(msg.Order.Signature, makerAddress, orderHash) {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid signature")
	} else if quantity == nil || quantity.Cmp(big.NewInt(0)) <= 0 {
		return sdkerrors.Wrap(ErrInsufficientOrderQuantity, "insufficient quantity")
	} else if price == nil || price.Cmp(big.NewInt(0)) <= 0 {
		return sdkerrors.Wrap(ErrInsufficientOrderQuantity, "insufficient price")
	}

	return nil
}

func (msg *MsgSoftCancelDerivativeOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgSoftCancelDerivativeOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgCreateSpotOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateSpotOrder) Type() string { return "createSpotOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateSpotOrder) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	sOrder := msg.Order.ToSignedOrder()

	if msg.Order == nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no make order specified")
	}

	orderHash, err := msg.Order.ToSignedOrder().ComputeOrderHash()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("hash check failed: %v", err))
	}
	makerAddress, _ := sOrder.GetMakerAddressAndNonce()

	// TODO: also allow for exchange without signatures (just cosmos)
	if !isValidSignature(msg.Order.Signature, makerAddress, orderHash) {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid signature")
	}

	return nil
}

// isValidSignature checks that the signature of the order is correct
func isValidSignature(sig string, makerAddr common.Address, hash common.Hash) bool {
	signature := common.FromHex(sig)
	signatureType := zeroex.SignatureType(signature[len(signature)-1])
	switch signatureType {
	case zeroex.InvalidSignature, zeroex.IllegalSignature:
		return false

	case zeroex.EIP712Signature:
		if len(signature) != 66 {
			return false
		}

		v := signature[0]
		r := signature[1:33]
		s := signature[33:65]

		//build Eth signature
		EthSignature := make([]byte, 65)
		copy(EthSignature[0:32], r)
		copy(EthSignature[32:64], s)
		EthSignature[64] = v - 27

		//validate signature
		pubKey, err := ethcrypto.SigToPub(hash[:], EthSignature)
		if err != nil {
			return false
		}

		//compare recoveredAddr with makerAddress
		recoveredAddr := ethcrypto.PubkeyToAddress(*pubKey)
		return bytes.Equal(recoveredAddr.Bytes(), makerAddr.Bytes())
	case zeroex.EthSignSignature:
		if len(signature) != 66 {
			return false
		}

		//validate signature
		EthSignature := signature[:65]
		EthSignature[64] = EthSignature[64] - 27
		pubKey, err := ethcrypto.SigToPub(hash[:], EthSignature)
		if err != nil {
			return false
		}

		//compare recoveredAddr with makerAddress
		recoveredAddr := ethcrypto.PubkeyToAddress(*pubKey)
		return bytes.Equal(recoveredAddr.Bytes(), makerAddr.Bytes())
	case zeroex.ValidatorSignature:
		if len(signature) < 21 {
			return false
		}
		// TODO: not supported yet
		return false
	case zeroex.PreSignedSignature, zeroex.WalletSignature, zeroex.EIP1271WalletSignature:
		// TODO: not supported yet
		return false
	default:
		return false
	}
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateSpotOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateSpotOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgCreateDerivativeOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateDerivativeOrder) Type() string { return "createDerivativeOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateDerivativeOrder) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order == nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no make order specified")
	}

	order := msg.Order.ToSignedOrder()
	quantity := order.TakerAssetAmount
	price := order.MakerAssetAmount
	orderHash, err := order.ComputeOrderHash()
	makerAddress := common.HexToAddress(msg.Order.MakerAddress)

	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("hash check failed: %v", err))
	} else if !isValidSignature(msg.Order.Signature, makerAddress, orderHash) {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid signature")
	} else if quantity == nil || quantity.Cmp(big.NewInt(0)) <= 0 {
		return sdkerrors.Wrap(ErrInsufficientOrderQuantity, "insufficient quantity")
	} else if price == nil || price.Cmp(big.NewInt(0)) <= 0 {
		return sdkerrors.Wrap(ErrInsufficientOrderQuantity, "insufficient price")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateDerivativeOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateDerivativeOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgRegisterSpotMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterSpotMarket) Type() string { return "registerSpotMarket" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterSpotMarket) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if len(msg.Ticker) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no trade pair name specified")
	} else if parts := strings.Split(msg.Ticker, "/"); len(parts) != 2 ||
		len(strings.TrimSpace(parts[0])) == 0 || len(strings.TrimSpace(parts[1])) == 0 {
		return sdkerrors.Wrap(ErrBadField, "pair name must be in format AAA/BBB")
	}

	if !common.IsHexAddress(msg.BaseAsset) {
		return sdkerrors.Wrap(ErrBadField, "no valid BaseAsset specified")
	} else if !common.IsHexAddress(msg.QuoteAsset) {
		return sdkerrors.Wrap(ErrBadField, "no valid QuoteAsset specified")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRegisterSpotMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterSpotMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgRegisterDerivativeMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterDerivativeMarket) Type() string { return "registerDerivativeMarket" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterDerivativeMarket) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	market := msg.Market

	if len(market.Ticker) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no market ticker name specified")
	} else if parts := strings.Split(market.Ticker, "/"); len(parts) != 2 ||
		len(strings.TrimSpace(parts[0])) == 0 || len(strings.TrimSpace(parts[1])) == 0 {
		return sdkerrors.Wrap(ErrBadField, "market ticker must be in format AAA/BBB")
	} else if len(msg.Market.GetOracle()) != ADDRESS_LENGTH {
		return sdkerrors.Wrap(ErrBadField, "oracle address must be of length 42")
	} else if len(msg.Market.GetBaseCurrency()) != ADDRESS_LENGTH {
		return sdkerrors.Wrap(ErrBadField, "base currency address must be of length 42")
	} else if len(msg.Market.GetExchangeAddress()) != ADDRESS_LENGTH {
		return sdkerrors.Wrap(ErrBadField, "exchange address must be of length 42")
	} else if len(msg.Market.GetMarketId()) != BYTES32_LENGTH {
		return sdkerrors.Wrapf(ErrBadField, "marketID must be of length 66 but got length %d", len(msg.Market.GetMarketId()))
	}

	// TODO: (albertchon) proper validation here
	//hash, err := market.Hash()
	//if err != nil {
	//	return sdkerrors.Wrap(ErrMarketInvalid, err.Error())
	//}
	//if hash.String() != market.MarketId {
	//	errMsg := "The MarketID " + market.MarketId + " provided does not match the MarketID " + hash.String() + " computed"
	//	errMsg += "\n for Ticker: " + market.Ticker + "\nOracle: " + market.Oracle + "\nBaseCurrency: " + market.BaseCurrency + "\nNonce: " + market.Nonce
	//	return sdkerrors.Wrap(ErrMarketInvalid, errMsg)
	//}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRegisterDerivativeMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterDerivativeMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgSuspendDerivativeMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSuspendDerivativeMarket) Type() string {
	return "suspendDerivativeMarket"
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSuspendDerivativeMarket) ValidateBasic() error {
	// TODO: albertchon proper length checks here
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	} else if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrBadField, "no derivative market ID specified")
	} else if msg.ExchangeAddress == "" {
		return sdkerrors.Wrap(ErrBadField, "no derivative exchange address specified")

	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgSuspendDerivativeMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSuspendDerivativeMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgResumeDerivativeMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgResumeDerivativeMarket) Type() string {
	return "resumeDerivativeMarket"
}

// ValidateBasic runs stateless checks on the message
func (msg MsgResumeDerivativeMarket) ValidateBasic() error {
	// TODO: albertchon proper validation
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	} else if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrBadField, "no derivative market ID specified")
	} else if msg.ExchangeAddress == "" {
		return sdkerrors.Wrap(ErrBadField, "no derivative market ID specified")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgResumeDerivativeMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgResumeDerivativeMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgSuspendSpotMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSuspendSpotMarket) Type() string { return "suspendSpotMarket" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSuspendSpotMarket) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no trade pair name specified")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgSuspendSpotMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSuspendSpotMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgResumeSpotMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgResumeSpotMarket) Type() string { return "resumeSpotMarket" }

// ValidateBasic runs stateless checks on the message
func (msg MsgResumeSpotMarket) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no trade pair name specified")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgResumeSpotMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgResumeSpotMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgExecuteTECTransaction) Route() string {
	return RouterKey
}

func (msg *MsgExecuteTECTransaction) Type() string {
	return "executeTECTransaction"
}

func (msg *MsgExecuteTECTransaction) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	// TODO: hash this and validate
	transactionHash := common.Hash{}

	if len(msg.TecTransaction.Salt) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no salt specified")
	} else if msg.TecTransaction.SignerAddress == "" {
		return sdkerrors.Wrap(ErrBadField, "no signerAddress address specified")
	} else if msg.TecTransaction.Domain.VerifyingContract == "" {
		return sdkerrors.Wrap(ErrBadField, "no verifyingContract address specified")
	} else if msg.TecTransaction.Domain.ChainId != "888" {
		return sdkerrors.Wrap(ErrBadField, "wrong chainID specified")
	} else if len(msg.TecTransaction.GasPrice) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no gasPrice specified")
	} else if len(msg.TecTransaction.ExpirationTimeSeconds) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no expirationTimeSeconds specified")
	} else if !isValidSignature(msg.TecTransaction.Signature, common.HexToAddress(msg.TecTransaction.SignerAddress), transactionHash) {
		// TODO: fix this later
		return nil
		//return sdkerrors.Wrap(ErrBadField, "invalid transaction signature")
	}
	return nil
}

func (msg *MsgExecuteTECTransaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgExecuteTECTransaction) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgExecuteDerivativeTakeOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgExecuteDerivativeTakeOrder) Type() string { return "executeDerivativeTakeOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgExecuteDerivativeTakeOrder) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order == nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no make order specified")
	}

	order := msg.Order.ToSignedOrder()
	quantity := order.TakerAssetAmount
	margin := order.MakerFee
	orderHash, err := order.ComputeOrderHash()
	takerAddress := common.HexToAddress(msg.Order.TakerAddress)

	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("hash check failed: %v", err))
	} else if !isValidSignature(msg.Order.Signature, takerAddress, orderHash) {
		// TODO @albert @venkatesh: delete return nil once this is ready
		return nil
		//return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid signature")
	} else if quantity == nil || quantity.Cmp(big.NewInt(0)) <= 0 {
		return sdkerrors.Wrap(ErrInsufficientOrderQuantity, "insufficient quantity")
	} else if margin == nil || margin.Cmp(big.NewInt(0)) <= 0 {
		return sdkerrors.Wrap(ErrInsufficientTakerMargin, "insufficient margin")
	} else if !bytes.Equal(order.MakerAddress.Bytes(), common.Address{}.Bytes()) {
		return sdkerrors.Wrap(ErrBadField, "maker address field must be empty")
	} else if bytes.Equal(order.TakerAddress.Bytes(), common.Address{}.Bytes()) {
		return sdkerrors.Wrap(ErrBadField, "taker address field must not be empty")
	}
	return nil
}

func (msg *MsgExecuteDerivativeTakeOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgExecuteDerivativeTakeOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
