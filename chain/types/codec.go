package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// RegisterInterfaces registers the tendermint concrete client-related
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("injective.types.v1beta1.EthAccount", (*authtypes.AccountI)(nil))

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&EthAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&EthAccount{},
	)

	registry.RegisterInterface("injective.types.v1beta1.ExtensionOptionsWeb3Tx", (*tx.ExtensionOptionI)(nil))
	registry.RegisterImplementations(
		(*tx.ExtensionOptionI)(nil),
		&ExtensionOptionsWeb3Tx{},
	)
}
