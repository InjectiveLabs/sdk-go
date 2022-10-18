package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewSubaccountOrderbookMetadata() *SubaccountOrderbookMetadata {
	return &SubaccountOrderbookMetadata{
		VanillaLimitOrderCount:          0,
		ReduceOnlyLimitOrderCount:       0,
		AggregateReduceOnlyQuantity:     sdk.ZeroDec(),
		AggregateVanillaQuantity:        sdk.ZeroDec(),
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
	}
}

func (o *SubaccountOrder) IsVanilla() bool {
	return !o.IsReduceOnly
}

func (m *SubaccountOrderbookMetadata) AssertValid() {
	errStr := ""
	if m.AggregateVanillaQuantity.IsNegative() {
		errStr += "m.AggregateVanillaQuantity is negative with " + m.AggregateVanillaQuantity.String() + "\n"
	}
	if m.AggregateReduceOnlyQuantity.IsNegative() {
		errStr += "m.AggregateReduceOnlyQuantity is negative with " + m.AggregateReduceOnlyQuantity.String() + "\n"
	}
	if m.VanillaLimitOrderCount > 20 {
		errStr += fmt.Sprintf("m.AggregateVanillaQuantity is GT 20 %d\n", m.VanillaLimitOrderCount)
	}
	if m.ReduceOnlyLimitOrderCount > 20 {
		errStr += fmt.Sprintf("m.ReduceOnlyLimitOrderCount is GT 20 %d\n", m.ReduceOnlyLimitOrderCount)
	}
	if errStr != "" {
		panic(errStr)
	}
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
