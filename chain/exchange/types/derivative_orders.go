package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (order *DerivativeLimitOrder) GetCancelDepositDelta(makerFeeRate sdk.Dec) *DepositDelta {
	depositDelta := &DepositDelta{
		AvailableBalanceDelta: sdk.ZeroDec(),
		TotalBalanceDelta:     sdk.ZeroDec(),
	}
	if order.IsVanilla() {
		// Refund = (Fillable / Quantity) * (Margin + Price * Quantity * MakerFeeRate)
		fillableFraction := order.Fillable.Quo(order.OrderInfo.Quantity)
		notional := order.OrderInfo.Price.Mul(order.OrderInfo.Quantity)
		marginHoldRefund := fillableFraction.Mul(order.Margin.Add(notional.Mul(makerFeeRate)))
		depositDelta.AvailableBalanceDelta = marginHoldRefund
	}
	return depositDelta
}

func (c *DerivativeMarketOrderCancel) GetCancelDepositDelta() *DepositDelta {
	order := c.MarketOrder
	// no market order quantity was executed, so refund the entire margin hold
	if order.IsVanilla() && c.CancelQuantity.Equal(order.OrderInfo.Quantity) {
		return &DepositDelta{
			AvailableBalanceDelta: order.MarginHold,
			TotalBalanceDelta:     sdk.ZeroDec(),
		}
	}
	return nil
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

func (o *DerivativeOrder) IsVanilla() bool {
	return !o.IsReduceOnly()
}

func (o *DerivativeMarketOrder) IsVanilla() bool {
	return !o.IsReduceOnly()
}

func (o *DerivativeLimitOrder) IsVanilla() bool {
	return !o.IsReduceOnly()
}

func (o *DerivativeOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeMarketOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeLimitOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *OrderInfo) SubaccountID() common.Hash {
	return common.HexToHash(o.SubaccountId)
}
