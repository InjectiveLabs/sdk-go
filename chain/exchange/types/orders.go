package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

var ZeroHash = common.Hash{}

func (o *SpotOrder) GetNewSpotLimitOrder(hash common.Hash) *SpotLimitOrder {
	return &SpotLimitOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Fillable:     o.OrderInfo.Quantity,
		TriggerPrice: o.TriggerPrice,
		Hash:         hash.Bytes(),
	}
}

type SpotOrderStateExpansion struct {
	BaseChangeAmount   sdk.Dec
	QuoteChangeAmount  sdk.Dec
	QuoteRefundAmount  sdk.Dec
	FeeRecipient       common.Address
	FeeRecipientReward sdk.Dec
	AuctionFeeReward   sdk.Dec
	Hash               common.Hash
	SubaccountID       common.Hash
	// for market orders, FillableAmount refers to the fillable quantity of the market order execution (if any)
	FillableAmount sdk.Dec
}

type SpotMarketBatchExecutionData struct {
	Market                         *SpotMarket
	BaseDenomDepositMap            map[common.Hash]*DepositDelta
	QuoteDenomDepositMap           map[common.Hash]*DepositDelta
	BaseDenomDepositSubaccountIDs  []common.Hash
	QuoteDenomDepositSubaccountIDs []common.Hash
	LimitOrderFilledDeltas         []*SpotLimitOrderFilledDelta
	MarketOrderExecutionEvent      *EventBatchSpotExecution
	LimitOrderExecutionEvent       []*EventBatchSpotExecution
	NewOrdersEvent                 *EventNewSpotOrders
}

type SpotLimitOrderFilledDelta struct {
	SubaccountIndexKey []byte
	FillableAmount     sdk.Dec
}

type DepositDelta struct {
	AvailableBalanceDelta sdk.Dec
	TotalBalanceDelta     sdk.Dec
}

// Note: clearingPrice should be set to sdk.Dec{} for normal fills
func ProcessBothRestingLimitOrderExpansions(
	orderbookStateChange *OrderbookStateChange,
	marketID common.Hash,
	clearingPrice sdk.Dec,
	tradeFeeRate, relayerFeeShareRate sdk.Dec,
	baseDenomDepositMap map[common.Hash]*DepositDelta,
	quoteDenomDepositMap map[common.Hash]*DepositDelta,
) (limitBuyRestingOrderBatchEvent *EventBatchSpotExecution, limitSellRestingOrderBatchEvent *EventBatchSpotExecution, filledDeltas []*SpotLimitOrderFilledDelta) {
	spotLimitBuyOrderStateExpansions := make([]*SpotOrderStateExpansion, 0)
	spotLimitSellOrderStateExpansions := make([]*SpotOrderStateExpansion, 0)
	filledDeltas = make([]*SpotLimitOrderFilledDelta, 0)

	var currFilledDeltas []*SpotLimitOrderFilledDelta

	if orderbookStateChange.RestingBuyOrderbookFills != nil {
		spotLimitBuyOrderStateExpansions = ProcessRestingLimitOrderExpansions(orderbookStateChange.RestingBuyOrderbookFills, true, clearingPrice, tradeFeeRate, relayerFeeShareRate)
		// Process limit order events and filledDeltas
		limitBuyRestingOrderBatchEvent, currFilledDeltas = GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
			true,
			marketID,
			ExecutionType_LimitMatchRestingOrder,
			spotLimitBuyOrderStateExpansions,
			baseDenomDepositMap, quoteDenomDepositMap,
		)
		filledDeltas = append(filledDeltas, currFilledDeltas...)
	}

	if orderbookStateChange.RestingSellOrderbookFills != nil {
		spotLimitSellOrderStateExpansions = ProcessRestingLimitOrderExpansions(orderbookStateChange.RestingSellOrderbookFills, false, clearingPrice, tradeFeeRate, relayerFeeShareRate)
		// Process limit order events and filledDeltas
		limitSellRestingOrderBatchEvent, currFilledDeltas = GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
			false,
			marketID,
			ExecutionType_LimitMatchRestingOrder,
			spotLimitSellOrderStateExpansions,
			baseDenomDepositMap, quoteDenomDepositMap,
		)
		filledDeltas = append(filledDeltas, currFilledDeltas...)
	}
	return
}

