package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type (
	ExtensionOptionsWeb3TxI interface{}
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

	registry.RegisterInterface("injective.types.v1beta1.ExtensionOptionsWeb3Tx", (*ExtensionOptionsWeb3TxI)(nil))
	registry.RegisterImplementations(
		(*ExtensionOptionsWeb3TxI)(nil),
		&ExtensionOptionsWeb3Tx{},
	)

	registry.RegisterInterface("injective.types.v1beta1.ExtensionOptionI", (*tx.TxExtensionOptionI)(nil))
	registry.RegisterImplementations(
		(*tx.TxExtensionOptionI)(nil),
		&ExtensionOptionsWeb3Tx{},
	)
}
