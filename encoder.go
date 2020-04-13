package zeroex

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/InjectiveLabs/zeroex-go/wrappers"
)

func OrdersToTrimmed(orders []*SignedOrder) []wrappers.Order {
	trimmedOrders := make([]wrappers.Order, len(orders))
	for idx, o := range orders {
		trimmedOrders[idx] = o.Trim()
	}

	return trimmedOrders
}

func EncodeOrdersToExchangeData(fnName ExchangeFunctionName, signedOrders []*SignedOrder) (data []byte, err error) {
	orders := OrdersToTrimmed(signedOrders)

	switch {
	case SingleFillFnNames[fnName]:
		data, err = IExchangeABIPack(fnName, orders[0], orders[0].TakerAssetAmount, signedOrders[0].Signature)
	case BatchFillFnNames[fnName]:
		takerAssetAmounts := make([]*big.Int, len(orders))
		signatures := make([][]byte, len(orders))

		for idx, o := range orders {
			takerAssetAmounts[idx] = o.TakerAssetAmount
			signatures[idx] = signedOrders[idx].Signature
		}

		data, err = IExchangeABIPack(fnName, orders, takerAssetAmounts, signatures)
	case MarketFillFnNames[fnName]:
		totalFillAmount := new(big.Int)
		signatures := make([][]byte, len(orders))

		for idx, o := range orders {
			if fnName.HasPart("Buy") {
				totalFillAmount.Add(totalFillAmount, o.MakerAssetAmount)
			} else {
				totalFillAmount.Add(totalFillAmount, o.TakerAssetAmount)
			}
			signatures[idx] = signedOrders[idx].Signature
		}

		data, err = IExchangeABIPack(fnName, orders, totalFillAmount, signatures)
	case MatchOrderFnNames[fnName]:
		data, err = IExchangeABIPack(fnName, orders[0], orders[1], signedOrders[0].Signature, signedOrders[1].Signature)
	case fnName == CancelOrder:
		data, err = IExchangeABIPack(fnName, orders[0])
	case fnName == BatchCancelOrders:
		data, err = IExchangeABIPack(fnName, orders)
	case fnName == CancelOrdersUpTo:
		data, err = IExchangeABIPack(fnName, new(big.Int))
	case fnName == PreSign:
		orderHash, _ := signedOrders[0].ComputeOrderHash()
		data, err = IExchangeABIPack(fnName, orderHash.Bytes())
	case fnName == SetSignatureValidatorApproval:
		data, err = IExchangeABIPack(fnName, common.Address{}, true)
	default:
		err = errors.Errorf("IExchange function is not supported: %s", fnName)
		return nil, err
	}

	if err != nil {
		err = errors.Wrapf(err, "failed to pack %s", fnName)
	}

	return data, err
}

func DecodeFromTransactionData(data []byte) (txData *ZeroExTransactionData, err error) {
	if len(data) < 4 {
		err := errors.New("data must be at least 4 bytes long")
		return nil, err
	}

	id := data[:4]
	method, err := exchangeABI.MethodById(id)
	if err != nil {
		err = errors.Wrap(err, "failed to get method name")
		return nil, err
	}

	switch ExchangeFunctionName(method.Name) {
	case FillOrder:
		inputs := struct {
			Order                wrappers.Order
			TakerAssetFillAmount *big.Int
			Signature            []byte
		}{}

		if err = method.Inputs.Unpack(&inputs, data[4:]); err != nil {
			err = errors.Wrap(err, "failed to unpack method inputs")
			return nil, err
		}

		txData = &ZeroExTransactionData{
			FunctionName:          FillOrder,
			Orders:                make([]*Order, 1),
			TakerAssetFillAmounts: make([]*big.Int, 1),
			Signatures:            make([][]byte, 1),
		}
		txData.Orders[0] = FromTrimmedOrder(inputs.Order)
		txData.TakerAssetFillAmounts[0] = inputs.TakerAssetFillAmount
		txData.Signatures[0] = inputs.Signature
	case BatchFillOrders:
		inputs := struct {
			Orders                []wrappers.Order
			TakerAssetFillAmounts []*big.Int
			Signatures            [][]byte
		}{}

		if err = method.Inputs.Unpack(&inputs, data[4:]); err != nil {
			err = errors.Wrap(err, "failed to unpack method inputs")
			return nil, err
		}

		txData = &ZeroExTransactionData{
			FunctionName:          FillOrder,
			Orders:                make([]*Order, len(inputs.Orders)),
			TakerAssetFillAmounts: make([]*big.Int, len(inputs.TakerAssetFillAmounts)),
			Signatures:            make([][]byte, len(inputs.Signatures)),
		}
		for idx, order := range inputs.Orders {
			txData.Orders[idx] = FromTrimmedOrder(order)
		}
		for idx, takerAssetFillAmount := range inputs.TakerAssetFillAmounts {
			txData.TakerAssetFillAmounts[idx] = takerAssetFillAmount
		}
		for idx, signature := range inputs.Signatures {
			txData.Signatures[idx] = signature
		}
	default:
		panic("not supported: " + method.Name)
	}

	return txData, nil
}

