package v2

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

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

func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	for i, record := range gs.SubaccountRiskProfiles {
		if err := validateRiskProfileRecord(record); err != nil {
			return fmt.Errorf("subaccount_risk_profiles[%d]: %w", i, err)
		}
	}
	for _, record := range gs.SubaccountRiskProfiles {
		if record == nil {
			continue
		}
		profile := record.RiskProfile
		_, err := validateRiskMode(&profile)
		if err != nil {
			continue
		}
	}

	return nil
}

func validateRiskProfileRecord(record *SubaccountRiskProfileRecord) error {
	if record == nil {
		return errors.New("nil record")
	}

	if _, ok := types.IsValidSubaccountID(record.SubaccountId); !ok {
		return fmt.Errorf("invalid subaccount_id %q: must be a 32-byte hex hash", record.SubaccountId)
	}

	mode, err := validateRiskMode(&record.RiskProfile)
	if err != nil {
		return err
	}

	if mode == RiskMode_RISK_MODE_CROSS && types.IsDefaultSubaccountID(common.HexToHash(record.SubaccountId)) {
		return fmt.Errorf("default subaccount %s cannot use cross-margin mode", record.SubaccountId)
	}

	return nil
}

func validateRiskMode(p *SubaccountRiskProfile) (RiskMode, error) {
	mode := p.Mode
	if mode == RiskMode_RISK_MODE_UNSPECIFIED {
		mode = RiskMode_RISK_MODE_ISOLATED
	}
	if mode != RiskMode_RISK_MODE_ISOLATED && mode != RiskMode_RISK_MODE_CROSS {
		return 0, fmt.Errorf("unsupported risk mode %v", p.Mode)
	}

	policy := p.ReservationPolicy
	if policy == ReservationPolicy_RESERVATION_POLICY_UNSPECIFIED {
		policy = ReservationPolicy_RESERVATION_POLICY_FULL_HOLD
	}
	if policy != ReservationPolicy_RESERVATION_POLICY_FULL_HOLD {
		return 0, fmt.Errorf("unsupported reservation policy %v", p.ReservationPolicy)
	}

	if p.CreditLineId != "" {
		return 0, errors.New("credit lines are not supported")
	}

	return mode, nil
}
