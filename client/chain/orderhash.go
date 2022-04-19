package chain

import (
	"context"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
	"golang.org/x/crypto/sha3"
	"strconv"
)

var AuctionSubaccountID = common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")

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

var domain = gethsigner.TypedDataDomain{
	Name:              "Injective Protocol",
	Version:           "2.0.0",
	ChainId:           ethmath.NewHexOrDecimal256(888),
	VerifyingContract: "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
	Salt:              "0x0000000000000000000000000000000000000000000000000000000000000000",
}

func (c *chainClient) ComputeSpotOrderHash(orders []exchangetypes.SpotOrder) ([]common.Hash, error) {
	if len(orders) == 0 {
		return nil, nil
	}

	orderHashes := []common.Hash{}

	// get nonce
	res, err := c.GetSubAccountNonce(context.Background(), orders[0].SubaccountID())
	if err != nil {
		return nil, err
	}
	nonce := res.Nonce + 1

	for _, o := range orders {
		triggerPrice := ""
		if o.TriggerPrice != nil {
			triggerPrice = o.TriggerPrice.String()
		}

		message := map[string]interface{}{
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

		typedData := gethsigner.TypedData{
			Types:       eip712OrderTypes,
			PrimaryType: "SpotOrder",
			Domain:      domain,
			Message:     message,
		}

		domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
		if err != nil {
			return nil, err
		}
		typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
		if err != nil {
			return nil, err
		}

		w := sha3.NewLegacyKeccak256()
		w.Write([]byte("\x19\x01"))
		w.Write([]byte(domainSeparator))
		w.Write([]byte(typedDataHash))

		hash := common.BytesToHash(w.Sum(nil))
		orderHashes = append(orderHashes, hash)

		nonce += 1
	}

	return orderHashes, nil
}

func (c *chainClient) ComputeDerivativeOrderHash(orders []exchangetypes.DerivativeOrder) ([]common.Hash, error) {
	if len(orders) == 0 {
		return nil, nil
	}

	orderHashes := []common.Hash{}

	// get nonce
	res, err := c.GetSubAccountNonce(context.Background(), orders[0].SubaccountID())
	if err != nil {
		return nil, err
	}
	nonce := res.Nonce + 1

	for _, o := range orders {
		triggerPrice := ""
		if o.TriggerPrice != nil {
			triggerPrice = o.TriggerPrice.String()
		}

		message := map[string]interface{}{
			"MarketId": o.MarketId,
			"OrderInfo": map[string]interface{}{
				"SubaccountId": o.OrderInfo.SubaccountId,
				"FeeRecipient": o.OrderInfo.FeeRecipient,
				"Price":        o.OrderInfo.Price.String(),
				"Quantity":     o.OrderInfo.Quantity.String(),
			},
			"Margin":       o.Margin.String(),
			"OrderType":    string(o.OrderType),
			"TriggerPrice": triggerPrice,
			"Salt":         strconv.Itoa(int(nonce)),
		}

		typedData := gethsigner.TypedData{
			Types:       eip712OrderTypes,
			PrimaryType: "DerivativeOrder",
			Domain:      domain,
			Message:     message,
		}

		domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
		if err != nil {
			return nil, err
		}
		typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
		if err != nil {
			return nil, err
		}

		w := sha3.NewLegacyKeccak256()
		w.Write([]byte("\x19\x01"))
		w.Write([]byte(domainSeparator))
		w.Write([]byte(typedDataHash))

		hash := common.BytesToHash(w.Sum(nil))
		orderHashes = append(orderHashes, hash)

		nonce += 1
	}

	return orderHashes, nil
}
