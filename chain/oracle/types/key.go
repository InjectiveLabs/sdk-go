package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	ModuleName = "oracle"
	StoreKey   = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	//// Version defines the current version the IBC module supports
	// Version = "bandchain-1"
	//
	//// PortID is the default port id that module binds to
	// PortID = "oracle"
)

var (
	// Keys for band store prefixes
	BandPriceKey   = []byte{0x01}
	BandRelayerKey = []byte{0x02}

	// Keys for pricefeeder store prefixes
	PricefeedInfoKey    = []byte{0x11}
	PricefeedPriceKey   = []byte{0x12}
	PricefeedRelayerKey = []byte{0x13}

	CoinbasePriceKey = []byte{0x21}

	// Band IBC
	BandIBCPriceKey           = []byte{0x31}
	LatestClientIDKey         = []byte{0x32}
	BandIBCCallDataRecordKey  = []byte{0x33}
	BandIBCOracleRequestIDKey = []byte{0x34}
	BandIBCParamsKey          = []byte{0x35}
	LatestRequestIDKey        = []byte{0x36}

	// Prefixes for chainlink keys
	ChainlinkPriceKey = []byte{0x41}

	SymbolHistoricalPriceRecordsPrefix = []byte{0x51} // prefix for each key to a symbols's historical price records
	SymbolsMapLastPriceTimestampsKey   = []byte{0x52} // key for symbols map with latest price update timestamps

	// ProviderInfoPrefix is the prefix for the Provider => ProviderInfo store.
	ProviderInfoPrefix = []byte{0x61}
	// ProviderIndexPrefix is the prefix for the ProviderAddress => Provider index store.
	ProviderIndexPrefix = []byte{0x62}
	// ProviderPricePrefix is the prefix for the Provider + symbol => PriceState store.
	ProviderPricePrefix = []byte{0x63}
)

func GetBandPriceStoreKey(symbol string) []byte {
	return append(BandPriceKey, []byte(symbol)...)
}

func GetBandRelayerStoreKey(relayer sdk.AccAddress) []byte {
	return append(BandRelayerKey, relayer.Bytes()...)
}

func GetBandIBCOracleRequestIDKey(requestID uint64) []byte {
	return append(BandIBCOracleRequestIDKey, sdk.Uint64ToBigEndian(requestID)...)
}

func GetBandIBCPriceStoreKey(symbol string) []byte {
	return append(BandIBCPriceKey, []byte(symbol)...)
}

func GetBandIBCCallDataRecordKey(clientID uint64) []byte {
	return append(BandIBCCallDataRecordKey, sdk.Uint64ToBigEndian(clientID)...)
}

func GetBaseQuoteHash(oracleBase, oracleQuote string) common.Hash {
	return crypto.Keccak256Hash([]byte(oracleBase + oracleQuote))
}

func GetPriceFeedInfoKey(priceFeedInfo *PriceFeedInfo) []byte {
	return append(PricefeedInfoKey, GetBaseQuoteHash(priceFeedInfo.Base, priceFeedInfo.Quote).Bytes()...)
}

func GetPriceFeedPriceStoreKey(baseQuoteHash common.Hash) []byte {
	return append(PricefeedPriceKey, baseQuoteHash.Bytes()...)
}

func GetPricefeedRelayerStoreKey(oracleBase, oracleQuote string, relayer sdk.AccAddress) []byte {
	return append(GetPricefeedRelayerStorePrefix(GetBaseQuoteHash(oracleBase, oracleQuote)), relayer.Bytes()...)
}

func GetPricefeedRelayerStorePrefix(baseQuoteHash common.Hash) []byte {
	return append(PricefeedRelayerKey, baseQuoteHash.Bytes()...)
}

func GetCoinbasePriceStoreKey(key string, timestamp uint64) []byte {
	timeKey := sdk.Uint64ToBigEndian(timestamp)
	return append(append(CoinbasePriceKey, []byte(key)...), timeKey...)
}

func GetCoinbasePriceStoreIterationKey(key string) []byte {
	return append(CoinbasePriceKey, []byte(key)...)
}

func GetChainlinkPriceStoreKey(feedId string) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	buf := make([]byte, 0, len(ChainlinkPriceKey)+len(feedIdBz))
	buf = append(buf, ChainlinkPriceKey...)
	buf = append(buf, feedIdBz...)
	return buf
}

func getPaddedFeedIdBz(feedId string) string {
	return fmt.Sprintf("%20s", feedId)
}

func GetSymbolHistoricalPriceRecordsKey(oracleType OracleType, symbol string) []byte {
	return append(SymbolHistoricalPriceRecordsPrefix, []byte(fmt.Sprintf("%s_%s", oracleType.String(), symbol))...)
}

// providerDelimiter is the delimiter used to enforce uniqueness in the provider symbol.
const providerDelimiter = "@@@"

func GetDelimitedProvider(provider string) string {
	return fmt.Sprintf("%s%s", provider, providerDelimiter)
}

func GetProviderInfoKey(provider string) []byte {
	return append(ProviderInfoPrefix, []byte(GetDelimitedProvider(provider))...)
}

func GetProviderIndexKey(providerAddress sdk.AccAddress) []byte {
	return append(ProviderIndexPrefix, providerAddress.Bytes()...)
}

func GetProviderPricePrefix(provider string) []byte {
	p := GetDelimitedProvider(provider)
	buf := make([]byte, 0, len(provider))
	buf = append(buf, ProviderPricePrefix...)
	buf = append(buf, []byte(p)...)
	return buf
}

func GetProviderPriceKey(provider, symbol string) []byte {
	p := GetDelimitedProvider(provider)
	buf := make([]byte, 0, len(p)+len(symbol))
	buf = append(buf, ProviderPricePrefix...)
	buf = append(buf, []byte(p)...)
	buf = append(buf, []byte(symbol)...)
	return buf
}
