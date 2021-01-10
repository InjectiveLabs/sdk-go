package sdk

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/InjectiveLabs/sdk-go/wrappers"
)

func DecodeFromTransactionData(data []byte) (txData *ZeroExTransactionData, err error) {
	if len(data) < 4 {
		err := errors.New("data must be at least 4 bytes long")
		return nil, err
	}

	id := data[:4]
	method, err := futuresABI.MethodById(id)
	if err != nil {
		err = errors.Wrap(err, "failed to get method name")
		return nil, err
	}

	switch FuturesFunctionName(method.Name) {
	case ClosePosition:
		inputs := struct {
			PositionID                *big.Int
			Orders                    []wrappers.Order
			Quantity                  *big.Int
			IsRevertingOnPartialFills bool
			Signatures                [][]byte
		}{}

		if err = method.Inputs.Unpack(&inputs, data[4:]); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}

		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			Orders:       make([]*Order, len(inputs.Orders)),
			Signatures:   make([][]byte, len(inputs.Signatures)),
		}

		txData.PositionID = inputs.PositionID
		txData.Quantity = inputs.Quantity
		txData.IsRevertingOnPartialFills = inputs.IsRevertingOnPartialFills

		for idx, order := range inputs.Orders {
			txData.Orders[idx] = FromTrimmedOrder(order)
		}

		for idx, signature := range inputs.Signatures {
			txData.Signatures[idx] = signature
		}
	case LiquidatePosition, VaporizePosition:
		inputs := struct {
			PositionID *big.Int
			Orders     []wrappers.Order
			Quantity   *big.Int
			Signatures [][]byte
		}{}

		if err = method.Inputs.Unpack(&inputs, data[4:]); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}

		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			Orders:       make([]*Order, len(inputs.Orders)),
			Signatures:   make([][]byte, len(inputs.Signatures)),
		}

		txData.PositionID = inputs.PositionID
		txData.Quantity = inputs.Quantity

		for idx, order := range inputs.Orders {
			txData.Orders[idx] = FromTrimmedOrder(order)
		}

		for idx, signature := range inputs.Signatures {
			txData.Signatures[idx] = signature
		}
	case BatchCheckFunding:
		inputs := struct {
			MarketIDs []common.Hash
		}{}
		if err = method.Inputs.Unpack(&inputs, data[4:]); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}
		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			MarketIDs:    make([]common.Hash, len(inputs.MarketIDs)),
		}
		for idx, marketID := range inputs.MarketIDs {
			txData.MarketIDs[idx] = marketID
		}
	case WithdrawForSubAccount:
		inputs := struct {
			BaseCurrency common.Address
			SubAccountID common.Hash
			Amount       *big.Int
		}{}
		if err = method.Inputs.Unpack(&inputs, data[4:]); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}
		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			BaseCurrency: inputs.BaseCurrency,
			SubAccountID: inputs.SubAccountID,
			Amount:       inputs.Amount,
		}
	case DepositForSubaccount:
		inputs := struct {
			BaseCurrency common.Address
			SubaccountID common.Hash
			Amount       *big.Int
		}{}
		if err = method.Inputs.Unpack(&inputs, data[4:]); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}
		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			BaseCurrency: inputs.BaseCurrency,
			SubaccountID: inputs.SubaccountID,
			Amount:       inputs.Amount,
		}
	case CreateMarketWithFixedMarketId:
		inputs := struct {
			Ticker                 string
			BaseCurrency           common.Address
			Oracle                 common.Address
			InitialMarginRatio     Permyriad
			MaintenanceMarginRatio Permyriad
			FundingInterval        *big.Int
			ExpirationTime         *big.Int
			MakerTxFee             Permyriad
			TakerTxFee             Permyriad
			RelayerFeePercentage   Permyriad
			MarketID               common.Hash
		}{}
		if err = method.Inputs.Unpack(&inputs, data[4:]); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}
		txData = &ZeroExTransactionData{
			FunctionName:           FuturesFunctionName(method.Name),
			Ticker:                 inputs.Ticker,
			BaseCurrency:           inputs.BaseCurrency,
			Oracle:                 inputs.Oracle,
			InitialMarginRatio:     inputs.InitialMarginRatio,
			MaintenanceMarginRatio: inputs.MaintenanceMarginRatio,
			FundingInterval:        inputs.FundingInterval,
			ExpirationTime:         inputs.ExpirationTime,
			MakerTxFee:             inputs.MakerTxFee,
			TakerTxFee:             inputs.TakerTxFee,
			RelayerFeePercentage:   inputs.RelayerFeePercentage,
			MarketID:               inputs.MarketID,
		}
	case BatchCancelOrders:
		inputs := struct {
			Orders []wrappers.Order
		}{}

		if err = method.Inputs.Unpack(&inputs, data[4:]); err != nil {
			err = errors.Wrap(err, "failed to unpack method inputs")
			return nil, err
		}

		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			Orders:       make([]*Order, len(inputs.Orders)),
		}
		for idx, order := range inputs.Orders {
			txData.Orders[idx] = FromTrimmedOrder(order)
		}

	}

	return txData, nil
}

