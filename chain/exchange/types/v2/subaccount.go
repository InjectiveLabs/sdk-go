package v2

import (
	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
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

func NewSubaccountOrderMetadata() *SubaccountOrderMetadata {
	return &SubaccountOrderMetadata{
		CumulativeEOBVanillaQuantity:       math.LegacyZeroDec(),
		CumulativeEOBReduceOnlyQuantity:    math.LegacyZeroDec(),
		CumulativeBetterReduceOnlyQuantity: math.LegacyZeroDec(),
	}
}

type SubaccountOrderMetadata struct {
	CumulativeEOBVanillaQuantity       math.LegacyDec
	CumulativeEOBReduceOnlyQuantity    math.LegacyDec
	CumulativeBetterReduceOnlyQuantity math.LegacyDec
}

func NewSubaccountOrderResults() *SubaccountOrderResults {
	return &SubaccountOrderResults{
		ReduceOnlyOrders: make([]*SubaccountOrderData, 0),
		VanillaOrders:    make([]*SubaccountOrderData, 0),
		metadata:         NewSubaccountOrderMetadata(),
	}
}

type SubaccountOrderResults struct {
	ReduceOnlyOrders    []*SubaccountOrderData
	VanillaOrders       []*SubaccountOrderData
	metadata            *SubaccountOrderMetadata
	LastFoundOrderPrice *math.LegacyDec
	LastFoundOrderHash  *common.Hash
}

func (r *SubaccountOrderResults) AddSubaccountOrder(d *SubaccountOrderData) {
	if d.Order.IsReduceOnly {
		r.ReduceOnlyOrders = append(r.ReduceOnlyOrders, d)
		r.metadata.CumulativeEOBReduceOnlyQuantity = r.metadata.CumulativeEOBReduceOnlyQuantity.Add(d.Order.Quantity)
	} else {
		r.VanillaOrders = append(r.VanillaOrders, d)
		r.metadata.CumulativeEOBVanillaQuantity = r.metadata.CumulativeEOBVanillaQuantity.Add(d.Order.Quantity)
	}
}

func (r *SubaccountOrderResults) IncrementCumulativeBetterReduceOnlyQuantity(quantity math.LegacyDec) {
	r.metadata.CumulativeBetterReduceOnlyQuantity = r.metadata.CumulativeBetterReduceOnlyQuantity.Add(quantity)
}

func (r *SubaccountOrderResults) GetCumulativeEOBVanillaQuantity() math.LegacyDec {
	return r.metadata.CumulativeEOBVanillaQuantity
}

func (r *SubaccountOrderResults) GetCumulativeEOBReduceOnlyQuantity() math.LegacyDec {
	return r.metadata.CumulativeEOBReduceOnlyQuantity
}

func (r *SubaccountOrderResults) GetCumulativeBetterReduceOnlyQuantity() math.LegacyDec {
	return r.metadata.CumulativeBetterReduceOnlyQuantity
}
