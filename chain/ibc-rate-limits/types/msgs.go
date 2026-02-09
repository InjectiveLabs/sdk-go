package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.LegacyMsg = &MsgCreateRateLimit{}
	_ sdk.LegacyMsg = &MsgUpdateRateLimit{}
	_ sdk.LegacyMsg = &MsgRemoveRateLimit{}
)

func (*MsgCreateRateLimit) Route() string { return RouterKey }

func (*MsgCreateRateLimit) Type() string { return "create_ibc_rate_limit" }

func (m *MsgCreateRateLimit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return err
	}

	return m.RateLimit.ValidateBasic()
}

func (m *MsgCreateRateLimit) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshal(m))
}

func (m *MsgCreateRateLimit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Authority)}
}

func (*MsgUpdateRateLimit) Route() string { return RouterKey }

func (*MsgUpdateRateLimit) Type() string { return "update_ibc_rate_limit" }

func (m *MsgUpdateRateLimit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return err
	}

	return m.NewRateLimit.ValidateBasic()
}

func (m *MsgUpdateRateLimit) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshal(m))
}

func (m *MsgUpdateRateLimit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Authority)}
}

func (*MsgRemoveRateLimit) Route() string { return RouterKey }

func (*MsgRemoveRateLimit) Type() string { return "remove_ibc_rate_limit" }

func (m *MsgRemoveRateLimit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return err
	}

	if m.Token == "" {
		return sdkerrors.Wrap(ErrInvalidRateLimit, "token cannot be empty")
	}

	return nil
}

func (m *MsgRemoveRateLimit) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshal(m))
}

func (m *MsgRemoveRateLimit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Authority)}
}
