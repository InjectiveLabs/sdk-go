package zeroex

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"

	"github.com/InjectiveLabs/zeroex-go/wrappers"
)

// Order represents an unsigned 0x order
type Order struct {
	ChainID               *big.Int       `json:"chainId"`
	ExchangeAddress       common.Address `json:"exchangeAddress"`
	MakerAddress          common.Address `json:"makerAddress"`
	MakerAssetData        []byte         `json:"makerAssetData"`
	MakerFeeAssetData     []byte         `json:"makerFeeAssetData"`
	MakerAssetAmount      *big.Int       `json:"makerAssetAmount"`
	MakerFee              *big.Int       `json:"makerFee"`
	TakerAddress          common.Address `json:"takerAddress"`
	TakerAssetData        []byte         `json:"takerAssetData"`
	TakerFeeAssetData     []byte         `json:"takerFeeAssetData"`
	TakerAssetAmount      *big.Int       `json:"takerAssetAmount"`
	TakerFee              *big.Int       `json:"takerFee"`
	SenderAddress         common.Address `json:"senderAddress"`
	FeeRecipientAddress   common.Address `json:"feeRecipientAddress"`
	ExpirationTimeSeconds *big.Int       `json:"expirationTimeSeconds"`
	Salt                  *big.Int       `json:"salt"`

	// Cache hash for performance
	hash *common.Hash
}

// SignedOrder represents a signed 0x order
type SignedOrder struct {
	Order
	Signature []byte `json:"signature"`
}

// ResetHash resets the cached order hash. Usually only required for testing.
func (o *Order) ResetHash() {
	o.hash = nil
}

