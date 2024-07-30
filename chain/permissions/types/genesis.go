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
		Namespaces: []Namespace{},
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

	for _, ns := range gs.GetNamespaces() {
		if _, ok := seenDenoms[ns.GetDenom()]; ok {
			return errors.Wrapf(ErrInvalidGenesis, "duplicate denom: %s", ns.GetDenom())
		}
		seenDenoms[ns.GetDenom()] = struct{}{}
	}

	return nil
}
