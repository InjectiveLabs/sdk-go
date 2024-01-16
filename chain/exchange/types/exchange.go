package types

import (
	"fmt"

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
	MarkPrice          sdk.Dec
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
	Points  sdk.Dec
}

func (p *PointsMultiplier) GetMultiplier(e ExecutionType) sdk.Dec {
	if e.IsMaker() {
		return p.MakerPointsMultiplier
	}

	return p.TakerPointsMultiplier
}

type IOrder interface {
	GetPrice() sdk.Dec
	GetQuantity() sdk.Dec
	GetFillable() sdk.Dec
	IsBuy() bool
	GetSubaccountID() common.Hash
}

// IDerivativeOrder proto interface for wrapping all different variations of representations of derivative orders. Methods can be added as needed (make sure to add to every implementor)
type IDerivativeOrder interface {
	IOrder
	GetMargin() sdk.Dec
	IsReduceOnly() bool
	IsVanilla() bool
}

type IMutableDerivativeOrder interface {
	IDerivativeOrder
	SetPrice(sdk.Dec)
	SetQuantity(sdk.Dec)
	SetMargin(sdk.Dec)
}

// DerivativeOrder - IMutableDerivativeOrder implementation

func (o *DerivativeOrder) GetPrice() sdk.Dec {
	return o.OrderInfo.Price
}

func (o *DerivativeOrder) GetQuantity() sdk.Dec {
	return o.OrderInfo.Quantity
}
func (o *DerivativeOrder) GetFillable() sdk.Dec {
	return o.GetQuantity()
}

func (o *DerivativeOrder) GetMargin() sdk.Dec {
	return o.Margin
}
func (o *DerivativeOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeOrder) SetPrice(price sdk.Dec) {
	o.OrderInfo.Price = price
}

func (o *DerivativeOrder) SetQuantity(quantity sdk.Dec) {
	o.OrderInfo.Quantity = quantity
}

func (o *DerivativeOrder) SetMargin(margin sdk.Dec) {
	o.Margin = margin
}

// DerivativeLimitOrder - IMutableDerivativeOrder implementation

func (o *DerivativeLimitOrder) GetPrice() sdk.Dec {
	return o.OrderInfo.Price
}

func (o *DerivativeLimitOrder) GetQuantity() sdk.Dec {
	return o.OrderInfo.Quantity
}
func (o *DerivativeLimitOrder) GetFillable() sdk.Dec {
	return o.Fillable
}
func (o *DerivativeLimitOrder) GetMargin() sdk.Dec {
	return o.Margin
}

func (o *DerivativeLimitOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeLimitOrder) SetPrice(price sdk.Dec) {
	o.OrderInfo.Price = price
}

func (o *DerivativeLimitOrder) SetQuantity(quantity sdk.Dec) {
	o.OrderInfo.Quantity = quantity
	o.Fillable = quantity
}

func (o *DerivativeLimitOrder) SetMargin(margin sdk.Dec) {
	o.Margin = margin
}

// DerivativeMarketOrder - IMutableDerivativeOrder implementation

func (o *DerivativeMarketOrder) GetPrice() sdk.Dec {
	return o.OrderInfo.Price
}

func (o *DerivativeMarketOrder) GetQuantity() sdk.Dec {
	return o.OrderInfo.Quantity
}
func (o *DerivativeMarketOrder) GetFillable() sdk.Dec {
	return o.OrderInfo.Quantity
}
func (o *DerivativeMarketOrder) GetMargin() sdk.Dec {
	return o.Margin
}
func (o *DerivativeMarketOrder) GetSubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}
func (o *DerivativeMarketOrder) SetPrice(price sdk.Dec) {
	o.OrderInfo.Price = price
}

func (o *DerivativeMarketOrder) SetQuantity(quantity sdk.Dec) {
	o.OrderInfo.Quantity = quantity
}

func (o *DerivativeMarketOrder) SetMargin(margin sdk.Dec) {
	o.Margin = margin
}

func (o *DerivativeMarketOrder) DebugString() string {
	return fmt.Sprintf("(q:%v, p:%v, m:%v, isLong: %v)", o.Quantity(), o.Price(), o.Margin, o.IsBuy())
}

// spot orders

func (o *SpotOrder) GetPrice() sdk.Dec {
	return o.OrderInfo.Price
}
func (o *SpotLimitOrder) GetPrice() sdk.Dec {
	return o.OrderInfo.Price
}
func (o *SpotMarketOrder) GetPrice() sdk.Dec {
	return o.OrderInfo.Price
}

func (o *SpotOrder) GetQuantity() sdk.Dec {
	return o.OrderInfo.Quantity
}
func (o *SpotLimitOrder) GetQuantity() sdk.Dec {
	return o.OrderInfo.Quantity
}

func (o *SpotMarketOrder) GetQuantity() sdk.Dec {
	return o.OrderInfo.Quantity
}
func (o *SpotMarketOrder) IsBuy() bool {
	return o.OrderType.IsBuy()
}
func (o *SpotOrder) GetFillable() sdk.Dec {
	return o.OrderInfo.Quantity
}
func (o *SpotMarketOrder) GetFillable() sdk.Dec {
	// no fillable for market order, but quantity works same in this case
	return o.OrderInfo.Quantity
}
func (o *SpotLimitOrder) GetFillable() sdk.Dec {
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