// ComputeOrderHash computes a 0x order hash
func (o *Order) ComputeOrderHash() (common.Hash, error) {
	if o.hash != nil {
		return *o.hash, nil
	}

	chainID := math.NewHexOrDecimal256(o.ChainID.Int64())
	var domain = gethsigner.TypedDataDomain{
		Name:              "0x Protocol",
		Version:           "3.0.0",
		ChainId:           chainID,
		VerifyingContract: o.ExchangeAddress.Hex(),
	}

	var message = map[string]interface{}{
		"makerAddress":          o.MakerAddress.Hex(),
		"takerAddress":          o.TakerAddress.Hex(),
		"senderAddress":         o.SenderAddress.Hex(),
		"feeRecipientAddress":   o.FeeRecipientAddress.Hex(),
		"makerAssetData":        o.MakerAssetData,
		"makerFeeAssetData":     o.MakerFeeAssetData,
		"takerAssetData":        o.TakerAssetData,
		"takerFeeAssetData":     o.TakerFeeAssetData,
		"salt":                  o.Salt.String(),
		"makerFee":              o.MakerFee.String(),
		"takerFee":              o.TakerFee.String(),
		"makerAssetAmount":      o.MakerAssetAmount.String(),
		"takerAssetAmount":      o.TakerAssetAmount.String(),
		"expirationTimeSeconds": o.ExpirationTimeSeconds.String(),
	}

	var typedData = gethsigner.TypedData{
		Types:       eip712OrderTypes,
		PrimaryType: "Order",
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

// SignOrder signs the 0x order with the supplied Signer.
func SignOrder(signer Signer, order *Order) (*SignedOrder, error) {
	if order == nil {
		return nil, errors.New("cannot sign nil order")
	}

	orderHash, err := order.ComputeOrderHash()
	if err != nil {
		return nil, err
	}

	ecSignature, err := signer.EthSign(orderHash.Bytes(), order.MakerAddress)
	if err != nil {
		return nil, err
	}

	// Generate 0x EthSign Signature (append the signature type byte)
	signature := make([]byte, 66)
	signature[0] = ecSignature.V
	copy(signature[1:33], ecSignature.R[:])
	copy(signature[33:65], ecSignature.S[:])
	signature[65] = byte(EthSignSignature)
	signedOrder := &SignedOrder{
		Order:     *order,
		Signature: signature,
	}

	return signedOrder, nil
}

// Trim converts the order to a TrimmedOrder, which is the format expected by
// our smart contracts. It removes the ChainID and ExchangeAddress fields.
func (s *SignedOrder) Trim() wrappers.Order {
	return wrappers.Order{
		MakerAddress:          s.MakerAddress,
		TakerAddress:          s.TakerAddress,
		FeeRecipientAddress:   s.FeeRecipientAddress,
		SenderAddress:         s.SenderAddress,
		MakerAssetAmount:      s.MakerAssetAmount,
		TakerAssetAmount:      s.TakerAssetAmount,
		MakerFee:              s.MakerFee,
		TakerFee:              s.TakerFee,
		ExpirationTimeSeconds: s.ExpirationTimeSeconds,
		Salt:                  s.Salt,
		MakerAssetData:        s.MakerAssetData,
		MakerFeeAssetData:     s.MakerFeeAssetData,
		TakerAssetData:        s.TakerAssetData,
		TakerFeeAssetData:     s.TakerFeeAssetData,
	}
}

func FromTrimmedOrder(order wrappers.Order) *Order {
	return &Order{
		MakerAddress:          order.MakerAddress,
		MakerAssetData:        order.MakerAssetData,
		MakerFeeAssetData:     order.MakerFeeAssetData,
		MakerAssetAmount:      order.MakerAssetAmount,
		MakerFee:              order.MakerFee,
		TakerAddress:          order.TakerAddress,
		TakerAssetData:        order.TakerAssetData,
		TakerFeeAssetData:     order.TakerFeeAssetData,
		TakerAssetAmount:      order.TakerAssetAmount,
		TakerFee:              order.TakerFee,
		SenderAddress:         order.SenderAddress,
		FeeRecipientAddress:   order.FeeRecipientAddress,
		ExpirationTimeSeconds: order.ExpirationTimeSeconds,
		Salt:                  order.Salt,
	}
}

// TrimmedOrderJSON is an unmodified JSON representation of a Trimmed Order
type TrimmedOrderJSON struct {
	MakerAddress          string `json:"makerAddress"`
	MakerAssetData        string `json:"makerAssetData"`
	MakerFeeAssetData     string `json:"makerFeeAssetData"`
	MakerAssetAmount      string `json:"makerAssetAmount"`
	MakerFee              string `json:"makerFee"`
	TakerAddress          string `json:"takerAddress"`
	TakerAssetData        string `json:"takerAssetData"`
	TakerFeeAssetData     string `json:"takerFeeAssetData"`
	TakerAssetAmount      string `json:"takerAssetAmount"`
	TakerFee              string `json:"takerFee"`
	SenderAddress         string `json:"senderAddress"`
	FeeRecipientAddress   string `json:"feeRecipientAddress"`
	ExpirationTimeSeconds string `json:"expirationTimeSeconds"`
	Salt                  string `json:"salt"`
}

// MarshalJSON implements a custom JSON marshaller for the Order type into Trimmed Order
func (o *Order) MarshalIntoTrimmedOrderJSON() ([]byte, error) {
	makerAssetData := "0x"
	if len(o.MakerAssetData) != 0 {
		makerAssetData = fmt.Sprintf("0x%s", common.Bytes2Hex(o.MakerAssetData))
	}
	// Note(albrow): Because of how our smart contracts work, most fields of an
	// order cannot be null. However, makerAssetFeeData and takerAssetFeeData are
	// the exception. For these fields, "0x" is used to indicate a null value.
	makerFeeAssetData := "0x"
	if len(o.MakerFeeAssetData) != 0 {
		makerFeeAssetData = fmt.Sprintf("0x%s", common.Bytes2Hex(o.MakerFeeAssetData))
	}
	takerAssetData := "0x"
	if len(o.TakerAssetData) != 0 {
		takerAssetData = fmt.Sprintf("0x%s", common.Bytes2Hex(o.TakerAssetData))
	}
	takerFeeAssetData := "0x"
	if len(o.TakerFeeAssetData) != 0 {
		takerFeeAssetData = fmt.Sprintf("0x%s", common.Bytes2Hex(o.TakerFeeAssetData))
	}

	trimmedOrderBytes, err := json.Marshal(TrimmedOrderJSON{
		MakerAddress:          strings.ToLower(o.MakerAddress.Hex()),
		MakerAssetData:        makerAssetData,
		MakerFeeAssetData:     makerFeeAssetData,
		MakerAssetAmount:      o.MakerAssetAmount.String(),
		MakerFee:              o.MakerFee.String(),
		TakerAddress:          strings.ToLower(o.TakerAddress.Hex()),
		TakerAssetData:        takerAssetData,
		TakerFeeAssetData:     takerFeeAssetData,
		TakerAssetAmount:      o.TakerAssetAmount.String(),
		TakerFee:              o.TakerFee.String(),
		SenderAddress:         strings.ToLower(o.SenderAddress.Hex()),
		FeeRecipientAddress:   strings.ToLower(o.FeeRecipientAddress.Hex()),
		ExpirationTimeSeconds: o.ExpirationTimeSeconds.String(),
		Salt:                  o.Salt.String(),
	})
	return trimmedOrderBytes, err
}

// SignedOrderJSON is an unmodified JSON representation of a SignedOrder
type SignedOrderJSON struct {
	ChainID               int64  `json:"chainId"`
	ExchangeAddress       string `json:"exchangeAddress"`
	MakerAddress          string `json:"makerAddress"`
	MakerAssetData        string `json:"makerAssetData"`
	MakerFeeAssetData     string `json:"makerFeeAssetData"`
	MakerAssetAmount      string `json:"makerAssetAmount"`
	MakerFee              string `json:"makerFee"`
	TakerAddress          string `json:"takerAddress"`
	TakerAssetData        string `json:"takerAssetData"`
	TakerFeeAssetData     string `json:"takerFeeAssetData"`
	TakerAssetAmount      string `json:"takerAssetAmount"`
	TakerFee              string `json:"takerFee"`
	SenderAddress         string `json:"senderAddress"`
	FeeRecipientAddress   string `json:"feeRecipientAddress"`
	ExpirationTimeSeconds string `json:"expirationTimeSeconds"`
	Salt                  string `json:"salt"`
	Signature             string `json:"signature"`
}

// MarshalJSON implements a custom JSON marshaller for the SignedOrder type
func (s SignedOrder) MarshalJSON() ([]byte, error) {
	makerAssetData := "0x"
	if len(s.MakerAssetData) != 0 {
		makerAssetData = fmt.Sprintf("0x%s", common.Bytes2Hex(s.MakerAssetData))
	}
	// Note(albrow): Because of how our smart contracts work, most fields of an
	// order cannot be null. However, makerAssetFeeData and takerAssetFeeData are
	// the exception. For these fields, "0x" is used to indicate a null value.
	makerFeeAssetData := "0x"
	if len(s.MakerFeeAssetData) != 0 {
		makerFeeAssetData = fmt.Sprintf("0x%s", common.Bytes2Hex(s.MakerFeeAssetData))
	}
	takerAssetData := "0x"
	if len(s.TakerAssetData) != 0 {
		takerAssetData = fmt.Sprintf("0x%s", common.Bytes2Hex(s.TakerAssetData))
	}
	takerFeeAssetData := "0x"
	if len(s.TakerFeeAssetData) != 0 {
		takerFeeAssetData = fmt.Sprintf("0x%s", common.Bytes2Hex(s.TakerFeeAssetData))
	}
	signature := "0x"
	if len(s.Signature) != 0 {
		signature = fmt.Sprintf("0x%s", common.Bytes2Hex(s.Signature))
	}

	signedOrderBytes, err := json.Marshal(SignedOrderJSON{
		ChainID:               s.ChainID.Int64(),
		ExchangeAddress:       strings.ToLower(s.ExchangeAddress.Hex()),
		MakerAddress:          strings.ToLower(s.MakerAddress.Hex()),
		MakerAssetData:        makerAssetData,
		MakerFeeAssetData:     makerFeeAssetData,
		MakerAssetAmount:      s.MakerAssetAmount.String(),
		MakerFee:              s.MakerFee.String(),
		TakerAddress:          strings.ToLower(s.TakerAddress.Hex()),
		TakerAssetData:        takerAssetData,
		TakerFeeAssetData:     takerFeeAssetData,
		TakerAssetAmount:      s.TakerAssetAmount.String(),
		TakerFee:              s.TakerFee.String(),
		SenderAddress:         strings.ToLower(s.SenderAddress.Hex()),
		FeeRecipientAddress:   strings.ToLower(s.FeeRecipientAddress.Hex()),
		ExpirationTimeSeconds: s.ExpirationTimeSeconds.String(),
		Salt:                  s.Salt.String(),
		Signature:             signature,
	})
	return signedOrderBytes, err
}

const addressHexLength = 42

// UnmarshalJSON implements a custom JSON unmarshaller for the SignedOrder type
func (s *SignedOrder) UnmarshalJSON(data []byte) error {
	var signedOrderJSON SignedOrderJSON
	err := json.Unmarshal(data, &signedOrderJSON)
	if err != nil {
		return err
	}
	var ok bool
	s.ChainID = big.NewInt(signedOrderJSON.ChainID)
	s.ExchangeAddress = common.HexToAddress(signedOrderJSON.ExchangeAddress)
	s.MakerAddress = common.HexToAddress(signedOrderJSON.MakerAddress)
	s.MakerAssetData = common.FromHex(signedOrderJSON.MakerAssetData)
	s.MakerFeeAssetData = common.FromHex(signedOrderJSON.MakerFeeAssetData)
	if signedOrderJSON.MakerAssetAmount != "" {
		s.MakerAssetAmount, ok = math.ParseBig256(signedOrderJSON.MakerAssetAmount)
		if !ok {
			s.MakerAssetAmount = nil
		}
	}
	if signedOrderJSON.MakerFee != "" {
		s.MakerFee, ok = math.ParseBig256(signedOrderJSON.MakerFee)
		if !ok {
			s.MakerFee = nil
		}
	}
	s.TakerAddress = common.HexToAddress(signedOrderJSON.TakerAddress)
	s.TakerAssetData = common.FromHex(signedOrderJSON.TakerAssetData)
	s.TakerFeeAssetData = common.FromHex(signedOrderJSON.TakerFeeAssetData)
	if signedOrderJSON.TakerAssetAmount != "" {
		s.TakerAssetAmount, ok = math.ParseBig256(signedOrderJSON.TakerAssetAmount)
		if !ok {
			s.TakerAssetAmount = nil
		}
	}
	if signedOrderJSON.TakerFee != "" {
		s.TakerFee, ok = math.ParseBig256(signedOrderJSON.TakerFee)
		if !ok {
			s.TakerFee = nil
		}
	}
	s.SenderAddress = common.HexToAddress(signedOrderJSON.SenderAddress)
	s.FeeRecipientAddress = common.HexToAddress(signedOrderJSON.FeeRecipientAddress)
	if signedOrderJSON.ExpirationTimeSeconds != "" {
		s.ExpirationTimeSeconds, ok = math.ParseBig256(signedOrderJSON.ExpirationTimeSeconds)
		if !ok {
			s.ExpirationTimeSeconds = nil
		}
	}
	if signedOrderJSON.Salt != "" {
		s.Salt, ok = math.ParseBig256(signedOrderJSON.Salt)
		if !ok {
			s.Salt = nil
		}
	}
	s.Signature = common.FromHex(signedOrderJSON.Signature)
	return nil
}

// OrderStatus represents the status of an order as returned from the 0x smart contracts
// as part of OrderInfo
type OrderStatus uint8

// OrderStatus values
const (
	OSInvalid OrderStatus = iota
	OSInvalidMakerAssetAmount
	OSInvalidTakerAssetAmount
	OSFillable
	OSExpired
	OSFullyFilled
	OSCancelled
	OSSignatureInvalid
	OSInvalidMakerAssetData
	OSInvalidTakerAssetData
)

// ContractEvent is an event emitted by a smart contract
type ContractEvent struct {
	BlockHash  common.Hash
	TxHash     common.Hash
	TxIndex    uint
	LogIndex   uint
	IsRemoved  bool
	Address    common.Address
	Kind       string
	Parameters interface{}
}

type contractEventJSON struct {
	BlockHash  common.Hash
	TxHash     common.Hash
	TxIndex    uint
	LogIndex   uint
	IsRemoved  bool
	Address    common.Address
	Kind       string
	Parameters json.RawMessage
}

// MarshalJSON implements a custom JSON marshaller for the ContractEvent type
func (c ContractEvent) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"blockHash":  c.BlockHash.Hex(),
		"txHash":     c.TxHash.Hex(),
		"txIndex":    c.TxIndex,
		"logIndex":   c.LogIndex,
		"isRemoved":  c.IsRemoved,
		"address":    c.Address,
		"kind":       c.Kind,
		"parameters": c.Parameters,
	}
	return json.Marshal(m)
}

