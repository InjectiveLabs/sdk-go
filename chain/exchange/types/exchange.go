package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type MatchedMarketDirection struct {
	MarketId    common.Hash
	BuysExists  bool
	SellsExists bool
}

func (m *DerivativeMarket) MarketID() common.Hash {
	return common.HexToHash(m.MarketId)
}

func (m *SpotMarket) MarketID() common.Hash {
	return common.HexToHash(m.MarketId)
}

func (e ExecutionType) IsMarket() bool {
	return e == ExecutionType_Market
}

func (e ExecutionType) IsMaker() bool {
	return !e.IsTaker()
}

func (e ExecutionType) IsTaker() bool {
	return e == ExecutionType_Market || e == ExecutionType_LimitMatchNewOrder
}

func (m *SpotMarket) StatusSupportsOrderCancellations() bool {
	if m == nil {
		return false
	}
	return m.Status.supportsOrderCancellations()
}

func (m *DerivativeMarket) StatusSupportsOrderCancellations() bool {
	if m == nil {
		return false
	}
	return m.Status.supportsOrderCancellations()
}

func (s MarketStatus) supportsOrderCancellations() bool {
	switch s {
	case MarketStatus_Active, MarketStatus_Suspended, MarketStatus_Demolished, MarketStatus_Expired:
		return true
	case MarketStatus_Paused:
		return false
	default:
		return false
	}
}

type TradingRewardAccountPoints struct {
	Account sdk.AccAddress
	Points  sdk.Dec
}

func (p *PointsMultiplier) GetEffectiveMultiplier(e ExecutionType, isRegisteredDMMCache map[string]bool, account sdk.AccAddress) sdk.Dec {
	if isRegisteredDMMCache[account.String()] {
		return sdk.ZeroDec()
	}

	if e.IsMaker() {
		return p.MakerPointsMultiplier
	}

	return p.TakerPointsMultiplier
}

func (p *PointsMultiplier) GetEffectiveMultipliers(isRegisteredDMMCache map[string]bool, account sdk.AccAddress) PointsMultiplier {
	if isRegisteredDMMCache[account.String()] {
		return PointsMultiplier{
			MakerPointsMultiplier: sdk.ZeroDec(),
			TakerPointsMultiplier: sdk.ZeroDec(),
		}
	}

	return PointsMultiplier{
		MakerPointsMultiplier: p.MakerPointsMultiplier,
		TakerPointsMultiplier: p.TakerPointsMultiplier,
	}
}
