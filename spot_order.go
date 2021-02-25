package sdk

import (
	"github.com/InjectiveLabs/sdk-go/typeddata"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"math/big"
)

// SpotOrder represents an unsigned Injective Spot order
type SpotOrder struct {
	ChainID      int64          `json:"chainId"`
	SubaccountID common.Hash    `json:"subaccountID"`
	Sender       common.Address `json:"sender"`
	FeeRecipient common.Address `json:"feeRecipient"`
	Expiry       uint64         `json:"expiry"`
	MarketID     common.Hash    `json:"marketID"`
	SupplyAmount *big.Int       `json:"supplyAmount"`
	DemandAmount *big.Int       `json:"demandAmount"`
	Salt         uint64         `json:"salt"`
	OrderType    uint8          `json:"orderType"`
	TriggerPrice *big.Int       `json:"triggerPrice"`

	// Cache hash for performance
	hash  *common.Hash
	price *big.Int
}

func (s *SignedSpotOrder) ComputeSpotPrice() *decimal.Decimal {
	var price decimal.Decimal
	supplyAmount, receiveAmount := decimal.NewFromBigInt(s.SupplyAmount, 0), decimal.NewFromBigInt(s.DemandAmount, 0)
	if s.OrderType%2 == 0 {
		price = supplyAmount.Div(receiveAmount)
	} else {
		price = receiveAmount.Div(supplyAmount)
	}
	return &price
}

// SignedSpotOrder represents a signed Injective Spot order
type SignedSpotOrder struct {
	SpotOrder
	Signature []byte `json:"signature"`
}

// ResetHash resets the cached order hash. Usually only required for testing.
func (o *SpotOrder) ResetHash() {
	o.hash = nil
}

func (o *SpotOrder) GetMakerAddressAndNonce() (common.Address, *big.Int) {
	return common.BytesToAddress(o.SubaccountID[:common.AddressLength]), new(big.Int).SetBytes(o.SubaccountID[common.AddressLength:common.HashLength])
}

// ComputeOrderHash computes a Injective Spot Order hash
func (o *SpotOrder) ComputeOrderHash() (common.Hash, error) {
	if o.hash != nil {
		return *o.hash, nil
	}

	chainID := math.NewHexOrDecimal256(o.ChainID)
	var domain = typeddata.TypedDataDomain{
		Name:              "Injective Protocol",
		Version:           "1.0.0",
		ChainId:           chainID,
		VerifyingContract: "Injective Chain",
	}

	var message = map[string]interface{}{
		"subaccountID": o.SubaccountID.Hex(),
		"sender":       o.Sender.Hex(),
		"feeRecipient": o.FeeRecipient.Hex(),
		"expiry":       o.Expiry,
		"marketID":     o.MarketID,
		"supplyAmount": o.SupplyAmount,
		"demandAmount": o.DemandAmount,
		"salt":         o.Salt,
		"orderType":    o.OrderType,
		"triggerPrice": o.TriggerPrice,
	}

	var typedData = typeddata.TypedData{
		Types:       eip712SpotOrderTypes,
		PrimaryType: "SpotOrder",
		Domain:      domain,
		Message:     message,
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return common.Hash{}, err
	}

	w := sha3.NewLegacyKeccak256()
	w.Write([]byte("\x19\x01"))
	w.Write([]byte(domainSeparator))
	w.Write([]byte(typedDataHash))

	hash := common.BytesToHash(w.Sum(nil))
	o.hash = &hash

	return hash, nil
}

// SignSpotOrder signs the Injective Spot order with the supplied Signer.
func SignSpotOrder(signer Signer, order *SpotOrder) (*SignedSpotOrder, error) {
	if order == nil {
		return nil, errors.New("cannot sign nil order")
	}

	orderHash, err := order.ComputeOrderHash()
	if err != nil {
		return nil, err
	}

	makerAddress, _ := order.GetMakerAddressAndNonce()

	ecSignature, err := signer.EthSign(orderHash.Bytes(), makerAddress)
	if err != nil {
		return nil, err
	}

	// Generate Injective EthSign Signature (append the signature type byte)
	signature := make([]byte, 66)
	signature[0] = ecSignature.V
	copy(signature[1:33], ecSignature.R[:])
	copy(signature[33:65], ecSignature.S[:])
	signature[65] = byte(EthSignSignature)
	signedOrder := &SignedSpotOrder{
		SpotOrder: *order,
		Signature: signature,
	}

	return signedOrder, nil
}
