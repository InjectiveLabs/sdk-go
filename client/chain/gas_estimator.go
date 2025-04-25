package chain

import (
	"math"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// refer to python sdk: https://github.com/InjectiveLabs/sdk-python/blob/master/pyinjective/core/gas_limit_estimator.py#L15
const (
	GENERAL_MESSAGE_GAS_LIMIT = 25_000
	BASIC_REFERENCE_GAS_LIMIT = 150_000

	SPOT_ORDER_CREATION_GAS_LIMIT          = 52_000
	DERIVATIVE_ORDER_CREATION_GAS_LIMIT    = 84_000
	SPOT_ORDER_CANCELATION_GAS_LIMIT       = 50_000
	DERIVATIVE_ORDER_CANCELATION_GAS_LIMIT = 68_000
	// POST ONLY orders take around 50% more gas to create than normal orders due to the required validations
	SPOT_POST_ONLY_ORDER_MULTIPLIER       = 0.62
	DERIVATIVE_POST_ONLY_ORDER_MULTIPLIER = 0.35

	// for batch update
	CANCEL_ALL_SPOT_MARKET_GAS_LIMIT       = 40_000
	CANCEL_ALL_DERIVATIVE_MARKET_GAS_LIMIT = 50_000
	MESSAGE_GAS_LIMIT                      = 30_000

	AVERAGE_CANCEL_ALL_AFFECTED_ORDERS = 20

	TRANSACTION_GAS_LIMIT = 60_000

	TRANSACTION_PRESERVED_MULTIPLIER = 0.2
)

type msgGasEstimator interface {
	appliesTo(msg sdk.Msg) bool
	estimateMsgGas(msgs sdk.Msg) uint64
}

type TXGasEstimator struct {
	msg_gas_estimators []msgGasEstimator
}

func newTXGasEstimator() *TXGasEstimator {
	return &TXGasEstimator{
		msg_gas_estimators: []msgGasEstimator{
			newMsgCreateSpotLimitOrderGasEstimator(),
			newMsgCreateDerivativeLimitOrderGasEstimator(),
			newMsgBatchCreateSpotLimitOrdersGasEstimator(),
			newMsgBatchCreateDerivativeLimitOrdersGasEstimator(),
			newMsgBatchCancelSpotOrdersGasEstimator(),
			newMsgBatchCancelDerivativeOrdersGasEstimator(),
			newMsgCancelSpotOrderGasEstimator(),
			newMsgCancelDerivativeOrderGasEstimator(),
			newMsgBatchUpdateOrdersGasEstimator(),
		},
	}
}

func (g *TXGasEstimator) EstimateTXGas(msgs ...sdk.Msg) uint64 {

	// add up all types of messages' gas
	// messages_gas_limit = Decimal("0")
	// for message in messages:
	//    estimator = GasLimitEstimator.for_message(message=message)
	//    messages_gas_limit += estimator.gas_limit()
	// messages_gas_limit = math.ceil(messages_gas_limit)

	// add the transaction gas limit
	// gas_wanted = messages_gas_limit + self.TRANSACTION_GAS_LIMIT
	// transaction.with_gas(gas=gas_wanted)

	// msg types:
	// - spot order creation, with spot post only order multiplier
	// - derivative order creation, with derivative post only order multiplier
	// - spot order cancellation
	// - derivative order cancellation

	var totalMsgGas uint64
	for i := range msgs {
		msg := msgs[i]
		totalMsgGas += g.estimateMsgGas(msg)
	}
	txGas := totalMsgGas + TRANSACTION_GAS_LIMIT
	txGas += uint64(math.Ceil(float64(txGas) * TRANSACTION_PRESERVED_MULTIPLIER))

	// for i := range msgs {
	// 	msg := msgs[i]
	// 	fmt.Printf("TypeURL: %s\n", proto.MessageName(msg))
	// }
	return txGas
}

func (g *TXGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// use chain responsibility pattern
	for i := range g.msg_gas_estimators {
		estimator := g.msg_gas_estimators[i]
		if estimator.appliesTo(msg) {
			return estimator.estimateMsgGas(msg)
		}
	}
	return GENERAL_MESSAGE_GAS_LIMIT
}

func toPointer[T any](slice []T) []*T {
	ptrs := make([]*T, len(slice))
	for i := range slice {
		ptrs[i] = &slice[i]
	}
	return ptrs
}

// [T interface { GetOrderType() exchangetypes.OrderType}]
// [T exchangetypes.SpotOrder | exchangetypes.DerivativeOrder]
func selectPostOnlyOrders[T interface {
	GetOrderType() exchangetypes.OrderType
}](orders []T) []T {
	var filtered []T
	for _, order := range orders {
		if order.GetOrderType() == exchangetypes.OrderType_BUY_PO || order.GetOrderType() == exchangetypes.OrderType_SELL_PO {
			filtered = append(filtered, order)
		}
	}
	return filtered
}

// MsgCreate[Spot/Derivative]LimitOrder
type MsgCreateSpotLimitOrderGasEstimator struct {
}

func newMsgCreateSpotLimitOrderGasEstimator() *MsgCreateSpotLimitOrderGasEstimator {
	return &MsgCreateSpotLimitOrderGasEstimator{}
}

func (e *MsgCreateSpotLimitOrderGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgCreateSpotLimitOrder)
	return ok
}

func (e *MsgCreateSpotLimitOrderGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// postOnlyOrder := selectPostOnlyOrders([]exchangetypes.SpotOrder{msg.(*exchangetypes.MsgCreateSpotLimitOrder).Order})
	postOnlyOrder := selectPostOnlyOrders([]*exchangetypes.SpotOrder{&msg.(*exchangetypes.MsgCreateSpotLimitOrder).Order})
	total := uint64(SPOT_ORDER_CREATION_GAS_LIMIT)
	if len(postOnlyOrder) > 0 {
		total += uint64(math.Ceil(SPOT_ORDER_CREATION_GAS_LIMIT * SPOT_POST_ONLY_ORDER_MULTIPLIER))
	}
	return total
}

type MsgCreateDerivativeLimitOrderGasEstimator struct {
}

func newMsgCreateDerivativeLimitOrderGasEstimator() *MsgCreateDerivativeLimitOrderGasEstimator {
	return &MsgCreateDerivativeLimitOrderGasEstimator{}
}

func (e *MsgCreateDerivativeLimitOrderGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgCreateDerivativeLimitOrder)
	return ok
}

func (e *MsgCreateDerivativeLimitOrderGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// type assert exchangetypes.MsgCreateDerivativeLimitOrder
	postOnlyOrder := selectPostOnlyOrders([]*exchangetypes.DerivativeOrder{&msg.(*exchangetypes.MsgCreateDerivativeLimitOrder).Order})
	total := uint64(DERIVATIVE_ORDER_CREATION_GAS_LIMIT)
	if len(postOnlyOrder) > 0 {
		total += uint64(math.Ceil(DERIVATIVE_ORDER_CREATION_GAS_LIMIT * DERIVATIVE_POST_ONLY_ORDER_MULTIPLIER))
	}
	return total
}

// MsgBatchCreate[Spot/Derivative]LimitOrders

type MsgBatchCreateSpotLimitOrdersGasEstimator struct {
}

func newMsgBatchCreateSpotLimitOrdersGasEstimator() *MsgBatchCreateSpotLimitOrdersGasEstimator {
	return &MsgBatchCreateSpotLimitOrdersGasEstimator{}
}

func (e *MsgBatchCreateSpotLimitOrdersGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgBatchCreateSpotLimitOrders)
	return ok
}

func (e *MsgBatchCreateSpotLimitOrdersGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// type assert exchangetypes.MsgBatchCreateSpotLimitOrders
	batchCreateMsg := msg.(*exchangetypes.MsgBatchCreateSpotLimitOrders)
	var total uint64
	postOnlyOrders := selectPostOnlyOrders(toPointer(batchCreateMsg.Orders))
	total += GENERAL_MESSAGE_GAS_LIMIT
	total += uint64(len(batchCreateMsg.Orders)) * SPOT_ORDER_CREATION_GAS_LIMIT
	total += uint64(math.Ceil(float64(len(postOnlyOrders)) * float64(SPOT_ORDER_CREATION_GAS_LIMIT) * SPOT_POST_ONLY_ORDER_MULTIPLIER))
	return total
}

