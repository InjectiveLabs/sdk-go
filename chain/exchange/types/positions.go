package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type positionPayout struct {
	Payout       sdk.Dec
	PnlNotional  sdk.Dec
	IsProfitable bool
}

func (m *Position) IsShort() bool { return !m.IsLong }

func (m *PositionDelta) IsShort() bool { return !m.IsLong }

// NewPosition initializes a new position with a given cumulativeFundingEntry (should be nil for non-perpetual markets)
func NewPosition(isLong bool, cumulativeFundingEntry sdk.Dec) *Position {
	position := &Position{
		IsLong:     isLong,
		Quantity:   sdk.ZeroDec(),
		EntryPrice: sdk.ZeroDec(),
		Margin:     sdk.ZeroDec(),
	}
	if !cumulativeFundingEntry.IsNil() {
		position.CumulativeFundingEntry = cumulativeFundingEntry
	}
	return position
}

// GetEffectiveMarginRatio returns the effective margin ratio of the position, based on the input mark price.
// CONTRACT: position must already be funding-adjusted (if perpetual) and have positive quantity.
func (p *Position) GetEffectiveMarginRatio(markPrice sdk.Dec) (marginRatio sdk.Dec) {
	// marginRatio = (margin + quantity * PnlPerContract) / (markPrice * quantity)
	effectiveMargin := p.Margin.Add(p.GetPayoutFromPnl(markPrice, p.Quantity))
	return effectiveMargin.Quo(markPrice.Mul(p.Quantity))
}

// ApplyProfitHaircut results in reducing the payout (pnl * quantity) by the given rate (e.g. 0.1=10%) by modifying the entry price.
// Formula for adjustment:
// newPayoutFromPnl = oldPayoutFromPnl * (1 - missingFundsRate)
// => Entry price adjustment for buys
// (newEntryPrice - settlementPrice) * quantity = (entryPrice - settlementPrice) * quantity * (1 - missingFundsRate)
// newEntryPrice = entryPrice - entryPrice * haircutPercentage + settlementPrice * haircutPercentage
// => Entry price adjustment for sells
// (settlementPrice - newEntryPrice) * quantity = (settlementPrice - entryPrice) * quantity * (1 - missingFundsRate)
// newEntryPrice = entryPrice - entryPrice * haircutPercentage + settlementPrice * haircutPercentage
func (p *Position) ApplyProfitHaircut(deficitAmount, totalProfits, settlementPrice sdk.Dec) {
	// haircutPercentage = deficitAmount / totalProfits
	// To preserve precision, the division by totalProfits is done last.
	// newEntryPrice =  haircutPercentage * (settlementPrice - entryPrice) + entryPrice
	newEntryPrice := deficitAmount.Mul(settlementPrice.Sub(p.EntryPrice)).Quo(totalProfits).Add(p.EntryPrice)
	p.EntryPrice = newEntryPrice
}

func (p *Position) ClosePositionWithSettlePrice(settlementPrice, closingFeeRate sdk.Dec) (payout, closeTradingFee sdk.Dec, positionDelta *PositionDelta) {
	closingDirection := !p.IsLong
	fullyClosingQuantity := p.Quantity

	closeTradingFee = settlementPrice.Mul(fullyClosingQuantity).Mul(closingFeeRate)
	positionDelta = &PositionDelta{
		IsLong:            closingDirection,
		ExecutionQuantity: fullyClosingQuantity,
		ExecutionMargin:   sdk.ZeroDec(),
		ExecutionPrice:    settlementPrice,
	}

	// there should not be positions with 0 quantity
	if fullyClosingQuantity.IsZero() {
		return sdk.ZeroDec(), closeTradingFee, positionDelta
	}

	payout, _, _ = p.ApplyPositionDelta(positionDelta, closeTradingFee)

	return payout, closeTradingFee, positionDelta
}

func (p *Position) ClosePositionByRefunding(closingFeeRate sdk.Dec) (payout, closeTradingFee sdk.Dec, positionDelta *PositionDelta) {
	return p.ClosePositionWithSettlePrice(p.EntryPrice, closingFeeRate)
}

func (p *Position) GetDirectionString() string {
	directionStr := "Long"
	if p.IsShort() {
		directionStr = "Short"
	}
	return directionStr
}

