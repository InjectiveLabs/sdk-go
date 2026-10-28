package types

import (
	"context"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
)

var (
	_ authztypes.Authorization         = &ContractExecutionCompatAuthorization{}
	_ cdctypes.UnpackInterfacesMessage = &ContractExecutionCompatAuthorization{}
)

// NewContractExecutionCompatAuthorization constructor
func NewContractExecutionCompatAuthorization(grants ...wasmtypes.ContractGrant) *ContractExecutionCompatAuthorization {
	return &ContractExecutionCompatAuthorization{
		Grants: grants,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (ContractExecutionCompatAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgExecuteContractCompat{})
}

// NewAuthz factory method to create an Authorization with updated grants
func (ContractExecutionCompatAuthorization) NewAuthz(g []wasmtypes.ContractGrant) authztypes.Authorization {
	return NewContractExecutionCompatAuthorization(g...)
}

// Accept implements Authorization.Accept.
func (a *ContractExecutionCompatAuthorization) Accept(goCtx context.Context, msg sdk.Msg) (authztypes.AcceptResponse, error) {
	compatMsg, ok := msg.(*MsgExecuteContractCompat)
	if !ok {
		return authztypes.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrap("msg is not a MsgExecuteContractCompat")
	}
	wasmMsg, err := compatMsg.ToWasmExecuteContract()
	if err != nil {
		return authztypes.AcceptResponse{}, err
	}
	response, err := wasmtypes.NewContractExecutionAuthorization(a.Grants...).Accept(goCtx, wasmMsg)
	if response.Updated != nil {
		response.Updated = NewContractExecutionCompatAuthorization(response.Updated.(*wasmtypes.ContractExecutionAuthorization).Grants...)
	}
	return response, err
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a ContractExecutionCompatAuthorization) ValidateBasic() error {
	return wasmtypes.NewContractExecutionAuthorization(a.Grants...).ValidateBasic()
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (a ContractExecutionCompatAuthorization) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	for _, g := range a.Grants {
		if err := g.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}
	return nil
}
