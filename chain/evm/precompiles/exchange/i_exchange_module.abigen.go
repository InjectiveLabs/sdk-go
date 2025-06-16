// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchange

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CosmosCoin is an auto generated low-level Go binding around an user-defined struct.
type CosmosCoin struct {
	Amount *big.Int
	Denom  string
}

// IExchangeModuleAuthorization is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleAuthorization struct {
	Method     uint8
	SpendLimit []CosmosCoin
	Duration   *big.Int
}

// IExchangeModuleBatchCreateDerivativeLimitOrdersResponse is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleBatchCreateDerivativeLimitOrdersResponse struct {
	OrderHashes       []string
	CreatedOrdersCids []string
	FailedOrdersCids  []string
}

// IExchangeModuleBatchCreateSpotLimitOrdersResponse is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleBatchCreateSpotLimitOrdersResponse struct {
	OrderHashes       []string
	CreatedOrdersCids []string
	FailedOrdersCids  []string
}

// IExchangeModuleBatchUpdateOrdersRequest is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleBatchUpdateOrdersRequest struct {
	SubaccountID                   string
	SpotMarketIDsToCancelAll       []string
	SpotOrdersToCancel             []IExchangeModuleOrderData
	SpotOrdersToCreate             []IExchangeModuleSpotOrder
	DerivativeMarketIDsToCancelAll []string
	DerivativeOrdersToCancel       []IExchangeModuleOrderData
	DerivativeOrdersToCreate       []IExchangeModuleDerivativeOrder
}

// IExchangeModuleBatchUpdateOrdersResponse is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleBatchUpdateOrdersResponse struct {
	SpotCancelSuccess           []bool
	SpotOrderHashes             []string
	CreatedSpotOrdersCids       []string
	FailedSpotOrdersCids        []string
	DerivativeCancelSuccess     []bool
	DerivativeOrderHashes       []string
	CreatedDerivativeOrdersCids []string
	FailedDerivativeOrdersCids  []string
}

// IExchangeModuleCreateDerivativeLimitOrderResponse is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleCreateDerivativeLimitOrderResponse struct {
	OrderHash string
	Cid       string
}

// IExchangeModuleCreateDerivativeMarketOrderResponse is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleCreateDerivativeMarketOrderResponse struct {
	OrderHash              string
	Cid                    string
	Quantity               *big.Int
	Price                  *big.Int
	Fee                    *big.Int
	Payout                 *big.Int
	DeltaExecutionQuantity *big.Int
	DeltaExecutionMargin   *big.Int
	DeltaExecutionPrice    *big.Int
	DeltaIsLong            bool
}

// IExchangeModuleCreateSpotLimitOrderResponse is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleCreateSpotLimitOrderResponse struct {
	OrderHash string
	Cid       string
}

// IExchangeModuleCreateSpotMarketOrderResponse is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleCreateSpotMarketOrderResponse struct {
	OrderHash string
	Cid       string
	Quantity  *big.Int
	Price     *big.Int
	Fee       *big.Int
}

// IExchangeModuleDerivativeOrder is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleDerivativeOrder struct {
	MarketID     string
	SubaccountID string
	FeeRecipient string
	Price        *big.Int
	Quantity     *big.Int
	Cid          string
	OrderType    string
	Margin       *big.Int
	TriggerPrice *big.Int
}

// IExchangeModuleDerivativeOrdersRequest is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleDerivativeOrdersRequest struct {
	MarketID     string
	SubaccountID string
	OrderHashes  []string
}

// IExchangeModuleDerivativePosition is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleDerivativePosition struct {
	SubaccountID           string
	MarketID               string
	IsLong                 bool
	Quantity               *big.Int
	EntryPrice             *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
}

// IExchangeModuleOrderData is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleOrderData struct {
	MarketID     string
	SubaccountID string
	OrderHash    string
	OrderMask    int32
	Cid          string
}

// IExchangeModuleSpotOrder is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleSpotOrder struct {
	MarketID     string
	SubaccountID string
	FeeRecipient string
	Price        *big.Int
	Quantity     *big.Int
	Cid          string
	OrderType    string
	TriggerPrice *big.Int
}

// IExchangeModuleSpotOrdersRequest is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleSpotOrdersRequest struct {
	MarketID     string
	SubaccountID string
	OrderHashes  []string
}

// IExchangeModuleSubaccountDepositData is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleSubaccountDepositData struct {
	Denom            string
	AvailableBalance *big.Int
	TotalBalance     *big.Int
}

// IExchangeModuleTrimmedDerivativeLimitOrder is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleTrimmedDerivativeLimitOrder struct {
	Price     *big.Int
	Quantity  *big.Int
	Margin    *big.Int
	Fillable  *big.Int
	IsBuy     bool
	OrderHash string
	Cid       string
}

// IExchangeModuleTrimmedSpotLimitOrder is an auto generated low-level Go binding around an user-defined struct.
type IExchangeModuleTrimmedSpotLimitOrder struct {
	Price     *big.Int
	Quantity  *big.Int
	Fillable  *big.Int
	IsBuy     bool
	OrderHash string
	Cid       string
}

