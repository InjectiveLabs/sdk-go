package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type SpotBatchExecutionData struct {
	Market                         *SpotMarket
	BaseDenomDepositDeltas         DepositDeltas
	QuoteDenomDepositDeltas        DepositDeltas
	BaseDenomDepositSubaccountIDs  []common.Hash
	QuoteDenomDepositSubaccountIDs []common.Hash
	LimitOrderFilledDeltas         []*SpotLimitOrderDelta
	MarketOrderExecutionEvent      *EventBatchSpotExecution
	LimitOrderExecutionEvent       []*EventBatchSpotExecution
	NewOrdersEvent                 *EventNewSpotOrders
}

func GetSpotMarketOrderBatchExecution(
	isMarketBuy bool,
	market *SpotMarket,
	spotLimitOrderStateExpansions, spotMarketOrderStateExpansions []*SpotOrderStateExpansion,
	clearingPrice sdk.Dec,
) *SpotBatchExecutionData {
	baseDenomDepositDeltas := NewDepositDeltas()
	quoteDenomDepositDeltas := NewDepositDeltas()

	// Step 3a: Process market order events
	marketOrderBatchEvent := &EventBatchSpotExecution{
		MarketId:      market.MarketID().Hex(),
		IsBuy:         isMarketBuy,
		ExecutionType: ExecutionType_Market,
	}

	trades := make([]*TradeLog, len(spotMarketOrderStateExpansions))

	for idx := range spotMarketOrderStateExpansions {
		expansion := spotMarketOrderStateExpansions[idx]
		expansion.UpdateDepositDeltas(baseDenomDepositDeltas, quoteDenomDepositDeltas)

		trades[idx] = &TradeLog{
			Quantity:     expansion.BaseChangeAmount.Abs(),
			Price:        clearingPrice,
			SubaccountId: expansion.SubaccountID.Bytes(),
			Fee:          expansion.FeeRecipientReward.Add(expansion.AuctionFeeReward),
			OrderHash:    expansion.OrderHash.Bytes(),
		}
	}
	marketOrderBatchEvent.Trades = trades

	if len(trades) == 0 {
		marketOrderBatchEvent = nil
	}

	// Stage 3b: Process limit order events
	limitOrderBatchEvent, filledDeltas := GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
		!isMarketBuy,
		market.MarketID(),
		ExecutionType_LimitFill,
		spotLimitOrderStateExpansions,
		baseDenomDepositDeltas, quoteDenomDepositDeltas,
	)

	limitOrderExecutionEvent := make([]*EventBatchSpotExecution, 0)
	if limitOrderBatchEvent != nil {
		limitOrderExecutionEvent = append(limitOrderExecutionEvent, limitOrderBatchEvent)
	}

	// Final Step: Store the SpotBatchExecutionData for future reduction/processing
	batch := &SpotBatchExecutionData{
		Market:                         market,
		BaseDenomDepositDeltas:         baseDenomDepositDeltas,
		QuoteDenomDepositDeltas:        quoteDenomDepositDeltas,
		BaseDenomDepositSubaccountIDs:  baseDenomDepositDeltas.GetSortedSubaccountKeys(),
		QuoteDenomDepositSubaccountIDs: quoteDenomDepositDeltas.GetSortedSubaccountKeys(),
		LimitOrderFilledDeltas:         filledDeltas,
		MarketOrderExecutionEvent:      marketOrderBatchEvent,
		LimitOrderExecutionEvent:       limitOrderExecutionEvent,
	}
	return batch
}

