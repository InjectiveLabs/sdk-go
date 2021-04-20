package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/crypto/sha3"

	"strconv"
)

func (t OrderType) IsBuy() bool {
	switch t {
	case OrderType_BUY, OrderType_STOP_BUY, OrderType_TAKE_BUY:
		return true
	case OrderType_SELL, OrderType_STOP_SELL, OrderType_TAKE_SELL:
		return false
	}
	return false
}

func (m *OrderInfo) GetNotional() sdk.Dec {
	return m.Quantity.Mul(m.Price)
}

func (m *OrderInfo) GetFeeAmount(fee sdk.Dec) sdk.Dec {
	return m.GetNotional().Mul(fee)
}

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

func computeOrderHash(marketId, subaccountId, feeRecipient, price, quantity, margin, triggerPrice, orderType string, nonce uint32) (common.Hash, error) {
	chainID := ethmath.NewHexOrDecimal256(888)
	var domain = gethsigner.TypedDataDomain{
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

func (m *Direction) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Direction_value, data, "Direction")
	if err != nil {
		return err
	}
	*m = Direction(value)
	return nil
}
