package v2

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:                       DefaultParams(),
		IsSpotExchangeEnabled:        true,
		IsDerivativesExchangeEnabled: true,
	}
}
