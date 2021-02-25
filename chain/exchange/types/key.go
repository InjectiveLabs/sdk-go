package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

const (
	// module name
	ModuleName = "exchange"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)
const SpotPriceDecimalPlaces = 18

var (
	// Keys for store prefixes
	MarginInfoKey              = []byte{0x01} // prefix for each key to a MarginInfo
	UnknownOrderHashKey        = []byte{0x02} // prefix for each key to an unknown order hash
	TECSeenHashStoreKeyPrefix  = []byte{0x03} // prefix for the seen state of a TEC transaction hash
	PositionInfoStoreKeyPrefix = []byte{0x04} // prefix for each key to a PositionInfo

	SpotMarketKey                       = []byte{0x11} // prefix for each key to a spot market by (exchange address, isEnabled, marketID)
	SpotLimitOrdersKey                  = []byte{0x12} // prefix for each key to an active spot order, by (marketID, direction, price level, order hash)
	SpotMarketOrdersKey                 = []byte{0x12} // prefix for each key to an active spot order, by (marketID, direction, price level, order hash)
	SpotLimitOrdersBySubaccountIndexKey = []byte{0x13} // prefix for each key to an active spot order index, by (marketID, direction, subaccountID, order hash)
	SpotMarketOrderQuantityKey          = []byte{0x14} // prefix for each key to a cumulative spot market order quantity key, by marketID and direction
	ArchiveOrderHashKey                 = []byte{0x15} // prefix for each key to an archived spot order, by order hash

	DerivativeMarketKey               = []byte{0x21} // prefix for each key to a derivative market by (exchange address, isEnabled, marketID)
	ActiveOrdersKey                   = []byte{0x22} // prefix for each key to an active order, by (marketID, direction, price level, order hash)
	ActiveOrdersBySubaccountIndexKey  = []byte{0x23} // prefix for each key to an active order index, by (marketID, direction, subaccountID, order hash)
	ArchiveOrdersKey                  = []byte{0x24} // prefix for each key to an archived order, by (marketID, direction, price level, order hash)
	ArchiveOrdersBySubaccountIndexKey = []byte{0x25} // prefix for each key to an archived order index, by (marketID, direction, subaccountID, order hash)
)

func GetSpotMarketKey(exchangeAddress common.Address, isEnabled bool) []byte {
	return append(SpotMarketKey, marketByExchangeAddressEnabledPrefix(exchangeAddress, isEnabled)...)
}

func getPaddedPrice(price *decimal.Decimal) string {
	priceString := price.StringFixed(SpotPriceDecimalPlaces)
	leftSide := fmt.Sprintf("%032d", price.Floor().String())
	priceComponents := strings.Split(priceString, ".")
	return leftSide + "." + priceComponents[2]
}

// SpotMarketDirectionPriceHashPrefix turns a marketID + direction + price + order hash to prefix used to get a spot order from the store.
func SpotMarketDirectionPriceHashPrefix(marketID common.Hash, isBuy bool, price *decimal.Decimal, orderHash common.Hash) []byte {
	return append(append(MarketDirectionPrefix(marketID, isBuy), []byte(getPaddedPrice(price))...), orderHash.Bytes()...)
}

func GetDerivativeMarketKey(exchangeAddress common.Address, isEnabled bool) []byte {
	return append(DerivativeMarketKey, marketByExchangeAddressEnabledPrefix(exchangeAddress, isEnabled)...)
}

func marketByExchangeAddressEnabledPrefix(exchangeAddress common.Address, isEnabled bool) []byte {
	isEnabledByte := byte(0)
	if isEnabled {
		isEnabledByte = byte(1)
	}
	return append(exchangeAddress.Bytes(), isEnabledByte)
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

// MarketDirectionHashPrefix allows to obtain prefix against a particular marketID, direction, order hash
func MarketDirectionHashPrefix(marketID common.Hash, isLong bool, hash common.Hash) []byte {
	direction := byte(0)
	if isLong {
		direction = byte(1)
	}
	return append(append(marketID.Bytes(), direction), hash.Bytes()...)
}

// MarketDirectionPrefix allows to obtain prefix against a particular marketID, direction
func MarketDirectionPrefix(marketID common.Hash, isLong bool) []byte {
	direction := byte(0)
	if isLong {
		direction = byte(1)
	}
	return append(marketID.Bytes(), direction)
}

// MarginInfoByExchangeSubaccountBaseCurrencyPrefix provides the prefix key to obtain a given subaccount's base currency
// margin info in a given exchange
func MarginInfoByExchangeSubaccountBaseCurrencyPrefix(exchangeAddress common.Address, subaccountID common.Hash, baseCurrencyAddress common.Address) []byte {
	return append(MarginInfoByExchangeSubaccountPrefix(exchangeAddress, subaccountID), baseCurrencyAddress.Bytes()...)
}

// MarginInfoByExchangeSubaccountPrefix provides the prefix key to obtain a given subaccount's margin info in a given exchange
func MarginInfoByExchangeSubaccountPrefix(exchangeAddress common.Address, subaccountID common.Hash) []byte {
	return append(exchangeAddress.Bytes(), subaccountID.Bytes()...)
}

// PositionInfoByExchangeSubaccountMarketIDPrefix provides the prefix key to obtain a given subaccount's position info
// in a given market in a given exchange
func PositionInfoByExchangeSubaccountMarketIDPrefix(exchangeAddress common.Address, subaccountID common.Hash, marketID common.Hash) []byte {
	return append(exchangeAddress.Bytes(), append(subaccountID.Bytes(), marketID.Bytes()...)...)
}

// SpotMarketsStoreKey turns a pair hash to key used to get it from the store.
func SpotMarketsStoreKey(hash common.Hash) []byte {
	return append(SpotMarketKey, hash.Bytes()...)
}