func GetSpotLimitMatchingBatchExecution(
	isCanaryV2 bool,
	market *SpotMarket,
	orderbookStateChange *SpotOrderbookStateChange,
	clearingPrice sdk.Dec,
) *SpotBatchExecutionData {

	// Initialize map DepositKey subaccountID => Deposit Delta (availableBalanceDelta, totalDepositsDelta)
	baseDenomDepositDeltas := NewDepositDeltas()
	quoteDenomDepositDeltas := NewDepositDeltas()

	limitBuyRestingOrderBatchEvent, limitSellRestingOrderBatchEvent, filledDeltas := orderbookStateChange.ProcessBothRestingSpotLimitOrderExpansions(
		isCanaryV2, market.MarketID(), clearingPrice, market.MakerFeeRate, market.RelayerFeeShareRate,
		baseDenomDepositDeltas, quoteDenomDepositDeltas,
	)

	// filled deltas are handled implicitly with the new resting spot limit orders
	limitBuyNewOrderBatchEvent, limitSellNewOrderBatchEvent, newRestingBuySpotLimitOrders, newRestingSellSpotLimitOrders := orderbookStateChange.ProcessBothTransientSpotLimitOrderExpansions(
		isCanaryV2,
		market.MarketID(),
		clearingPrice,
		market.MakerFeeRate, market.TakerFeeRate, market.RelayerFeeShareRate,
		baseDenomDepositDeltas, quoteDenomDepositDeltas,
	)

	eventBatchSpotExecution := make([]*EventBatchSpotExecution, 0)

	if limitBuyRestingOrderBatchEvent != nil {
		eventBatchSpotExecution = append(eventBatchSpotExecution, limitBuyRestingOrderBatchEvent)
	}

	if limitSellRestingOrderBatchEvent != nil {
		eventBatchSpotExecution = append(eventBatchSpotExecution, limitSellRestingOrderBatchEvent)
	}

	if limitBuyNewOrderBatchEvent != nil {
		eventBatchSpotExecution = append(eventBatchSpotExecution, limitBuyNewOrderBatchEvent)
	}

	if limitSellNewOrderBatchEvent != nil {
		eventBatchSpotExecution = append(eventBatchSpotExecution, limitSellNewOrderBatchEvent)
	}

	// Final Step: Store the SpotBatchExecutionData for future reduction/processing
	batch := &SpotBatchExecutionData{
		Market:                         market,
		BaseDenomDepositDeltas:         baseDenomDepositDeltas,
		QuoteDenomDepositDeltas:        quoteDenomDepositDeltas,
		BaseDenomDepositSubaccountIDs:  baseDenomDepositDeltas.GetSortedSubaccountKeys(),
		QuoteDenomDepositSubaccountIDs: quoteDenomDepositDeltas.GetSortedSubaccountKeys(),
		LimitOrderFilledDeltas:         filledDeltas,
		LimitOrderExecutionEvent:       eventBatchSpotExecution,
	}

	if len(newRestingBuySpotLimitOrders) > 0 || len(newRestingSellSpotLimitOrders) > 0 {
		batch.NewOrdersEvent = &EventNewSpotOrders{
			MarketId:   market.MarketId,
			BuyOrders:  newRestingBuySpotLimitOrders,
			SellOrders: newRestingSellSpotLimitOrders,
		}
	}
	return batch
}

func GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
	isBuy bool,
	marketID common.Hash,
	executionType ExecutionType,
	spotLimitOrderStateExpansions []*SpotOrderStateExpansion,
	baseDenomDepositDeltas DepositDeltas, quoteDenomDepositDeltas DepositDeltas,
) (*EventBatchSpotExecution, []*SpotLimitOrderDelta) {
	limitOrderBatchEvent := &EventBatchSpotExecution{
		MarketId:      marketID.Hex(),
		IsBuy:         isBuy,
		ExecutionType: executionType,
	}

	trades := make([]*TradeLog, 0, len(spotLimitOrderStateExpansions))

	// array of (SubaccountIndexKey, fillableAmount) to update/delete
	filledDeltas := make([]*SpotLimitOrderDelta, 0, len(spotLimitOrderStateExpansions))

	for idx := range spotLimitOrderStateExpansions {
		expansion := spotLimitOrderStateExpansions[idx]
		expansion.UpdateDepositDeltas(baseDenomDepositDeltas, quoteDenomDepositDeltas)

		// skip adding trade data if there was no trade (unfilled new order)
		fillQuantity := spotLimitOrderStateExpansions[idx].BaseChangeAmount
		if fillQuantity.IsZero() {
			continue
		}

		filledDeltas = append(filledDeltas, &SpotLimitOrderDelta{
			Order:        expansion.LimitOrder,
			FillQuantity: expansion.LimitOrderFillQuantity,
		})

		hasNegativeMakerFee := expansion.AuctionFeeReward.IsNegative()
		var traderFee sdk.Dec

		// For limit buys, QuoteChangeAmount is negative (selling quote), but also was used to pay the fee. To get the actual price, add the fee.
		// For limit sells, QuoteChangeAmount is positive (receiving quote), but already includes the paid fees. To get the actual price, add the fee.

		if hasNegativeMakerFee {
			// TraderFeeReward is positive, subtract the value from QuoteChangeAmount
			traderFee = expansion.TraderFeeReward.Neg()
		} else {
			traderFee = expansion.FeeRecipientReward.Add(expansion.AuctionFeeReward)
		}

		price := expansion.QuoteChangeAmount.Add(traderFee).Quo(expansion.BaseChangeAmount).Abs()
		var totalTradeFee sdk.Dec

		if hasNegativeMakerFee {
			totalTradeFee = expansion.FeeRecipientReward.Add(expansion.TraderFeeReward)
		} else {
			totalTradeFee = traderFee
		}

		trades = append(trades, &TradeLog{
			Quantity:     expansion.BaseChangeAmount.Abs(),
			Price:        price,
			SubaccountId: expansion.SubaccountID.Bytes(),
			Fee:          totalTradeFee,
			OrderHash:    expansion.OrderHash.Bytes(),
		})
	}
	limitOrderBatchEvent.Trades = trades

	if len(trades) == 0 {
		limitOrderBatchEvent = nil
	}
	return limitOrderBatchEvent, filledDeltas
}