// ExchangeModuleMetaData contains all meta data concerning the ExchangeModule contract.
var ExchangeModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"internalType\":\"ExchangeTypes.MsgType\",\"name\":\"method\",\"type\":\"uint8\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"ExchangeTypes.MsgType\",\"name\":\"method\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"spendLimit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.Authorization[]\",\"name\":\"authorizations\",\"type\":\"tuple[]\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"int32\",\"name\":\"orderMask\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"internalType\":\"structIExchangeModule.OrderData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"batchCancelDerivativeOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"success\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"int32\",\"name\":\"orderMask\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"internalType\":\"structIExchangeModule.OrderData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"batchCancelSpotOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"success\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"feeRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.DerivativeOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"batchCreateDerivativeLimitOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"orderHashes\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"createdOrdersCids\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"failedOrdersCids\",\"type\":\"string[]\"}],\"internalType\":\"structIExchangeModule.BatchCreateDerivativeLimitOrdersResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"feeRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.SpotOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"batchCreateSpotLimitOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"orderHashes\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"createdOrdersCids\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"failedOrdersCids\",\"type\":\"string[]\"}],\"internalType\":\"structIExchangeModule.BatchCreateSpotLimitOrdersResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"spotMarketIDsToCancelAll\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"int32\",\"name\":\"orderMask\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"internalType\":\"structIExchangeModule.OrderData[]\",\"name\":\"spotOrdersToCancel\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"feeRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.SpotOrder[]\",\"name\":\"spotOrdersToCreate\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"derivativeMarketIDsToCancelAll\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"int32\",\"name\":\"orderMask\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"internalType\":\"structIExchangeModule.OrderData[]\",\"name\":\"derivativeOrdersToCancel\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"feeRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.DerivativeOrder[]\",\"name\":\"derivativeOrdersToCreate\",\"type\":\"tuple[]\"}],\"internalType\":\"structIExchangeModule.BatchUpdateOrdersRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"batchUpdateOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"bool[]\",\"name\":\"spotCancelSuccess\",\"type\":\"bool[]\"},{\"internalType\":\"string[]\",\"name\":\"spotOrderHashes\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"createdSpotOrdersCids\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"failedSpotOrdersCids\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"derivativeCancelSuccess\",\"type\":\"bool[]\"},{\"internalType\":\"string[]\",\"name\":\"derivativeOrderHashes\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"createdDerivativeOrdersCids\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"failedDerivativeOrdersCids\",\"type\":\"string[]\"}],\"internalType\":\"structIExchangeModule.BatchUpdateOrdersResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"int32\",\"name\":\"orderMask\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"cancelDerivativeOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"cancelSpotOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"feeRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.DerivativeOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"createDerivativeLimitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"internalType\":\"structIExchangeModule.CreateDerivativeLimitOrderResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"feeRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.DerivativeOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"createDerivativeMarketOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deltaExecutionQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deltaExecutionMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deltaExecutionPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"deltaIsLong\",\"type\":\"bool\"}],\"internalType\":\"structIExchangeModule.CreateDerivativeMarketOrderResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"feeRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.SpotOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"createSpotLimitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"internalType\":\"structIExchangeModule.CreateSpotLimitOrderResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"feeRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.SpotOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"createSpotMarketOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.CreateSpotMarketOrderResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceSubaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationSubaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreasePositionMargin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"orderHashes\",\"type\":\"string[]\"}],\"internalType\":\"structIExchangeModule.DerivativeOrdersRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"derivativeOrdersByHashes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillable\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"internalType\":\"structIExchangeModule.TrimmedDerivativeLimitOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceSubaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationSubaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"externalTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceSubaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationSubaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increasePositionMargin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"ExchangeTypes.MsgType[]\",\"name\":\"methods\",\"type\":\"uint8[]\"}],\"name\":\"revoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"orderHashes\",\"type\":\"string[]\"}],\"internalType\":\"structIExchangeModule.SpotOrdersRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"spotOrdersByHashes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillable\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"orderHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"internalType\":\"structIExchangeModule.TrimmedSpotLimitOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"subaccountDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"trader\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"subaccountNonce\",\"type\":\"uint32\"}],\"name\":\"subaccountDeposits\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"availableBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.SubaccountDepositData[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"}],\"name\":\"subaccountPositions\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"marketID\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeModule.DerivativePosition[]\",\"name\":\"positions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceSubaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationSubaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subaccountTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"subaccountID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExchangeModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeModuleMetaData.ABI instead.
var ExchangeModuleABI = ExchangeModuleMetaData.ABI

// ExchangeModule is an auto generated Go binding around an Ethereum contract.
type ExchangeModule struct {
	ExchangeModuleCaller     // Read-only binding to the contract
	ExchangeModuleTransactor // Write-only binding to the contract
	ExchangeModuleFilterer   // Log filterer for contract events
}

// ExchangeModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeModuleSession struct {
	Contract     *ExchangeModule   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeModuleCallerSession struct {
	Contract *ExchangeModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ExchangeModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeModuleTransactorSession struct {
	Contract     *ExchangeModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ExchangeModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeModuleRaw struct {
	Contract *ExchangeModule // Generic contract binding to access the raw methods on
}

// ExchangeModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeModuleCallerRaw struct {
	Contract *ExchangeModuleCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeModuleTransactorRaw struct {
	Contract *ExchangeModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeModule creates a new instance of ExchangeModule, bound to a specific deployed contract.
func NewExchangeModule(address common.Address, backend bind.ContractBackend) (*ExchangeModule, error) {
	contract, err := bindExchangeModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeModule{ExchangeModuleCaller: ExchangeModuleCaller{contract: contract}, ExchangeModuleTransactor: ExchangeModuleTransactor{contract: contract}, ExchangeModuleFilterer: ExchangeModuleFilterer{contract: contract}}, nil
}

// NewExchangeModuleCaller creates a new read-only instance of ExchangeModule, bound to a specific deployed contract.
func NewExchangeModuleCaller(address common.Address, caller bind.ContractCaller) (*ExchangeModuleCaller, error) {
	contract, err := bindExchangeModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeModuleCaller{contract: contract}, nil
}

// NewExchangeModuleTransactor creates a new write-only instance of ExchangeModule, bound to a specific deployed contract.
func NewExchangeModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeModuleTransactor, error) {
	contract, err := bindExchangeModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeModuleTransactor{contract: contract}, nil
}

// NewExchangeModuleFilterer creates a new log filterer instance of ExchangeModule, bound to a specific deployed contract.
func NewExchangeModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeModuleFilterer, error) {
	contract, err := bindExchangeModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeModuleFilterer{contract: contract}, nil
}

// bindExchangeModule binds a generic wrapper to an already deployed contract.
func bindExchangeModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExchangeModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeModule *ExchangeModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeModule.Contract.ExchangeModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeModule *ExchangeModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeModule.Contract.ExchangeModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeModule *ExchangeModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeModule.Contract.ExchangeModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeModule *ExchangeModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeModule *ExchangeModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeModule *ExchangeModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeModule.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x27be18a8.
//
// Solidity: function allowance(address grantee, address granter, uint8 method) view returns(bool allowed)
func (_ExchangeModule *ExchangeModuleCaller) Allowance(opts *bind.CallOpts, grantee common.Address, granter common.Address, method uint8) (bool, error) {
	var out []interface{}
	err := _ExchangeModule.contract.Call(opts, &out, "allowance", grantee, granter, method)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x27be18a8.
//
// Solidity: function allowance(address grantee, address granter, uint8 method) view returns(bool allowed)
func (_ExchangeModule *ExchangeModuleSession) Allowance(grantee common.Address, granter common.Address, method uint8) (bool, error) {
	return _ExchangeModule.Contract.Allowance(&_ExchangeModule.CallOpts, grantee, granter, method)
}

// Allowance is a free data retrieval call binding the contract method 0x27be18a8.
//
// Solidity: function allowance(address grantee, address granter, uint8 method) view returns(bool allowed)
func (_ExchangeModule *ExchangeModuleCallerSession) Allowance(grantee common.Address, granter common.Address, method uint8) (bool, error) {
	return _ExchangeModule.Contract.Allowance(&_ExchangeModule.CallOpts, grantee, granter, method)
}

// SubaccountDeposit is a free data retrieval call binding the contract method 0x9e96621f.
//
// Solidity: function subaccountDeposit(string subaccountID, string denom) view returns(uint256 availableBalance, uint256 totalBalance)
func (_ExchangeModule *ExchangeModuleCaller) SubaccountDeposit(opts *bind.CallOpts, subaccountID string, denom string) (struct {
	AvailableBalance *big.Int
	TotalBalance     *big.Int
}, error) {
	var out []interface{}
	err := _ExchangeModule.contract.Call(opts, &out, "subaccountDeposit", subaccountID, denom)

	outstruct := new(struct {
		AvailableBalance *big.Int
		TotalBalance     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AvailableBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SubaccountDeposit is a free data retrieval call binding the contract method 0x9e96621f.
//
// Solidity: function subaccountDeposit(string subaccountID, string denom) view returns(uint256 availableBalance, uint256 totalBalance)
func (_ExchangeModule *ExchangeModuleSession) SubaccountDeposit(subaccountID string, denom string) (struct {
	AvailableBalance *big.Int
	TotalBalance     *big.Int
}, error) {
	return _ExchangeModule.Contract.SubaccountDeposit(&_ExchangeModule.CallOpts, subaccountID, denom)
}

// SubaccountDeposit is a free data retrieval call binding the contract method 0x9e96621f.
//
// Solidity: function subaccountDeposit(string subaccountID, string denom) view returns(uint256 availableBalance, uint256 totalBalance)
func (_ExchangeModule *ExchangeModuleCallerSession) SubaccountDeposit(subaccountID string, denom string) (struct {
	AvailableBalance *big.Int
	TotalBalance     *big.Int
}, error) {
	return _ExchangeModule.Contract.SubaccountDeposit(&_ExchangeModule.CallOpts, subaccountID, denom)
}

// SubaccountDeposits is a free data retrieval call binding the contract method 0x12433f4b.
//
// Solidity: function subaccountDeposits(string subaccountID, string trader, uint32 subaccountNonce) view returns((string,uint256,uint256)[] deposits)
func (_ExchangeModule *ExchangeModuleCaller) SubaccountDeposits(opts *bind.CallOpts, subaccountID string, trader string, subaccountNonce uint32) ([]IExchangeModuleSubaccountDepositData, error) {
	var out []interface{}
	err := _ExchangeModule.contract.Call(opts, &out, "subaccountDeposits", subaccountID, trader, subaccountNonce)

	if err != nil {
		return *new([]IExchangeModuleSubaccountDepositData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IExchangeModuleSubaccountDepositData)).(*[]IExchangeModuleSubaccountDepositData)

	return out0, err

}

// SubaccountDeposits is a free data retrieval call binding the contract method 0x12433f4b.
//
// Solidity: function subaccountDeposits(string subaccountID, string trader, uint32 subaccountNonce) view returns((string,uint256,uint256)[] deposits)
func (_ExchangeModule *ExchangeModuleSession) SubaccountDeposits(subaccountID string, trader string, subaccountNonce uint32) ([]IExchangeModuleSubaccountDepositData, error) {
	return _ExchangeModule.Contract.SubaccountDeposits(&_ExchangeModule.CallOpts, subaccountID, trader, subaccountNonce)
}

// SubaccountDeposits is a free data retrieval call binding the contract method 0x12433f4b.
//
// Solidity: function subaccountDeposits(string subaccountID, string trader, uint32 subaccountNonce) view returns((string,uint256,uint256)[] deposits)
func (_ExchangeModule *ExchangeModuleCallerSession) SubaccountDeposits(subaccountID string, trader string, subaccountNonce uint32) ([]IExchangeModuleSubaccountDepositData, error) {
	return _ExchangeModule.Contract.SubaccountDeposits(&_ExchangeModule.CallOpts, subaccountID, trader, subaccountNonce)
}

// SubaccountPositions is a free data retrieval call binding the contract method 0x9bb15b3c.
//
// Solidity: function subaccountPositions(string subaccountID) view returns((string,string,bool,uint256,uint256,uint256,uint256)[] positions)
func (_ExchangeModule *ExchangeModuleCaller) SubaccountPositions(opts *bind.CallOpts, subaccountID string) ([]IExchangeModuleDerivativePosition, error) {
	var out []interface{}
	err := _ExchangeModule.contract.Call(opts, &out, "subaccountPositions", subaccountID)

	if err != nil {
		return *new([]IExchangeModuleDerivativePosition), err
	}

	out0 := *abi.ConvertType(out[0], new([]IExchangeModuleDerivativePosition)).(*[]IExchangeModuleDerivativePosition)

	return out0, err

}

// SubaccountPositions is a free data retrieval call binding the contract method 0x9bb15b3c.
//
// Solidity: function subaccountPositions(string subaccountID) view returns((string,string,bool,uint256,uint256,uint256,uint256)[] positions)
func (_ExchangeModule *ExchangeModuleSession) SubaccountPositions(subaccountID string) ([]IExchangeModuleDerivativePosition, error) {
	return _ExchangeModule.Contract.SubaccountPositions(&_ExchangeModule.CallOpts, subaccountID)
}

// SubaccountPositions is a free data retrieval call binding the contract method 0x9bb15b3c.
//
// Solidity: function subaccountPositions(string subaccountID) view returns((string,string,bool,uint256,uint256,uint256,uint256)[] positions)
func (_ExchangeModule *ExchangeModuleCallerSession) SubaccountPositions(subaccountID string) ([]IExchangeModuleDerivativePosition, error) {
	return _ExchangeModule.Contract.SubaccountPositions(&_ExchangeModule.CallOpts, subaccountID)
}

// Approve is a paid mutator transaction binding the contract method 0xd7ca8d06.
//
// Solidity: function approve(address grantee, (uint8,(uint256,string)[],uint256)[] authorizations) returns(bool approved)
func (_ExchangeModule *ExchangeModuleTransactor) Approve(opts *bind.TransactOpts, grantee common.Address, authorizations []IExchangeModuleAuthorization) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "approve", grantee, authorizations)
}

// Approve is a paid mutator transaction binding the contract method 0xd7ca8d06.
//
// Solidity: function approve(address grantee, (uint8,(uint256,string)[],uint256)[] authorizations) returns(bool approved)
func (_ExchangeModule *ExchangeModuleSession) Approve(grantee common.Address, authorizations []IExchangeModuleAuthorization) (*types.Transaction, error) {
	return _ExchangeModule.Contract.Approve(&_ExchangeModule.TransactOpts, grantee, authorizations)
}

// Approve is a paid mutator transaction binding the contract method 0xd7ca8d06.
//
// Solidity: function approve(address grantee, (uint8,(uint256,string)[],uint256)[] authorizations) returns(bool approved)
func (_ExchangeModule *ExchangeModuleTransactorSession) Approve(grantee common.Address, authorizations []IExchangeModuleAuthorization) (*types.Transaction, error) {
	return _ExchangeModule.Contract.Approve(&_ExchangeModule.TransactOpts, grantee, authorizations)
}

// BatchCancelDerivativeOrders is a paid mutator transaction binding the contract method 0x8b073525.
//
// Solidity: function batchCancelDerivativeOrders(address sender, (string,string,string,int32,string)[] data) returns(bool[] success)
func (_ExchangeModule *ExchangeModuleTransactor) BatchCancelDerivativeOrders(opts *bind.TransactOpts, sender common.Address, data []IExchangeModuleOrderData) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "batchCancelDerivativeOrders", sender, data)
}

// BatchCancelDerivativeOrders is a paid mutator transaction binding the contract method 0x8b073525.
//
// Solidity: function batchCancelDerivativeOrders(address sender, (string,string,string,int32,string)[] data) returns(bool[] success)
func (_ExchangeModule *ExchangeModuleSession) BatchCancelDerivativeOrders(sender common.Address, data []IExchangeModuleOrderData) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchCancelDerivativeOrders(&_ExchangeModule.TransactOpts, sender, data)
}

// BatchCancelDerivativeOrders is a paid mutator transaction binding the contract method 0x8b073525.
//
// Solidity: function batchCancelDerivativeOrders(address sender, (string,string,string,int32,string)[] data) returns(bool[] success)
func (_ExchangeModule *ExchangeModuleTransactorSession) BatchCancelDerivativeOrders(sender common.Address, data []IExchangeModuleOrderData) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchCancelDerivativeOrders(&_ExchangeModule.TransactOpts, sender, data)
}

