package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
	"golang.org/x/crypto/sha3"
)

func (o *SpotOrder) ToSpotMarketOrder(balanceHold sdk.Dec, orderHash common.Hash) *SpotMarketOrder {
	return &SpotMarketOrder{
		OrderInfo:    o.OrderInfo,
		BalanceHold:  balanceHold,
		OrderHash:    orderHash.Bytes(),
		OrderType:    o.OrderType,
		TriggerPrice: o.TriggerPrice,
	}
}

func (o *SpotOrder) GetNewSpotLimitOrder(orderHash common.Hash) *SpotLimitOrder {
	return &SpotLimitOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Fillable:     o.OrderInfo.Quantity,
		TriggerPrice: o.TriggerPrice,
		OrderHash:    orderHash.Bytes(),
	}
}

func (o *SpotOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *SpotOrder) IsFromDefaultSubaccount() bool {
	return o.OrderInfo.IsFromDefaultSubaccount()
}

func (o *SpotLimitOrder) IsFromDefaultSubaccount() bool {
	return o.OrderInfo.IsFromDefaultSubaccount()
}

func (o *SpotLimitOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *SpotMarketOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *SpotMarketOrder) IsFromDefaultSubaccount() bool {
	return o.OrderInfo.IsFromDefaultSubaccount()
}

func (o *SpotMarketOrder) SdkAccAddress() sdk.AccAddress {
	return sdk.AccAddress(o.SubaccountID().Bytes()[:common.AddressLength])
}

func (o *SpotLimitOrder) SdkAccAddress() sdk.AccAddress {
	return sdk.AccAddress(o.SubaccountID().Bytes()[:common.AddressLength])
}

func (o *SpotLimitOrder) FeeRecipient() common.Address {
	return o.OrderInfo.FeeRecipientAddress()
}

func (o *SpotMarketOrder) FeeRecipient() common.Address {
	return o.OrderInfo.FeeRecipientAddress()
}

func (o *SpotOrder) CheckTickSize(minPriceTickSize, minQuantityTickSize sdk.Dec) error {
	if BreachesMinimumTickSize(o.OrderInfo.Price, minPriceTickSize) {
		return sdkerrors.Wrapf(ErrInvalidPrice, "price %s must be a multiple of the minimum price tick size %s", o.OrderInfo.Price.String(), minPriceTickSize.String())
	}
	if BreachesMinimumTickSize(o.OrderInfo.Quantity, minQuantityTickSize) {
		return sdkerrors.Wrapf(ErrInvalidQuantity, "quantity %s must be a multiple of the minimum quantity tick size %s", o.OrderInfo.Quantity.String(), minQuantityTickSize.String())
	}
	return nil
}

func (o *SpotOrder) IsBuy() bool {
	return o.OrderType.IsBuy()
}

func (m *SpotLimitOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *SpotLimitOrder) Hash() common.Hash {
	return common.BytesToHash(m.OrderHash)
}

func (m *SpotMarketOrder) Hash() common.Hash {
	return common.BytesToHash(m.OrderHash)
}

func (o *SpotOrder) IsConditional() bool {
	return o.OrderType.IsConditional()
}

func (o *SpotLimitOrder) IsConditional() bool {
	return o.OrderType.IsConditional()
}

func (o *SpotMarketOrder) IsConditional() bool {
	return o.OrderType.IsConditional()
}

func (m *SpotLimitOrder) GetUnfilledNotional() sdk.Dec {
	return m.Fillable.Mul(m.OrderInfo.Price)
}
func (m *SpotLimitOrder) GetUnfilledFeeAmount(fee sdk.Dec) sdk.Dec {
	return m.GetUnfilledNotional().Mul(fee)
}

func (m *SpotOrder) GetBalanceHoldAndMarginDenom(market *SpotMarket) (sdk.Dec, string) {
	var denom string
	var balanceHold sdk.Dec
	if m.IsBuy() {
		denom = market.QuoteDenom
		if m.OrderType.IsPostOnly() {
			// for a PO limit buy in the ETH/USDT market, denom is USDT and balanceHold is (1 + makerFee)*(price * quantity)
			balanceHold = m.OrderInfo.GetNotional()
			if market.MakerFeeRate.IsPositive() {
				balanceHold = balanceHold.Add(m.OrderInfo.GetFeeAmount(market.MakerFeeRate))
			}
		} else {
			// for a normal limit buy in the ETH/USDT market, denom is USDT and balanceHold is (1 + takerFee)*(price * quantity)
			balanceHold = m.OrderInfo.GetNotional().Add(m.OrderInfo.GetFeeAmount(market.TakerFeeRate))
		}
	} else {
		// for a limit sell in the ETH/USDT market, denom is ETH and balanceHold is just quantity
		denom = market.BaseDenom
		balanceHold = m.OrderInfo.Quantity
	}

	return balanceHold, denom
}

