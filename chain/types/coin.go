package types

import (
	"math/big"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// INJ defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	InjectiveCoin string = "inj"

	// BaseDenomUnit defines the base denomination unit for Photons.
	// 1 photon = 1x10^{BaseDenomUnit} inj
	BaseDenomUnit = 18
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = math.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewInjectiveCoin is a utility function that returns an "inj" coin with the given math.Int amount.
// The function will panic if the provided amount is negative.
func NewInjectiveCoin(amount math.Int) sdk.Coin {
	return sdk.NewCoin(InjectiveCoin, amount)
}

// NewInjectiveCoinInt64 is a utility function that returns an "inj" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewInjectiveCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(InjectiveCoin, amount)
}

func NewDefaultCoinInt64(amount int64) sdk.Coin {
	return NewInjectiveCoinInt64(amount)
}
