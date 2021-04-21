package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type PositionState struct {
	Position       *Position
	FundingPayment sdk.Dec
}

type DerivativeOrderStateExpansion struct {
	SubaccountID  common.Hash
	PositionDelta *PositionDelta
	Payout        sdk.Dec

	TotalBalanceDelta     sdk.Dec
	AvailableBalanceDelta sdk.Dec

	AuctionFeeReward   sdk.Dec
	FeeRecipientReward sdk.Dec
	FeeRecipient       common.Address
	OrderHash          common.Hash
	// For market orders, FillableAmount refers to the fillable quantity of the market order execution (if any)
	FillableAmount sdk.Dec
}

func GetDerivativeOrderBatchEvent(
	isBuy bool,
	executionType ExecutionType,
	market *DerivativeMarket,
	funding *PerpetualMarketFunding,
	stateExpansions []*DerivativeOrderStateExpansion,
	depositDeltas DepositDeltas,
) (batch *EventBatchDerivativeExecution, filledDeltas []*LimitOrderFilledDelta) {
	if len(stateExpansions) == 0 {
		return
	}

	trades := make([]*DerivativeTradeLog, 0, len(stateExpansions))

	if executionType == ExecutionType_LimitMatchRestingOrder || executionType == ExecutionType_LimitFill {
		filledDeltas = make([]*LimitOrderFilledDelta, 0, len(stateExpansions))
	}

	for idx := range stateExpansions {
		expansion := stateExpansions[idx]
		depositDeltas.ApplyDepositDelta(expansion.SubaccountID, &DepositDelta{TotalBalanceDelta: expansion.TotalBalanceDelta, AvailableBalanceDelta: expansion.AvailableBalanceDelta})
		depositDeltas.ApplyUniformDelta(EthAddressToSubaccountID(expansion.FeeRecipient), expansion.FeeRecipientReward)
		depositDeltas.ApplyUniformDelta(ZeroHash, expansion.AuctionFeeReward)

		if executionType == ExecutionType_LimitMatchRestingOrder || executionType == ExecutionType_LimitFill {
			filledDelta := &LimitOrderFilledDelta{
				SubaccountIndexKey: GetLimitOrderIndexKey(market.MarketID(), isBuy, expansion.SubaccountID, expansion.OrderHash),
				FillableAmount:     expansion.FillableAmount,
			}
			filledDeltas = append(filledDeltas, filledDelta)
		}

		if expansion.PositionDelta != nil {
			tradeLog := &DerivativeTradeLog{
				SubaccountId:  expansion.SubaccountID.Bytes(),
				PositionDelta: expansion.PositionDelta,
				Payout:        expansion.Payout,
				Fee:           expansion.FeeRecipientReward.Add(expansion.AuctionFeeReward),
				OrderHash:     expansion.OrderHash.Bytes(),
			}
			trades = append(trades, tradeLog)
		}
	}
	batch = &EventBatchDerivativeExecution{
		MarketId:          market.MarketId,
		IsBuy:             isBuy,
		IsLiquidation:     false,
		CumulativeFunding: &funding.CumulativeFunding,
		ExecutionType:     executionType,
		Trades:            trades,
	}
	return batch, filledDeltas
}

// NOTE: clearingPrice may be Nil
func ProcessDerivativeMarketOrderStateExpansions(
	isMarketBuy bool,
	marketOrders []*DerivativeMarketOrder,
	marketFillQuantities []sdk.Dec,
	positionStates map[common.Hash]*PositionState,
	clearingPrice sdk.Dec,
	tradeFeeRate, relayerFeeShareRate sdk.Dec,
) ([]*DerivativeOrderStateExpansion, []*DerivativeMarketOrderCancel) {
	stateExpansions := make([]*DerivativeOrderStateExpansion, len(marketOrders))
	ordersToCancel := make([]*DerivativeMarketOrderCancel, 0, len(marketOrders))
	for idx := range marketOrders {
		stateExpansions[idx] = GetDerivativeMarketOrderStateExpansion(
			isMarketBuy,
			marketOrders[idx],
			positionStates,
			marketFillQuantities[idx],
			clearingPrice,
			tradeFeeRate,
			relayerFeeShareRate,
		)
		if !stateExpansions[idx].FillableAmount.IsZero() {
			ordersToCancel = append(ordersToCancel, &DerivativeMarketOrderCancel{
				MarketOrder:    marketOrders[idx],
				CancelQuantity: stateExpansions[idx].FillableAmount,
			})
		}
	}
	return stateExpansions, ordersToCancel
}

