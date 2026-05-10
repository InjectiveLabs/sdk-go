package v2

import (
	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

var BinaryOptionsMarketRefundFlagPrice = math.LegacyNewDec(-1)

type MarketIDQuoteDenomMakerFee struct {
	MarketID   common.Hash
	QuoteDenom string
	MakerFee   math.LegacyDec
}

type DerivativeMarketInfo struct {
	Market    *DerivativeMarket
	MarkPrice math.LegacyDec
	Funding   *PerpetualMarketFunding
}

func (m *SpotMarket) IsActive() bool {
	return m.Status == MarketStatus_Active
}

func (m *SpotMarket) IsInactive() bool {
	return !m.IsActive()
}

func (m *SpotMarket) MarketID() common.Hash {
	return common.HexToHash(m.MarketId)
}

func (m *SpotMarket) StatusSupportsOrderCancellations() bool {
	if m == nil {
		return false
	}
	return m.Status.SupportsOrderCancellations()
}

func (*SpotMarket) GetMarketType() types.MarketType {
	return types.MarketType_Spot
}

func (m *SpotMarket) GetMakerFeeRate() math.LegacyDec {
	return m.MakerFeeRate
}

func (m *SpotMarket) GetTakerFeeRate() math.LegacyDec {
	return m.TakerFeeRate
}

func (m *SpotMarket) GetRelayerFeeShareRate() math.LegacyDec {
	return m.RelayerFeeShareRate
}

func (m *SpotMarket) GetMinPriceTickSize() math.LegacyDec {
	return m.MinPriceTickSize
}

func (m *SpotMarket) GetMinQuantityTickSize() math.LegacyDec {
	return m.MinQuantityTickSize
}

func (m *SpotMarket) GetMinNotional() math.LegacyDec {
	return m.MinNotional
}

func (m *SpotMarket) GetMarketStatus() MarketStatus {
	return m.Status
}

func (m *SpotMarket) GetDisabledMinimalProtocolFee() bool {
	return m.HasDisabledMinimalProtocolFee
}

func (m *SpotMarket) PriceFromChainFormat(price math.LegacyDec) math.LegacyDec {
	return types.PriceFromChainFormat(price, m.BaseDecimals, m.QuoteDecimals)
}

func (m *SpotMarket) QuantityFromChainFormat(quantity math.LegacyDec) math.LegacyDec {
	return types.QuantityFromChainFormat(quantity, m.BaseDecimals)
}

func (m *SpotMarket) NotionalFromChainFormat(notional math.LegacyDec) math.LegacyDec {
	return types.NotionalFromChainFormat(notional, m.QuoteDecimals)
}

func (m *SpotMarket) PriceToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.PriceToChainFormat(humanReadableValue, m.BaseDecimals, m.QuoteDecimals)
}

func (m *SpotMarket) QuantityToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.QuantityToChainFormat(humanReadableValue, m.BaseDecimals)
}

func (m *SpotMarket) NotionalToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.NotionalToChainFormat(humanReadableValue, m.QuoteDecimals)
}

func (m *ExpiryFuturesMarketInfo) IsPremature(currBlockTime int64) bool {
	return currBlockTime < m.TwapStartTimestamp
}

func (m *ExpiryFuturesMarketInfo) IsStartingMaturation(currBlockTime int64) bool {
	return currBlockTime >= m.TwapStartTimestamp &&
		(m.ExpirationTwapStartBaseCumulativePrice.IsNil() || m.ExpirationTwapStartBaseCumulativePrice.IsZero())
}

func (m *ExpiryFuturesMarketInfo) IsMatured(currBlockTime int64) bool {
	return currBlockTime >= m.ExpirationTimestamp
}

func (m *DerivativeMarket) MarketID() common.Hash {
	return common.HexToHash(m.MarketId)
}

func (m *DerivativeMarket) StatusSupportsOrderCancellations() bool {
	if m == nil {
		return false
	}
	return m.Status.SupportsOrderCancellations()
}

func (m *DerivativeMarket) IsTimeExpiry() bool {
	return !m.IsPerpetual
}

func (m *DerivativeMarket) IsActive() bool {
	return m.Status == MarketStatus_Active
}

func (m *DerivativeMarket) IsInactive() bool {
	return !m.IsActive()
}

func (m *DerivativeMarket) GetMinQuantityTickSize() math.LegacyDec {
	return m.MinQuantityTickSize
}

