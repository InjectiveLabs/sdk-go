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
const PriceDecimalPlaces = 18
const DefaultQueryOrderbookLimit = 20
const Uint64BytesLen = 8

var (
	// Keys for store prefixes
	DepositsPrefix                       = []byte{0x01} // prefix for each key to a Deposit
	SubaccountTradeNoncePrefix           = []byte{0x02} // prefix for each key to a Subaccount Trade Nonce
	SubaccountOrderbookMetadataPrefix    = []byte{0x03} // prefix for each key to a Subaccount Orderbook Metadata
	SubaccountOrderPrefix                = []byte{0x04} // prefix for each key to a Subaccount derivative limit Order
	SubaccountMarketOrderIndicatorPrefix = []byte{0x05} // prefix for each key to a Subaccount market order indicator
	SubaccountLimitOrderIndicatorPrefix  = []byte{0x06} // prefix for each key to a Subaccount limit order indicator
	SpotExchangeEnabledKey               = []byte{0x07} // key for whether spot exchange is enabled
	DerivativeExchangeEnabledKey         = []byte{0x08} // key for whether derivative exchange is enabled

	SpotMarketsPrefix                = []byte{0x11} // prefix for each key to a spot market by (isEnabled, marketID)
	SpotLimitOrdersPrefix            = []byte{0x12} // prefix for each key to a spot order, by (marketID, direction, price level, order hash)
	SpotMarketOrdersPrefix           = []byte{0x13} // prefix for each key to a spot order, by (marketID, direction, price level, order hash)
	SpotLimitOrdersIndexPrefix       = []byte{0x14} // prefix for each key to a spot order index, by (marketID, direction, subaccountID, order hash)
	SpotMarketOrderIndicatorPrefix   = []byte{0x15} // prefix for each key to a spot market order indicator, by marketID and direction
	SpotMarketParamUpdateScheduleKey = []byte{0x16} // prefix for a key to save scheduled spot market params update

	DerivativeMarketPrefix                  = []byte{0x21} // prefix for each key to a derivative market by (exchange address, isEnabled, marketID)
	DerivativeLimitOrdersPrefix             = []byte{0x22} // prefix for each key to a derivative limit order, by (marketID, direction, price level, order hash)
	DerivativeMarketOrdersPrefix            = []byte{0x23} // prefix for each key to a derivative order, by (marketID, direction, price level, order hash)
	DerivativeLimitOrdersIndexPrefix        = []byte{0x24} // prefix for each key to a derivative order index, by (marketID, direction, subaccountID, order hash)
	DerivativeLimitOrderIndicatorPrefix     = []byte{0x25} // prefix for each key to a derivative limit order indicator, by marketID and direction
	DerivativeMarketOrderIndicatorPrefix    = []byte{0x26} // prefix for each key to a derivative market order indicator, by marketID and direction
	DerivativePositionsPrefix               = []byte{0x27} // prefix for each key to a Position
	DerivativeMarketParamUpdateScheduleKey  = []byte{0x28} // prefix for a key to save scheduled derivative market params update
	DerivativeMarketScheduledSettlementInfo = []byte{0x29} // prefix for a key to save scheduled derivative market settlements

	PerpetualMarketFundingPrefix             = []byte{0x31} // prefix for each key to a perpetual market's funding state
	PerpetualMarketInfoPrefix                = []byte{0x32} // prefix for each key to a perpetual market's market info
	ExpiryFuturesMarketInfoPrefix            = []byte{0x33} // prefix for each key to a expiry futures market's market info
	ExpiryFuturesMarketInfoByTimestampPrefix = []byte{0x34} // prefix for each index key to a expiry futures market's market info

	IsFirstFeeCycleFinishedKey = []byte{0x3c} // key to the fee discount is first cycle finished

	TradingRewardCampaignInfoKey                  = []byte{0x40} // key to the TradingRewardCampaignInfo
	TradingRewardMarketQualificationPrefix        = []byte{0x41} // prefix for each key to a market's qualification/disqualification status
	TradingRewardMarketPointsMultiplierPrefix     = []byte{0x42} // prefix for each key to a market's FeePaidMultiplier
	TradingRewardCampaignRewardPoolPrefix         = []byte{0x43} // prefix for each key to a campaign's reward pool
	TradingRewardCurrentCampaignEndTimeKey        = []byte{0x44} // key to the current campaign's end time
	TradingRewardCampaignTotalPointsKey           = []byte{0x45} // key to the total trading reward points for the current campaign
	TradingRewardAccountPointsPrefix              = []byte{0x46} // prefix for each key to an account's current campaign reward points
	TradingRewardCampaignRewardPendingPoolPrefix  = []byte{0x47} // prefix for each key to a campaign's reward pending pool
	TradingRewardAccountPendingPointsPrefix       = []byte{0x48} // prefix for each key to an account's current campaign reward points
	TradingRewardCampaignTotalPendingPointsPrefix = []byte{0x49} // prefix to the total trading reward points for the current campaign

	FeeDiscountMarketQualificationPrefix                  = []byte{0x50} // prefix for each key to a market's qualification/disqualification status
	FeeDiscountBucketCountKey                             = []byte{0x51} // key to the fee discount bucket count
	FeeDiscountBucketDurationKey                          = []byte{0x52} // key to the fee discount bucket duration
	FeeDiscountCurrentBucketStartTimeKey                  = []byte{0x53} // key to the current bucket start timestamp
	FeeDiscountScheduleKey                                = []byte{0x54} // key to the fee discount schedule
	FeeDiscountTierInfoPrefix                             = []byte{0x55} // prefix to the fee discount tier info
	FeeDiscountAccountTierPrefix                          = []byte{0x56} // prefix to each account's fee discount tier and TTL timestamp
	FeeDiscountBucketAccountFeesPaidPrefix                = []byte{0x57} // prefix to each account's fee paid amount for a given bucket
	FeeDiscountAccountPastBucketTotalFeesPaidAmountPrefix = []byte{0x58} // prefix to each account's total past bucket fees paid amount FeeDiscountAccountIndicatorPrefix
	FeeDiscountAccountOrderIndicatorPrefix                = []byte{0x59} // prefix to each account's transient indicator if the account has placed an order that block that is relevant for fee discounts

	IsRegisteredDMMPrefix = []byte{0x60} // prefix to each account's is registered DMM address key
)

