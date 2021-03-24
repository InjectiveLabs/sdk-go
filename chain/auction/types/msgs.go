package types

import (
	chaintypes "github.com/InjectiveLabs/injective-core/injective-chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgBid{}
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgBid) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgBid) Type() string { return "msgBid" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgBid) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !msg.BidAmount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.BidAmount.String())
	}

	if msg.BidAmount.Denom != chaintypes.InjectiveCoin {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.BidAmount.Denom)
	}

	if !msg.BidAmount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.BidAmount.String())
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgBid) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgBid) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
