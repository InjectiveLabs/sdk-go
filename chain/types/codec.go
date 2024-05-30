package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// RegisterInterfaces registers the tendermint concrete client-related
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("injective.types.v1beta1.EthAccount", (*types.AccountI)(nil))

	registry.RegisterImplementations(
		(*types.AccountI)(nil),
		&EthAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&EthAccount{},
	)

	registry.RegisterInterface("injective.types.v1beta1.ExtensionOptionsWeb3Tx", (*tx.TxExtensionOptionI)(nil))
	registry.RegisterImplementations(
		(*tx.TxExtensionOptionI)(nil),
		&ExtensionOptionsWeb3Tx{},
	)
}