// GetFeeDiscountAccountFeesPaidInBucketKey provides the key for the account's fees paid in the given bucket
func GetFeeDiscountAccountFeesPaidInBucketKey(bucketStartTimestamp int64, account sdk.AccAddress) []byte {
	timeBz := sdk.Uint64ToBigEndian(uint64(bucketStartTimestamp))
	accountBz := account.Bytes()

	buf := make([]byte, 0, len(FeeDiscountBucketAccountFeesPaidPrefix)+len(timeBz)+len(accountBz))
	buf = append(buf, FeeDiscountBucketAccountFeesPaidPrefix...)
	buf = append(buf, timeBz...)
	buf = append(buf, accountBz...)
	return buf
}

// GetFeeDiscountTierKey provides the key for the fee discount tier for a given tier level
func GetFeeDiscountTierKey(tierLevel uint64) []byte {
	return append(FeeDiscountTierInfoPrefix, sdk.Uint64ToBigEndian(tierLevel)...)
}

func ParseFeeDiscountBucketAccountFeesPaidIteratorKey(key []byte) (bucketStartTimestamp int64, account sdk.AccAddress) {
	timeBz := key[:Uint64BytesLen]
	accountBz := key[Uint64BytesLen:]
	return int64(sdk.BigEndianToUint64(timeBz)), sdk.AccAddress(accountBz)
}

func ParseTradingRewardAccountPendingPointsKey(key []byte) (bucketStartTimestamp int64, account sdk.AccAddress) {
	timeBz := key[:Uint64BytesLen]
	accountBz := key[Uint64BytesLen:]
	return int64(sdk.BigEndianToUint64(timeBz)), sdk.AccAddress(accountBz)
}

// GetFeeDiscountPastBucketAccountFeesPaidKey provides the key for the account's total past bucket fees paid.
func GetFeeDiscountPastBucketAccountFeesPaidKey(account sdk.AccAddress) []byte {
	accountBz := account.Bytes()
	return append(FeeDiscountAccountPastBucketTotalFeesPaidAmountPrefix, accountBz...)
}

// GetFeeDiscountAccountOrderIndicatorKey provides the key for the transient indicator if the account has placed an order that block
func GetFeeDiscountAccountOrderIndicatorKey(account sdk.AccAddress) []byte {
	accountBz := account.Bytes()
	return append(FeeDiscountAccountOrderIndicatorPrefix, accountBz...)
}

// GetFeeDiscountAccountTierKey provides the key for the account's fee discount tier.
func GetFeeDiscountAccountTierKey(account sdk.AccAddress) []byte {
	accountBz := account.Bytes()

	buf := make([]byte, 0, len(FeeDiscountAccountTierPrefix)+len(accountBz))
	buf = append(buf, FeeDiscountAccountTierPrefix...)
	buf = append(buf, accountBz...)
	return buf
}

// GetIsRegisteredDMMKey provides the key for the registered DMM address
func GetIsRegisteredDMMKey(account sdk.AccAddress) []byte {
	accountBz := account.Bytes()

	buf := make([]byte, 0, len(IsRegisteredDMMPrefix)+len(accountBz))
	buf = append(buf, IsRegisteredDMMPrefix...)
	buf = append(buf, accountBz...)
	return buf
}

