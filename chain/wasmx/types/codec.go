package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// RegisterLegacyAminoCodec registers the necessary x/wasmx interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateContract{}, "wasmx/MsgUpdateContract", nil)
	cdc.RegisterConcrete(&MsgActivateContract{}, "wasmx/MsgActivateContract", nil)
	cdc.RegisterConcrete(&MsgDeactivateContract{}, "wasmx/MsgDeactivateContract", nil)
	cdc.RegisterConcrete(&MsgExecuteContractCompat{}, "wasmx/MsgExecuteContractCompat", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "wasmx/MsgUpdateParams", nil)

	cdc.RegisterConcrete(&ContractRegistrationRequestProposal{}, "wasmx/ContractRegistrationRequestProposal", nil)
	cdc.RegisterConcrete(&BatchContractRegistrationRequestProposal{}, "wasmx/BatchContractRegistrationRequestProposal", nil)
	cdc.RegisterConcrete(&BatchContractDeregistrationProposal{}, "wasmx/BatchContractDeregistrationProposal", nil)
	cdc.RegisterConcrete(&BatchStoreCodeProposal{}, "wasmx/BatchStoreCodeProposal", nil)

	cdc.RegisterConcrete(&ContractExecutionCompatAuthorization{}, "wasmx/ContractExecutionCompatAuthorization", nil)

	cdc.RegisterConcrete(&Params{}, "wasmx/Params", nil)

}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ContractRegistrationRequestProposal{},
		&BatchContractRegistrationRequestProposal{},
		&BatchContractDeregistrationProposal{},
		&BatchStoreCodeProposal{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateContract{},
		&MsgActivateContract{},
		&MsgActivateContract{},
		&MsgDeactivateContract{},
		&MsgUpdateParams{},
	)

	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&ContractExecutionCompatAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	// ModuleCdc references the global x/wasmx module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/wasmx and
	// defined at the application level.
	ModuleCdc = codec.NewLegacyAmino()
)

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	cryptocodec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()

	RegisterLegacyAminoCodec(authzcdc.Amino)

	// TODO: Check
	// RegisterLegacyAminoCodec(govcdc.Amino)
	// RegisterLegacyAminoCodec(groupcdc.Amino)
}