func (m *DerivativeMarket) GetMinNotional() math.LegacyDec {
	return m.MinNotional
}

func (m *DerivativeMarket) GetMarketType() types.MarketType {
	if m.IsPerpetual {
		return types.MarketType_Perpetual
	} else {
		return types.MarketType_Expiry
	}
}

func (m *DerivativeMarket) GetMakerFeeRate() math.LegacyDec {
	return m.MakerFeeRate
}

func (m *DerivativeMarket) GetTakerFeeRate() math.LegacyDec {
	return m.TakerFeeRate
}

func (m *DerivativeMarket) GetRelayerFeeShareRate() math.LegacyDec {
	return m.RelayerFeeShareRate
}

func (m *DerivativeMarket) GetInitialMarginRatio() math.LegacyDec {
	return m.InitialMarginRatio
}

func (m *DerivativeMarket) GetMinPriceTickSize() math.LegacyDec {
	return m.MinPriceTickSize
}

func (m *DerivativeMarket) GetQuoteDenom() string {
	return m.QuoteDenom
}

func (m *DerivativeMarket) GetTicker() string {
	return m.Ticker
}

func (m *DerivativeMarket) GetIsPerpetual() bool {
	return m.IsPerpetual
}

func (m *DerivativeMarket) GetOracleScaleFactor() uint32 {
	return m.OracleScaleFactor
}

func (m *DerivativeMarket) GetMarketStatus() MarketStatus {
	return m.Status
}

func (m *DerivativeMarket) GetDisabledMinimalProtocolFee() bool {
	return m.HasDisabledMinimalProtocolFee
}

func (m *DerivativeMarket) GetQuoteDecimals() uint32 {
	return m.QuoteDecimals
}

func (m *DerivativeMarket) GetOpenNotionalCap() OpenNotionalCap {
	return m.OpenNotionalCap
}

func (m *DerivativeMarket) IsCrossMarginEligible() bool {
	return m.CrossMarginEligible
}

func (m *DerivativeMarket) PriceFromChainFormat(price math.LegacyDec) math.LegacyDec {
	return types.PriceFromChainFormat(price, 0, m.QuoteDecimals)
}

func (*DerivativeMarket) QuantityFromChainFormat(quantity math.LegacyDec) math.LegacyDec {
	return types.QuantityFromChainFormat(quantity, 0)
}

func (m *DerivativeMarket) NotionalFromChainFormat(notional math.LegacyDec) math.LegacyDec {
	return types.NotionalFromChainFormat(notional, m.QuoteDecimals)
}

func (m *DerivativeMarket) PriceToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.PriceToChainFormat(humanReadableValue, 0, m.QuoteDecimals)
}

func (*DerivativeMarket) QuantityToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.QuantityToChainFormat(humanReadableValue, 0)
}

func (m *DerivativeMarket) NotionalToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.NotionalToChainFormat(humanReadableValue, m.QuoteDecimals)
}

/// Binary Options Markets
//

func (m *BinaryOptionsMarket) GetMarketType() types.MarketType {
	return types.MarketType_BinaryOption
}

func (m *BinaryOptionsMarket) GetInitialMarginRatio() math.LegacyDec {
	return math.LegacyOneDec()
}
func (m *BinaryOptionsMarket) IsInactive() bool {
	return !m.IsActive()
}

func (m *BinaryOptionsMarket) IsActive() bool {
	return m.Status == MarketStatus_Active
}

func (m *BinaryOptionsMarket) MarketID() common.Hash {
	return common.HexToHash(m.MarketId)
}

func (m *BinaryOptionsMarket) GetMinPriceTickSize() math.LegacyDec {
	return m.MinPriceTickSize
}

func (m *BinaryOptionsMarket) GetMinQuantityTickSize() math.LegacyDec {
	return m.MinQuantityTickSize
}

func (m *BinaryOptionsMarket) GetMinNotional() math.LegacyDec {
	return m.MinNotional
}

func (m *BinaryOptionsMarket) GetTicker() string {
	return m.Ticker
}

func (m *BinaryOptionsMarket) GetQuoteDenom() string {
	return m.QuoteDenom
}

func (m *BinaryOptionsMarket) GetMakerFeeRate() math.LegacyDec {
	return m.MakerFeeRate
}

