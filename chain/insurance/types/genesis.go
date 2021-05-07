package types

import (
	"errors"
)

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func (gs GenesisState) Validate() error {
	if gs.NextRedemptionScheduleId == 0 {
		return errors.New("NextRedemptionScheduleId should NOT be zero")
	}
	if gs.NextShareDenomId == 0 {
		return errors.New("NextShareDenomId should NOT be zero")
	}

	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:                   DefaultParams(),
		NextShareDenomId:         1,
		NextRedemptionScheduleId: 1,
		RedemptionSchedule:       []RedemptionSchedule{},
		InsuranceFunds:           []InsuranceFund{},
	}
}
