package types

import (
	"github.com/ethereum/go-ethereum/common"
)

func (o *DerivativeOrder) GetNewDerivativeLimitOrder(hash common.Hash) *DerivativeLimitOrder {
	return &DerivativeLimitOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		Fillable:     o.OrderInfo.Quantity,
		TriggerPrice: o.TriggerPrice,
		Hash:         hash.Bytes(),
	}
}

func (o *DerivativeOrder) IsReduceOnly() bool {
	return o.Margin.IsZero()
}
