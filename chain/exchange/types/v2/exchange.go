package v2

import (
	"fmt"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
)

type TriggeredOrdersInMarket struct {
	Market             *DerivativeMarket
	MarkPrice          math.LegacyDec
	MarketOrders       []*DerivativeMarketOrder
	LimitOrders        []*DerivativeLimitOrder
	HasLimitBuyOrders  bool
	HasLimitSellOrders bool
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

func (s MarketStatus) SupportsOrderCancellations() bool {
	switch s {
	case MarketStatus_Active, MarketStatus_Demolished, MarketStatus_Expired, MarketStatus_Paused:
		return true
	default:
		return false
	}
}

func (p *PointsMultiplier) GetMultiplier(e ExecutionType) math.LegacyDec {
	if e.IsMaker() {
		return p.MakerPointsMultiplier
	}

	return p.TakerPointsMultiplier
}

// DerivativeOrder - IMutableDerivativeOrder implementation

func (o *DerivativeOrder) GetPrice() math.LegacyDec {
	return o.OrderInfo.Price
}

func (o *DerivativeOrder) GetQuantity() math.LegacyDec {
	return o.OrderInfo.Quantity
}
func (o *DerivativeOrder) GetFillable() math.LegacyDec {
	return o.GetQuantity()
}

func (o *DerivativeOrder) GetMargin() math.LegacyDec {
	return o.Margin
}
func (o *DerivativeOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeOrder) SetPrice(price math.LegacyDec) {
	o.OrderInfo.Price = price
}

func (o *DerivativeOrder) SetQuantity(quantity math.LegacyDec) {
	o.OrderInfo.Quantity = quantity
}

func (o *DerivativeOrder) SetMargin(margin math.LegacyDec) {
	o.Margin = margin
}

// DerivativeLimitOrder - IMutableDerivativeOrder implementation

func (m *DerivativeLimitOrder) GetPrice() math.LegacyDec {
	return m.OrderInfo.Price
}

func (m *DerivativeLimitOrder) GetQuantity() math.LegacyDec {
	return m.OrderInfo.Quantity
}
func (m *DerivativeLimitOrder) GetFillable() math.LegacyDec {
	return m.Fillable
}
func (m *DerivativeLimitOrder) GetMargin() math.LegacyDec {
	return m.Margin
}

func (m *DerivativeLimitOrder) GetSubaccountID() common.Hash {
	return m.OrderInfo.SubaccountID()
}

func (m *DerivativeLimitOrder) SetPrice(price math.LegacyDec) {
	m.OrderInfo.Price = price
}

func (m *DerivativeLimitOrder) SetQuantity(quantity math.LegacyDec) {
	m.OrderInfo.Quantity = quantity
	m.Fillable = quantity
}

func (m *DerivativeLimitOrder) SetMargin(margin math.LegacyDec) {
	m.Margin = margin
}

// DerivativeMarketOrder - IMutableDerivativeOrder implementation

func (o *DerivativeMarketOrder) GetPrice() math.LegacyDec {
	return o.OrderInfo.Price
}

func (o *DerivativeMarketOrder) GetQuantity() math.LegacyDec {
	return o.OrderInfo.Quantity
}
func (o *DerivativeMarketOrder) GetFillable() math.LegacyDec {
	return o.OrderInfo.Quantity
}
func (o *DerivativeMarketOrder) GetMargin() math.LegacyDec {
	return o.Margin
}
func (o *DerivativeMarketOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}
func (o *DerivativeMarketOrder) SetPrice(price math.LegacyDec) {
	o.OrderInfo.Price = price
}

func (o *DerivativeMarketOrder) SetQuantity(quantity math.LegacyDec) {
	o.OrderInfo.Quantity = quantity
}

func (o *DerivativeMarketOrder) SetMargin(margin math.LegacyDec) {
	o.Margin = margin
}

func (o *DerivativeMarketOrder) DebugString() string {
	return fmt.Sprintf("(q:%v, p:%v, m:%v, isLong: %v)", o.Quantity(), o.Price(), o.Margin, o.IsBuy())
}

// spot orders

func (m *SpotOrder) GetPrice() math.LegacyDec {
	return m.OrderInfo.Price
}
func (m *SpotLimitOrder) GetPrice() math.LegacyDec {
	return m.OrderInfo.Price
}
func (o *SpotMarketOrder) GetPrice() math.LegacyDec {
	return o.OrderInfo.Price
}

func (m *SpotOrder) GetQuantity() math.LegacyDec {
	return m.OrderInfo.Quantity
}
func (m *SpotLimitOrder) GetQuantity() math.LegacyDec {
	return m.OrderInfo.Quantity
}

func (o *SpotMarketOrder) GetQuantity() math.LegacyDec {
	return o.OrderInfo.Quantity
}
func (o *SpotMarketOrder) IsBuy() bool {
	return o.OrderType.IsBuy()
}
func (m *SpotOrder) GetFillable() math.LegacyDec {
	return m.OrderInfo.Quantity
}
func (o *SpotMarketOrder) GetFillable() math.LegacyDec {
	// no fillable for market order, but quantity works same in this case
	return o.OrderInfo.Quantity
}
func (m *SpotLimitOrder) GetFillable() math.LegacyDec {
	return m.Fillable
}

func (m *SpotOrder) GetSubaccountID() common.Hash {
	return m.OrderInfo.SubaccountID()
}
func (o *SpotMarketOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}
func (m *SpotLimitOrder) GetSubaccountID() common.Hash {
	return m.OrderInfo.SubaccountID()
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
