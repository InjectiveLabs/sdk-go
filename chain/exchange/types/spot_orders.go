package types

import (
	"strconv"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"golang.org/x/crypto/sha3"
)

type IOrderInfo interface {
	GetSubaccountId() string
	GetFeeRecipient() string
	GetPrice() math.LegacyDec
	GetQuantity() math.LegacyDec
	GetCid() string
}

var _ IOrderInfo = &OrderInfo{}

func ComputeSpotOrderHash(marketId, orderType, triggerPrice string, orderInfo IOrderInfo, nonce uint32) (common.Hash, error) {
	chainID := ethmath.NewHexOrDecimal256(888)
	var domain = apitypes.TypedDataDomain{
		Name:              "Injective Protocol",
		Version:           "2.0.0",
		ChainId:           chainID,
		VerifyingContract: "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
		Salt:              "0x0000000000000000000000000000000000000000000000000000000000000000",
	}

	var message = map[string]interface{}{
		"MarketId": marketId,
		"OrderInfo": map[string]interface{}{
			"SubaccountId": orderInfo.GetSubaccountId(),
			"FeeRecipient": orderInfo.GetFeeRecipient(),
			"Price":        orderInfo.GetPrice().String(),
			"Quantity":     orderInfo.GetQuantity().String(),
		},
		"Salt":         strconv.Itoa(int(nonce)),
		"OrderType":    orderType,
		"TriggerPrice": triggerPrice,
	}

	var typedData = apitypes.TypedData{
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
	w.Write(domainSeparator)
	w.Write(typedDataHash)

	hash := common.BytesToHash(w.Sum(nil))
	return hash, nil
}

func (m *SpotOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *SpotOrder) GetBalanceHoldAndMarginDenom(market *SpotMarket) (math.LegacyDec, string) {
	var denom string
	var balanceHold math.LegacyDec
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
