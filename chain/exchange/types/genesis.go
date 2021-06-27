package types

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
		Params:                       DefaultParams(),
		IsSpotExchangeEnabled:        true,
		IsDerivativesExchangeEnabled: true,
	}
}
