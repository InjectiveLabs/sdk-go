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
	LimitOrderFilledDeltas         []*LimitOrderFilledDelta
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
	marketOrderBatchEvent := EventBatchSpotExecution{
		MarketId:      market.MarketID().Hex(),
		IsBuy:         isMarketBuy,
		ExecutionType: ExecutionType_Market,
	}

	trades := make([]*TradeLog, len(spotMarketOrderStateExpansions))

	for idx := range spotMarketOrderStateExpansions {
		expansion := spotMarketOrderStateExpansions[idx]
		expansion.UpdateDepositDeltas(baseDenomDepositDeltas, quoteDenomDepositDeltas)

		trades[idx] = &TradeLog{
			Quantity:     expansion.BaseChangeAmount,
			Price:        clearingPrice,
			SubaccountId: expansion.SubaccountID.Bytes(),
			Fee:          expansion.FeeRecipientReward.Add(expansion.AuctionFeeReward),
			OrderHash:    expansion.OrderHash.Bytes(),
		}
	}
	marketOrderBatchEvent.Trades = trades

	// Stage 3b: Process limit order events
	limitOrderBatchEvent, filledDeltas := GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
		!isMarketBuy,
		market.MarketID(),
		ExecutionType_LimitFill,
		spotLimitOrderStateExpansions,
		baseDenomDepositDeltas, quoteDenomDepositDeltas,
	)

	// Final Step: Store the SpotBatchExecutionData for future reduction/processing
	batch := &SpotBatchExecutionData{
		Market:                         market,
		BaseDenomDepositDeltas:         baseDenomDepositDeltas,
		QuoteDenomDepositDeltas:        quoteDenomDepositDeltas,
		BaseDenomDepositSubaccountIDs:  baseDenomDepositDeltas.GetSortedSubaccountKeys(),
		QuoteDenomDepositSubaccountIDs: quoteDenomDepositDeltas.GetSortedSubaccountKeys(),
		LimitOrderFilledDeltas:         filledDeltas,
		MarketOrderExecutionEvent:      &marketOrderBatchEvent,
		LimitOrderExecutionEvent:       []*EventBatchSpotExecution{limitOrderBatchEvent},
	}
	return batch
}

func GetSpotLimitMatchingBatchExecution(
	market *SpotMarket,
	orderbookStateChange *SpotOrderbookStateChange,
	clearingPrice sdk.Dec,
) *SpotBatchExecutionData {

	// Initialize map DepositKey subaccountID => Deposit Delta (availableBalanceDelta, totalDepositsDelta)
	baseDenomDepositDeltas := NewDepositDeltas()
	quoteDenomDepositDeltas := NewDepositDeltas()

	limitBuyRestingOrderBatchEvent, limitSellRestingOrderBatchEvent, filledDeltas := orderbookStateChange.ProcessBothRestingSpotLimitOrderExpansions(
		market.MarketID(), clearingPrice, market.MakerFeeRate, market.RelayerFeeShareRate,
		baseDenomDepositDeltas, quoteDenomDepositDeltas,
	)

	// filled deltas are handled implicitly with the new resting spot limit orders
	limitBuyNewOrderBatchEvent, limitSellNewOrderBatchEvent, newRestingBuySpotLimitOrders, newRestingSellSpotLimitOrders := orderbookStateChange.ProcessBothTransientSpotLimitOrderExpansions(
		market.MarketID(),
		clearingPrice,
		market.MakerFeeRate, market.TakerFeeRate, market.RelayerFeeShareRate,
		baseDenomDepositDeltas, quoteDenomDepositDeltas,
	)

	// Final Step: Store the SpotBatchExecutionData for future reduction/processing
	batch := &SpotBatchExecutionData{
		Market:                         market,
		BaseDenomDepositDeltas:         baseDenomDepositDeltas,
		QuoteDenomDepositDeltas:        quoteDenomDepositDeltas,
		BaseDenomDepositSubaccountIDs:  baseDenomDepositDeltas.GetSortedSubaccountKeys(),
		QuoteDenomDepositSubaccountIDs: quoteDenomDepositDeltas.GetSortedSubaccountKeys(),
		LimitOrderFilledDeltas:         filledDeltas,
		LimitOrderExecutionEvent: []*EventBatchSpotExecution{
			limitBuyRestingOrderBatchEvent,
			limitSellRestingOrderBatchEvent,
			limitBuyNewOrderBatchEvent,
			limitSellNewOrderBatchEvent,
		},
		NewOrdersEvent: &EventNewSpotOrders{
			MarketId:   market.MarketId,
			BuyOrders:  newRestingBuySpotLimitOrders,
			SellOrders: newRestingSellSpotLimitOrders,
		},
	}
	return batch
}

func GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
	isBuy bool,
	marketID common.Hash,
	executionType ExecutionType,
	spotLimitOrderStateExpansions []*SpotOrderStateExpansion,
	baseDenomDepositDeltas DepositDeltas, quoteDenomDepositDeltas DepositDeltas,
) (*EventBatchSpotExecution, []*LimitOrderFilledDelta) {
	limitOrderBatchEvent := EventBatchSpotExecution{
		MarketId:      marketID.Hex(),
		IsBuy:         isBuy,
		ExecutionType: executionType,
	}

	trades := make([]*TradeLog, 0, len(spotLimitOrderStateExpansions))

	// array of (SubaccountIndexKey, fillableAmount) to update/delete
	filledDeltas := make([]*LimitOrderFilledDelta, 0, len(spotLimitOrderStateExpansions))

	for idx := range spotLimitOrderStateExpansions {
		expansion := spotLimitOrderStateExpansions[idx]
		expansion.UpdateDepositDeltas(baseDenomDepositDeltas, quoteDenomDepositDeltas)

		// skip adding trade data if there was no trade (unfilled new order)
		fillQuantity := spotLimitOrderStateExpansions[idx].BaseChangeAmount
		if fillQuantity.IsZero() {
			continue
		}

		filledDeltas = append(filledDeltas, &LimitOrderFilledDelta{
			SubaccountIndexKey: GetLimitOrderIndexKey(
				marketID,
				limitOrderBatchEvent.IsBuy,
				expansion.SubaccountID,
				expansion.OrderHash,
			),
			FillableAmount: expansion.FillableAmount,
		})

		fee := expansion.FeeRecipientReward.Add(expansion.AuctionFeeReward)
		// Fee is always positive, so for both cases can just be added to the quote change amount.
		// For limit sells, QuoteChangeAmount is positive (receiving quote), but already includes the paid fees. To get the actual price, add the fee.
		// For limit buys, QuoteChangeAmount is negative (selling quote), but also was used to pay the fee. To get the actual price, add the fee.
		price := expansion.QuoteChangeAmount.Add(fee).Quo(expansion.BaseChangeAmount).Abs()

		trades = append(trades, &TradeLog{
			Quantity:     expansion.BaseChangeAmount.Abs(),
			Price:        price,
			SubaccountId: expansion.SubaccountID.Bytes(),
			Fee:          fee,
			OrderHash:    expansion.OrderHash.Bytes(),
		})
	}
	limitOrderBatchEvent.Trades = trades
	return &limitOrderBatchEvent, filledDeltas
}
