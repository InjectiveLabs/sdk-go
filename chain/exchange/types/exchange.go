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
	case MarketStatus_Active, MarketStatus_Demolished, MarketStatus_Expired:
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

func (p *PointsMultiplier) GetMultiplier(e ExecutionType) sdk.Dec {
	if e.IsMaker() {
		return p.MakerPointsMultiplier
	}

	return p.TakerPointsMultiplier
}

type IOrder interface {
	GetPrice() sdk.Dec
	GetQuantity() sdk.Dec
	IsBuy() bool
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

func (o *DerivativeOrder) GetMargin() sdk.Dec {
	return o.Margin
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

func (o *DerivativeLimitOrder) GetMargin() sdk.Dec {
	return o.Margin
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

func (o *DerivativeMarketOrder) GetMargin() sdk.Dec {
	return o.Margin
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
