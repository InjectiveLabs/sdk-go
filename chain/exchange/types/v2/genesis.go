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

	if err := gs.validateDerivativeMarketSettlementScheduled(); err != nil {
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

func (gs GenesisState) validateSpotOrderbookMarkets() error {
	spotMarkets := gs.spotMarketIDs()
	for i, orderbook := range gs.SpotOrderbook {
		marketID := orderbook.MarketId
		if !types.IsHexHash(marketID) {
			return fmt.Errorf("spot_orderbook[%d]: invalid market_id %q", i, marketID)
		}
		if _, ok := spotMarkets[marketID]; !ok {
			return fmt.Errorf("spot_orderbook[%d]: unknown market_id %s", i, marketID)
		}
	}

	return nil
}

func (gs GenesisState) spotMarketIDs() map[string]struct{} {
	markets := make(map[string]struct{}, len(gs.SpotMarkets))
	for _, market := range gs.SpotMarkets {
		if market == nil {
			continue
		}
		markets[market.MarketId] = struct{}{}
	}

	return markets
}

func (gs GenesisState) validateDerivativeMarketSettlementScheduled() error {
	expiryInfos, err := gs.validatedExpiryInfoMarketIDs()
	if err != nil {
		return err
	}

	derivativeMarkets := gs.derivativeMarketByID()
	binaryMarkets := gs.binaryMarketIDs()
	seen := make(map[string]struct{}, len(gs.DerivativeMarketSettlementScheduled))
	for i, marker := range gs.DerivativeMarketSettlementScheduled {
		if err := validateScheduledSettlementMarkerID(i, marker, seen); err != nil {
			return err
		}
		if err := validateScheduledSettlementMarkerMarket(i, marker, derivativeMarkets, binaryMarkets, expiryInfos); err != nil {
			return err
		}
	}

	return nil
}

func (gs GenesisState) derivativeMarketByID() map[string]*DerivativeMarket {
	markets := make(map[string]*DerivativeMarket, len(gs.DerivativeMarkets))
	for _, market := range gs.DerivativeMarkets {
		if market == nil {
			continue
		}
		markets[market.MarketId] = market
	}

	return markets
}

func (gs GenesisState) binaryMarketIDs() map[string]struct{} {
	markets := make(map[string]struct{}, len(gs.BinaryOptionsMarkets))
	for _, market := range gs.BinaryOptionsMarkets {
		if market == nil {
			continue
		}
		markets[market.MarketId] = struct{}{}
	}

	return markets
}

func (gs GenesisState) validatedExpiryInfoMarketIDs() (map[string]struct{}, error) {
	expiryInfos := make(map[string]struct{}, len(gs.ExpiryFuturesMarketInfoState))
	for i, info := range gs.ExpiryFuturesMarketInfoState {
		if err := validateExpiryInfoState(i, info); err != nil {
			return nil, err
		}
		expiryInfos[info.MarketId] = struct{}{}
	}

	return expiryInfos, nil
}

func validateExpiryInfoState(i int, info ExpiryFuturesMarketInfoState) error {
	if !types.IsHexHash(info.MarketId) {
		return fmt.Errorf("expiry_futures_market_info_state[%d]: invalid market_id %q", i, info.MarketId)
	}
	if info.MarketInfo == nil {
		return fmt.Errorf("expiry_futures_market_info_state[%d]: missing market_info", i)
	}
	if info.MarketInfo.MarketId != "" && info.MarketInfo.MarketId != info.MarketId {
		return fmt.Errorf("expiry_futures_market_info_state[%d]: market_id mismatch %q != %q",
			i, info.MarketInfo.MarketId, info.MarketId)
	}

	return nil
}

func validateScheduledSettlementMarkerID(
	i int,
	marker DerivativeMarketSettlementInfo,
	seen map[string]struct{},
) error {
	marketID := marker.MarketId
	if !types.IsHexHash(marketID) {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: invalid market_id %q", i, marketID)
	}
	if _, ok := seen[marketID]; ok {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: duplicate market_id %s", i, marketID)
	}
	seen[marketID] = struct{}{}

	return nil
}

func validateScheduledSettlementMarkerMarket(
	i int,
	marker DerivativeMarketSettlementInfo,
	derivativeMarkets map[string]*DerivativeMarket,
	binaryMarkets map[string]struct{},
	expiryInfos map[string]struct{},
) error {
	marketID := marker.MarketId
	if derivativeMarket, ok := derivativeMarkets[marketID]; ok {
		return validateDerivativeSettlementMarker(i, marker, derivativeMarket, expiryInfos)
	}
	if _, ok := binaryMarkets[marketID]; ok {
		return validateBinarySettlementMarker(i, marker)
	}

	return fmt.Errorf("derivative_market_settlement_scheduled[%d]: unknown market_id %s", i, marketID)
}

func validateDerivativeSettlementMarker(
	i int,
	marker DerivativeMarketSettlementInfo,
	market *DerivativeMarket,
	expiryInfos map[string]struct{},
) error {
	marketID := marker.MarketId
	if market.IsTimeExpiry() {
		if _, hasInfo := expiryInfos[marketID]; !hasInfo {
			return fmt.Errorf("derivative_market_settlement_scheduled[%d]: expiry market %s missing expiry info",
				i, marketID)
		}
	}
	if marker.SettlementPrice.IsNil() {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: derivative market %s missing settlement price",
			i, marketID)
	}
	if marker.IsForcedSettlement && marker.SettlementPrice.IsNegative() {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: forced derivative market %s has negative settlement price %s",
			i, marketID, marker.SettlementPrice.String())
	}
	if !marker.IsForcedSettlement && marker.SettlementPrice.IsNegative() {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: derivative market %s has negative settlement price %s",
			i, marketID, marker.SettlementPrice.String())
	}

	return nil
}

func validateBinarySettlementMarker(i int, marker DerivativeMarketSettlementInfo) error {
	marketID := marker.MarketId
	if marker.SettlementPrice.IsNil() {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: binary options market %s missing settlement price",
			i, marketID)
	}
	if marker.IsForcedSettlement && !isValidForcedBinarySettlementMarkerPrice(marker) {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: forced binary options market %s has invalid settlement price %s",
			i, marketID, marker.SettlementPrice.String())
	}
	if !marker.IsForcedSettlement && !isValidNonForcedBinarySettlementMarkerPrice(marker) {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: binary options market %s has invalid settlement price %s",
			i, marketID, marker.SettlementPrice.String())
	}

	return nil
}

func isValidForcedBinarySettlementMarkerPrice(marker DerivativeMarketSettlementInfo) bool {
	return marker.SettlementPrice.IsPositive() &&
		marker.SettlementPrice.LTE(types.MaxBinaryOptionsOrderPrice)
}

func isValidNonForcedBinarySettlementMarkerPrice(marker DerivativeMarketSettlementInfo) bool {
	return marker.SettlementPrice.Equal(BinaryOptionsMarketRefundFlagPrice) ||
		(!marker.SettlementPrice.IsNegative() &&
			marker.SettlementPrice.LTE(types.MaxBinaryOptionsOrderPrice))
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
