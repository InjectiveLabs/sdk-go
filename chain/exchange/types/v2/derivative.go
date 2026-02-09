package v2 //nolint:revive // ok

import (
	"bytes"
	"sort"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

// GetIsOrderLess returns true if the order is less than the other order
func GetIsOrderLess( //nolint:revive // ok
	referencePrice,
	order1Price,
	order2Price math.LegacyDec,
	order1IsBuy,
	order2IsBuy,
	isSortingFromWorstToBest bool,
) bool {
	var firstDistanceToReferencePrice, secondDistanceToReferencePrice math.LegacyDec

	if order1IsBuy {
		firstDistanceToReferencePrice = referencePrice.Sub(order1Price)
	} else {
		firstDistanceToReferencePrice = order1Price.Sub(referencePrice)
	}

	if order2IsBuy {
		secondDistanceToReferencePrice = referencePrice.Sub(order2Price)
	} else {
		secondDistanceToReferencePrice = order2Price.Sub(referencePrice)
	}

	if isSortingFromWorstToBest {
		return firstDistanceToReferencePrice.GT(secondDistanceToReferencePrice)
	}

	return firstDistanceToReferencePrice.LT(secondDistanceToReferencePrice)
}

// GetDerivativeOrdersToCancelUpToAmount returns the Derivative orders to cancel up to a given amount
func GetDerivativeOrdersToCancelUpToAmount(
	market *DerivativeMarket,
	orders []*TrimmedDerivativeLimitOrder,
	strategy CancellationStrategy,
	referencePrice *math.LegacyDec,
	quoteAmount math.LegacyDec,
) ([]*TrimmedDerivativeLimitOrder, bool) {
	switch strategy {
	case CancellationStrategy_FromWorstToBest:
		sort.SliceStable(orders, func(i, j int) bool {
			return GetIsOrderLess(*referencePrice, orders[i].Price, orders[j].Price, orders[i].IsBuy, orders[j].IsBuy, true)
		})
	case CancellationStrategy_FromBestToWorst:
		sort.SliceStable(orders, func(i, j int) bool {
			return GetIsOrderLess(*referencePrice, orders[i].Price, orders[j].Price, orders[i].IsBuy, orders[j].IsBuy, false)
		})
	default:
		// do nothing
	}

	positiveMakerFeePart := math.LegacyMaxDec(math.LegacyZeroDec(), market.MakerFeeRate)

	ordersToCancel := make([]*TrimmedDerivativeLimitOrder, 0)
	cumulativeQuoteAmount := math.LegacyZeroDec()

	for _, order := range orders {
		hasSufficientQuote := cumulativeQuoteAmount.GTE(quoteAmount)
		if hasSufficientQuote {
			break
		}

		ordersToCancel = append(ordersToCancel, order)

		notional := order.Fillable.Mul(order.Price)
		fee := notional.Mul(positiveMakerFeePart)
		remainingMargin := order.Margin.Mul(order.Fillable).Quo(order.Quantity)
		cumulativeQuoteAmount = cumulativeQuoteAmount.Add(remainingMargin).Add(fee)
	}

	hasProcessedFullAmount := cumulativeQuoteAmount.GTE(quoteAmount)
	return ordersToCancel, hasProcessedFullAmount
}

// ReduceOnlyOrdersTracker maps subaccountID => orders
type ReduceOnlyOrdersTracker map[common.Hash][]*DerivativeLimitOrder

func NewReduceOnlyOrdersTracker() ReduceOnlyOrdersTracker {
	return make(map[common.Hash][]*DerivativeLimitOrder)
}

func (r ReduceOnlyOrdersTracker) GetSortedSubaccountIDs() []common.Hash {
	subaccountIDs := make([]common.Hash, 0, len(r))
	for subaccountID := range r {
		subaccountIDs = append(subaccountIDs, subaccountID)
	}

	sort.SliceStable(subaccountIDs, func(i, j int) bool {
		return bytes.Compare(subaccountIDs[i].Bytes(), subaccountIDs[j].Bytes()) < 0
	})
	return subaccountIDs
}

func (r ReduceOnlyOrdersTracker) GetCumulativeOrderQuantity(subaccountID common.Hash) math.LegacyDec {
	cumulativeQuantity := math.LegacyZeroDec()
	orders := r[subaccountID]

	for idx := range orders {
		cumulativeQuantity = cumulativeQuantity.Add(orders[idx].Fillable)
	}

	return cumulativeQuantity
}

func (r ReduceOnlyOrdersTracker) AppendOrder(subaccountID common.Hash, order *DerivativeLimitOrder) {
	orders, ok := r[subaccountID]
	if !ok {
		r[subaccountID] = []*DerivativeLimitOrder{order}
		return
	}

	r[subaccountID] = append(orders, order)
}

// ModifiedPositionCache maps marketID => subaccountID => position or nil indicator
type ModifiedPositionCache map[common.Hash]map[common.Hash]*Position

func NewModifiedPositionCache() ModifiedPositionCache {
	return make(map[common.Hash]map[common.Hash]*Position)
}

func (c ModifiedPositionCache) SetPosition(marketID, subaccountID common.Hash, position *Position) {
	if position == nil {
		return
	}

	v, ok := c[marketID]
	if !ok {
		v = make(map[common.Hash]*Position)
		c[marketID] = v
	}

	v[subaccountID] = position
}

func (c ModifiedPositionCache) SetPositionIndicator(marketID, subaccountID common.Hash) {
	v, ok := c[marketID]
	if !ok {
		v = make(map[common.Hash]*Position)
		c[marketID] = v
	}

	v[subaccountID] = nil
}

func (c ModifiedPositionCache) GetPosition(marketID, subaccountID common.Hash) *Position {
	v, ok := c[marketID]
	if !ok {
		return nil
	}
	return v[subaccountID]
}

func (c ModifiedPositionCache) GetSortedSubaccountIDsByMarket(marketID common.Hash) []common.Hash {
	v, ok := c[marketID]
	if !ok {
		return nil
	}

	subaccountIDs := make([]common.Hash, 0, len(v))
	for subaccountID := range v {
		subaccountIDs = append(subaccountIDs, subaccountID)
	}

	sort.SliceStable(subaccountIDs, func(i, j int) bool {
		return bytes.Compare(subaccountIDs[i].Bytes(), subaccountIDs[j].Bytes()) < 0
	})

	return subaccountIDs
}

func (c ModifiedPositionCache) HasAnyModifiedPositionsInMarket(marketID common.Hash) bool {
	_, found := c[marketID]
	return found
}

func (c ModifiedPositionCache) HasPositionBeenModified(marketID, subaccountID common.Hash) bool {
	v, ok := c[marketID]
	if !ok {
		return false
	}
	_, found := v[subaccountID]
	return found
}

type DerivativeMatchingExpansionData struct {
	TransientLimitBuyExpansions    []*DerivativeOrderStateExpansion
	TransientLimitSellExpansions   []*DerivativeOrderStateExpansion
	RestingLimitBuyExpansions      []*DerivativeOrderStateExpansion
	RestingLimitSellExpansions     []*DerivativeOrderStateExpansion
	RestingLimitBuyOrderCancels    []*DerivativeLimitOrder
	RestingLimitSellOrderCancels   []*DerivativeLimitOrder
	TransientLimitBuyOrderCancels  []*DerivativeLimitOrder
	TransientLimitSellOrderCancels []*DerivativeLimitOrder
	ClearingPrice                  math.LegacyDec
	ClearingQuantity               math.LegacyDec
	MarketBalanceDelta             math.LegacyDec
	OpenInterestDelta              math.LegacyDec
	NewRestingLimitBuyOrders       []*DerivativeLimitOrder // transient buy orders that become new resting limit orders
	NewRestingLimitSellOrders      []*DerivativeLimitOrder // transient sell orders that become new resting limit orders
}

func NewDerivativeMatchingExpansionData(clearingPrice, clearingQuantity math.LegacyDec) *DerivativeMatchingExpansionData {
	return &DerivativeMatchingExpansionData{
		TransientLimitBuyExpansions:    make([]*DerivativeOrderStateExpansion, 0),
		TransientLimitSellExpansions:   make([]*DerivativeOrderStateExpansion, 0),
		RestingLimitBuyExpansions:      make([]*DerivativeOrderStateExpansion, 0),
		RestingLimitSellExpansions:     make([]*DerivativeOrderStateExpansion, 0),
		RestingLimitBuyOrderCancels:    make([]*DerivativeLimitOrder, 0),
		RestingLimitSellOrderCancels:   make([]*DerivativeLimitOrder, 0),
		TransientLimitBuyOrderCancels:  make([]*DerivativeLimitOrder, 0),
		TransientLimitSellOrderCancels: make([]*DerivativeLimitOrder, 0),
		ClearingPrice:                  clearingPrice,
		ClearingQuantity:               clearingQuantity,
		OpenInterestDelta:              math.LegacyZeroDec(),
		MarketBalanceDelta:             math.LegacyZeroDec(),
		NewRestingLimitBuyOrders:       make([]*DerivativeLimitOrder, 0),
		NewRestingLimitSellOrders:      make([]*DerivativeLimitOrder, 0),
	}
}

func (e *DerivativeMatchingExpansionData) AddExpansion(isBuy, isTransient bool, expansion *DerivativeOrderStateExpansion) {
	e.MarketBalanceDelta = e.MarketBalanceDelta.Add(expansion.MarketBalanceDelta)

	switch {
	case isBuy && isTransient:
		e.TransientLimitBuyExpansions = append(e.TransientLimitBuyExpansions, expansion)
	case isBuy && !isTransient:
		e.RestingLimitBuyExpansions = append(e.RestingLimitBuyExpansions, expansion)
	case !isBuy && isTransient:
		e.TransientLimitSellExpansions = append(e.TransientLimitSellExpansions, expansion)
	case !isBuy && !isTransient:
		e.RestingLimitSellExpansions = append(e.RestingLimitSellExpansions, expansion)
	}
}

func (e *DerivativeMatchingExpansionData) AddNewBuyRestingLimitOrder(order *DerivativeLimitOrder) {
	e.NewRestingLimitBuyOrders = append(e.NewRestingLimitBuyOrders, order)
}

func (e *DerivativeMatchingExpansionData) AddNewSellRestingLimitOrder(order *DerivativeLimitOrder) {
	e.NewRestingLimitSellOrders = append(e.NewRestingLimitSellOrders, order)
}

func (e *DerivativeMatchingExpansionData) GetLimitMatchingDerivativeBatchExecutionData(
	market DerivativeMarketI,
	markPrice math.LegacyDec,
	funding *PerpetualMarketFunding,
	positionStates map[common.Hash]*PositionState,
) *DerivativeBatchExecutionData {
	depositDeltas := types.NewDepositDeltas()
	tradingRewardPoints := types.NewTradingRewardPoints()

	// process undermargined resting limit order forced cancellations
	cancelLimitOrdersEvents, restingOrderCancelledDeltas, transientOrderCancelledDeltas :=
		e.applyCancellationsAndGetDerivativeLimitCancelEvents(
			market,
			market.GetMakerFeeRate(),
			market.GetTakerFeeRate(),
			depositDeltas,
		)

	positions, positionSubaccountIDs := GetPositionSliceData(positionStates)

	transientLimitBuyOrderBatchEvent, transientLimitBuyFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(
		true,
		ExecutionType_LimitMatchNewOrder,
		market,
		funding,
		e.TransientLimitBuyExpansions,
		depositDeltas,
		tradingRewardPoints,
		false,
	)
	restingLimitBuyOrderBatchEvent, restingLimitBuyFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(
		true,
		ExecutionType_LimitMatchRestingOrder,
		market,
		funding,
		e.RestingLimitBuyExpansions,
		depositDeltas,
		tradingRewardPoints,
		false,
	)
	transientLimitSellOrderBatchEvent, transientLimitSellFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(
		false,
		ExecutionType_LimitMatchNewOrder,
		market,
		funding,
		e.TransientLimitSellExpansions,
		depositDeltas,
		tradingRewardPoints,
		false,
	)
	restingLimitSellOrderBatchEvent, restingLimitSellFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(
		false,
		ExecutionType_LimitMatchRestingOrder,
		market,
		funding,
		e.RestingLimitSellExpansions,
		depositDeltas,
		tradingRewardPoints,
		false,
	)

	restingOrderFilledDeltas := mergeDerivativeLimitOrderFilledDeltas(restingLimitBuyFilledDeltas, restingLimitSellFilledDeltas)
	transientOrderFilledDeltas := mergeDerivativeLimitOrderFilledDeltas(transientLimitBuyFilledDeltas, transientLimitSellFilledDeltas)

	// sort keys since map iteration is non-deterministic
	depositDeltaKeys := depositDeltas.GetSortedSubaccountKeys()

	vwapData := NewVwapData()
	vwapData = vwapData.ApplyExecution(e.ClearingPrice, e.ClearingQuantity)

	var newOrdersEvent *EventNewDerivativeOrders
	if len(e.NewRestingLimitBuyOrders) > 0 || len(e.NewRestingLimitSellOrders) > 0 {
		newOrdersEvent = &EventNewDerivativeOrders{
			MarketId:   market.MarketID().String(),
			BuyOrders:  e.NewRestingLimitBuyOrders,
			SellOrders: e.NewRestingLimitSellOrders,
		}
	}

	// Final Step: Store the DerivativeBatchExecutionData for future reduction/processing
	batch := &DerivativeBatchExecutionData{
		Market:                                market,
		MarkPrice:                             markPrice,
		Funding:                               funding,
		DepositDeltas:                         depositDeltas,
		DepositSubaccountIDs:                  depositDeltaKeys,
		TradingRewards:                        tradingRewardPoints,
		Positions:                             positions,
		MarketBalanceDelta:                    market.NotionalToChainFormat(e.MarketBalanceDelta),
		OpenInterestDelta:                     e.OpenInterestDelta,
		PositionSubaccountIDs:                 positionSubaccountIDs,
		RestingLimitOrderFilledDeltas:         restingOrderFilledDeltas,
		TransientLimitOrderFilledDeltas:       transientOrderFilledDeltas,
		RestingLimitOrderCancelledDeltas:      restingOrderCancelledDeltas,
		TransientLimitOrderCancelledDeltas:    transientOrderCancelledDeltas,
		MarketBuyOrderExecutionEvent:          nil,
		MarketSellOrderExecutionEvent:         nil,
		RestingLimitBuyOrderExecutionEvent:    restingLimitBuyOrderBatchEvent,
		RestingLimitSellOrderExecutionEvent:   restingLimitSellOrderBatchEvent,
		TransientLimitBuyOrderExecutionEvent:  transientLimitBuyOrderBatchEvent,
		TransientLimitSellOrderExecutionEvent: transientLimitSellOrderBatchEvent,
		NewOrdersEvent:                        newOrdersEvent,
		CancelLimitOrderEvents:                cancelLimitOrdersEvents,
		CancelMarketOrderEvents:               nil,
		VwapData:                              vwapData,
	}

	return batch
}

func (e *DerivativeMatchingExpansionData) applyCancellationsAndGetDerivativeLimitCancelEvents(
	market DerivativeMarketI,
	makerFeeRate math.LegacyDec,
	takerFeeRate math.LegacyDec,
	depositDeltas types.DepositDeltas,
) (
	cancelOrdersEvent []*EventCancelDerivativeOrder,
	restingOrderCancelledDeltas []*DerivativeLimitOrderDelta,
	transientOrderCancelledDeltas []*DerivativeLimitOrderDelta,
) {
	marketIDHex := market.MarketID().Hex()

	cancelOrdersEvent = make(
		[]*EventCancelDerivativeOrder,
		0,
		len(e.RestingLimitBuyOrderCancels)+
			len(e.RestingLimitSellOrderCancels)+
			len(e.TransientLimitBuyOrderCancels)+
			len(e.TransientLimitSellOrderCancels),
	)
	restingOrderCancelledDeltas = make(
		[]*DerivativeLimitOrderDelta,
		0,
		len(e.RestingLimitBuyOrderCancels)+len(e.RestingLimitSellOrderCancels),
	)
	transientOrderCancelledDeltas = make(
		[]*DerivativeLimitOrderDelta,
		0,
		len(e.TransientLimitBuyOrderCancels)+len(e.TransientLimitSellOrderCancels),
	)

	for idx := range e.RestingLimitBuyOrderCancels {
		order := e.RestingLimitBuyOrderCancels[idx]

		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas, market)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
		restingOrderCancelledDeltas = append(restingOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   math.LegacyZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	for idx := range e.RestingLimitSellOrderCancels {
		order := e.RestingLimitSellOrderCancels[idx]

		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas, market)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
		restingOrderCancelledDeltas = append(restingOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   math.LegacyZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	for idx := range e.TransientLimitBuyOrderCancels {
		order := e.TransientLimitBuyOrderCancels[idx]

		applyDerivativeLimitCancellation(order, takerFeeRate, depositDeltas, market)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
		transientOrderCancelledDeltas = append(transientOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   math.LegacyZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	for idx := range e.TransientLimitSellOrderCancels {
		order := e.TransientLimitSellOrderCancels[idx]
		applyDerivativeLimitCancellation(order, takerFeeRate, depositDeltas, market)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
		transientOrderCancelledDeltas = append(transientOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   math.LegacyZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	return cancelOrdersEvent, restingOrderCancelledDeltas, transientOrderCancelledDeltas
}

func applyDerivativeLimitCancellation(
	order *DerivativeLimitOrder,
	orderFeeRate math.LegacyDec,
	depositDeltas types.DepositDeltas,
	market DerivativeMarketI,
) {
	// For vanilla orders, increment the available balance
	if order.IsVanilla() {
		depositDelta := order.GetCancelDepositDelta(orderFeeRate)
		chainFormatDepositDelta := types.DepositDelta{
			AvailableBalanceDelta: market.NotionalToChainFormat(depositDelta.AvailableBalanceDelta),
			TotalBalanceDelta:     market.NotionalToChainFormat(depositDelta.TotalBalanceDelta),
		}
		depositDeltas.ApplyDepositDelta(order.SubaccountID(), &chainFormatDepositDelta)
	}
}

type DerivativeMarketOrderExpansionData struct {
	MarketBuyExpansions          []*DerivativeOrderStateExpansion
	MarketSellExpansions         []*DerivativeOrderStateExpansion
	LimitBuyExpansions           []*DerivativeOrderStateExpansion
	LimitSellExpansions          []*DerivativeOrderStateExpansion
	RestingLimitBuyOrderCancels  []*DerivativeLimitOrder
	RestingLimitSellOrderCancels []*DerivativeLimitOrder
	MarketBuyOrderCancels        []*DerivativeMarketOrderCancel
	MarketSellOrderCancels       []*DerivativeMarketOrderCancel
	MarketBuyClearingPrice       math.LegacyDec
	MarketSellClearingPrice      math.LegacyDec
	MarketBuyClearingQuantity    math.LegacyDec
	MarketSellClearingQuantity   math.LegacyDec
	MarketBalanceDelta           math.LegacyDec
	OpenInterestDelta            math.LegacyDec
}

func (e *DerivativeMarketOrderExpansionData) SetSellExecutionData(
	marketOrderClearingPrice, marketOrderClearingQuantity math.LegacyDec,
	restingLimitOrderCancels []*DerivativeLimitOrder,
	marketOrderStateExpansions,
	restingLimitOrderStateExpansions []*DerivativeOrderStateExpansion,
	marketOrderCancels []*DerivativeMarketOrderCancel,
) {
	e.MarketSellClearingPrice = marketOrderClearingPrice
	e.MarketSellClearingQuantity = marketOrderClearingQuantity
	e.RestingLimitBuyOrderCancels = restingLimitOrderCancels
	e.MarketSellExpansions = marketOrderStateExpansions
	e.LimitBuyExpansions = restingLimitOrderStateExpansions
	e.MarketSellOrderCancels = marketOrderCancels

	e.setExecutionData(marketOrderStateExpansions, restingLimitOrderStateExpansions)
}

func (e *DerivativeMarketOrderExpansionData) SetBuyExecutionData(
	marketOrderClearingPrice, marketOrderClearingQuantity math.LegacyDec,
	restingLimitOrderCancels []*DerivativeLimitOrder,
	marketOrderStateExpansions,
	restingLimitOrderStateExpansions []*DerivativeOrderStateExpansion,
	marketOrderCancels []*DerivativeMarketOrderCancel,
) {
	e.MarketBuyClearingPrice = marketOrderClearingPrice
	e.MarketBuyClearingQuantity = marketOrderClearingQuantity
	e.RestingLimitSellOrderCancels = restingLimitOrderCancels
	e.MarketBuyExpansions = marketOrderStateExpansions
	e.LimitSellExpansions = restingLimitOrderStateExpansions
	e.MarketBuyOrderCancels = marketOrderCancels

	e.setExecutionData(marketOrderStateExpansions, restingLimitOrderStateExpansions)
}

func (e *DerivativeMarketOrderExpansionData) setExecutionData(
	marketOrderStateExpansions,
	restingLimitOrderStateExpansions []*DerivativeOrderStateExpansion,
) {
	if e.MarketBalanceDelta.IsNil() {
		e.MarketBalanceDelta = math.LegacyZeroDec()
	}

	for idx := range marketOrderStateExpansions {
		stateExpansion := marketOrderStateExpansions[idx]
		e.MarketBalanceDelta = e.MarketBalanceDelta.Add(stateExpansion.MarketBalanceDelta)
	}
	for idx := range restingLimitOrderStateExpansions {
		stateExpansion := restingLimitOrderStateExpansions[idx]
		e.MarketBalanceDelta = e.MarketBalanceDelta.Add(stateExpansion.MarketBalanceDelta)
	}
}

func (e *DerivativeMarketOrderExpansionData) GetMarketDerivativeBatchExecutionData( //nolint:revive // ok
	market DerivativeMarketI,
	markPrice math.LegacyDec,
	funding *PerpetualMarketFunding,
	positionStates map[common.Hash]*PositionState,
	isLiquidation bool,
) *DerivativeBatchExecutionData {
	depositDeltas := types.NewDepositDeltas()
	tradingRewardPoints := types.NewTradingRewardPoints()

	// process undermargined limit order forced cancellations
	cancelLimitOrdersEvents, restingOrderCancelledDeltas := e.applyCancellationsAndGetDerivativeLimitCancelEvents(
		market,
		market.GetMakerFeeRate(),
		depositDeltas,
	)

	// process unfilled market order cancellations
	cancelMarketOrdersEvents := e.getDerivativeMarketCancelEvents(market.MarketID())

	positions, positionSubaccountIDs := GetPositionSliceData(positionStates)

	buyMarketOrderBatchEvent, _ := ApplyDeltasAndGetDerivativeOrderBatchEvent(
		true,
		ExecutionType_Market,
		market,
		funding,
		e.MarketBuyExpansions,
		depositDeltas,
		tradingRewardPoints,
		isLiquidation,
	)
	sellMarketOrderBatchEvent, _ := ApplyDeltasAndGetDerivativeOrderBatchEvent(
		false,
		ExecutionType_Market,
		market,
		funding,
		e.MarketSellExpansions,
		depositDeltas,
		tradingRewardPoints,
		isLiquidation,
	)

	restingLimitBuyOrderBatchEvent, limitBuyFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(
		true,
		ExecutionType_LimitFill,
		market,
		funding,
		e.LimitBuyExpansions,
		depositDeltas,
		tradingRewardPoints,
		false,
	)
	restingLimitSellOrderBatchEvent, limitSellFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(
		false,
		ExecutionType_LimitFill,
		market,
		funding,
		e.LimitSellExpansions,
		depositDeltas,
		tradingRewardPoints,
		false,
	)

	filledDeltas := mergeDerivativeLimitOrderFilledDeltas(limitBuyFilledDeltas, limitSellFilledDeltas)

	// sort keys since map iteration is non-deterministic
	depositDeltaKeys := depositDeltas.GetSortedSubaccountKeys()

	vwapData := NewVwapData()
	vwapData = vwapData.ApplyExecution(e.MarketBuyClearingPrice, e.MarketBuyClearingQuantity)
	vwapData = vwapData.ApplyExecution(e.MarketSellClearingPrice, e.MarketSellClearingQuantity)

	// Final Step: Store the DerivativeBatchExecutionData for future reduction/processing
	batch := &DerivativeBatchExecutionData{
		Market:                                market,
		MarkPrice:                             markPrice,
		Funding:                               funding,
		DepositDeltas:                         depositDeltas,
		DepositSubaccountIDs:                  depositDeltaKeys,
		TradingRewards:                        tradingRewardPoints,
		Positions:                             positions,
		MarketBalanceDelta:                    market.NotionalToChainFormat(e.MarketBalanceDelta),
		PositionSubaccountIDs:                 positionSubaccountIDs,
		TransientLimitOrderFilledDeltas:       nil,
		RestingLimitOrderFilledDeltas:         filledDeltas,
		TransientLimitOrderCancelledDeltas:    nil,
		RestingLimitOrderCancelledDeltas:      restingOrderCancelledDeltas,
		MarketBuyOrderExecutionEvent:          buyMarketOrderBatchEvent,
		MarketSellOrderExecutionEvent:         sellMarketOrderBatchEvent,
		RestingLimitBuyOrderExecutionEvent:    restingLimitBuyOrderBatchEvent,
		RestingLimitSellOrderExecutionEvent:   restingLimitSellOrderBatchEvent,
		TransientLimitBuyOrderExecutionEvent:  nil,
		TransientLimitSellOrderExecutionEvent: nil,
		NewOrdersEvent:                        nil,
		OpenInterestDelta:                     e.OpenInterestDelta,
		CancelLimitOrderEvents:                cancelLimitOrdersEvents,
		CancelMarketOrderEvents:               cancelMarketOrdersEvents,
		VwapData:                              vwapData,
	}
	return batch
}

func (e *DerivativeMarketOrderExpansionData) getDerivativeMarketCancelEvents(marketID common.Hash) []*EventCancelDerivativeOrder {
	marketIDHex := marketID.Hex()
	cancelOrdersEvent := make([]*EventCancelDerivativeOrder, 0, len(e.MarketBuyOrderCancels)+len(e.MarketSellOrderCancels))

	for idx := range e.MarketBuyOrderCancels {
		orderCancel := e.MarketBuyOrderCancels[idx]
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:          marketIDHex,
			IsLimitCancel:     false,
			MarketOrderCancel: orderCancel,
		})
	}

	for idx := range e.MarketSellOrderCancels {
		orderCancel := e.MarketSellOrderCancels[idx]
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:          marketIDHex,
			IsLimitCancel:     false,
			MarketOrderCancel: orderCancel,
		})
	}
	return cancelOrdersEvent
}

func (e *DerivativeMarketOrderExpansionData) applyCancellationsAndGetDerivativeLimitCancelEvents(
	market DerivativeMarketI,
	makerFeeRate math.LegacyDec,
	depositDeltas types.DepositDeltas,
) (
	cancelOrdersEvent []*EventCancelDerivativeOrder,
	restingOrderCancelledDeltas []*DerivativeLimitOrderDelta,
) {
	marketIDHex := market.MarketID().Hex()
	cancelOrdersEvent = make([]*EventCancelDerivativeOrder, 0, len(e.RestingLimitBuyOrderCancels)+len(e.RestingLimitSellOrderCancels))

	restingOrderCancelledDeltas = make(
		[]*DerivativeLimitOrderDelta, 0, len(e.RestingLimitBuyOrderCancels)+len(e.RestingLimitSellOrderCancels),
	)

	for idx := range e.RestingLimitBuyOrderCancels {
		order := e.RestingLimitBuyOrderCancels[idx]
		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas, market)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})

		restingOrderCancelledDeltas = append(restingOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   math.LegacyZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	for idx := range e.RestingLimitSellOrderCancels {
		order := e.RestingLimitSellOrderCancels[idx]
		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas, market)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})

		restingOrderCancelledDeltas = append(restingOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   math.LegacyZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}
	return cancelOrdersEvent, restingOrderCancelledDeltas
}

