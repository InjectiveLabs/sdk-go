package types

import (
	"strconv"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/crypto/sha3"
)

// GetRequiredBinaryOptionsOrderMargin returns the required margin for a binary options trade (or order) at a given price
func GetRequiredBinaryOptionsOrderMargin(
	price math.LegacyDec,
	quantity math.LegacyDec,
	oracleScaleFactor uint32,
	isBuy bool,
	isReduceOnly bool,
) math.LegacyDec {
	if isReduceOnly {
		return math.LegacyZeroDec()
	}

	if isBuy {
		return price.Mul(quantity)
	}
	return GetScaledPrice(math.LegacyOneDec(), oracleScaleFactor).Sub(price).Mul(quantity)
}

func (t OrderType) IsBuy() bool {
	switch t {
	case OrderType_BUY, OrderType_STOP_BUY, OrderType_TAKE_BUY, OrderType_BUY_PO, OrderType_BUY_ATOMIC:
		return true
	case OrderType_SELL, OrderType_STOP_SELL, OrderType_TAKE_SELL, OrderType_SELL_PO, OrderType_SELL_ATOMIC:
		return false
	}
	return false
}

func (t OrderType) IsPostOnly() bool {
	switch t {
	case OrderType_BUY_PO, OrderType_SELL_PO:
		return true
	default:
		return false
	}
}

func (t OrderType) IsConditional() bool {
	switch t {
	case OrderType_STOP_BUY,
		OrderType_STOP_SELL,
		OrderType_TAKE_BUY,
		OrderType_TAKE_SELL:
		return true
	}
	return false
}

func (t OrderType) IsAtomic() bool {
	switch t {
	case OrderType_BUY_ATOMIC,
		OrderType_SELL_ATOMIC:
		return true
	}
	return false
}

func (m *OrderInfo) GetNotional() math.LegacyDec {
	return m.Quantity.Mul(m.Price)
}

func (m *OrderInfo) GetFeeAmount(fee math.LegacyDec) math.LegacyDec {
	return m.GetNotional().Mul(fee)
}

func (m *OrderInfo) IsFromDefaultSubaccount() bool {
	return IsDefaultSubaccountID(common.HexToHash(m.SubaccountId))
}

func (m *OrderInfo) GetPrice() math.LegacyDec {
	return m.Price
}

func (m *OrderInfo) GetQuantity() math.LegacyDec {
	return m.Quantity
}

var eip712OrderTypes = apitypes.Types{
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

func ComputeOrderHash(marketId, subaccountId, feeRecipient, price, quantity, margin, triggerPrice, orderType string, nonce uint32) (common.Hash, error) {
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
			"SubaccountId": subaccountId,
			"FeeRecipient": feeRecipient,
			"Price":        price,
			"Quantity":     quantity,
		},
		"Margin":       margin,
		"OrderType":    orderType,
		"TriggerPrice": triggerPrice,
		"Salt":         strconv.Itoa(int(nonce)),
	}

	var typedData = apitypes.TypedData{
		Types:       eip712OrderTypes,
		PrimaryType: "DerivativeOrder",
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

func (m *MarketStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MarketStatus_value, data, "MarketStatus")
	if err != nil {
		return err
	}
	*m = MarketStatus(value)
	return nil
}

func (m *ExecutionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExecutionType_value, data, "ExecutionType")
	if err != nil {
		return err
	}
	*m = ExecutionType(value)
	return nil
}

func (m *OrderType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OrderType_value, data, "OrderType")
	if err != nil {
		return err
	}
	*m = OrderType(value)
	return nil
}

// GetOrderIdentifier returns the cid of an order if it exists, otherwise returns the order hash
func GetOrderIdentifier(orderHash, cid string) any {
	if cid != "" {
		return cid
	}
	return common.HexToHash(orderHash)
}
