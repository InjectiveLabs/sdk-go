package types

import (
	"errors"
	chainsdk "github.com/InjectiveLabs/sdk-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

// ToSignedOrder returns an appropriate SignedSpotOrder
func (b *BaseSpotOrder) ToSignedOrder() *chainsdk.SignedSpotOrder {
	o, err := bo2so(b)
	if err != nil {
		panic(err)
	}
	return o
}

// SpotOrderType encodes spot order type
type SpotOrderType uint8

const (
	LimitBuy   SpotOrderType = 1
	LimitSell  SpotOrderType = 2
	MarketBuy  SpotOrderType = 3
	MarketSell SpotOrderType = 4
)

type SpotMarketQuantity struct {
	MarketID common.Hash
	IsBuy    bool
	Quantity *big.Int
}

// demand amount - filled amount
func (o *SpotLimitOrder) GetFillableDemandAmount() *big.Int {
	return new(big.Int).Sub(BigNum(o.DemandAmount).Int(), BigNum(o.FilledAmount).Int())
}

func (o *SpotLimitOrder) GetSuppliableQuantity() *big.Int {
	// supplyQuantity * GetFillableDemandAmount / receiveQuantity
	return new(big.Int).Div(new(big.Int).Mul(BigNum(o.SupplyAmount).Int(), o.GetFillableDemandAmount()), BigNum(o.DemandAmount).Int())
}

func (b *BaseSpotOrder) GetSpotSupplyAsset(market *SpotMarket) common.Address {
	if b.OrderType == BaseSpotOrder_LIMIT_BUY || b.OrderType == BaseSpotOrder_MARKET_BUY ||
		b.OrderType == BaseSpotOrder_STOP_LIMIT_BUY || b.OrderType == BaseSpotOrder_TAKE_LIMIT_BUY ||
		b.OrderType == BaseSpotOrder_STOP_MARKET_BUY || b.OrderType == BaseSpotOrder_TAKE_MARKET_BUY {
		return common.HexToAddress(market.QuoteAsset)
	}
	return common.HexToAddress(market.BaseAsset)
}

func GetLimitOrderRequiredFunds(spotOrderType BaseSpotOrder_OrderType, sOrder *chainsdk.SignedSpotOrder, market *SpotMarket) *big.Int {
	if spotOrderType == BaseSpotOrder_LIMIT_BUY {
		return GetAssetAmountWithFeesApplied(sOrder.SupplyAmount, BigNum(market.GetTakerTxFee()).Int())
	}
	return sOrder.SupplyAmount
}

// ParseBig128 parses s as a 128 bit integer in decimal or hexadecimal syntax.
// Leading zeros are accepted. The empty string parses as zero.
func ParseBig128(s string) (*big.Int, error) {
	if v, ok := math.ParseBig256(s); !ok {
		return nil, errors.New("ParseBig128 failed for " + s)
	} else if v.BitLen() > 128 {
		return nil, errors.New("ParseBig128 failed for " + s + "due to BitLength error")
	} else {
		return v, nil
	}
}

// bo2so internal function converts model from *BaseSpotOrder to *chainsdk.SignedSpotOrder.
func bo2so(o *BaseSpotOrder) (*chainsdk.SignedSpotOrder, error) {
	if o == nil {
		return nil, nil
	}

	order := chainsdk.SpotOrder{
		ChainID:      o.ChainId,
		SubaccountID: common.HexToHash(o.SubaccountID),
		Sender:       common.HexToAddress(o.Sender),
		FeeRecipient: common.HexToAddress(o.FeeRecipient),
		Expiry:       o.Expiry,
		MarketID:     common.HexToHash(o.MarketID),
		Salt:         o.Salt,
		OrderType:    uint8(o.OrderType),
	}

	var err error
	order.SupplyAmount, err = ParseBig128(o.SupplyAmount)
	if err != nil {
		return nil, err
	}
	order.DemandAmount, err = ParseBig128(o.DemandAmount)
	if err != nil {
		return nil, err
	}
	order.TriggerPrice, err = ParseBig128(o.TriggerPrice)
	if err != nil {
		return nil, err
	}

	signedOrder := &chainsdk.SignedSpotOrder{
		SpotOrder: order,
	}

	// Orders do not need a signature if the maker address equals the MsgCreateSpotOrder sender address
	if len(o.Signature) != 0 {
		signedOrder.Signature = common.FromHex(o.Signature)
	}
	return signedOrder, nil
}