// BatchCancelSpotOrders is a paid mutator transaction binding the contract method 0x438051ab.
//
// Solidity: function batchCancelSpotOrders(address sender, (string,string,string,int32,string)[] data) returns(bool[] success)
func (_ExchangeModule *ExchangeModuleTransactor) BatchCancelSpotOrders(opts *bind.TransactOpts, sender common.Address, data []IExchangeModuleOrderData) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "batchCancelSpotOrders", sender, data)
}

// BatchCancelSpotOrders is a paid mutator transaction binding the contract method 0x438051ab.
//
// Solidity: function batchCancelSpotOrders(address sender, (string,string,string,int32,string)[] data) returns(bool[] success)
func (_ExchangeModule *ExchangeModuleSession) BatchCancelSpotOrders(sender common.Address, data []IExchangeModuleOrderData) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchCancelSpotOrders(&_ExchangeModule.TransactOpts, sender, data)
}

// BatchCancelSpotOrders is a paid mutator transaction binding the contract method 0x438051ab.
//
// Solidity: function batchCancelSpotOrders(address sender, (string,string,string,int32,string)[] data) returns(bool[] success)
func (_ExchangeModule *ExchangeModuleTransactorSession) BatchCancelSpotOrders(sender common.Address, data []IExchangeModuleOrderData) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchCancelSpotOrders(&_ExchangeModule.TransactOpts, sender, data)
}

