package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcdc "github.com/cosmos/cosmos-sdk/x/gov/codec"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	groupcdc "github.com/cosmos/cosmos-sdk/x/group/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)

	RegisterLegacyAminoCodec(authzcdc.Amino)
	RegisterLegacyAminoCodec(govcdc.Amino)
	RegisterLegacyAminoCodec(groupcdc.Amino)
}

// RegisterInterfaces registers the interfaces for the proto stuff
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgValsetConfirm{},
		&MsgSendToEth{},
		&MsgRequestBatch{},
		&MsgConfirmBatch{},
		&MsgDepositClaim{},
		&MsgWithdrawClaim{},
		&MsgERC20DeployedClaim{},
		&MsgSetOrchestratorAddresses{},
		&MsgValsetUpdatedClaim{},
		&MsgCancelSendToEth{},
		&MsgSubmitBadSignatureEvidence{},
		&MsgUpdateParams{},
	)

	registry.RegisterImplementations((*govtypes.Content)(nil),
		&BlacklistEthereumAddressesProposal{},
		&RevokeEthereumBlacklistProposal{},
	)

	registry.RegisterInterface(
		"peggy.v1beta1.EthereumClaim",
		(*EthereumClaim)(nil),
		&MsgDepositClaim{},
		&MsgWithdrawClaim{},
		&MsgERC20DeployedClaim{},
		&MsgValsetUpdatedClaim{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers concrete types on the Amino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetOrchestratorAddresses{}, "peggy/MsgSetOrchestratorAddresses", nil)
	cdc.RegisterConcrete(&MsgValsetConfirm{}, "peggy/MsgValsetConfirm", nil)
	cdc.RegisterConcrete(&MsgSendToEth{}, "peggy/MsgSendToEth", nil)
	cdc.RegisterConcrete(&MsgCancelSendToEth{}, "peggy/MsgCancelSendToEth", nil)
	cdc.RegisterConcrete(&MsgRequestBatch{}, "peggy/MsgRequestBatch", nil)
	cdc.RegisterConcrete(&MsgConfirmBatch{}, "peggy/MsgConfirmBatch", nil)
	cdc.RegisterConcrete(&Valset{}, "peggy/Valset", nil)
	cdc.RegisterConcrete(&MsgDepositClaim{}, "peggy/MsgDepositClaim", nil)
	cdc.RegisterConcrete(&MsgWithdrawClaim{}, "peggy/MsgWithdrawClaim", nil)
	cdc.RegisterConcrete(&MsgERC20DeployedClaim{}, "peggy/MsgERC20DeployedClaim", nil)
	cdc.RegisterConcrete(&MsgValsetUpdatedClaim{}, "peggy/MsgValsetUpdatedClaim", nil)
	cdc.RegisterConcrete(&OutgoingTxBatch{}, "peggy/OutgoingTxBatch", nil)
	cdc.RegisterConcrete(&OutgoingTransferTx{}, "peggy/OutgoingTransferTx", nil)
	cdc.RegisterConcrete(&ERC20Token{}, "peggy/ERC20Token", nil)
	cdc.RegisterConcrete(&IDSet{}, "peggy/IDSet", nil)
	cdc.RegisterConcrete(&Attestation{}, "peggy/Attestation", nil)
	cdc.RegisterConcrete(&MsgSubmitBadSignatureEvidence{}, "peggy/MsgSubmitBadSignatureEvidence", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "peggy/MsgUpdateParams", nil)

	cdc.RegisterConcrete(&BlacklistEthereumAddressesProposal{}, "peggy/BlacklistEthereumAddressesProposal", nil)
	cdc.RegisterConcrete(&RevokeEthereumBlacklistProposal{}, "peggy/RevokeEthereumBlacklistProposal", nil)
}
