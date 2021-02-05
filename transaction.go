package sdk

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"

	"github.com/InjectiveLabs/sdk-go/typeddata"
)

type Transaction struct {
	Salt                  *big.Int       `json:"salt"`
	SignerAddress         common.Address `json:"signerAddress"`
	Data                  []byte         `json:"data"`
	ExpirationTimeSeconds *big.Int       `json:"expirationTimeSeconds"`
	GasPrice              *big.Int       `json:"gasPrice"`
	Domain                EIP712Domain   `json:"domain"`

	decodedData *ZeroExTransactionData `json:"-"`
	hash        *common.Hash           `json:"-"`
}

type SignedTransaction struct {
	Transaction
	Signature []byte `json:"signature"`
}

// ComputeTransactionHash computes a 0x transaction hash
func (tx *Transaction) ComputeTransactionHash() (common.Hash, error) {
	if tx.hash != nil {
		return *tx.hash, nil
	}

	var message = map[string]interface{}{
		"salt":                  tx.Salt.String(),
		"expirationTimeSeconds": tx.ExpirationTimeSeconds.String(),
		"gasPrice":              tx.GasPrice.String(),
		"signerAddress":         tx.SignerAddress.Hex(),
		"data":                  tx.Data,
	}

	var typedData = typeddata.TypedData{
		Types:       eip712TransactionTypes,
		PrimaryType: "ZeroExTransaction",
		Domain:      tx.ExchangeDomain(),
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
	tx.hash = &hash

	return hash, nil
}

type Permyriad struct {
	Value *big.Int
}

type ZeroExTransactionData struct {
	FunctionName              FuturesFunctionName
	Orders                    []*Order
	MarketIDs                 []common.Hash
	SubAccountID              common.Hash
	ExchangeAddress           common.Address
	CloseQuantity             *big.Int
	SubaccountID              common.Hash
	Signatures                [][]byte
	PositionID                *big.Int
	Quantity                  *big.Int
	BaseCurrency              common.Address
	Amount                    *big.Int
	IsRevertingOnPartialFills bool
	// market specific parameters
	Ticker                 string
	Oracle                 common.Address
	InitialMarginRatio     Permyriad
	MaintenanceMarginRatio Permyriad
	FundingInterval        *big.Int
	ExpirationTime         *big.Int
	MakerTxFee             Permyriad
	TakerTxFee             Permyriad
	RelayerFeePercentage   Permyriad
	MarketID               common.Hash
	// parameters from 0x
	LeftOrders            []*Order
	RightOrders           []*Order
	TakerAssetFillAmounts []*big.Int
	MakerAssetFillAmount  *big.Int
	TakerAssetFillAmount  *big.Int
	LeftSignatures        [][]byte
	RightSignatures       [][]byte
}

func (tx *Transaction) DecodeTransactionData() (*ZeroExTransactionData, error) {
	if tx.decodedData != nil {
		return tx.decodedData, nil
	}

	txData, err := DecodeFromTransactionData(tx.Data)
	if err != nil {
		err := errors.Wrap(err, "failed to decode from transaction data")
		return nil, err
	}

	if txData.Orders != nil {
		for _, order := range txData.Orders {
			// these two fields make the order a complete version
			order.ChainID = tx.Domain.ChainID
			order.ExchangeAddress = tx.Domain.VerifyingContract
		}
	} else if txData.LeftOrders != nil {
		for _, order := range txData.LeftOrders {
			// these two fields make the order a complete version
			order.ChainID = tx.Domain.ChainID
			order.ExchangeAddress = tx.Domain.VerifyingContract
		}
	}

	tx.decodedData = txData

	return tx.decodedData, nil
}

func (data *ZeroExTransactionData) ValidateAssetFillAmounts() error {
	if data.isMarketBuyFn() {
		if len(data.TakerAssetFillAmounts) > 0 {
			err := errors.Errorf("tx is %s but TakerAssetFillAmounts provided", data.FunctionName)
			return err
		} else if data.MakerAssetFillAmount == nil {
			err := errors.Errorf("tx is %s but MakerAssetFillAmount not provided", data.FunctionName)
			return err
		}
		return nil
	} else if data.isMarketSellFn() {
		if len(data.TakerAssetFillAmounts) > 0 {
			err := errors.Errorf("tx is %s but TakerAssetFillAmounts provided", data.FunctionName)
			return err
		} else if data.TakerAssetFillAmount == nil {
			err := errors.Errorf("tx is %s but TakerAssetFillAmount not provided", data.FunctionName)
			return err
		}
		return nil
	} else if data.isMatchOrdersFn() {
		if (len(data.RightOrders) != len(data.LeftOrders)) || (len(data.RightOrders) != len(data.RightSignatures)) || (len(data.LeftSignatures) != len(data.RightSignatures)) {
			err := errors.Errorf("tx is %s but length of RightOrders/LeftOrders/RightSignatures/LeftSignatures do not match", data.FunctionName)
			return err
		}
	} else if len(data.TakerAssetFillAmounts) != len(data.Orders) {
		// TODO: add more validation cases
		// otherwise fill or something
		err := errors.New("incorrect TakerAssetFillAmounts length: must match Orders length")
		return err
	}
	return nil
}

func (data *ZeroExTransactionData) isMarketBuyFn() bool {
	return strings.HasPrefix(string(data.FunctionName), "marketBuy")
}

func (data *ZeroExTransactionData) isMarketSellFn() bool {
	return strings.HasPrefix(string(data.FunctionName), "marketSell")
}

func (data *ZeroExTransactionData) isMatchOrdersFn() bool {
	return strings.Contains(string(data.FunctionName), "atchOrders")
}

func (data *ZeroExTransactionData) isBuy() bool {
	return strings.Contains(string(data.FunctionName), "Buy")
}

func (data *ZeroExTransactionData) isSell() bool {
	return strings.Contains(string(data.FunctionName), "Sell")
}

// SignTransaction signs the 0x transaction with the supplied Signer
func SignTransaction(
	sender common.Address,
	ethSigner Signer,
	tx *Transaction,
) (*SignedTransaction, error) {
	if tx == nil {
		return nil, errors.New("cannot sign nil transaction")
	}

	txHash, err := tx.ComputeTransactionHash()
	if err != nil {
		return nil, err
	}

	ecSignature, err := ethSigner.EthSign(txHash.Bytes(), sender)
	if err != nil {
		return nil, err
	}

	// Generate 0x Ethereum Signature (append the signature type byte)
	signature := make([]byte, 66)
	signature[0] = ecSignature.V
	copy(signature[1:33], ecSignature.R[:])
	copy(signature[33:65], ecSignature.S[:])
	signature[65] = byte(EthSignSignature)
	signedTransaction := &SignedTransaction{
		Transaction: *tx,
		Signature:   signature,
	}

	return signedTransaction, nil
}

func (tx *Transaction) ExchangeDomain() typeddata.TypedDataDomain {
	return makeExchangeDomain(tx.Domain.ChainID, tx.Domain.VerifyingContract)
}

func makeExchangeDomain(
	chainID *big.Int,
	verifyingContract common.Address,
) typeddata.TypedDataDomain {
	return typeddata.TypedDataDomain{
		Name:              "0x Protocol",
		Version:           "3.0.0",
		ChainId:           math.NewHexOrDecimal256(chainID.Int64()),
		VerifyingContract: verifyingContract.Hex(),
	}
}
