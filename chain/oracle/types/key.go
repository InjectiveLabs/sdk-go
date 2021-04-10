package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	ModuleName = "oracle"
	StoreKey   = ModuleName
)

var (
	// Keys for band store prefixes
	BandRefKey     = []byte{0x01}
	BandRelayerKey = []byte{0x02}

	// Keys for pricefeeder store prefixes
	PricefeedPriceKey   = []byte{0x11}
	PricefeedRelayerKey = []byte{0x12}
)

func GetBandRefStoreKey(symbol string) []byte {
	return append(BandRefKey, []byte(symbol)...)
}

func GetBandRelayerStoreKey(relayer sdk.AccAddress) []byte {
	return append(BandRelayerKey, relayer.Bytes()...)
}

func GetBaseQuoteHash(oracleBase, oracleQuote string) common.Hash {
	return crypto.Keccak256Hash([]byte(oracleBase + oracleQuote))
}

func GetPriceFeedPriceStoreKeyFromBaseQuote(oracleBase, oracleQuote string) []byte {
	return GetPriceFeedPriceStoreKey(GetBaseQuoteHash(oracleBase, oracleQuote))
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
