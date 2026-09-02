package types

import (
	"math/big"
	"strconv"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	peggytypes "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

const (
	SPOT_MARKET_ID_PREFIX           = "-SPOT-MARKET-"
	PERPETUAL_MARKET_ID_PREFIX      = "-PERPETUAL-MARKET-"
	EXPIRY_FUTURES_MARKET_ID_PREFIX = "-EXPIRY-FUTURES-MARKET-"
	BINARY_OPTIONS_MARKET_ID_PREFIX = "-BINARY-OPTIONS-MARKET-"
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

	return crypto.Keccak256Hash([]byte((SPOT_MARKET_ID_PREFIX + baseDenom + quoteDenom)))
}

func NewPerpetualMarketID(ticker, quoteDenom, oracleBase, oracleQuote string, oracleType oracletypes.OracleType) common.Hash {
	quotePeggyDenom, err := peggytypes.NewPeggyDenomFromString(quoteDenom)
	if err == nil {
		quoteDenom = quotePeggyDenom.String()
	}

	return crypto.Keccak256Hash([]byte((PERPETUAL_MARKET_ID_PREFIX + oracleType.String() + ticker + quoteDenom + oracleBase + oracleQuote)))
}

func NewBinaryOptionsMarketID(ticker, quoteDenom, oracleSymbol, oracleProvider string, oracleType oracletypes.OracleType) common.Hash {
	quotePeggyDenom, err := peggytypes.NewPeggyDenomFromString(quoteDenom)
	if err == nil {
		quoteDenom = quotePeggyDenom.String()
	}

	return crypto.Keccak256Hash([]byte((BINARY_OPTIONS_MARKET_ID_PREFIX +
		oracleType.String() +
		ticker +
		quoteDenom +
		oracleSymbol +
		oracleProvider)))
}

func NewExpiryFuturesMarketID(ticker, quoteDenom, oracleBase, oracleQuote string, oracleType oracletypes.OracleType, expiry int64) common.Hash {
	quotePeggyDenom, err := peggytypes.NewPeggyDenomFromString(quoteDenom)
	if err == nil {
		quoteDenom = quotePeggyDenom.String()
	}
	return crypto.Keccak256Hash([]byte((EXPIRY_FUTURES_MARKET_ID_PREFIX +
		oracleType.String() +
		ticker +
		quoteDenom +
		oracleBase +
		oracleQuote +
		strconv.Itoa(int(expiry)))))
}

func NewDerivativesMarketID(ticker, quoteDenom, oracleBase, oracleQuote string, oracleType oracletypes.OracleType, expiry int64) common.Hash {
	if expiry == -1 {
		return NewPerpetualMarketID(ticker, quoteDenom, oracleBase, oracleQuote, oracleType)
	} else {
		return NewExpiryFuturesMarketID(ticker, quoteDenom, oracleBase, oracleQuote, oracleType, expiry)
	}
}

func PriceFromChainFormat(price math.LegacyDec, baseDecimals, quoteDecimals uint32) math.LegacyDec {
	if price.IsNil() {
		return price
	}
	baseMultiplier := math.LegacyNewDec(10).Power(uint64(baseDecimals))
	quoteMultiplier := math.LegacyNewDec(10).Power(uint64(quoteDecimals))
	return price.Mul(baseMultiplier).Quo(quoteMultiplier)
}

func QuantityFromChainFormat(quantity math.LegacyDec, decimals uint32) math.LegacyDec {
	if quantity.IsNil() {
		return quantity
	}
	multiplier := math.LegacyNewDec(10).Power(uint64(decimals))
	return quantity.Quo(multiplier)
}

func NotionalFromChainFormat(notional math.LegacyDec, decimals uint32) math.LegacyDec {
	if notional.IsNil() {
		return notional
	}
	multiplier := math.LegacyNewDec(10).Power(uint64(decimals))
	return notional.Quo(multiplier)
}

func PriceToChainFormat(humanReadableValue math.LegacyDec, baseDecimals, quoteDecimals uint32) math.LegacyDec {
	if humanReadableValue.IsNil() {
		return humanReadableValue
	}
	baseMultiplier := math.LegacyNewDec(10).Power(uint64(baseDecimals))
	quoteMultiplier := math.LegacyNewDec(10).Power(uint64(quoteDecimals))
	return humanReadableValue.Mul(quoteMultiplier).Quo(baseMultiplier)
}

func QuantityToChainFormat(humanReadableValue math.LegacyDec, decimals uint32) math.LegacyDec {
	if humanReadableValue.IsNil() {
		return humanReadableValue
	}
	multiplier := math.LegacyNewDec(10).Power(uint64(decimals))
	return humanReadableValue.Mul(multiplier)
}

func NotionalToChainFormat(humanReadableValue math.LegacyDec, decimals uint32) math.LegacyDec {
	if humanReadableValue.IsNil() {
		return humanReadableValue
	}
	multiplier := math.LegacyNewDec(10).Power(uint64(decimals))
	return humanReadableValue.Mul(multiplier)
}

// legacyDecUpperLimit is the largest value a LegacyDec may hold, in the same internal
// representation LegacyDec.BigInt() returns: 2^256 scaled by 10^LegacyPrecision. cosmossdk.io/math
// keeps its own copy of this unexported, so it is recomputed here.
var legacyDecUpperLimit = new(big.Int).Mul(
	new(big.Int).Lsh(big.NewInt(1), 256),
	new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(math.LegacyPrecision)), nil),
)

// CanRepresentNotionalInChainFormat reports whether NotionalToChainFormat can scale this value by
// 10^decimals without leaving LegacyDec's valid range.
//
// It exists because NotionalToChainFormat panics rather than returning an error when the result is
// out of range, so any caller that can be handed an attacker-influenced magnitude has to ask first.
// The scaling is done on big.Int, which grows instead of panicking, so the check itself is safe.
func CanRepresentNotionalInChainFormat(humanReadableValue math.LegacyDec, decimals uint32) bool {
	if humanReadableValue.IsNil() {
		return false
	}

	scaled := new(big.Int).Mul(
		new(big.Int).Abs(humanReadableValue.BigInt()),
		new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil),
	)

	return scaled.Cmp(legacyDecUpperLimit) < 0
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
