package types

import (
	"math/big"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IsValidChainID returns false if the given chain identifier is incorrectly formatted.
func IsValidChainID(chainID string) bool {
	return chainID == "888"
}


// ParseChainID parses a string chain identifier's epoch to an Ethereum-compatible
// chain-id in *big.Int format. The function returns an error if the chain-id has an invalid format
func ParseChainID(chainID string) (*big.Int, error) {
	chainID = strings.TrimSpace(chainID)

	// verify that the chain-id entered is a base 10 integer
	chainIDInt, ok := new(big.Int).SetString(chainID, 10)
	if !ok {
		return nil, sdkerrors.Wrapf(ErrInvalidChainID, "epoch %s must be base-10 integer format", chainID)
	}

	return chainIDInt, nil
}