// OrderEvent is the order event emitted by Mesh nodes on the "orders" topic
// when calling JSON-RPC method `mesh_subscribe`
type OrderEvent struct {
	// Timestamp is an order event timestamp that can be used for bookkeeping purposes.
	// If the OrderEvent represents a Mesh-specific event (e.g., ADDED, STOPPED_WATCHING),
	// the timestamp is when the event was generated. If the event was generated after
	// re-validating an order at the latest block height (e.g., FILLED, UNFUNDED, CANCELED),
	// then it is set to the latest block timestamp at which the order was re-validated.
	Timestamp time.Time `json:"timestamp"`
	// OrderHash is the EIP712 hash of the 0x order
	OrderHash common.Hash `json:"orderHash"`
	// SignedOrder is the signed 0x order struct
	SignedOrder *SignedOrder `json:"signedOrder"`
	// EndState is the end state of this order at the time this event was generated
	EndState OrderEventEndState `json:"endState"`
	// FillableTakerAssetAmount is the amount for which this order is still fillable
	FillableTakerAssetAmount *big.Int `json:"fillableTakerAssetAmount"`
	// ContractEvents contains all the contract events that triggered this orders re-evaluation.
	// They did not all necessarily cause the orders state change itself, only it's re-evaluation.
	// Since it's state _did_ change, at least one of them did cause the actual state change.
	ContractEvents []*ContractEvent `json:"contractEvents"`
}