func (p *Position) CheckValidPositionToReduce(
	marketType MarketType,
	reducePrice sdk.Dec,
	isBuyOrder bool,
	tradeFeeRate sdk.Dec,
	funding *PerpetualMarketFunding,
	orderMargin sdk.Dec,
) error {
	if isBuyOrder == p.IsLong {
		return ErrInvalidReduceOnlyPositionDirection
	}

	if marketType == MarketType_BinaryOption {
		return nil
	}

	if err := p.checkValidClosingPrice(reducePrice, tradeFeeRate, funding, orderMargin); err != nil {
		return err
	}

	return nil
}

func (p *Position) checkValidClosingPrice(closingPrice, tradeFeeRate sdk.Dec, funding *PerpetualMarketFunding, orderMargin sdk.Dec) error {
	bankruptcyPrice := p.getBankruptcyPriceWithAddedMargin(funding, orderMargin)

	if p.IsLong {
		// For long positions, Price ≥ BankruptcyPrice / (1 - TradeFeeRate) must hold
		feeAdjustedBankruptcyPrice := bankruptcyPrice.Quo(sdk.OneDec().Sub(tradeFeeRate))

		if closingPrice.LT(feeAdjustedBankruptcyPrice) {
			return ErrPriceSurpassesBankruptcyPrice
		}
	} else {
		// For short positions, Price ≤ BankruptcyPrice / (1 + TradeFeeRate) must hold
		feeAdjustedBankruptcyPrice := bankruptcyPrice.Quo(sdk.OneDec().Add(tradeFeeRate))

		if closingPrice.GT(feeAdjustedBankruptcyPrice) {
			return ErrPriceSurpassesBankruptcyPrice
		}
	}
	return nil
}

func (p *Position) GetBankruptcyPrice(funding *PerpetualMarketFunding) (bankruptcyPrice sdk.Dec) {
	return p.GetLiquidationPrice(sdk.ZeroDec(), funding)
}

func (p *Position) getBankruptcyPriceWithAddedMargin(funding *PerpetualMarketFunding, addedMargin sdk.Dec) (bankruptcyPrice sdk.Dec) {
	return p.getLiquidationPriceWithAddedMargin(sdk.ZeroDec(), funding, addedMargin)
}

func (p *Position) GetLiquidationPrice(maintenanceMarginRatio sdk.Dec, funding *PerpetualMarketFunding) sdk.Dec {
	return p.getLiquidationPriceWithAddedMargin(maintenanceMarginRatio, funding, sdk.ZeroDec())
}

func (p *Position) getLiquidationPriceWithAddedMargin(maintenanceMarginRatio sdk.Dec, funding *PerpetualMarketFunding, addedMargin sdk.Dec) sdk.Dec {
	adjustedUnitMargin := p.getFundingAdjustedUnitMarginWithAddedMargin(funding, addedMargin)

	// TODO include closing fee for reduce only ?

	var liquidationPrice sdk.Dec
	if p.IsLong {
		// liquidation price = (entry price - unit margin) / (1 - maintenanceMarginRatio)
		liquidationPrice = p.EntryPrice.Sub(adjustedUnitMargin).Quo(sdk.OneDec().Sub(maintenanceMarginRatio))
	} else {
		// liquidation price = (entry price + unit margin) / (1 + maintenanceMarginRatio)
		liquidationPrice = p.EntryPrice.Add(adjustedUnitMargin).Quo(sdk.OneDec().Add(maintenanceMarginRatio))
	}
	return liquidationPrice
}

func (p *Position) GetEffectiveMargin(funding *PerpetualMarketFunding, closingPrice sdk.Dec) sdk.Dec {
	fundingAdjustedMargin := p.Margin
	if funding != nil {
		fundingAdjustedMargin = p.getFundingAdjustedMargin(funding)
	}
	pnlNotional := sdk.ZeroDec()
	if !closingPrice.IsNil() {
		pnlNotional = p.GetPayoutFromPnl(closingPrice, p.Quantity)
	}
	effectiveMargin := fundingAdjustedMargin.Add(pnlNotional)
	return effectiveMargin
}

// ApplyFunding updates the position to account for any funding payment.
func (p *Position) ApplyFunding(funding *PerpetualMarketFunding) {
	if funding != nil {
		p.Margin = p.getFundingAdjustedMargin(funding)

		// update the cumulative funding entry to current
		p.CumulativeFundingEntry = funding.CumulativeFunding
	}
}

func (p *Position) getFundingAdjustedMargin(funding *PerpetualMarketFunding) sdk.Dec {
	return p.getFundingAdjustedMarginWithAddedMargin(funding, sdk.ZeroDec())
}

