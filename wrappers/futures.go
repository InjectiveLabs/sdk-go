// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DerivativeOrderInfo is an auto generated low-level Go binding around an user-defined struct.
type DerivativeOrderInfo struct {
	OrderType                   uint8
	OrderStatus                 uint8
	OrderHash                   [32]byte
	OrderTakerAssetFilledAmount *big.Int
	SubAccountID                [32]byte
	Direction                   uint8
	MarketID                    [32]byte
	ContractPrice               *big.Int
}

// MixinOrdersCloseResults is an auto generated low-level Go binding around an user-defined struct.
type MixinOrdersCloseResults struct {
	Payout         *big.Int
	QuantityClosed *big.Int
}

// MixinOrdersFillResults is an auto generated low-level Go binding around an user-defined struct.
type MixinOrdersFillResults struct {
	MakerPositionID *big.Int
	TakerPositionID *big.Int
	MakerMarginUsed *big.Int
	TakerMarginUsed *big.Int
	QuantityFilled  *big.Int
	MakerFeePaid    *big.Int
	TakerFeePaid    *big.Int
}

// MixinOrdersMatchResults is an auto generated low-level Go binding around an user-defined struct.
type MixinOrdersMatchResults struct {
	LeftPositionID  *big.Int
	RightPositionID *big.Int
	LeftMarginUsed  *big.Int
	RightMarginUsed *big.Int
	QuantityFilled  *big.Int
	LeftFeePaid     *big.Int
	RightFeePaid    *big.Int
}

// MixinOrdersPositionResults is an auto generated low-level Go binding around an user-defined struct.
type MixinOrdersPositionResults struct {
	PositionID *big.Int
	MarginUsed *big.Int
	Quantity   *big.Int
	Fee        *big.Int
}

// PermyriadMathPermyriad is an auto generated low-level Go binding around an user-defined struct.
type PermyriadMathPermyriad struct {
	Value *big.Int
}

// TypesPosition is an auto generated low-level Go binding around an user-defined struct.
type TypesPosition struct {
	SubAccountID           [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	ContractPrice          *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
	OrderHash              [32]byte
}

// TypesTransactionFees is an auto generated low-level Go binding around an user-defined struct.
type TypesTransactionFees struct {
	Maker   PermyriadMathPermyriad
	Taker   PermyriadMathPermyriad
	Relayer PermyriadMathPermyriad
}

// FuturesABI is the input ABI used to generate the binding from.
const FuturesABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"_minimumMarginRatio\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subAccountNonce\",\"type\":\"uint256\"}],\"name\":\"AccountCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecrementSubaccountDeposits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"}],\"name\":\"FuturesCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"contractPNL\",\"type\":\"int256\"}],\"name\":\"FuturesClose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"contractPNL\",\"type\":\"int256\"}],\"name\":\"FuturesLiquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"leftOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rightOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"FuturesMatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFilled\",\"type\":\"uint256\"}],\"name\":\"FuturesOrderFill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"name\":\"FuturesPosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncrementSubaccountDeposits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addedMargin\",\"type\":\"uint256\"}],\"name\":\"MarginAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maintenanceMarginRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialMarginRatioFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerTxFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerTxFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayerFeePercentage\",\"type\":\"uint256\"}],\"name\":\"MarketCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingFee\",\"type\":\"int256\"}],\"name\":\"SetFunding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"SignatureValidatorApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_EXCHANGE_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_MARGIN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"subAccountNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedMargin\",\"type\":\"uint256\"}],\"name\":\"addMarginIntoPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressToSubAccountIDs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"margins\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"subAccountIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchFillOrKillOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.FillResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"margins\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchFillOrKillOrdersSinglePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.FillResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"margins\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"subAccountIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchFillOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.FillResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"margins\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchFillOrdersSinglePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.FillResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"leftOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"rightOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"leftSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"rightSignatures\",\"type\":\"bytes[]\"}],\"name\":\"batchMatchOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"leftPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.MatchResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"name\":\"calcCumulativeFunding\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"calcLiquidationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"calcMinMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"closePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.PositionResults[]\",\"name\":\"pResults\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"payout\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"quantityClosed\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.CloseResults\",\"name\":\"cResults\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"closePositionOrKill\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.PositionResults[]\",\"name\":\"pResults\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"payout\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"quantityClosed\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.CloseResults\",\"name\":\"cResults\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subAccountNonce\",\"type\":\"uint256\"}],\"name\":\"computeSubAccountIdFromNonce\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createDefaultSubAccountAndDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"initialMarginRatioFactor\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maintenanceMarginRatio\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"makerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"takerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"relayerFeePercentage\",\"type\":\"tuple\"}],\"name\":\"createMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"initialMarginRatioFactor\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maintenanceMarginRatio\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"makerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"takerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"relayerFeePercentage\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"createMarketWithFixedMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subAccountNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createSubAccountAndDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subAccountNonce\",\"type\":\"uint256\"}],\"name\":\"createSubAccountForTraderWithNonce\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositIntoSubAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyStopFutures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.FillResults\",\"name\":\"results\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.FillResults\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getDefaultSubAccountDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getDefaultSubAccountIdForTrader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"}],\"name\":\"getOrderRelevantState\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.OrderStatus\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"}],\"internalType\":\"structLibOrder.DerivativeOrderInfo\",\"name\":\"orderInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fillableTakerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidSignature\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"getOrderRelevantStates\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.OrderStatus\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"}],\"internalType\":\"structLibOrder.DerivativeOrderInfo[]\",\"name\":\"ordersInfo\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fillableTakerAssetAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"isValidSignature\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getPositionIDsForTrader\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"positionIDs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getPositionsForTrader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Position[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getTraderSubAccountsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"}],\"name\":\"getUnitPositionValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"unitPositionValue\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"insurancePools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidBaseCurrency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidOrderSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"liquidatePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.PositionResults[]\",\"name\":\"pResults\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"payout\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"quantityClosed\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.CloseResults\",\"name\":\"cResults\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"marketOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.FillResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"marketOrdersOrKill\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.FillResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketSerialToID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"initialMarginRatioFactor\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maintenanceMarginRatio\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextFundingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFunding\",\"type\":\"int256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maker\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"taker\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"relayer\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransactionFees\",\"name\":\"transactionFees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"leftOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"leftSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"leftPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.MatchResults\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"mostRecentEpochQuantity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"mostRecentEpochVWAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"mostRecentEpochVolume\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"leftOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"leftSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"multiMatchOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"leftPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.MatchResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderHashToPositionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payIntoInsurancePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"preSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"restrictedDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeFutures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"setFundingRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"subAccountDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subAccountIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subAccountToMarketToPositionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"vaporizePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.PositionResults[]\",\"name\":\"pResults\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"payout\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"quantityClosed\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.CloseResults\",\"name\":\"cResults\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromDefaultSubAcount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromSubAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Futures is an auto generated Go binding around an Ethereum contract.
type Futures struct {
	FuturesCaller     // Read-only binding to the contract
	FuturesTransactor // Write-only binding to the contract
	FuturesFilterer   // Log filterer for contract events
}

// FuturesCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuturesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuturesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuturesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuturesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuturesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuturesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuturesSession struct {
	Contract     *Futures          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuturesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuturesCallerSession struct {
	Contract *FuturesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FuturesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuturesTransactorSession struct {
	Contract     *FuturesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FuturesRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuturesRaw struct {
	Contract *Futures // Generic contract binding to access the raw methods on
}

// FuturesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuturesCallerRaw struct {
	Contract *FuturesCaller // Generic read-only contract binding to access the raw methods on
}

// FuturesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuturesTransactorRaw struct {
	Contract *FuturesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFutures creates a new instance of Futures, bound to a specific deployed contract.
func NewFutures(address common.Address, backend bind.ContractBackend) (*Futures, error) {
	contract, err := bindFutures(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Futures{FuturesCaller: FuturesCaller{contract: contract}, FuturesTransactor: FuturesTransactor{contract: contract}, FuturesFilterer: FuturesFilterer{contract: contract}}, nil
}

// NewFuturesCaller creates a new read-only instance of Futures, bound to a specific deployed contract.
func NewFuturesCaller(address common.Address, caller bind.ContractCaller) (*FuturesCaller, error) {
	contract, err := bindFutures(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuturesCaller{contract: contract}, nil
}

// NewFuturesTransactor creates a new write-only instance of Futures, bound to a specific deployed contract.
func NewFuturesTransactor(address common.Address, transactor bind.ContractTransactor) (*FuturesTransactor, error) {
	contract, err := bindFutures(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuturesTransactor{contract: contract}, nil
}

// NewFuturesFilterer creates a new log filterer instance of Futures, bound to a specific deployed contract.
func NewFuturesFilterer(address common.Address, filterer bind.ContractFilterer) (*FuturesFilterer, error) {
	contract, err := bindFutures(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuturesFilterer{contract: contract}, nil
}

// bindFutures binds a generic wrapper to an already deployed contract.
func bindFutures(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuturesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (f *FuturesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return f.Contract.FuturesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (f *FuturesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.Contract.FuturesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (f *FuturesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return f.Contract.FuturesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (f *FuturesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return f.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (f *FuturesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (f *FuturesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return f.Contract.contract.Transact(opts, method, params...)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (f *FuturesCaller) EIP712EXCHANGEDOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := f.contract.Call(opts, out, "EIP712_EXCHANGE_DOMAIN_HASH")
	return *ret0, err
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (f *FuturesSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return f.Contract.EIP712EXCHANGEDOMAINHASH(&f.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (f *FuturesCallerSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return f.Contract.EIP712EXCHANGEDOMAINHASH(&f.CallOpts)
}

// MINIMUMMARGINRATIO is a free data retrieval call binding the contract method 0xe63f9a7d.
//
// Solidity: function MINIMUM_MARGIN_RATIO() view returns(uint256 value)
func (f *FuturesCaller) MINIMUMMARGINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "MINIMUM_MARGIN_RATIO")
	return *ret0, err
}

// MINIMUMMARGINRATIO is a free data retrieval call binding the contract method 0xe63f9a7d.
//
// Solidity: function MINIMUM_MARGIN_RATIO() view returns(uint256 value)
func (f *FuturesSession) MINIMUMMARGINRATIO() (*big.Int, error) {
	return f.Contract.MINIMUMMARGINRATIO(&f.CallOpts)
}

// MINIMUMMARGINRATIO is a free data retrieval call binding the contract method 0xe63f9a7d.
//
// Solidity: function MINIMUM_MARGIN_RATIO() view returns(uint256 value)
func (f *FuturesCallerSession) MINIMUMMARGINRATIO() (*big.Int, error) {
	return f.Contract.MINIMUMMARGINRATIO(&f.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bytes32 subAccountID, uint256 subAccountNonce)
func (f *FuturesCaller) Accounts(opts *bind.CallOpts, arg0 [32]byte) (struct {
	SubAccountID    [32]byte
	SubAccountNonce *big.Int
}, error) {
	ret := new(struct {
		SubAccountID    [32]byte
		SubAccountNonce *big.Int
	})
	out := ret
	err := f.contract.Call(opts, out, "accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bytes32 subAccountID, uint256 subAccountNonce)
func (f *FuturesSession) Accounts(arg0 [32]byte) (struct {
	SubAccountID    [32]byte
	SubAccountNonce *big.Int
}, error) {
	return f.Contract.Accounts(&f.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bytes32 subAccountID, uint256 subAccountNonce)
func (f *FuturesCallerSession) Accounts(arg0 [32]byte) (struct {
	SubAccountID    [32]byte
	SubAccountNonce *big.Int
}, error) {
	return f.Contract.Accounts(&f.CallOpts, arg0)
}

// AddressToSubAccountIDs is a free data retrieval call binding the contract method 0x07294a8e.
//
// Solidity: function addressToSubAccountIDs(address , uint256 ) view returns(bytes32)
func (f *FuturesCaller) AddressToSubAccountIDs(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := f.contract.Call(opts, out, "addressToSubAccountIDs", arg0, arg1)
	return *ret0, err
}

// AddressToSubAccountIDs is a free data retrieval call binding the contract method 0x07294a8e.
//
// Solidity: function addressToSubAccountIDs(address , uint256 ) view returns(bytes32)
func (f *FuturesSession) AddressToSubAccountIDs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return f.Contract.AddressToSubAccountIDs(&f.CallOpts, arg0, arg1)
}

// AddressToSubAccountIDs is a free data retrieval call binding the contract method 0x07294a8e.
//
// Solidity: function addressToSubAccountIDs(address , uint256 ) view returns(bytes32)
func (f *FuturesCallerSession) AddressToSubAccountIDs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return f.Contract.AddressToSubAccountIDs(&f.CallOpts, arg0, arg1)
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) view returns(bool)
func (f *FuturesCaller) AllowedValidators(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := f.contract.Call(opts, out, "allowedValidators", arg0, arg1)
	return *ret0, err
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) view returns(bool)
func (f *FuturesSession) AllowedValidators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return f.Contract.AllowedValidators(&f.CallOpts, arg0, arg1)
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) view returns(bool)
func (f *FuturesCallerSession) AllowedValidators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return f.Contract.AllowedValidators(&f.CallOpts, arg0, arg1)
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) view returns(int256)
func (f *FuturesCaller) CalcCumulativeFunding(opts *bind.CallOpts, marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "calcCumulativeFunding", marketID, cumulativeFundingEntry)
	return *ret0, err
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) view returns(int256)
func (f *FuturesSession) CalcCumulativeFunding(marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	return f.Contract.CalcCumulativeFunding(&f.CallOpts, marketID, cumulativeFundingEntry)
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) view returns(int256)
func (f *FuturesCallerSession) CalcCumulativeFunding(marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	return f.Contract.CalcCumulativeFunding(&f.CallOpts, marketID, cumulativeFundingEntry)
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 marketID, uint256 quantity) view returns(uint256)
func (f *FuturesCaller) CalcLiquidationFee(opts *bind.CallOpts, marketID [32]byte, quantity *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "calcLiquidationFee", marketID, quantity)
	return *ret0, err
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 marketID, uint256 quantity) view returns(uint256)
func (f *FuturesSession) CalcLiquidationFee(marketID [32]byte, quantity *big.Int) (*big.Int, error) {
	return f.Contract.CalcLiquidationFee(&f.CallOpts, marketID, quantity)
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 marketID, uint256 quantity) view returns(uint256)
func (f *FuturesCallerSession) CalcLiquidationFee(marketID [32]byte, quantity *big.Int) (*big.Int, error) {
	return f.Contract.CalcLiquidationFee(&f.CallOpts, marketID, quantity)
}

// CalcMinMargin is a free data retrieval call binding the contract method 0x35c43c4e.
//
// Solidity: function calcMinMargin(bytes32 marketID, uint256 quantity, uint256 price) view returns(uint256)
func (f *FuturesCaller) CalcMinMargin(opts *bind.CallOpts, marketID [32]byte, quantity *big.Int, price *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "calcMinMargin", marketID, quantity, price)
	return *ret0, err
}

// CalcMinMargin is a free data retrieval call binding the contract method 0x35c43c4e.
//
// Solidity: function calcMinMargin(bytes32 marketID, uint256 quantity, uint256 price) view returns(uint256)
func (f *FuturesSession) CalcMinMargin(marketID [32]byte, quantity *big.Int, price *big.Int) (*big.Int, error) {
	return f.Contract.CalcMinMargin(&f.CallOpts, marketID, quantity, price)
}

// CalcMinMargin is a free data retrieval call binding the contract method 0x35c43c4e.
//
// Solidity: function calcMinMargin(bytes32 marketID, uint256 quantity, uint256 price) view returns(uint256)
func (f *FuturesCallerSession) CalcMinMargin(marketID [32]byte, quantity *big.Int, price *big.Int) (*big.Int, error) {
	return f.Contract.CalcMinMargin(&f.CallOpts, marketID, quantity, price)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (f *FuturesCaller) Cancelled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := f.contract.Call(opts, out, "cancelled", arg0)
	return *ret0, err
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (f *FuturesSession) Cancelled(arg0 [32]byte) (bool, error) {
	return f.Contract.Cancelled(&f.CallOpts, arg0)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (f *FuturesCallerSession) Cancelled(arg0 [32]byte) (bool, error) {
	return f.Contract.Cancelled(&f.CallOpts, arg0)
}

// ComputeSubAccountIdFromNonce is a free data retrieval call binding the contract method 0x1103b304.
//
// Solidity: function computeSubAccountIdFromNonce(address trader, uint256 subAccountNonce) pure returns(bytes32)
func (f *FuturesCaller) ComputeSubAccountIdFromNonce(opts *bind.CallOpts, trader common.Address, subAccountNonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := f.contract.Call(opts, out, "computeSubAccountIdFromNonce", trader, subAccountNonce)
	return *ret0, err
}

// ComputeSubAccountIdFromNonce is a free data retrieval call binding the contract method 0x1103b304.
//
// Solidity: function computeSubAccountIdFromNonce(address trader, uint256 subAccountNonce) pure returns(bytes32)
func (f *FuturesSession) ComputeSubAccountIdFromNonce(trader common.Address, subAccountNonce *big.Int) ([32]byte, error) {
	return f.Contract.ComputeSubAccountIdFromNonce(&f.CallOpts, trader, subAccountNonce)
}

// ComputeSubAccountIdFromNonce is a free data retrieval call binding the contract method 0x1103b304.
//
// Solidity: function computeSubAccountIdFromNonce(address trader, uint256 subAccountNonce) pure returns(bytes32)
func (f *FuturesCallerSession) ComputeSubAccountIdFromNonce(trader common.Address, subAccountNonce *big.Int) ([32]byte, error) {
	return f.Contract.ComputeSubAccountIdFromNonce(&f.CallOpts, trader, subAccountNonce)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (f *FuturesCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "filled", arg0)
	return *ret0, err
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (f *FuturesSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.Filled(&f.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.Filled(&f.CallOpts, arg0)
}

// GetDefaultSubAccountDeposits is a free data retrieval call binding the contract method 0x1883e458.
//
// Solidity: function getDefaultSubAccountDeposits(address baseCurrency, address trader) view returns(uint256)
func (f *FuturesCaller) GetDefaultSubAccountDeposits(opts *bind.CallOpts, baseCurrency common.Address, trader common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "getDefaultSubAccountDeposits", baseCurrency, trader)
	return *ret0, err
}

// GetDefaultSubAccountDeposits is a free data retrieval call binding the contract method 0x1883e458.
//
// Solidity: function getDefaultSubAccountDeposits(address baseCurrency, address trader) view returns(uint256)
func (f *FuturesSession) GetDefaultSubAccountDeposits(baseCurrency common.Address, trader common.Address) (*big.Int, error) {
	return f.Contract.GetDefaultSubAccountDeposits(&f.CallOpts, baseCurrency, trader)
}

// GetDefaultSubAccountDeposits is a free data retrieval call binding the contract method 0x1883e458.
//
// Solidity: function getDefaultSubAccountDeposits(address baseCurrency, address trader) view returns(uint256)
func (f *FuturesCallerSession) GetDefaultSubAccountDeposits(baseCurrency common.Address, trader common.Address) (*big.Int, error) {
	return f.Contract.GetDefaultSubAccountDeposits(&f.CallOpts, baseCurrency, trader)
}

// GetDefaultSubAccountIdForTrader is a free data retrieval call binding the contract method 0x80755948.
//
// Solidity: function getDefaultSubAccountIdForTrader(address trader) pure returns(bytes32)
func (f *FuturesCaller) GetDefaultSubAccountIdForTrader(opts *bind.CallOpts, trader common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := f.contract.Call(opts, out, "getDefaultSubAccountIdForTrader", trader)
	return *ret0, err
}

// GetDefaultSubAccountIdForTrader is a free data retrieval call binding the contract method 0x80755948.
//
// Solidity: function getDefaultSubAccountIdForTrader(address trader) pure returns(bytes32)
func (f *FuturesSession) GetDefaultSubAccountIdForTrader(trader common.Address) ([32]byte, error) {
	return f.Contract.GetDefaultSubAccountIdForTrader(&f.CallOpts, trader)
}

// GetDefaultSubAccountIdForTrader is a free data retrieval call binding the contract method 0x80755948.
//
// Solidity: function getDefaultSubAccountIdForTrader(address trader) pure returns(bytes32)
func (f *FuturesCallerSession) GetDefaultSubAccountIdForTrader(trader common.Address) ([32]byte, error) {
	return f.Contract.GetDefaultSubAccountIdForTrader(&f.CallOpts, trader)
}

// GetOrderRelevantState is a free data retrieval call binding the contract method 0xef3a29b3.
//
// Solidity: function getOrderRelevantState((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature, uint256 indexPrice) view returns((uint8,uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256) orderInfo, uint256 fillableTakerAssetAmount, bool isValidSignature)
func (f *FuturesCaller) GetOrderRelevantState(opts *bind.CallOpts, order Order, signature []byte, indexPrice *big.Int) (struct {
	OrderInfo                DerivativeOrderInfo
	FillableTakerAssetAmount *big.Int
	IsValidSignature         bool
}, error) {
	ret := new(struct {
		OrderInfo                DerivativeOrderInfo
		FillableTakerAssetAmount *big.Int
		IsValidSignature         bool
	})
	out := ret
	err := f.contract.Call(opts, out, "getOrderRelevantState", order, signature, indexPrice)
	return *ret, err
}

// GetOrderRelevantState is a free data retrieval call binding the contract method 0xef3a29b3.
//
// Solidity: function getOrderRelevantState((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature, uint256 indexPrice) view returns((uint8,uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256) orderInfo, uint256 fillableTakerAssetAmount, bool isValidSignature)
func (f *FuturesSession) GetOrderRelevantState(order Order, signature []byte, indexPrice *big.Int) (struct {
	OrderInfo                DerivativeOrderInfo
	FillableTakerAssetAmount *big.Int
	IsValidSignature         bool
}, error) {
	return f.Contract.GetOrderRelevantState(&f.CallOpts, order, signature, indexPrice)
}

// GetOrderRelevantState is a free data retrieval call binding the contract method 0xef3a29b3.
//
// Solidity: function getOrderRelevantState((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature, uint256 indexPrice) view returns((uint8,uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256) orderInfo, uint256 fillableTakerAssetAmount, bool isValidSignature)
func (f *FuturesCallerSession) GetOrderRelevantState(order Order, signature []byte, indexPrice *big.Int) (struct {
	OrderInfo                DerivativeOrderInfo
	FillableTakerAssetAmount *big.Int
	IsValidSignature         bool
}, error) {
	return f.Contract.GetOrderRelevantState(&f.CallOpts, order, signature, indexPrice)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, bytes[] signatures) view returns((uint8,uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256)[] ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (f *FuturesCaller) GetOrderRelevantStates(opts *bind.CallOpts, orders []Order, signatures [][]byte) (struct {
	OrdersInfo                []DerivativeOrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	ret := new(struct {
		OrdersInfo                []DerivativeOrderInfo
		FillableTakerAssetAmounts []*big.Int
		IsValidSignature          []bool
	})
	out := ret
	err := f.contract.Call(opts, out, "getOrderRelevantStates", orders, signatures)
	return *ret, err
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, bytes[] signatures) view returns((uint8,uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256)[] ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (f *FuturesSession) GetOrderRelevantStates(orders []Order, signatures [][]byte) (struct {
	OrdersInfo                []DerivativeOrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	return f.Contract.GetOrderRelevantStates(&f.CallOpts, orders, signatures)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, bytes[] signatures) view returns((uint8,uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256)[] ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (f *FuturesCallerSession) GetOrderRelevantStates(orders []Order, signatures [][]byte) (struct {
	OrdersInfo                []DerivativeOrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	return f.Contract.GetOrderRelevantStates(&f.CallOpts, orders, signatures)
}

// GetPositionIDsForTrader is a free data retrieval call binding the contract method 0xc3e49bb7.
//
// Solidity: function getPositionIDsForTrader(address trader, bytes32 marketID) view returns(uint256[] positionIDs)
func (f *FuturesCaller) GetPositionIDsForTrader(opts *bind.CallOpts, trader common.Address, marketID [32]byte) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "getPositionIDsForTrader", trader, marketID)
	return *ret0, err
}

// GetPositionIDsForTrader is a free data retrieval call binding the contract method 0xc3e49bb7.
//
// Solidity: function getPositionIDsForTrader(address trader, bytes32 marketID) view returns(uint256[] positionIDs)
func (f *FuturesSession) GetPositionIDsForTrader(trader common.Address, marketID [32]byte) ([]*big.Int, error) {
	return f.Contract.GetPositionIDsForTrader(&f.CallOpts, trader, marketID)
}

// GetPositionIDsForTrader is a free data retrieval call binding the contract method 0xc3e49bb7.
//
// Solidity: function getPositionIDsForTrader(address trader, bytes32 marketID) view returns(uint256[] positionIDs)
func (f *FuturesCallerSession) GetPositionIDsForTrader(trader common.Address, marketID [32]byte) ([]*big.Int, error) {
	return f.Contract.GetPositionIDsForTrader(&f.CallOpts, trader, marketID)
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,uint256,int256,bytes32)[])
func (f *FuturesCaller) GetPositionsForTrader(opts *bind.CallOpts, trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	var (
		ret0 = new([]TypesPosition)
	)
	out := ret0
	err := f.contract.Call(opts, out, "getPositionsForTrader", trader, marketID)
	return *ret0, err
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,uint256,int256,bytes32)[])
func (f *FuturesSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return f.Contract.GetPositionsForTrader(&f.CallOpts, trader, marketID)
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,uint256,int256,bytes32)[])
func (f *FuturesCallerSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return f.Contract.GetPositionsForTrader(&f.CallOpts, trader, marketID)
}

// GetTraderSubAccountsCount is a free data retrieval call binding the contract method 0x603ca5dc.
//
// Solidity: function getTraderSubAccountsCount(address trader) view returns(uint256)
func (f *FuturesCaller) GetTraderSubAccountsCount(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "getTraderSubAccountsCount", trader)
	return *ret0, err
}

// GetTraderSubAccountsCount is a free data retrieval call binding the contract method 0x603ca5dc.
//
// Solidity: function getTraderSubAccountsCount(address trader) view returns(uint256)
func (f *FuturesSession) GetTraderSubAccountsCount(trader common.Address) (*big.Int, error) {
	return f.Contract.GetTraderSubAccountsCount(&f.CallOpts, trader)
}

// GetTraderSubAccountsCount is a free data retrieval call binding the contract method 0x603ca5dc.
//
// Solidity: function getTraderSubAccountsCount(address trader) view returns(uint256)
func (f *FuturesCallerSession) GetTraderSubAccountsCount(trader common.Address) (*big.Int, error) {
	return f.Contract.GetTraderSubAccountsCount(&f.CallOpts, trader)
}

// GetUnitPositionValue is a free data retrieval call binding the contract method 0x574e2080.
//
// Solidity: function getUnitPositionValue(uint256 positionID, uint256 indexPrice) view returns(int256 unitPositionValue)
func (f *FuturesCaller) GetUnitPositionValue(opts *bind.CallOpts, positionID *big.Int, indexPrice *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "getUnitPositionValue", positionID, indexPrice)
	return *ret0, err
}

// GetUnitPositionValue is a free data retrieval call binding the contract method 0x574e2080.
//
// Solidity: function getUnitPositionValue(uint256 positionID, uint256 indexPrice) view returns(int256 unitPositionValue)
func (f *FuturesSession) GetUnitPositionValue(positionID *big.Int, indexPrice *big.Int) (*big.Int, error) {
	return f.Contract.GetUnitPositionValue(&f.CallOpts, positionID, indexPrice)
}

// GetUnitPositionValue is a free data retrieval call binding the contract method 0x574e2080.
//
// Solidity: function getUnitPositionValue(uint256 positionID, uint256 indexPrice) view returns(int256 unitPositionValue)
func (f *FuturesCallerSession) GetUnitPositionValue(positionID *big.Int, indexPrice *big.Int) (*big.Int, error) {
	return f.Contract.GetUnitPositionValue(&f.CallOpts, positionID, indexPrice)
}

// InsurancePools is a free data retrieval call binding the contract method 0x2514c1f1.
//
// Solidity: function insurancePools(bytes32 ) view returns(uint256)
func (f *FuturesCaller) InsurancePools(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "insurancePools", arg0)
	return *ret0, err
}

// InsurancePools is a free data retrieval call binding the contract method 0x2514c1f1.
//
// Solidity: function insurancePools(bytes32 ) view returns(uint256)
func (f *FuturesSession) InsurancePools(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.InsurancePools(&f.CallOpts, arg0)
}

// InsurancePools is a free data retrieval call binding the contract method 0x2514c1f1.
//
// Solidity: function insurancePools(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) InsurancePools(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.InsurancePools(&f.CallOpts, arg0)
}

// IsValidBaseCurrency is a free data retrieval call binding the contract method 0x227cdb85.
//
// Solidity: function isValidBaseCurrency(address ) view returns(bool)
func (f *FuturesCaller) IsValidBaseCurrency(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := f.contract.Call(opts, out, "isValidBaseCurrency", arg0)
	return *ret0, err
}

// IsValidBaseCurrency is a free data retrieval call binding the contract method 0x227cdb85.
//
// Solidity: function isValidBaseCurrency(address ) view returns(bool)
func (f *FuturesSession) IsValidBaseCurrency(arg0 common.Address) (bool, error) {
	return f.Contract.IsValidBaseCurrency(&f.CallOpts, arg0)
}

// IsValidBaseCurrency is a free data retrieval call binding the contract method 0x227cdb85.
//
// Solidity: function isValidBaseCurrency(address ) view returns(bool)
func (f *FuturesCallerSession) IsValidBaseCurrency(arg0 common.Address) (bool, error) {
	return f.Contract.IsValidBaseCurrency(&f.CallOpts, arg0)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature) view returns(bool isValid)
func (f *FuturesCaller) IsValidOrderSignature(opts *bind.CallOpts, order Order, signature []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := f.contract.Call(opts, out, "isValidOrderSignature", order, signature)
	return *ret0, err
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature) view returns(bool isValid)
func (f *FuturesSession) IsValidOrderSignature(order Order, signature []byte) (bool, error) {
	return f.Contract.IsValidOrderSignature(&f.CallOpts, order, signature)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature) view returns(bool isValid)
func (f *FuturesCallerSession) IsValidOrderSignature(order Order, signature []byte) (bool, error) {
	return f.Contract.IsValidOrderSignature(&f.CallOpts, order, signature)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (f *FuturesCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "marketCount")
	return *ret0, err
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (f *FuturesSession) MarketCount() (*big.Int, error) {
	return f.Contract.MarketCount(&f.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (f *FuturesCallerSession) MarketCount() (*big.Int, error) {
	return f.Contract.MarketCount(&f.CallOpts)
}

// MarketSerialToID is a free data retrieval call binding the contract method 0xbae18473.
//
// Solidity: function marketSerialToID(uint256 ) view returns(bytes32)
func (f *FuturesCaller) MarketSerialToID(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := f.contract.Call(opts, out, "marketSerialToID", arg0)
	return *ret0, err
}

// MarketSerialToID is a free data retrieval call binding the contract method 0xbae18473.
//
// Solidity: function marketSerialToID(uint256 ) view returns(bytes32)
func (f *FuturesSession) MarketSerialToID(arg0 *big.Int) ([32]byte, error) {
	return f.Contract.MarketSerialToID(&f.CallOpts, arg0)
}

// MarketSerialToID is a free data retrieval call binding the contract method 0xbae18473.
//
// Solidity: function marketSerialToID(uint256 ) view returns(bytes32)
func (f *FuturesCallerSession) MarketSerialToID(arg0 *big.Int) ([32]byte, error) {
	return f.Contract.MarketSerialToID(&f.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) view returns(bytes32 marketID, address baseCurrency, string ticker, address oracle, (uint256) initialMarginRatioFactor, (uint256) maintenanceMarginRatio, uint256 indexPrice, uint256 nextFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding, ((uint256),(uint256),(uint256)) transactionFees)
func (f *FuturesCaller) Markets(opts *bind.CallOpts, arg0 [32]byte) (struct {
	MarketID                 [32]byte
	BaseCurrency             common.Address
	Ticker                   string
	Oracle                   common.Address
	InitialMarginRatioFactor PermyriadMathPermyriad
	MaintenanceMarginRatio   PermyriadMathPermyriad
	IndexPrice               *big.Int
	NextFundingTimestamp     *big.Int
	FundingInterval          *big.Int
	CumulativeFunding        *big.Int
	TransactionFees          TypesTransactionFees
}, error) {
	ret := new(struct {
		MarketID                 [32]byte
		BaseCurrency             common.Address
		Ticker                   string
		Oracle                   common.Address
		InitialMarginRatioFactor PermyriadMathPermyriad
		MaintenanceMarginRatio   PermyriadMathPermyriad
		IndexPrice               *big.Int
		NextFundingTimestamp     *big.Int
		FundingInterval          *big.Int
		CumulativeFunding        *big.Int
		TransactionFees          TypesTransactionFees
	})
	out := ret
	err := f.contract.Call(opts, out, "markets", arg0)
	return *ret, err
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) view returns(bytes32 marketID, address baseCurrency, string ticker, address oracle, (uint256) initialMarginRatioFactor, (uint256) maintenanceMarginRatio, uint256 indexPrice, uint256 nextFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding, ((uint256),(uint256),(uint256)) transactionFees)
func (f *FuturesSession) Markets(arg0 [32]byte) (struct {
	MarketID                 [32]byte
	BaseCurrency             common.Address
	Ticker                   string
	Oracle                   common.Address
	InitialMarginRatioFactor PermyriadMathPermyriad
	MaintenanceMarginRatio   PermyriadMathPermyriad
	IndexPrice               *big.Int
	NextFundingTimestamp     *big.Int
	FundingInterval          *big.Int
	CumulativeFunding        *big.Int
	TransactionFees          TypesTransactionFees
}, error) {
	return f.Contract.Markets(&f.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) view returns(bytes32 marketID, address baseCurrency, string ticker, address oracle, (uint256) initialMarginRatioFactor, (uint256) maintenanceMarginRatio, uint256 indexPrice, uint256 nextFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding, ((uint256),(uint256),(uint256)) transactionFees)
func (f *FuturesCallerSession) Markets(arg0 [32]byte) (struct {
	MarketID                 [32]byte
	BaseCurrency             common.Address
	Ticker                   string
	Oracle                   common.Address
	InitialMarginRatioFactor PermyriadMathPermyriad
	MaintenanceMarginRatio   PermyriadMathPermyriad
	IndexPrice               *big.Int
	NextFundingTimestamp     *big.Int
	FundingInterval          *big.Int
	CumulativeFunding        *big.Int
	TransactionFees          TypesTransactionFees
}, error) {
	return f.Contract.Markets(&f.CallOpts, arg0)
}

// MostRecentEpochQuantity is a free data retrieval call binding the contract method 0xf7a28a1a.
//
// Solidity: function mostRecentEpochQuantity(bytes32 ) view returns(uint256)
func (f *FuturesCaller) MostRecentEpochQuantity(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "mostRecentEpochQuantity", arg0)
	return *ret0, err
}

// MostRecentEpochQuantity is a free data retrieval call binding the contract method 0xf7a28a1a.
//
// Solidity: function mostRecentEpochQuantity(bytes32 ) view returns(uint256)
func (f *FuturesSession) MostRecentEpochQuantity(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentEpochQuantity(&f.CallOpts, arg0)
}

// MostRecentEpochQuantity is a free data retrieval call binding the contract method 0xf7a28a1a.
//
// Solidity: function mostRecentEpochQuantity(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) MostRecentEpochQuantity(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentEpochQuantity(&f.CallOpts, arg0)
}

// MostRecentEpochVWAP is a free data retrieval call binding the contract method 0xcf03660b.
//
// Solidity: function mostRecentEpochVWAP(bytes32 ) view returns(uint256)
func (f *FuturesCaller) MostRecentEpochVWAP(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "mostRecentEpochVWAP", arg0)
	return *ret0, err
}

// MostRecentEpochVWAP is a free data retrieval call binding the contract method 0xcf03660b.
//
// Solidity: function mostRecentEpochVWAP(bytes32 ) view returns(uint256)
func (f *FuturesSession) MostRecentEpochVWAP(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentEpochVWAP(&f.CallOpts, arg0)
}

// MostRecentEpochVWAP is a free data retrieval call binding the contract method 0xcf03660b.
//
// Solidity: function mostRecentEpochVWAP(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) MostRecentEpochVWAP(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentEpochVWAP(&f.CallOpts, arg0)
}

// MostRecentEpochVolume is a free data retrieval call binding the contract method 0x13626422.
//
// Solidity: function mostRecentEpochVolume(bytes32 ) view returns(uint256)
func (f *FuturesCaller) MostRecentEpochVolume(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "mostRecentEpochVolume", arg0)
	return *ret0, err
}

// MostRecentEpochVolume is a free data retrieval call binding the contract method 0x13626422.
//
// Solidity: function mostRecentEpochVolume(bytes32 ) view returns(uint256)
func (f *FuturesSession) MostRecentEpochVolume(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentEpochVolume(&f.CallOpts, arg0)
}

// MostRecentEpochVolume is a free data retrieval call binding the contract method 0x13626422.
//
// Solidity: function mostRecentEpochVolume(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) MostRecentEpochVolume(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentEpochVolume(&f.CallOpts, arg0)
}

// OrderHashToPositionID is a free data retrieval call binding the contract method 0x7c261203.
//
// Solidity: function orderHashToPositionID(bytes32 ) view returns(uint256)
func (f *FuturesCaller) OrderHashToPositionID(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "orderHashToPositionID", arg0)
	return *ret0, err
}

// OrderHashToPositionID is a free data retrieval call binding the contract method 0x7c261203.
//
// Solidity: function orderHashToPositionID(bytes32 ) view returns(uint256)
func (f *FuturesSession) OrderHashToPositionID(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.OrderHashToPositionID(&f.CallOpts, arg0)
}

// OrderHashToPositionID is a free data retrieval call binding the contract method 0x7c261203.
//
// Solidity: function orderHashToPositionID(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) OrderHashToPositionID(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.OrderHashToPositionID(&f.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (f *FuturesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := f.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (f *FuturesSession) Owner() (common.Address, error) {
	return f.Contract.Owner(&f.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (f *FuturesCallerSession) Owner() (common.Address, error) {
	return f.Contract.Owner(&f.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (f *FuturesCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := f.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (f *FuturesSession) Paused() (bool, error) {
	return f.Contract.Paused(&f.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (f *FuturesCallerSession) Paused() (bool, error) {
	return f.Contract.Paused(&f.CallOpts)
}

// PositionCount is a free data retrieval call binding the contract method 0xe7702d05.
//
// Solidity: function positionCount() view returns(uint256)
func (f *FuturesCaller) PositionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "positionCount")
	return *ret0, err
}

// PositionCount is a free data retrieval call binding the contract method 0xe7702d05.
//
// Solidity: function positionCount() view returns(uint256)
func (f *FuturesSession) PositionCount() (*big.Int, error) {
	return f.Contract.PositionCount(&f.CallOpts)
}

// PositionCount is a free data retrieval call binding the contract method 0xe7702d05.
//
// Solidity: function positionCount() view returns(uint256)
func (f *FuturesCallerSession) PositionCount() (*big.Int, error) {
	return f.Contract.PositionCount(&f.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 subAccountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, uint256 margin, int256 cumulativeFundingEntry, bytes32 orderHash)
func (f *FuturesCaller) Positions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SubAccountID           [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	ContractPrice          *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
	OrderHash              [32]byte
}, error) {
	ret := new(struct {
		SubAccountID           [32]byte
		MarketID               [32]byte
		Direction              uint8
		Quantity               *big.Int
		ContractPrice          *big.Int
		Margin                 *big.Int
		CumulativeFundingEntry *big.Int
		OrderHash              [32]byte
	})
	out := ret
	err := f.contract.Call(opts, out, "positions", arg0)
	return *ret, err
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 subAccountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, uint256 margin, int256 cumulativeFundingEntry, bytes32 orderHash)
func (f *FuturesSession) Positions(arg0 *big.Int) (struct {
	SubAccountID           [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	ContractPrice          *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
	OrderHash              [32]byte
}, error) {
	return f.Contract.Positions(&f.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 subAccountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, uint256 margin, int256 cumulativeFundingEntry, bytes32 orderHash)
func (f *FuturesCallerSession) Positions(arg0 *big.Int) (struct {
	SubAccountID           [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	ContractPrice          *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
	OrderHash              [32]byte
}, error) {
	return f.Contract.Positions(&f.CallOpts, arg0)
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (f *FuturesCaller) PreSigned(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := f.contract.Call(opts, out, "preSigned", arg0, arg1)
	return *ret0, err
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (f *FuturesSession) PreSigned(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return f.Contract.PreSigned(&f.CallOpts, arg0, arg1)
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (f *FuturesCallerSession) PreSigned(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return f.Contract.PreSigned(&f.CallOpts, arg0, arg1)
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x475ca0bb.
//
// Solidity: function restrictedDeposits(bytes32 , address ) view returns(uint256)
func (f *FuturesCaller) RestrictedDeposits(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "restrictedDeposits", arg0, arg1)
	return *ret0, err
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x475ca0bb.
//
// Solidity: function restrictedDeposits(bytes32 , address ) view returns(uint256)
func (f *FuturesSession) RestrictedDeposits(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return f.Contract.RestrictedDeposits(&f.CallOpts, arg0, arg1)
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x475ca0bb.
//
// Solidity: function restrictedDeposits(bytes32 , address ) view returns(uint256)
func (f *FuturesCallerSession) RestrictedDeposits(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return f.Contract.RestrictedDeposits(&f.CallOpts, arg0, arg1)
}

// SubAccountDeposits is a free data retrieval call binding the contract method 0x666ffb9b.
//
// Solidity: function subAccountDeposits(bytes32 , address ) view returns(uint256)
func (f *FuturesCaller) SubAccountDeposits(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "subAccountDeposits", arg0, arg1)
	return *ret0, err
}

// SubAccountDeposits is a free data retrieval call binding the contract method 0x666ffb9b.
//
// Solidity: function subAccountDeposits(bytes32 , address ) view returns(uint256)
func (f *FuturesSession) SubAccountDeposits(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return f.Contract.SubAccountDeposits(&f.CallOpts, arg0, arg1)
}

// SubAccountDeposits is a free data retrieval call binding the contract method 0x666ffb9b.
//
// Solidity: function subAccountDeposits(bytes32 , address ) view returns(uint256)
func (f *FuturesCallerSession) SubAccountDeposits(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return f.Contract.SubAccountDeposits(&f.CallOpts, arg0, arg1)
}

// SubAccountIdToAddress is a free data retrieval call binding the contract method 0x234842eb.
//
// Solidity: function subAccountIdToAddress(bytes32 ) view returns(address)
func (f *FuturesCaller) SubAccountIdToAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := f.contract.Call(opts, out, "subAccountIdToAddress", arg0)
	return *ret0, err
}

// SubAccountIdToAddress is a free data retrieval call binding the contract method 0x234842eb.
//
// Solidity: function subAccountIdToAddress(bytes32 ) view returns(address)
func (f *FuturesSession) SubAccountIdToAddress(arg0 [32]byte) (common.Address, error) {
	return f.Contract.SubAccountIdToAddress(&f.CallOpts, arg0)
}

// SubAccountIdToAddress is a free data retrieval call binding the contract method 0x234842eb.
//
// Solidity: function subAccountIdToAddress(bytes32 ) view returns(address)
func (f *FuturesCallerSession) SubAccountIdToAddress(arg0 [32]byte) (common.Address, error) {
	return f.Contract.SubAccountIdToAddress(&f.CallOpts, arg0)
}

// SubAccountToMarketToPositionID is a free data retrieval call binding the contract method 0x1ebcc120.
//
// Solidity: function subAccountToMarketToPositionID(bytes32 , bytes32 ) view returns(uint256)
func (f *FuturesCaller) SubAccountToMarketToPositionID(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "subAccountToMarketToPositionID", arg0, arg1)
	return *ret0, err
}

// SubAccountToMarketToPositionID is a free data retrieval call binding the contract method 0x1ebcc120.
//
// Solidity: function subAccountToMarketToPositionID(bytes32 , bytes32 ) view returns(uint256)
func (f *FuturesSession) SubAccountToMarketToPositionID(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return f.Contract.SubAccountToMarketToPositionID(&f.CallOpts, arg0, arg1)
}

// SubAccountToMarketToPositionID is a free data retrieval call binding the contract method 0x1ebcc120.
//
// Solidity: function subAccountToMarketToPositionID(bytes32 , bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) SubAccountToMarketToPositionID(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return f.Contract.SubAccountToMarketToPositionID(&f.CallOpts, arg0, arg1)
}

// AddMarginIntoPosition is a paid mutator transaction binding the contract method 0x62c6985e.
//
// Solidity: function addMarginIntoPosition(bytes32 subAccountID, uint256 positionID, uint256 addedMargin) returns()
func (f *FuturesTransactor) AddMarginIntoPosition(opts *bind.TransactOpts, subAccountID [32]byte, positionID *big.Int, addedMargin *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "addMarginIntoPosition", subAccountID, positionID, addedMargin)
}

// AddMarginIntoPosition is a paid mutator transaction binding the contract method 0x62c6985e.
//
// Solidity: function addMarginIntoPosition(bytes32 subAccountID, uint256 positionID, uint256 addedMargin) returns()
func (f *FuturesSession) AddMarginIntoPosition(subAccountID [32]byte, positionID *big.Int, addedMargin *big.Int) (*types.Transaction, error) {
	return f.Contract.AddMarginIntoPosition(&f.TransactOpts, subAccountID, positionID, addedMargin)
}

// AddMarginIntoPosition is a paid mutator transaction binding the contract method 0x62c6985e.
//
// Solidity: function addMarginIntoPosition(bytes32 subAccountID, uint256 positionID, uint256 addedMargin) returns()
func (f *FuturesTransactorSession) AddMarginIntoPosition(subAccountID [32]byte, positionID *big.Int, addedMargin *big.Int) (*types.Transaction, error) {
	return f.Contract.AddMarginIntoPosition(&f.TransactOpts, subAccountID, positionID, addedMargin)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0xc8be86b5.
//
// Solidity: function batchFillOrKillOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32[] subAccountIDs, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) BatchFillOrKillOrders(opts *bind.TransactOpts, orders []Order, quantities []*big.Int, margins []*big.Int, subAccountIDs [][32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchFillOrKillOrders", orders, quantities, margins, subAccountIDs, signatures)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0xc8be86b5.
//
// Solidity: function batchFillOrKillOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32[] subAccountIDs, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) BatchFillOrKillOrders(orders []Order, quantities []*big.Int, margins []*big.Int, subAccountIDs [][32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchFillOrKillOrders(&f.TransactOpts, orders, quantities, margins, subAccountIDs, signatures)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0xc8be86b5.
//
// Solidity: function batchFillOrKillOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32[] subAccountIDs, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) BatchFillOrKillOrders(orders []Order, quantities []*big.Int, margins []*big.Int, subAccountIDs [][32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchFillOrKillOrders(&f.TransactOpts, orders, quantities, margins, subAccountIDs, signatures)
}

// BatchFillOrKillOrdersSinglePosition is a paid mutator transaction binding the contract method 0xfdcb3d4f.
//
// Solidity: function batchFillOrKillOrdersSinglePosition((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) BatchFillOrKillOrdersSinglePosition(opts *bind.TransactOpts, orders []Order, quantities []*big.Int, margins []*big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchFillOrKillOrdersSinglePosition", orders, quantities, margins, subAccountID, signatures)
}

// BatchFillOrKillOrdersSinglePosition is a paid mutator transaction binding the contract method 0xfdcb3d4f.
//
// Solidity: function batchFillOrKillOrdersSinglePosition((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) BatchFillOrKillOrdersSinglePosition(orders []Order, quantities []*big.Int, margins []*big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchFillOrKillOrdersSinglePosition(&f.TransactOpts, orders, quantities, margins, subAccountID, signatures)
}

// BatchFillOrKillOrdersSinglePosition is a paid mutator transaction binding the contract method 0xfdcb3d4f.
//
// Solidity: function batchFillOrKillOrdersSinglePosition((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) BatchFillOrKillOrdersSinglePosition(orders []Order, quantities []*big.Int, margins []*big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchFillOrKillOrdersSinglePosition(&f.TransactOpts, orders, quantities, margins, subAccountID, signatures)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0x2e401b78.
//
// Solidity: function batchFillOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32[] subAccountIDs, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) BatchFillOrders(opts *bind.TransactOpts, orders []Order, quantities []*big.Int, margins []*big.Int, subAccountIDs [][32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchFillOrders", orders, quantities, margins, subAccountIDs, signatures)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0x2e401b78.
//
// Solidity: function batchFillOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32[] subAccountIDs, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) BatchFillOrders(orders []Order, quantities []*big.Int, margins []*big.Int, subAccountIDs [][32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchFillOrders(&f.TransactOpts, orders, quantities, margins, subAccountIDs, signatures)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0x2e401b78.
//
// Solidity: function batchFillOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32[] subAccountIDs, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) BatchFillOrders(orders []Order, quantities []*big.Int, margins []*big.Int, subAccountIDs [][32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchFillOrders(&f.TransactOpts, orders, quantities, margins, subAccountIDs, signatures)
}

// BatchFillOrdersSinglePosition is a paid mutator transaction binding the contract method 0x0dc32e0d.
//
// Solidity: function batchFillOrdersSinglePosition((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) BatchFillOrdersSinglePosition(opts *bind.TransactOpts, orders []Order, quantities []*big.Int, margins []*big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchFillOrdersSinglePosition", orders, quantities, margins, subAccountID, signatures)
}

// BatchFillOrdersSinglePosition is a paid mutator transaction binding the contract method 0x0dc32e0d.
//
// Solidity: function batchFillOrdersSinglePosition((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) BatchFillOrdersSinglePosition(orders []Order, quantities []*big.Int, margins []*big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchFillOrdersSinglePosition(&f.TransactOpts, orders, quantities, margins, subAccountID, signatures)
}

// BatchFillOrdersSinglePosition is a paid mutator transaction binding the contract method 0x0dc32e0d.
//
// Solidity: function batchFillOrdersSinglePosition((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256[] quantities, uint256[] margins, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) BatchFillOrdersSinglePosition(orders []Order, quantities []*big.Int, margins []*big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchFillOrdersSinglePosition(&f.TransactOpts, orders, quantities, margins, subAccountID, signatures)
}

// BatchMatchOrders is a paid mutator transaction binding the contract method 0x6fcf3e9e.
//
// Solidity: function batchMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) BatchMatchOrders(opts *bind.TransactOpts, leftOrders []Order, rightOrders []Order, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchMatchOrders", leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// BatchMatchOrders is a paid mutator transaction binding the contract method 0x6fcf3e9e.
//
// Solidity: function batchMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) BatchMatchOrders(leftOrders []Order, rightOrders []Order, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchMatchOrders(&f.TransactOpts, leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// BatchMatchOrders is a paid mutator transaction binding the contract method 0x6fcf3e9e.
//
// Solidity: function batchMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) BatchMatchOrders(leftOrders []Order, rightOrders []Order, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchMatchOrders(&f.TransactOpts, leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) returns()
func (f *FuturesTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return f.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) returns()
func (f *FuturesSession) CancelOrder(order Order) (*types.Transaction, error) {
	return f.Contract.CancelOrder(&f.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) returns()
func (f *FuturesTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return f.Contract.CancelOrder(&f.TransactOpts, order)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x89667e07.
//
// Solidity: function closePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesTransactor) ClosePosition(opts *bind.TransactOpts, positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "closePosition", positionID, orders, quantity, signatures)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x89667e07.
//
// Solidity: function closePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesSession) ClosePosition(positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.ClosePosition(&f.TransactOpts, positionID, orders, quantity, signatures)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x89667e07.
//
// Solidity: function closePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesTransactorSession) ClosePosition(positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.ClosePosition(&f.TransactOpts, positionID, orders, quantity, signatures)
}

// ClosePositionOrKill is a paid mutator transaction binding the contract method 0x1b5841de.
//
// Solidity: function closePositionOrKill(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesTransactor) ClosePositionOrKill(opts *bind.TransactOpts, positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "closePositionOrKill", positionID, orders, quantity, signatures)
}

// ClosePositionOrKill is a paid mutator transaction binding the contract method 0x1b5841de.
//
// Solidity: function closePositionOrKill(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesSession) ClosePositionOrKill(positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.ClosePositionOrKill(&f.TransactOpts, positionID, orders, quantity, signatures)
}

// ClosePositionOrKill is a paid mutator transaction binding the contract method 0x1b5841de.
//
// Solidity: function closePositionOrKill(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesTransactorSession) ClosePositionOrKill(positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.ClosePositionOrKill(&f.TransactOpts, positionID, orders, quantity, signatures)
}

// CreateDefaultSubAccountAndDeposit is a paid mutator transaction binding the contract method 0xcaa5cf87.
//
// Solidity: function createDefaultSubAccountAndDeposit(address baseCurrency, uint256 amount) returns(bytes32)
func (f *FuturesTransactor) CreateDefaultSubAccountAndDeposit(opts *bind.TransactOpts, baseCurrency common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createDefaultSubAccountAndDeposit", baseCurrency, amount)
}

// CreateDefaultSubAccountAndDeposit is a paid mutator transaction binding the contract method 0xcaa5cf87.
//
// Solidity: function createDefaultSubAccountAndDeposit(address baseCurrency, uint256 amount) returns(bytes32)
func (f *FuturesSession) CreateDefaultSubAccountAndDeposit(baseCurrency common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateDefaultSubAccountAndDeposit(&f.TransactOpts, baseCurrency, amount)
}

// CreateDefaultSubAccountAndDeposit is a paid mutator transaction binding the contract method 0xcaa5cf87.
//
// Solidity: function createDefaultSubAccountAndDeposit(address baseCurrency, uint256 amount) returns(bytes32)
func (f *FuturesTransactorSession) CreateDefaultSubAccountAndDeposit(baseCurrency common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateDefaultSubAccountAndDeposit(&f.TransactOpts, baseCurrency, amount)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xa592e451.
//
// Solidity: function createMarket(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatioFactor, (uint256) maintenanceMarginRatio, uint256 fundingInterval, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage) returns()
func (f *FuturesTransactor) CreateMarket(opts *bind.TransactOpts, ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatioFactor PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createMarket", ticker, baseCurrency, oracle, initialMarginRatioFactor, maintenanceMarginRatio, fundingInterval, makerTxFee, takerTxFee, relayerFeePercentage)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xa592e451.
//
// Solidity: function createMarket(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatioFactor, (uint256) maintenanceMarginRatio, uint256 fundingInterval, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage) returns()
func (f *FuturesSession) CreateMarket(ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatioFactor PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.CreateMarket(&f.TransactOpts, ticker, baseCurrency, oracle, initialMarginRatioFactor, maintenanceMarginRatio, fundingInterval, makerTxFee, takerTxFee, relayerFeePercentage)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xa592e451.
//
// Solidity: function createMarket(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatioFactor, (uint256) maintenanceMarginRatio, uint256 fundingInterval, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage) returns()
func (f *FuturesTransactorSession) CreateMarket(ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatioFactor PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.CreateMarket(&f.TransactOpts, ticker, baseCurrency, oracle, initialMarginRatioFactor, maintenanceMarginRatio, fundingInterval, makerTxFee, takerTxFee, relayerFeePercentage)
}

// CreateMarketWithFixedMarketId is a paid mutator transaction binding the contract method 0x266389c2.
//
// Solidity: function createMarketWithFixedMarketId(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatioFactor, (uint256) maintenanceMarginRatio, uint256 fundingInterval, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage, bytes32 marketID) returns()
func (f *FuturesTransactor) CreateMarketWithFixedMarketId(opts *bind.TransactOpts, ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatioFactor PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad, marketID [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createMarketWithFixedMarketId", ticker, baseCurrency, oracle, initialMarginRatioFactor, maintenanceMarginRatio, fundingInterval, makerTxFee, takerTxFee, relayerFeePercentage, marketID)
}

// CreateMarketWithFixedMarketId is a paid mutator transaction binding the contract method 0x266389c2.
//
// Solidity: function createMarketWithFixedMarketId(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatioFactor, (uint256) maintenanceMarginRatio, uint256 fundingInterval, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage, bytes32 marketID) returns()
func (f *FuturesSession) CreateMarketWithFixedMarketId(ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatioFactor PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad, marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.CreateMarketWithFixedMarketId(&f.TransactOpts, ticker, baseCurrency, oracle, initialMarginRatioFactor, maintenanceMarginRatio, fundingInterval, makerTxFee, takerTxFee, relayerFeePercentage, marketID)
}

// CreateMarketWithFixedMarketId is a paid mutator transaction binding the contract method 0x266389c2.
//
// Solidity: function createMarketWithFixedMarketId(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatioFactor, (uint256) maintenanceMarginRatio, uint256 fundingInterval, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage, bytes32 marketID) returns()
func (f *FuturesTransactorSession) CreateMarketWithFixedMarketId(ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatioFactor PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad, marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.CreateMarketWithFixedMarketId(&f.TransactOpts, ticker, baseCurrency, oracle, initialMarginRatioFactor, maintenanceMarginRatio, fundingInterval, makerTxFee, takerTxFee, relayerFeePercentage, marketID)
}

// CreateSubAccountAndDeposit is a paid mutator transaction binding the contract method 0x3bc22e90.
//
// Solidity: function createSubAccountAndDeposit(address baseCurrency, uint256 subAccountNonce, uint256 amount) returns(bytes32)
func (f *FuturesTransactor) CreateSubAccountAndDeposit(opts *bind.TransactOpts, baseCurrency common.Address, subAccountNonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createSubAccountAndDeposit", baseCurrency, subAccountNonce, amount)
}

// CreateSubAccountAndDeposit is a paid mutator transaction binding the contract method 0x3bc22e90.
//
// Solidity: function createSubAccountAndDeposit(address baseCurrency, uint256 subAccountNonce, uint256 amount) returns(bytes32)
func (f *FuturesSession) CreateSubAccountAndDeposit(baseCurrency common.Address, subAccountNonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateSubAccountAndDeposit(&f.TransactOpts, baseCurrency, subAccountNonce, amount)
}

// CreateSubAccountAndDeposit is a paid mutator transaction binding the contract method 0x3bc22e90.
//
// Solidity: function createSubAccountAndDeposit(address baseCurrency, uint256 subAccountNonce, uint256 amount) returns(bytes32)
func (f *FuturesTransactorSession) CreateSubAccountAndDeposit(baseCurrency common.Address, subAccountNonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateSubAccountAndDeposit(&f.TransactOpts, baseCurrency, subAccountNonce, amount)
}

// CreateSubAccountForTraderWithNonce is a paid mutator transaction binding the contract method 0x2d1fb098.
//
// Solidity: function createSubAccountForTraderWithNonce(address trader, uint256 subAccountNonce) returns(bytes32)
func (f *FuturesTransactor) CreateSubAccountForTraderWithNonce(opts *bind.TransactOpts, trader common.Address, subAccountNonce *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createSubAccountForTraderWithNonce", trader, subAccountNonce)
}

// CreateSubAccountForTraderWithNonce is a paid mutator transaction binding the contract method 0x2d1fb098.
//
// Solidity: function createSubAccountForTraderWithNonce(address trader, uint256 subAccountNonce) returns(bytes32)
func (f *FuturesSession) CreateSubAccountForTraderWithNonce(trader common.Address, subAccountNonce *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateSubAccountForTraderWithNonce(&f.TransactOpts, trader, subAccountNonce)
}

// CreateSubAccountForTraderWithNonce is a paid mutator transaction binding the contract method 0x2d1fb098.
//
// Solidity: function createSubAccountForTraderWithNonce(address trader, uint256 subAccountNonce) returns(bytes32)
func (f *FuturesTransactorSession) CreateSubAccountForTraderWithNonce(trader common.Address, subAccountNonce *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateSubAccountForTraderWithNonce(&f.TransactOpts, trader, subAccountNonce)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address baseCurrency, uint256 amount) returns()
func (f *FuturesTransactor) Deposit(opts *bind.TransactOpts, baseCurrency common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "deposit", baseCurrency, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address baseCurrency, uint256 amount) returns()
func (f *FuturesSession) Deposit(baseCurrency common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Deposit(&f.TransactOpts, baseCurrency, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address baseCurrency, uint256 amount) returns()
func (f *FuturesTransactorSession) Deposit(baseCurrency common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Deposit(&f.TransactOpts, baseCurrency, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0xb3db428b.
//
// Solidity: function depositFor(address baseCurrency, address trader, uint256 amount) returns()
func (f *FuturesTransactor) DepositFor(opts *bind.TransactOpts, baseCurrency common.Address, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "depositFor", baseCurrency, trader, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0xb3db428b.
//
// Solidity: function depositFor(address baseCurrency, address trader, uint256 amount) returns()
func (f *FuturesSession) DepositFor(baseCurrency common.Address, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositFor(&f.TransactOpts, baseCurrency, trader, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0xb3db428b.
//
// Solidity: function depositFor(address baseCurrency, address trader, uint256 amount) returns()
func (f *FuturesTransactorSession) DepositFor(baseCurrency common.Address, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositFor(&f.TransactOpts, baseCurrency, trader, amount)
}

// DepositIntoSubAccount is a paid mutator transaction binding the contract method 0xf2848b01.
//
// Solidity: function depositIntoSubAccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesTransactor) DepositIntoSubAccount(opts *bind.TransactOpts, baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "depositIntoSubAccount", baseCurrency, subAccountID, amount)
}

// DepositIntoSubAccount is a paid mutator transaction binding the contract method 0xf2848b01.
//
// Solidity: function depositIntoSubAccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesSession) DepositIntoSubAccount(baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositIntoSubAccount(&f.TransactOpts, baseCurrency, subAccountID, amount)
}

// DepositIntoSubAccount is a paid mutator transaction binding the contract method 0xf2848b01.
//
// Solidity: function depositIntoSubAccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesTransactorSession) DepositIntoSubAccount(baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositIntoSubAccount(&f.TransactOpts, baseCurrency, subAccountID, amount)
}

// EmergencyStopFutures is a paid mutator transaction binding the contract method 0x0d64dc42.
//
// Solidity: function emergencyStopFutures() returns()
func (f *FuturesTransactor) EmergencyStopFutures(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.contract.Transact(opts, "emergencyStopFutures")
}

// EmergencyStopFutures is a paid mutator transaction binding the contract method 0x0d64dc42.
//
// Solidity: function emergencyStopFutures() returns()
func (f *FuturesSession) EmergencyStopFutures() (*types.Transaction, error) {
	return f.Contract.EmergencyStopFutures(&f.TransactOpts)
}

// EmergencyStopFutures is a paid mutator transaction binding the contract method 0x0d64dc42.
//
// Solidity: function emergencyStopFutures() returns()
func (f *FuturesTransactorSession) EmergencyStopFutures() (*types.Transaction, error) {
	return f.Contract.EmergencyStopFutures(&f.TransactOpts)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0x4024fd0e.
//
// Solidity: function fillOrKillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes signature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactor) FillOrKillOrder(opts *bind.TransactOpts, order Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "fillOrKillOrder", order, quantity, margin, subAccountID, signature)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0x4024fd0e.
//
// Solidity: function fillOrKillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes signature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) results)
func (f *FuturesSession) FillOrKillOrder(order Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signature []byte) (*types.Transaction, error) {
	return f.Contract.FillOrKillOrder(&f.TransactOpts, order, quantity, margin, subAccountID, signature)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0x4024fd0e.
//
// Solidity: function fillOrKillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes signature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactorSession) FillOrKillOrder(order Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signature []byte) (*types.Transaction, error) {
	return f.Contract.FillOrKillOrder(&f.TransactOpts, order, quantity, margin, subAccountID, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x57b6abf8.
//
// Solidity: function fillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes signature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (f *FuturesTransactor) FillOrder(opts *bind.TransactOpts, order Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "fillOrder", order, quantity, margin, subAccountID, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x57b6abf8.
//
// Solidity: function fillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes signature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (f *FuturesSession) FillOrder(order Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signature []byte) (*types.Transaction, error) {
	return f.Contract.FillOrder(&f.TransactOpts, order, quantity, margin, subAccountID, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x57b6abf8.
//
// Solidity: function fillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes signature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (f *FuturesTransactorSession) FillOrder(order Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signature []byte) (*types.Transaction, error) {
	return f.Contract.FillOrder(&f.TransactOpts, order, quantity, margin, subAccountID, signature)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0x472275aa.
//
// Solidity: function liquidatePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesTransactor) LiquidatePosition(opts *bind.TransactOpts, positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "liquidatePosition", positionID, orders, quantity, signatures)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0x472275aa.
//
// Solidity: function liquidatePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesSession) LiquidatePosition(positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.LiquidatePosition(&f.TransactOpts, positionID, orders, quantity, signatures)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0x472275aa.
//
// Solidity: function liquidatePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesTransactorSession) LiquidatePosition(positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.LiquidatePosition(&f.TransactOpts, positionID, orders, quantity, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0x5209a8e7.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) MarketOrders(opts *bind.TransactOpts, orders []Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "marketOrders", orders, quantity, margin, subAccountID, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0x5209a8e7.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) MarketOrders(orders []Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.MarketOrders(&f.TransactOpts, orders, quantity, margin, subAccountID, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0x5209a8e7.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) MarketOrders(orders []Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.MarketOrders(&f.TransactOpts, orders, quantity, margin, subAccountID, signatures)
}

// MarketOrdersOrKill is a paid mutator transaction binding the contract method 0x52863e4f.
//
// Solidity: function marketOrdersOrKill((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) MarketOrdersOrKill(opts *bind.TransactOpts, orders []Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "marketOrdersOrKill", orders, quantity, margin, subAccountID, signatures)
}

// MarketOrdersOrKill is a paid mutator transaction binding the contract method 0x52863e4f.
//
// Solidity: function marketOrdersOrKill((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) MarketOrdersOrKill(orders []Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.MarketOrdersOrKill(&f.TransactOpts, orders, quantity, margin, subAccountID, signatures)
}

// MarketOrdersOrKill is a paid mutator transaction binding the contract method 0x52863e4f.
//
// Solidity: function marketOrdersOrKill((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes32 subAccountID, bytes[] signatures) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) MarketOrdersOrKill(orders []Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.MarketOrdersOrKill(&f.TransactOpts, orders, quantity, margin, subAccountID, signatures)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (f *FuturesTransactor) MatchOrders(opts *bind.TransactOpts, leftOrder Order, rightOrder Order, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "matchOrders", leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (f *FuturesSession) MatchOrders(leftOrder Order, rightOrder Order, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MatchOrders(&f.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (f *FuturesTransactorSession) MatchOrders(leftOrder Order, rightOrder Order, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MatchOrders(&f.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes[] leftSignatures, bytes rightSignature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) MultiMatchOrders(opts *bind.TransactOpts, leftOrders []Order, rightOrder Order, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "multiMatchOrders", leftOrders, rightOrder, leftSignatures, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes[] leftSignatures, bytes rightSignature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) MultiMatchOrders(leftOrders []Order, rightOrder Order, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MultiMatchOrders(&f.TransactOpts, leftOrders, rightOrder, leftSignatures, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes[] leftSignatures, bytes rightSignature) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) MultiMatchOrders(leftOrders []Order, rightOrder Order, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MultiMatchOrders(&f.TransactOpts, leftOrders, rightOrder, leftSignatures, rightSignature)
}

// PayIntoInsurancePool is a paid mutator transaction binding the contract method 0x95563906.
//
// Solidity: function payIntoInsurancePool(bytes32 marketID, uint256 amount) returns()
func (f *FuturesTransactor) PayIntoInsurancePool(opts *bind.TransactOpts, marketID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "payIntoInsurancePool", marketID, amount)
}

// PayIntoInsurancePool is a paid mutator transaction binding the contract method 0x95563906.
//
// Solidity: function payIntoInsurancePool(bytes32 marketID, uint256 amount) returns()
func (f *FuturesSession) PayIntoInsurancePool(marketID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.PayIntoInsurancePool(&f.TransactOpts, marketID, amount)
}

// PayIntoInsurancePool is a paid mutator transaction binding the contract method 0x95563906.
//
// Solidity: function payIntoInsurancePool(bytes32 marketID, uint256 amount) returns()
func (f *FuturesTransactorSession) PayIntoInsurancePool(marketID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.PayIntoInsurancePool(&f.TransactOpts, marketID, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (f *FuturesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (f *FuturesSession) RenounceOwnership() (*types.Transaction, error) {
	return f.Contract.RenounceOwnership(&f.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (f *FuturesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return f.Contract.RenounceOwnership(&f.TransactOpts)
}

// ResumeFutures is a paid mutator transaction binding the contract method 0x55585bce.
//
// Solidity: function resumeFutures() returns()
func (f *FuturesTransactor) ResumeFutures(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.contract.Transact(opts, "resumeFutures")
}

// ResumeFutures is a paid mutator transaction binding the contract method 0x55585bce.
//
// Solidity: function resumeFutures() returns()
func (f *FuturesSession) ResumeFutures() (*types.Transaction, error) {
	return f.Contract.ResumeFutures(&f.TransactOpts)
}

// ResumeFutures is a paid mutator transaction binding the contract method 0x55585bce.
//
// Solidity: function resumeFutures() returns()
func (f *FuturesTransactorSession) ResumeFutures() (*types.Transaction, error) {
	return f.Contract.ResumeFutures(&f.TransactOpts)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x33db9b03.
//
// Solidity: function setFundingRate(bytes32 marketID) returns()
func (f *FuturesTransactor) SetFundingRate(opts *bind.TransactOpts, marketID [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "setFundingRate", marketID)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x33db9b03.
//
// Solidity: function setFundingRate(bytes32 marketID) returns()
func (f *FuturesSession) SetFundingRate(marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.SetFundingRate(&f.TransactOpts, marketID)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x33db9b03.
//
// Solidity: function setFundingRate(bytes32 marketID) returns()
func (f *FuturesTransactorSession) SetFundingRate(marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.SetFundingRate(&f.TransactOpts, marketID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (f *FuturesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return f.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (f *FuturesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return f.Contract.TransferOwnership(&f.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (f *FuturesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return f.Contract.TransferOwnership(&f.TransactOpts, newOwner)
}

// VaporizePosition is a paid mutator transaction binding the contract method 0x20f3c024.
//
// Solidity: function vaporizePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesTransactor) VaporizePosition(opts *bind.TransactOpts, positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "vaporizePosition", positionID, orders, quantity, signatures)
}

// VaporizePosition is a paid mutator transaction binding the contract method 0x20f3c024.
//
// Solidity: function vaporizePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesSession) VaporizePosition(positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.VaporizePosition(&f.TransactOpts, positionID, orders, quantity, signatures)
}

// VaporizePosition is a paid mutator transaction binding the contract method 0x20f3c024.
//
// Solidity: function vaporizePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256)[] pResults, (int256,uint256) cResults)
func (f *FuturesTransactorSession) VaporizePosition(positionID *big.Int, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.VaporizePosition(&f.TransactOpts, positionID, orders, quantity, signatures)
}

// WithdrawFromDefaultSubAcount is a paid mutator transaction binding the contract method 0x1f630454.
//
// Solidity: function withdrawFromDefaultSubAcount(address baseCurrency, uint256 amount) returns()
func (f *FuturesTransactor) WithdrawFromDefaultSubAcount(opts *bind.TransactOpts, baseCurrency common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "withdrawFromDefaultSubAcount", baseCurrency, amount)
}

// WithdrawFromDefaultSubAcount is a paid mutator transaction binding the contract method 0x1f630454.
//
// Solidity: function withdrawFromDefaultSubAcount(address baseCurrency, uint256 amount) returns()
func (f *FuturesSession) WithdrawFromDefaultSubAcount(baseCurrency common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawFromDefaultSubAcount(&f.TransactOpts, baseCurrency, amount)
}

// WithdrawFromDefaultSubAcount is a paid mutator transaction binding the contract method 0x1f630454.
//
// Solidity: function withdrawFromDefaultSubAcount(address baseCurrency, uint256 amount) returns()
func (f *FuturesTransactorSession) WithdrawFromDefaultSubAcount(baseCurrency common.Address, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawFromDefaultSubAcount(&f.TransactOpts, baseCurrency, amount)
}

// WithdrawFromSubAccount is a paid mutator transaction binding the contract method 0xbf59e058.
//
// Solidity: function withdrawFromSubAccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesTransactor) WithdrawFromSubAccount(opts *bind.TransactOpts, baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "withdrawFromSubAccount", baseCurrency, subAccountID, amount)
}

// WithdrawFromSubAccount is a paid mutator transaction binding the contract method 0xbf59e058.
//
// Solidity: function withdrawFromSubAccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesSession) WithdrawFromSubAccount(baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawFromSubAccount(&f.TransactOpts, baseCurrency, subAccountID, amount)
}

// WithdrawFromSubAccount is a paid mutator transaction binding the contract method 0xbf59e058.
//
// Solidity: function withdrawFromSubAccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesTransactorSession) WithdrawFromSubAccount(baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawFromSubAccount(&f.TransactOpts, baseCurrency, subAccountID, amount)
}

// FuturesAccountCreationIterator is returned from FilterAccountCreation and is used to iterate over the raw logs and unpacked data for AccountCreation events raised by the Futures contract.
type FuturesAccountCreationIterator struct {
	Event *FuturesAccountCreation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesAccountCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesAccountCreation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesAccountCreation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesAccountCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesAccountCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesAccountCreation represents a AccountCreation event raised by the Futures contract.
type FuturesAccountCreation struct {
	Creator         common.Address
	SubAccountID    [32]byte
	SubAccountNonce *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountCreation is a free log retrieval operation binding the contract event 0x4a6d2d5a911a10d00446d1413ca969c249074711c144e93aeddfa4d77de64784.
//
// Solidity: event AccountCreation(address indexed creator, bytes32 subAccountID, uint256 subAccountNonce)
func (f *FuturesFilterer) FilterAccountCreation(opts *bind.FilterOpts, creator []common.Address) (*FuturesAccountCreationIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "AccountCreation", creatorRule)
	if err != nil {
		return nil, err
	}
	return &FuturesAccountCreationIterator{contract: f.contract, event: "AccountCreation", logs: logs, sub: sub}, nil
}

// WatchAccountCreation is a free log subscription operation binding the contract event 0x4a6d2d5a911a10d00446d1413ca969c249074711c144e93aeddfa4d77de64784.
//
// Solidity: event AccountCreation(address indexed creator, bytes32 subAccountID, uint256 subAccountNonce)
func (f *FuturesFilterer) WatchAccountCreation(opts *bind.WatchOpts, sink chan<- *FuturesAccountCreation, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "AccountCreation", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesAccountCreation)
				if err := f.contract.UnpackLog(event, "AccountCreation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountCreation is a log parse operation binding the contract event 0x4a6d2d5a911a10d00446d1413ca969c249074711c144e93aeddfa4d77de64784.
//
// Solidity: event AccountCreation(address indexed creator, bytes32 subAccountID, uint256 subAccountNonce)
func (f *FuturesFilterer) ParseAccountCreation(log types.Log) (*FuturesAccountCreation, error) {
	event := new(FuturesAccountCreation)
	if err := f.contract.UnpackLog(event, "AccountCreation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesDecrementSubaccountDepositsIterator is returned from FilterDecrementSubaccountDeposits and is used to iterate over the raw logs and unpacked data for DecrementSubaccountDeposits events raised by the Futures contract.
type FuturesDecrementSubaccountDepositsIterator struct {
	Event *FuturesDecrementSubaccountDeposits // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesDecrementSubaccountDepositsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesDecrementSubaccountDeposits)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesDecrementSubaccountDeposits)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesDecrementSubaccountDepositsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesDecrementSubaccountDepositsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesDecrementSubaccountDeposits represents a DecrementSubaccountDeposits event raised by the Futures contract.
type FuturesDecrementSubaccountDeposits struct {
	SubAccountID [32]byte
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDecrementSubaccountDeposits is a free log retrieval operation binding the contract event 0x07cbd2a0bd322d5c23e10d71ecca0a9f41b9c9597b2f8736607b9435324e4357.
//
// Solidity: event DecrementSubaccountDeposits(bytes32 indexed subAccountID, uint256 amount)
func (f *FuturesFilterer) FilterDecrementSubaccountDeposits(opts *bind.FilterOpts, subAccountID [][32]byte) (*FuturesDecrementSubaccountDepositsIterator, error) {

	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "DecrementSubaccountDeposits", subAccountIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesDecrementSubaccountDepositsIterator{contract: f.contract, event: "DecrementSubaccountDeposits", logs: logs, sub: sub}, nil
}

// WatchDecrementSubaccountDeposits is a free log subscription operation binding the contract event 0x07cbd2a0bd322d5c23e10d71ecca0a9f41b9c9597b2f8736607b9435324e4357.
//
// Solidity: event DecrementSubaccountDeposits(bytes32 indexed subAccountID, uint256 amount)
func (f *FuturesFilterer) WatchDecrementSubaccountDeposits(opts *bind.WatchOpts, sink chan<- *FuturesDecrementSubaccountDeposits, subAccountID [][32]byte) (event.Subscription, error) {

	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "DecrementSubaccountDeposits", subAccountIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesDecrementSubaccountDeposits)
				if err := f.contract.UnpackLog(event, "DecrementSubaccountDeposits", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecrementSubaccountDeposits is a log parse operation binding the contract event 0x07cbd2a0bd322d5c23e10d71ecca0a9f41b9c9597b2f8736607b9435324e4357.
//
// Solidity: event DecrementSubaccountDeposits(bytes32 indexed subAccountID, uint256 amount)
func (f *FuturesFilterer) ParseDecrementSubaccountDeposits(log types.Log) (*FuturesDecrementSubaccountDeposits, error) {
	event := new(FuturesDecrementSubaccountDeposits)
	if err := f.contract.UnpackLog(event, "DecrementSubaccountDeposits", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesCancelIterator is returned from FilterFuturesCancel and is used to iterate over the raw logs and unpacked data for FuturesCancel events raised by the Futures contract.
type FuturesCancelIterator struct {
	Event *FuturesCancel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesCancel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesCancel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesCancel represents a FuturesCancel event raised by the Futures contract.
type FuturesCancel struct {
	MakerAddress   common.Address
	OrderHash      [32]byte
	MarketID       [32]byte
	ContractPrice  *big.Int
	QuantityFilled *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFuturesCancel is a free log retrieval operation binding the contract event 0x414118d90fd71dbfe3eebc508a8edaebe20d4e43ac23c65ba56fe87edb7c61ca.
//
// Solidity: event FuturesCancel(address indexed makerAddress, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled)
func (f *FuturesFilterer) FilterFuturesCancel(opts *bind.FilterOpts, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (*FuturesCancelIterator, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesCancel", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesCancelIterator{contract: f.contract, event: "FuturesCancel", logs: logs, sub: sub}, nil
}

// WatchFuturesCancel is a free log subscription operation binding the contract event 0x414118d90fd71dbfe3eebc508a8edaebe20d4e43ac23c65ba56fe87edb7c61ca.
//
// Solidity: event FuturesCancel(address indexed makerAddress, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled)
func (f *FuturesFilterer) WatchFuturesCancel(opts *bind.WatchOpts, sink chan<- *FuturesCancel, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (event.Subscription, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesCancel", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesCancel)
				if err := f.contract.UnpackLog(event, "FuturesCancel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFuturesCancel is a log parse operation binding the contract event 0x414118d90fd71dbfe3eebc508a8edaebe20d4e43ac23c65ba56fe87edb7c61ca.
//
// Solidity: event FuturesCancel(address indexed makerAddress, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled)
func (f *FuturesFilterer) ParseFuturesCancel(log types.Log) (*FuturesCancel, error) {
	event := new(FuturesCancel)
	if err := f.contract.UnpackLog(event, "FuturesCancel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesCloseIterator is returned from FilterFuturesClose and is used to iterate over the raw logs and unpacked data for FuturesClose events raised by the Futures contract.
type FuturesCloseIterator struct {
	Event *FuturesClose // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesClose)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesClose)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesClose represents a FuturesClose event raised by the Futures contract.
type FuturesClose struct {
	PositionID   *big.Int
	MarketID     [32]byte
	SubAccountID [32]byte
	Quantity     *big.Int
	ContractPNL  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFuturesClose is a free log retrieval operation binding the contract event 0x552b29e4ee304f6809a963a25ef66f4922981189771e56a21c4639b51b3ebff2.
//
// Solidity: event FuturesClose(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed subAccountID, uint256 quantity, int256 contractPNL)
func (f *FuturesFilterer) FilterFuturesClose(opts *bind.FilterOpts, positionID []*big.Int, marketID [][32]byte, subAccountID [][32]byte) (*FuturesCloseIterator, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesClose", positionIDRule, marketIDRule, subAccountIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesCloseIterator{contract: f.contract, event: "FuturesClose", logs: logs, sub: sub}, nil
}

// WatchFuturesClose is a free log subscription operation binding the contract event 0x552b29e4ee304f6809a963a25ef66f4922981189771e56a21c4639b51b3ebff2.
//
// Solidity: event FuturesClose(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed subAccountID, uint256 quantity, int256 contractPNL)
func (f *FuturesFilterer) WatchFuturesClose(opts *bind.WatchOpts, sink chan<- *FuturesClose, positionID []*big.Int, marketID [][32]byte, subAccountID [][32]byte) (event.Subscription, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesClose", positionIDRule, marketIDRule, subAccountIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesClose)
				if err := f.contract.UnpackLog(event, "FuturesClose", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFuturesClose is a log parse operation binding the contract event 0x552b29e4ee304f6809a963a25ef66f4922981189771e56a21c4639b51b3ebff2.
//
// Solidity: event FuturesClose(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed subAccountID, uint256 quantity, int256 contractPNL)
func (f *FuturesFilterer) ParseFuturesClose(log types.Log) (*FuturesClose, error) {
	event := new(FuturesClose)
	if err := f.contract.UnpackLog(event, "FuturesClose", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesLiquidationIterator is returned from FilterFuturesLiquidation and is used to iterate over the raw logs and unpacked data for FuturesLiquidation events raised by the Futures contract.
type FuturesLiquidationIterator struct {
	Event *FuturesLiquidation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesLiquidation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesLiquidation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesLiquidation represents a FuturesLiquidation event raised by the Futures contract.
type FuturesLiquidation struct {
	PositionID   *big.Int
	MarketID     [32]byte
	SubAccountID [32]byte
	Quantity     *big.Int
	ContractPNL  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFuturesLiquidation is a free log retrieval operation binding the contract event 0xf7def0e827224d4fdf4354d5d330812ba6798aa95fc60d5ed8864290cb04d451.
//
// Solidity: event FuturesLiquidation(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed subAccountID, uint256 quantity, int256 contractPNL)
func (f *FuturesFilterer) FilterFuturesLiquidation(opts *bind.FilterOpts, positionID []*big.Int, marketID [][32]byte, subAccountID [][32]byte) (*FuturesLiquidationIterator, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesLiquidation", positionIDRule, marketIDRule, subAccountIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesLiquidationIterator{contract: f.contract, event: "FuturesLiquidation", logs: logs, sub: sub}, nil
}

// WatchFuturesLiquidation is a free log subscription operation binding the contract event 0xf7def0e827224d4fdf4354d5d330812ba6798aa95fc60d5ed8864290cb04d451.
//
// Solidity: event FuturesLiquidation(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed subAccountID, uint256 quantity, int256 contractPNL)
func (f *FuturesFilterer) WatchFuturesLiquidation(opts *bind.WatchOpts, sink chan<- *FuturesLiquidation, positionID []*big.Int, marketID [][32]byte, subAccountID [][32]byte) (event.Subscription, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesLiquidation", positionIDRule, marketIDRule, subAccountIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesLiquidation)
				if err := f.contract.UnpackLog(event, "FuturesLiquidation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFuturesLiquidation is a log parse operation binding the contract event 0xf7def0e827224d4fdf4354d5d330812ba6798aa95fc60d5ed8864290cb04d451.
//
// Solidity: event FuturesLiquidation(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed subAccountID, uint256 quantity, int256 contractPNL)
func (f *FuturesFilterer) ParseFuturesLiquidation(log types.Log) (*FuturesLiquidation, error) {
	event := new(FuturesLiquidation)
	if err := f.contract.UnpackLog(event, "FuturesLiquidation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesMatchIterator is returned from FilterFuturesMatch and is used to iterate over the raw logs and unpacked data for FuturesMatch events raised by the Futures contract.
type FuturesMatchIterator struct {
	Event *FuturesMatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesMatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesMatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesMatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesMatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesMatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesMatch represents a FuturesMatch event raised by the Futures contract.
type FuturesMatch struct {
	LeftOrderHash  [32]byte
	RightOrderHash [32]byte
	MarketID       [32]byte
	Quantity       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFuturesMatch is a free log retrieval operation binding the contract event 0xf6dc3b6914c1ade29b4a0c34e7d47e49d51690d622548649e7f67ced15b09a06.
//
// Solidity: event FuturesMatch(bytes32 indexed leftOrderHash, bytes32 indexed rightOrderHash, bytes32 indexed marketID, uint256 quantity)
func (f *FuturesFilterer) FilterFuturesMatch(opts *bind.FilterOpts, leftOrderHash [][32]byte, rightOrderHash [][32]byte, marketID [][32]byte) (*FuturesMatchIterator, error) {

	var leftOrderHashRule []interface{}
	for _, leftOrderHashItem := range leftOrderHash {
		leftOrderHashRule = append(leftOrderHashRule, leftOrderHashItem)
	}
	var rightOrderHashRule []interface{}
	for _, rightOrderHashItem := range rightOrderHash {
		rightOrderHashRule = append(rightOrderHashRule, rightOrderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesMatch", leftOrderHashRule, rightOrderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesMatchIterator{contract: f.contract, event: "FuturesMatch", logs: logs, sub: sub}, nil
}

// WatchFuturesMatch is a free log subscription operation binding the contract event 0xf6dc3b6914c1ade29b4a0c34e7d47e49d51690d622548649e7f67ced15b09a06.
//
// Solidity: event FuturesMatch(bytes32 indexed leftOrderHash, bytes32 indexed rightOrderHash, bytes32 indexed marketID, uint256 quantity)
func (f *FuturesFilterer) WatchFuturesMatch(opts *bind.WatchOpts, sink chan<- *FuturesMatch, leftOrderHash [][32]byte, rightOrderHash [][32]byte, marketID [][32]byte) (event.Subscription, error) {

	var leftOrderHashRule []interface{}
	for _, leftOrderHashItem := range leftOrderHash {
		leftOrderHashRule = append(leftOrderHashRule, leftOrderHashItem)
	}
	var rightOrderHashRule []interface{}
	for _, rightOrderHashItem := range rightOrderHash {
		rightOrderHashRule = append(rightOrderHashRule, rightOrderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesMatch", leftOrderHashRule, rightOrderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesMatch)
				if err := f.contract.UnpackLog(event, "FuturesMatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFuturesMatch is a log parse operation binding the contract event 0xf6dc3b6914c1ade29b4a0c34e7d47e49d51690d622548649e7f67ced15b09a06.
//
// Solidity: event FuturesMatch(bytes32 indexed leftOrderHash, bytes32 indexed rightOrderHash, bytes32 indexed marketID, uint256 quantity)
func (f *FuturesFilterer) ParseFuturesMatch(log types.Log) (*FuturesMatch, error) {
	event := new(FuturesMatch)
	if err := f.contract.UnpackLog(event, "FuturesMatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesOrderFillIterator is returned from FilterFuturesOrderFill and is used to iterate over the raw logs and unpacked data for FuturesOrderFill events raised by the Futures contract.
type FuturesOrderFillIterator struct {
	Event *FuturesOrderFill // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesOrderFillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesOrderFill)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesOrderFill)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesOrderFillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesOrderFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesOrderFill represents a FuturesOrderFill event raised by the Futures contract.
type FuturesOrderFill struct {
	OrderHash   [32]byte
	MarketID    [32]byte
	TotalFilled *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFuturesOrderFill is a free log retrieval operation binding the contract event 0xa7c00a70596823b554350f2addbfe1f5b49630f6ff4b8cfebb76e5c37ea402d8.
//
// Solidity: event FuturesOrderFill(bytes32 indexed orderHash, bytes32 indexed marketID, uint256 totalFilled)
func (f *FuturesFilterer) FilterFuturesOrderFill(opts *bind.FilterOpts, orderHash [][32]byte, marketID [][32]byte) (*FuturesOrderFillIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesOrderFill", orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesOrderFillIterator{contract: f.contract, event: "FuturesOrderFill", logs: logs, sub: sub}, nil
}

// WatchFuturesOrderFill is a free log subscription operation binding the contract event 0xa7c00a70596823b554350f2addbfe1f5b49630f6ff4b8cfebb76e5c37ea402d8.
//
// Solidity: event FuturesOrderFill(bytes32 indexed orderHash, bytes32 indexed marketID, uint256 totalFilled)
func (f *FuturesFilterer) WatchFuturesOrderFill(opts *bind.WatchOpts, sink chan<- *FuturesOrderFill, orderHash [][32]byte, marketID [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesOrderFill", orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesOrderFill)
				if err := f.contract.UnpackLog(event, "FuturesOrderFill", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFuturesOrderFill is a log parse operation binding the contract event 0xa7c00a70596823b554350f2addbfe1f5b49630f6ff4b8cfebb76e5c37ea402d8.
//
// Solidity: event FuturesOrderFill(bytes32 indexed orderHash, bytes32 indexed marketID, uint256 totalFilled)
func (f *FuturesFilterer) ParseFuturesOrderFill(log types.Log) (*FuturesOrderFill, error) {
	event := new(FuturesOrderFill)
	if err := f.contract.UnpackLog(event, "FuturesOrderFill", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesPositionIterator is returned from FilterFuturesPosition and is used to iterate over the raw logs and unpacked data for FuturesPosition events raised by the Futures contract.
type FuturesPositionIterator struct {
	Event *FuturesPosition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesPositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesPosition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesPosition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesPositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesPositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesPosition represents a FuturesPosition event raised by the Futures contract.
type FuturesPosition struct {
	MakerAddress           common.Address
	AccountId              [32]byte
	OrderHash              [32]byte
	MarketID               [32]byte
	ContractPrice          *big.Int
	QuantityFilled         *big.Int
	TotalQuantity          *big.Int
	InitialMargin          *big.Int
	CumulativeFundingEntry *big.Int
	PositionID             *big.Int
	IsLong                 bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFuturesPosition is a free log retrieval operation binding the contract event 0x3f15c65539a543c3d4f963f69449601ff9cb5921e1cf2872cef61f57127ded65.
//
// Solidity: event FuturesPosition(address indexed makerAddress, bytes32 accountId, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled, uint256 totalQuantity, uint256 initialMargin, int256 cumulativeFundingEntry, uint256 positionID, bool isLong)
func (f *FuturesFilterer) FilterFuturesPosition(opts *bind.FilterOpts, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (*FuturesPositionIterator, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesPosition", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesPositionIterator{contract: f.contract, event: "FuturesPosition", logs: logs, sub: sub}, nil
}

// WatchFuturesPosition is a free log subscription operation binding the contract event 0x3f15c65539a543c3d4f963f69449601ff9cb5921e1cf2872cef61f57127ded65.
//
// Solidity: event FuturesPosition(address indexed makerAddress, bytes32 accountId, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled, uint256 totalQuantity, uint256 initialMargin, int256 cumulativeFundingEntry, uint256 positionID, bool isLong)
func (f *FuturesFilterer) WatchFuturesPosition(opts *bind.WatchOpts, sink chan<- *FuturesPosition, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (event.Subscription, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesPosition", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesPosition)
				if err := f.contract.UnpackLog(event, "FuturesPosition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFuturesPosition is a log parse operation binding the contract event 0x3f15c65539a543c3d4f963f69449601ff9cb5921e1cf2872cef61f57127ded65.
//
// Solidity: event FuturesPosition(address indexed makerAddress, bytes32 accountId, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled, uint256 totalQuantity, uint256 initialMargin, int256 cumulativeFundingEntry, uint256 positionID, bool isLong)
func (f *FuturesFilterer) ParseFuturesPosition(log types.Log) (*FuturesPosition, error) {
	event := new(FuturesPosition)
	if err := f.contract.UnpackLog(event, "FuturesPosition", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesIncrementSubaccountDepositsIterator is returned from FilterIncrementSubaccountDeposits and is used to iterate over the raw logs and unpacked data for IncrementSubaccountDeposits events raised by the Futures contract.
type FuturesIncrementSubaccountDepositsIterator struct {
	Event *FuturesIncrementSubaccountDeposits // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesIncrementSubaccountDepositsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesIncrementSubaccountDeposits)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesIncrementSubaccountDeposits)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesIncrementSubaccountDepositsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesIncrementSubaccountDepositsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesIncrementSubaccountDeposits represents a IncrementSubaccountDeposits event raised by the Futures contract.
type FuturesIncrementSubaccountDeposits struct {
	SubAccountID [32]byte
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIncrementSubaccountDeposits is a free log retrieval operation binding the contract event 0xa645c0b79353c4e5184d0db3b4af9170887ad392f0b06f835cfdab50e329b37f.
//
// Solidity: event IncrementSubaccountDeposits(bytes32 indexed subAccountID, uint256 amount)
func (f *FuturesFilterer) FilterIncrementSubaccountDeposits(opts *bind.FilterOpts, subAccountID [][32]byte) (*FuturesIncrementSubaccountDepositsIterator, error) {

	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "IncrementSubaccountDeposits", subAccountIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesIncrementSubaccountDepositsIterator{contract: f.contract, event: "IncrementSubaccountDeposits", logs: logs, sub: sub}, nil
}

// WatchIncrementSubaccountDeposits is a free log subscription operation binding the contract event 0xa645c0b79353c4e5184d0db3b4af9170887ad392f0b06f835cfdab50e329b37f.
//
// Solidity: event IncrementSubaccountDeposits(bytes32 indexed subAccountID, uint256 amount)
func (f *FuturesFilterer) WatchIncrementSubaccountDeposits(opts *bind.WatchOpts, sink chan<- *FuturesIncrementSubaccountDeposits, subAccountID [][32]byte) (event.Subscription, error) {

	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "IncrementSubaccountDeposits", subAccountIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesIncrementSubaccountDeposits)
				if err := f.contract.UnpackLog(event, "IncrementSubaccountDeposits", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncrementSubaccountDeposits is a log parse operation binding the contract event 0xa645c0b79353c4e5184d0db3b4af9170887ad392f0b06f835cfdab50e329b37f.
//
// Solidity: event IncrementSubaccountDeposits(bytes32 indexed subAccountID, uint256 amount)
func (f *FuturesFilterer) ParseIncrementSubaccountDeposits(log types.Log) (*FuturesIncrementSubaccountDeposits, error) {
	event := new(FuturesIncrementSubaccountDeposits)
	if err := f.contract.UnpackLog(event, "IncrementSubaccountDeposits", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesMarginAddedIterator is returned from FilterMarginAdded and is used to iterate over the raw logs and unpacked data for MarginAdded events raised by the Futures contract.
type FuturesMarginAddedIterator struct {
	Event *FuturesMarginAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesMarginAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesMarginAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesMarginAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesMarginAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesMarginAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesMarginAdded represents a MarginAdded event raised by the Futures contract.
type FuturesMarginAdded struct {
	PositionID  *big.Int
	AddedMargin *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarginAdded is a free log retrieval operation binding the contract event 0xe1f22f90ef2dc8cc8b52420487417e2d986cc8001133dff107fda9b9d8a7395b.
//
// Solidity: event MarginAdded(uint256 indexed positionID, uint256 addedMargin)
func (f *FuturesFilterer) FilterMarginAdded(opts *bind.FilterOpts, positionID []*big.Int) (*FuturesMarginAddedIterator, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "MarginAdded", positionIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesMarginAddedIterator{contract: f.contract, event: "MarginAdded", logs: logs, sub: sub}, nil
}

// WatchMarginAdded is a free log subscription operation binding the contract event 0xe1f22f90ef2dc8cc8b52420487417e2d986cc8001133dff107fda9b9d8a7395b.
//
// Solidity: event MarginAdded(uint256 indexed positionID, uint256 addedMargin)
func (f *FuturesFilterer) WatchMarginAdded(opts *bind.WatchOpts, sink chan<- *FuturesMarginAdded, positionID []*big.Int) (event.Subscription, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "MarginAdded", positionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesMarginAdded)
				if err := f.contract.UnpackLog(event, "MarginAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarginAdded is a log parse operation binding the contract event 0xe1f22f90ef2dc8cc8b52420487417e2d986cc8001133dff107fda9b9d8a7395b.
//
// Solidity: event MarginAdded(uint256 indexed positionID, uint256 addedMargin)
func (f *FuturesFilterer) ParseMarginAdded(log types.Log) (*FuturesMarginAdded, error) {
	event := new(FuturesMarginAdded)
	if err := f.contract.UnpackLog(event, "MarginAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesMarketCreationIterator is returned from FilterMarketCreation and is used to iterate over the raw logs and unpacked data for MarketCreation events raised by the Futures contract.
type FuturesMarketCreationIterator struct {
	Event *FuturesMarketCreation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesMarketCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesMarketCreation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesMarketCreation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesMarketCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesMarketCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesMarketCreation represents a MarketCreation event raised by the Futures contract.
type FuturesMarketCreation struct {
	MarketID                 [32]byte
	Ticker                   common.Hash
	Oracle                   common.Address
	BaseCurrency             common.Address
	MaintenanceMarginRatio   *big.Int
	InitialMarginRatioFactor *big.Int
	MakerTxFee               *big.Int
	TakerTxFee               *big.Int
	RelayerFeePercentage     *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterMarketCreation is a free log retrieval operation binding the contract event 0x76b32833b43d38848ea77d0d66b06aa8011c537b6693cb0109470c967b098f9c.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle, address baseCurrency, uint256 maintenanceMarginRatio, uint256 initialMarginRatioFactor, uint256 makerTxFee, uint256 takerTxFee, uint256 relayerFeePercentage)
func (f *FuturesFilterer) FilterMarketCreation(opts *bind.FilterOpts, marketID [][32]byte, ticker []string, oracle []common.Address) (*FuturesMarketCreationIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "MarketCreation", marketIDRule, tickerRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &FuturesMarketCreationIterator{contract: f.contract, event: "MarketCreation", logs: logs, sub: sub}, nil
}

// WatchMarketCreation is a free log subscription operation binding the contract event 0x76b32833b43d38848ea77d0d66b06aa8011c537b6693cb0109470c967b098f9c.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle, address baseCurrency, uint256 maintenanceMarginRatio, uint256 initialMarginRatioFactor, uint256 makerTxFee, uint256 takerTxFee, uint256 relayerFeePercentage)
func (f *FuturesFilterer) WatchMarketCreation(opts *bind.WatchOpts, sink chan<- *FuturesMarketCreation, marketID [][32]byte, ticker []string, oracle []common.Address) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "MarketCreation", marketIDRule, tickerRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesMarketCreation)
				if err := f.contract.UnpackLog(event, "MarketCreation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketCreation is a log parse operation binding the contract event 0x76b32833b43d38848ea77d0d66b06aa8011c537b6693cb0109470c967b098f9c.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle, address baseCurrency, uint256 maintenanceMarginRatio, uint256 initialMarginRatioFactor, uint256 makerTxFee, uint256 takerTxFee, uint256 relayerFeePercentage)
func (f *FuturesFilterer) ParseMarketCreation(log types.Log) (*FuturesMarketCreation, error) {
	event := new(FuturesMarketCreation)
	if err := f.contract.UnpackLog(event, "MarketCreation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Futures contract.
type FuturesOwnershipTransferredIterator struct {
	Event *FuturesOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesOwnershipTransferred represents a OwnershipTransferred event raised by the Futures contract.
type FuturesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (f *FuturesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FuturesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FuturesOwnershipTransferredIterator{contract: f.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (f *FuturesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FuturesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesOwnershipTransferred)
				if err := f.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (f *FuturesFilterer) ParseOwnershipTransferred(log types.Log) (*FuturesOwnershipTransferred, error) {
	event := new(FuturesOwnershipTransferred)
	if err := f.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Futures contract.
type FuturesPausedIterator struct {
	Event *FuturesPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesPaused represents a Paused event raised by the Futures contract.
type FuturesPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (f *FuturesFilterer) FilterPaused(opts *bind.FilterOpts) (*FuturesPausedIterator, error) {

	logs, sub, err := f.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FuturesPausedIterator{contract: f.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (f *FuturesFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FuturesPaused) (event.Subscription, error) {

	logs, sub, err := f.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesPaused)
				if err := f.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (f *FuturesFilterer) ParsePaused(log types.Log) (*FuturesPaused, error) {
	event := new(FuturesPaused)
	if err := f.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesSetFundingIterator is returned from FilterSetFunding and is used to iterate over the raw logs and unpacked data for SetFunding events raised by the Futures contract.
type FuturesSetFundingIterator struct {
	Event *FuturesSetFunding // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesSetFundingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesSetFunding)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesSetFunding)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesSetFundingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesSetFundingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesSetFunding represents a SetFunding event raised by the Futures contract.
type FuturesSetFunding struct {
	MarketID   [32]byte
	FundingFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetFunding is a free log retrieval operation binding the contract event 0x51f338889900f0109c0f0f69c3ec20862654ee807dd48394843129d257069320.
//
// Solidity: event SetFunding(bytes32 indexed marketID, int256 fundingFee)
func (f *FuturesFilterer) FilterSetFunding(opts *bind.FilterOpts, marketID [][32]byte) (*FuturesSetFundingIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "SetFunding", marketIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesSetFundingIterator{contract: f.contract, event: "SetFunding", logs: logs, sub: sub}, nil
}

// WatchSetFunding is a free log subscription operation binding the contract event 0x51f338889900f0109c0f0f69c3ec20862654ee807dd48394843129d257069320.
//
// Solidity: event SetFunding(bytes32 indexed marketID, int256 fundingFee)
func (f *FuturesFilterer) WatchSetFunding(opts *bind.WatchOpts, sink chan<- *FuturesSetFunding, marketID [][32]byte) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "SetFunding", marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesSetFunding)
				if err := f.contract.UnpackLog(event, "SetFunding", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFunding is a log parse operation binding the contract event 0x51f338889900f0109c0f0f69c3ec20862654ee807dd48394843129d257069320.
//
// Solidity: event SetFunding(bytes32 indexed marketID, int256 fundingFee)
func (f *FuturesFilterer) ParseSetFunding(log types.Log) (*FuturesSetFunding, error) {
	event := new(FuturesSetFunding)
	if err := f.contract.UnpackLog(event, "SetFunding", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesSignatureValidatorApprovalIterator is returned from FilterSignatureValidatorApproval and is used to iterate over the raw logs and unpacked data for SignatureValidatorApproval events raised by the Futures contract.
type FuturesSignatureValidatorApprovalIterator struct {
	Event *FuturesSignatureValidatorApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesSignatureValidatorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesSignatureValidatorApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesSignatureValidatorApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesSignatureValidatorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesSignatureValidatorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesSignatureValidatorApproval represents a SignatureValidatorApproval event raised by the Futures contract.
type FuturesSignatureValidatorApproval struct {
	SignerAddress    common.Address
	ValidatorAddress common.Address
	IsApproved       bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSignatureValidatorApproval is a free log retrieval operation binding the contract event 0xa8656e308026eeabce8f0bc18048433252318ab80ac79da0b3d3d8697dfba891.
//
// Solidity: event SignatureValidatorApproval(address indexed signerAddress, address indexed validatorAddress, bool isApproved)
func (f *FuturesFilterer) FilterSignatureValidatorApproval(opts *bind.FilterOpts, signerAddress []common.Address, validatorAddress []common.Address) (*FuturesSignatureValidatorApprovalIterator, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "SignatureValidatorApproval", signerAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &FuturesSignatureValidatorApprovalIterator{contract: f.contract, event: "SignatureValidatorApproval", logs: logs, sub: sub}, nil
}

// WatchSignatureValidatorApproval is a free log subscription operation binding the contract event 0xa8656e308026eeabce8f0bc18048433252318ab80ac79da0b3d3d8697dfba891.
//
// Solidity: event SignatureValidatorApproval(address indexed signerAddress, address indexed validatorAddress, bool isApproved)
func (f *FuturesFilterer) WatchSignatureValidatorApproval(opts *bind.WatchOpts, sink chan<- *FuturesSignatureValidatorApproval, signerAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "SignatureValidatorApproval", signerAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesSignatureValidatorApproval)
				if err := f.contract.UnpackLog(event, "SignatureValidatorApproval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignatureValidatorApproval is a log parse operation binding the contract event 0xa8656e308026eeabce8f0bc18048433252318ab80ac79da0b3d3d8697dfba891.
//
// Solidity: event SignatureValidatorApproval(address indexed signerAddress, address indexed validatorAddress, bool isApproved)
func (f *FuturesFilterer) ParseSignatureValidatorApproval(log types.Log) (*FuturesSignatureValidatorApproval, error) {
	event := new(FuturesSignatureValidatorApproval)
	if err := f.contract.UnpackLog(event, "SignatureValidatorApproval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FuturesUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Futures contract.
type FuturesUnpausedIterator struct {
	Event *FuturesUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FuturesUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FuturesUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FuturesUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesUnpaused represents a Unpaused event raised by the Futures contract.
type FuturesUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (f *FuturesFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FuturesUnpausedIterator, error) {

	logs, sub, err := f.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FuturesUnpausedIterator{contract: f.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (f *FuturesFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FuturesUnpaused) (event.Subscription, error) {

	logs, sub, err := f.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesUnpaused)
				if err := f.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (f *FuturesFilterer) ParseUnpaused(log types.Log) (*FuturesUnpaused, error) {
	event := new(FuturesUnpaused)
	if err := f.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
