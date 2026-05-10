package v2

import (
	"cosmossdk.io/math"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	v1 "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

func (p *DerivativePosition) Copy() *DerivativePosition {
	return &DerivativePosition{
		SubaccountId: p.SubaccountId,
		MarketId:     p.MarketId,
		Position:     p.Position.Copy(),
	}
}

func (m *PositionDelta) IsShort() bool { return !m.IsLong }

// NewPosition initializes a new position with a given cumulativeFundingEntry (should be nil for non-perpetual markets)
func NewPosition(isLong bool, cumulativeFundingEntry math.LegacyDec) *Position {
	position := &Position{
		IsLong:     isLong,
		Quantity:   math.LegacyZeroDec(),
		EntryPrice: math.LegacyZeroDec(),
		Margin:     math.LegacyZeroDec(),
	}
	if !cumulativeFundingEntry.IsNil() {
		position.CumulativeFundingEntry = cumulativeFundingEntry
	}
	return position
}

func (p *Position) IsShort() bool { return !p.IsLong }

func (p *Position) Copy() *Position {
	c := &Position{
		IsLong:   p.IsLong,
		Quantity: p.Quantity.Clone(),
	}
	if !p.EntryPrice.IsNil() {
		c.EntryPrice = p.EntryPrice.Clone()
	}
	if !p.Margin.IsNil() {
		c.Margin = p.Margin.Clone()
	}
	if !p.CumulativeFundingEntry.IsNil() {
		c.CumulativeFundingEntry = p.CumulativeFundingEntry.Clone()
	}
	return c
}

// GetEffectiveMarginRatio returns the effective margin ratio of the position, based on the input closing price.
// CONTRACT: position must already be funding-adjusted (if perpetual) and have positive quantity.
func (p *Position) GetEffectiveMarginRatio(closingPrice, closingFee math.LegacyDec) (marginRatio math.LegacyDec) {
	// nolint:all
	// marginRatio = (margin + quantity * PnlPerContract) / (closingPrice * quantity)
	effectiveMargin := p.Margin.Add(p.GetPayoutFromPnl(closingPrice, p.Quantity)).Sub(closingFee)
	return effectiveMargin.Quo(closingPrice.Mul(p.Quantity))
}

// HaircutProfitBasis returns the per-position contribution to TotalProfits used
// during scheduled-settlement profit-haircut accounting. The basis is the
// payout-capped post-fee PnL: `min(post_fee_pnl, post_fee_payout)` when both
// are positive, else zero. This is the SINGLE source of truth used by both
// the scan path (`getPositionFundsStatus`) and the apply path
// (`ApplyProfitHaircutForDerivatives`); they MUST NOT recompute it
// independently or use a different basis (e.g. gross PnL).
//
// The realised-PnL cap applies uniformly across all paths that distribute haircut.
func (p *Position) HaircutProfitBasis(settlementPrice, closingFeeRate math.LegacyDec) math.LegacyDec {
	payout := p.GetPayoutIfFullyClosing(settlementPrice, closingFeeRate)
	if !payout.PnlNotional.IsPositive() || !payout.Payout.IsPositive() {
		return math.LegacyZeroDec()
	}
	return math.LegacyMinDec(payout.PnlNotional, payout.Payout)
}

// ApplyProfitHaircutForDerivatives reduces the position's payout by exactly
// `(deficitAmount / totalProfits) * HaircutProfitBasis(settlementPrice, closingFeeRate)`,
// achieved by an entry-price adjustment. The basis is the capped post-fee PnL
// — see `HaircutProfitBasis`.
//
// Formula:
//
//	rate          = deficitAmount / totalProfits
//	cappedBasis   = HaircutProfitBasis(settlementPrice, closingFeeRate)
//	haircut       = rate * cappedBasis
//	newPayoutFromPnl(gross) = oldPayoutFromPnl(gross) - haircut
//	=> entry-price adjustment for longs: newEntry = entry + haircut/quantity
//	   entry-price adjustment for shorts: newEntry = entry - haircut/quantity
//
// (Both directions move entry "toward" settlement by `haircut/quantity`.)
//
// Rejected historical form: scaling gross PnL by `1 - rate` (i.e.
// `newPayoutFromPnl = oldPayoutFromPnl * (1 - rate)`). For a negative-margin
// winner with `pnl=10, margin=-5, payout=5, rate=0.5`, gross-PnL scaling
// produces post-haircut payout 0; capped-basis scaling produces 2.5, matching
// the deficit-conservation invariant that sum of per-position haircuts equals
// the ProfitHaircutAmount.
func (p *Position) ApplyProfitHaircutForDerivatives(
	deficitAmount, totalProfits, settlementPrice, closingFeeRate math.LegacyDec,
) {
	cappedBasis := p.HaircutProfitBasis(settlementPrice, closingFeeRate)
	if !cappedBasis.IsPositive() {
		return
	}
	haircutAmount := deficitAmount.Mul(cappedBasis).Quo(totalProfits)
	if !haircutAmount.IsPositive() || p.Quantity.IsZero() {
		return
	}

	entryDelta := haircutAmount.Quo(p.Quantity)
	if p.IsLong {
		p.EntryPrice = p.EntryPrice.Add(entryDelta)
	} else {
		p.EntryPrice = p.EntryPrice.Sub(entryDelta)
	}

	// Defensive: the capped-basis formula cannot produce a negative payout
	// (haircut <= cappedBasis <= payout). The clamp is retained as a guard
	// against rounding-driven sub-atom drift.
	newPositionPayout := p.GetPayoutIfFullyClosing(settlementPrice, closingFeeRate).Payout
	if newPositionPayout.IsNegative() {
		p.Margin = p.Margin.Add(newPositionPayout.Abs())
	}
}

func (p *Position) ApplyTotalPositionPayoutHaircut(deficitAmount, totalPayouts, settlementPrice, closingFeeRate math.LegacyDec) {
	if totalPayouts.IsZero() {
		return
	}

	payoutBefore := p.GetPayoutIfFullyClosing(settlementPrice, closingFeeRate).Payout
	if !payoutBefore.IsPositive() {
		return
	}

	removedMargin := payoutBefore.Mul(deficitAmount).Quo(totalPayouts)
	p.Margin = p.Margin.Sub(removedMargin)
}

func (p *Position) ApplyProfitHaircutForBinaryOptions(deficitAmount, totalAssets math.LegacyDec, oracleScaleFactor uint32) {
	// haircutPercentage = deficitAmount / totalAssets
	// To preserve precision, the division by totalAssets is done last.
	// newMargin =  p.Margin - p.Margin * haircutPercentage
	newMargin := p.Margin.Sub(deficitAmount.Mul(p.Margin).Quo(totalAssets))
	p.Margin = newMargin

	// updating entry price just for consistency, but it has no effect since applied haircut is on margin, not on entry price during binary options refunds
	if p.IsLong {
		p.EntryPrice = p.Margin.Quo(p.Quantity)
	} else {
		scaledOne := types.GetScaledPrice(math.LegacyOneDec(), oracleScaleFactor)
		p.EntryPrice = scaledOne.Sub(p.Margin.Quo(p.Quantity))
	}
}

func (p *Position) ClosePositionWithSettlePrice(settlementPrice, closingFeeRate math.LegacyDec) (
	payout, closeTradingFee math.LegacyDec, positionDelta *PositionDelta, pnl math.LegacyDec,
) {
	closingDirection := !p.IsLong
	fullyClosingQuantity := p.Quantity

	closeTradingFee = settlementPrice.Mul(fullyClosingQuantity).Mul(closingFeeRate)
	positionDelta = &PositionDelta{
		IsLong:            closingDirection,
		ExecutionQuantity: fullyClosingQuantity,
		ExecutionMargin:   math.LegacyZeroDec(),
		ExecutionPrice:    settlementPrice,
	}

	// there should not be positions with 0 quantity
	if fullyClosingQuantity.IsZero() {
		return math.LegacyZeroDec(), closeTradingFee, positionDelta, math.LegacyZeroDec()
	}

	payout, _, _, pnl = p.ApplyPositionDelta(positionDelta, closeTradingFee)

	return payout, closeTradingFee, positionDelta, pnl
}

func (p *Position) ClosePositionWithoutPayouts() {
	p.IsLong = false
	p.EntryPrice = math.LegacyZeroDec()
	p.Quantity = math.LegacyZeroDec()
	p.Margin = math.LegacyZeroDec()
	p.CumulativeFundingEntry = math.LegacyZeroDec()
}

func (p *Position) ClosePositionByRefunding(closingFeeRate math.LegacyDec) (
	payout, closeTradingFee math.LegacyDec, positionDelta *PositionDelta, pnl math.LegacyDec,
) {
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
	marketType v1.MarketType,
	reducePrice math.LegacyDec,
	isBuyOrder bool,
	tradeFeeRate math.LegacyDec,
	funding *PerpetualMarketFunding,
	orderMargin math.LegacyDec,
) error {
	if isBuyOrder == p.IsLong {
		return types.ErrInvalidReduceOnlyPositionDirection
	}

	if marketType == v1.MarketType_BinaryOption {
		return nil
	}

	if err := p.checkValidClosingPrice(reducePrice, tradeFeeRate, funding, orderMargin); err != nil {
		return err
	}

	return nil
}

func (p *Position) checkValidClosingPrice(
	closingPrice, tradeFeeRate math.LegacyDec, funding *PerpetualMarketFunding, orderMargin math.LegacyDec,
) error {
	bankruptcyPrice := p.GetBankruptcyPriceWithAddedMargin(funding, orderMargin)

	if p.IsLong {
		// For long positions, Price ≥ BankruptcyPrice / (1 - TradeFeeRate) must hold
		feeAdjustedBankruptcyPrice := bankruptcyPrice.Quo(math.LegacyOneDec().Sub(tradeFeeRate))

		if closingPrice.LT(feeAdjustedBankruptcyPrice) {
			return types.ErrPriceSurpassesBankruptcyPrice
		}
	} else {
		// For short positions, Price ≤ BankruptcyPrice / (1 + TradeFeeRate) must hold
		feeAdjustedBankruptcyPrice := bankruptcyPrice.Quo(math.LegacyOneDec().Add(tradeFeeRate))

		if closingPrice.GT(feeAdjustedBankruptcyPrice) {
			return types.ErrPriceSurpassesBankruptcyPrice
		}
	}
	return nil
}

func (p *Position) GetLiquidationMarketOrderWorstPrice(markPrice math.LegacyDec, funding *PerpetualMarketFunding) *math.LegacyDec {
	bankruptcyPrice := p.GetBankruptcyPrice(funding)
	hasNegativeEquity := (p.IsLong && markPrice.LT(bankruptcyPrice)) || (p.IsShort() && markPrice.GT(bankruptcyPrice))
	if hasNegativeEquity {
		return &markPrice
	}

	return &bankruptcyPrice
}

func (p *Position) GetBankruptcyPrice(funding *PerpetualMarketFunding) (bankruptcyPrice math.LegacyDec) {
	return p.GetLiquidationPrice(math.LegacyZeroDec(), funding)
}

func (p *Position) GetBankruptcyPriceWithAddedMargin(
	funding *PerpetualMarketFunding, addedMargin math.LegacyDec,
) (bankruptcyPrice math.LegacyDec) {
	return p.getLiquidationPriceWithAddedMargin(math.LegacyZeroDec(), funding, addedMargin)
}

func (p *Position) GetLiquidationPrice(maintenanceMarginRatio math.LegacyDec, funding *PerpetualMarketFunding) math.LegacyDec {
	return p.getLiquidationPriceWithAddedMargin(maintenanceMarginRatio, funding, math.LegacyZeroDec())
}

func (p *Position) getLiquidationPriceWithAddedMargin(
	maintenanceMarginRatio math.LegacyDec, funding *PerpetualMarketFunding, addedMargin math.LegacyDec,
) math.LegacyDec {
	adjustedUnitMargin := p.getFundingAdjustedUnitMarginWithAddedMargin(funding, addedMargin)

	// TODO include closing fee for reduce only ?

	var liquidationPrice math.LegacyDec
	if p.IsLong {
		// liquidation price = (entry price - unit margin) / (1 - maintenanceMarginRatio)
		liquidationPrice = p.EntryPrice.Sub(adjustedUnitMargin).Quo(math.LegacyOneDec().Sub(maintenanceMarginRatio))
	} else {
		// liquidation price = (entry price + unit margin) / (1 + maintenanceMarginRatio)
		liquidationPrice = p.EntryPrice.Add(adjustedUnitMargin).Quo(math.LegacyOneDec().Add(maintenanceMarginRatio))
	}
	return liquidationPrice
}

func (p *Position) GetEffectiveMargin(funding *PerpetualMarketFunding, closingPrice math.LegacyDec) math.LegacyDec {
	fundingAdjustedMargin := p.Margin
	if funding != nil {
		fundingAdjustedMargin = p.getFundingAdjustedMargin(funding)
	}
	pnlNotional := math.LegacyZeroDec()
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

func (p *Position) getFundingAdjustedMargin(funding *PerpetualMarketFunding) math.LegacyDec {
	return p.getFundingAdjustedMarginWithAddedMargin(funding, math.LegacyZeroDec())
}

func (p *Position) getFundingAdjustedMarginWithAddedMargin(funding *PerpetualMarketFunding, addedMargin math.LegacyDec) math.LegacyDec {
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

func (p *Position) getFundingAdjustedUnitMarginWithAddedMargin(funding *PerpetualMarketFunding, addedMargin math.LegacyDec) math.LegacyDec {
	adjustedMargin := p.getFundingAdjustedMarginWithAddedMargin(funding, addedMargin)

	// Unit Margin = PositionMargin / PositionQuantity
	fundingAdjustedUnitMargin := adjustedMargin.Quo(p.Quantity)
	return fundingAdjustedUnitMargin
}

func (p *Position) GetAverageWeightedEntryPrice(executionQuantity, executionPrice math.LegacyDec) math.LegacyDec {
	num := p.Quantity.Mul(p.EntryPrice).Add(executionQuantity.Mul(executionPrice))
	denom := p.Quantity.Add(executionQuantity)

	return num.Quo(denom)
}

func (p *Position) GetPayoutIfFullyClosing(closingPrice, closingFeeRate math.LegacyDec) *v1.PositionPayout {
	isProfitable := (p.IsLong && p.EntryPrice.LT(closingPrice)) || (!p.IsLong && p.EntryPrice.GT(closingPrice))

	fullyClosingQuantity := p.Quantity
	positionMargin := p.Margin

	closeTradingFee := closingPrice.Mul(fullyClosingQuantity).Mul(closingFeeRate)
	payoutFromPnl := p.GetPayoutFromPnl(closingPrice, fullyClosingQuantity)
	pnlNotional := payoutFromPnl.Sub(closeTradingFee)
	payout := pnlNotional.Add(positionMargin)

	return &v1.PositionPayout{
		Payout:       payout,
		PnlNotional:  pnlNotional,
		IsProfitable: isProfitable,
	}
}

func (p *Position) GetPayoutFromPnl(closingPrice, closingQuantity math.LegacyDec) math.LegacyDec {
	var pnlNotional math.LegacyDec

	if p.IsLong {
		// nolint:all
		// pnl = closingQuantity * (executionPrice - entryPrice)
		pnlNotional = closingQuantity.Mul(closingPrice.Sub(p.EntryPrice))
	} else {
		// nolint:all
		// pnl = -closingQuantity * (executionPrice - entryPrice)
		pnlNotional = closingQuantity.Mul(closingPrice.Sub(p.EntryPrice)).Neg()
	}

	return pnlNotional
}

func splitPositionMargin(totalMargin, totalQuantity, closingQuantity math.LegacyDec) (
	closingMargin, remainingMargin math.LegacyDec,
) {
	if totalQuantity.IsZero() || closingQuantity.IsZero() {
		return math.LegacyZeroDec(), totalMargin
	}

	if closingQuantity.Equal(totalQuantity) {
		return totalMargin, math.LegacyZeroDec()
	}

	remainingQuantity := totalQuantity.Sub(closingQuantity)
	remainingMargin = totalMargin.Mul(remainingQuantity).Quo(totalQuantity)
	closingMargin = totalMargin.Sub(remainingMargin)

	return closingMargin, remainingMargin
}

// ApplyBankruptCloseWithoutPayouts closes up to closingQuantity at closingPrice with an explicit
// zero payout, while preserving the remaining position state.
func (p *Position) ApplyBankruptCloseWithoutPayouts(
	closingPrice, closingQuantity math.LegacyDec,
) (pnl math.LegacyDec, positionDelta *PositionDelta) {
	if p == nil {
		return math.LegacyZeroDec(), nil
	}

	if p.Quantity.IsZero() || closingQuantity.IsZero() {
		return math.LegacyZeroDec(), &PositionDelta{
			IsLong:            !p.IsLong,
			ExecutionQuantity: math.LegacyZeroDec(),
			ExecutionMargin:   math.LegacyZeroDec(),
			ExecutionPrice:    closingPrice,
		}
	}

	closingQuantity = math.LegacyMinDec(p.Quantity, closingQuantity)
	positionDelta = &PositionDelta{
		IsLong:            !p.IsLong,
		ExecutionQuantity: closingQuantity,
		ExecutionMargin:   math.LegacyZeroDec(),
		ExecutionPrice:    closingPrice,
	}

	pnl = p.GetPayoutFromPnl(closingPrice, closingQuantity)
	remainingMargin := p.Margin.Add(pnl)
	remainingQuantity := p.Quantity.Sub(closingQuantity)

	if remainingQuantity.IsZero() {
		p.ClosePositionWithoutPayouts()
		return pnl, positionDelta
	}

	p.Quantity = remainingQuantity
	p.Margin = remainingMargin

	return pnl, positionDelta
}

func (p *Position) ApplyPositionDelta(delta *PositionDelta, tradingFeeForReduceOnly math.LegacyDec) (
	payout, closeExecutionMargin, collateralizationMargin, pnl math.LegacyDec,
) {
	// No payouts or margin changes if the position delta is nil
	if delta == nil || p == nil {
		return math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()
	}

	if p.Quantity.IsZero() {
		p.IsLong = delta.IsLong
	}

	payout, closeExecutionMargin, collateralizationMargin = math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()
	isNettingInSameDirection := (p.IsLong && delta.IsLong) || (p.IsShort() && delta.IsShort())

	if isNettingInSameDirection {
		p.EntryPrice = p.GetAverageWeightedEntryPrice(delta.ExecutionQuantity, delta.ExecutionPrice)
		p.Quantity = p.Quantity.Add(delta.ExecutionQuantity)
		p.Margin = p.Margin.Add(delta.ExecutionMargin)
		collateralizationMargin = delta.ExecutionMargin

		return payout, closeExecutionMargin, collateralizationMargin, math.LegacyZeroDec()
	}

	// netting in opposing direction
	closingQuantity := math.LegacyMinDec(p.Quantity, delta.ExecutionQuantity)
	// closeExecutionMargin = execution margin * closing quantity / execution quantity
	closeExecutionMargin = delta.ExecutionMargin.Mul(closingQuantity).Quo(delta.ExecutionQuantity)

	pnlNotional := p.GetPayoutFromPnl(delta.ExecutionPrice, closingQuantity)
	isReduceOnlyTrade := delta.ExecutionMargin.IsZero()

	if isReduceOnlyTrade {
		// deduct fees from PNL (position margin) for reduce-only orders

		// only use the closing trading fee for now
		pnlNotional = pnlNotional.Sub(tradingFeeForReduceOnly)
	}

	positionClosingMargin, remainingMargin := splitPositionMargin(p.Margin, p.Quantity, closingQuantity)
	payout = pnlNotional.Add(positionClosingMargin)

	// for netting opposite direction
	newPositionQuantity := p.Quantity.Sub(closingQuantity)
	p.Margin = remainingMargin
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
		_, _, collateralizationMargin, _ = p.ApplyPositionDelta(newPositionDelta, tradingFeeForReduceOnly)
	}

	return payout, closeExecutionMargin, collateralizationMargin, pnlNotional
}