func mergeDerivativeLimitOrderFilledDeltas(d1, d2 []*DerivativeLimitOrderDelta) []*DerivativeLimitOrderDelta {
	filledDeltas := make([]*DerivativeLimitOrderDelta, 0, len(d1)+len(d2))
	if len(d1) > 0 {
		filledDeltas = append(filledDeltas, d1...)
	}
	if len(d2) > 0 {
		filledDeltas = append(filledDeltas, d2...)
	}
	return filledDeltas
}

type VwapData struct {
	Price    math.LegacyDec
	Quantity math.LegacyDec
}

func NewVwapData() *VwapData {
	return &VwapData{
		Price:    math.LegacyZeroDec(),
		Quantity: math.LegacyZeroDec(),
	}
}

func (p *VwapData) ApplyExecution(price, quantity math.LegacyDec) *VwapData {
	if p == nil {
		p = NewVwapData()
	}

	if price.IsNil() || quantity.IsNil() || quantity.IsZero() {
		return p
	}

	newQuantity := p.Quantity.Add(quantity)
	newPrice := p.Price.Mul(p.Quantity).Add(price.Mul(quantity)).Quo(newQuantity)

	return &VwapData{
		Price:    newPrice,
		Quantity: newQuantity,
	}
}

type VwapInfo struct {
	MarkPrice *math.LegacyDec
	VwapData  *VwapData
}