// Note: clearingPrice should be set to sdk.Dec{} for normal fills
func ProcessBothTransientLimitOrderExpansions(
	orderbookStateChange *OrderbookStateChange,
	marketID common.Hash,
	clearingPrice sdk.Dec,
	makerFeeRate, takerFeeRate, relayerFeeShareRate sdk.Dec,
	baseDenomDepositMap map[common.Hash]*DepositDelta,
	quoteDenomDepositMap map[common.Hash]*DepositDelta,
) (limitBuyNewOrderBatchEvent *EventBatchSpotExecution, limitSellNewOrderBatchEvent *EventBatchSpotExecution,
	newRestingBuySpotLimitOrders []*SpotLimitOrder, newRestingSellSpotLimitOrders []*SpotLimitOrder,
) {
	var expansions []*SpotOrderStateExpansion
	if orderbookStateChange.NewBuyOrderbookFills != nil {
		expansions, newRestingBuySpotLimitOrders = ProcessNewLimitBuyExpansions(orderbookStateChange.NewBuyOrderbookFills, clearingPrice, makerFeeRate, takerFeeRate, relayerFeeShareRate)
		limitBuyNewOrderBatchEvent, _ = GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
			true,
			marketID,
			ExecutionType_LimitMatchNewOrder,
			expansions,
			baseDenomDepositMap, quoteDenomDepositMap,
		)
	}

	if orderbookStateChange.NewSellOrderbookFills != nil {
		expansions, newRestingSellSpotLimitOrders = ProcessNewLimitSellExpansions(orderbookStateChange.NewSellOrderbookFills, clearingPrice, takerFeeRate, relayerFeeShareRate)
		limitSellNewOrderBatchEvent, _ = GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
			false,
			marketID,
			ExecutionType_LimitMatchNewOrder,
			expansions,
			baseDenomDepositMap, quoteDenomDepositMap,
		)
	}
	return
}

// Note: clearingPrice should be set to sdk.Dec{} for normal fills
func ProcessRestingLimitOrderExpansions(
	orderbookState *OrderbookFills,
	isLimitBuy bool,
	clearingPrice sdk.Dec,
	makerFeeRate, relayerFeeShare sdk.Dec,
) []*SpotOrderStateExpansion {
	stateExpansions := make([]*SpotOrderStateExpansion, len(orderbookState.Orders))

	for idx, order := range orderbookState.Orders {
		fillQuantity, fillPrice := orderbookState.FillQuantities[idx], order.OrderInfo.Price
		if !clearingPrice.IsNil() {
			fillPrice = clearingPrice
		}

		if isLimitBuy {
			stateExpansions[idx] = GetRestingLimitBuyStateExpansion(
				order,
				common.BytesToHash(order.Hash),
				fillQuantity,
				fillPrice,
				makerFeeRate,
				relayerFeeShare,
			)
		} else {
			stateExpansions[idx] = GetLimitSellStateExpansion(
				order,
				fillQuantity,
				fillPrice,
				makerFeeRate,
				relayerFeeShare,
			)
		}
	}
	return stateExpansions
}

// Note: clearingPrice should be set to sdk.Dec{} for normal fills
func ProcessNewLimitSellExpansions(
	orderbookFills *OrderbookFills,
	clearingPrice sdk.Dec,
	takerFeeRate, relayerFeeShare sdk.Dec,
) ([]*SpotOrderStateExpansion, []*SpotLimitOrder) {
	stateExpansions := make([]*SpotOrderStateExpansion, len(orderbookFills.Orders))
	newRestingOrders := make([]*SpotLimitOrder, 0, len(orderbookFills.Orders))

	for idx, order := range orderbookFills.Orders {
		fillQuantity, fillPrice := orderbookFills.FillQuantities[idx], order.OrderInfo.Price
		if !clearingPrice.IsNil() {
			fillPrice = clearingPrice
		}
		stateExpansions[idx] = GetLimitSellStateExpansion(
			order,
			fillQuantity,
			fillPrice,
			takerFeeRate,
			relayerFeeShare,
		)
		if fillQuantity.LT(order.OrderInfo.Quantity) {
			order.Fillable = order.Fillable.Sub(fillQuantity)
			newRestingOrders = append(newRestingOrders, order)
		}
	}
	return stateExpansions, newRestingOrders
}