// BatchCreateDerivativeLimitOrders is a paid mutator transaction binding the contract method 0x79374eab.
//
// Solidity: function batchCreateDerivativeLimitOrders(address sender, (string,string,string,uint256,uint256,string,string,uint256,uint256)[] orders) returns((string[],string[],string[]) response)
func (_ExchangeModule *ExchangeModuleTransactor) BatchCreateDerivativeLimitOrders(opts *bind.TransactOpts, sender common.Address, orders []IExchangeModuleDerivativeOrder) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "batchCreateDerivativeLimitOrders", sender, orders)
}

// BatchCreateDerivativeLimitOrders is a paid mutator transaction binding the contract method 0x79374eab.
//
// Solidity: function batchCreateDerivativeLimitOrders(address sender, (string,string,string,uint256,uint256,string,string,uint256,uint256)[] orders) returns((string[],string[],string[]) response)
func (_ExchangeModule *ExchangeModuleSession) BatchCreateDerivativeLimitOrders(sender common.Address, orders []IExchangeModuleDerivativeOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchCreateDerivativeLimitOrders(&_ExchangeModule.TransactOpts, sender, orders)
}

// BatchCreateDerivativeLimitOrders is a paid mutator transaction binding the contract method 0x79374eab.
//
// Solidity: function batchCreateDerivativeLimitOrders(address sender, (string,string,string,uint256,uint256,string,string,uint256,uint256)[] orders) returns((string[],string[],string[]) response)
func (_ExchangeModule *ExchangeModuleTransactorSession) BatchCreateDerivativeLimitOrders(sender common.Address, orders []IExchangeModuleDerivativeOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchCreateDerivativeLimitOrders(&_ExchangeModule.TransactOpts, sender, orders)
}

