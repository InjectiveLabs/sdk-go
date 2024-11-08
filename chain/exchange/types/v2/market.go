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

type MarketInterface interface {
	PriceFromChainFormat(price math.LegacyDec) math.LegacyDec
	QuantityFromChainFormat(quantity math.LegacyDec) math.LegacyDec
	NotionalFromChainFormat(notional math.LegacyDec) math.LegacyDec
	PriceToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec
	QuantityToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec
	NotionalToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec
}

func NewV1MarketVolumeFromV2(market MarketInterface, v2MarketVolume MarketVolume) types.MarketVolume {
	return types.MarketVolume{
		MarketId: v2MarketVolume.MarketId,
		Volume:   NewV1VolumeRecordFromV2(market, v2MarketVolume.Volume),
	}
}

func NewV1VolumeRecordFromV2(market MarketInterface, v2VolumeRecord VolumeRecord) types.VolumeRecord {
	chainFormatMakerVolume := market.NotionalToChainFormat(v2VolumeRecord.MakerVolume)
	chainFormatTakerVolume := market.NotionalToChainFormat(v2VolumeRecord.TakerVolume)
	return types.VolumeRecord{
		MakerVolume: chainFormatMakerVolume,
		TakerVolume: chainFormatTakerVolume,
	}
}

