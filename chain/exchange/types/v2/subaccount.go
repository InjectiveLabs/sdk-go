package v2

import (
	"cosmossdk.io/math"
)

func NewSubaccountOrderbookMetadata() *SubaccountOrderbookMetadata {
	return &SubaccountOrderbookMetadata{
		VanillaLimitOrderCount:          0,
		ReduceOnlyLimitOrderCount:       0,
		AggregateReduceOnlyQuantity:     math.LegacyZeroDec(),
		AggregateVanillaQuantity:        math.LegacyZeroDec(),
		VanillaConditionalOrderCount:    0,
		ReduceOnlyConditionalOrderCount: 0,
	}
}

func (m *SubaccountOrderbookMetadata) GetOrderSideCount() uint32 {
	return m.VanillaLimitOrderCount + m.ReduceOnlyLimitOrderCount + m.VanillaConditionalOrderCount + m.ReduceOnlyConditionalOrderCount
}

func NewSubaccountOrder(o *DerivativeLimitOrder) *SubaccountOrder {
	return &SubaccountOrder{
		Price:        o.OrderInfo.Price,
		Quantity:     o.Fillable,
		IsReduceOnly: o.IsReduceOnly(),
		Cid:          o.Cid(),
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