// BatchCreateSpotLimitOrders is a paid mutator transaction binding the contract method 0x4881c7c6.
//
// Solidity: function batchCreateSpotLimitOrders(address sender, (string,string,string,uint256,uint256,string,string,uint256)[] orders) returns((string[],string[],string[]) response)
func (_ExchangeModule *ExchangeModuleTransactor) BatchCreateSpotLimitOrders(opts *bind.TransactOpts, sender common.Address, orders []IExchangeModuleSpotOrder) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "batchCreateSpotLimitOrders", sender, orders)
}

// BatchCreateSpotLimitOrders is a paid mutator transaction binding the contract method 0x4881c7c6.
//
// Solidity: function batchCreateSpotLimitOrders(address sender, (string,string,string,uint256,uint256,string,string,uint256)[] orders) returns((string[],string[],string[]) response)
func (_ExchangeModule *ExchangeModuleSession) BatchCreateSpotLimitOrders(sender common.Address, orders []IExchangeModuleSpotOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchCreateSpotLimitOrders(&_ExchangeModule.TransactOpts, sender, orders)
}

// BatchCreateSpotLimitOrders is a paid mutator transaction binding the contract method 0x4881c7c6.
//
// Solidity: function batchCreateSpotLimitOrders(address sender, (string,string,string,uint256,uint256,string,string,uint256)[] orders) returns((string[],string[],string[]) response)
func (_ExchangeModule *ExchangeModuleTransactorSession) BatchCreateSpotLimitOrders(sender common.Address, orders []IExchangeModuleSpotOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchCreateSpotLimitOrders(&_ExchangeModule.TransactOpts, sender, orders)
}

// BatchUpdateOrders is a paid mutator transaction binding the contract method 0xcb0b6590.
//
// Solidity: function batchUpdateOrders(address sender, (string,string[],(string,string,string,int32,string)[],(string,string,string,uint256,uint256,string,string,uint256)[],string[],(string,string,string,int32,string)[],(string,string,string,uint256,uint256,string,string,uint256,uint256)[]) request) returns((bool[],string[],string[],string[],bool[],string[],string[],string[]) response)
func (_ExchangeModule *ExchangeModuleTransactor) BatchUpdateOrders(opts *bind.TransactOpts, sender common.Address, request IExchangeModuleBatchUpdateOrdersRequest) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "batchUpdateOrders", sender, request)
}

// BatchUpdateOrders is a paid mutator transaction binding the contract method 0xcb0b6590.
//
// Solidity: function batchUpdateOrders(address sender, (string,string[],(string,string,string,int32,string)[],(string,string,string,uint256,uint256,string,string,uint256)[],string[],(string,string,string,int32,string)[],(string,string,string,uint256,uint256,string,string,uint256,uint256)[]) request) returns((bool[],string[],string[],string[],bool[],string[],string[],string[]) response)
func (_ExchangeModule *ExchangeModuleSession) BatchUpdateOrders(sender common.Address, request IExchangeModuleBatchUpdateOrdersRequest) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchUpdateOrders(&_ExchangeModule.TransactOpts, sender, request)
}

// BatchUpdateOrders is a paid mutator transaction binding the contract method 0xcb0b6590.
//
// Solidity: function batchUpdateOrders(address sender, (string,string[],(string,string,string,int32,string)[],(string,string,string,uint256,uint256,string,string,uint256)[],string[],(string,string,string,int32,string)[],(string,string,string,uint256,uint256,string,string,uint256,uint256)[]) request) returns((bool[],string[],string[],string[],bool[],string[],string[],string[]) response)
func (_ExchangeModule *ExchangeModuleTransactorSession) BatchUpdateOrders(sender common.Address, request IExchangeModuleBatchUpdateOrdersRequest) (*types.Transaction, error) {
	return _ExchangeModule.Contract.BatchUpdateOrders(&_ExchangeModule.TransactOpts, sender, request)
}

// CancelDerivativeOrder is a paid mutator transaction binding the contract method 0x44b9bf3a.
//
// Solidity: function cancelDerivativeOrder(address sender, string marketID, string subaccountID, string orderHash, int32 orderMask, string cid) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactor) CancelDerivativeOrder(opts *bind.TransactOpts, sender common.Address, marketID string, subaccountID string, orderHash string, orderMask int32, cid string) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "cancelDerivativeOrder", sender, marketID, subaccountID, orderHash, orderMask, cid)
}

// CancelDerivativeOrder is a paid mutator transaction binding the contract method 0x44b9bf3a.
//
// Solidity: function cancelDerivativeOrder(address sender, string marketID, string subaccountID, string orderHash, int32 orderMask, string cid) returns(bool success)
func (_ExchangeModule *ExchangeModuleSession) CancelDerivativeOrder(sender common.Address, marketID string, subaccountID string, orderHash string, orderMask int32, cid string) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CancelDerivativeOrder(&_ExchangeModule.TransactOpts, sender, marketID, subaccountID, orderHash, orderMask, cid)
}

// CancelDerivativeOrder is a paid mutator transaction binding the contract method 0x44b9bf3a.
//
// Solidity: function cancelDerivativeOrder(address sender, string marketID, string subaccountID, string orderHash, int32 orderMask, string cid) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactorSession) CancelDerivativeOrder(sender common.Address, marketID string, subaccountID string, orderHash string, orderMask int32, cid string) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CancelDerivativeOrder(&_ExchangeModule.TransactOpts, sender, marketID, subaccountID, orderHash, orderMask, cid)
}

