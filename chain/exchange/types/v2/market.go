package v2

import (
	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

var BinaryOptionsMarketRefundFlagPrice = math.LegacyNewDec(-1)

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
	return currBlockTime >= m.TwapStartTimestamp && (m.ExpirationTwapStartPriceCumulative.IsNil() || m.ExpirationTwapStartPriceCumulative.IsZero())
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

func (m *DerivativeMarket) GetQuoteDecimals() uint32 {
	return m.QuoteDecimals
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

func (m *BinaryOptionsMarket) GetQuoteDecimals() uint32 {
	return m.QuoteDecimals
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
