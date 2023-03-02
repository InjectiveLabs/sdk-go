package chain

import (
	"context"
	"strconv"
	"time"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
	"golang.org/x/crypto/sha3"
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

type OrderHashes struct {
	Spot       []common.Hash
	Derivative []common.Hash
}

var domain = gethsigner.TypedDataDomain{
	Name:              "Injective Protocol",
	Version:           "2.0.0",
	ChainId:           ethmath.NewHexOrDecimal256(888),
	VerifyingContract: "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
	Salt:              "0x0000000000000000000000000000000000000000000000000000000000000000",
}

func (c *chainClient) ComputeOrderHashes(spotOrders []exchangetypes.SpotOrder, derivativeOrders []exchangetypes.DerivativeOrder) (OrderHashes, error) {
	if len(spotOrders)+len(derivativeOrders) == 0 {
		return OrderHashes{}, nil
	}

	orderHashes := OrderHashes{}

	// protect nonce used in this function
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	for _, o := range spotOrders {
		c.nonce++

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
			"Salt":         strconv.Itoa(int(c.nonce)),
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
			return OrderHashes{}, err
		}
		typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
		if err != nil {
			return OrderHashes{}, err
		}

		w := sha3.NewLegacyKeccak256()
		w.Write([]byte("\x19\x01"))
		w.Write([]byte(domainSeparator))
		w.Write([]byte(typedDataHash))

		hash := common.BytesToHash(w.Sum(nil))
		orderHashes.Spot = append(orderHashes.Spot, hash)
	}

	for _, o := range derivativeOrders {
		c.nonce++

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
			"Salt":         strconv.Itoa(int(c.nonce)),
		}
		typedData := gethsigner.TypedData{
			Types:       eip712OrderTypes,
			PrimaryType: "DerivativeOrder",
			Domain:      domain,
			Message:     message,
		}
		domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
		if err != nil {
			return OrderHashes{}, err
		}
		typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
		if err != nil {
			return OrderHashes{}, err
		}

		w := sha3.NewLegacyKeccak256()
		w.Write([]byte("\x19\x01"))
		w.Write([]byte(domainSeparator))
		w.Write([]byte(typedDataHash))

		hash := common.BytesToHash(w.Sum(nil))
		orderHashes.Derivative = append(orderHashes.Derivative, hash)
	}

	return orderHashes, nil
}

func (c *chainClient) RoutineUpdateNounce() context.CancelFunc {
	c.updateNounce()

	ctx, cancel := context.WithCancel(context.Background())

	ticker := time.NewTicker(10 * time.Second)

	go func(ctx context.Context) {
		for {
			select {
			case <-ticker.C:
				c.updateNounce()
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	return cancel
}

func (c *chainClient) updateNounce() {
	// get nonce
	subaccountId := c.DefaultSubaccount(c.ctx.FromAddress)
	res, err := c.GetSubAccountNonce(context.Background(), subaccountId)
	if err != nil {
		c.logger.Errorln("[INJ-GO-SDK] Failed to get nonce: ", err)
	} else {
		c.syncMux.Lock()
		defer c.syncMux.Unlock()

		c.nonce = res.Nonce
	}
}
