package types

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func (gs GenesisState) Validate() error {
	// TODO: validate stuff in genesis
	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}