func GetLimitSellStateExpansion(
	sellOrder *SpotLimitOrder,
	fillQuantity, fillPrice, tradeFeeRate, relayerFeeShare sdk.Dec,
) *SpotOrderStateExpansion {
	orderNotional := fillQuantity.Mul(fillPrice)

	tradingFee := orderNotional.Mul(tradeFeeRate)
	feeRecipientReward := relayerFeeShare.Mul(tradingFee)
	auctionFeeReward := tradingFee.Sub(feeRecipientReward)

	// limit sells are credited with the (fillQuantity * price) * (1 - tradeFeeRate) in quote denom
	quoteChangeAmount := orderNotional.Sub(tradingFee)

	stateExpansion := SpotOrderStateExpansion{
		// limit sells are debited by fillQuantity in base denom
		BaseChangeAmount:   fillQuantity.Neg(),
		QuoteChangeAmount:  quoteChangeAmount,
		QuoteRefundAmount:  sdk.ZeroDec(),
		FeeRecipient:       common.HexToAddress(sellOrder.OrderInfo.FeeRecipient),
		FeeRecipientReward: feeRecipientReward,
		AuctionFeeReward:   auctionFeeReward,
		Hash:               common.BytesToHash(sellOrder.Hash),
		SubaccountID:       common.HexToHash(sellOrder.OrderInfo.SubaccountId),
		FillableAmount:     sellOrder.Fillable.Sub(fillQuantity),
	}
	return &stateExpansion
}

func GetRestingLimitBuyStateExpansion(
	buyOrder *SpotLimitOrder,
	orderHash common.Hash,
	fillQuantity, fillPrice, makerFeeRate, relayerFeeShare sdk.Dec,
) *SpotOrderStateExpansion {
	var baseChangeAmount, quoteChangeAmount sdk.Dec
	fillableAmount := buyOrder.Fillable.Sub(fillQuantity)
	orderNotional := fillQuantity.Mul(fillPrice)
	tradingFee := orderNotional.Mul(makerFeeRate)
	feeRecipientReward := relayerFeeShare.Mul(tradingFee)
	auctionFeeReward := tradingFee.Sub(feeRecipientReward)
	// limit buys are credited with the order fill quantity in base denom
	baseChangeAmount = fillQuantity
	// limit buys are debited with (fillQuantity * Price) * (1 + makerFee) in quote denom
	quoteChangeAmount = orderNotional.Add(tradingFee).Neg()
	quoteRefund := sdk.ZeroDec()
	if !fillPrice.Equal(buyOrder.OrderInfo.Price) {
		priceDelta := buyOrder.OrderInfo.Price.Sub(fillPrice)
		clearingRefund := fillQuantity.Mul(priceDelta)
		matchedFeeRefund := fillQuantity.Mul(makerFeeRate).Mul(priceDelta)
		quoteRefund = clearingRefund.Add(matchedFeeRefund)
	}
	stateExpansion := SpotOrderStateExpansion{
		BaseChangeAmount:   baseChangeAmount,
		QuoteChangeAmount:  quoteChangeAmount,
		QuoteRefundAmount:  quoteRefund,
		FeeRecipient:       common.HexToAddress(buyOrder.OrderInfo.FeeRecipient),
		FeeRecipientReward: feeRecipientReward,
		AuctionFeeReward:   auctionFeeReward,
		Hash:               orderHash,
		SubaccountID:       common.HexToHash(buyOrder.OrderInfo.SubaccountId),
		FillableAmount:     fillableAmount,
	}
	return &stateExpansion
}

func ProcessNewLimitBuyExpansions(
	orderbookState *OrderbookFills,
	clearingPrice sdk.Dec,
	makerFeeRate, takerFeeRate, relayerFeeShare sdk.Dec,
) ([]*SpotOrderStateExpansion, []*SpotLimitOrder) {
	stateExpansions := make([]*SpotOrderStateExpansion, len(orderbookState.Orders))
	newRestingOrders := make([]*SpotLimitOrder, 0, len(orderbookState.Orders))

	for idx, order := range orderbookState.Orders {
		fillQuantity := sdk.ZeroDec()
		if orderbookState.FillQuantities != nil {
			fillQuantity = orderbookState.FillQuantities[idx]
		}
		stateExpansions[idx] = GetNewLimitBuyStateExpansion(
			order,
			common.BytesToHash(order.Hash),
			clearingPrice, fillQuantity,
			makerFeeRate, takerFeeRate, relayerFeeShare,
		)

		if fillQuantity.LT(order.OrderInfo.Quantity) {
			order.Fillable = order.Fillable.Sub(fillQuantity)
			newRestingOrders = append(newRestingOrders, order)
		}
	}
	return stateExpansions, newRestingOrders
}