// CancelSpotOrder is a paid mutator transaction binding the contract method 0x25bf6b92.
//
// Solidity: function cancelSpotOrder(address sender, string marketID, string subaccountID, string orderHash, string cid) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactor) CancelSpotOrder(opts *bind.TransactOpts, sender common.Address, marketID string, subaccountID string, orderHash string, cid string) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "cancelSpotOrder", sender, marketID, subaccountID, orderHash, cid)
}

// CancelSpotOrder is a paid mutator transaction binding the contract method 0x25bf6b92.
//
// Solidity: function cancelSpotOrder(address sender, string marketID, string subaccountID, string orderHash, string cid) returns(bool success)
func (_ExchangeModule *ExchangeModuleSession) CancelSpotOrder(sender common.Address, marketID string, subaccountID string, orderHash string, cid string) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CancelSpotOrder(&_ExchangeModule.TransactOpts, sender, marketID, subaccountID, orderHash, cid)
}

// CancelSpotOrder is a paid mutator transaction binding the contract method 0x25bf6b92.
//
// Solidity: function cancelSpotOrder(address sender, string marketID, string subaccountID, string orderHash, string cid) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactorSession) CancelSpotOrder(sender common.Address, marketID string, subaccountID string, orderHash string, cid string) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CancelSpotOrder(&_ExchangeModule.TransactOpts, sender, marketID, subaccountID, orderHash, cid)
}

// CreateDerivativeLimitOrder is a paid mutator transaction binding the contract method 0x20c69837.
//
// Solidity: function createDerivativeLimitOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256,uint256) order) returns((string,string) response)
func (_ExchangeModule *ExchangeModuleTransactor) CreateDerivativeLimitOrder(opts *bind.TransactOpts, sender common.Address, order IExchangeModuleDerivativeOrder) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "createDerivativeLimitOrder", sender, order)
}

// CreateDerivativeLimitOrder is a paid mutator transaction binding the contract method 0x20c69837.
//
// Solidity: function createDerivativeLimitOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256,uint256) order) returns((string,string) response)
func (_ExchangeModule *ExchangeModuleSession) CreateDerivativeLimitOrder(sender common.Address, order IExchangeModuleDerivativeOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CreateDerivativeLimitOrder(&_ExchangeModule.TransactOpts, sender, order)
}

// CreateDerivativeLimitOrder is a paid mutator transaction binding the contract method 0x20c69837.
//
// Solidity: function createDerivativeLimitOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256,uint256) order) returns((string,string) response)
func (_ExchangeModule *ExchangeModuleTransactorSession) CreateDerivativeLimitOrder(sender common.Address, order IExchangeModuleDerivativeOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CreateDerivativeLimitOrder(&_ExchangeModule.TransactOpts, sender, order)
}

// CreateDerivativeMarketOrder is a paid mutator transaction binding the contract method 0xb84857a1.
//
// Solidity: function createDerivativeMarketOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256,uint256) order) returns((string,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) response)
func (_ExchangeModule *ExchangeModuleTransactor) CreateDerivativeMarketOrder(opts *bind.TransactOpts, sender common.Address, order IExchangeModuleDerivativeOrder) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "createDerivativeMarketOrder", sender, order)
}

// CreateDerivativeMarketOrder is a paid mutator transaction binding the contract method 0xb84857a1.
//
// Solidity: function createDerivativeMarketOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256,uint256) order) returns((string,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) response)
func (_ExchangeModule *ExchangeModuleSession) CreateDerivativeMarketOrder(sender common.Address, order IExchangeModuleDerivativeOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CreateDerivativeMarketOrder(&_ExchangeModule.TransactOpts, sender, order)
}

// CreateDerivativeMarketOrder is a paid mutator transaction binding the contract method 0xb84857a1.
//
// Solidity: function createDerivativeMarketOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256,uint256) order) returns((string,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) response)
func (_ExchangeModule *ExchangeModuleTransactorSession) CreateDerivativeMarketOrder(sender common.Address, order IExchangeModuleDerivativeOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CreateDerivativeMarketOrder(&_ExchangeModule.TransactOpts, sender, order)
}

// CreateSpotLimitOrder is a paid mutator transaction binding the contract method 0xf642485e.
//
// Solidity: function createSpotLimitOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256) order) returns((string,string) response)
func (_ExchangeModule *ExchangeModuleTransactor) CreateSpotLimitOrder(opts *bind.TransactOpts, sender common.Address, order IExchangeModuleSpotOrder) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "createSpotLimitOrder", sender, order)
}

// CreateSpotLimitOrder is a paid mutator transaction binding the contract method 0xf642485e.
//
// Solidity: function createSpotLimitOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256) order) returns((string,string) response)
func (_ExchangeModule *ExchangeModuleSession) CreateSpotLimitOrder(sender common.Address, order IExchangeModuleSpotOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CreateSpotLimitOrder(&_ExchangeModule.TransactOpts, sender, order)
}

// CreateSpotLimitOrder is a paid mutator transaction binding the contract method 0xf642485e.
//
// Solidity: function createSpotLimitOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256) order) returns((string,string) response)
func (_ExchangeModule *ExchangeModuleTransactorSession) CreateSpotLimitOrder(sender common.Address, order IExchangeModuleSpotOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CreateSpotLimitOrder(&_ExchangeModule.TransactOpts, sender, order)
}

// CreateSpotMarketOrder is a paid mutator transaction binding the contract method 0x29d3d0e4.
//
// Solidity: function createSpotMarketOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256) order) returns((string,string,uint256,uint256,uint256) response)
func (_ExchangeModule *ExchangeModuleTransactor) CreateSpotMarketOrder(opts *bind.TransactOpts, sender common.Address, order IExchangeModuleSpotOrder) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "createSpotMarketOrder", sender, order)
}

