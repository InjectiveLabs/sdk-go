package codec

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/InjectiveLabs/sdk-go/chain/crypto/ethsecp256k1"
)

var (
	_ cryptotypes.PubKey  = &ethsecp256k1.PubKey{}
	_ cryptotypes.PrivKey = &ethsecp256k1.PrivKey{}
)

// RegisterInterfaces registers the ethsecp256k1 implementations of tendermint crypto types.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PrivKey)(nil), &ethsecp256k1.PrivKey{})
}