func GetDerivativeMarketOrderStateExpansion(
	isBuy bool,
	order *DerivativeMarketOrder,
	positionStates map[common.Hash]*PositionState,
	fillQuantity, clearingPrice sdk.Dec,
	takerFeeRate, relayerFeeShareRate sdk.Dec,
) *DerivativeOrderStateExpansion {
	if fillQuantity.IsNil() {
		fillQuantity = sdk.ZeroDec()
	}

	orderFillNotional := fillQuantity.Mul(clearingPrice)
	tradingFee, feeRecipientReward, auctionFeeReward := getOrderFillFeeInfo(orderFillNotional, takerFeeRate, relayerFeeShareRate)
	subaccountID := order.SubaccountID()

	var position *Position
	positionState := positionStates[subaccountID]
	if positionState != nil {
		position = positionState.Position
	}

	// The current approach using the entire order.Margin, regardless of the fillQuantity (which can be less than the order quantity)
	// TODO: Consider using order.Margin.Mul(fillQuantity).Quo(order.OrderInfo.Quantity) for the executionMargin
	// and then refunding the unused margin instead?
	var positionDelta *PositionDelta

	if !fillQuantity.IsZero() {
		positionDelta = &PositionDelta{
			IsLong:            isBuy,
			ExecutionQuantity: fillQuantity,
			ExecutionMargin:   order.Margin,
			ExecutionPrice:    clearingPrice,
		}
	}

	payout, closeExecutionMargin, collateralizationMargin := position.ApplyPositionDelta(positionDelta, takerFeeRate)

	clearingFeeChargeOrRefund := sdk.ZeroDec()
	feeCharge := sdk.ZeroDec()

	if order.IsVanilla() {
		feeCharge = tradingFee
		priceDelta := sdk.ZeroDec()

		if !clearingPrice.IsNil() {
			// ΔPrice = |OrderPrice - ClearingPrice|
			priceDelta = order.OrderInfo.Price.Sub(clearingPrice).Abs()

			// refund or charge = Quantity * ΔPrice * γ_taker
			clearingFeeChargeOrRefund = fillQuantity.Mul(priceDelta).Mul(takerFeeRate)
		}
	}

	totalBalanceChange := payout.Sub(collateralizationMargin.Add(feeCharge))
	availableBalanceChange := payout.Add(closeExecutionMargin)

	if order.IsBuy() {
		availableBalanceChange = availableBalanceChange.Add(clearingFeeChargeOrRefund)
	} else {
		availableBalanceChange = availableBalanceChange.Sub(clearingFeeChargeOrRefund)
	}

	stateExpansion := DerivativeOrderStateExpansion{
		SubaccountID:          subaccountID,
		PositionDelta:         positionDelta,
		Payout:                payout,
		TotalBalanceDelta:     totalBalanceChange,
		AvailableBalanceDelta: availableBalanceChange,
		AuctionFeeReward:      auctionFeeReward,
		FeeRecipientReward:    feeRecipientReward,
		FeeRecipient:          common.HexToAddress(order.OrderInfo.FeeRecipient),
		OrderHash:             common.BytesToHash(order.OrderHash),
		FillableAmount:        order.OrderInfo.Quantity.Sub(fillQuantity),
	}

	return &stateExpansion
}

func getOrderFillFeeInfo(orderFillNotional, tradeFeeRate, relayerFeeShareRate sdk.Dec) (
	tradingFee, feeRecipientReward, auctionFeeReward sdk.Dec,
) {
	tradingFee = orderFillNotional.Mul(tradeFeeRate)
	feeRecipientReward = relayerFeeShareRate.Mul(tradingFee)
	auctionFeeReward = tradingFee.Sub(feeRecipientReward)
	return tradingFee, feeRecipientReward, auctionFeeReward
}

