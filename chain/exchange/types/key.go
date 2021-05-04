package types

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

const (
	// module name
	ModuleName = "exchange"

	// StoreKey to be used when creating the KVStore
	StoreKey  = ModuleName
	TStoreKey = "transient_exchange"
)
const SpotPriceDecimalPlaces = 18

var (
	// Keys for store prefixes
	DepositsPrefix             = []byte{0x01} // prefix for each key to a Deposit
	SubaccountTradeNoncePrefix = []byte{0x02} // prefix for each key to a Subaccount Trade Nonce

	SpotMarketsPrefix                = []byte{0x11} // prefix for each key to a spot market by (isEnabled, marketID)
	SpotLimitOrdersPrefix            = []byte{0x12} // prefix for each key to a spot order, by (marketID, direction, price level, order hash)
	SpotMarketOrdersPrefix           = []byte{0x13} // prefix for each key to a spot order, by (marketID, direction, price level, order hash)
	SpotLimitOrdersIndexPrefix       = []byte{0x14} // prefix for each key to a spot order index, by (marketID, direction, subaccountID, order hash)
	SpotMarketOrderIndicatorPrefix   = []byte{0x15} // prefix for each key to a spot market order indicator, by marketID and direction
	SpotMarketParamUpdateScheduleKey = []byte{0x16} // prefix for a key to save scheduled spot market params update

	DerivativeMarketPrefix               = []byte{0x21} // prefix for each key to a derivative market by (exchange address, isEnabled, marketID)
	DerivativeLimitOrdersPrefix          = []byte{0x22} // prefix for each key to an derivative limit order, by (marketID, direction, price level, order hash)
	DerivativeMarketOrdersPrefix         = []byte{0x23} // prefix for each key to a derivative order, by (marketID, direction, price level, order hash)
	DerivativeLimitOrdersIndexPrefix     = []byte{0x24} // prefix for each key to a derivative order index, by (marketID, direction, subaccountID, order hash)
	DerivativeLimitOrderIndicatorPrefix  = []byte{0x25} // prefix for each key to a derivative limit order indicator, by marketID and direction
	DerivativeMarketOrderIndicatorPrefix = []byte{0x26} // prefix for each key to a derivative market order indicator, by marketID and direction
	DerivativePositionsPrefix            = []byte{0x27} // prefix for each key to a Position
	DerivativePositionsIndexPrefix       = []byte{0x28} // prefix for each index key to a Position Key

	PerpetualMarketFundingPrefix  = []byte{0x31} // prefix for each key to a perpetual market's funding state
	PerpetualMarketInfoPrefix     = []byte{0x32} // prefix for each key to a perpetual market's market info
	ExpiryFuturesMarketInfoPrefix = []byte{0x33} // prefix for each key to a expiry futures market's market info
	NextFundingTimestampKey       = []byte{0x34} // key to the next perpetual market funding timestamp

)

// GetDepositKey provides the key to obtain a given subaccount's deposits for a given denom
func GetDepositKey(subaccountID common.Hash, denom string) []byte {
	return append(GetDepositKeyPrefixBySubaccountID(subaccountID), []byte(denom)...)
}

func GetDepositKeyPrefixBySubaccountID(subaccountID common.Hash) []byte {
	return append(DepositsPrefix, subaccountID.Bytes()...)
}

func ParseDepositStoreKey(key []byte) (subaccountID common.Hash, denom string) {
	subaccountEndIdx := common.HashLength
	subaccountID = common.BytesToHash(key[:subaccountEndIdx])
	denom = string(key[subaccountEndIdx:])
	return subaccountID, denom
}

// ParseDepositTransientStoreKey parses the deposit transient store key.
func ParseDepositTransientStoreKey(prefix, key []byte) (subaccountID common.Hash, denom string) {
	return ParseDepositStoreKey(key[len(prefix):])
}

// GetSubaccountTradeNonceKey provides the prefix to obtain a given subaccount's trade nonce.
func GetSubaccountTradeNonceKey(subaccountID common.Hash) []byte {
	return append(SubaccountTradeNoncePrefix, subaccountID.Bytes()...)
}

func GetSpotMarketKey(isEnabled bool) []byte {
	return append(SpotMarketsPrefix, enabledPrefix(isEnabled)...)
}

func GetSpotMarketTransientMarketsKeyPrefix(marketID common.Hash, isBuy bool) []byte {
	return append(SpotMarketsPrefix, MarketDirectionPrefix(marketID, isBuy)...)
}

func GetDerivativeMarketTransientMarketsKeyPrefix(marketID common.Hash, isBuy bool) []byte {
	return append(DerivativeLimitOrderIndicatorPrefix, MarketDirectionPrefix(marketID, isBuy)...)
}

// TODO: properly compute this without using decimal.Decimal
func getPaddedPrice(price sdk.Dec) string {
	priceString := price.String()
	temp, _ := decimal.NewFromString(priceString)
	priceString = temp.StringFixed(SpotPriceDecimalPlaces)
	leftSide := fmt.Sprintf("%032s", temp.Floor().String())
	priceComponents := strings.Split(priceString, ".")
	return leftSide + "." + priceComponents[1]
}

func GetLimitOrderByPriceKeyPrefix(marketID common.Hash, isBuy bool, price sdk.Dec, orderHash common.Hash) []byte {
	return GetOrderByPriceKeyPrefix(marketID, isBuy, price, orderHash)
}

func GetSpotLimitOrderIndexPrefix(marketID common.Hash, isBuy bool, subaccountID common.Hash) []byte {
	return append(SpotLimitOrdersIndexPrefix, getLimitOrderIndexSubaccountPrefix(marketID, isBuy, subaccountID)...)
}