// CreateSpotMarketOrder is a paid mutator transaction binding the contract method 0x29d3d0e4.
//
// Solidity: function createSpotMarketOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256) order) returns((string,string,uint256,uint256,uint256) response)
func (_ExchangeModule *ExchangeModuleSession) CreateSpotMarketOrder(sender common.Address, order IExchangeModuleSpotOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CreateSpotMarketOrder(&_ExchangeModule.TransactOpts, sender, order)
}

// CreateSpotMarketOrder is a paid mutator transaction binding the contract method 0x29d3d0e4.
//
// Solidity: function createSpotMarketOrder(address sender, (string,string,string,uint256,uint256,string,string,uint256) order) returns((string,string,uint256,uint256,uint256) response)
func (_ExchangeModule *ExchangeModuleTransactorSession) CreateSpotMarketOrder(sender common.Address, order IExchangeModuleSpotOrder) (*types.Transaction, error) {
	return _ExchangeModule.Contract.CreateSpotMarketOrder(&_ExchangeModule.TransactOpts, sender, order)
}

// DecreasePositionMargin is a paid mutator transaction binding the contract method 0xaf78360b.
//
// Solidity: function decreasePositionMargin(address sender, string sourceSubaccountID, string destinationSubaccountID, string marketID, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactor) DecreasePositionMargin(opts *bind.TransactOpts, sender common.Address, sourceSubaccountID string, destinationSubaccountID string, marketID string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "decreasePositionMargin", sender, sourceSubaccountID, destinationSubaccountID, marketID, amount)
}

// DecreasePositionMargin is a paid mutator transaction binding the contract method 0xaf78360b.
//
// Solidity: function decreasePositionMargin(address sender, string sourceSubaccountID, string destinationSubaccountID, string marketID, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleSession) DecreasePositionMargin(sender common.Address, sourceSubaccountID string, destinationSubaccountID string, marketID string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.DecreasePositionMargin(&_ExchangeModule.TransactOpts, sender, sourceSubaccountID, destinationSubaccountID, marketID, amount)
}

// DecreasePositionMargin is a paid mutator transaction binding the contract method 0xaf78360b.
//
// Solidity: function decreasePositionMargin(address sender, string sourceSubaccountID, string destinationSubaccountID, string marketID, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactorSession) DecreasePositionMargin(sender common.Address, sourceSubaccountID string, destinationSubaccountID string, marketID string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.DecreasePositionMargin(&_ExchangeModule.TransactOpts, sender, sourceSubaccountID, destinationSubaccountID, marketID, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe441dec9.
//
// Solidity: function deposit(address sender, string subaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactor) Deposit(opts *bind.TransactOpts, sender common.Address, subaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "deposit", sender, subaccountID, denom, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe441dec9.
//
// Solidity: function deposit(address sender, string subaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleSession) Deposit(sender common.Address, subaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.Deposit(&_ExchangeModule.TransactOpts, sender, subaccountID, denom, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe441dec9.
//
// Solidity: function deposit(address sender, string subaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactorSession) Deposit(sender common.Address, subaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.Deposit(&_ExchangeModule.TransactOpts, sender, subaccountID, denom, amount)
}

// DerivativeOrdersByHashes is a paid mutator transaction binding the contract method 0xd6673c03.
//
// Solidity: function derivativeOrdersByHashes((string,string,string[]) request) returns((uint256,uint256,uint256,uint256,bool,string,string)[] orders)
func (_ExchangeModule *ExchangeModuleTransactor) DerivativeOrdersByHashes(opts *bind.TransactOpts, request IExchangeModuleDerivativeOrdersRequest) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "derivativeOrdersByHashes", request)
}

// DerivativeOrdersByHashes is a paid mutator transaction binding the contract method 0xd6673c03.
//
// Solidity: function derivativeOrdersByHashes((string,string,string[]) request) returns((uint256,uint256,uint256,uint256,bool,string,string)[] orders)
func (_ExchangeModule *ExchangeModuleSession) DerivativeOrdersByHashes(request IExchangeModuleDerivativeOrdersRequest) (*types.Transaction, error) {
	return _ExchangeModule.Contract.DerivativeOrdersByHashes(&_ExchangeModule.TransactOpts, request)
}

// DerivativeOrdersByHashes is a paid mutator transaction binding the contract method 0xd6673c03.
//
// Solidity: function derivativeOrdersByHashes((string,string,string[]) request) returns((uint256,uint256,uint256,uint256,bool,string,string)[] orders)
func (_ExchangeModule *ExchangeModuleTransactorSession) DerivativeOrdersByHashes(request IExchangeModuleDerivativeOrdersRequest) (*types.Transaction, error) {
	return _ExchangeModule.Contract.DerivativeOrdersByHashes(&_ExchangeModule.TransactOpts, request)
}

// ExternalTransfer is a paid mutator transaction binding the contract method 0xc01307d2.
//
// Solidity: function externalTransfer(address sender, string sourceSubaccountID, string destinationSubaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactor) ExternalTransfer(opts *bind.TransactOpts, sender common.Address, sourceSubaccountID string, destinationSubaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "externalTransfer", sender, sourceSubaccountID, destinationSubaccountID, denom, amount)
}

// ExternalTransfer is a paid mutator transaction binding the contract method 0xc01307d2.
//
// Solidity: function externalTransfer(address sender, string sourceSubaccountID, string destinationSubaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleSession) ExternalTransfer(sender common.Address, sourceSubaccountID string, destinationSubaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.ExternalTransfer(&_ExchangeModule.TransactOpts, sender, sourceSubaccountID, destinationSubaccountID, denom, amount)
}

// ExternalTransfer is a paid mutator transaction binding the contract method 0xc01307d2.
//
// Solidity: function externalTransfer(address sender, string sourceSubaccountID, string destinationSubaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactorSession) ExternalTransfer(sender common.Address, sourceSubaccountID string, destinationSubaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.ExternalTransfer(&_ExchangeModule.TransactOpts, sender, sourceSubaccountID, destinationSubaccountID, denom, amount)
}

// IncreasePositionMargin is a paid mutator transaction binding the contract method 0x8ff96af4.
//
// Solidity: function increasePositionMargin(address sender, string sourceSubaccountID, string destinationSubaccountID, string marketID, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactor) IncreasePositionMargin(opts *bind.TransactOpts, sender common.Address, sourceSubaccountID string, destinationSubaccountID string, marketID string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "increasePositionMargin", sender, sourceSubaccountID, destinationSubaccountID, marketID, amount)
}