// NOTE: clearingPrice can be nil
func GetDerivativeLimitOrderStateExpansion(
	isBuy bool,
	isTransient bool,
	order *DerivativeLimitOrder,
	positionStates map[common.Hash]*PositionState,
	fillQuantity, clearingPrice sdk.Dec,
	makerFeeRate, takerFeeRate, relayerFeeShareRate sdk.Dec,
) *DerivativeOrderStateExpansion {
	var executionPrice sdk.Dec
	if clearingPrice.IsNil() {
		executionPrice = order.OrderInfo.Price
	} else {
		executionPrice = clearingPrice
	}
	orderFillNotional := fillQuantity.Mul(executionPrice)
	var tradeFeeRate sdk.Dec
	if isTransient {
		tradeFeeRate = takerFeeRate
	} else {
		tradeFeeRate = makerFeeRate
	}
	tradingFee, feeRecipientReward, auctionFeeReward := getOrderFillFeeInfo(orderFillNotional, tradeFeeRate, relayerFeeShareRate)

	subaccountID := order.SubaccountID()
	var position *Position
	positionState := positionStates[subaccountID]
	if positionState != nil {
		position = positionState.Position
	}
	var positionDelta *PositionDelta

	if !fillQuantity.IsZero() {
		executionMargin := order.Margin.Mul(fillQuantity).Quo(order.OrderInfo.Quantity)
		positionDelta = &PositionDelta{
			IsLong:            isBuy,
			ExecutionQuantity: fillQuantity,
			ExecutionMargin:   executionMargin,
			ExecutionPrice:    executionPrice,
		}
	}

	payout, closeExecutionMargin, collateralizationMargin := position.ApplyPositionDelta(positionDelta, takerFeeRate)
	tradingFeeBalanceDebitAmount := sdk.ZeroDec()
	clearingChargeOrRefund := sdk.ZeroDec()
	fillableQuantity := order.OrderInfo.Quantity.Sub(fillQuantity)

	if order.IsVanilla() {
		tradingFeeBalanceDebitAmount = tradingFee

		priceDelta, clearingFeeChargeOrRefund, unmatchedFeeRefund := sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()
		// ΔPrice = |OrderPrice - ClearingPrice|
		priceDelta = order.OrderInfo.Price.Sub(executionPrice).Abs()

		if isTransient {
			clearingFeeChargeOrRefund = fillQuantity.Mul(priceDelta).Mul(takerFeeRate)
			unmatchedFeeRefund = fillableQuantity.Mul(order.OrderInfo.Price).Mul(takerFeeRate.Sub(makerFeeRate))
		} else {
			clearingFeeChargeOrRefund = fillQuantity.Mul(priceDelta).Mul(makerFeeRate)
			unmatchedFeeRefund = sdk.ZeroDec()
		}

		if order.IsBuy() {
			clearingChargeOrRefund = unmatchedFeeRefund.Add(clearingFeeChargeOrRefund)
		} else {
			clearingChargeOrRefund = unmatchedFeeRefund.Sub(clearingFeeChargeOrRefund)
		}
	}

	totalBalanceChange := payout.Sub(collateralizationMargin.Add(tradingFeeBalanceDebitAmount))
	availableBalanceChange := payout.Add(closeExecutionMargin).Add(clearingChargeOrRefund)

	stateExpansion := DerivativeOrderStateExpansion{
		SubaccountID:          subaccountID,
		PositionDelta:         positionDelta,
		Payout:                payout,
		TotalBalanceDelta:     totalBalanceChange,
		AvailableBalanceDelta: availableBalanceChange,
		AuctionFeeReward:      auctionFeeReward,
		FeeRecipientReward:    feeRecipientReward,
		FeeRecipient:          common.HexToAddress(order.OrderInfo.FeeRecipient),
		OrderHash:             common.BytesToHash(order.OrderHash),
		FillableAmount:        fillableQuantity,
	}
	return &stateExpansion
}
