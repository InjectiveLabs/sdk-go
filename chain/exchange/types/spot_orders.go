package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
	"golang.org/x/crypto/sha3"
	"strconv"
)

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

func (o *SpotLimitOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *SpotMarketOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *SpotOrder) CheckDecimalScale(maxPriceScaleDecimals, maxQuantityScaleDecimals uint32) error {
	if checkIfExceedDecimals(o.OrderInfo.Price, maxPriceScaleDecimals) {
		return sdkerrors.Wrapf(ErrInvalidPrice, "price exceeds market maximum price decimals %d", maxPriceScaleDecimals)
	}
	if checkIfExceedDecimals(o.OrderInfo.Quantity, maxQuantityScaleDecimals) {
		return sdkerrors.Wrapf(ErrInvalidQuantity, "quantity exceeds market maximum quantity decimals %d", maxQuantityScaleDecimals)
	}
	return nil
}

func (m *SpotOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *SpotLimitOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
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
		// for a limit buy in the ETH/USDT market, denom is USDT and balanceHold is (1 + takerFee)*(price * quantity)
		denom = market.QuoteDenom
		balanceHold = m.OrderInfo.GetNotional().Add(m.OrderInfo.GetFeeAmount(market.TakerFeeRate))
	} else {
		// for a limit sell in the ETH/USDT market, denom is ETH and balanceHold is just quantity
		denom = market.BaseDenom
		balanceHold = m.OrderInfo.Quantity
	}

	return balanceHold, denom
}

func (m *SpotLimitOrder) GetUnfilledMarginHoldAndMarginDenom(market *SpotMarket) (sdk.Dec, string) {
	var denom string
	var balanceHold sdk.Dec
	if m.IsBuy() {
		// for a limit buy in the ETH/USDT market, denom is USDT and fillable amount is BalanceHold is (1 + makerFee)*(price * quantity)
		// (takerFee - makerFee) is already refunded
		denom = market.QuoteDenom
		balanceHold = m.GetUnfilledNotional().Add(m.GetUnfilledFeeAmount(market.MakerFeeRate))
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

// Calculate the balance hold for the market order.
// availableBalance should be in the margin denom
func (m *SpotOrder) CheckMarketOrderBalanceHold(market *SpotMarket, availableBalance, bestPrice sdk.Dec) (sdk.Dec, error) {
	var balanceHold sdk.Dec

	if m.IsBuy() {
		// required margin for best sell price = bestPrice * quantity * (1 + takerFeeRate)
		requiredMarginForBestPrice := bestPrice.Mul(m.OrderInfo.Quantity).Mul(sdk.OneDec().Add(market.TakerFeeRate))
		requiredMarginForWorstPrice := m.OrderInfo.Price.Mul(m.OrderInfo.Quantity).Mul(sdk.OneDec().Add(market.TakerFeeRate))
		requiredMargin := sdk.MaxDec(requiredMarginForBestPrice, requiredMarginForWorstPrice)
		if requiredMargin.GT(availableBalance) {
			return sdk.Dec{}, sdkerrors.Wrapf(ErrInsufficientDeposit, "Required Margin %s exceeds availableBalance %s", requiredMargin.String(), availableBalance.String())
		}
		balanceHold = requiredMargin
	} else {
		// required margin for market sells just equals the quantity being sold
		if availableBalance.LT(m.OrderInfo.Quantity) {
			return sdk.Dec{}, sdkerrors.Wrapf(ErrInsufficientDeposit, "Required Sell Quantity %s exceeds availableBalance %s", m.OrderInfo.Quantity.String(), availableBalance.String())
		}
		balanceHold = m.OrderInfo.Quantity
	}
	return balanceHold, nil
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
