package types

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

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

type TradingRewardAccountPoints struct {
	Account sdk.AccAddress
	Points  math.LegacyDec
}

func (p *PointsMultiplier) GetMultiplier(e ExecutionType) math.LegacyDec {
	if e.IsMaker() {
		return p.MakerPointsMultiplier
	}

	return p.TakerPointsMultiplier
}

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

func (o *DerivativeLimitOrder) GetPrice() math.LegacyDec {
	return o.OrderInfo.Price
}

func (o *DerivativeLimitOrder) GetQuantity() math.LegacyDec {
	return o.OrderInfo.Quantity
}
func (o *DerivativeLimitOrder) GetFillable() math.LegacyDec {
	return o.Fillable
}
func (o *DerivativeLimitOrder) GetMargin() math.LegacyDec {
	return o.Margin
}

func (o *DerivativeLimitOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeLimitOrder) SetPrice(price math.LegacyDec) {
	o.OrderInfo.Price = price
}

func (o *DerivativeLimitOrder) SetQuantity(quantity math.LegacyDec) {
	o.OrderInfo.Quantity = quantity
	o.Fillable = quantity
}

func (o *DerivativeLimitOrder) SetMargin(margin math.LegacyDec) {
	o.Margin = margin
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

func (o *SpotOrder) GetPrice() math.LegacyDec {
	return o.OrderInfo.Price
}
func (o *SpotLimitOrder) GetPrice() math.LegacyDec {
	return o.OrderInfo.Price
}
func (o *SpotMarketOrder) GetPrice() math.LegacyDec {
	return o.OrderInfo.Price
}

func (o *SpotOrder) GetQuantity() math.LegacyDec {
	return o.OrderInfo.Quantity
}
func (o *SpotLimitOrder) GetQuantity() math.LegacyDec {
	return o.OrderInfo.Quantity
}

func (o *SpotMarketOrder) GetQuantity() math.LegacyDec {
	return o.OrderInfo.Quantity
}
func (o *SpotMarketOrder) IsBuy() bool {
	return o.OrderType.IsBuy()
}
func (o *SpotOrder) GetFillable() math.LegacyDec {
	return o.OrderInfo.Quantity
}
func (o *SpotMarketOrder) GetFillable() math.LegacyDec {
	// no fillable for market order, but quantity works same in this case
	return o.OrderInfo.Quantity
}
func (o *SpotLimitOrder) GetFillable() math.LegacyDec {
	return o.Fillable
}

func (o *SpotOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}
func (o *SpotMarketOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}
func (o *SpotLimitOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
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
