package zeroex

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
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


// StringSignedOrder is a special signed order structure
// for including in Msgs, because it consists of just string types besides ChainID.
type StringSignedOrder struct {
	// ChainID is a network identifier of the order.
	ChainID int64 `json:"chainId,omitempty"`
	// Exchange v3 contract address.
	ExchangeAddress string `json:"exchangeAddress,omitempty"`
	// Address that created the order.
	MakerAddress string `json:"makerAddress,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerAsset.
	MakerAssetData string `json:"makerAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerFee.
	MakerFeeAssetData string `json:"makerFeeAssetData,omitempty"`
	// Amount of makerAsset being offered by maker. Must be greater than 0.
	MakerAssetAmount string `json:"makerAssetAmount,omitempty"`
	// Amount of Fee Asset paid to feeRecipientAddress by maker when order is filled. If set to
	// 0, no transfer of Fee Asset from maker to feeRecipientAddress will be attempted.
	MakerFee string `json:"makerFee,omitempty"`
	// Address that is allowed to fill the order. If set to "0x0", any address is
	// allowed to fill the order.
	TakerAddress string `json:"takerAddress,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerAsset.
	TakerAssetData string `json:"takerAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerFee.
	TakerFeeAssetData string `json:"takerFeeAssetData,omitempty"`
	// Amount of takerAsset being bid on by maker. Must be greater than 0.
	TakerAssetAmount string `json:"takerAssetAmount,omitempty"`
	// Amount of Fee Asset paid to feeRecipientAddress by taker when order is filled. If set to
	// 0, no transfer of Fee Asset from taker to feeRecipientAddress will be attempted.
	TakerFee string `json:"takerFee,omitempty"`
	// Address that is allowed to call Exchange contract methods that affect this
	// order. If set to "0x0", any address is allowed to call these methods.
	SenderAddress string `json:"senderAddress,omitempty"`
	// Address that will receive fees when order is filled.
	FeeRecipientAddress string `json:"feeRecipientAddress,omitempty"`
	// Timestamp in seconds at which order expires.
	ExpirationTimeSeconds string `json:"expirationTimeSeconds,omitempty"`
	// Arbitrary number to facilitate uniqueness of the order's hash.
	Salt string `json:"salt,omitempty"`
	// Order signature.
	Signature string `json:"signature,omitempty"`
}

// StringUnsignedBareOrder is a special unsigned signed order structure
// for including in Msgs, because it consists of just string types for a bare 0x order without ChainId or ExchangeAddress.
type StringUnsignedBareOrder struct {
	// Address that created the order.
	MakerAddress string `json:"makerAddress,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerAsset.
	MakerAssetData string `json:"makerAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerFee.
	MakerFeeAssetData string `json:"makerFeeAssetData,omitempty"`
	// Amount of makerAsset being offered by maker. Must be greater than 0.
	MakerAssetAmount string `json:"makerAssetAmount,omitempty"`
	// Amount of Fee Asset paid to feeRecipientAddress by maker when order is filled. If set to
	// 0, no transfer of Fee Asset from maker to feeRecipientAddress will be attempted.
	MakerFee string `json:"makerFee,omitempty"`
	// Address that is allowed to fill the order. If set to "0x0", any address is
	// allowed to fill the order.
	TakerAddress string `json:"takerAddress,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerAsset.
	TakerAssetData string `json:"takerAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerFee.
	TakerFeeAssetData string `json:"takerFeeAssetData,omitempty"`
	// Amount of takerAsset being bid on by maker. Must be greater than 0.
	TakerAssetAmount string `json:"takerAssetAmount,omitempty"`
	// Amount of Fee Asset paid to feeRecipientAddress by taker when order is filled. If set to
	// 0, no transfer of Fee Asset from taker to feeRecipientAddress will be attempted.
	TakerFee string `json:"takerFee,omitempty"`
	// Address that is allowed to call Exchange contract methods that affect this
	// order. If set to "0x0", any address is allowed to call these methods.
	SenderAddress string `json:"senderAddress,omitempty"`
	// Address that will receive fees when order is filled.
	FeeRecipientAddress string `json:"feeRecipientAddress,omitempty"`
	// Timestamp in seconds at which order expires.
	ExpirationTimeSeconds string `json:"expirationTimeSeconds,omitempty"`
	// Arbitrary number to facilitate uniqueness of the order's hash.
	Salt string `json:"salt,omitempty"`
	// Order signature.
	Signature string `json:"signature,omitempty"`
}