func GetNewLimitBuyStateExpansion(
	buyOrder *SpotLimitOrder,
	orderHash common.Hash,
	clearingPrice, fillQuantity,
	makerFeeRate, takerFeeRate, relayerFeeShare sdk.Dec,
) *SpotOrderStateExpansion {
	// TODO: optimize for the case when fillQuantity is 0
	var baseChangeAmount, quoteChangeAmount sdk.Dec
	fillableAmount := buyOrder.Fillable.Sub(fillQuantity)

	orderNotional := sdk.ZeroDec()
	clearingRefund := sdk.ZeroDec()
	matchedFeeRefund := sdk.ZeroDec()
	if !fillQuantity.IsZero() {
		orderNotional = fillQuantity.Mul(clearingPrice)
		priceDelta := buyOrder.OrderInfo.Price.Sub(clearingPrice)
		// Clearing Refund = FillQuantity * (Price - ClearingPrice)
		clearingRefund = fillQuantity.Mul(priceDelta)
		// Matched Fee Refund = FillQuantity * TakerFeeRate * (Price - ClearingPrice)
		matchedFeeRefund = fillQuantity.Mul(takerFeeRate).Mul(priceDelta)
	}
	tradingFee := orderNotional.Mul(takerFeeRate)
	feeRecipientReward := relayerFeeShare.Mul(tradingFee)
	auctionFeeReward := tradingFee.Sub(feeRecipientReward)
	// limit buys are credited with the order fill quantity in base denom
	baseChangeAmount = fillQuantity
	// limit buys are debited with (fillQuantity * Price) * (1 + makerFee) in quote denom
	quoteChangeAmount = orderNotional.Add(tradingFee).Neg()
	// Unmatched Fee Refund = (Quantity - FillQuantity) * Price * (TakerFeeRate - MakerFeeRate)
	unmatchedFeeRefund := buyOrder.OrderInfo.Quantity.Sub(fillQuantity).Mul(buyOrder.OrderInfo.Price).Mul(takerFeeRate.Sub(makerFeeRate))
	// Fee Refund = Matched Fee Refund + Unmatched Fee Refund
	feeRefund := matchedFeeRefund.Add(unmatchedFeeRefund)
	// refund amount = clearing refund + matched fee refund + unmatched fee refund
	quoteRefundAmount := clearingRefund.Add(feeRefund)
	stateExpansion := SpotOrderStateExpansion{
		BaseChangeAmount:   baseChangeAmount,
		QuoteChangeAmount:  quoteChangeAmount,
		QuoteRefundAmount:  quoteRefundAmount,
		FeeRecipient:       common.HexToAddress(buyOrder.OrderInfo.FeeRecipient),
		FeeRecipientReward: feeRecipientReward,
		AuctionFeeReward:   auctionFeeReward,
		Hash:               orderHash,
		SubaccountID:       common.HexToHash(buyOrder.OrderInfo.SubaccountId),
		FillableAmount:     fillableAmount,
	}
	return &stateExpansion
}

// NOTE: clearingPrice may be Nil
func ProcessMarketOrderStateExpansions(
	isMarketBuy bool,
	marketOrders []*SpotMarketOrder,
	marketFillQuantities []sdk.Dec,
	clearingPrice sdk.Dec,
	tradeFeeRate, relayerFeeShareRate sdk.Dec,
) []*SpotOrderStateExpansion {
	stateExpansions := make([]*SpotOrderStateExpansion, len(marketOrders))

	for idx := range marketOrders {
		stateExpansions[idx] = GetMarketOrderStateExpansion(
			marketOrders[idx],
			isMarketBuy,
			marketFillQuantities[idx],
			clearingPrice,
			tradeFeeRate,
			relayerFeeShareRate,
		)
	}
	return stateExpansions
}