func NewVwapInfo(markPrice *math.LegacyDec) *VwapInfo {
	return &VwapInfo{
		MarkPrice: markPrice,
		VwapData:  NewVwapData(),
	}
}

type DerivativeVwapInfo struct {
	PerpetualVwapInfo     map[common.Hash]*VwapInfo
	ExpiryVwapInfo        map[common.Hash]*VwapInfo
	BinaryOptionsVwapInfo map[common.Hash]*VwapInfo
}

func NewDerivativeVwapInfo() DerivativeVwapInfo {
	return DerivativeVwapInfo{
		PerpetualVwapInfo:     make(map[common.Hash]*VwapInfo),
		ExpiryVwapInfo:        make(map[common.Hash]*VwapInfo),
		BinaryOptionsVwapInfo: make(map[common.Hash]*VwapInfo),
	}
}

func (p *DerivativeVwapInfo) ApplyVwap(marketID common.Hash, markPrice *math.LegacyDec, vwapData *VwapData, marketType types.MarketType) {
	var vwapInfo *VwapInfo

	switch marketType {
	case types.MarketType_Perpetual:
		vwapInfo = p.PerpetualVwapInfo[marketID]
		if vwapInfo == nil {
			vwapInfo = NewVwapInfo(markPrice)
			p.PerpetualVwapInfo[marketID] = vwapInfo
		}
	case types.MarketType_Expiry:
		vwapInfo = p.ExpiryVwapInfo[marketID]
		if vwapInfo == nil {
			vwapInfo = NewVwapInfo(markPrice)
			p.ExpiryVwapInfo[marketID] = vwapInfo
		}
	case types.MarketType_BinaryOption:
		vwapInfo = p.BinaryOptionsVwapInfo[marketID]
		if vwapInfo == nil {
			vwapInfo = NewVwapInfo(markPrice)
			p.BinaryOptionsVwapInfo[marketID] = vwapInfo
		}
	default:
	}

	if !vwapData.Quantity.IsZero() {
		vwapInfo.VwapData = vwapInfo.VwapData.ApplyExecution(vwapData.Price, vwapData.Quantity)
	}
}