var exchangeABI, _ = abi.JSON(strings.NewReader(wrappers.ExchangeABI))

func IExchangeABIPack(fnName ExchangeFunctionName, args ...interface{}) (data []byte, err error) {
	return exchangeABI.Pack(string(fnName), args...)
}

func IExchangeABIUnpack(fnName ExchangeFunctionName, data []byte, out interface{}) error {
	return exchangeABI.Unpack(out, string(fnName), data)
}

type ExchangeFunctionName string

func (e ExchangeFunctionName) HasPart(part string) bool {
	return strings.Contains(string(e), part)
}

const (
	BatchCancelOrders                 ExchangeFunctionName = "batchCancelOrders"
	BatchExecuteTransactions          ExchangeFunctionName = "batchExecuteTransactions"
	BatchFillOrKillOrders             ExchangeFunctionName = "batchFillOrKillOrders"
	BatchFillOrders                   ExchangeFunctionName = "batchFillOrders"
	BatchFillOrdersNoThrow            ExchangeFunctionName = "batchFillOrdersNoThrow"
	BatchMatchOrders                  ExchangeFunctionName = "batchMatchOrders"
	BatchMatchOrdersWithMaximalFill   ExchangeFunctionName = "batchMatchOrdersWithMaximalFill"
	CancelOrder                       ExchangeFunctionName = "cancelOrder"
	CancelOrdersUpTo                  ExchangeFunctionName = "cancelOrdersUpTo"
	ExecuteTransaction                ExchangeFunctionName = "executeTransaction"
	FillOrKillOrder                   ExchangeFunctionName = "fillOrKillOrder"
	FillOrder                         ExchangeFunctionName = "fillOrder"
	FillOrderNoThrow                  ExchangeFunctionName = "fillOrderNoThrow"
	MarketBuyOrdersNoThrow            ExchangeFunctionName = "marketBuyOrdersNoThrow"
	MarketSellOrdersNoThrow           ExchangeFunctionName = "marketSellOrdersNoThrow"
	MarketBuyOrdersFillOrKill         ExchangeFunctionName = "marketBuyOrdersFillOrKill"
	MarketSellOrdersFillOrKill        ExchangeFunctionName = "marketSellOrdersFillOrKill"
	MatchOrders                       ExchangeFunctionName = "matchOrders"
	MatchOrdersWithMaximalFill        ExchangeFunctionName = "matchOrdersWithMaximalFill"
	PreSign                           ExchangeFunctionName = "preSign"
	RegisterAssetProxy                ExchangeFunctionName = "registerAssetProxy"
	SetSignatureValidatorApproval     ExchangeFunctionName = "setSignatureValidatorApproval"
	SimulateDispatchTransferFromCalls ExchangeFunctionName = "simulateDispatchTransferFromCalls"
	TransferOwnership                 ExchangeFunctionName = "transferOwnership"
	SetProtocolFeeMultiplier          ExchangeFunctionName = "setProtocolFeeMultiplier"
	SetProtocolFeeCollectorAddress    ExchangeFunctionName = "setProtocolFeeCollectorAddress"
	DetachProtocolFeeCollector        ExchangeFunctionName = "detachProtocolFeeCollector"
)

var SingleFillFnNames = map[ExchangeFunctionName]bool{
	FillOrder:       true,
	FillOrKillOrder: true,
}

var BatchFillFnNames = map[ExchangeFunctionName]bool{
	BatchFillOrders:        true,
	BatchFillOrKillOrders:  true,
	BatchFillOrdersNoThrow: true,
}

var MarketFillFnNames = map[ExchangeFunctionName]bool{
	MarketBuyOrdersFillOrKill:  true,
	MarketSellOrdersFillOrKill: true,
	MarketBuyOrdersNoThrow:     true,
	MarketSellOrdersNoThrow:    true,
}

var MatchOrderFnNames = map[ExchangeFunctionName]bool{
	MatchOrders:                true,
	MatchOrdersWithMaximalFill: true,
}

var BatchMatchOrderFnNames = map[ExchangeFunctionName]bool{
	BatchMatchOrders:                true,
	BatchMatchOrdersWithMaximalFill: true,
}

var cancelOrderFnNames = map[ExchangeFunctionName]bool{
	CancelOrder:       true,
	BatchCancelOrders: true,
	CancelOrdersUpTo:  true,
}
