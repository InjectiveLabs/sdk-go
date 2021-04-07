package types

import (
	"github.com/ethereum/go-ethereum/common"
)

func (o *DerivativeOrder) GetNewDerivativeLimitOrder(orderHash common.Hash) *DerivativeLimitOrder {
	return &DerivativeLimitOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		Fillable:     o.OrderInfo.Quantity,
		TriggerPrice: o.TriggerPrice,
		OrderHash:    orderHash.Bytes(),
	}
}

func (o *DerivativeOrder) IsReduceOnly() bool {
	return o.Margin.IsZero()
}

func (o *DerivativeMarketOrder) IsReduceOnly() bool {
	return o.MarginHold.IsZero()
}

func (o *DerivativeLimitOrder) IsReduceOnly() bool {
	return o.Margin.IsZero()
}
