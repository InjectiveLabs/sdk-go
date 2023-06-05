package types

import (
	"strings"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// oracle message types
const (
	RouterKey = ModuleName

	TypeMsgRelayPriceFeedPrice   = "relayPriceFeedPrice"
	TypeMsgRelayBandRates        = "relayBandRates"
	TypeMsgRelayCoinbaseMessages = "relayCoinbaseMessages"
	TypeMsgRequestBandIBCRates   = "requestBandIBCRates"
	TypeMsgRelayProviderPrices   = "relayProviderPrices"
	TypeMsgRelayPythPrices       = "relayPythPrices"
	TypeMsgUpdateParams          = "updateParams"
)

var (
	_ sdk.Msg = &MsgRelayPriceFeedPrice{}
	_ sdk.Msg = &MsgRelayBandRates{}
	_ sdk.Msg = &MsgRelayCoinbaseMessages{}
	_ sdk.Msg = &MsgRequestBandIBCRates{}
	_ sdk.Msg = &MsgRelayProviderPrices{}
	_ sdk.Msg = &MsgRelayPythPrices{}
	_ sdk.Msg = &MsgUpdateParams{}
)

func (msg MsgUpdateParams) Route() string { return RouterKey }

func (msg MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

func (msg MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	if err := msg.Params.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(msg))
}

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRelayPriceFeedPrice) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRelayPriceFeedPrice) Type() string { return TypeMsgRelayPriceFeedPrice }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgRelayPriceFeedPrice) ValidateBasic() error {
	if msg.Sender == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return err
	}
	priceCount := len(msg.Price)
	if len(msg.Base) != priceCount {
		return ErrBadPriceFeedBaseCount
	}
	if len(msg.Quote) != priceCount {
		return ErrBadPriceFeedQuoteCount
	}
	for _, price := range msg.Price {
		if !price.IsPositive() {
			return ErrBadPrice
		}
		if price.GT(LargestDecPrice) {
			return ErrPriceTooLarge
		}
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgRelayPriceFeedPrice) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgRelayPriceFeedPrice) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRelayBandRates) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRelayBandRates) Type() string { return TypeMsgRelayBandRates }

// ValidateBasic implements the sdk.Msg interface for MsgRelay.
func (msg MsgRelayBandRates) ValidateBasic() error {
	if msg.Relayer == "" {
		return ErrEmptyRelayerAddr
	}

	// check that the sizes of symbols,rates,resolveTimes,requestIDs are equal
	symbolsCount := len(msg.Symbols)
	if len(msg.Rates) != symbolsCount {
		return ErrBadRatesCount
	}
	if len(msg.ResolveTimes) != symbolsCount {
		return ErrBadResolveTimesCount
	}
	if len(msg.RequestIDs) != symbolsCount {
		return ErrBadRequestIDsCount
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgRelayBandRates) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgRelayBandRates) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Relayer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRelayCoinbaseMessages) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRelayCoinbaseMessages) Type() string { return TypeMsgRelayCoinbaseMessages }

// ValidateBasic implements the sdk.Msg interface for MsgRelay.
func (msg MsgRelayCoinbaseMessages) ValidateBasic() error {
	if msg.Sender == "" {
		return ErrEmptyRelayerAddr
	}

	// check that the sizes of messages and signatures are equal
	if len(msg.Signatures) != len(msg.Messages) || len(msg.Messages) == 0 {
		return ErrBadMessagesCount
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgRelayCoinbaseMessages) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgRelayCoinbaseMessages) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// NewMsgRequestBandIBCRates creates a new MsgRequestBandIBCRates instance.
func NewMsgRequestBandIBCRates(
	sender sdk.AccAddress,
	requestID uint64,
) *MsgRequestBandIBCRates {
	return &MsgRequestBandIBCRates{
		Sender:    sender.String(),
		RequestId: requestID,
	}
}

// Route implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestBandIBCRates) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestBandIBCRates) Type() string { return TypeMsgRequestBandIBCRates }

// ValidateBasic implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestBandIBCRates) ValidateBasic() error {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	if sender.Empty() {
		return errors.Wrapf(ErrInvalidBandIBCRequest, "MsgRequestBandIBCRates: Sender address must not be empty.")
	}

	if msg.RequestId == 0 {
		return errors.Wrapf(ErrInvalidBandIBCRequest, "MsgRequestBandIBCRates: requestID should be greater than zero")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestBandIBCRates) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestBandIBCRates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRelayProviderPrices) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRelayProviderPrices) Type() string { return TypeMsgRelayProviderPrices }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgRelayProviderPrices) ValidateBasic() error {
	if msg.Sender == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return err
	}

	if msg.Provider == "" {
		return ErrEmptyProvider
	}

	if len(msg.Symbols) != len(msg.Prices) || len(msg.Prices) == 0 {
		return ErrBadRatesCount
	}

	for _, symbol := range msg.Symbols {
		if strings.Contains(symbol, providerDelimiter) {
			return ErrInvalidSymbol
		}
	}

	for _, price := range msg.Prices {
		// zero prices are allowed for provider oracles
		if price.IsNegative() {
			return ErrBadPrice
		}
		if price.GT(LargestDecPrice) {
			return ErrPriceTooLarge
		}
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgRelayProviderPrices) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgRelayProviderPrices) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRelayPythPrices) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRelayPythPrices) Type() string { return TypeMsgRelayPythPrices }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgRelayPythPrices) ValidateBasic() error {
	if msg.Sender == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return err
	}

	if len(msg.PriceAttestations) == 0 {
		return ErrEmptyPriceAttestations
	}

	// the ValidateBasic method intentiontally does not check the validity of the price attestations since
	// we don't want to prevent attesting valid prices just because other price attestations are invalid
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgRelayPythPrices) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgRelayPythPrices) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