func GetMarketOrderStateExpansion(
	marketOrder *SpotMarketOrder,
	isMarketBuy bool,
	fillQuantity, clearingPrice sdk.Dec,
	takerFeeRate, relayerFeeShare sdk.Dec,
) *SpotOrderStateExpansion {
	var baseChangeAmount, quoteChangeAmount sdk.Dec

	if fillQuantity.IsNil() {
		fillQuantity = sdk.ZeroDec()
	}
	orderNotional := sdk.ZeroDec()
	if !clearingPrice.IsNil() {
		orderNotional = fillQuantity.Mul(clearingPrice)
	}
	tradingFee := orderNotional.Mul(takerFeeRate)
	feeRecipientReward := relayerFeeShare.Mul(tradingFee)
	auctionFeeReward := tradingFee.Sub(feeRecipientReward)
	quoteRefundAmount, quoteChangeAmount := sdk.ZeroDec(), sdk.ZeroDec()
	if isMarketBuy {
		// market buys are credited with the order fill quantity in base denom
		baseChangeAmount = fillQuantity
		// market buys are debited with (fillQuantity * clearingPrice) * (1 + takerFee) in quote denom
		if !clearingPrice.IsNil() {
			quoteChangeAmount = fillQuantity.Mul(clearingPrice).Add(tradingFee).Neg()
		}
		quoteRefundAmount = marketOrder.BalanceHold.Add(quoteChangeAmount)
	} else {
		// market sells are debited by fillQuantity in base denom
		baseChangeAmount = fillQuantity.Neg()
		// market sells are credited with the (fillQuantity * clearingPrice) * (1 - TakerFee) in quote denom
		if !clearingPrice.IsNil() {
			quoteChangeAmount = orderNotional.Sub(tradingFee)
		}
	}
	stateExpansion := SpotOrderStateExpansion{
		BaseChangeAmount:   baseChangeAmount,
		QuoteChangeAmount:  quoteChangeAmount,
		QuoteRefundAmount:  quoteRefundAmount,
		FeeRecipient:       common.HexToAddress(marketOrder.OrderInfo.FeeRecipient),
		FeeRecipientReward: feeRecipientReward,
		AuctionFeeReward:   auctionFeeReward,
		Hash:               common.HexToHash(marketOrder.Hash),
		SubaccountID:       common.HexToHash(marketOrder.OrderInfo.SubaccountId),
		FillableAmount:     marketOrder.OrderInfo.Quantity.Sub(fillQuantity),
	}
	return &stateExpansion
}

func GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
	isBuy bool,
	marketID common.Hash,
	executionType ExecutionType,
	spotLimitOrderStateExpansions []*SpotOrderStateExpansion,
	baseDenomDepositMap map[common.Hash]*DepositDelta, quoteDenomDepositMap map[common.Hash]*DepositDelta,
) (*EventBatchSpotExecution, []*SpotLimitOrderFilledDelta) {
	limitOrderBatchEvent := EventBatchSpotExecution{
		MarketId:      marketID.Hex(),
		IsBuy:         isBuy,
		ExecutionType: executionType,
	}

	trades := make([]*TradeLog, 0, len(spotLimitOrderStateExpansions))

	// array of (SubaccountIndexKey, fillableAmount) to update/delete
	filledDeltas := make([]*SpotLimitOrderFilledDelta, 0, len(spotLimitOrderStateExpansions))

	for idx := range spotLimitOrderStateExpansions {
		expansion := spotLimitOrderStateExpansions[idx]
		UpdateDepositMap(baseDenomDepositMap, quoteDenomDepositMap, expansion)
		// skip adding trade data if there was no trade (unfilled new order)
		fillQuantity := spotLimitOrderStateExpansions[idx].BaseChangeAmount
		if fillQuantity.IsZero() {
			continue
		}

		filledDeltas = append(filledDeltas, &SpotLimitOrderFilledDelta{
			SubaccountIndexKey: GetLimitOrderBySubaccountKey(marketID, limitOrderBatchEvent.IsBuy, expansion.SubaccountID, expansion.Hash),
			FillableAmount:     expansion.FillableAmount,
		})

		fee := expansion.FeeRecipientReward.Add(expansion.AuctionFeeReward)
		// Fee is always positive, so for both cases can just be added to the quote change amount.
		// For limit sells, QuoteChangeAmount is positive (receiving quote), but already includes the paid fees. To get the actual price, add the fee.
		// For limit buys, QuoteChangeAmount is negative (selling quote), but also was used to pay the fee. To get the actual price, add the fee.
		price := expansion.QuoteChangeAmount.Add(fee).Quo(expansion.BaseChangeAmount).Abs()

		trades = append(trades, &TradeLog{
			Quantity:     expansion.BaseChangeAmount.Abs(),
			Price:        price,
			SubaccountId: expansion.SubaccountID.Hex(),
			Fee:          fee,
			Hash:         expansion.Hash.Hex(),
		})
	}
	limitOrderBatchEvent.Trades = trades
	return &limitOrderBatchEvent, filledDeltas
}

