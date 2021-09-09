package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type DerivativeBatchExecutionData struct {
	Market *DerivativeMarket

	MarkPrice sdk.Dec
	Funding   *PerpetualMarketFunding

	// update the deposits for margin deductions, payouts and refunds
	DepositDeltas        DepositDeltas
	DepositSubaccountIDs []common.Hash

	LiquidityMiningRewards         LiquidityMiningRewards

	// updated positions
	Positions             []*Position
	PositionSubaccountIDs []common.Hash

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

type DerivativeMatchingExpansionData struct {
	TransientLimitBuyExpansions    []*DerivativeOrderStateExpansion
	TransientLimitSellExpansions   []*DerivativeOrderStateExpansion
	RestingLimitBuyExpansions      []*DerivativeOrderStateExpansion
	RestingLimitSellExpansions     []*DerivativeOrderStateExpansion
	RestingLimitBuyOrderCancels    []*DerivativeLimitOrder
	RestingLimitSellOrderCancels   []*DerivativeLimitOrder
	TransientLimitBuyOrderCancels  []*DerivativeLimitOrder
	TransientLimitSellOrderCancels []*DerivativeLimitOrder
	ClearingPrice                  sdk.Dec
	ClearingQuantity               sdk.Dec
	NewRestingLimitBuyOrders       []*DerivativeLimitOrder // transient buy orders that become new resting limit orders
	NewRestingLimitSellOrders      []*DerivativeLimitOrder // transient sell orders that become new resting limit orders
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
	MarketBuyClearingPrice       sdk.Dec
	MarketSellClearingPrice      sdk.Dec
	MarketBuyClearingQuantity    sdk.Dec
	MarketSellClearingQuantity   sdk.Dec
}

func (d *DerivativeMarketOrderExpansionData) SetExecutionData(
	isMarketBuy bool,
	marketOrderClearingPrice, marketOrderClearingQuantity sdk.Dec,
	restingLimitOrderCancels []*DerivativeLimitOrder,
	marketOrderStateExpansions,
	restingLimitOrderStateExpansions []*DerivativeOrderStateExpansion,
	marketOrderCancels []*DerivativeMarketOrderCancel,
) {
	if isMarketBuy {
		d.MarketBuyClearingPrice = marketOrderClearingPrice
		d.MarketBuyClearingQuantity = marketOrderClearingQuantity
		d.RestingLimitSellOrderCancels = restingLimitOrderCancels
		d.MarketBuyExpansions = marketOrderStateExpansions
		d.LimitSellExpansions = restingLimitOrderStateExpansions
		d.MarketBuyOrderCancels = marketOrderCancels
	} else {
		d.MarketSellClearingPrice = marketOrderClearingPrice
		d.MarketSellClearingQuantity = marketOrderClearingQuantity
		d.RestingLimitBuyOrderCancels = restingLimitOrderCancels
		d.MarketSellExpansions = marketOrderStateExpansions
		d.LimitBuyExpansions = restingLimitOrderStateExpansions
		d.MarketSellOrderCancels = marketOrderCancels
	}
}

func (e *DerivativeMatchingExpansionData) GetLimitMatchingDerivativeBatchExecutionData(
	market *DerivativeMarket,
	markPrice sdk.Dec,
	funding *PerpetualMarketFunding,
	positionStates map[common.Hash]*PositionState,
) *DerivativeBatchExecutionData {
	depositDeltas := NewDepositDeltas()
	liquidityMiningRewards := NewLiquidityMiningRewards()

	// process undermargined resting limit order forced cancellations
	cancelLimitOrdersEvents, restingOrderCancelledDeltas, transientOrderCancelledDeltas := e.applyCancellationsAndGetDerivativeLimitCancelEvents(market.MarketID(), market.MakerFeeRate, market.TakerFeeRate, depositDeltas)

	positions, positionSubaccountIDs := GetPositionSliceData(positionStates)

	transientLimitBuyOrderBatchEvent, transientLimitBuyFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(true, ExecutionType_LimitMatchNewOrder, market, funding, e.TransientLimitBuyExpansions, depositDeltas, liquidityMiningRewards)
	restingLimitBuyOrderBatchEvent, restingLimitBuyFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(true, ExecutionType_LimitMatchRestingOrder, market, funding, e.RestingLimitBuyExpansions, depositDeltas, liquidityMiningRewards)
	transientLimitSellOrderBatchEvent, transientLimitSellFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(false, ExecutionType_LimitMatchNewOrder, market, funding, e.TransientLimitSellExpansions, depositDeltas, liquidityMiningRewards)
	restingLimitSellOrderBatchEvent, restingLimitSellFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(false, ExecutionType_LimitMatchRestingOrder, market, funding, e.RestingLimitSellExpansions, depositDeltas, liquidityMiningRewards)

	restingOrderFilledDeltas := mergeDerivativeLimitOrderFilledDeltas(restingLimitBuyFilledDeltas, restingLimitSellFilledDeltas)
	transientOrderFilledDeltas := mergeDerivativeLimitOrderFilledDeltas(transientLimitBuyFilledDeltas, transientLimitSellFilledDeltas)

	// sort keys since map iteration is non-deterministic
	depositDeltaKeys := depositDeltas.GetSortedSubaccountKeys()

	var vwapData *VwapData
	if market.IsPerpetual {
		vwapData = vwapData.ApplyExecution(e.ClearingPrice, e.ClearingQuantity)
	}

	var newOrdersEvent *EventNewDerivativeOrders
	if len(e.NewRestingLimitBuyOrders) > 0 || len(e.NewRestingLimitSellOrders) > 0 {
		newOrdersEvent = &EventNewDerivativeOrders{
			MarketId:   market.MarketId,
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
		LiquidityMiningRewards:                liquidityMiningRewards,
		Positions:                             positions,
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

func (e *DerivativeMarketOrderExpansionData) getDerivativeMarketCancelEvents(
	marketID common.Hash,
) []*EventCancelDerivativeOrder {
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

func applyDerivativeLimitCancellation(
	order *DerivativeLimitOrder,
	orderFeeRate sdk.Dec,
	depositDeltas DepositDeltas,
) {
	// For vanilla orders, increment the available balance
	if order.IsVanilla() {
		depositDelta := order.GetCancelDepositDelta(orderFeeRate)
		depositDeltas.ApplyDepositDelta(order.SubaccountID(), depositDelta)
	}
}

func (e *DerivativeMatchingExpansionData) applyCancellationsAndGetDerivativeLimitCancelEvents(
	marketID common.Hash,
	makerFeeRate sdk.Dec,
	takerFeeRate sdk.Dec,
	depositDeltas DepositDeltas,
) (
	cancelOrdersEvent []*EventCancelDerivativeOrder,
	restingOrderCancelledDeltas []*DerivativeLimitOrderDelta,
	transientOrderCancelledDeltas []*DerivativeLimitOrderDelta,
) {
	marketIDHex := marketID.Hex()

	cancelOrdersEvent = make([]*EventCancelDerivativeOrder, 0, len(e.RestingLimitBuyOrderCancels)+len(e.RestingLimitSellOrderCancels)+len(e.TransientLimitBuyOrderCancels)+len(e.TransientLimitSellOrderCancels))
	restingOrderCancelledDeltas = make([]*DerivativeLimitOrderDelta, 0, len(e.RestingLimitBuyOrderCancels)+len(e.RestingLimitSellOrderCancels))
	transientOrderCancelledDeltas = make([]*DerivativeLimitOrderDelta, 0, len(e.TransientLimitBuyOrderCancels)+len(e.TransientLimitSellOrderCancels))

	for idx := range e.RestingLimitBuyOrderCancels {
		order := e.RestingLimitBuyOrderCancels[idx]

		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
		restingOrderCancelledDeltas = append(restingOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   sdk.ZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	for idx := range e.RestingLimitSellOrderCancels {
		order := e.RestingLimitSellOrderCancels[idx]

		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
		restingOrderCancelledDeltas = append(restingOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   sdk.ZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	for idx := range e.TransientLimitBuyOrderCancels {
		order := e.TransientLimitBuyOrderCancels[idx]

		applyDerivativeLimitCancellation(order, takerFeeRate, depositDeltas)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
		transientOrderCancelledDeltas = append(transientOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   sdk.ZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	for idx := range e.TransientLimitSellOrderCancels {
		order := e.TransientLimitSellOrderCancels[idx]
		applyDerivativeLimitCancellation(order, takerFeeRate, depositDeltas)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
		transientOrderCancelledDeltas = append(transientOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   sdk.ZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	return cancelOrdersEvent, restingOrderCancelledDeltas, transientOrderCancelledDeltas
}

func (e *DerivativeMarketOrderExpansionData) applyCancellationsAndGetDerivativeLimitCancelEvents(
	marketID common.Hash,
	makerFeeRate sdk.Dec,
	depositDeltas DepositDeltas,
) (
	cancelOrdersEvent []*EventCancelDerivativeOrder,
	restingOrderCancelledDeltas []*DerivativeLimitOrderDelta,
) {
	marketIDHex := marketID.Hex()
	cancelOrdersEvent = make([]*EventCancelDerivativeOrder, 0, len(e.RestingLimitBuyOrderCancels)+len(e.RestingLimitSellOrderCancels))

	restingOrderCancelledDeltas = make([]*DerivativeLimitOrderDelta, 0, len(e.RestingLimitBuyOrderCancels)+len(e.RestingLimitSellOrderCancels))

	for idx := range e.RestingLimitBuyOrderCancels {
		order := e.RestingLimitBuyOrderCancels[idx]
		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})

		restingOrderCancelledDeltas = append(restingOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   sdk.ZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}

	for idx := range e.RestingLimitSellOrderCancels {
		order := e.RestingLimitSellOrderCancels[idx]
		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})

		restingOrderCancelledDeltas = append(restingOrderCancelledDeltas, &DerivativeLimitOrderDelta{
			Order:          order,
			FillQuantity:   sdk.ZeroDec(),
			CancelQuantity: order.Fillable,
		})
	}
	return cancelOrdersEvent, restingOrderCancelledDeltas
}

func (e *DerivativeMarketOrderExpansionData) GetMarketDerivativeBatchExecutionData(
	market *DerivativeMarket,
	markPrice sdk.Dec,
	funding *PerpetualMarketFunding,
	positionStates map[common.Hash]*PositionState,
) *DerivativeBatchExecutionData {
	depositDeltas := NewDepositDeltas()
	liquidityMiningRewards := NewLiquidityMiningRewards()

	// process undermargined limit order forced cancellations
	cancelLimitOrdersEvents, restingOrderCancelledDeltas := e.applyCancellationsAndGetDerivativeLimitCancelEvents(market.MarketID(), market.MakerFeeRate, depositDeltas)

	// process unfilled market order cancellations
	cancelMarketOrdersEvents := e.getDerivativeMarketCancelEvents(market.MarketID())

	positions, positionSubaccountIDs := GetPositionSliceData(positionStates)

	buyMarketOrderBatchEvent, _ := ApplyDeltasAndGetDerivativeOrderBatchEvent(true, ExecutionType_Market, market, funding, e.MarketBuyExpansions, depositDeltas, liquidityMiningRewards)
	sellMarketOrderBatchEvent, _ := ApplyDeltasAndGetDerivativeOrderBatchEvent(false, ExecutionType_Market, market, funding, e.MarketSellExpansions, depositDeltas, liquidityMiningRewards)
	restingLimitBuyOrderBatchEvent, limitBuyFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(true, ExecutionType_LimitFill, market, funding, e.LimitBuyExpansions, depositDeltas, liquidityMiningRewards)
	restingLimitSellOrderBatchEvent, limitSellFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(false, ExecutionType_LimitFill, market, funding, e.LimitSellExpansions, depositDeltas, liquidityMiningRewards)

	filledDeltas := mergeDerivativeLimitOrderFilledDeltas(limitBuyFilledDeltas, limitSellFilledDeltas)

	// sort keys since map iteration is non-deterministic
	depositDeltaKeys := depositDeltas.GetSortedSubaccountKeys()

	var vwapData *VwapData
	if market.IsPerpetual {
		vwapData = vwapData.ApplyExecution(e.MarketBuyClearingPrice, e.MarketBuyClearingQuantity)
		vwapData = vwapData.ApplyExecution(e.MarketSellClearingPrice, e.MarketSellClearingQuantity)
	}

	// Final Step: Store the DerivativeBatchExecutionData for future reduction/processing
	batch := &DerivativeBatchExecutionData{
		Market:                                market,
		MarkPrice:                             markPrice,
		Funding:                               funding,
		DepositDeltas:                         depositDeltas,
		DepositSubaccountIDs:                  depositDeltaKeys,
		LiquidityMiningRewards:                liquidityMiningRewards,
		Positions:                             positions,
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
		CancelLimitOrderEvents:                cancelLimitOrdersEvents,
		CancelMarketOrderEvents:               cancelMarketOrdersEvents,
		VwapData:                              vwapData,
	}
	return batch
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