func NewV1SpotMarketFromV2(spotMarket SpotMarket) types.SpotMarket {
	chainFormattedMinPriceTickSize := spotMarket.PriceToChainFormat(spotMarket.MinPriceTickSize)
	chainFormattedMinQuantityTickSize := spotMarket.QuantityToChainFormat(spotMarket.MinQuantityTickSize)
	chainFormattedMinNotional := spotMarket.NotionalToChainFormat(spotMarket.MinNotional)
	return types.SpotMarket{
		Ticker:              spotMarket.Ticker,
		BaseDenom:           spotMarket.BaseDenom,
		QuoteDenom:          spotMarket.QuoteDenom,
		MakerFeeRate:        spotMarket.MakerFeeRate,
		TakerFeeRate:        spotMarket.TakerFeeRate,
		RelayerFeeShareRate: spotMarket.RelayerFeeShareRate,
		MarketId:            spotMarket.MarketId,
		Status:              types.MarketStatus(spotMarket.Status),
		MinPriceTickSize:    chainFormattedMinPriceTickSize,
		MinQuantityTickSize: chainFormattedMinQuantityTickSize,
		MinNotional:         chainFormattedMinNotional,
		Admin:               spotMarket.Admin,
		AdminPermissions:    spotMarket.AdminPermissions,
		BaseDecimals:        spotMarket.BaseDecimals,
		QuoteDecimals:       spotMarket.QuoteDecimals,
	}
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

func (m *SpotMarket) GetMarketType() types.MarketType {
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

func NewV1ExpiryFuturesMarketInfoStateFromV2(market DerivativeMarket, marketInfoState ExpiryFuturesMarketInfoState) types.ExpiryFuturesMarketInfoState {
	v1MarketInfo := NewV1ExpiryFuturesMarketInfoFromV2(market, *marketInfoState.MarketInfo)
	return types.ExpiryFuturesMarketInfoState{
		MarketId:   marketInfoState.MarketId,
		MarketInfo: &v1MarketInfo,
	}
}

func NewV1ExpiryFuturesMarketInfoFromV2(market DerivativeMarket, marketInfo ExpiryFuturesMarketInfo) types.ExpiryFuturesMarketInfo {
	return types.ExpiryFuturesMarketInfo{
		MarketId:                           marketInfo.MarketId,
		ExpirationTimestamp:                marketInfo.ExpirationTimestamp,
		TwapStartTimestamp:                 marketInfo.TwapStartTimestamp,
		ExpirationTwapStartPriceCumulative: market.PriceToChainFormat(marketInfo.ExpirationTwapStartPriceCumulative),
		SettlementPrice:                    market.PriceToChainFormat(marketInfo.SettlementPrice),
	}
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

func NewV1DerivativeMarketFromV2(derivativeMarket DerivativeMarket) types.DerivativeMarket {
	return types.DerivativeMarket{
		Ticker:                 derivativeMarket.Ticker,
		OracleBase:             derivativeMarket.OracleBase,
		OracleQuote:            derivativeMarket.OracleQuote,
		OracleType:             derivativeMarket.OracleType,
		OracleScaleFactor:      derivativeMarket.OracleScaleFactor + derivativeMarket.QuoteDecimals,
		QuoteDenom:             derivativeMarket.QuoteDenom,
		MarketId:               derivativeMarket.MarketId,
		InitialMarginRatio:     derivativeMarket.InitialMarginRatio,
		MaintenanceMarginRatio: derivativeMarket.MaintenanceMarginRatio,
		MakerFeeRate:           derivativeMarket.MakerFeeRate,
		TakerFeeRate:           derivativeMarket.TakerFeeRate,
		RelayerFeeShareRate:    derivativeMarket.RelayerFeeShareRate,
		IsPerpetual:            derivativeMarket.IsPerpetual,
		Status:                 types.MarketStatus(derivativeMarket.Status),
		MinPriceTickSize:       derivativeMarket.PriceToChainFormat(derivativeMarket.MinPriceTickSize),
		MinQuantityTickSize:    derivativeMarket.QuantityToChainFormat(derivativeMarket.MinQuantityTickSize),
		MinNotional:            derivativeMarket.NotionalToChainFormat(derivativeMarket.MinNotional),
		Admin:                  derivativeMarket.Admin,
		AdminPermissions:       derivativeMarket.AdminPermissions,
		QuoteDecimals:          derivativeMarket.QuoteDecimals,
	}
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

func (m *DerivativeMarket) QuantityFromChainFormat(quantity math.LegacyDec) math.LegacyDec {
	return types.QuantityFromChainFormat(quantity, 0)
}

func (m *DerivativeMarket) NotionalFromChainFormat(notional math.LegacyDec) math.LegacyDec {
	return types.NotionalFromChainFormat(notional, m.QuoteDecimals)
}

func (m *DerivativeMarket) PriceToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.PriceToChainFormat(humanReadableValue, 0, m.QuoteDecimals)
}

func (m *DerivativeMarket) QuantityToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.QuantityToChainFormat(humanReadableValue, 0)
}

func (m *DerivativeMarket) NotionalToChainFormat(humanReadableValue math.LegacyDec) math.LegacyDec {
	return types.NotionalToChainFormat(humanReadableValue, m.QuoteDecimals)
}

/// Binary Options Markets
//

func NewV1BinaryOptionsMarketFromV2(market BinaryOptionsMarket) types.BinaryOptionsMarket {
	v1Market := types.BinaryOptionsMarket{
		Ticker:              market.Ticker,
		OracleSymbol:        market.OracleSymbol,
		OracleProvider:      market.OracleProvider,
		OracleType:          market.OracleType,
		OracleScaleFactor:   market.OracleScaleFactor + market.QuoteDecimals,
		ExpirationTimestamp: market.ExpirationTimestamp,
		SettlementTimestamp: market.SettlementTimestamp,
		Admin:               market.Admin,
		QuoteDenom:          market.QuoteDenom,
		MarketId:            market.MarketId,
		MakerFeeRate:        market.MakerFeeRate,
		TakerFeeRate:        market.TakerFeeRate,
		RelayerFeeShareRate: market.RelayerFeeShareRate,
		Status:              types.MarketStatus(market.Status),
		MinPriceTickSize:    market.PriceToChainFormat(market.MinPriceTickSize),
		MinQuantityTickSize: market.QuantityToChainFormat(market.MinQuantityTickSize),
		MinNotional:         market.NotionalToChainFormat(market.MinNotional),
		AdminPermissions:    market.AdminPermissions,
		QuoteDecimals:       market.QuoteDecimals,
	}

	if market.SettlementPrice != nil {
		chainFormatSettlementPrice := market.PriceToChainFormat(*market.SettlementPrice)
		v1Market.SettlementPrice = &chainFormatSettlementPrice
	}

	return v1Market
}

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

func NewV1PerpetualMarketFundingStateFromV2(market DerivativeMarket, fundingState PerpetualMarketFundingState) types.PerpetualMarketFundingState {
	v1Funding := NewV1PerpetualMarketFundingFromV2(market, *fundingState.Funding)
	return types.PerpetualMarketFundingState{
		MarketId: fundingState.MarketId,
		Funding:  &v1Funding,
	}
}

func NewV1PerpetualMarketFundingFromV2(market DerivativeMarket, funding PerpetualMarketFunding) types.PerpetualMarketFunding {
	return types.PerpetualMarketFunding{
		CumulativeFunding: market.NotionalToChainFormat(funding.CumulativeFunding),
		CumulativePrice:   market.PriceToChainFormat(funding.CumulativePrice),
		LastTimestamp:     funding.LastTimestamp,
	}
}

func NewV1DerivativeMarketSettlementInfoFromV2(market DerivativeMarket, settlementInfo DerivativeMarketSettlementInfo) types.DerivativeMarketSettlementInfo {
	return types.DerivativeMarketSettlementInfo{
		MarketId:        settlementInfo.MarketId,
		SettlementPrice: market.PriceToChainFormat(settlementInfo.SettlementPrice),
	}
}