func (p *Position) getFundingAdjustedMarginWithAddedMargin(funding *PerpetualMarketFunding, addedMargin sdk.Dec) sdk.Dec {
	adjustedMargin := p.Margin.Add(addedMargin)

	// Compute the adjusted position margin for positions in perpetual markets
	if funding != nil {
		unrealizedFundingPayment := p.Quantity.Mul(funding.CumulativeFunding.Sub(p.CumulativeFundingEntry))

		// For longs, Margin -= Funding
		// For shorts, Margin += Funding
		if p.IsLong {
			adjustedMargin = adjustedMargin.Sub(unrealizedFundingPayment)
		} else {
			adjustedMargin = adjustedMargin.Add(unrealizedFundingPayment)
		}
	}

	return adjustedMargin
}

func (p *Position) getFundingAdjustedUnitMarginWithAddedMargin(funding *PerpetualMarketFunding, addedMargin sdk.Dec) sdk.Dec {
	adjustedMargin := p.getFundingAdjustedMarginWithAddedMargin(funding, addedMargin)

	// Unit Margin = PositionMargin / PositionQuantity
	fundingAdjustedUnitMargin := adjustedMargin.Quo(p.Quantity)
	return fundingAdjustedUnitMargin
}

func (p *Position) GetAverageWeightedEntryPrice(executionQuantity, executionPrice sdk.Dec) sdk.Dec {
	num := p.Quantity.Mul(p.EntryPrice).Add(executionQuantity.Mul(executionPrice))
	denom := p.Quantity.Add(executionQuantity)

	return num.Quo(denom)
}

func (p *Position) GetPayoutIfFullyClosing(closingPrice, closingFeeRate sdk.Dec) *positionPayout {
	isProfitable := (p.IsLong && p.EntryPrice.LT(closingPrice)) || (!p.IsLong && p.EntryPrice.GT(closingPrice))

	fullyClosingQuantity := p.Quantity
	positionMargin := p.Margin

	closeTradingFee := closingPrice.Mul(fullyClosingQuantity).Mul(closingFeeRate)
	payoutFromPnl := p.GetPayoutFromPnl(closingPrice, fullyClosingQuantity)
	pnlNotional := payoutFromPnl.Sub(closeTradingFee)
	payout := pnlNotional.Add(positionMargin)

	return &positionPayout{
		Payout:       payout,
		PnlNotional:  pnlNotional,
		IsProfitable: isProfitable,
	}
}

func (p *Position) GetPayoutFromPnl(closingPrice, closingQuantity sdk.Dec) sdk.Dec {
	var pnlNotional sdk.Dec

	if p.IsLong {
		// pnl = closingQuantity * (executionPrice - entryPrice)
		pnlNotional = closingQuantity.Mul(closingPrice.Sub(p.EntryPrice))
	} else {
		// pnl = -closingQuantity * (executionPrice - entryPrice)
		pnlNotional = closingQuantity.Mul(closingPrice.Sub(p.EntryPrice)).Neg()
	}

	return pnlNotional
}

func (p *Position) ApplyPositionDelta(delta *PositionDelta, tradingFeeForReduceOnly sdk.Dec) (
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
		p.EntryPrice = p.GetAverageWeightedEntryPrice(delta.ExecutionQuantity, delta.ExecutionPrice)
		p.Quantity = p.Quantity.Add(delta.ExecutionQuantity)
		p.Margin = p.Margin.Add(delta.ExecutionMargin)
		collateralizationMargin = delta.ExecutionMargin

		return payout, closeExecutionMargin, collateralizationMargin
	}

	// netting in opposing direction
	closingQuantity := sdk.MinDec(p.Quantity, delta.ExecutionQuantity)
	// closeExecutionMargin = execution margin * closing quantity / execution quantity
	closeExecutionMargin = delta.ExecutionMargin.Mul(closingQuantity).Quo(delta.ExecutionQuantity)

	pnlNotional := p.GetPayoutFromPnl(delta.ExecutionPrice, closingQuantity)
	isReduceOnlyTrade := delta.ExecutionMargin.IsZero()

	if isReduceOnlyTrade {
		// deduct fees from PNL (position margin) for reduce-only orders

		// only use the closing trading fee for now
		pnlNotional = pnlNotional.Sub(tradingFeeForReduceOnly)
	}

	positionClosingMargin := p.Margin.Mul(closingQuantity).Quo(p.Quantity)
	payout = pnlNotional.Add(positionClosingMargin)

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
		_, _, collateralizationMargin = p.ApplyPositionDelta(newPositionDelta, tradingFeeForReduceOnly)
	}

	return payout, closeExecutionMargin, collateralizationMargin
}
