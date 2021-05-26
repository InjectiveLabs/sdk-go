package types

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type SpotOrderStateExpansion struct {
	BaseChangeAmount        sdk.Dec
	BaseRefundAmount        sdk.Dec
	QuoteChangeAmount       sdk.Dec
	QuoteRefundAmount       sdk.Dec
	FeeRecipient            common.Address
	FeeRecipientReward      sdk.Dec
	AuctionFeeReward        sdk.Dec
	LimitOrder              *SpotLimitOrder
	LimitOrderFillQuantity  sdk.Dec
	MarketOrder             *SpotMarketOrder
	MarketOrderFillQuantity sdk.Dec
	OrderHash               common.Hash
	OrderPrice              sdk.Dec
	SubaccountID            common.Hash
}

// ProcessSpotMarketOrderStateExpansions processes the spot market order state expansions.
// NOTE: clearingPrice may be Nil
func ProcessSpotMarketOrderStateExpansions(
	isMarketBuy bool,
	marketOrders []*SpotMarketOrder,
	marketFillQuantities []sdk.Dec,
	clearingPrice sdk.Dec,
	tradeFeeRate, relayerFeeShareRate sdk.Dec,
) []*SpotOrderStateExpansion {
	stateExpansions := make([]*SpotOrderStateExpansion, len(marketOrders))

	for idx := range marketOrders {
		stateExpansions[idx] = getSpotMarketOrderStateExpansion(
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

func ProcessRestingSpotLimitOrderExpansions(
	fills *OrderbookFills,
	isLimitBuy bool,
	clearingPrice sdk.Dec,
	makerFeeRate, relayerFeeShare sdk.Dec,
) []*SpotOrderStateExpansion {
	stateExpansions := make([]*SpotOrderStateExpansion, len(fills.Orders))
	for idx, order := range fills.Orders {
		fillQuantity, fillPrice := fills.FillQuantities[idx], order.OrderInfo.Price
		if !clearingPrice.IsNil() {
			fillPrice = clearingPrice
		}

		if isLimitBuy {
			stateExpansions[idx] = getRestingSpotLimitBuyStateExpansion(
				order,
				order.Hash(),
				fillQuantity,
				fillPrice,
				makerFeeRate,
				relayerFeeShare,
			)
		} else {
			stateExpansions[idx] = getSpotLimitSellStateExpansion(
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

func (e *SpotOrderStateExpansion) UpdateDepositDeltas(baseDenomDepositDeltas DepositDeltas, quoteDenomDepositDeltas DepositDeltas) {
	traderBaseDepositDelta := &DepositDelta{
		AvailableBalanceDelta: e.BaseRefundAmount,
		TotalBalanceDelta:     e.BaseChangeAmount,
	}

	traderQuoteDepositDelta := &DepositDelta{
		AvailableBalanceDelta: e.QuoteRefundAmount,
		TotalBalanceDelta:     e.QuoteChangeAmount,
	}

	// increment availableBalanceDelta in tandem with TotalBalanceDelta if positive
	if e.BaseChangeAmount.IsPositive() {
		traderBaseDepositDelta.AddAvailableBalance(e.BaseChangeAmount)
	}

	// increment availableBalanceDelta in tandem with TotalBalanceDelta if positive
	if e.QuoteChangeAmount.IsPositive() {
		traderQuoteDepositDelta.AddAvailableBalance(e.QuoteChangeAmount)
	}

	feeRecipientSubaccount := EthAddressToSubaccountID(e.FeeRecipient)
	if bytes.Equal(feeRecipientSubaccount.Bytes(), ZeroSubaccountID.Bytes()) {
		feeRecipientSubaccount = AuctionSubaccountID
	}

	// update trader's base and quote balances
	baseDenomDepositDeltas.ApplyDepositDelta(e.SubaccountID, traderBaseDepositDelta)
	quoteDenomDepositDeltas.ApplyDepositDelta(e.SubaccountID, traderQuoteDepositDelta)

	// increment fee recipient's balances
	quoteDenomDepositDeltas.ApplyUniformDelta(feeRecipientSubaccount, e.FeeRecipientReward)

	// increment auction fee balance
	quoteDenomDepositDeltas.ApplyUniformDelta(AuctionSubaccountID, e.AuctionFeeReward)

}

func getSpotLimitSellStateExpansion(
	order *SpotLimitOrder,
	fillQuantity, fillPrice, tradeFeeRate, relayerFeeShare sdk.Dec,
) *SpotOrderStateExpansion {
	orderNotional := fillQuantity.Mul(fillPrice)

	tradingFee := orderNotional.Mul(tradeFeeRate)
	feeRecipientReward := relayerFeeShare.Mul(tradingFee)
	auctionFeeReward := tradingFee.Sub(feeRecipientReward)

	// limit sells are credited with the (fillQuantity * price) * (1 - tradeFeeRate) in quote denom
	quoteChangeAmount := orderNotional.Sub(tradingFee)
	order.Fillable = order.Fillable.Sub(fillQuantity)

	stateExpansion := SpotOrderStateExpansion{
		// limit sells are debited by fillQuantity in base denom
		BaseChangeAmount:       fillQuantity.Neg(),
		BaseRefundAmount:       sdk.ZeroDec(),
		QuoteChangeAmount:      quoteChangeAmount,
		QuoteRefundAmount:      sdk.ZeroDec(),
		FeeRecipient:           common.HexToAddress(order.OrderInfo.FeeRecipient),
		FeeRecipientReward:     feeRecipientReward,
		AuctionFeeReward:       auctionFeeReward,
		LimitOrder:             order,
		LimitOrderFillQuantity: fillQuantity,
		OrderPrice:             order.OrderInfo.Price,
		OrderHash:              common.BytesToHash(order.OrderHash),
		SubaccountID:           common.HexToHash(order.OrderInfo.SubaccountId),
	}
	return &stateExpansion
}

func getRestingSpotLimitBuyStateExpansion(
	order *SpotLimitOrder,
	orderHash common.Hash,
	fillQuantity, fillPrice, makerFeeRate, relayerFeeShare sdk.Dec,
) *SpotOrderStateExpansion {
	var baseChangeAmount, quoteChangeAmount sdk.Dec
	orderNotional := fillQuantity.Mul(fillPrice)
	tradingFee := orderNotional.Mul(makerFeeRate)
	feeRecipientReward := relayerFeeShare.Mul(tradingFee)
	auctionFeeReward := tradingFee.Sub(feeRecipientReward)
	// limit buys are credited with the order fill quantity in base denom
	baseChangeAmount = fillQuantity
	// limit buys are debited with (fillQuantity * Price) * (1 + makerFee) in quote denom
	quoteChangeAmount = orderNotional.Add(tradingFee).Neg()
	quoteRefund := sdk.ZeroDec()

	if !fillPrice.Equal(order.OrderInfo.Price) {
		priceDelta := order.OrderInfo.Price.Sub(fillPrice)
		clearingRefund := fillQuantity.Mul(priceDelta)
		matchedFeeRefund := fillQuantity.Mul(makerFeeRate).Mul(priceDelta)
		quoteRefund = clearingRefund.Add(matchedFeeRefund)
	}
	order.Fillable = order.Fillable.Sub(fillQuantity)
	stateExpansion := SpotOrderStateExpansion{
		BaseChangeAmount:       baseChangeAmount,
		BaseRefundAmount:       sdk.ZeroDec(),
		QuoteChangeAmount:      quoteChangeAmount,
		QuoteRefundAmount:      quoteRefund,
		FeeRecipient:           common.HexToAddress(order.OrderInfo.FeeRecipient),
		FeeRecipientReward:     feeRecipientReward,
		AuctionFeeReward:       auctionFeeReward,
		LimitOrder:             order,
		LimitOrderFillQuantity: fillQuantity,
		OrderPrice:             order.OrderInfo.Price,
		OrderHash:              orderHash,
		SubaccountID:           order.SubaccountID(),
	}
	return &stateExpansion
}

func getNewSpotLimitBuyStateExpansion(
	buyOrder *SpotLimitOrder,
	orderHash common.Hash,
	clearingPrice, fillQuantity,
	makerFeeRate, takerFeeRate, relayerFeeShare sdk.Dec,
) *SpotOrderStateExpansion {
	// TODO: optimize for the case when fillQuantity is 0
	var baseChangeAmount, quoteChangeAmount sdk.Dec

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
	buyOrder.Fillable = buyOrder.Fillable.Sub(fillQuantity)
	stateExpansion := SpotOrderStateExpansion{
		BaseChangeAmount:       baseChangeAmount,
		BaseRefundAmount:       sdk.ZeroDec(),
		QuoteChangeAmount:      quoteChangeAmount,
		QuoteRefundAmount:      quoteRefundAmount,
		FeeRecipient:           common.HexToAddress(buyOrder.OrderInfo.FeeRecipient),
		FeeRecipientReward:     feeRecipientReward,
		AuctionFeeReward:       auctionFeeReward,
		LimitOrder:             buyOrder,
		LimitOrderFillQuantity: fillQuantity,
		OrderPrice:             buyOrder.OrderInfo.Price,
		OrderHash:              orderHash,
		SubaccountID:           buyOrder.SubaccountID(),
	}
	return &stateExpansion
}

func getSpotMarketOrderStateExpansion(
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
	baseRefundAmount, quoteRefundAmount, quoteChangeAmount := sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()
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
		// base denom refund unfilled market order quantity
		if fillQuantity.LT(marketOrder.OrderInfo.Quantity) {
			baseRefundAmount = marketOrder.OrderInfo.Quantity.Sub(fillQuantity)
		}
	}

	stateExpansion := SpotOrderStateExpansion{
		BaseChangeAmount:        baseChangeAmount,
		BaseRefundAmount:        baseRefundAmount,
		QuoteChangeAmount:       quoteChangeAmount,
		QuoteRefundAmount:       quoteRefundAmount,
		FeeRecipient:            common.HexToAddress(marketOrder.OrderInfo.FeeRecipient),
		FeeRecipientReward:      feeRecipientReward,
		AuctionFeeReward:        auctionFeeReward,
		MarketOrder:             marketOrder,
		MarketOrderFillQuantity: fillQuantity,
		OrderPrice:              marketOrder.OrderInfo.Price,
		OrderHash:               common.BytesToHash(marketOrder.OrderHash),
		SubaccountID:            marketOrder.SubaccountID(),
	}
	return &stateExpansion
}
