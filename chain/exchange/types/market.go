package types

import (
	"strconv"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	peggytypes "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

var BinaryOptionsMarketRefundFlagPrice = math.LegacyNewDec(-1)

type DerivativeMarketInfo struct {
	Market    *DerivativeMarket
	MarkPrice math.LegacyDec
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

func PriceFromChainFormat(price math.LegacyDec, baseDecimals, quoteDecimals uint32) math.LegacyDec {
	baseMultiplier := math.LegacyNewDec(10).Power(uint64(baseDecimals))
	quoteMultiplier := math.LegacyNewDec(10).Power(uint64(quoteDecimals))
	return price.Mul(baseMultiplier).Quo(quoteMultiplier)
}

func QuantityFromChainFormat(quantity math.LegacyDec, decimals uint32) math.LegacyDec {
	multiplier := math.LegacyNewDec(10).Power(uint64(decimals))
	return quantity.Quo(multiplier)
}

func NotionalFromChainFormat(notional math.LegacyDec, decimals uint32) math.LegacyDec {
	multiplier := math.LegacyNewDec(10).Power(uint64(decimals))
	return notional.Quo(multiplier)
}

func PriceToChainFormat(humanReadableValue math.LegacyDec, baseDecimals, quoteDecimals uint32) math.LegacyDec {
	baseMultiplier := math.LegacyNewDec(10).Power(uint64(baseDecimals))
	quoteMultiplier := math.LegacyNewDec(10).Power(uint64(quoteDecimals))
	return humanReadableValue.Mul(quoteMultiplier).Quo(baseMultiplier)
}

func QuantityToChainFormat(humanReadableValue math.LegacyDec, decimals uint32) math.LegacyDec {
	multiplier := math.LegacyNewDec(10).Power(uint64(decimals))
	return humanReadableValue.Mul(multiplier)
}

func NotionalToChainFormat(humanReadableValue math.LegacyDec, decimals uint32) math.LegacyDec {
	multiplier := math.LegacyNewDec(10).Power(uint64(decimals))
	return humanReadableValue.Mul(multiplier)
}

type MarketType byte

// nolint:all
const (
	MarketType_Spot MarketType = iota
	MarketType_Perpetual
	MarketType_Expiry
	MarketType_BinaryOption
)

func (m MarketType) IsPerpetual() bool {
	return m == MarketType_Perpetual
}

func (m MarketType) IsBinaryOptions() bool {
	return m == MarketType_BinaryOption
}

func (m *BinaryOptionsMarket) MarketID() common.Hash {
	return common.HexToHash(m.MarketId)
}

func (m *BinaryOptionsMarket) GetOracleScaleFactor() uint32 {
	return m.OracleScaleFactor
}
