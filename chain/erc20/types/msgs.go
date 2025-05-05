package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// constants
const (
	// RouterKey is the message route for slashing
	routerKey = ModuleName

	TypeMsgUpdateParams    = "update_params"
	TypeMsgCreateTokenPair = "create_token_pair"
	TypeMsgDeleteTokenPair = "delete_token_pair"
)

var (
	_ sdk.Msg = &MsgUpdateParams{}
	_ sdk.Msg = &MsgCreateTokenPair{}
	_ sdk.Msg = &MsgDeleteTokenPair{}
)

func (m MsgUpdateParams) Route() string { return routerKey }

func (m MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

func (m MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return err
	}

	if err := m.Params.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgCreateTokenPair) Route() string { return routerKey }

func (m MsgCreateTokenPair) Type() string { return TypeMsgCreateTokenPair }

func (msg MsgCreateTokenPair) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return err
	}

	if err := msg.TokenPair.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *MsgCreateTokenPair) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgCreateTokenPair) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m MsgDeleteTokenPair) Route() string { return routerKey }

func (m MsgDeleteTokenPair) Type() string { return TypeMsgDeleteTokenPair }

func (m MsgDeleteTokenPair) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return err
	}

	return nil
}

func (m *MsgDeleteTokenPair) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgDeleteTokenPair) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}