func (p *DerivativeVwapInfo) GetSortedPerpetualMarketIDs() []common.Hash {
	perpetualMarketIDs := make([]common.Hash, 0)
	for k := range p.PerpetualVwapInfo {
		perpetualMarketIDs = append(perpetualMarketIDs, k)
	}

	sort.SliceStable(perpetualMarketIDs, func(i, j int) bool {
		return bytes.Compare(perpetualMarketIDs[i].Bytes(), perpetualMarketIDs[j].Bytes()) < 0
	})
	return perpetualMarketIDs
}

func (p *DerivativeVwapInfo) GetSortedExpiryFutureMarketIDs() []common.Hash {
	expiryFutureMarketIDs := make([]common.Hash, 0)
	for k := range p.ExpiryVwapInfo {
		expiryFutureMarketIDs = append(expiryFutureMarketIDs, k)
	}

	sort.SliceStable(expiryFutureMarketIDs, func(i, j int) bool {
		return bytes.Compare(expiryFutureMarketIDs[i].Bytes(), expiryFutureMarketIDs[j].Bytes()) < 0
	})
	return expiryFutureMarketIDs
}

func (p *DerivativeVwapInfo) GetSortedBinaryOptionsMarketIDs() []common.Hash {
	binaryOptionsMarketIDs := make([]common.Hash, 0)
	for k := range p.BinaryOptionsVwapInfo {
		binaryOptionsMarketIDs = append(binaryOptionsMarketIDs, k)
	}

	sort.SliceStable(binaryOptionsMarketIDs, func(i, j int) bool {
		return bytes.Compare(binaryOptionsMarketIDs[i].Bytes(), binaryOptionsMarketIDs[j].Bytes()) < 0
	})
	return binaryOptionsMarketIDs
}

