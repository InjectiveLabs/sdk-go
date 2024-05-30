package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// RegisterLegacyAminoCodec registers the necessary x/oracle interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRelayPriceFeedPrice{}, "oracle/MsgRelayPriceFeedPrice", nil)
	cdc.RegisterConcrete(&MsgRelayBandRates{}, "oracle/MsgRelayBandRates", nil)
	cdc.RegisterConcrete(&MsgRelayCoinbaseMessages{}, "oracle/MsgRelayCoinbaseMessages", nil)
	cdc.RegisterConcrete(&MsgRequestBandIBCRates{}, "oracle/MsgRequestBandIBCRates", nil)
	cdc.RegisterConcrete(&MsgRelayProviderPrices{}, "oracle/MsgRelayProviderPrices", nil)
	cdc.RegisterConcrete(&MsgRelayPythPrices{}, "oracle/MsgRelayPythPrices", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "oracle/MsgUpdateParams", nil)

	cdc.RegisterConcrete(&GrantBandOraclePrivilegeProposal{}, "oracle/GrantBandOraclePrivilegeProposal", nil)
	cdc.RegisterConcrete(&RevokeBandOraclePrivilegeProposal{}, "oracle/RevokeBandOraclePrivilegeProposal", nil)
	cdc.RegisterConcrete(&GrantPriceFeederPrivilegeProposal{}, "oracle/GrantPriceFeederPrivilegeProposal", nil)
	cdc.RegisterConcrete(&RevokePriceFeederPrivilegeProposal{}, "oracle/RevokePriceFeederPrivilegeProposal", nil)
	cdc.RegisterConcrete(&AuthorizeBandOracleRequestProposal{}, "oracle/AuthorizeBandOracleRequestProposal", nil)
	cdc.RegisterConcrete(&UpdateBandOracleRequestProposal{}, "oracle/UpdateBandOracleRequestProposal", nil)
	cdc.RegisterConcrete(&EnableBandIBCProposal{}, "oracle/EnableBandIBCProposal", nil)
	cdc.RegisterConcrete(&GrantProviderPrivilegeProposal{}, "oracle/GrantProviderPrivilegeProposal", nil)
	cdc.RegisterConcrete(&RevokeProviderPrivilegeProposal{}, "oracle/RevokeProviderPrivilegeProposal", nil)
	cdc.RegisterConcrete(&Params{}, "oracle/Params", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRelayPriceFeedPrice{},
		&MsgRelayBandRates{},
		&MsgRelayCoinbaseMessages{},
		&MsgRequestBandIBCRates{},
		&MsgRelayProviderPrices{},
		&MsgRelayPythPrices{},
		&MsgUpdateParams{},
	)

	registry.RegisterImplementations((*govtypes.Content)(nil),
		&GrantBandOraclePrivilegeProposal{},
		&RevokeBandOraclePrivilegeProposal{},
		&GrantPriceFeederPrivilegeProposal{},
		&RevokePriceFeederPrivilegeProposal{},
		&AuthorizeBandOracleRequestProposal{},
		&UpdateBandOracleRequestProposal{},
		&EnableBandIBCProposal{},
		&GrantProviderPrivilegeProposal{},
		&RevokeProviderPrivilegeProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/oracle module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/oracle and
	// defined at the application level.
	// ModuleCdc = codec.NewAminoCodec(amino)

	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()

	RegisterLegacyAminoCodec(authzcdc.Amino)
	// TODO: check
	// RegisterLegacyAminoCodec(govcdc.Amino)
	// RegisterLegacyAminoCodec(groupcdc.Amino)
}
