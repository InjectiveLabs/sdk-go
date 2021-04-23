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
	PricefeedInfoKey    = []byte{0x11}
	PricefeedPriceKey   = []byte{0x12}
	PricefeedRelayerKey = []byte{0x13}
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
