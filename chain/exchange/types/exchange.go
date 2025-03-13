package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type IOrder interface {
	GetPrice() math.LegacyDec
	GetQuantity() math.LegacyDec
	GetFillable() math.LegacyDec
	IsBuy() bool
	GetSubaccountID() common.Hash
}

// IDerivativeOrder proto interface for wrapping all different variations of representations of derivative orders. Methods can be added as needed (make sure to add to every implementor)
type IDerivativeOrder interface {
	IOrder
	GetMargin() math.LegacyDec
	IsReduceOnly() bool
	IsVanilla() bool
}

type IMutableDerivativeOrder interface {
	IDerivativeOrder
	SetPrice(math.LegacyDec)
	SetQuantity(math.LegacyDec)
	SetMargin(math.LegacyDec)
}

func (al AtomicMarketOrderAccessLevel) IsValid() bool {
	switch al {
	case AtomicMarketOrderAccessLevel_Nobody,
		AtomicMarketOrderAccessLevel_SmartContractsOnly,
		AtomicMarketOrderAccessLevel_BeginBlockerSmartContractsOnly,
		AtomicMarketOrderAccessLevel_Everyone:
		return true
	default:
		return false
	}
}

type MatchedMarketDirection struct {
	MarketId    common.Hash
	BuysExists  bool
	SellsExists bool
}

type TriggeredOrdersInMarket struct {
	Market             *DerivativeMarket
	MarkPrice          math.LegacyDec
	MarketOrders       []*DerivativeMarketOrder
	LimitOrders        []*DerivativeLimitOrder
	HasLimitBuyOrders  bool
	HasLimitSellOrders bool
}

type TradingRewardAccountPoints struct {
	Account sdk.AccAddress
	Points  math.LegacyDec
}

func (m *DerivativeOrder) GetQuantity() math.LegacyDec {
	return m.OrderInfo.Quantity
}
func (m *DerivativeOrder) GetFillable() math.LegacyDec {
	return m.GetQuantity()
}

func (m *DerivativeOrder) GetMargin() math.LegacyDec {
	return m.Margin
}

func (m *SpotOrder) GetQuantity() math.LegacyDec {
	return m.OrderInfo.Quantity
}

func (m *SpotOrder) GetFillable() math.LegacyDec {
	return m.OrderInfo.Quantity
}