// ComputeSyntheticVwapUnitDelta returns (price - markPrice) / markPrice
func (p *DerivativeVwapInfo) ComputeSyntheticVwapUnitDelta(marketID common.Hash) math.LegacyDec {
	vwapInfo := p.PerpetualVwapInfo[marketID]
	return vwapInfo.VwapData.Price.Sub(*vwapInfo.MarkPrice).Quo(*vwapInfo.MarkPrice)
}

// MergeAtomicPerpetualVwap merges accumulated atomic order VWAP data into this DerivativeVwapInfo.
func (p *DerivativeVwapInfo) MergeAtomicPerpetualVwap(atomicVwapData map[common.Hash]*VwapInfo) {
	for marketID, atomicInfo := range atomicVwapData {
		if atomicInfo == nil || atomicInfo.MarkPrice == nil || atomicInfo.VwapData == nil {
			continue
		}

		if atomicInfo.VwapData.Quantity.IsZero() {
			continue
		}

		existingInfo := p.PerpetualVwapInfo[marketID]
		if existingInfo == nil {
			// No existing VWAP for this market, just use the atomic data
			p.PerpetualVwapInfo[marketID] = atomicInfo
			continue
		}

		// Merge the VWAP data: newVwap = (existingPrice * existingQty + atomicPrice * atomicQty) / (existingQty + atomicQty)
		existingInfo.VwapData = existingInfo.VwapData.ApplyExecution(atomicInfo.VwapData.Price, atomicInfo.VwapData.Quantity)
	}
}

