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

	if err := gs.validateSpotOrderbookMarkets(); err != nil {
		return err
	}

	for i, record := range gs.SubaccountRiskProfiles {
		if err := validateRiskProfileRecord(record); err != nil {
			return fmt.Errorf("subaccount_risk_profiles[%d]: %w", i, err)
		}
	}

	return nil
}

func (gs GenesisState) validateSpotOrderbookMarkets() error {
	spotMarkets, err := gs.validatedSpotMarkets()
	if err != nil {
		return err
	}
	for i, orderbook := range gs.SpotOrderbook {
		marketID := orderbook.MarketId
		if !types.IsHexHash(marketID) {
			return fmt.Errorf("spot_orderbook[%d]: invalid market_id %q", i, marketID)
		}
		market, ok := spotMarkets[marketID]
		if !ok {
			return fmt.Errorf("spot_orderbook[%d]: unknown market_id %s", i, marketID)
		}
		if market.IsActive() {
			if err := validateSpotOrderbookTicks(i, orderbook, market); err != nil {
				return err
			}
		}
	}

	return nil
}

func (gs GenesisState) validatedSpotMarkets() (map[string]*SpotMarket, error) {
	markets := make(map[string]*SpotMarket, len(gs.SpotMarkets))
	for i, market := range gs.SpotMarkets {
		if market == nil {
			continue
		}
		if market.IsActive() {
			if err := ValidateSpotMarketTickSizes(market.MinPriceTickSize, market.MinQuantityTickSize); err != nil {
				return nil, fmt.Errorf("spot_markets[%d]: %w", i, err)
			}
		}
		markets[market.MarketId] = market
	}

	return markets, nil
}

func validateSpotOrderbookTicks(i int, orderbook SpotOrderBook, market *SpotMarket) error {
	for j, order := range orderbook.Orders {
		if order == nil {
			return fmt.Errorf("spot_orderbook[%d].orders[%d]: missing order", i, j)
		}
		if order.OrderInfo.Price.IsNil() || types.BreachesMinimumTickSize(order.OrderInfo.Price, market.MinPriceTickSize) {
			return fmt.Errorf("spot_orderbook[%d].orders[%d]: price does not match market tick size", i, j)
		}
		if order.Fillable.IsNil() || types.BreachesMinimumTickSize(order.Fillable, market.MinQuantityTickSize) {
			return fmt.Errorf("spot_orderbook[%d].orders[%d]: fillable does not match market tick size", i, j)
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
