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
	TraderFeeReward         sdk.Dec
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
	makerFeeRate, relayerFeeShareRate sdk.Dec,
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
				relayerFeeShareRate,
			)
		} else {
			stateExpansions[idx] = getSpotLimitSellStateExpansion(
				order,
				fillQuantity,
				fillPrice,
				makerFeeRate,
				relayerFeeShareRate,
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
	feeRecipientReward := relayerFeeShare.Mul(tradingFee).Abs()

	var auctionFeeReward, traderFee sdk.Dec
	if tradingFee.IsNegative() {
		traderFee = tradingFee.Add(feeRecipientReward)
		auctionFeeReward = tradingFee // taker auction fees pay for maker
	} else {
		traderFee = tradingFee
		auctionFeeReward = tradingFee.Sub(feeRecipientReward)
	}

	// limit sells are credited with the (fillQuantity * price) * traderFee in quote denom
	// traderFee can be positive or negative
	quoteChangeAmount := orderNotional.Sub(traderFee)
	order.Fillable = order.Fillable.Sub(fillQuantity)

	stateExpansion := SpotOrderStateExpansion{
		// limit sells are debited by fillQuantity in base denom
		BaseChangeAmount:       fillQuantity.Neg(),
		BaseRefundAmount:       sdk.ZeroDec(),
		QuoteChangeAmount:      quoteChangeAmount,
		QuoteRefundAmount:      sdk.ZeroDec(),
		FeeRecipient:           order.FeeRecipient(),
		FeeRecipientReward:     feeRecipientReward,
		AuctionFeeReward:       auctionFeeReward,
		TraderFeeReward:        traderFee.Abs(),
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
	fillQuantity, fillPrice, makerFeeRate, relayerFeeShareRate sdk.Dec,
) *SpotOrderStateExpansion {
	var baseChangeAmount, quoteChangeAmount, auctionFeeReward, feeRebate sdk.Dec

	orderNotional := fillQuantity.Mul(fillPrice)
	tradingFee := orderNotional.Mul(makerFeeRate)
	// relayer fee reward = relayerFeeShareRate * |trading fee|
	feeRecipientReward := relayerFeeShareRate.Mul(tradingFee.Abs())

	if tradingFee.IsNegative() {
		auctionFeeReward = tradingFee // taker auction fees pay for maker
		feeRebate = tradingFee.Abs().Sub(feeRecipientReward)
	} else {
		// auction fee reward = (1 - relayerFeeShareRate) * trading fee
		auctionFeeReward = tradingFee.Abs().Sub(feeRecipientReward)
		feeRebate = sdk.ZeroDec()
	}

	// limit buys are credited with the order fill quantity in base denom
	baseChangeAmount = fillQuantity

	// limit buys are debited with (fillQuantity * Price) * (1 + makerFee) in quote denom
	if tradingFee.IsNegative() {
		quoteChangeAmount = orderNotional.Neg().Add(feeRebate)
	} else {
		quoteChangeAmount = orderNotional.Add(tradingFee).Neg()
	}

	quoteRefund := feeRebate

	if !fillPrice.Equal(order.OrderInfo.Price) {
		// priceDelta = price - fill price
		priceDelta := order.OrderInfo.Price.Sub(fillPrice)
		// clearingRefund = fillQuantity * priceDelta
		clearingRefund := fillQuantity.Mul(priceDelta)
		// matchedFeeRefund = max(makerFeeRate, 0) * fillQuantity * priceDelta
		matchedFeeRefund := sdk.MaxDec(makerFeeRate, sdk.ZeroDec()).Mul(fillQuantity.Mul(priceDelta))
		// quoteRefund += (1 + max(makerFeeRate, 0)) * fillQuantity * priceDelta
		quoteRefund = quoteRefund.Add(clearingRefund.Add(matchedFeeRefund))
	}
	order.Fillable = order.Fillable.Sub(fillQuantity)


	stateExpansion := SpotOrderStateExpansion{
		BaseChangeAmount:       baseChangeAmount,
		BaseRefundAmount:       sdk.ZeroDec(),
		QuoteChangeAmount:      quoteChangeAmount,
		QuoteRefundAmount:      quoteRefund,
		FeeRecipient:           order.FeeRecipient(),
		FeeRecipientReward:     feeRecipientReward,
		AuctionFeeReward:       auctionFeeReward,
		TraderFeeReward:        feeRebate,
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
	positiveMakerFeePart := sdk.MaxDec(sdk.ZeroDec(), makerFeeRate)
	unmatchedFeeRefund := buyOrder.OrderInfo.Quantity.Sub(fillQuantity).Mul(buyOrder.OrderInfo.Price).Mul(takerFeeRate.Sub(positiveMakerFeePart))
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
		FeeRecipient:           buyOrder.FeeRecipient(),
		FeeRecipientReward:     feeRecipientReward,
		AuctionFeeReward:       auctionFeeReward,
		TraderFeeReward:        sdk.ZeroDec(),
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
		FeeRecipient:            marketOrder.FeeRecipient(),
		FeeRecipientReward:      feeRecipientReward,
		AuctionFeeReward:        auctionFeeReward,
		TraderFeeReward:         sdk.ZeroDec(),
		MarketOrder:             marketOrder,
		MarketOrderFillQuantity: fillQuantity,
		OrderPrice:              marketOrder.OrderInfo.Price,
		OrderHash:               common.BytesToHash(marketOrder.OrderHash),
		SubaccountID:            marketOrder.SubaccountID(),
	}
	return &stateExpansion
}