type PositionState struct {
	Position *Position
}

func NewPositionStates() map[common.Hash]*PositionState {
	return make(map[common.Hash]*PositionState)
}

// ApplyFundingAndGetUpdatedPositionState updates the position to account for any funding payment and returns a PositionState.
func ApplyFundingAndGetUpdatedPositionState(p *Position, funding *PerpetualMarketFunding) *PositionState {
	p.ApplyFunding(funding)
	positionState := &PositionState{
		Position: p,
	}
	return positionState
}

func GetSortedSubaccountKeys(p map[common.Hash]*PositionState) []common.Hash {
	subaccountKeys := make([]common.Hash, 0)
	for k := range p {
		subaccountKeys = append(subaccountKeys, k)
	}
	sort.SliceStable(subaccountKeys, func(i, j int) bool {
		return bytes.Compare(subaccountKeys[i].Bytes(), subaccountKeys[j].Bytes()) < 0
	})
	return subaccountKeys
}

func GetPositionSliceData(p map[common.Hash]*PositionState) ([]*Position, []common.Hash) {
	positionSubaccountIDs := GetSortedSubaccountKeys(p)
	positions := make([]*Position, 0, len(positionSubaccountIDs))

	nonNilPositionSubaccountIDs := make([]common.Hash, 0)
	for idx := range positionSubaccountIDs {
		subaccountID := positionSubaccountIDs[idx]
		position := p[subaccountID]
		if position.Position != nil {
			positions = append(positions, position.Position)
			nonNilPositionSubaccountIDs = append(nonNilPositionSubaccountIDs, subaccountID)
		}

		// else {
		// 	fmt.Println("❌ position is nil for subaccount", subaccountID.Hex())
		// }
	}

	return positions, nonNilPositionSubaccountIDs
}