type orderEventJSON struct {
	Timestamp                time.Time            `json:"timestamp"`
	OrderHash                string               `json:"orderHash"`
	SignedOrder              *SignedOrder         `json:"signedOrder"`
	EndState                 string               `json:"endState"`
	FillableTakerAssetAmount string               `json:"fillableTakerAssetAmount"`
	ContractEvents           []*contractEventJSON `json:"contractEvents"`
}

// MarshalJSON implements a custom JSON marshaller for the OrderEvent type
func (o OrderEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"timestamp":                o.Timestamp,
		"orderHash":                o.OrderHash.Hex(),
		"signedOrder":              o.SignedOrder,
		"endState":                 o.EndState,
		"fillableTakerAssetAmount": o.FillableTakerAssetAmount.String(),
		"contractEvents":           o.ContractEvents,
	})
}

// OrderEventEndState enumerates all the possible order event types. An OrderEventEndState describes the
// end state of a 0x order after revalidation
type OrderEventEndState string

// OrderEventEndState values
const (
	// ESInvalid is an event that is never emitted. It is here to discern between a declared but uninitialized OrderEventEndState
	ESInvalid = OrderEventEndState("INVALID")
	// ESOrderAdded means an order was successfully added to the Mesh node
	ESOrderAdded = OrderEventEndState("ADDED")
	// ESOrderFilled means an order was filled for a partial amount
	ESOrderFilled = OrderEventEndState("FILLED")
	// ESOrderFullyFilled means an order was fully filled such that it's remaining fillableTakerAssetAmount is 0
	ESOrderFullyFilled = OrderEventEndState("FULLY_FILLED")
	// ESOrderCancelled means an order was cancelled on-chain
	ESOrderCancelled = OrderEventEndState("CANCELLED")
	// ESOrderExpired means an order expired according to the latest block timestamp
	ESOrderExpired = OrderEventEndState("EXPIRED")
	// ESOrderUnexpired means an order is no longer expired. This can happen if a block re-org causes the latest
	// block timestamp to decline below the order's expirationTimestamp (rare and usually short-lived)
	ESOrderUnexpired = OrderEventEndState("UNEXPIRED")
	// ESOrderBecameUnfunded means an order has become unfunded. This happens if the maker transfers the balance /
	// changes their allowance backing an order
	ESOrderBecameUnfunded = OrderEventEndState("UNFUNDED")
	// ESOrderFillabilityIncreased means the fillability of an order has increased. Fillability for an order can
	// increase if a previously processed fill event gets reverted, or if a maker tops up their balance/allowance
	// backing an order
	ESOrderFillabilityIncreased = OrderEventEndState("FILLABILITY_INCREASED")
	// ESStoppedWatching means an order is potentially still valid but was removed for a different reason (e.g.
	// the database is full or the peer that sent the order was misbehaving). The order will no longer be watched
	// and no further events for this order will be emitted. In some cases, the order may be re-added in the
	// future.
	ESStoppedWatching = OrderEventEndState("STOPPED_WATCHING")
)
