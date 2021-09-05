package types

import (
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
	//Version = "bandchain-1"
	//
	//// PortID is the default port id that module binds to
	//PortID = "oracle"
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
	return append(append(CoinbasePriceKey, []byte(key)...))
}
