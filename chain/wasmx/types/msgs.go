package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
)

const RouterKey = ModuleName

const (
	TypeMsgUpdate     = "msgUpdate"
	TypeMsgActivate   = "msgActivate"
	TypeMsgDeactivate = "msgDeactivate"
	TypeMsgTransfer   = "transfer"
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgUpdateContract) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgUpdateContract) Type() string { return TypeMsgUpdate }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgUpdateContract) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if _, err := sdk.AccAddressFromBech32(msg.ContractAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ContractAddress)
	}

	if msg.AdminAddress != "" {
		if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.AdminAddress)
		}
	}

	if msg.GasLimit == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "GasLimit must be > 0")
	}

	if msg.GasPrice == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "GasPrice must be > 0")
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgUpdateContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgUpdateContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgActivateContract) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgActivateContract) Type() string { return TypeMsgActivate }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgActivateContract) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ContractAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ContractAddress)
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgActivateContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgActivateContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgDeactivateContract) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgDeactivateContract) Type() string { return TypeMsgDeactivate }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgDeactivateContract) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ContractAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ContractAddress)
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgDeactivateContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgDeactivateContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements sdk.Msg
func (MsgTransfer) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgTransfer) Type() string {
	return TypeMsgTransfer
}

// ValidateBasic performs a basic check of the MsgTransfer fields.
// NOTE: timeout height or timestamp values can be 0 to disable the timeout.
// NOTE: The recipient addresses format is not validated as the format defined by
// the chain is not known to IBC.
func (msg MsgTransfer) ValidateBasic() error {
	if err := host.PortIdentifierValidator(msg.SourcePort); err != nil {
		return sdkerrors.Wrap(err, "invalid source port ID")
	}
	if err := host.ChannelIdentifierValidator(msg.SourceChannel); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}
	if !msg.Token.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Token.String())
	}
	if !msg.Token.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, msg.Token.String())
	}
	// NOTE: sender format must be validated as it is required by the GetSigners function.
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	if strings.TrimSpace(msg.Receiver) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	return ibctransfertypes.ValidateIBCDenom(msg.Token.Denom)
}

// GetSignBytes implements sdk.Msg.
func (msg MsgTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgTransfer) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}
