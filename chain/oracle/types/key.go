package types

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

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
	ParamsKey      = []byte{0x03}

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
	CleanupCursorKey                   = []byte{0x53} // incremental cleanup cursor for historical price records

	// ProviderInfoPrefix is the prefix for the Provider => ProviderInfo store.
	ProviderInfoPrefix = []byte{0x61}
	// ProviderIndexPrefix is the prefix for the ProviderAddress => Provider index store.
	ProviderIndexPrefix = []byte{0x62}
	// ProviderPricePrefix is the prefix for the Provider + symbol => PriceState store.
	ProviderPricePrefix = []byte{0x63}

	// PythPriceKey is the prefix for the priceID => PythPriceState store.
	PythPriceKey = []byte{0x71}

	// StorkPriceKey is the prefix for the priceID => StorkPriceState store.
	StorkPriceKey     = []byte{0x81}
	StorkPublisherKey = []byte{0x82}

	// ChainlinkDataStreamsPriceKey is the prefix for the feedID => ChainlinkDataStreamsPriceState store.
	ChainlinkDataStreamsPriceKey = []byte{0x91}
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

// ParseCoinbasePriceStoreIterKey parses a prefix-store iterator key under CoinbasePriceKey:
// []byte(symbol) || BigEndian(timestamp).
func ParseCoinbasePriceStoreIterKey(iterKey []byte) (symbol string, timestamp uint64, ok bool) {
	if len(iterKey) < 8 {
		return "", 0, false
	}
	return string(iterKey[:len(iterKey)-8]), binary.BigEndian.Uint64(iterKey[len(iterKey)-8:]), true
}

