package types

import (
	"cosmossdk.io/errors"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = ModuleName

const (
	TypeMsgUpdate                = "msgUpdate"
	TypeMsgActivate              = "msgActivate"
	TypeMsgDeactivate            = "msgDeactivate"
	TypeMsgExecuteContractCompat = "executeContractCompat"
	TypeMsgUpdateParams          = "updateParams"
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgUpdateParams) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	if err := msg.Params.Validate(); err != nil {
		return err
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgUpdateContract) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgUpdateContract) Type() string { return TypeMsgUpdate }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgUpdateContract) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if _, err := sdk.AccAddressFromBech32(msg.ContractAddress); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.ContractAddress)
	}

	if msg.AdminAddress != "" {
		if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.AdminAddress)
		}
	}

	if msg.GasLimit == 0 {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "GasLimit must be > 0")
	}

	if msg.GasPrice == 0 {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "GasPrice must be > 0")
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
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ContractAddress); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.ContractAddress)
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
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ContractAddress); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.ContractAddress)
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

func (msg MsgExecuteContractCompat) Route() string {
	return RouterKey
}

func (msg MsgExecuteContractCompat) Type() string {
	return TypeMsgExecuteContractCompat
}

func (msg MsgExecuteContractCompat) ValidateBasic() error {
	funds := sdk.Coins{}
	if msg.Funds != "0" {
		var err error
		funds, err = sdk.ParseCoinsNormalized(msg.Funds)
		if err != nil {
			return err
		}
	}

	oMsg := &wasmtypes.MsgExecuteContract{
		Sender:   msg.Sender,
		Contract: msg.Contract,
		Msg:      []byte(msg.Msg),
		Funds:    funds,
	}

	if err := oMsg.ValidateBasic(); err != nil {
		return err
	}
	return nil
}

func (msg MsgExecuteContractCompat) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgExecuteContractCompat) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{senderAddr}
}