func GetDerivativeLimitOrderIndexPrefix(marketID common.Hash, isBuy bool, subaccountID common.Hash) []byte {
	return append(DerivativeLimitOrdersIndexPrefix, getLimitOrderIndexSubaccountPrefix(marketID, isBuy, subaccountID)...)
}

func GetLimitOrderIndexKey(marketID common.Hash, isBuy bool, subaccountID, orderHash common.Hash) []byte {
	return append(getLimitOrderIndexSubaccountPrefix(marketID, isBuy, subaccountID), orderHash.Bytes()...)
}

// prefix containing marketID + isBuy + subaccountID
func getLimitOrderIndexSubaccountPrefix(marketID common.Hash, isBuy bool, subaccountID common.Hash) []byte {
	return append(MarketDirectionPrefix(marketID, isBuy), subaccountID.Bytes()...)
}

func GetOrderByPriceKeyPrefix(marketID common.Hash, isBuy bool, price sdk.Dec, orderHash common.Hash) []byte {
	return append(append(MarketDirectionPrefix(marketID, isBuy), []byte(getPaddedPrice(price))...), orderHash.Bytes()...)
}

// SpotMarketDirectionPriceHashPrefix turns a marketID + direction + price + order hash to prefix used to get a spot order from the store.
func SpotMarketDirectionPriceHashPrefix(marketID common.Hash, isBuy bool, price sdk.Dec, orderHash common.Hash) []byte {
	return append(append(MarketDirectionPrefix(marketID, isBuy), []byte(getPaddedPrice(price))...), orderHash.Bytes()...)
}

func GetDerivativeMarketKey(isEnabled bool) []byte {
	return append(DerivativeMarketPrefix, enabledPrefix(isEnabled)...)
}

func enabledPrefix(isEnabled bool) []byte {
	isEnabledByte := byte(0)
	if isEnabled {
		isEnabledByte = byte(1)
	}
	return []byte{isEnabledByte}
}

// OrdersByMarketDirectionPriceOrderHashPrefix turns a marketID + direction + price + order hash to prefix used to get an order from the store.
func OrdersByMarketDirectionPriceOrderHashPrefix(marketID common.Hash, orderHash common.Hash, price *big.Int, isLong bool) []byte {
	return append(ordersByMarketDirectionPricePrefix(marketID, price, isLong), orderHash.Bytes()...)
}

// orderIndexByMarketDirectionSubaccountPrefix allows to obtain prefix of exchange against a particular marketID, direction and price
func ordersByMarketDirectionPricePrefix(marketID common.Hash, price *big.Int, isLong bool) []byte {
	return append(MarketDirectionPrefix(marketID, isLong), common.LeftPadBytes(price.Bytes(), 32)...)
}

// OrderIndexByMarketDirectionSubaccountOrderHashPrefix turns a marketID + direction + subaccountID + order hash to prefix used to get an order from the store.
func OrderIndexByMarketDirectionSubaccountOrderHashPrefix(marketID common.Hash, isLong bool, subaccountID common.Hash, orderHash common.Hash) []byte {
	return append(OrderIndexByMarketDirectionSubaccountPrefix(marketID, subaccountID, isLong), orderHash.Bytes()...)
}

// OrderIndexByMarketDirectionSubaccountPrefix allows to obtain prefix of exchange against a particular marketID, subaccountID and direction
func OrderIndexByMarketDirectionSubaccountPrefix(marketID common.Hash, subaccountID common.Hash, isLong bool) []byte {
	return append(MarketDirectionPrefix(marketID, isLong), subaccountID.Bytes()...)
}

// MarketDirectionPrefix allows to obtain prefix against a particular marketID, direction
func MarketDirectionPrefix(marketID common.Hash, isLong bool) []byte {
	direction := byte(0)
	if isLong {
		direction = byte(1)
	}
	return append(marketID.Bytes(), direction)
}

// GetMarketIdDirectionFromTransientKey parses the marketID and direction from a transient Key.
// NOTE: this will not work for a normal key.
func GetMarketIdDirectionFromTransientKey(prefix, key []byte) (marketID common.Hash, isBuy bool) {
	marketID = common.BytesToHash(key[len(prefix) : common.HashLength+len(prefix)])
	isBuyBytes := key[common.HashLength+len(prefix):]
	return marketID, bytes.Equal(isBuyBytes, []byte{byte(0)})
}

// PositionByMarketSubaccountPrefix provides the prefix key to obtain a position for a given market and subaccount
func PositionByMarketSubaccountPrefix(marketID, subaccountID common.Hash) []byte {
	return append(marketID.Bytes(), subaccountID.Bytes()...)
}

// PositionIndexBySubaccountMarketPrefix provides the prefix key to obtain a position key for a given market and subaccount
func PositionIndexBySubaccountMarketPrefix(subaccountID, marketID common.Hash) []byte {
	return append(subaccountID.Bytes(), marketID.Bytes()...)
}

func ParsePositionTransientStoreKey(key []byte) (marketID, subaccountID common.Hash) {
	prefixLen := len(DerivativePositionsPrefix)
	marketIDEndIdx := common.HashLength + prefixLen
	marketID = common.BytesToHash(key[prefixLen:marketIDEndIdx])
	subaccountID = common.BytesToHash(key[marketIDEndIdx : marketIDEndIdx+common.HashLength])
	return marketID, subaccountID
}

func GetSubaccountAndMarketIDFromPositionKey(key []byte) (subaccountID, marketID common.Hash) {
	subaccountOffsetLen := common.HashLength
	subaccountID = common.BytesToHash(key[:subaccountOffsetLen])
	marketID = common.BytesToHash(key[subaccountOffsetLen : subaccountOffsetLen+common.HashLength])
	return subaccountID, marketID
}