type MsgBatchCreateDerivativeLimitOrdersGasEstimator struct {
}

func newMsgBatchCreateDerivativeLimitOrdersGasEstimator() *MsgBatchCreateDerivativeLimitOrdersGasEstimator {
	return &MsgBatchCreateDerivativeLimitOrdersGasEstimator{}
}

func (e *MsgBatchCreateDerivativeLimitOrdersGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgBatchCreateDerivativeLimitOrders)
	return ok
}

func (e *MsgBatchCreateDerivativeLimitOrdersGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// type assert exchangetypes.MsgBatchCreateDerivativeLimitOrders
	batchCreateMsg := msg.(*exchangetypes.MsgBatchCreateDerivativeLimitOrders)
	var total uint64
	postOnlyOrders := selectPostOnlyOrders(toPointer(batchCreateMsg.Orders))
	total += GENERAL_MESSAGE_GAS_LIMIT
	total += uint64(len(batchCreateMsg.Orders)) * DERIVATIVE_ORDER_CREATION_GAS_LIMIT
	total += uint64(math.Ceil(float64(len(postOnlyOrders)) * float64(DERIVATIVE_ORDER_CREATION_GAS_LIMIT) * DERIVATIVE_POST_ONLY_ORDER_MULTIPLIER))
	return total
}

// MsgCancel[Spot/Derivative]OrderGasEstimator
type MsgCancelSpotOrderGasEstimator struct {
}

func newMsgCancelSpotOrderGasEstimator() *MsgCancelSpotOrderGasEstimator {
	return &MsgCancelSpotOrderGasEstimator{}
}

func (e *MsgCancelSpotOrderGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgCancelSpotOrder)
	return ok
}

func (e *MsgCancelSpotOrderGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	return SPOT_ORDER_CANCELATION_GAS_LIMIT
}

type MsgCancelDerivativeOrderGasEstimator struct {
}

func newMsgCancelDerivativeOrderGasEstimator() *MsgCancelDerivativeOrderGasEstimator {
	return &MsgCancelDerivativeOrderGasEstimator{}
}

func (e *MsgCancelDerivativeOrderGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgCancelDerivativeOrder)
	return ok
}

func (e *MsgCancelDerivativeOrderGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	return DERIVATIVE_ORDER_CANCELATION_GAS_LIMIT
}

// MsgBatchCancel[Spot/Derivative]OrdersGasEstimator
type MsgBatchCancelSpotOrdersGasEstimator struct {
}

func newMsgBatchCancelSpotOrdersGasEstimator() *MsgBatchCancelSpotOrdersGasEstimator {
	return &MsgBatchCancelSpotOrdersGasEstimator{}
}

func (e *MsgBatchCancelSpotOrdersGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgBatchCancelSpotOrders)
	return ok
}

func (e *MsgBatchCancelSpotOrdersGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// type assert exchangetypes.MsgBatchCancelSpotOrders
	batchCancelMsg := msg.(*exchangetypes.MsgBatchCancelSpotOrders)
	var total uint64
	total += GENERAL_MESSAGE_GAS_LIMIT
	total += uint64(len(batchCancelMsg.Data)) * SPOT_ORDER_CANCELATION_GAS_LIMIT // Calculate gas based on number of orders to cancel
	return total
}

type MsgBatchCancelDerivativeOrdersGasEstimator struct {
}

func newMsgBatchCancelDerivativeOrdersGasEstimator() *MsgBatchCancelDerivativeOrdersGasEstimator {
	return &MsgBatchCancelDerivativeOrdersGasEstimator{}
}

func (e *MsgBatchCancelDerivativeOrdersGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgBatchCancelDerivativeOrders)
	return ok
}

