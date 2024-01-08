package types

import (
	"fmt"
	"strings"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	ModuleDenomPrefix = "factory"
	// See the TokenFactory readme for a derivation of these.
	// TL;DR, MaxSubdenomLength + MaxHrpLength = 60 comes from SDK max denom length = 128
	// and the structure of tokenfactory denoms.
	MaxSubdenomLength = 44
	MaxHrpLength      = 16
	// MaxCreatorLength = 59 + MaxHrpLength
	MaxCreatorLength  = 59 + MaxHrpLength
	MinSubdenomLength = 1
)

// GetTokenDenom constructs a denom string for tokens created by tokenfactory
// based on an input creator address and a subdenom
// The denom constructed is factory/{creator}/{subdenom}
func GetTokenDenom(creator, subdenom string) (string, error) {
	if len(subdenom) > MaxSubdenomLength {
		return "", ErrSubdenomTooLong
	}
	if len(subdenom) < MinSubdenomLength {
		return "", ErrSubdenomTooShort
	}

	strParts := strings.Split(subdenom, "/")

	for idx := range strParts {
		if len(strParts[idx]) < MinSubdenomLength {
			return "", ErrSubdenomNestedTooShort
		}
	}

	if len(creator) > MaxCreatorLength {
		return "", ErrCreatorTooLong
	}
	if strings.Contains(creator, "/") || len(creator) < 1 {
		return "", ErrInvalidCreator
	}

	// ensure creator address is in lowercase Bech32 form
	creatorAddr, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return "", err
	}

	denom := strings.Join([]string{ModuleDenomPrefix, creatorAddr.String(), subdenom}, "/")
	return denom, sdk.ValidateDenom(denom)
}

// DeconstructDenom takes a token denom string and verifies that it is a valid
// denom of the tokenfactory module, and is of the form `factory/{creator}/{subdenom}`
// If valid, it returns the creator address and subdenom
func DeconstructDenom(denom string) (creator, subdenom string, err error) {
	err = sdk.ValidateDenom(denom)
	if err != nil {
		return "", "", err
	}

	strParts := strings.Split(denom, "/")
	if len(strParts) < 3 {
		return "", "", errors.Wrapf(ErrInvalidDenom, "not enough parts of denom %s", denom)
	}

	if strParts[0] != ModuleDenomPrefix {
		return "", "", errors.Wrapf(ErrInvalidDenom, "denom prefix is incorrect. Is: %s.  Should be: %s", strParts[0], ModuleDenomPrefix)
	}

	creator = strParts[1]
	creatorAddr, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return "", "", errors.Wrapf(ErrInvalidDenom, "Invalid creator address (%s)", err)
	}

	if len(strParts[2]) < MinSubdenomLength {
		return "", "", errors.Wrapf(ErrSubdenomTooShort, "subdenom too short. Has length of: %d.  Should have length at least of: %d", len(strParts[2]), MinSubdenomLength)
	}

	if len(strParts) > 3 {
		for i := 2; i <= len(strParts)-1; i++ {
			if len(strParts[i]) < MinSubdenomLength {
				return "", "", errors.Wrapf(ErrSubdenomNestedTooShort, "nested subdenom too short. Has length of: %d.  Should have length at least of: %d", len(strParts[i]), MinSubdenomLength)
			}
		}
	}

	// Handle the case where a denom has a slash in its subdenom. For example,
	// when we did the split, we'd turn factory/accaddr/atomderivative/sikka into ["factory", "accaddr", "atomderivative", "sikka"]
	// So we have to join [2:] with a "/" as the delimiter to get back the correct subdenom which should be "atomderivative/sikka"
	subdenom = strings.Join(strParts[2:], "/")

	return creatorAddr.String(), subdenom, nil
}

// NewTokenFactoryDenomMintCoinsRestriction creates and returns a MintingRestrictionFn that only allows minting of
// valid tokenfactory denoms
func NewTokenFactoryDenomMintCoinsRestriction() banktypes.MintingRestrictionFn {
	return func(ctx sdk.Context, coinsToMint sdk.Coins) error {
		for _, coin := range coinsToMint {
			_, _, err := DeconstructDenom(coin.Denom)
			if err != nil {
				return fmt.Errorf("does not have permission to mint %s", coin.Denom)
			}
		}
		return nil
	}
}
