package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// cumulativeFundingEntry should be nil for non-perpetual markets
func NewPosition(isLong bool, cumulativeFundingEntry sdk.Dec) *Position {
	position := &Position{
		IsLong:       isLong,
		Quantity:     sdk.ZeroDec(),
		EntryPrice:   sdk.ZeroDec(),
		Margin:       sdk.ZeroDec(),
		HoldQuantity: sdk.ZeroDec(),
	}
	if !cumulativeFundingEntry.IsNil() {
		position.CumulativeFundingEntry = cumulativeFundingEntry
	}
	return position
}

// GetUpdatedPositionState updates the position to account for any funding payment and returns a PositionState
func (p *Position) GetUpdatedPositionState(fundingState *PerpetualMarketFundingState) *PositionState {
	fundingPayment := sdk.ZeroDec()

	if fundingState != nil {
		fundingPayment = p.Quantity.Mul(p.CumulativeFundingEntry.Sub(fundingState.CumulativeFunding))
		// For longs, Margin -= Funding
		// For shorts, Margin += Funding
		if p.IsLong {
			fundingPayment = fundingPayment.Neg()
		}
		p.Margin = p.Margin.Add(fundingPayment)
	}
	positionState := PositionState{
		Position:       p,
		FundingPayment: fundingPayment,
	}
	return &positionState
}

func (p *Position) GetAverageWeightedEntryPrice(executionQuantity, executionPrice sdk.Dec) sdk.Dec {
	num := p.Quantity.Mul(p.EntryPrice).Add(executionQuantity.Mul(executionPrice))
	denom := p.Quantity.Add(executionQuantity)
	return num.Quo(denom)
}

func (p *Position) ApplyPositionDelta(delta *PositionDelta, tradeFeeRate sdk.Dec) (
	payout, closeExecutionMargin, collateralizationMargin sdk.Dec,
) {
	// No payouts or margin changes if the position delta is nil
	if delta == nil || p == nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()
	}

	if p.Quantity.IsZero() {
		p.IsLong = delta.IsLong
	}

	payout, closeExecutionMargin, collateralizationMargin = sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()
	isNettingInSameDirection := (p.IsLong && delta.IsLong) || (p.IsShort() && delta.IsShort())

	if isNettingInSameDirection {
		p.Quantity = p.Quantity.Add(delta.ExecutionQuantity)
		p.Margin = p.Margin.Add(delta.ExecutionMargin)
		p.EntryPrice = p.GetAverageWeightedEntryPrice(delta.ExecutionQuantity, delta.ExecutionPrice)
		collateralizationMargin = delta.ExecutionMargin

		return payout, closeExecutionMargin, collateralizationMargin
	}

	// netting in opposing direction
	closingQuantity := sdk.MinDec(p.Quantity, delta.ExecutionQuantity)
	// closeExecutionMargin = execution margin * closing quantity / execution quantity
	closeExecutionMargin = delta.ExecutionMargin.Mul(closingQuantity).Quo(delta.ExecutionQuantity)

	var pnl sdk.Dec
	if p.IsLong {
		// pnl = closingQuantity * (executionPrice - entryPrice)
		pnl = closingQuantity.Mul(delta.ExecutionPrice.Sub(p.EntryPrice))
	} else {
		// pnl = -closingQuantity * (executionPrice - entryPrice)
		pnl = closingQuantity.Mul(delta.ExecutionPrice.Sub(p.EntryPrice)).Neg()
	}

	isReduceOnlyTrade := delta.ExecutionMargin.IsZero()

	if isReduceOnlyTrade {
		// deduct fees from PNL (position margin) for reduce-only orders

		// only compute the closing trading fee for now
		tradingFee := delta.ExecutionPrice.Mul(closingQuantity).Mul(tradeFeeRate)
		pnl = pnl.Sub(tradingFee)
		p.HoldQuantity = p.HoldQuantity.Sub(closingQuantity)
	}

	positionClosingMargin := p.Margin.Mul(closingQuantity).Quo(p.Quantity)
	payout = pnl.Add(positionClosingMargin)

	// for netting opposite direction
	newPositionQuantity := p.Quantity.Sub(closingQuantity)
	p.Margin = p.Margin.Mul(newPositionQuantity).Quo(p.Quantity)
	p.Quantity = newPositionQuantity

	isFlippingPosition := delta.ExecutionQuantity.GT(closingQuantity)

	if isFlippingPosition {
		remainingExecutionQuantity := delta.ExecutionQuantity.Sub(closingQuantity)
		remainingExecutionMargin := delta.ExecutionMargin.Sub(closeExecutionMargin)
		newPositionDelta := &PositionDelta{
			IsLong:            !p.IsLong,
			ExecutionQuantity: remainingExecutionQuantity,
			ExecutionMargin:   remainingExecutionMargin,
			ExecutionPrice:    delta.ExecutionPrice,
		}
		// recurse
		_, _, collateralizationMargin = p.ApplyPositionDelta(newPositionDelta, tradeFeeRate)
	}

	return payout, closeExecutionMargin, collateralizationMargin
}