type DerivativeBatchExecutionData struct {
	Market DerivativeMarketI

	MarkPrice math.LegacyDec
	Funding   *PerpetualMarketFunding

	// update the deposits for margin deductions, payouts and refunds
	DepositDeltas        types.DepositDeltas
	DepositSubaccountIDs []common.Hash

	TradingRewards types.TradingRewardPoints

	// updated positions
	Positions             []*Position
	MarketBalanceDelta    math.LegacyDec
	PositionSubaccountIDs []common.Hash
	OpenInterestDelta     math.LegacyDec

	// resting limit order filled deltas to apply
	TransientLimitOrderFilledDeltas []*DerivativeLimitOrderDelta
	// resting limit order filled deltas to apply
	RestingLimitOrderFilledDeltas []*DerivativeLimitOrderDelta
	// transient limit order cancelled deltas to apply
	TransientLimitOrderCancelledDeltas []*DerivativeLimitOrderDelta
	// resting limit order cancelled deltas to apply
	RestingLimitOrderCancelledDeltas []*DerivativeLimitOrderDelta

	// events for batch market order and limit order execution
	MarketBuyOrderExecutionEvent          *EventBatchDerivativeExecution
	MarketSellOrderExecutionEvent         *EventBatchDerivativeExecution
	RestingLimitBuyOrderExecutionEvent    *EventBatchDerivativeExecution
	RestingLimitSellOrderExecutionEvent   *EventBatchDerivativeExecution
	TransientLimitBuyOrderExecutionEvent  *EventBatchDerivativeExecution
	TransientLimitSellOrderExecutionEvent *EventBatchDerivativeExecution

	// event for new orders to add to the orderbook
	NewOrdersEvent          *EventNewDerivativeOrders
	CancelLimitOrderEvents  []*EventCancelDerivativeOrder
	CancelMarketOrderEvents []*EventCancelDerivativeOrder

	VwapData *VwapData
}