func UpdateDepositMap(baseDenomDepositMap map[common.Hash]*DepositDelta, quoteDenomDepositMap map[common.Hash]*DepositDelta, expansion *SpotOrderStateExpansion) {
	baseDenomDeposit := baseDenomDepositMap[expansion.SubaccountID]
	if baseDenomDeposit == nil {
		availableBalanceDelta := sdk.ZeroDec()
		// increment availableBalanceDelta in tandem with TotalBalanceDelta if positive
		if expansion.BaseChangeAmount.IsPositive() {
			availableBalanceDelta = expansion.BaseChangeAmount
		}
		baseDenomDepositMap[expansion.SubaccountID] = &DepositDelta{
			TotalBalanceDelta:     expansion.BaseChangeAmount,
			AvailableBalanceDelta: availableBalanceDelta,
		}
	} else {
		baseDenomDeposit.TotalBalanceDelta = expansion.BaseChangeAmount.Add(baseDenomDeposit.TotalBalanceDelta)
	}

	traderQuoteDepositDelta := quoteDenomDepositMap[expansion.SubaccountID]
	if traderQuoteDepositDelta == nil {
		quoteDenomDepositMap[expansion.SubaccountID] = &DepositDelta{
			TotalBalanceDelta:     sdk.ZeroDec(),
			AvailableBalanceDelta: sdk.ZeroDec(),
		}
		traderQuoteDepositDelta = quoteDenomDepositMap[expansion.SubaccountID]
	}

	traderQuoteDepositDelta.TotalBalanceDelta = expansion.QuoteChangeAmount.Add(traderQuoteDepositDelta.TotalBalanceDelta)
	traderQuoteDepositDelta.AvailableBalanceDelta = expansion.QuoteRefundAmount.Add(traderQuoteDepositDelta.AvailableBalanceDelta)

	// increment availableBalanceDelta in tandem with TotalBalanceDelta if positive
	if expansion.QuoteChangeAmount.IsPositive() {
		traderQuoteDepositDelta.AvailableBalanceDelta = expansion.QuoteChangeAmount.Add(traderQuoteDepositDelta.AvailableBalanceDelta)
	}

	// increment fee recipient's balances
	feeSubaccount := common.BytesToHash(common.RightPadBytes(expansion.FeeRecipient.Bytes(), common.HashLength))

	feeRecipientQuoteDepositDelta := quoteDenomDepositMap[feeSubaccount]
	if feeRecipientQuoteDepositDelta == nil {
		quoteDenomDepositMap[feeSubaccount] = &DepositDelta{
			TotalBalanceDelta:     sdk.ZeroDec(),
			AvailableBalanceDelta: sdk.ZeroDec(),
		}
		feeRecipientQuoteDepositDelta = quoteDenomDepositMap[feeSubaccount]

	}
	feeRecipientQuoteDepositDelta.TotalBalanceDelta = feeRecipientQuoteDepositDelta.TotalBalanceDelta.Add(expansion.FeeRecipientReward)
	feeRecipientQuoteDepositDelta.AvailableBalanceDelta = feeRecipientQuoteDepositDelta.AvailableBalanceDelta.Add(expansion.FeeRecipientReward)

	// increment auction fee balance
	auctionQuoteDepositDelta := quoteDenomDepositMap[ZeroHash]
	if auctionQuoteDepositDelta == nil {
		quoteDenomDepositMap[ZeroHash] = &DepositDelta{
			AvailableBalanceDelta: sdk.ZeroDec(),
			TotalBalanceDelta:     sdk.ZeroDec(),
		}
		auctionQuoteDepositDelta = quoteDenomDepositMap[ZeroHash]
	}
	auctionQuoteDepositDelta.TotalBalanceDelta = auctionQuoteDepositDelta.TotalBalanceDelta.Add(expansion.AuctionFeeReward)
	auctionQuoteDepositDelta.AvailableBalanceDelta = auctionQuoteDepositDelta.AvailableBalanceDelta.Add(expansion.AuctionFeeReward)
}
