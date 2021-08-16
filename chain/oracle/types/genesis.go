package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func (gs GenesisState) Validate() error {
	// TODO: validate stuff in genesis
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:        DefaultParams(),
	}
}
