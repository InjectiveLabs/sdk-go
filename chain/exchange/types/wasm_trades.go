package types

import (
	"math/big"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
)

type PositionTransfer struct {
	MarketID                common.Hash    `json:"market_id"`
	SourceSubaccountID      common.Hash    `json:"source_subaccount_id"`
	DestinationSubaccountID common.Hash    `json:"destination_subaccount_id"`
	Quantity                math.LegacyDec `json:"quantity"`
}

// ValidateBasic is version-agnostic, unlike SyntheticTrade.Validate: there is no v1->v2 conversion
// for a position transfer, and a position quantity has the same representation in both versions
// because DerivativeMarket.QuantityFromChainFormat scales by 10^0. The cap means the same thing
// wherever it is checked.
func (p *PositionTransfer) ValidateBasic() error {
	if p.Quantity.IsNil() || !p.Quantity.IsPositive() {
		return ErrInvalidQuantity
	}

	if p.Quantity.GT(MaxOrderQuantity) {
		return errors.Wrap(ErrInvalidQuantity, p.Quantity.String())
	}

	if p.SourceSubaccountID == p.DestinationSubaccountID {
		return ErrBadSubaccountID
	}

	return nil
}

type SyntheticTradeAction struct {
	UserTrades     []*SyntheticTrade `json:"user_trades"`
	ContractTrades []*SyntheticTrade `json:"contract_trades"`
}

func (a *SyntheticTradeAction) ValidateBasic() error {
	if len(a.UserTrades) == 0 || len(a.ContractTrades) == 0 {
		return ErrInvalidTrade
	}

	for _, t := range a.UserTrades {
		if t == nil {
			return ErrInvalidTrade
		}
		if err := t.Validate(); err != nil {
			return err
		}
	}

	for _, t := range a.ContractTrades {
		if t == nil {
			return ErrInvalidTrade
		}
		if err := t.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// ValidateRFQLiquidationBasic applies the dedicated full-pool liquidation wire
// contract. Target legs are pure closes with zero supplied margin; provider legs
// may carry enough isolated collateral for the maximum representable notional.
func (a *SyntheticTradeAction) ValidateRFQLiquidationBasic() error {
	if a == nil || len(a.UserTrades) == 0 || len(a.UserTrades) != len(a.ContractTrades) {
		return ErrInvalidTrade
	}

	for _, trade := range a.UserTrades {
		if err := validateRFQLiquidationTrade(trade, true); err != nil {
			return err
		}
	}
	for _, trade := range a.ContractTrades {
		if err := validateRFQLiquidationTrade(trade, false); err != nil {
			return err
		}
	}

	return a.validateTrades()
}

func validateRFQLiquidationTrade(trade *SyntheticTrade, isTarget bool) error {
	if trade == nil {
		return ErrInvalidTrade
	}
	if trade.Quantity.IsNil() || !trade.Quantity.IsPositive() {
		return ErrInvalidQuantity
	}
	if trade.Quantity.GT(MaxOrderQuantity) {
		return errors.Wrap(ErrInvalidQuantity, trade.Quantity.String())
	}
	if trade.Price.IsNil() || !trade.Price.IsPositive() {
		return ErrInvalidPrice
	}
	if !rfqLiquidationTradeNotionalRepresentable(trade.Quantity, trade.Price) {
		return errors.Wrap(ErrInvalidTrade, "RFQ liquidation trade notional cannot be represented")
	}
	if trade.Margin.IsNil() || trade.Margin.IsNegative() {
		return ErrInvalidMargin
	}
	if isTarget && !trade.Margin.IsZero() {
		return errors.Wrap(ErrInvalidMargin, "RFQ liquidation target trade margin must be zero")
	}
	return nil
}

// rfqLiquidationTradeNotionalRepresentable mirrors LegacyDec.Mul's banker rounding on
// arbitrary-precision integers, avoiding the panic that the multiplication itself emits
// outside LegacyDec's range. RFQ prices are deliberately not capped at MaxOrderPrice:
// oracle and funding moves can put a protocol-required liquidation price above the
// ordinary public-order domain.
func rfqLiquidationTradeNotionalRepresentable(quantity, price math.LegacyDec) bool {
	if quantity.IsNil() || price.IsNil() {
		return false
	}
	precision := math.LegacyOneDec().BigInt()
	product := new(big.Int).Mul(quantity.BigInt(), price.BigInt())
	quotient, remainder := new(big.Int), new(big.Int)
	quotient.QuoRem(product, precision, remainder)
	halfPrecision := new(big.Int).Quo(precision, big.NewInt(2))
	if cmp := remainder.Cmp(halfPrecision); cmp > 0 || cmp == 0 && quotient.Bit(0) == 1 {
		quotient.Add(quotient, big.NewInt(1))
	}
	return quotient.Cmp(legacyDecUpperLimit) < 0
}

type SyntheticTrade struct {
	MarketID     common.Hash    `json:"market_id"`
	SubaccountID common.Hash    `json:"subaccount_id"`
	IsBuy        bool           `json:"is_buy"`
	Quantity     math.LegacyDec `json:"quantity"`
	Price        math.LegacyDec `json:"price"`
	Margin       math.LegacyDec `json:"margin"`
}

// Validate bounds each field the way OrderInfo.ValidateBasic bounds a regular order. The upper
// bounds matter as much as the sign checks: these values come from a contract's reply bytes rather
// than from a signed message, and LegacyDec accepts anything up to 2^256 (~1.16e77). Left
// unbounded, a price/quantity pair reaches DerivativeOrder.CheckNotional and panics with
// "Int overflow" instead of being rejected.
//
// The caps are applied twice, and under ExchangeTypeVersionV1 the two passes see different units.
// ParseRequest runs this on the raw chain-format reply; ConvertSyntheticTradesV1ToV2 then rescales
// price and margin by 1/10^QuoteDecimals, and HandleSyntheticTradeAction runs it again on the
// converted values. The first pass is therefore stricter for v1 by exactly that factor.
//
// That asymmetry is deliberate, and it is the same arrangement v1 orders already use: the raw v1
// message is bounded by OrderInfo.ValidateBasic (msgs.go), whose own comment notes the constants
// are a deliberately loose ceiling on a scaled value because stateless validation cannot know the
// market's scale factor, and DerivativesV1MsgServer then re-applies them to the converted v2 order
// (msg_server_derivative_v1.go).
//
// Neither pass subsumes the other. Dividing cannot push a value back over a cap, so the second
// pass never catches an oversized value the first missed — but it does catch a raw price small
// enough to round to zero on conversion, which the first pass sees as positive.
func (t *SyntheticTrade) Validate() error {
	if t.Quantity.IsNil() || !t.Quantity.IsPositive() {
		return ErrInvalidQuantity
	}

	if t.Quantity.GT(MaxOrderQuantity) {
		return errors.Wrap(ErrInvalidQuantity, t.Quantity.String())
	}

	if t.Price.IsNil() || !t.Price.IsPositive() {
		return ErrInvalidPrice
	}

	if t.Price.GT(MaxOrderPrice) {
		return errors.Wrap(ErrInvalidPrice, t.Price.String())
	}

	if t.Margin.IsNil() || t.Margin.IsNegative() {
		return ErrInvalidMargin
	}

	if t.Margin.GT(MaxOrderMargin) {
		return errors.Wrap(ErrTooMuchOrderMargin, t.Margin.String())
	}

	return nil
}

func (t *SyntheticTrade) IsReduceOnly() bool {
	return t.Margin.IsZero()
}
