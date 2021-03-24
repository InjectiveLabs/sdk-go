package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgCreateInsuranceFund{}
	_ sdk.Msg = &MsgUnderwrite{}
	_ sdk.Msg = &MsgRequestRedemption{}
	_ sdk.Msg = &MsgClaimRedemption{}
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgCreateInsuranceFund) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgCreateInsuranceFund) Type() string { return "createInsuranceFund" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgCreateInsuranceFund) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.OracleId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.OracleId)
	}
	if msg.MarketId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.MarketId)
	}
	if !msg.DepositAmount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.DepositAmount.String())
	}
	if !msg.DepositAmount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.DepositAmount.String())
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgCreateInsuranceFund) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgCreateInsuranceFund) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgUnderwrite) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgUnderwrite) Type() string { return "underwrite" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgUnderwrite) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.MarketId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.MarketId)
	}
	if msg.OracleId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.OracleId)
	}
	if !msg.DepositAmount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.DepositAmount.String())
	}
	if !msg.DepositAmount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.DepositAmount.String())
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgUnderwrite) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgUnderwrite) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRequestRedemption) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRequestRedemption) Type() string { return "requestRedemption" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgRequestRedemption) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.MarketId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.MarketId)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgRequestRedemption) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgRequestRedemption) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgClaimRedemption) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgClaimRedemption) Type() string { return "claimRedemption" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgClaimRedemption) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.MarketId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.MarketId)
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgClaimRedemption) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgClaimRedemption) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
