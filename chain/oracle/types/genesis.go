package types

import "errors"

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func (gs GenesisState) Validate() error {
	if len(gs.HistoricalPriceRecords) != 0 {
		return errors.New("historical oracle prices cannot be imported after removal")
	}
	if len(gs.CoinbasePriceStates) != 0 {
		return errors.New("coinbase oracle genesis state is no longer supported")
	}
	if len(gs.PriceFeedPriceStates) != 0 {
		return errors.New("pricefeed oracle genesis state is no longer supported")
	}
	if len(gs.StorkPriceStates) != 0 || len(gs.StorkPublishers) != 0 {
		return errors.New("stork oracle genesis state is no longer supported")
	}
	// TODO: validate stuff in genesis
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:        DefaultParams(),
		BandIbcParams: DefaultBandIBCParams(),
	}
}
