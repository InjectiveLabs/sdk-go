package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = ModuleName

// oracle message types
const (
	TypeMsgRelayPriceFeedPrice   = "relayPriceFeedPrice"
	TypeMsgRelayBandRates        = "relayBandRates"
	TypeMsgRelayCoinbaseMessages = "relayCoinbaseMessages"
)

var (
	_ sdk.Msg = &MsgRelayPriceFeedPrice{}
	_ sdk.Msg = &MsgRelayBandRates{}
	_ sdk.Msg = &MsgRelayCoinbaseMessages{}
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRelayPriceFeedPrice) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRelayPriceFeedPrice) Type() string { return TypeMsgRelayPriceFeedPrice }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgRelayPriceFeedPrice) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
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