// IncreasePositionMargin is a paid mutator transaction binding the contract method 0x8ff96af4.
//
// Solidity: function increasePositionMargin(address sender, string sourceSubaccountID, string destinationSubaccountID, string marketID, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleSession) IncreasePositionMargin(sender common.Address, sourceSubaccountID string, destinationSubaccountID string, marketID string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.IncreasePositionMargin(&_ExchangeModule.TransactOpts, sender, sourceSubaccountID, destinationSubaccountID, marketID, amount)
}

// IncreasePositionMargin is a paid mutator transaction binding the contract method 0x8ff96af4.
//
// Solidity: function increasePositionMargin(address sender, string sourceSubaccountID, string destinationSubaccountID, string marketID, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactorSession) IncreasePositionMargin(sender common.Address, sourceSubaccountID string, destinationSubaccountID string, marketID string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.IncreasePositionMargin(&_ExchangeModule.TransactOpts, sender, sourceSubaccountID, destinationSubaccountID, marketID, amount)
}

// Revoke is a paid mutator transaction binding the contract method 0x3f79e2b1.
//
// Solidity: function revoke(address grantee, uint8[] methods) returns(bool revoked)
func (_ExchangeModule *ExchangeModuleTransactor) Revoke(opts *bind.TransactOpts, grantee common.Address, methods []uint8) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "revoke", grantee, methods)
}

// Revoke is a paid mutator transaction binding the contract method 0x3f79e2b1.
//
// Solidity: function revoke(address grantee, uint8[] methods) returns(bool revoked)
func (_ExchangeModule *ExchangeModuleSession) Revoke(grantee common.Address, methods []uint8) (*types.Transaction, error) {
	return _ExchangeModule.Contract.Revoke(&_ExchangeModule.TransactOpts, grantee, methods)
}

// Revoke is a paid mutator transaction binding the contract method 0x3f79e2b1.
//
// Solidity: function revoke(address grantee, uint8[] methods) returns(bool revoked)
func (_ExchangeModule *ExchangeModuleTransactorSession) Revoke(grantee common.Address, methods []uint8) (*types.Transaction, error) {
	return _ExchangeModule.Contract.Revoke(&_ExchangeModule.TransactOpts, grantee, methods)
}

// SpotOrdersByHashes is a paid mutator transaction binding the contract method 0x57d90abb.
//
// Solidity: function spotOrdersByHashes((string,string,string[]) request) returns((uint256,uint256,uint256,bool,string,string)[] orders)
func (_ExchangeModule *ExchangeModuleTransactor) SpotOrdersByHashes(opts *bind.TransactOpts, request IExchangeModuleSpotOrdersRequest) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "spotOrdersByHashes", request)
}

// SpotOrdersByHashes is a paid mutator transaction binding the contract method 0x57d90abb.
//
// Solidity: function spotOrdersByHashes((string,string,string[]) request) returns((uint256,uint256,uint256,bool,string,string)[] orders)
func (_ExchangeModule *ExchangeModuleSession) SpotOrdersByHashes(request IExchangeModuleSpotOrdersRequest) (*types.Transaction, error) {
	return _ExchangeModule.Contract.SpotOrdersByHashes(&_ExchangeModule.TransactOpts, request)
}

// SpotOrdersByHashes is a paid mutator transaction binding the contract method 0x57d90abb.
//
// Solidity: function spotOrdersByHashes((string,string,string[]) request) returns((uint256,uint256,uint256,bool,string,string)[] orders)
func (_ExchangeModule *ExchangeModuleTransactorSession) SpotOrdersByHashes(request IExchangeModuleSpotOrdersRequest) (*types.Transaction, error) {
	return _ExchangeModule.Contract.SpotOrdersByHashes(&_ExchangeModule.TransactOpts, request)
}

// SubaccountTransfer is a paid mutator transaction binding the contract method 0x42eba2ed.
//
// Solidity: function subaccountTransfer(address sender, string sourceSubaccountID, string destinationSubaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactor) SubaccountTransfer(opts *bind.TransactOpts, sender common.Address, sourceSubaccountID string, destinationSubaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "subaccountTransfer", sender, sourceSubaccountID, destinationSubaccountID, denom, amount)
}

// SubaccountTransfer is a paid mutator transaction binding the contract method 0x42eba2ed.
//
// Solidity: function subaccountTransfer(address sender, string sourceSubaccountID, string destinationSubaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleSession) SubaccountTransfer(sender common.Address, sourceSubaccountID string, destinationSubaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.SubaccountTransfer(&_ExchangeModule.TransactOpts, sender, sourceSubaccountID, destinationSubaccountID, denom, amount)
}

// SubaccountTransfer is a paid mutator transaction binding the contract method 0x42eba2ed.
//
// Solidity: function subaccountTransfer(address sender, string sourceSubaccountID, string destinationSubaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactorSession) SubaccountTransfer(sender common.Address, sourceSubaccountID string, destinationSubaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.SubaccountTransfer(&_ExchangeModule.TransactOpts, sender, sourceSubaccountID, destinationSubaccountID, denom, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xeb28c205.
//
// Solidity: function withdraw(address sender, string subaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactor) Withdraw(opts *bind.TransactOpts, sender common.Address, subaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.contract.Transact(opts, "withdraw", sender, subaccountID, denom, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xeb28c205.
//
// Solidity: function withdraw(address sender, string subaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleSession) Withdraw(sender common.Address, subaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.Withdraw(&_ExchangeModule.TransactOpts, sender, subaccountID, denom, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xeb28c205.
//
// Solidity: function withdraw(address sender, string subaccountID, string denom, uint256 amount) returns(bool success)
func (_ExchangeModule *ExchangeModuleTransactorSession) Withdraw(sender common.Address, subaccountID string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeModule.Contract.Withdraw(&_ExchangeModule.TransactOpts, sender, subaccountID, denom, amount)
}
