package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// constants
const (
	TypeMsgCreateDenom      = "create_denom"
	TypeMsgMint             = "tf_mint"
	TypeMsgBurn             = "tf_burn"
	TypeMsgChangeAdmin      = "change_admin"
	TypeMsgSetDenomMetadata = "set_denom_metadata"
	TypeMsgForceTransfer    = "force_transfer"
	TypeMsgUpdateParams     = "update_params"
)

var _ sdk.Msg = &MsgCreateDenom{}
var _ sdk.Msg = &MsgMint{}
var _ sdk.Msg = &MsgBurn{}
var _ sdk.Msg = &MsgSetDenomMetadata{}
var _ sdk.Msg = &MsgChangeAdmin{}
var _ sdk.Msg = &MsgUpdateParams{}

func (m MsgUpdateParams) Route() string { return RouterKey }

func (m MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

func (m MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
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

// NewMsgCreateDenom creates a msg to create a new denom
func NewMsgCreateDenom(sender, subdenom, name, symbol string) *MsgCreateDenom {
	return &MsgCreateDenom{
		Sender:   sender,
		Subdenom: subdenom,
		Name:     name,
		Symbol:   symbol,
	}
}

func (m MsgCreateDenom) Route() string { return RouterKey }
func (m MsgCreateDenom) Type() string  { return TypeMsgCreateDenom }
func (m MsgCreateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err = GetTokenDenom(m.Sender, m.Subdenom)
	if err != nil {
		return errors.Wrap(ErrInvalidDenom, err.Error())
	}

	return nil
}

func (m *MsgCreateDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgCreateDenom) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

// NewMsgMint creates a message to mint tokens
func NewMsgMint(sender string, amount sdk.Coin) *MsgMint {
	return &MsgMint{
		Sender: sender,
		Amount: amount,
	}
}

func (m MsgMint) Route() string { return RouterKey }
func (m MsgMint) Type() string  { return TypeMsgMint }
func (m MsgMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	if !m.Amount.IsValid() || m.Amount.Amount.Equal(sdk.ZeroInt()) {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}

	return nil
}

func (m MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgMint) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

// NewMsgBurn creates a message to burn tokens
func NewMsgBurn(sender string, amount sdk.Coin) *MsgBurn {
	return &MsgBurn{
		Sender: sender,
		Amount: amount,
	}
}

func (m MsgBurn) Route() string { return RouterKey }
func (m MsgBurn) Type() string  { return TypeMsgBurn }
func (m MsgBurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	if !m.Amount.IsValid() || !m.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}

	return nil
}

func (m *MsgBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgBurn) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

// NewMsgChangeAdmin creates a message to change the admin
func NewMsgChangeAdmin(sender, denom, newAdmin string) *MsgChangeAdmin {
	return &MsgChangeAdmin{
		Sender:   sender,
		Denom:    denom,
		NewAdmin: newAdmin,
	}
}

func (m MsgChangeAdmin) Route() string { return RouterKey }
func (m MsgChangeAdmin) Type() string  { return TypeMsgChangeAdmin }
func (m MsgChangeAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(m.NewAdmin)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid address (%s)", err)
	}

	_, _, err = DeconstructDenom(m.Denom)
	if err != nil {
		return err
	}

	return nil
}

func (m *MsgChangeAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgChangeAdmin) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

// NewMsgSetDenomMetadata creates a message to set the denom metadata
func NewMsgSetDenomMetadata(sender string, metadata banktypes.Metadata) *MsgSetDenomMetadata {
	return &MsgSetDenomMetadata{
		Sender:   sender,
		Metadata: metadata,
	}
}

func (m MsgSetDenomMetadata) Route() string { return RouterKey }
func (m MsgSetDenomMetadata) Type() string  { return TypeMsgSetDenomMetadata }
func (m MsgSetDenomMetadata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	err = m.Metadata.Validate()
	if err != nil {
		return err
	}

	_, _, err = DeconstructDenom(m.Metadata.Base)
	if err != nil {
		return err
	}

	return nil
}

func (m *MsgSetDenomMetadata) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgSetDenomMetadata) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

// var _ sdk.Msg = &MsgForceTransfer{}

// // NewMsgForceTransfer creates a transfer funds from one account to another
// func NewMsgForceTransfer(sender string, amount sdk.Coin, fromAddr, toAddr string) *MsgForceTransfer {
// 	return &MsgForceTransfer{
// 		Sender:              sender,
// 		Amount:              amount,
// 		TransferFromAddress: fromAddr,
// 		TransferToAddress:   toAddr,
// 	}
// }

// func (m MsgForceTransfer) Route() string { return RouterKey }
// func (m MsgForceTransfer) Type() string  { return TypeMsgForceTransfer }
// func (m MsgForceTransfer) ValidateBasic() error {
// 	_, err := sdk.AccAddressFromBech32(m.Sender)
// 	if err != nil {
// 		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
// 	}

// 	_, err = sdk.AccAddressFromBech32(m.TransferFromAddress)
// 	if err != nil {
// 		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid address (%s)", err)
// 	}
// 	_, err = sdk.AccAddressFromBech32(m.TransferToAddress)
// 	if err != nil {
// 		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid address (%s)", err)
// 	}

// 	if !m.Amount.IsValid() {
// 		return errors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
// 	}

// 	return nil
// }

// func (m MsgForceTransfer) GetSignBytes() []byte {
// 	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
// }

// func (m MsgForceTransfer) GetSigners() []sdk.AccAddress {
// 	sender, _ := sdk.AccAddressFromBech32(m.Sender)
// 	return []sdk.AccAddress{sender}
// }