var futuresABI, _ = abi.JSON(strings.NewReader(wrappers.FuturesABI))

func IFuturesABIPack(fnName FuturesFunctionName, args ...interface{}) (data []byte, err error) {
	return futuresABI.Pack(string(fnName), args...)
}

func IFuturesABIUnPack(fnName FuturesFunctionName, data []byte, out interface{}) error {
	return futuresABI.Unpack(out, string(fnName), data)
}

type FuturesFunctionName string

func (e FuturesFunctionName) HasPart(part string) bool {
	return strings.Contains(string(e), part)
}

const (
	ClosePosition                    FuturesFunctionName = "closePosition"
	LiquidatePosition                FuturesFunctionName = "liquidatePosition"
	VaporizePosition                 FuturesFunctionName = "vaporizePosition"
	BatchCheckFunding                FuturesFunctionName = "batchCheckFunding"
	WithdrawForSubAccount            FuturesFunctionName = "withdrawForSubAccount"
	DepositForSubaccount             FuturesFunctionName = "depositForSubaccount"
	CreateMarketWithFixedMarketId    FuturesFunctionName = "createMarketWithFixedMarketId"
	BatchCancelOrders                FuturesFunctionName = "batchCancelOrders"
	AddMarginIntoPosition            FuturesFunctionName = "addMarginIntoPosition"
	SettleExpiryFuturesPosition      FuturesFunctionName = "settleExpiryFuturesPosition"
	BatchSettleExpiryFuturesPosition FuturesFunctionName = "batchSettleExpiryFuturesPosition"
	TransferPosition                 FuturesFunctionName = "transferPosition"
	ApproveForReceiving              FuturesFunctionName = "approveForReceiving"
	SetReceiptApprovalForMarket      FuturesFunctionName = "setReceiptApprovalForMarket"
	SetReceiptApprovalForAll         FuturesFunctionName = "setReceiptApprovalForAll"
	Approve                          FuturesFunctionName = "approve"
	SetApprovalForMarket             FuturesFunctionName = "setApprovalForMarket"
	SetApprovalForAll                FuturesFunctionName = "setApprovalForAll"
)

const (
	//BatchCancelOrders                 FuturesFunctionName = "batchCancelOrders"
	BatchExecuteTransactions          FuturesFunctionName = "batchExecuteTransactions"
	BatchFillOrKillOrders             FuturesFunctionName = "batchFillOrKillOrders"
	BatchFillOrders                   FuturesFunctionName = "batchFillOrders"
	BatchFillOrdersNoThrow            FuturesFunctionName = "batchFillOrdersNoThrow"
	BatchMatchOrders                  FuturesFunctionName = "batchMatchOrders"
	BatchMatchOrdersWithMaximalFill   FuturesFunctionName = "batchMatchOrdersWithMaximalFill"
	CancelOrder                       FuturesFunctionName = "cancelOrder"
	CancelOrdersUpTo                  FuturesFunctionName = "cancelOrdersUpTo"
	ExecuteTransaction                FuturesFunctionName = "executeTransaction"
	FillOrKillOrder                   FuturesFunctionName = "fillOrKillOrder"
	FillOrder                         FuturesFunctionName = "fillOrder"
	FillOrderNoThrow                  FuturesFunctionName = "fillOrderNoThrow"
	MarketBuyOrdersNoThrow            FuturesFunctionName = "marketBuyOrdersNoThrow"
	MarketSellOrdersNoThrow           FuturesFunctionName = "marketSellOrdersNoThrow"
	MarketBuyOrdersFillOrKill         FuturesFunctionName = "marketBuyOrdersFillOrKill"
	MarketSellOrdersFillOrKill        FuturesFunctionName = "marketSellOrdersFillOrKill"
	MatchOrders                       FuturesFunctionName = "matchOrders"
	MatchOrdersWithMaximalFill        FuturesFunctionName = "matchOrdersWithMaximalFill"
	PreSign                           FuturesFunctionName = "preSign"
	RegisterAssetProxy                FuturesFunctionName = "registerAssetProxy"
	SetSignatureValidatorApproval     FuturesFunctionName = "setSignatureValidatorApproval"
	SimulateDispatchTransferFromCalls FuturesFunctionName = "simulateDispatchTransferFromCalls"
	TransferOwnership                 FuturesFunctionName = "transferOwnership"
	SetProtocolFeeMultiplier          FuturesFunctionName = "setProtocolFeeMultiplier"
	SetProtocolFeeCollectorAddress    FuturesFunctionName = "setProtocolFeeCollectorAddress"
	DetachProtocolFeeCollector        FuturesFunctionName = "detachProtocolFeeCollector"
)

