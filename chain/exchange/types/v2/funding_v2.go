package v2

import (
	"fmt"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

const (
	// MaxFundingV2Markets bounds the number of perpetual markets using funding v2 at once.
	MaxFundingV2Markets = 64
	// MaxFundingImpactOrdersPerSide bounds the resting orders visited per side per impact sample,
	// skipped orders included.
	MaxFundingImpactOrdersPerSide = 128
	// MaxFundingSampleAgeSeconds is how long a funding v2 observation keeps contributing.
	MaxFundingSampleAgeSeconds = 60
)

// FundingImpactTakerSubaccountID is the reserved virtual taker of the funding v2 impact price
// simulation. It lives on the zero address, so no one can sign for it: it can receive deposits by
// transfer, but it cannot place orders or hold positions.
var FundingImpactTakerSubaccountID = common.HexToHash("0x0000000000000000000000000000000000000000ffffffffffffffffffffffff")

// GetFundingImpactNotional returns the market's funding v2 impact notional, treating an unset value
// as zero (legacy funding).
func (m *PerpetualMarketInfo) GetFundingImpactNotional() math.LegacyDec {
	if m == nil || m.FundingImpactNotional.IsNil() {
		return math.LegacyZeroDec()
	}
	return m.FundingImpactNotional
}

// IsFundingV2 reports whether the perpetual market uses funding v2.
func (m *PerpetualMarketInfo) IsFundingV2() bool {
	return m.GetFundingImpactNotional().IsPositive()
}

// GetLastPremium returns the last funding v2 observation, treating an unset value as zero.
func (m *PerpetualMarketFunding) GetLastPremium() math.LegacyDec {
	if m == nil || m.LastPremium.IsNil() {
		return math.LegacyZeroDec()
	}
	return m.LastPremium
}

// ValidateFundingImpactNotional checks a funding v2 impact notional.
func ValidateFundingImpactNotional(notional math.LegacyDec) error {
	if notional.IsNil() || notional.IsNegative() || !notional.IsInValidRange() {
		return errors.Wrapf(types.ErrInvalidFundingImpactNotional, "funding impact notional %s must be non-negative", notional)
	}
	return nil
}

// ValidateSyntheticTradeFeeRate checks that a synthetic trade fee rate is in [0,1].
func ValidateSyntheticTradeFeeRate(rate math.LegacyDec) error {
	if err := types.ValidateFee(rate); err != nil {
		return errors.Wrap(types.ErrInvalidSyntheticTradeFeeRate, err.Error())
	}
	return nil
}

// SyntheticTradeFeeRateOrZero returns the optional synthetic trade fee rate of a market launch, which
// defaults to zero.
func SyntheticTradeFeeRateOrZero(rate *math.LegacyDec) math.LegacyDec {
	if rate == nil {
		return math.LegacyZeroDec()
	}
	return *rate
}

// ValidateFundingV2Genesis checks the funding v2 and synthetic fee state of a genesis. Unset
// synthetic rates are accepted: they belong to a legacy genesis, and import sets them to the
// market's taker rate. genesisTime is nil when the genesis time is unknown, skipping the
// timestamp upper bounds.
//
//nolint:revive // genesis validation is intentionally whole-state
func (gs GenesisState) ValidateFundingV2Genesis(genesisTime *int64) error {
	perpetualMarkets := make(map[common.Hash]struct{}, len(gs.DerivativeMarkets))
	for i, market := range gs.DerivativeMarkets {
		if market == nil {
			continue
		}
		if market.IsPerpetual {
			perpetualMarkets[common.HexToHash(market.MarketId)] = struct{}{}
		}
		if market.SyntheticTradeFeeRate.IsNil() {
			continue
		}
		if err := ValidateSyntheticTradeFeeRate(market.SyntheticTradeFeeRate); err != nil {
			return fmt.Errorf("derivative_markets[%d] (%s): %w", i, market.MarketId, err)
		}
	}

	fundingV2Markets := 0
	for i := range gs.PerpetualMarketInfo {
		info := &gs.PerpetualMarketInfo[i]
		if info.FundingImpactNotional.IsNil() {
			continue
		}
		if err := ValidateFundingImpactNotional(info.FundingImpactNotional); err != nil {
			return fmt.Errorf("perpetual_market_info[%d] (%s): %w", i, info.MarketId, err)
		}
		if !info.IsFundingV2() {
			continue
		}
		if _, ok := perpetualMarkets[common.HexToHash(info.MarketId)]; !ok {
			return errors.Wrapf(types.ErrInvalidFundingImpactNotional, "perpetual_market_info[%d] (%s) is not a perpetual market", i, info.MarketId)
		}
		fundingV2Markets++
	}
	if fundingV2Markets > MaxFundingV2Markets {
		return errors.Wrapf(types.ErrTooManyFundingV2Markets, "%d markets use funding v2, at most %d allowed", fundingV2Markets, MaxFundingV2Markets)
	}

	for i, state := range gs.PerpetualMarketFundingState {
		funding := state.Funding
		if funding == nil {
			continue
		}
		if !funding.LastPremium.IsNil() && !funding.LastPremium.IsInValidRange() {
			return fmt.Errorf("perpetual_market_funding_state[%d] (%s): last_premium is out of range", i, state.MarketId)
		}
		if funding.PeriodStart > funding.LastTimestamp || funding.LastSampleTimestamp > funding.LastTimestamp {
			return fmt.Errorf(
				"perpetual_market_funding_state[%d] (%s): period_start %d and last_sample_timestamp %d must not exceed last_timestamp %d",
				i, state.MarketId, funding.PeriodStart, funding.LastSampleTimestamp, funding.LastTimestamp,
			)
		}
		if genesisTime != nil && funding.LastTimestamp > *genesisTime {
			return fmt.Errorf(
				"perpetual_market_funding_state[%d] (%s): last_timestamp %d is later than the genesis time %d",
				i, state.MarketId, funding.LastTimestamp, *genesisTime,
			)
		}
	}

	return gs.validateFundingImpactTakerState()
}

//nolint:revive // genesis validation is intentionally whole-state
func (gs GenesisState) validateFundingImpactTakerState() error {
	reserved := FundingImpactTakerSubaccountID
	for _, orderbook := range gs.SpotOrderbook {
		for _, order := range orderbook.Orders {
			if order != nil && order.SubaccountID() == reserved {
				return fundingImpactTakerGenesisError("spot order")
			}
		}
	}
	for _, orderbook := range gs.DerivativeOrderbook {
		for _, order := range orderbook.Orders {
			if order != nil && order.SubaccountID() == reserved {
				return fundingImpactTakerGenesisError("derivative order")
			}
		}
	}
	for _, orderbook := range gs.ConditionalDerivativeOrderbooks {
		if orderbook == nil {
			continue
		}
		for _, orders := range [][]*DerivativeLimitOrder{orderbook.LimitBuyOrders, orderbook.LimitSellOrders} {
			for _, order := range orders {
				if order != nil && order.SubaccountID() == reserved {
					return fundingImpactTakerGenesisError("conditional order")
				}
			}
		}
		for _, orders := range [][]*DerivativeMarketOrder{orderbook.MarketBuyOrders, orderbook.MarketSellOrders} {
			for _, order := range orders {
				if order != nil && order.SubaccountID() == reserved {
					return fundingImpactTakerGenesisError("conditional order")
				}
			}
		}
	}
	for _, position := range gs.Positions {
		if common.HexToHash(position.SubaccountId) == reserved {
			return fundingImpactTakerGenesisError("position")
		}
	}
	for _, record := range gs.SubaccountRiskProfiles {
		if record != nil && common.HexToHash(record.SubaccountId) == reserved {
			return fundingImpactTakerGenesisError("risk profile")
		}
	}
	for _, record := range gs.SubaccountMarketRiskModes {
		if record != nil && common.HexToHash(record.SubaccountId) == reserved {
			return fundingImpactTakerGenesisError("market risk mode")
		}
	}
	return nil
}

func fundingImpactTakerGenesisError(kind string) error {
	return errors.Wrapf(types.ErrBadSubaccountID, "the reserved funding impact taker subaccount %s cannot hold a %s", FundingImpactTakerSubaccountID.Hex(), kind)
}