func (e *MsgBatchCancelDerivativeOrdersGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// type assert exchangetypes.MsgBatchCancelDerivativeOrders
	batchCancelMsg := msg.(*exchangetypes.MsgBatchCancelDerivativeOrders)
	var total uint64
	total += GENERAL_MESSAGE_GAS_LIMIT
	total += uint64(len(batchCancelMsg.Data)) * DERIVATIVE_ORDER_CANCELATION_GAS_LIMIT // Calculate gas based on number of orders to cancel
	return total
}

// MsgBatchUpdateOrdersGasEstimator
type MsgBatchUpdateOrdersGasEstimator struct {
}

func newMsgBatchUpdateOrdersGasEstimator() *MsgBatchUpdateOrdersGasEstimator {
	return &MsgBatchUpdateOrdersGasEstimator{}
}

func (e *MsgBatchUpdateOrdersGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgBatchUpdateOrders)
	return ok
}

func (e *MsgBatchUpdateOrdersGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// type assert exchangetypes.MsgBatchUpdateOrders
	batchUpdateMsg := msg.(*exchangetypes.MsgBatchUpdateOrders)

	// Helper functions to count post-only orders for different order types
	postOnlySpotOrders := selectPostOnlyOrders(batchUpdateMsg.SpotOrdersToCreate)
	postOnlyDerivativeOrders := selectPostOnlyOrders(batchUpdateMsg.DerivativeOrdersToCreate)
	postOnlyBinaryOptionsOrders := selectPostOnlyOrders(batchUpdateMsg.BinaryOptionsOrdersToCreate)

	var total uint64
	total += MESSAGE_GAS_LIMIT
	total += uint64(len(batchUpdateMsg.SpotOrdersToCreate)) * SPOT_ORDER_CREATION_GAS_LIMIT
	total += uint64(len(batchUpdateMsg.DerivativeOrdersToCreate)) * DERIVATIVE_ORDER_CREATION_GAS_LIMIT
	total += uint64(len(batchUpdateMsg.BinaryOptionsOrdersToCreate)) * DERIVATIVE_ORDER_CREATION_GAS_LIMIT

	total += uint64(math.Ceil(float64(len(postOnlySpotOrders)) * float64(SPOT_ORDER_CREATION_GAS_LIMIT) * SPOT_POST_ONLY_ORDER_MULTIPLIER))
	total += uint64(math.Ceil(float64(len(postOnlyDerivativeOrders)) * float64(DERIVATIVE_ORDER_CREATION_GAS_LIMIT) * DERIVATIVE_POST_ONLY_ORDER_MULTIPLIER))
	total += uint64(math.Ceil(float64(len(postOnlyBinaryOptionsOrders)) * float64(DERIVATIVE_ORDER_CREATION_GAS_LIMIT) * DERIVATIVE_POST_ONLY_ORDER_MULTIPLIER))

	total += uint64(len(batchUpdateMsg.SpotOrdersToCancel)) * SPOT_ORDER_CANCELATION_GAS_LIMIT
	total += uint64(len(batchUpdateMsg.DerivativeOrdersToCancel)) * DERIVATIVE_ORDER_CANCELATION_GAS_LIMIT
	total += uint64(len(batchUpdateMsg.BinaryOptionsOrdersToCancel)) * DERIVATIVE_ORDER_CANCELATION_GAS_LIMIT

	total += uint64(len(batchUpdateMsg.SpotMarketIdsToCancelAll)) * CANCEL_ALL_SPOT_MARKET_GAS_LIMIT * AVERAGE_CANCEL_ALL_AFFECTED_ORDERS
	total += uint64(len(batchUpdateMsg.DerivativeMarketIdsToCancelAll)) * CANCEL_ALL_DERIVATIVE_MARKET_GAS_LIMIT * AVERAGE_CANCEL_ALL_AFFECTED_ORDERS
	total += uint64(len(batchUpdateMsg.BinaryOptionsMarketIdsToCancelAll)) * CANCEL_ALL_DERIVATIVE_MARKET_GAS_LIMIT * AVERAGE_CANCEL_ALL_AFFECTED_ORDERS

	return total
}
