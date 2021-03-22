package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/crypto/sha3"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
)

var eip712OrderTypes = gethsigner.Types{
	"EIP712Domain": {
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
		{Name: "verifyingContract", Type: "address"},
		{Name: "salt", Type: "bytes32"},
	},
	"OrderInfo": {
		{Name: "SubaccountId", Type: "string"},
		{Name: "FeeRecipient", Type: "string"},
		{Name: "Price", Type: "string"},
		{Name: "Quantity", Type: "string"},
	},
	"SpotOrder": {
		{Name: "MarketId", Type: "string"},
		{Name: "OrderInfo", Type: "OrderInfo"},
		{Name: "Salt", Type: "string"},
		{Name: "OrderType", Type: "string"},
		{Name: "TriggerPrice", Type: "string"},
	},
	"DerivativeOrder": {
		{Name: "MarketId", Type: "string"},
		{Name: "OrderInfo", Type: "OrderInfo"},
		{Name: "OrderType", Type: "string"},
		{Name: "Margin", Type: "string"},
		{Name: "TriggerPrice", Type: "string"},
		{Name: "Salt", Type: "string"},
	},
}

// ComputeOrderHash computes the order hash for given derivative limit order
func (o *DerivativeOrder) ComputeOrderHash() (common.Hash, error) {
	chainID := ethmath.NewHexOrDecimal256(888)
	var domain = gethsigner.TypedDataDomain{
		Name:              "Injective Protocol",
		Version:           "2.0.0",
		ChainId:           chainID,
		VerifyingContract: "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
		Salt:              "0x0000000000000000000000000000000000000000000000000000000000000000",
	}

	var message = map[string]interface{}{
		"MarketId": o.MarketId,
		"OrderInfo": map[string]interface{}{
			"SubaccountId": o.OrderInfo.SubaccountId,
			"FeeRecipient": o.OrderInfo.FeeRecipient,
			"Price":        o.OrderInfo.Price.String(),
			"Quantity":     o.OrderInfo.Quantity.String(),
		},
		"Margin":       o.Margin.String(),
		"OrderType":    string(o.OrderType),
		"TriggerPrice": o.TriggerPrice.String(),
		"Salt":         strconv.Itoa(int(o.Salt)),
	}

	var typedData = gethsigner.TypedData{
		Types:       eip712OrderTypes,
		PrimaryType: "DerivativeOrder",
		Domain:      domain,
		Message:     message,
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return ZeroHash, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return ZeroHash, err
	}

	w := sha3.NewLegacyKeccak256()
	w.Write([]byte("\x19\x01"))
	w.Write([]byte(domainSeparator))
	w.Write([]byte(typedDataHash))

	hash := common.BytesToHash(w.Sum(nil))
	return hash, nil
}

// ComputeOrderHash computes the order hash for given spot limit order
func (o *SpotOrder) ComputeOrderHash() (common.Hash, error) {
	chainID := ethmath.NewHexOrDecimal256(888)
	var domain = gethsigner.TypedDataDomain{
		Name:              "Injective Protocol",
		Version:           "2.0.0",
		ChainId:           chainID,
		VerifyingContract: "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
		Salt:              "0x0000000000000000000000000000000000000000000000000000000000000000",
	}

	var message = map[string]interface{}{
		"MarketId": o.MarketId,
		"OrderInfo": map[string]interface{}{
			"SubaccountId": o.OrderInfo.SubaccountId,
			"FeeRecipient": o.OrderInfo.FeeRecipient,
			"Price":        o.OrderInfo.Price.String(),
			"Quantity":     o.OrderInfo.Quantity.String(),
		},
		"Salt":         strconv.Itoa(int(o.Salt)),
		"OrderType":    string(o.OrderType),
		"TriggerPrice": o.TriggerPrice.String(),
	}

	var typedData = gethsigner.TypedData{
		Types:       eip712OrderTypes,
		PrimaryType: "SpotOrder",
		Domain:      domain,
		Message:     message,
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return ZeroHash, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return ZeroHash, err
	}

	w := sha3.NewLegacyKeccak256()
	w.Write([]byte("\x19\x01"))
	w.Write([]byte(domainSeparator))
	w.Write([]byte(typedDataHash))

	hash := common.BytesToHash(w.Sum(nil))
	return hash, nil
}

