package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	// module name
	ModuleName = "orders"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	// ActiveOrdersIndexStoreKeyPrefix for order-by-hash store of active orders.
	ActiveOrdersIndexStoreKeyPrefix = []byte{0x01}

	// ArchiveOrdersIndexStoreKeyPrefix for order-by-hash store of archive orders.
	ArchiveOrdersIndexStoreKeyPrefix = []byte{0x02}

	// UnknownOrderHashStoreKeyPrefix for order hash to key used to get the filled amount (or cancelled status)
	UnknownOrderHashStoreKeyPrefix = []byte{0x03}

	//// CloseOrdersStoreKeyPrefix
	//CloseOrdersStoreKeyPrefix = []byte{0x04}

	// MarginInfo store.
	MarginInfoStoreKeyPrefix = []byte{0x05}

	// PositionInfo store.
	PositionInfoStoreKeyPrefix = []byte{0x06}

	// TradePairsStoreKeyPrefix for pair-by-hash store of trade pairs with asset data.
	TradePairsStoreKeyPrefix       = []byte{0x07}
	DerivativeMarketStoreKeyPrefix = []byte{0x08}

	// ActiveOrdersStoreKeyPrefix for order-by-hash store of active orders.
	ActiveOrdersStoreKeyPrefix = []byte{0x09}

	// ArchiveOrdersStoreKeyPrefix for order-by-hash store of archive orders.
	ArchiveOrdersStoreKeyPrefix = []byte{0x0a}

	// TECSeenHashStoreKeyPrefix for bool-by-hash store of boolean seen TEC hashes
	TECSeenHashStoreKeyPrefix = []byte{0x0b}

	// ModuleStatePrefix defines a collection of module state-related entries.
	ModuleStatePrefix = []byte{0xff}
)

// OrdersByMarketDirectionPriceOrderHashPrefix turns a marketID + direction + price + order hash to prefix used to get an order from the store.
func OrdersByMarketDirectionPriceOrderHashPrefix(marketID common.Hash, orderHash common.Hash, price *big.Int, isLong bool) []byte {
	return append(OrdersByMarketDirectionPricePrefix(marketID, price, isLong), orderHash.Bytes()...)
}

// OrderIndexByMarketDirectionSubaccountPrefix allows to obtain prefix of orders against a particular marketID, direction and price
func OrdersByMarketDirectionPricePrefix(marketID common.Hash, price *big.Int, isLong bool) []byte {
	return append(MarketDirectionPrefix(marketID, isLong), common.LeftPadBytes(price.Bytes(), 32)...)
}

// OrderIndexByMarketDirectionSubaccountOrderHashPrefix turns a marketID + direction + subaccountID + order hash to prefix used to get an order from the store.
func OrderIndexByMarketDirectionSubaccountOrderHashPrefix(marketID common.Hash, orderHash common.Hash, subaccountID common.Hash, isLong bool) []byte {
	return append(OrderIndexByMarketDirectionSubaccountPrefix(marketID, subaccountID, isLong), orderHash.Bytes()...)
}

// OrderIndexByMarketDirectionSubaccountPrefix allows to obtain prefix of orders against a particular marketID, subaccountID and direction
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

// TradePairsStoreKey turns a pair hash to key used to get it from the store.
func TradePairsStoreKey(hash common.Hash) []byte {
	return append(TradePairsStoreKeyPrefix, hash.Bytes()...)
}
