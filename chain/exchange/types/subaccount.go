package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewSubaccountOrderbookMetadata() *SubaccountOrderbookMetadata {
	return &SubaccountOrderbookMetadata{
		VanillaLimitOrderCount:      0,
		ReduceOnlyLimitOrderCount:   0,
		AggregateReduceOnlyQuantity: sdk.ZeroDec(),
		AggregateVanillaQuantity:    sdk.ZeroDec(),
	}
}

func (m *SubaccountOrderbookMetadata) GetOrderSideCount() uint32 {
	return m.VanillaLimitOrderCount + m.ReduceOnlyLimitOrderCount
}

func NewSubaccountOrder(o *DerivativeLimitOrder) *SubaccountOrder {
	return &SubaccountOrder{
		Price:        o.OrderInfo.Price,
		Quantity:     o.Fillable,
		IsReduceOnly: o.IsReduceOnly(),
	}
}

func (o *SubaccountOrder) IsVanilla() bool {
	return !o.IsReduceOnly
}

func (m *SubaccountOrderbookMetadata) ApplyDelta(d *SubaccountOrderbookMetadata) {
	if !d.AggregateReduceOnlyQuantity.IsZero() {
		m.AggregateReduceOnlyQuantity = m.AggregateReduceOnlyQuantity.Add(d.AggregateReduceOnlyQuantity)
	}
	if !d.AggregateVanillaQuantity.IsZero() {
		m.AggregateVanillaQuantity = m.AggregateVanillaQuantity.Add(d.AggregateVanillaQuantity)
	}
	m.VanillaLimitOrderCount += d.VanillaLimitOrderCount
	m.ReduceOnlyLimitOrderCount += d.ReduceOnlyLimitOrderCount
}
