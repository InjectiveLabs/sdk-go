package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgInitHub{}
)

// NewMsgInitHub returns init hub msg instance
func NewMsgInitHub(hub string, proposer sdk.AccAddress) *MsgInitHub {
	return &MsgInitHub{
		HubAddress: hub,
		Proposer:   proposer,
	}
}

// Route should return the name of the module
func (msg MsgInitHub) Route() string { return RouterKey }

// Type should return the action
func (msg MsgInitHub) Type() string { return "inithub" }

// ValidateBasic runs stateless checks on the message
func (msg MsgInitHub) ValidateBasic() error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgInitHub) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgInitHub) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Proposer}
}