func GetStorkPriceStoreKey(key string) []byte {
	return append(StorkPriceKey, []byte(key)...)
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

// GetSymbolHistoricalPriceRecordPrefix returns SymbolHistoricalPriceRecordsPrefix || byte(oracleType) || symbol || 0x00.
// Per-record keys append an 8-byte big-endian int64 timestamp after this prefix.
func GetSymbolHistoricalPriceRecordPrefix(oracleType OracleType, symbol string) []byte {
	buf := make([]byte, 0, len(SymbolHistoricalPriceRecordsPrefix)+1+len(symbol)+1)
	buf = append(buf, SymbolHistoricalPriceRecordsPrefix...)
	buf = append(buf, byte(oracleType))
	buf = append(buf, symbol...)
	buf = append(buf, 0)
	return buf
}

// GetSymbolHistoricalPriceRecordKey returns the full store key for one historical price sample.
func GetSymbolHistoricalPriceRecordKey(oracleType OracleType, symbol string, timestamp int64) []byte {
	prefix := GetSymbolHistoricalPriceRecordPrefix(oracleType, symbol)
	ts := make([]byte, 8)
	binary.BigEndian.PutUint64(ts, uint64(timestamp))
	return append(prefix, ts...)
}

// ParseSymbolHistoricalPriceRecordKey parses a full oracle module key under SymbolHistoricalPriceRecordsPrefix.
//
//nolint:revive // four named return values are required to unambiguously convey (oracleType, symbol, timestamp, ok)
func ParseSymbolHistoricalPriceRecordKey(fullKey []byte) (oracleType OracleType, symbol string, timestamp int64, ok bool) {
	if len(fullKey) < len(SymbolHistoricalPriceRecordsPrefix)+1+1+1+8 {
		return 0, "", 0, false
	}
	if fullKey[0] != SymbolHistoricalPriceRecordsPrefix[0] {
		return 0, "", 0, false
	}
	rest := fullKey[len(SymbolHistoricalPriceRecordsPrefix):]
	if len(rest) < 1+1+1+8 {
		return 0, "", 0, false
	}
	ot := OracleType(rest[0])
	body := rest[1:]
	if len(body) < 9 {
		return 0, "", 0, false
	}
	tsOff := len(body) - 8
	if body[tsOff-1] != 0 {
		return 0, "", 0, false
	}
	sym := string(body[:tsOff-1])
	ts := int64(binary.BigEndian.Uint64(body[tsOff:]))
	return ot, sym, ts, true
}

// ProviderDelimiter is appended to provider identifiers in store keys and compound oracle keys.
const ProviderDelimiter = "@@@"

// ProviderCompoundKeyDelimiter separates the delimited provider prefix from the symbol in compound keys
// (e.g. AppendPriceRecord pair strings, ProviderAssistant keys). It is ProviderDelimiter + "/".
const ProviderCompoundKeyDelimiter = ProviderDelimiter + "/"

func GetDelimitedProvider(provider string) string {
	return fmt.Sprintf("%s%s", provider, ProviderDelimiter)
}

// JoinProviderCompoundKey returns the compound key for provider+symbol (price records, ProviderAssistant).
func JoinProviderCompoundKey(provider, symbol string) string {
	return GetDelimitedProvider(provider) + ProviderCompoundKeyDelimiter[len(ProviderDelimiter):] + symbol
}

// ParseProviderCompoundKey splits a compound provider price key (provider@@@/symbol).
func ParseProviderCompoundKey(key string) (provider, symbol string, ok bool) {
	provider, symbol, ok = strings.Cut(key, ProviderCompoundKeyDelimiter)
	return provider, symbol, ok && provider != "" && symbol != ""
}

// ValidateProviderDerivativeOracleLayout enforces canonical exchange-stored provider oracle fields:
// oracle base (or binary-options symbol) must be a plain symbol; compound keys are rejected, and if a
// compound-shaped string parses, oracle quote must match the embedded provider before rejecting as non-canonical.
// Symbols must not contain ProviderDelimiter, matching MsgRelayProviderPrices symbol validation.
func ValidateProviderDerivativeOracleLayout(oracleBaseOrSymbol, oracleQuoteOrProvider string) error {
	p, _, ok := ParseProviderCompoundKey(oracleBaseOrSymbol)
	if ok {
		if oracleQuoteOrProvider != p {
			return fmt.Errorf("oracle quote %q must match provider %q encoded in oracle base", oracleQuoteOrProvider, p)
		}
		return errors.New("oracle base must be a plain symbol for provider markets, not a compound key")
	}
	if strings.Contains(oracleBaseOrSymbol, ProviderCompoundKeyDelimiter) {
		return fmt.Errorf("oracle base must not contain %q", ProviderCompoundKeyDelimiter)
	}
	if strings.Contains(oracleBaseOrSymbol, ProviderDelimiter) {
		return fmt.Errorf("oracle base must not contain %q", ProviderDelimiter)
	}
	return nil
}

func GetProviderInfoKey(provider string) []byte {
	return append(ProviderInfoPrefix, []byte(GetDelimitedProvider(provider))...)
}

func GetProviderIndexKey(providerAddress sdk.AccAddress) []byte {
	return append(ProviderIndexPrefix, providerAddress.Bytes()...)
}

func GetProviderPricePrefix(provider string) []byte {
	p := GetDelimitedProvider(provider)
	buf := make([]byte, 0, len(ProviderPricePrefix)+len(p))
	buf = append(buf, ProviderPricePrefix...)
	buf = append(buf, []byte(p)...)
	return buf
}

func GetProviderPriceKey(provider, symbol string) []byte {
	p := GetDelimitedProvider(provider)
	buf := make([]byte, 0, len(ProviderPricePrefix)+len(p)+len(symbol))
	buf = append(buf, ProviderPricePrefix...)
	buf = append(buf, []byte(p)...)
	buf = append(buf, []byte(symbol)...)
	return buf
}

func GetPythPriceStoreKey(priceID common.Hash) []byte {
	return append(PythPriceKey, priceID.Bytes()...)
}

// GetChainlinkDataStreamsPriceStoreKey returns the store key for a Chainlink Data Streams price state.
// The feed ID is encoded as 32 raw bytes (same as Pyth price IDs), not as the UTF-8 bytes of the hex string.
func GetChainlinkDataStreamsPriceStoreKey(feedID string) []byte {
	return append(ChainlinkDataStreamsPriceKey, common.HexToHash(feedID).Bytes()...)
}

// GetChainlinkDataStreamsFeedIDFromIterKey returns the canonical "0x…" feed ID from a prefix-store iterator key suffix (32-byte hash).
func GetChainlinkDataStreamsFeedIDFromIterKey(iterKey []byte) string {
	if len(iterKey) == 32 {
		return common.BytesToHash(iterKey).Hex()
	}
	return string(iterKey)
}