func (m *SpotLimitOrder) GetUnfilledMarginHoldAndMarginDenom(market *SpotMarket, isTransient bool) (sdk.Dec, string) {
	var denom string
	var balanceHold sdk.Dec
	if m.IsBuy() {
		var tradeFeeRate sdk.Dec

		if isTransient {
			tradeFeeRate = market.TakerFeeRate
		} else {
			tradeFeeRate = sdk.MaxDec(sdk.ZeroDec(), market.MakerFeeRate)
		}

		// for a resting limit buy in the ETH/USDT market, denom is USDT and fillable amount is BalanceHold is (1 + makerFee)*(price * quantity) since (takerFee - makerFee) is already refunded
		denom = market.QuoteDenom
		balanceHold = m.GetUnfilledNotional().Add(m.GetUnfilledFeeAmount(tradeFeeRate))
	} else {
		// for a limit sell in the ETH/USDT market, denom is ETH and balanceHold is just quantity
		denom = market.BaseDenom
		balanceHold = m.Fillable
	}

	return balanceHold, denom
}

func (m *SpotOrder) GetMarginDenom(market *SpotMarket) string {
	var denom string
	if m.IsBuy() {
		// for a market buy in the ETH/USDT market, margin denom is USDT
		denom = market.QuoteDenom
	} else {
		// for a market buy in the ETH/USDT market, margin denom is ETH
		denom = market.BaseDenom
	}
	return denom
}

// GetMarketOrderBalanceHold calculates the balance hold for the market order.
func (m *SpotOrder) GetMarketOrderBalanceHold(feeRate, bestPrice sdk.Dec) sdk.Dec {
	var balanceHold sdk.Dec

	if m.IsBuy() {
		// required margin for best sell price = bestPrice * quantity * (1 + feeRate)
		requiredMarginForBestPrice := bestPrice.Mul(m.OrderInfo.Quantity).Mul(sdk.OneDec().Add(feeRate))
		requiredMarginForWorstPrice := m.OrderInfo.Price.Mul(m.OrderInfo.Quantity).Mul(sdk.OneDec().Add(feeRate))
		requiredMargin := sdk.MaxDec(requiredMarginForBestPrice, requiredMarginForWorstPrice)
		balanceHold = requiredMargin
	} else {
		// required margin for market sells just equals the quantity being sold
		balanceHold = m.OrderInfo.Quantity
	}
	return balanceHold
}

func (m *SpotLimitOrder) ToTrimmed() *TrimmedSpotLimitOrder {
	return &TrimmedSpotLimitOrder{
		Price:     m.OrderInfo.Price,
		Quantity:  m.OrderInfo.Quantity,
		Fillable:  m.Fillable,
		IsBuy:     m.IsBuy(),
		OrderHash: common.BytesToHash(m.OrderHash).Hex(),
	}
}

// ComputeOrderHash computes the order hash for given spot limit order
func (o *SpotOrder) ComputeOrderHash(nonce uint32) (common.Hash, error) {
	chainID := ethmath.NewHexOrDecimal256(888)
	var domain = gethsigner.TypedDataDomain{
		Name:              "Injective Protocol",
		Version:           "2.0.0",
		ChainId:           chainID,
		VerifyingContract: "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
		Salt:              "0x0000000000000000000000000000000000000000000000000000000000000000",
	}

	triggerPrice := ""
	if o.TriggerPrice != nil {
		triggerPrice = o.TriggerPrice.String()
	}

	var message = map[string]interface{}{
		"MarketId": o.MarketId,
		"OrderInfo": map[string]interface{}{
			"SubaccountId": o.OrderInfo.SubaccountId,
			"FeeRecipient": o.OrderInfo.FeeRecipient,
			"Price":        o.OrderInfo.Price.String(),
			"Quantity":     o.OrderInfo.Quantity.String(),
		},
		"Salt":         strconv.Itoa(int(nonce)),
		"OrderType":    string(o.OrderType),
		"TriggerPrice": triggerPrice,
	}

	var typedData = gethsigner.TypedData{
		Types:       eip712OrderTypes,
		PrimaryType: "SpotOrder",
		Domain:      domain,
		Message:     message,
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return AuctionSubaccountID, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return AuctionSubaccountID, err
	}

	w := sha3.NewLegacyKeccak256()
	w.Write([]byte("\x19\x01"))
	w.Write([]byte(domainSeparator))
	w.Write([]byte(typedDataHash))

	hash := common.BytesToHash(w.Sum(nil))
	return hash, nil
}