func (d *DerivativeBatchExecutionData) GetAtomicDerivativeMarketOrderResults() *DerivativeMarketOrderResults {
	var trade *DerivativeTradeLog

	switch {
	case d.MarketBuyOrderExecutionEvent != nil:
		trade = d.MarketBuyOrderExecutionEvent.Trades[0]
	case d.MarketSellOrderExecutionEvent != nil:
		trade = d.MarketSellOrderExecutionEvent.Trades[0]
	default:
		return EmptyDerivativeMarketOrderResults()
	}

	return &DerivativeMarketOrderResults{
		Quantity:      trade.PositionDelta.ExecutionQuantity,
		Price:         trade.PositionDelta.ExecutionPrice,
		Fee:           trade.Fee,
		PositionDelta: *trade.PositionDelta,
		Payout:        trade.Payout,
	}
}

func ApplyDeltasAndGetDerivativeOrderBatchEvent(
	isBuy bool,
	executionType ExecutionType,
	market DerivativeMarketI,
	funding *PerpetualMarketFunding,
	stateExpansions []*DerivativeOrderStateExpansion,
	depositDeltas types.DepositDeltas,
	tradingRewardPoints types.TradingRewardPoints,
	isLiquidation bool,
) (batch *EventBatchDerivativeExecution, filledDeltas []*DerivativeLimitOrderDelta) {
	if len(stateExpansions) == 0 {
		return
	}

	trades := make([]*DerivativeTradeLog, 0, len(stateExpansions))

	if !executionType.IsMarket() {
		filledDeltas = make([]*DerivativeLimitOrderDelta, 0, len(stateExpansions))
	}

	for idx := range stateExpansions {
		expansion := stateExpansions[idx]

		feeRecipientSubaccount := types.EthAddressToSubaccountID(expansion.FeeRecipient)
		if bytes.Equal(feeRecipientSubaccount.Bytes(), common.Hash{}.Bytes()) {
			feeRecipientSubaccount = types.AuctionSubaccountID
		}

		depositDeltas.ApplyDepositDelta(expansion.SubaccountID, &types.DepositDelta{
			TotalBalanceDelta:     market.NotionalToChainFormat(expansion.TotalBalanceDelta),
			AvailableBalanceDelta: market.NotionalToChainFormat(expansion.AvailableBalanceDelta),
		})
		chainFormatFeeRecipientReward := market.NotionalToChainFormat(expansion.FeeRecipientReward)
		chainFormatAuctionFeeReward := market.NotionalToChainFormat(expansion.AuctionFeeReward)
		depositDeltas.ApplyUniformDelta(feeRecipientSubaccount, chainFormatFeeRecipientReward)
		depositDeltas.ApplyUniformDelta(types.AuctionSubaccountID, chainFormatAuctionFeeReward)

		sender := types.SubaccountIDToSdkAddress(expansion.SubaccountID)
		tradingRewardPoints.AddPointsForAddress(sender.String(), expansion.TradingRewardPoints)

		if !executionType.IsMarket() {
			filledDeltas = append(filledDeltas, expansion.LimitOrderFilledDelta)
		}

		var realizedTradeFee math.LegacyDec

		isSelfRelayedTrade := expansion.FeeRecipient == types.SubaccountIDToEthAddress(expansion.SubaccountID)
		if isSelfRelayedTrade {
			// if negative fee, equals the full negative rebate
			// otherwise equals the fees going to auction
			realizedTradeFee = expansion.AuctionFeeReward
		} else {
			realizedTradeFee = expansion.FeeRecipientReward.Add(expansion.AuctionFeeReward)
		}

		if expansion.PositionDelta != nil {
			tradeLog := &DerivativeTradeLog{
				SubaccountId:        expansion.SubaccountID.Bytes(),
				PositionDelta:       expansion.PositionDelta,
				Payout:              expansion.Payout,
				Fee:                 realizedTradeFee,
				OrderHash:           expansion.OrderHash.Bytes(),
				FeeRecipientAddress: expansion.FeeRecipient.Bytes(),
				Cid:                 expansion.Cid,
				Pnl:                 expansion.Pnl,
			}
			trades = append(trades, tradeLog)
		}
	}

	if len(trades) == 0 {
		return nil, filledDeltas
	}

	batch = &EventBatchDerivativeExecution{
		MarketId:      market.MarketID().String(),
		IsBuy:         isBuy,
		IsLiquidation: isLiquidation,
		ExecutionType: executionType,
		Trades:        trades,
	}
	if funding != nil {
		batch.CumulativeFunding = &funding.CumulativeFunding
	}

	return batch, filledDeltas
}

type DerivativeOrderStateExpansion struct {
	SubaccountID           common.Hash
	PositionDelta          *PositionDelta
	Payout                 math.LegacyDec
	Pnl                    math.LegacyDec
	MarketBalanceDelta     math.LegacyDec
	OpenInterestDelta      math.LegacyDec
	TotalBalanceDelta      math.LegacyDec
	AvailableBalanceDelta  math.LegacyDec
	AuctionFeeReward       math.LegacyDec
	TradingRewardPoints    math.LegacyDec
	FeeRecipientReward     math.LegacyDec
	FeeRecipient           common.Address
	LimitOrderFilledDelta  *DerivativeLimitOrderDelta
	MarketOrderFilledDelta *DerivativeMarketOrderDelta
	OrderHash              common.Hash
	Cid                    string
}

type DeficitPositions struct {
	DerivativePosition *DerivativePosition
	DeficitAmountAbs   math.LegacyDec
}

type SocializedLossData struct {
	PositionsReceivingHaircut []*Position
	DeficitPositions          []DeficitPositions
	DeficitAmountAbs          math.LegacyDec
	SurplusAmount             math.LegacyDec
	TotalProfits              math.LegacyDec
	TotalPositivePayouts      math.LegacyDec
}