type MatchedMarketDirection struct {
	MarketId    common.Hash
	BuysExists  bool
	SellsExists bool
}

func (m *SpotOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (t OrderType) IsBuy() bool {
	switch t {
	case OrderType_BUY, OrderType_STOP_BUY, OrderType_TAKE_BUY:
		return true
	case OrderType_SELL, OrderType_STOP_SELL, OrderType_TAKE_SELL:
		return false
	}
	return false
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

func (m *OrderInfo) GetNotional() sdk.Dec {
	return m.Quantity.Mul(m.Price)
}

func (m *OrderInfo) GetFeeAmount(fee sdk.Dec) sdk.Dec {
	return m.GetNotional().Mul(fee)
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

func (o *DerivativeOrder) CheckMarginAndGetBalanceHold(market *DerivativeMarket) (balanceHold sdk.Dec, err error) {
	notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
	feeAmount := notional.Mul(market.TakerFeeRate)
	// Margin ≥ InitialMarginRatio * Price * Quantity
	if o.Margin.LT(market.InitialMarginRatio.Mul(notional)) {
		return sdk.Dec{}, ErrInsufficientOrderMargin
	}

	markPriceThreshold := o.ComputeInitialMarginRequirementMarkPriceThreshold(market.InitialMarginRatio)
	// For Buys: MarkPrice ≥ (Margin - Price * Quantity) / ((InitialMarginRatio - 1) * Quantity)
	// For Sells: MarkPrice ≤ (Margin + Price * Quantity) / ((1+ InitialMarginRatio) * Quantity)
	if o.OrderType.IsBuy() && market.MarkPrice.LT(markPriceThreshold) {
		return sdk.Dec{}, ErrInsufficientOrderMargin
	} else if !o.OrderType.IsBuy() && market.MarkPrice.GT(markPriceThreshold) {
		return sdk.Dec{}, ErrInsufficientOrderMargin
	}
	return o.Margin.Add(feeAmount), nil
}

func (o *DerivativeOrder) ComputeInitialMarginRequirementMarkPriceThreshold(initialMarginRatio sdk.Dec) sdk.Dec {
	notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
	var numerator, denominator sdk.Dec
	if o.OrderType.IsBuy() {
		numerator = o.Margin.Sub(notional)
		denominator = initialMarginRatio.Sub(sdk.OneDec()).Mul(o.OrderInfo.Quantity)
	} else {
		numerator = o.Margin.Add(notional)
		denominator = initialMarginRatio.Add(sdk.OneDec()).Mul(o.OrderInfo.Quantity)
	}
	return numerator.Quo(denominator)
}

// TODO do for market order
func (o *DerivativeOrder) CheckMarketOrderMarginHold(market *DerivativeMarket) (balanceHold sdk.Dec, err error) {
	//notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
	//feeAmount := notional.Mul(market.TakerFeeRate)
	// Margin ≥ InitialMarginRatio * Price * Quantity
	//if o.Margin.LT(market.InitialMarginRatio.Mul(notional)) {
	//	return sdk.Dec{}, ErrInsufficientOrderMargin
	//}

	//markPriceThreshold := o.ComputeInitialMarginRequirementMarkPriceThreshold(market.InitialMarginRatio)
	// For Buys: MarkPrice ≥ (Margin - Price * Quantity) / ((InitialMarginRatio - 1) * Quantity)
	// For Sells: MarkPrice ≤ (Margin + Price * Quantity) / ((1+ InitialMarginRatio) * Quantity)
	//if o.OrderType.IsBuy() && market.MarkPrice.LT(markPriceThreshold) {
	//	return sdk.Dec{}, ErrInsufficientOrderMargin
	//} else if !o.OrderType.IsBuy() && market.MarkPrice.GT(markPriceThreshold) {
	//	return sdk.Dec{}, ErrInsufficientOrderMargin
	//}
	//return o.Margin.Add(feeAmount), nil
	return
}