// GetFeeDiscountMarketQualificationKey provides the key for the market fee discount qualification status
func GetFeeDiscountMarketQualificationKey(marketID common.Hash) []byte {
	return append(FeeDiscountMarketQualificationPrefix, marketID.Bytes()...)
}

// GetCampaignRewardPoolKey provides the key for a reward pool for a given start time
func GetCampaignRewardPoolKey(startTimestamp int64) []byte {
	return append(TradingRewardCampaignRewardPoolPrefix, sdk.Uint64ToBigEndian(uint64(startTimestamp))...)
}

// GetCampaignRewardPendingPoolKey provides the key for a pending reward pool for a given start time
func GetCampaignRewardPendingPoolKey(startTimestamp int64) []byte {
	return append(TradingRewardCampaignRewardPendingPoolPrefix, sdk.Uint64ToBigEndian(uint64(startTimestamp))...)
}

// GetCampaignMarketQualificationKey provides the key for the market trading rewards qualification status
func GetCampaignMarketQualificationKey(marketID common.Hash) []byte {
	return append(TradingRewardMarketQualificationPrefix, marketID.Bytes()...)
}

// GetTradingRewardsMarketPointsMultiplierKey provides the key for the market trading rewards multiplier
func GetTradingRewardsMarketPointsMultiplierKey(marketID common.Hash) []byte {
	return append(TradingRewardMarketPointsMultiplierPrefix, marketID.Bytes()...)
}

// GetTradingRewardAccountPointsKey provides the key for the account's trading rewards points.
func GetTradingRewardAccountPointsKey(account sdk.AccAddress) []byte {
	return append(TradingRewardAccountPointsPrefix, account.Bytes()...)
}

// GetTradingRewardAccountPendingPointsPrefix provides the prefix for the account's trading rewards pending points.
func GetTradingRewardAccountPendingPointsPrefix(pendingPoolStartTimestamp int64) []byte {
	return append(TradingRewardAccountPendingPointsPrefix, sdk.Uint64ToBigEndian(uint64(pendingPoolStartTimestamp))...)
}

// GetTradingRewardAccountPendingPointsKey provides the key for the account's trading rewards pending points.
func GetTradingRewardAccountPendingPointsKey(account sdk.AccAddress, pendingPoolStartTimestamp int64) []byte {
	return append(append(TradingRewardAccountPendingPointsPrefix, sdk.Uint64ToBigEndian(uint64(pendingPoolStartTimestamp))...), account.Bytes()...)
}

// GetTradingRewardTotalPendingPointsKey provides the key for the total pending trading rewards points.
func GetTradingRewardTotalPendingPointsKey(pendingPoolStartTimestamp int64) []byte {
	return append(TradingRewardCampaignTotalPendingPointsPrefix, sdk.Uint64ToBigEndian(uint64(pendingPoolStartTimestamp))...)
}

// GetTradingRewardAccountPendingPointsStartTimestamp provides the start timestamp of the pending points pool.
func GetTradingRewardAccountPendingPointsStartTimestamp(pendingPoolStartTimestamp []byte) int64 {
	return int64(sdk.BigEndianToUint64(pendingPoolStartTimestamp))
}

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

func GetSubaccountOrderbookMetadataKey(marketID, subaccountID common.Hash, isBuy bool) []byte {
	return append(SubaccountOrderbookMetadataPrefix, getSubaccountOrderSuffix(marketID, subaccountID, isBuy)...)
}

func GetSubaccountMarketOrderIndicatorKey(marketID, subaccountID common.Hash) []byte {
	return append(SubaccountMarketOrderIndicatorPrefix, MarketSubaccountInfix(marketID, subaccountID)...)
}

func GetSubaccountLimitOrderIndicatorKey(marketID, subaccountID common.Hash) []byte {
	return append(SubaccountLimitOrderIndicatorPrefix, MarketSubaccountInfix(marketID, subaccountID)...)
}

func getSubaccountOrderSuffix(marketID, subaccountID common.Hash, isBuy bool) []byte {
	return append(MarketSubaccountInfix(marketID, subaccountID), getBoolPrefix(isBuy)...)
}

func GetSubaccountOrderKey(marketID, subaccountID common.Hash, isBuy bool, price sdk.Dec, orderHash common.Hash) []byte {
	// TODO use copy for greater efficiency
	return append(append(GetSubaccountOrderPrefixByMarketSubaccountDirection(marketID, subaccountID, isBuy), []byte(getPaddedPrice(price))...), orderHash.Bytes()...)
}

func GetSubaccountOrderPrefixByMarketSubaccountDirection(marketID, subaccountID common.Hash, isBuy bool) []byte {
	return append(SubaccountOrderPrefix, append(MarketSubaccountInfix(marketID, subaccountID), getBoolPrefix(isBuy)...)...)
}

