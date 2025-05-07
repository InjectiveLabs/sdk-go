package types

import (
	"cosmossdk.io/errors"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:     DefaultParams(),
		TokenPairs: []TokenPair{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	err := gs.Params.Validate()
	if err != nil {
		return err
	}

	seenDenoms := map[string]struct{}{}

	for _, pair := range gs.GetTokenPairs() {
		if _, ok := seenDenoms[pair.GetBankDenom()]; ok {
			return errors.Wrapf(ErrInvalidGenesis, "duplicate bank denom: %s", pair.GetBankDenom())
		}
		seenDenoms[pair.GetBankDenom()] = struct{}{}
	}

	return nil
}
