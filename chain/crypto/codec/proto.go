package codec

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"

	"github.com/InjectiveLabs/injective-exchange/client/chain/crypto/ethsecp256k1"
)

// RegisterInterfaces registers the ethsecp256k1 implementations of tendermint crypto types.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*tmcrypto.PubKey)(nil), &ethsecp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})

	registry.RegisterImplementations((*tmcrypto.PrivKey)(nil), &ethsecp256k1.PrivKey{})
	registry.RegisterImplementations((*cryptotypes.PrivKey)(nil), &ethsecp256k1.PrivKey{})
}
