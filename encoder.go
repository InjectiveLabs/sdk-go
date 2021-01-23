package sdk

import (
	"fmt"
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
	fmt.Println("KAWABANGAAAAAAA")
	fmt.Printf("method %s has id %s", method, common.Bytes2Hex(id))
	if err != nil {
		err = errors.Wrap(err, "failed to get method name")
		return nil, err
	}
	data = data[4:]

	switch FuturesFunctionName(method.Name) {
	case ClosePositionMetaTransaction:
		inputs := struct {
			ExchangeAddress           common.Address
			IsRevertingOnPartialFills bool
			SubAccountID              common.Hash
			MarketID                  common.Hash
			CloseQuantity             *big.Int
		}{}

		if err = futuresABI.UnpackIntoInterface(&inputs, string(ClosePositionMetaTransaction), data); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}

		txData = &ZeroExTransactionData{
			FunctionName:              FuturesFunctionName(method.Name),
			ExchangeAddress:           inputs.ExchangeAddress,
			IsRevertingOnPartialFills: inputs.IsRevertingOnPartialFills,
			SubAccountID:              inputs.SubAccountID,
			MarketID:                  inputs.MarketID,
			CloseQuantity:             inputs.CloseQuantity,
		}
	case ClosePosition:
		inputs := struct {
			SubAccountID              common.Hash
			MarketID                  common.Hash
			Orders                    []wrappers.Order
			Quantity                  *big.Int
			IsRevertingOnPartialFills bool
			Signatures                [][]byte
		}{}

		if err = futuresABI.UnpackIntoInterface(&inputs, string(ClosePosition), data); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}

		txData = &ZeroExTransactionData{
			FunctionName:              FuturesFunctionName(method.Name),
			Orders:                    make([]*Order, len(inputs.Orders)),
			Signatures:                make([][]byte, len(inputs.Signatures)),
			SubAccountID:              inputs.SubAccountID,
			MarketID:                  inputs.MarketID,
			Quantity:                  inputs.Quantity,
			IsRevertingOnPartialFills: inputs.IsRevertingOnPartialFills,
		}

		for idx, order := range inputs.Orders {
			txData.Orders[idx] = FromTrimmedOrder(order)
		}

		for idx, signature := range inputs.Signatures {
			txData.Signatures[idx] = signature
		}
	case LiquidatePosition:
		inputs := struct {
			SubAccountID      common.Hash
			MarketID          common.Hash
			LiquidationCaller common.Address
			Orders            []wrappers.Order
			Quantity          *big.Int
			Signatures        [][]byte
		}{}

		if err = futuresABI.UnpackIntoInterface(&inputs, string(LiquidatePosition), data); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}

		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			Orders:       make([]*Order, len(inputs.Orders)),
			Signatures:   make([][]byte, len(inputs.Signatures)),
			SubAccountID: inputs.SubAccountID,
			MarketID:     inputs.MarketID,
			Quantity:     inputs.Quantity,
		}

		for idx, order := range inputs.Orders {
			txData.Orders[idx] = FromTrimmedOrder(order)
		}

		for idx, signature := range inputs.Signatures {
			txData.Signatures[idx] = signature
		}
	case VaporizePosition:
		inputs := struct {
			SubAccountID common.Hash
			MarketID     common.Hash
			Orders       []wrappers.Order
			Quantity     *big.Int
			Signatures   [][]byte
		}{}

		if err = futuresABI.UnpackIntoInterface(&inputs, string(VaporizePosition), data); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}

		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			Orders:       make([]*Order, 0),
			Signatures:   make([][]byte, 0),
			SubAccountID: inputs.SubAccountID,
			MarketID:     inputs.MarketID,
			Quantity:     inputs.Quantity,
		}
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
		if err = futuresABI.UnpackIntoInterface(&inputs, string(BatchCheckFunding), data); err != nil {
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
		if err = futuresABI.UnpackIntoInterface(&inputs, string(WithdrawForSubAccount), data); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}
		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			BaseCurrency: inputs.BaseCurrency,
			SubAccountID: inputs.SubAccountID,
			Amount:       inputs.Amount,
		}
	case Deposit:
		inputs := struct {
			BaseCurrency common.Address
			Amount       *big.Int
		}{}
		if err = futuresABI.UnpackIntoInterface(&inputs, string(Deposit), data); err != nil {
			err = errors.Wrapf(err, "failed to unpack %s method inputs", method.Name)
			return nil, err
		}
		txData = &ZeroExTransactionData{
			FunctionName: FuturesFunctionName(method.Name),
			BaseCurrency: inputs.BaseCurrency,
			Amount:       inputs.Amount,
		}
	case DepositForSubaccount:
		inputs := struct {
			BaseCurrency common.Address
			SubaccountID common.Hash
			Amount       *big.Int
		}{}
		if err = futuresABI.UnpackIntoInterface(&inputs, string(DepositForSubaccount), data); err != nil {
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
		if err = futuresABI.UnpackIntoInterface(&inputs, string(CreateMarketWithFixedMarketId), data); err != nil {
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

		if err = futuresABI.UnpackIntoInterface(&inputs, string(BatchCancelOrders), data); err != nil {
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

type FuturesFunctionName string

func (e FuturesFunctionName) HasPart(part string) bool {
	return strings.Contains(string(e), part)
}

const (
	ClosePositionMetaTransaction     FuturesFunctionName = "closePositionMetaTransaction"
	ClosePosition                    FuturesFunctionName = "closePosition"
	LiquidatePosition                FuturesFunctionName = "liquidatePosition"
	VaporizePosition                 FuturesFunctionName = "vaporizePosition"
	BatchCheckFunding                FuturesFunctionName = "batchCheckFunding"
	WithdrawForSubAccount            FuturesFunctionName = "withdrawForSubAccount"
	DepositForSubaccount             FuturesFunctionName = "depositForSubaccount"
	Deposit                          FuturesFunctionName = "deposit"
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
