package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgSetPrice{}
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgSetPrice) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgSetPrice) Type() string { return "msgSetPrice" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgSetPrice) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgSetPrice) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgSetPrice) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