func GetSpotMarketKey(isEnabled bool) []byte {
	return append(SpotMarketsPrefix, getBoolPrefix(isEnabled)...)
}

func GetSpotMarketTransientMarketsKeyPrefix(marketID common.Hash, isBuy bool) []byte {
	return append(SpotMarketsPrefix, MarketDirectionPrefix(marketID, isBuy)...)
}

func GetDerivativeMarketTransientMarketsKeyPrefix(marketID common.Hash, isBuy bool) []byte {
	return append(DerivativeLimitOrderIndicatorPrefix, MarketDirectionPrefix(marketID, isBuy)...)
}

func getPaddedPrice(price sdk.Dec) string {
	dec := decimal.NewFromBigInt(price.BigInt(), -18).StringFixed(PriceDecimalPlaces)
	components := strings.Split(dec, ".")
	naturalPart, decimalPart := components[0], components[1]
	return fmt.Sprintf("%032s.%s", naturalPart, decimalPart)
}

func GetLimitOrderByPriceKeyPrefix(marketID common.Hash, isBuy bool, price sdk.Dec, orderHash common.Hash) []byte {
	return GetOrderByPriceKeyPrefix(marketID, isBuy, price, orderHash)
}

func GetSpotLimitOrderIndexPrefix(marketID common.Hash, isBuy bool, subaccountID common.Hash) []byte {
	return append(SpotLimitOrdersIndexPrefix, GetLimitOrderIndexSubaccountPrefix(marketID, isBuy, subaccountID)...)
}

func GetDerivativeLimitOrderIndexPrefix(marketID common.Hash, isBuy bool, subaccountID common.Hash) []byte {
	return append(DerivativeLimitOrdersIndexPrefix, GetLimitOrderIndexSubaccountPrefix(marketID, isBuy, subaccountID)...)
}

func GetLimitOrderIndexKey(marketID common.Hash, isBuy bool, subaccountID, orderHash common.Hash) []byte {
	return append(GetLimitOrderIndexSubaccountPrefix(marketID, isBuy, subaccountID), orderHash.Bytes()...)
}

func GetTransientLimitOrderIndexIteratorPrefix(marketID common.Hash, isBuy bool, subaccountID common.Hash) []byte {
	return append(SpotLimitOrdersIndexPrefix, GetLimitOrderIndexSubaccountPrefix(marketID, isBuy, subaccountID)...)
}

// GetLimitOrderIndexSubaccountPrefix returns a prefix containing marketID + isBuy + subaccountID
func GetLimitOrderIndexSubaccountPrefix(marketID common.Hash, isBuy bool, subaccountID common.Hash) []byte {
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
	return append(DerivativeMarketPrefix, getBoolPrefix(isEnabled)...)
}

func getBoolPrefix(isEnabled bool) []byte {
	isEnabledByte := byte(0)
	if isEnabled {
		isEnabledByte = TrueByte
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

const TrueByte byte = byte(1)
const FalseByte byte = byte(0)

func IsTrueByte(bz []byte) bool {
	return bytes.Equal(bz, []byte{TrueByte})
}

// MarketDirectionPrefix allows to obtain prefix against a particular marketID, direction
func MarketDirectionPrefix(marketID common.Hash, isLong bool) []byte {
	direction := byte(0)
	if isLong {
		direction = TrueByte
	}
	return append(marketID.Bytes(), direction)
}

// GetMarketIdDirectionFromTransientKey parses the marketID and direction from a transient Key.
// NOTE: this will not work for a normal key.
func GetMarketIdDirectionFromTransientKey(prefix, key []byte) (marketID common.Hash, isBuy bool) {
	marketID = common.BytesToHash(key[len(prefix) : common.HashLength+len(prefix)])
	isBuyByte := key[common.HashLength+len(prefix)]
	return marketID, isBuyByte == TrueByte
}

// MarketSubaccountInfix provides the infix given a marketID and subaccountID
func MarketSubaccountInfix(marketID, subaccountID common.Hash) []byte {
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
	marketID = common.BytesToHash(key[:subaccountOffsetLen])
	subaccountID = common.BytesToHash(key[subaccountOffsetLen : subaccountOffsetLen+common.HashLength])

	return subaccountID, marketID
}

func GetSubaccountIDFromPositionKey(key []byte) (subaccountID common.Hash) {
	subaccountOffsetLen := common.HashLength
	subaccountID = common.BytesToHash(key[:subaccountOffsetLen])

	return subaccountID
}

func GetExpiryFuturesMarketInfoByTimestampKey(timestamp int64, marketID common.Hash) []byte {
	return append(ExpiryFuturesMarketInfoByTimestampPrefix, append(sdk.Uint64ToBigEndian(uint64(timestamp)), marketID.Bytes()...)...)
}
