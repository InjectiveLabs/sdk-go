package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	oracletypes "github.com/InjectiveLabs/injective-core/injective-chain/modules/oracle/types"
	peggytypes "github.com/InjectiveLabs/injective-core/injective-chain/modules/peggy/types"
)

var BinaryOptionsMarketRefundFlagPrice = sdk.NewDec(-1)

type DerivativeMarketInfo struct {
	Market    *DerivativeMarket
	MarkPrice sdk.Dec
	Funding   *PerpetualMarketFunding
}

func NewSpotMarketID(baseDenom, quoteDenom string) common.Hash {
	basePeggyDenom, err := peggytypes.NewPeggyDenomFromString(baseDenom)
	if err == nil {
		baseDenom = basePeggyDenom.String()
	}

	quotePeggyDenom, err := peggytypes.NewPeggyDenomFromString(quoteDenom)
	if err == nil {
		quoteDenom = quotePeggyDenom.String()
	}

	return crypto.Keccak256Hash([]byte((baseDenom + quoteDenom)))
}

func NewPerpetualMarketID(ticker, quoteDenom, oracleBase, oracleQuote string, oracleType oracletypes.OracleType) common.Hash {
	quotePeggyDenom, err := peggytypes.NewPeggyDenomFromString(quoteDenom)
	if err == nil {
		quoteDenom = quotePeggyDenom.String()
	}

	return crypto.Keccak256Hash([]byte((oracleType.String() + ticker + quoteDenom + oracleBase + oracleQuote)))
}

func NewBinaryOptionsMarketID(ticker, quoteDenom, oracleSymbol, oracleProvider string, oracleType oracletypes.OracleType) common.Hash {
	quotePeggyDenom, err := peggytypes.NewPeggyDenomFromString(quoteDenom)
	if err == nil {
		quoteDenom = quotePeggyDenom.String()
	}

	return crypto.Keccak256Hash([]byte((oracleType.String() + ticker + quoteDenom + oracleSymbol + oracleProvider)))
}

func NewExpiryFuturesMarketID(ticker, quoteDenom, oracleBase, oracleQuote string, oracleType oracletypes.OracleType, expiry int64) common.Hash {
	quotePeggyDenom, err := peggytypes.NewPeggyDenomFromString(quoteDenom)
	if err == nil {
		quoteDenom = quotePeggyDenom.String()
	}
	return crypto.Keccak256Hash([]byte((oracleType.String() + ticker + quoteDenom + oracleBase + oracleQuote + strconv.Itoa(int(expiry)))))
}

func NewDerivativesMarketID(ticker, quoteDenom, oracleBase, oracleQuote string, oracleType oracletypes.OracleType, expiry int64) common.Hash {
	if expiry == -1 {
		return NewPerpetualMarketID(ticker, quoteDenom, oracleBase, oracleQuote, oracleType)
	} else {
		return NewExpiryFuturesMarketID(ticker, quoteDenom, oracleBase, oracleQuote, oracleType, expiry)
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

func (m *DerivativeMarket) GetMinQuantityTickSize() sdk.Dec {
	return m.MinQuantityTickSize
}

type MarketType byte

//nolint:all
const (
	MarketType_Spot MarketType = iota
	MarketType_Perpetual
	MarketType_Expiry
	MarketType_BinaryOption
)

func (m MarketType) IsSpot() bool {
	return m == MarketType_Spot
}

func (m MarketType) IsPerpetual() bool {
	return m == MarketType_Perpetual
}

func (m MarketType) IsExpiry() bool {
	return m == MarketType_Expiry
}

func (m MarketType) IsBinaryOptions() bool {
	return m == MarketType_BinaryOption
}

func (m *DerivativeMarket) GetMarketType() MarketType {
	if m.IsPerpetual {
		return MarketType_Perpetual
	} else {
		return MarketType_Expiry
	}
}

func (m *DerivativeMarket) GetMakerFeeRate() sdk.Dec {
	return m.MakerFeeRate
}

func (m *DerivativeMarket) GetTakerFeeRate() sdk.Dec {
	return m.TakerFeeRate
}

func (m *DerivativeMarket) GetRelayerFeeShareRate() sdk.Dec {
	return m.RelayerFeeShareRate
}

func (m *DerivativeMarket) GetInitialMarginRatio() sdk.Dec {
	return m.InitialMarginRatio
}

func (m *DerivativeMarket) GetMinPriceTickSize() sdk.Dec {
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

/// Binary Options Markets
//

func (m *BinaryOptionsMarket) GetMarketType() MarketType {
	return MarketType_BinaryOption
}

func (m *BinaryOptionsMarket) GetInitialMarginRatio() sdk.Dec {
	return sdk.OneDec()
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

func (m *BinaryOptionsMarket) GetMinPriceTickSize() sdk.Dec {
	return m.MinPriceTickSize
}

func (m *BinaryOptionsMarket) GetMinQuantityTickSize() sdk.Dec {
	return m.MinQuantityTickSize
}

func (m *BinaryOptionsMarket) GetTicker() string {
	return m.Ticker
}

func (m *BinaryOptionsMarket) GetQuoteDenom() string {
	return m.QuoteDenom
}

func (m *BinaryOptionsMarket) GetMakerFeeRate() sdk.Dec {
	return m.MakerFeeRate
}

func (m *BinaryOptionsMarket) GetTakerFeeRate() sdk.Dec {
	return m.TakerFeeRate
}

func (m *BinaryOptionsMarket) GetRelayerFeeShareRate() sdk.Dec {
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