var SingleFillFnNames = map[FuturesFunctionName]bool{
	FillOrder:       true,
	FillOrKillOrder: true,
}

var BatchFillFnNames = map[FuturesFunctionName]bool{
	BatchFillOrders:        true,
	BatchFillOrKillOrders:  true,
	BatchFillOrdersNoThrow: true,
}

var MarketFillFnNames = map[FuturesFunctionName]bool{
	MarketBuyOrdersFillOrKill:  true,
	MarketSellOrdersFillOrKill: true,
	MarketBuyOrdersNoThrow:     true,
	MarketSellOrdersNoThrow:    true,
}

var MatchOrderFnNames = map[FuturesFunctionName]bool{
	MatchOrders:                true,
	MatchOrdersWithMaximalFill: true,
}

var BatchMatchOrderFnNames = map[FuturesFunctionName]bool{
	BatchMatchOrders:                true,
	BatchMatchOrdersWithMaximalFill: true,
}

var cancelOrderFnNames = map[FuturesFunctionName]bool{
	CancelOrder:       true,
	BatchCancelOrders: true,
	CancelOrdersUpTo:  true,
}

func OrdersToTrimmed(orders []*SignedOrder) []wrappers.Order {
	trimmedOrders := make([]wrappers.Order, len(orders))
	for idx, o := range orders {
		trimmedOrders[idx] = o.Trim()
	}

	return trimmedOrders
}

//func EncodeOrdersToExchangeData(fnName FuturesFunctionName, signedOrders []*SignedOrder) (data []byte, err error) {
//	orders := OrdersToTrimmed(signedOrders)
//
//	switch {
//	case SingleFillFnNames[fnName]:
//		data, err = IExchangeABIPack(fnName, orders[0], orders[0].TakerAssetAmount, signedOrders[0].Signature)
//	case BatchFillFnNames[fnName]:
//		takerAssetAmounts := make([]*big.Int, len(orders))
//		signatures := make([][]byte, len(orders))
//
//		for idx, o := range orders {
//			takerAssetAmounts[idx] = o.TakerAssetAmount
//			signatures[idx] = signedOrders[idx].Signature
//		}
//
//		data, err = IExchangeABIPack(fnName, orders, takerAssetAmounts, signatures)
//	case MarketFillFnNames[fnName]:
//		totalFillAmount := new(big.Int)
//		signatures := make([][]byte, len(orders))
//
//		for idx, o := range orders {
//			if fnName.HasPart("Buy") {
//				totalFillAmount.Add(totalFillAmount, o.MakerAssetAmount)
//			} else {
//				totalFillAmount.Add(totalFillAmount, o.TakerAssetAmount)
//			}
//			signatures[idx] = signedOrders[idx].Signature
//		}
//
//		data, err = IExchangeABIPack(fnName, orders, totalFillAmount, signatures)
//	case MatchOrderFnNames[fnName]:
//		data, err = IExchangeABIPack(fnName, orders[0], orders[1], signedOrders[0].Signature, signedOrders[1].Signature)
//	case fnName == CancelOrder:
//		data, err = IExchangeABIPack(fnName, orders[0])
//	case fnName == BatchCancelOrders:
//		data, err = IExchangeABIPack(fnName, orders)
//	case fnName == CancelOrdersUpTo:
//		data, err = IExchangeABIPack(fnName, new(big.Int))
//	case fnName == PreSign:
//		orderHash, _ := signedOrders[0].ComputeOrderHash()
//		data, err = IExchangeABIPack(fnName, orderHash.Bytes())
//	case fnName == SetSignatureValidatorApproval:
//		data, err = IExchangeABIPack(fnName, common.Address{}, true)
//	default:
//		err = errors.Errorf("IExchange function is not supported: %s", fnName)
//		return nil, err
//	}
//
//	if err != nil {
//		err = errors.Wrapf(err, "failed to pack %s", fnName)
//	}
//
//	return data, err
//}