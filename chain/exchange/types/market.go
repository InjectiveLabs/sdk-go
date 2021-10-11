package types

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	peggytypes "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

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

func (m *ExpiryFuturesMarketInfo) IsPremature(currBlockTime int64) bool {
	return currBlockTime < m.TwapStartTimestamp
}

func (m *ExpiryFuturesMarketInfo) IsStartingMaturation(currBlockTime int64) bool {
	return currBlockTime >= m.TwapStartTimestamp && (m.ExpirationTwapStartPriceCumulative.IsNil() || m.ExpirationTwapStartPriceCumulative.IsZero())
}

func (m *ExpiryFuturesMarketInfo) IsMatured(currBlockTime int64) bool {
	return currBlockTime >= m.ExpirationTimestamp
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