func (m *BinaryOptionsMarket) GetTakerFeeRate() math.LegacyDec {
	return m.TakerFeeRate
}

func (m *BinaryOptionsMarket) GetRelayerFeeShareRate() math.LegacyDec {
	return m.RelayerFeeShareRate
}

func (m *BinaryOptionsMarket) GetIsPerpetual() bool {
	return false
}

func (m *BinaryOptionsMarket) StatusSupportsOrderCancellations() bool {
	if m == nil {
		return false
	}
	return m.Status.SupportsOrderCancellations()
}

func (m *BinaryOptionsMarket) GetOracleScaleFactor() uint32 {
	return m.OracleScaleFactor
}

func (m *BinaryOptionsMarket) GetMarketStatus() MarketStatus {
	return m.Status
}

func (m *BinaryOptionsMarket) GetDisabledMinimalProtocolFee() bool {
	return m.HasDisabledMinimalProtocolFee
}

func (m *BinaryOptionsMarket) GetQuoteDecimals() uint32 {
	return m.QuoteDecimals
}

func (m *BinaryOptionsMarket) GetOpenNotionalCap() OpenNotionalCap {
	return m.OpenNotionalCap
}

func (*BinaryOptionsMarket) IsCrossMarginEligible() bool {
	return false
}

func (m *BinaryOptionsMarket) PriceFromChainFormat(price math.LegacyDec) math.LegacyDec {
	return types.PriceFromChainFormat(price, 0, m.QuoteDecimals)
}

func (m *BinaryOptionsMarket) QuantityFromChainFormat(quantity math.LegacyDec) math.LegacyDec {
	return types.QuantityFromChainFormat(quantity, 0)
}

func (m *BinaryOptionsMarket) NotionalFromChainFormat(notional math.LegacyDec) math.LegacyDec {
	return types.NotionalFromChainFormat(notional, m.QuoteDecimals)
}

func (m *BinaryOptionsMarket) PriceToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.PriceToChainFormat(humanReadableValue, 0, m.QuoteDecimals)
}

func (m *BinaryOptionsMarket) QuantityToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.QuantityToChainFormat(humanReadableValue, 0)
}

func (m *BinaryOptionsMarket) NotionalToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.NotionalToChainFormat(humanReadableValue, m.QuoteDecimals)
}

type MarketI interface {
	MarketID() common.Hash
	GetMarketType() types.MarketType
	GetMinPriceTickSize() math.LegacyDec
	GetMinQuantityTickSize() math.LegacyDec
	GetMinNotional() math.LegacyDec
	GetTicker() string
	GetMakerFeeRate() math.LegacyDec
	GetTakerFeeRate() math.LegacyDec
	GetRelayerFeeShareRate() math.LegacyDec
	GetQuoteDenom() string
	StatusSupportsOrderCancellations() bool
	GetDisabledMinimalProtocolFee() bool
	GetMarketStatus() MarketStatus
	PriceFromChainFormat(price math.LegacyDec) math.LegacyDec
	QuantityFromChainFormat(quantity math.LegacyDec) math.LegacyDec
	NotionalFromChainFormat(notional math.LegacyDec) math.LegacyDec
	PriceToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec
	QuantityToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec
	NotionalToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec
}

type DerivativeMarketI interface {
	MarketI
	GetIsPerpetual() bool
	GetInitialMarginRatio() math.LegacyDec
	GetOracleScaleFactor() uint32
	GetQuoteDecimals() uint32
	GetOpenNotionalCap() OpenNotionalCap
	IsCrossMarginEligible() bool
}

func IsMarketSolvent(availableMarketFunds, marketBalanceDelta math.LegacyDec) bool {
	return availableMarketFunds.Add(marketBalanceDelta).GTE(math.LegacyZeroDec())
}

// nolint // ok
func GetMarketBalanceDelta(
	payout,
	collateralizationMargin,
	tradeFee math.LegacyDec,
	isReduceOnly bool,
) math.LegacyDec {
	if payout.IsNegative() {
		// if payout is negative, don't just add these to the market balance,
		// instead try to adjust market balance later when insurance fund is tapped
		payout = math.LegacyZeroDec()
	}

	if isReduceOnly {
		// trade fee is removed from payout for RO, but still should be removed from market balance
		payout = payout.Add(tradeFee)
	}

	return collateralizationMargin.Sub(payout)
}
