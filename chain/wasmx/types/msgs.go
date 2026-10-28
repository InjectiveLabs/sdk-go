package types

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

const (
	TypeMsgUpdate                = "msgUpdate"
	TypeMsgActivate              = "msgActivate"
	TypeMsgDeactivate            = "msgDeactivate"
	TypeMsgExecuteContractCompat = "executeContractCompat"
	TypeMsgUpdateParams          = "updateParams"
	TypeMsgRegisterContract      = "registerContract"
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgUpdateParams) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgUpdateParams) ValidateBasic() error {
	return ErrModuleRemoved
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
	return ErrModuleRemoved
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
	return ErrModuleRemoved
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
	return ErrModuleRemoved
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
	wasmMsg, err := msg.ToWasmExecuteContract()
	if err != nil {
		return err
	}
	return wasmMsg.ValidateBasic()
}

// ToWasmExecuteContract preserves the legacy string encoding of JSON and coins,
// including "0" and the empty string for calls without funds.
func (msg MsgExecuteContractCompat) ToWasmExecuteContract() (*wasmtypes.MsgExecuteContract, error) {
	funds := sdk.Coins{}
	if msg.Funds != "0" {
		var err error
		funds, err = sdk.ParseCoinsNormalized(msg.Funds)
		if err != nil {
			return nil, err
		}
	}
	return &wasmtypes.MsgExecuteContract{
		Sender: msg.Sender, Contract: msg.Contract, Msg: []byte(msg.Msg), Funds: funds,
	}, nil
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

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRegisterContract) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRegisterContract) Type() string { return TypeMsgRegisterContract }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgRegisterContract) ValidateBasic() error {
	return ErrModuleRemoved
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgRegisterContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgRegisterContract) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
