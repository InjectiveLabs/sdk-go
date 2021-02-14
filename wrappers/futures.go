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

// LibFillResultsBatchMatchedFillResults is an auto generated low-level Go binding around an user-defined struct.
type LibFillResultsBatchMatchedFillResults struct {
	Left  []LibFillResultsFillResults
	Right []LibFillResultsFillResults
}

// LibFillResultsFillResults is an auto generated low-level Go binding around an user-defined struct.
type LibFillResultsFillResults struct {
	MakerAssetFilledAmount *big.Int
	TakerAssetFilledAmount *big.Int
	TakerFeePaid           *big.Int
	RelayerFeePaid         *big.Int
}

// LibFillResultsMatchedFillResults is an auto generated low-level Go binding around an user-defined struct.
type LibFillResultsMatchedFillResults struct {
	Left  LibFillResultsFillResults
	Right LibFillResultsFillResults
}

// DerivativeOrderInfo is an auto generated low-level Go binding around an user-defined struct.
type DerivativeOrderInfo struct {
	OrderStatus                 uint8
	OrderHash                   [32]byte
	OrderTakerAssetFilledAmount *big.Int
	SubAccountID                [32]byte
	Direction                   uint8
	MarketID                    [32]byte
	EntryPrice                  *big.Int
}

// LibZeroExTransactionZeroExTransaction is an auto generated low-level Go binding around an user-defined struct.
type LibZeroExTransactionZeroExTransaction struct {
	Salt                  *big.Int
	ExpirationTimeSeconds *big.Int
	GasPrice              *big.Int
	SignerAddress         common.Address
	Data                  []byte
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
	EntryPrice             *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
}

// TypesPositionResults is an auto generated low-level Go binding around an user-defined struct.
type TypesPositionResults struct {
	PositionID *big.Int
	MarginUsed *big.Int
	Quantity   *big.Int
	Fee        *big.Int
}

// TypesTransactionFees is an auto generated low-level Go binding around an user-defined struct.
type TypesTransactionFees struct {
	Maker   PermyriadMathPermyriad
	Taker   PermyriadMathPermyriad
	Relayer PermyriadMathPermyriad
}

// FuturesABI is the input ABI used to generate the binding from.
const FuturesABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"_minimumMarginRatio\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auctionTimeInterval_\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"injectiveToken_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"subAccountNonce\",\"type\":\"uint96\"}],\"name\":\"AccountCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"name\":\"FuturesCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maintenanceMarginRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialMarginRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerTxFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerTxFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayerFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"FuturesMarketCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isResultingPositionLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resultingMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resultingEntryPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resultingQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTradeDirectionLong\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalForOrderFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTypes.TradeOrTransferType\",\"name\":\"tradeType\",\"type\":\"uint8\"}],\"name\":\"FuturesTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"changedMargin\",\"type\":\"int256\"}],\"name\":\"MarginChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"senderPositionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiverPositionID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fromSubAccountID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"toSubAccountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumTypes.DirectionalStatus\",\"name\":\"directionalStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcResultingPositionMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcResultingPositionEntryPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcResultingPositionQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"srcResultingPositionCumulativeFundingEntry\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destResultingPositionMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destResultingPositionEntryPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destResultingPositionQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"destResultingPositionCumulativeFundingEntry\",\"type\":\"int256\"}],\"name\":\"PositionTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingFee\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"cumulativeFunding\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiryOrNextFundingTimestamp\",\"type\":\"uint256\"}],\"name\":\"SetFunding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"SignatureValidatorApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerTxFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerTxFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayerFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SpotMarketCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sentAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gainedAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gainedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentSentAssetDeposits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentGainedAssetDeposits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalForOrderFilled\",\"type\":\"uint256\"}],\"name\":\"SpotTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumMixinAccounts.DepositChangeType\",\"name\":\"depositChangeType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"changeAmount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentAmount\",\"type\":\"uint256\"}],\"name\":\"SubaccountDepositsChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mostRecentEpochVolume\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mostRecentEpochQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"mostRecentEpochScaledContractIndexDiff\",\"type\":\"int256\"}],\"name\":\"UpdateValuesForVWAP\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP1271_MAGIC_VALUE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_EXCHANGE_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_MARGIN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEC_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"subAccountNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedMargin\",\"type\":\"uint256\"}],\"name\":\"addMarginIntoPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressToSubAccountIDs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiverSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"}],\"name\":\"approveForReceiving\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"contractIAuction\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionTimeInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"baseCurrencies\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"batchCancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"marketIDs\",\"type\":\"bytes32[]\"}],\"name\":\"batchCheckFunding\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"baseCurrencies\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"subAccountIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchDepositForSubAccounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"baseCurrencies\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"traders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchDepositForTraders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structLibZeroExTransaction.ZeroExTransaction[]\",\"name\":\"transactions\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchExecuteTransactions\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"positionIDs\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[][]\",\"name\":\"orders\",\"type\":\"tuple[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"signatures\",\"type\":\"bytes[][]\"},{\"internalType\":\"address\",\"name\":\"liquidationCaller\",\"type\":\"address\"}],\"name\":\"batchLiquidatePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.PositionResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"subAccountIDs\",\"type\":\"bytes32[]\"}],\"name\":\"batchSettleExpiryFuturesPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.PositionResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"name\":\"calcCumulativeFunding\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"calcLiquidationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"calcMinMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionTimeInterval_\",\"type\":\"uint256\"}],\"name\":\"changeAuctionTimeInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"}],\"name\":\"changeMarketFundingInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"initialMarginRatio\",\"type\":\"tuple\"}],\"name\":\"changeMarketInitialMarginRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maintenanceMarginRatio\",\"type\":\"tuple\"}],\"name\":\"changeMarketMaintenanceMarginRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"makerTxFee\",\"type\":\"tuple\"}],\"name\":\"changeMarketMakerTxFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"changeMarketOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"}],\"name\":\"changeMarketOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"relayerTxFee\",\"type\":\"tuple\"}],\"name\":\"changeMarketRelayerTxFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"takerTxFee\",\"type\":\"tuple\"}],\"name\":\"changeMarketTakerTxFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBaseCurrencyCap_\",\"type\":\"uint256\"}],\"name\":\"changeMaxBaseCurrencyCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"MINIMUM_MARGIN_RATIO_\",\"type\":\"tuple\"}],\"name\":\"changeMinimumMarginRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"checkFunding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRevertingOnPartialFills\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"closePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.PositionResults\",\"name\":\"results\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRevertingOnPartialFills\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"closeQuantity\",\"type\":\"uint256\"}],\"name\":\"closePositionMetaTransaction\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"subAccountNonce\",\"type\":\"uint96\"}],\"name\":\"computeSubAccountIdFromNonce\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createDefaultSubAccountAndDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"initialMarginRatio\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maintenanceMarginRatio\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"makerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"takerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"relayerFeePercentage\",\"type\":\"tuple\"}],\"name\":\"createFuturesMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"initialMarginRatio\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maintenanceMarginRatio\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"makerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"takerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"relayerFeePercentage\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"createFuturesMarketWithFixedMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"makerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"takerTxFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"relayerFeePercentage\",\"type\":\"tuple\"}],\"name\":\"createSpotMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"subAccountNonce\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createSubAccountAndDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"subAccountNonce\",\"type\":\"uint96\"}],\"name\":\"createSubAccountForTraderWithNonce\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentAuctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentContextAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositForSubaccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositIntoSubAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"}],\"name\":\"doesPositionExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyStopFutures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"emergencyStopMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structLibZeroExTransaction.ZeroExTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"futuresMarkets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"initialMarginRatio\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maintenanceMarginRatio\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryOrNextFundingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFunding\",\"type\":\"int256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maker\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"taker\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"relayer\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransactionFees\",\"name\":\"transactionFees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getDefaultSubAccountDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getDefaultSubAccountIdForTrader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getMaxFundingFeeAbs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.SpotOrderStatus\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"}],\"internalType\":\"structLibOrder.OrderInfo\",\"name\":\"orderInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"}],\"name\":\"getOrderRelevantState\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.OrderStatus\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"}],\"internalType\":\"structLibOrder.DerivativeOrderInfo\",\"name\":\"orderInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fillableTakerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidSignature\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"getOrderRelevantStates\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.OrderStatus\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"}],\"internalType\":\"structLibOrder.DerivativeOrderInfo[]\",\"name\":\"ordersInfo\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fillableTakerAssetAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"isValidSignature\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getPositionIDsForTrader\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"positionIDs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getPositionsForTrader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"internalType\":\"structTypes.Position[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiverSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"}],\"name\":\"getReceiptApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"getTraderSubAccountsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceToUse\",\"type\":\"uint256\"}],\"name\":\"getUnitPositionValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"unitPositionValue\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"insurancePools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"receiverSubAccountID\",\"type\":\"bytes32\"}],\"name\":\"isAllowedToReceivePosition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isAllowedToTransferPosition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fromSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fromSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"isApprovedForMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isFuturesMarketSettled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiverSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isReceiptApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiverSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"isReceiptApprovedForMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"isTimeExpiryFuturesMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidBaseCurrency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidHashSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidOrderSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structLibZeroExTransaction.ZeroExTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidTransactionSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lastValidVWAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"liquidationCaller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"liquidatePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.PositionResults\",\"name\":\"results\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isRevertingOnPartialFills\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"}],\"name\":\"marketOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.FillResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketSerialToID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"leftOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"leftSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"matchSpotOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relayerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relayerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structLibFillResults.MatchedFillResults\",\"name\":\"matchedFillResults\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBaseCurrencyCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"mostRecentEpochQuantity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"mostRecentEpochVolume\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"mostRecentEpochWeightedAverageContractIndexDiff\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"mostRecentmostRecentEpochVolumeEpochQuantity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"leftOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"leftSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"multiMatchOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"leftPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightPositionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightMarginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structMixinOrders.MatchResults[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"leftOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"rightOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"leftSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"rightSignatures\",\"type\":\"bytes[]\"}],\"name\":\"multiMatchSpotOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relayerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults[]\",\"name\":\"left\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relayerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults[]\",\"name\":\"right\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibFillResults.BatchMatchedFillResults\",\"name\":\"batchMatchedFillResults\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payIntoInsurancePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"positionCountInTEF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"preSign\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"preSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"restrictedDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeFutures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fromSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fromSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"setApprovalForMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"setFundingRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiverSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"setReceiptApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiverSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"setReceiptApprovalForMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"setSignatureValidatorApproval\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"}],\"name\":\"settleExpiryFuturesPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.PositionResults\",\"name\":\"results\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"settleMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"spotMarkets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"maker\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"taker\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPermyriadMath.Permyriad\",\"name\":\"relayer\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransactionFees\",\"name\":\"transactionFees\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"subAccountDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subAccountIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subAccountToMarketToPositionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactionsExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"transferQuantity\",\"type\":\"uint256\"}],\"name\":\"transferPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fromSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"toSubAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToSubAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"vaporizePosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.PositionResults\",\"name\":\"results\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subAccountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawForSubAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromInsurancePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
func (f *FuturesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (f *FuturesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (f *FuturesCaller) EIP1271MAGICVALUE(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "EIP1271_MAGIC_VALUE")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (f *FuturesSession) EIP1271MAGICVALUE() ([4]byte, error) {
	return f.Contract.EIP1271MAGICVALUE(&f.CallOpts)
}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (f *FuturesCallerSession) EIP1271MAGICVALUE() ([4]byte, error) {
	return f.Contract.EIP1271MAGICVALUE(&f.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (f *FuturesCaller) EIP712EXCHANGEDOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "EIP712_EXCHANGE_DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "MINIMUM_MARGIN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// TECADDRESS is a free data retrieval call binding the contract method 0xe6077ac2.
//
// Solidity: function TEC_ADDRESS() view returns(address)
func (f *FuturesCaller) TECADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "TEC_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TECADDRESS is a free data retrieval call binding the contract method 0xe6077ac2.
//
// Solidity: function TEC_ADDRESS() view returns(address)
func (f *FuturesSession) TECADDRESS() (common.Address, error) {
	return f.Contract.TECADDRESS(&f.CallOpts)
}

// TECADDRESS is a free data retrieval call binding the contract method 0xe6077ac2.
//
// Solidity: function TEC_ADDRESS() view returns(address)
func (f *FuturesCallerSession) TECADDRESS() (common.Address, error) {
	return f.Contract.TECADDRESS(&f.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bytes32 subAccountID, uint256 subAccountNonce)
func (f *FuturesCaller) Accounts(opts *bind.CallOpts, arg0 [32]byte) (struct {
	SubAccountID    [32]byte
	SubAccountNonce *big.Int
}, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "accounts", arg0)

	outstruct := new(struct {
		SubAccountID    [32]byte
		SubAccountNonce *big.Int
	})

	outstruct.SubAccountID = out[0].([32]byte)
	outstruct.SubAccountNonce = out[1].(*big.Int)

	return *outstruct, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "addressToSubAccountIDs", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "allowedValidators", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (f *FuturesCaller) Auction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "auction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (f *FuturesSession) Auction() (common.Address, error) {
	return f.Contract.Auction(&f.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (f *FuturesCallerSession) Auction() (common.Address, error) {
	return f.Contract.Auction(&f.CallOpts)
}

// AuctionTimeInterval is a free data retrieval call binding the contract method 0xc583f691.
//
// Solidity: function auctionTimeInterval() view returns(uint256)
func (f *FuturesCaller) AuctionTimeInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "auctionTimeInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionTimeInterval is a free data retrieval call binding the contract method 0xc583f691.
//
// Solidity: function auctionTimeInterval() view returns(uint256)
func (f *FuturesSession) AuctionTimeInterval() (*big.Int, error) {
	return f.Contract.AuctionTimeInterval(&f.CallOpts)
}

// AuctionTimeInterval is a free data retrieval call binding the contract method 0xc583f691.
//
// Solidity: function auctionTimeInterval() view returns(uint256)
func (f *FuturesCallerSession) AuctionTimeInterval() (*big.Int, error) {
	return f.Contract.AuctionTimeInterval(&f.CallOpts)
}

// BaseCurrencies is a free data retrieval call binding the contract method 0x95092e50.
//
// Solidity: function baseCurrencies(uint256 ) view returns(address)
func (f *FuturesCaller) BaseCurrencies(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "baseCurrencies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseCurrencies is a free data retrieval call binding the contract method 0x95092e50.
//
// Solidity: function baseCurrencies(uint256 ) view returns(address)
func (f *FuturesSession) BaseCurrencies(arg0 *big.Int) (common.Address, error) {
	return f.Contract.BaseCurrencies(&f.CallOpts, arg0)
}

// BaseCurrencies is a free data retrieval call binding the contract method 0x95092e50.
//
// Solidity: function baseCurrencies(uint256 ) view returns(address)
func (f *FuturesCallerSession) BaseCurrencies(arg0 *big.Int) (common.Address, error) {
	return f.Contract.BaseCurrencies(&f.CallOpts, arg0)
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) view returns(int256)
func (f *FuturesCaller) CalcCumulativeFunding(opts *bind.CallOpts, marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "calcCumulativeFunding", marketID, cumulativeFundingEntry)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "calcLiquidationFee", marketID, quantity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "calcMinMargin", marketID, quantity, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "cancelled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// ClosePositionMetaTransaction is a free data retrieval call binding the contract method 0x1c51d905.
//
// Solidity: function closePositionMetaTransaction(address exchangeAddress, bool isRevertingOnPartialFills, bytes32 subAccountID, bytes32 marketID, uint256 closeQuantity) pure returns()
func (f *FuturesCaller) ClosePositionMetaTransaction(opts *bind.CallOpts, exchangeAddress common.Address, isRevertingOnPartialFills bool, subAccountID [32]byte, marketID [32]byte, closeQuantity *big.Int) error {
	var out []interface{}
	err := f.contract.Call(opts, &out, "closePositionMetaTransaction", exchangeAddress, isRevertingOnPartialFills, subAccountID, marketID, closeQuantity)

	if err != nil {
		return err
	}

	return err

}

// ClosePositionMetaTransaction is a free data retrieval call binding the contract method 0x1c51d905.
//
// Solidity: function closePositionMetaTransaction(address exchangeAddress, bool isRevertingOnPartialFills, bytes32 subAccountID, bytes32 marketID, uint256 closeQuantity) pure returns()
func (f *FuturesSession) ClosePositionMetaTransaction(exchangeAddress common.Address, isRevertingOnPartialFills bool, subAccountID [32]byte, marketID [32]byte, closeQuantity *big.Int) error {
	return f.Contract.ClosePositionMetaTransaction(&f.CallOpts, exchangeAddress, isRevertingOnPartialFills, subAccountID, marketID, closeQuantity)
}

// ClosePositionMetaTransaction is a free data retrieval call binding the contract method 0x1c51d905.
//
// Solidity: function closePositionMetaTransaction(address exchangeAddress, bool isRevertingOnPartialFills, bytes32 subAccountID, bytes32 marketID, uint256 closeQuantity) pure returns()
func (f *FuturesCallerSession) ClosePositionMetaTransaction(exchangeAddress common.Address, isRevertingOnPartialFills bool, subAccountID [32]byte, marketID [32]byte, closeQuantity *big.Int) error {
	return f.Contract.ClosePositionMetaTransaction(&f.CallOpts, exchangeAddress, isRevertingOnPartialFills, subAccountID, marketID, closeQuantity)
}

// ComputeSubAccountIdFromNonce is a free data retrieval call binding the contract method 0x2488c4d8.
//
// Solidity: function computeSubAccountIdFromNonce(address trader, uint96 subAccountNonce) pure returns(bytes32)
func (f *FuturesCaller) ComputeSubAccountIdFromNonce(opts *bind.CallOpts, trader common.Address, subAccountNonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "computeSubAccountIdFromNonce", trader, subAccountNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeSubAccountIdFromNonce is a free data retrieval call binding the contract method 0x2488c4d8.
//
// Solidity: function computeSubAccountIdFromNonce(address trader, uint96 subAccountNonce) pure returns(bytes32)
func (f *FuturesSession) ComputeSubAccountIdFromNonce(trader common.Address, subAccountNonce *big.Int) ([32]byte, error) {
	return f.Contract.ComputeSubAccountIdFromNonce(&f.CallOpts, trader, subAccountNonce)
}

// ComputeSubAccountIdFromNonce is a free data retrieval call binding the contract method 0x2488c4d8.
//
// Solidity: function computeSubAccountIdFromNonce(address trader, uint96 subAccountNonce) pure returns(bytes32)
func (f *FuturesCallerSession) ComputeSubAccountIdFromNonce(trader common.Address, subAccountNonce *big.Int) ([32]byte, error) {
	return f.Contract.ComputeSubAccountIdFromNonce(&f.CallOpts, trader, subAccountNonce)
}

// CurrentAuctionEndTime is a free data retrieval call binding the contract method 0xe196dff1.
//
// Solidity: function currentAuctionEndTime() view returns(uint256)
func (f *FuturesCaller) CurrentAuctionEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "currentAuctionEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentAuctionEndTime is a free data retrieval call binding the contract method 0xe196dff1.
//
// Solidity: function currentAuctionEndTime() view returns(uint256)
func (f *FuturesSession) CurrentAuctionEndTime() (*big.Int, error) {
	return f.Contract.CurrentAuctionEndTime(&f.CallOpts)
}

// CurrentAuctionEndTime is a free data retrieval call binding the contract method 0xe196dff1.
//
// Solidity: function currentAuctionEndTime() view returns(uint256)
func (f *FuturesCallerSession) CurrentAuctionEndTime() (*big.Int, error) {
	return f.Contract.CurrentAuctionEndTime(&f.CallOpts)
}

// CurrentContextAddress is a free data retrieval call binding the contract method 0xeea086ba.
//
// Solidity: function currentContextAddress() view returns(address)
func (f *FuturesCaller) CurrentContextAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "currentContextAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentContextAddress is a free data retrieval call binding the contract method 0xeea086ba.
//
// Solidity: function currentContextAddress() view returns(address)
func (f *FuturesSession) CurrentContextAddress() (common.Address, error) {
	return f.Contract.CurrentContextAddress(&f.CallOpts)
}

// CurrentContextAddress is a free data retrieval call binding the contract method 0xeea086ba.
//
// Solidity: function currentContextAddress() view returns(address)
func (f *FuturesCallerSession) CurrentContextAddress() (common.Address, error) {
	return f.Contract.CurrentContextAddress(&f.CallOpts)
}

// DoesPositionExist is a free data retrieval call binding the contract method 0x2c12d600.
//
// Solidity: function doesPositionExist(uint256 positionID) view returns(bool)
func (f *FuturesCaller) DoesPositionExist(opts *bind.CallOpts, positionID *big.Int) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "doesPositionExist", positionID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoesPositionExist is a free data retrieval call binding the contract method 0x2c12d600.
//
// Solidity: function doesPositionExist(uint256 positionID) view returns(bool)
func (f *FuturesSession) DoesPositionExist(positionID *big.Int) (bool, error) {
	return f.Contract.DoesPositionExist(&f.CallOpts, positionID)
}

// DoesPositionExist is a free data retrieval call binding the contract method 0x2c12d600.
//
// Solidity: function doesPositionExist(uint256 positionID) view returns(bool)
func (f *FuturesCallerSession) DoesPositionExist(positionID *big.Int) (bool, error) {
	return f.Contract.DoesPositionExist(&f.CallOpts, positionID)
}

// EpochFees is a free data retrieval call binding the contract method 0xfe944a57.
//
// Solidity: function epochFees(address ) view returns(uint256)
func (f *FuturesCaller) EpochFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "epochFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochFees is a free data retrieval call binding the contract method 0xfe944a57.
//
// Solidity: function epochFees(address ) view returns(uint256)
func (f *FuturesSession) EpochFees(arg0 common.Address) (*big.Int, error) {
	return f.Contract.EpochFees(&f.CallOpts, arg0)
}

// EpochFees is a free data retrieval call binding the contract method 0xfe944a57.
//
// Solidity: function epochFees(address ) view returns(uint256)
func (f *FuturesCallerSession) EpochFees(arg0 common.Address) (*big.Int, error) {
	return f.Contract.EpochFees(&f.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (f *FuturesCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "filled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// FuturesMarkets is a free data retrieval call binding the contract method 0x8192ea00.
//
// Solidity: function futuresMarkets(bytes32 ) view returns(bytes32 marketID, address baseCurrency, string ticker, address oracle, (uint256) initialMarginRatio, (uint256) maintenanceMarginRatio, uint256 indexPrice, uint256 expiryOrNextFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding, ((uint256),(uint256),(uint256)) transactionFees)
func (f *FuturesCaller) FuturesMarkets(opts *bind.CallOpts, arg0 [32]byte) (struct {
	MarketID                     [32]byte
	BaseCurrency                 common.Address
	Ticker                       string
	Oracle                       common.Address
	InitialMarginRatio           PermyriadMathPermyriad
	MaintenanceMarginRatio       PermyriadMathPermyriad
	IndexPrice                   *big.Int
	ExpiryOrNextFundingTimestamp *big.Int
	FundingInterval              *big.Int
	CumulativeFunding            *big.Int
	TransactionFees              TypesTransactionFees
}, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "futuresMarkets", arg0)

	outstruct := new(struct {
		MarketID                     [32]byte
		BaseCurrency                 common.Address
		Ticker                       string
		Oracle                       common.Address
		InitialMarginRatio           PermyriadMathPermyriad
		MaintenanceMarginRatio       PermyriadMathPermyriad
		IndexPrice                   *big.Int
		ExpiryOrNextFundingTimestamp *big.Int
		FundingInterval              *big.Int
		CumulativeFunding            *big.Int
		TransactionFees              TypesTransactionFees
	})

	outstruct.MarketID = out[0].([32]byte)
	outstruct.BaseCurrency = out[1].(common.Address)
	outstruct.Ticker = out[2].(string)
	outstruct.Oracle = out[3].(common.Address)
	outstruct.InitialMarginRatio = out[4].(PermyriadMathPermyriad)
	outstruct.MaintenanceMarginRatio = out[5].(PermyriadMathPermyriad)
	outstruct.IndexPrice = out[6].(*big.Int)
	outstruct.ExpiryOrNextFundingTimestamp = out[7].(*big.Int)
	outstruct.FundingInterval = out[8].(*big.Int)
	outstruct.CumulativeFunding = out[9].(*big.Int)
	outstruct.TransactionFees = out[10].(TypesTransactionFees)

	return *outstruct, err

}

// FuturesMarkets is a free data retrieval call binding the contract method 0x8192ea00.
//
// Solidity: function futuresMarkets(bytes32 ) view returns(bytes32 marketID, address baseCurrency, string ticker, address oracle, (uint256) initialMarginRatio, (uint256) maintenanceMarginRatio, uint256 indexPrice, uint256 expiryOrNextFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding, ((uint256),(uint256),(uint256)) transactionFees)
func (f *FuturesSession) FuturesMarkets(arg0 [32]byte) (struct {
	MarketID                     [32]byte
	BaseCurrency                 common.Address
	Ticker                       string
	Oracle                       common.Address
	InitialMarginRatio           PermyriadMathPermyriad
	MaintenanceMarginRatio       PermyriadMathPermyriad
	IndexPrice                   *big.Int
	ExpiryOrNextFundingTimestamp *big.Int
	FundingInterval              *big.Int
	CumulativeFunding            *big.Int
	TransactionFees              TypesTransactionFees
}, error) {
	return f.Contract.FuturesMarkets(&f.CallOpts, arg0)
}

// FuturesMarkets is a free data retrieval call binding the contract method 0x8192ea00.
//
// Solidity: function futuresMarkets(bytes32 ) view returns(bytes32 marketID, address baseCurrency, string ticker, address oracle, (uint256) initialMarginRatio, (uint256) maintenanceMarginRatio, uint256 indexPrice, uint256 expiryOrNextFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding, ((uint256),(uint256),(uint256)) transactionFees)
func (f *FuturesCallerSession) FuturesMarkets(arg0 [32]byte) (struct {
	MarketID                     [32]byte
	BaseCurrency                 common.Address
	Ticker                       string
	Oracle                       common.Address
	InitialMarginRatio           PermyriadMathPermyriad
	MaintenanceMarginRatio       PermyriadMathPermyriad
	IndexPrice                   *big.Int
	ExpiryOrNextFundingTimestamp *big.Int
	FundingInterval              *big.Int
	CumulativeFunding            *big.Int
	TransactionFees              TypesTransactionFees
}, error) {
	return f.Contract.FuturesMarkets(&f.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 positionID) view returns(address)
func (f *FuturesCaller) GetApproved(opts *bind.CallOpts, positionID *big.Int) (common.Address, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getApproved", positionID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 positionID) view returns(address)
func (f *FuturesSession) GetApproved(positionID *big.Int) (common.Address, error) {
	return f.Contract.GetApproved(&f.CallOpts, positionID)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 positionID) view returns(address)
func (f *FuturesCallerSession) GetApproved(positionID *big.Int) (common.Address, error) {
	return f.Contract.GetApproved(&f.CallOpts, positionID)
}

// GetDefaultSubAccountDeposits is a free data retrieval call binding the contract method 0x1883e458.
//
// Solidity: function getDefaultSubAccountDeposits(address baseCurrency, address trader) view returns(uint256)
func (f *FuturesCaller) GetDefaultSubAccountDeposits(opts *bind.CallOpts, baseCurrency common.Address, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getDefaultSubAccountDeposits", baseCurrency, trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "getDefaultSubAccountIdForTrader", trader)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

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

// GetMaxFundingFeeAbs is a free data retrieval call binding the contract method 0xa9d56fe9.
//
// Solidity: function getMaxFundingFeeAbs(bytes32 marketID) view returns(uint256)
func (f *FuturesCaller) GetMaxFundingFeeAbs(opts *bind.CallOpts, marketID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getMaxFundingFeeAbs", marketID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxFundingFeeAbs is a free data retrieval call binding the contract method 0xa9d56fe9.
//
// Solidity: function getMaxFundingFeeAbs(bytes32 marketID) view returns(uint256)
func (f *FuturesSession) GetMaxFundingFeeAbs(marketID [32]byte) (*big.Int, error) {
	return f.Contract.GetMaxFundingFeeAbs(&f.CallOpts, marketID)
}

// GetMaxFundingFeeAbs is a free data retrieval call binding the contract method 0xa9d56fe9.
//
// Solidity: function getMaxFundingFeeAbs(bytes32 marketID) view returns(uint256)
func (f *FuturesCallerSession) GetMaxFundingFeeAbs(marketID [32]byte) (*big.Int, error) {
	return f.Contract.GetMaxFundingFeeAbs(&f.CallOpts, marketID)
}

// GetOrderInfo is a free data retrieval call binding the contract method 0x9d3fa4b9.
//
// Solidity: function getOrderInfo((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) view returns((uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256) orderInfo)
func (f *FuturesCaller) GetOrderInfo(opts *bind.CallOpts, order Order) (OrderInfo, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getOrderInfo", order)

	if err != nil {
		return *new(OrderInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderInfo)).(*OrderInfo)

	return out0, err

}

// GetOrderInfo is a free data retrieval call binding the contract method 0x9d3fa4b9.
//
// Solidity: function getOrderInfo((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) view returns((uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256) orderInfo)
func (f *FuturesSession) GetOrderInfo(order Order) (OrderInfo, error) {
	return f.Contract.GetOrderInfo(&f.CallOpts, order)
}

// GetOrderInfo is a free data retrieval call binding the contract method 0x9d3fa4b9.
//
// Solidity: function getOrderInfo((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) view returns((uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256) orderInfo)
func (f *FuturesCallerSession) GetOrderInfo(order Order) (OrderInfo, error) {
	return f.Contract.GetOrderInfo(&f.CallOpts, order)
}

// GetOrderRelevantState is a free data retrieval call binding the contract method 0xef3a29b3.
//
// Solidity: function getOrderRelevantState((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature, uint256 indexPrice) view returns((uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256) orderInfo, uint256 fillableTakerAssetAmount, bool isValidSignature)
func (f *FuturesCaller) GetOrderRelevantState(opts *bind.CallOpts, order Order, signature []byte, indexPrice *big.Int) (struct {
	OrderInfo                DerivativeOrderInfo
	FillableTakerAssetAmount *big.Int
	IsValidSignature         bool
}, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getOrderRelevantState", order, signature, indexPrice)

	outstruct := new(struct {
		OrderInfo                DerivativeOrderInfo
		FillableTakerAssetAmount *big.Int
		IsValidSignature         bool
	})

	outstruct.OrderInfo = out[0].(DerivativeOrderInfo)
	outstruct.FillableTakerAssetAmount = out[1].(*big.Int)
	outstruct.IsValidSignature = out[2].(bool)

	return *outstruct, err

}

// GetOrderRelevantState is a free data retrieval call binding the contract method 0xef3a29b3.
//
// Solidity: function getOrderRelevantState((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature, uint256 indexPrice) view returns((uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256) orderInfo, uint256 fillableTakerAssetAmount, bool isValidSignature)
func (f *FuturesSession) GetOrderRelevantState(order Order, signature []byte, indexPrice *big.Int) (struct {
	OrderInfo                DerivativeOrderInfo
	FillableTakerAssetAmount *big.Int
	IsValidSignature         bool
}, error) {
	return f.Contract.GetOrderRelevantState(&f.CallOpts, order, signature, indexPrice)
}

// GetOrderRelevantState is a free data retrieval call binding the contract method 0xef3a29b3.
//
// Solidity: function getOrderRelevantState((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature, uint256 indexPrice) view returns((uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256) orderInfo, uint256 fillableTakerAssetAmount, bool isValidSignature)
func (f *FuturesCallerSession) GetOrderRelevantState(order Order, signature []byte, indexPrice *big.Int) (struct {
	OrderInfo                DerivativeOrderInfo
	FillableTakerAssetAmount *big.Int
	IsValidSignature         bool
}, error) {
	return f.Contract.GetOrderRelevantState(&f.CallOpts, order, signature, indexPrice)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, bytes[] signatures) view returns((uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256)[] ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (f *FuturesCaller) GetOrderRelevantStates(opts *bind.CallOpts, orders []Order, signatures [][]byte) (struct {
	OrdersInfo                []DerivativeOrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getOrderRelevantStates", orders, signatures)

	outstruct := new(struct {
		OrdersInfo                []DerivativeOrderInfo
		FillableTakerAssetAmounts []*big.Int
		IsValidSignature          []bool
	})

	outstruct.OrdersInfo = out[0].([]DerivativeOrderInfo)
	outstruct.FillableTakerAssetAmounts = out[1].([]*big.Int)
	outstruct.IsValidSignature = out[2].([]bool)

	return *outstruct, err

}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, bytes[] signatures) view returns((uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256)[] ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (f *FuturesSession) GetOrderRelevantStates(orders []Order, signatures [][]byte) (struct {
	OrdersInfo                []DerivativeOrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	return f.Contract.GetOrderRelevantStates(&f.CallOpts, orders, signatures)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, bytes[] signatures) view returns((uint8,bytes32,uint256,bytes32,uint8,bytes32,uint256)[] ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
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
	var out []interface{}
	err := f.contract.Call(opts, &out, "getPositionIDsForTrader", trader, marketID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

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
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,uint256,int256)[])
func (f *FuturesCaller) GetPositionsForTrader(opts *bind.CallOpts, trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getPositionsForTrader", trader, marketID)

	if err != nil {
		return *new([]TypesPosition), err
	}

	out0 := *abi.ConvertType(out[0], new([]TypesPosition)).(*[]TypesPosition)

	return out0, err

}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,uint256,int256)[])
func (f *FuturesSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return f.Contract.GetPositionsForTrader(&f.CallOpts, trader, marketID)
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,uint256,int256)[])
func (f *FuturesCallerSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return f.Contract.GetPositionsForTrader(&f.CallOpts, trader, marketID)
}

// GetReceiptApproved is a free data retrieval call binding the contract method 0x228906c8.
//
// Solidity: function getReceiptApproved(bytes32 receiverSubAccountID, uint256 positionID) view returns(bool)
func (f *FuturesCaller) GetReceiptApproved(opts *bind.CallOpts, receiverSubAccountID [32]byte, positionID *big.Int) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getReceiptApproved", receiverSubAccountID, positionID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetReceiptApproved is a free data retrieval call binding the contract method 0x228906c8.
//
// Solidity: function getReceiptApproved(bytes32 receiverSubAccountID, uint256 positionID) view returns(bool)
func (f *FuturesSession) GetReceiptApproved(receiverSubAccountID [32]byte, positionID *big.Int) (bool, error) {
	return f.Contract.GetReceiptApproved(&f.CallOpts, receiverSubAccountID, positionID)
}

// GetReceiptApproved is a free data retrieval call binding the contract method 0x228906c8.
//
// Solidity: function getReceiptApproved(bytes32 receiverSubAccountID, uint256 positionID) view returns(bool)
func (f *FuturesCallerSession) GetReceiptApproved(receiverSubAccountID [32]byte, positionID *big.Int) (bool, error) {
	return f.Contract.GetReceiptApproved(&f.CallOpts, receiverSubAccountID, positionID)
}

// GetTraderSubAccountsCount is a free data retrieval call binding the contract method 0x603ca5dc.
//
// Solidity: function getTraderSubAccountsCount(address trader) view returns(uint256)
func (f *FuturesCaller) GetTraderSubAccountsCount(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getTraderSubAccountsCount", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
// Solidity: function getUnitPositionValue(uint256 positionID, uint256 priceToUse) view returns(int256 unitPositionValue)
func (f *FuturesCaller) GetUnitPositionValue(opts *bind.CallOpts, positionID *big.Int, priceToUse *big.Int) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "getUnitPositionValue", positionID, priceToUse)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnitPositionValue is a free data retrieval call binding the contract method 0x574e2080.
//
// Solidity: function getUnitPositionValue(uint256 positionID, uint256 priceToUse) view returns(int256 unitPositionValue)
func (f *FuturesSession) GetUnitPositionValue(positionID *big.Int, priceToUse *big.Int) (*big.Int, error) {
	return f.Contract.GetUnitPositionValue(&f.CallOpts, positionID, priceToUse)
}

// GetUnitPositionValue is a free data retrieval call binding the contract method 0x574e2080.
//
// Solidity: function getUnitPositionValue(uint256 positionID, uint256 priceToUse) view returns(int256 unitPositionValue)
func (f *FuturesCallerSession) GetUnitPositionValue(positionID *big.Int, priceToUse *big.Int) (*big.Int, error) {
	return f.Contract.GetUnitPositionValue(&f.CallOpts, positionID, priceToUse)
}

// InsurancePools is a free data retrieval call binding the contract method 0x2514c1f1.
//
// Solidity: function insurancePools(bytes32 ) view returns(uint256)
func (f *FuturesCaller) InsurancePools(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "insurancePools", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// IsAllowedToReceivePosition is a free data retrieval call binding the contract method 0x57971430.
//
// Solidity: function isAllowedToReceivePosition(uint256 positionID, address sender, bytes32 receiverSubAccountID) view returns(bool)
func (f *FuturesCaller) IsAllowedToReceivePosition(opts *bind.CallOpts, positionID *big.Int, sender common.Address, receiverSubAccountID [32]byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isAllowedToReceivePosition", positionID, sender, receiverSubAccountID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToReceivePosition is a free data retrieval call binding the contract method 0x57971430.
//
// Solidity: function isAllowedToReceivePosition(uint256 positionID, address sender, bytes32 receiverSubAccountID) view returns(bool)
func (f *FuturesSession) IsAllowedToReceivePosition(positionID *big.Int, sender common.Address, receiverSubAccountID [32]byte) (bool, error) {
	return f.Contract.IsAllowedToReceivePosition(&f.CallOpts, positionID, sender, receiverSubAccountID)
}

// IsAllowedToReceivePosition is a free data retrieval call binding the contract method 0x57971430.
//
// Solidity: function isAllowedToReceivePosition(uint256 positionID, address sender, bytes32 receiverSubAccountID) view returns(bool)
func (f *FuturesCallerSession) IsAllowedToReceivePosition(positionID *big.Int, sender common.Address, receiverSubAccountID [32]byte) (bool, error) {
	return f.Contract.IsAllowedToReceivePosition(&f.CallOpts, positionID, sender, receiverSubAccountID)
}

// IsAllowedToTransferPosition is a free data retrieval call binding the contract method 0x0e6c0912.
//
// Solidity: function isAllowedToTransferPosition(uint256 positionID, address sender) view returns(bool)
func (f *FuturesCaller) IsAllowedToTransferPosition(opts *bind.CallOpts, positionID *big.Int, sender common.Address) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isAllowedToTransferPosition", positionID, sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToTransferPosition is a free data retrieval call binding the contract method 0x0e6c0912.
//
// Solidity: function isAllowedToTransferPosition(uint256 positionID, address sender) view returns(bool)
func (f *FuturesSession) IsAllowedToTransferPosition(positionID *big.Int, sender common.Address) (bool, error) {
	return f.Contract.IsAllowedToTransferPosition(&f.CallOpts, positionID, sender)
}

// IsAllowedToTransferPosition is a free data retrieval call binding the contract method 0x0e6c0912.
//
// Solidity: function isAllowedToTransferPosition(uint256 positionID, address sender) view returns(bool)
func (f *FuturesCallerSession) IsAllowedToTransferPosition(positionID *big.Int, sender common.Address) (bool, error) {
	return f.Contract.IsAllowedToTransferPosition(&f.CallOpts, positionID, sender)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xb0698628.
//
// Solidity: function isApprovedForAll(bytes32 fromSubAccountID, address operator) view returns(bool)
func (f *FuturesCaller) IsApprovedForAll(opts *bind.CallOpts, fromSubAccountID [32]byte, operator common.Address) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isApprovedForAll", fromSubAccountID, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xb0698628.
//
// Solidity: function isApprovedForAll(bytes32 fromSubAccountID, address operator) view returns(bool)
func (f *FuturesSession) IsApprovedForAll(fromSubAccountID [32]byte, operator common.Address) (bool, error) {
	return f.Contract.IsApprovedForAll(&f.CallOpts, fromSubAccountID, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xb0698628.
//
// Solidity: function isApprovedForAll(bytes32 fromSubAccountID, address operator) view returns(bool)
func (f *FuturesCallerSession) IsApprovedForAll(fromSubAccountID [32]byte, operator common.Address) (bool, error) {
	return f.Contract.IsApprovedForAll(&f.CallOpts, fromSubAccountID, operator)
}

// IsApprovedForMarket is a free data retrieval call binding the contract method 0x1d37b559.
//
// Solidity: function isApprovedForMarket(bytes32 fromSubAccountID, address operator, bytes32 marketID) view returns(bool)
func (f *FuturesCaller) IsApprovedForMarket(opts *bind.CallOpts, fromSubAccountID [32]byte, operator common.Address, marketID [32]byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isApprovedForMarket", fromSubAccountID, operator, marketID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForMarket is a free data retrieval call binding the contract method 0x1d37b559.
//
// Solidity: function isApprovedForMarket(bytes32 fromSubAccountID, address operator, bytes32 marketID) view returns(bool)
func (f *FuturesSession) IsApprovedForMarket(fromSubAccountID [32]byte, operator common.Address, marketID [32]byte) (bool, error) {
	return f.Contract.IsApprovedForMarket(&f.CallOpts, fromSubAccountID, operator, marketID)
}

// IsApprovedForMarket is a free data retrieval call binding the contract method 0x1d37b559.
//
// Solidity: function isApprovedForMarket(bytes32 fromSubAccountID, address operator, bytes32 marketID) view returns(bool)
func (f *FuturesCallerSession) IsApprovedForMarket(fromSubAccountID [32]byte, operator common.Address, marketID [32]byte) (bool, error) {
	return f.Contract.IsApprovedForMarket(&f.CallOpts, fromSubAccountID, operator, marketID)
}

// IsFuturesMarketSettled is a free data retrieval call binding the contract method 0xd5a5c5e2.
//
// Solidity: function isFuturesMarketSettled(bytes32 ) view returns(bool)
func (f *FuturesCaller) IsFuturesMarketSettled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isFuturesMarketSettled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFuturesMarketSettled is a free data retrieval call binding the contract method 0xd5a5c5e2.
//
// Solidity: function isFuturesMarketSettled(bytes32 ) view returns(bool)
func (f *FuturesSession) IsFuturesMarketSettled(arg0 [32]byte) (bool, error) {
	return f.Contract.IsFuturesMarketSettled(&f.CallOpts, arg0)
}

// IsFuturesMarketSettled is a free data retrieval call binding the contract method 0xd5a5c5e2.
//
// Solidity: function isFuturesMarketSettled(bytes32 ) view returns(bool)
func (f *FuturesCallerSession) IsFuturesMarketSettled(arg0 [32]byte) (bool, error) {
	return f.Contract.IsFuturesMarketSettled(&f.CallOpts, arg0)
}

// IsReceiptApprovedForAll is a free data retrieval call binding the contract method 0x2d1792bd.
//
// Solidity: function isReceiptApprovedForAll(bytes32 receiverSubAccountID, address operator) view returns(bool)
func (f *FuturesCaller) IsReceiptApprovedForAll(opts *bind.CallOpts, receiverSubAccountID [32]byte, operator common.Address) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isReceiptApprovedForAll", receiverSubAccountID, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReceiptApprovedForAll is a free data retrieval call binding the contract method 0x2d1792bd.
//
// Solidity: function isReceiptApprovedForAll(bytes32 receiverSubAccountID, address operator) view returns(bool)
func (f *FuturesSession) IsReceiptApprovedForAll(receiverSubAccountID [32]byte, operator common.Address) (bool, error) {
	return f.Contract.IsReceiptApprovedForAll(&f.CallOpts, receiverSubAccountID, operator)
}

// IsReceiptApprovedForAll is a free data retrieval call binding the contract method 0x2d1792bd.
//
// Solidity: function isReceiptApprovedForAll(bytes32 receiverSubAccountID, address operator) view returns(bool)
func (f *FuturesCallerSession) IsReceiptApprovedForAll(receiverSubAccountID [32]byte, operator common.Address) (bool, error) {
	return f.Contract.IsReceiptApprovedForAll(&f.CallOpts, receiverSubAccountID, operator)
}

// IsReceiptApprovedForMarket is a free data retrieval call binding the contract method 0x218c106c.
//
// Solidity: function isReceiptApprovedForMarket(bytes32 receiverSubAccountID, address operator, bytes32 marketID) view returns(bool)
func (f *FuturesCaller) IsReceiptApprovedForMarket(opts *bind.CallOpts, receiverSubAccountID [32]byte, operator common.Address, marketID [32]byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isReceiptApprovedForMarket", receiverSubAccountID, operator, marketID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReceiptApprovedForMarket is a free data retrieval call binding the contract method 0x218c106c.
//
// Solidity: function isReceiptApprovedForMarket(bytes32 receiverSubAccountID, address operator, bytes32 marketID) view returns(bool)
func (f *FuturesSession) IsReceiptApprovedForMarket(receiverSubAccountID [32]byte, operator common.Address, marketID [32]byte) (bool, error) {
	return f.Contract.IsReceiptApprovedForMarket(&f.CallOpts, receiverSubAccountID, operator, marketID)
}

// IsReceiptApprovedForMarket is a free data retrieval call binding the contract method 0x218c106c.
//
// Solidity: function isReceiptApprovedForMarket(bytes32 receiverSubAccountID, address operator, bytes32 marketID) view returns(bool)
func (f *FuturesCallerSession) IsReceiptApprovedForMarket(receiverSubAccountID [32]byte, operator common.Address, marketID [32]byte) (bool, error) {
	return f.Contract.IsReceiptApprovedForMarket(&f.CallOpts, receiverSubAccountID, operator, marketID)
}

// IsTimeExpiryFuturesMarket is a free data retrieval call binding the contract method 0x6e664084.
//
// Solidity: function isTimeExpiryFuturesMarket(bytes32 marketID) view returns(bool)
func (f *FuturesCaller) IsTimeExpiryFuturesMarket(opts *bind.CallOpts, marketID [32]byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isTimeExpiryFuturesMarket", marketID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimeExpiryFuturesMarket is a free data retrieval call binding the contract method 0x6e664084.
//
// Solidity: function isTimeExpiryFuturesMarket(bytes32 marketID) view returns(bool)
func (f *FuturesSession) IsTimeExpiryFuturesMarket(marketID [32]byte) (bool, error) {
	return f.Contract.IsTimeExpiryFuturesMarket(&f.CallOpts, marketID)
}

// IsTimeExpiryFuturesMarket is a free data retrieval call binding the contract method 0x6e664084.
//
// Solidity: function isTimeExpiryFuturesMarket(bytes32 marketID) view returns(bool)
func (f *FuturesCallerSession) IsTimeExpiryFuturesMarket(marketID [32]byte) (bool, error) {
	return f.Contract.IsTimeExpiryFuturesMarket(&f.CallOpts, marketID)
}

// IsValidBaseCurrency is a free data retrieval call binding the contract method 0x227cdb85.
//
// Solidity: function isValidBaseCurrency(address ) view returns(bool)
func (f *FuturesCaller) IsValidBaseCurrency(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isValidBaseCurrency", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// IsValidHashSignature is a free data retrieval call binding the contract method 0x8171c407.
//
// Solidity: function isValidHashSignature(bytes32 hash, address signerAddress, bytes signature) view returns(bool isValid)
func (f *FuturesCaller) IsValidHashSignature(opts *bind.CallOpts, hash [32]byte, signerAddress common.Address, signature []byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isValidHashSignature", hash, signerAddress, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidHashSignature is a free data retrieval call binding the contract method 0x8171c407.
//
// Solidity: function isValidHashSignature(bytes32 hash, address signerAddress, bytes signature) view returns(bool isValid)
func (f *FuturesSession) IsValidHashSignature(hash [32]byte, signerAddress common.Address, signature []byte) (bool, error) {
	return f.Contract.IsValidHashSignature(&f.CallOpts, hash, signerAddress, signature)
}

// IsValidHashSignature is a free data retrieval call binding the contract method 0x8171c407.
//
// Solidity: function isValidHashSignature(bytes32 hash, address signerAddress, bytes signature) view returns(bool isValid)
func (f *FuturesCallerSession) IsValidHashSignature(hash [32]byte, signerAddress common.Address, signature []byte) (bool, error) {
	return f.Contract.IsValidHashSignature(&f.CallOpts, hash, signerAddress, signature)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature) view returns(bool isValid)
func (f *FuturesCaller) IsValidOrderSignature(opts *bind.CallOpts, order Order, signature []byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isValidOrderSignature", order, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// IsValidTransactionSignature is a free data retrieval call binding the contract method 0x8d45cd23.
//
// Solidity: function isValidTransactionSignature((uint256,uint256,uint256,address,bytes) transaction, bytes signature) view returns(bool isValid)
func (f *FuturesCaller) IsValidTransactionSignature(opts *bind.CallOpts, transaction LibZeroExTransactionZeroExTransaction, signature []byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "isValidTransactionSignature", transaction, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidTransactionSignature is a free data retrieval call binding the contract method 0x8d45cd23.
//
// Solidity: function isValidTransactionSignature((uint256,uint256,uint256,address,bytes) transaction, bytes signature) view returns(bool isValid)
func (f *FuturesSession) IsValidTransactionSignature(transaction LibZeroExTransactionZeroExTransaction, signature []byte) (bool, error) {
	return f.Contract.IsValidTransactionSignature(&f.CallOpts, transaction, signature)
}

// IsValidTransactionSignature is a free data retrieval call binding the contract method 0x8d45cd23.
//
// Solidity: function isValidTransactionSignature((uint256,uint256,uint256,address,bytes) transaction, bytes signature) view returns(bool isValid)
func (f *FuturesCallerSession) IsValidTransactionSignature(transaction LibZeroExTransactionZeroExTransaction, signature []byte) (bool, error) {
	return f.Contract.IsValidTransactionSignature(&f.CallOpts, transaction, signature)
}

// LastValidVWAP is a free data retrieval call binding the contract method 0x42151440.
//
// Solidity: function lastValidVWAP(bytes32 ) view returns(uint256)
func (f *FuturesCaller) LastValidVWAP(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "lastValidVWAP", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastValidVWAP is a free data retrieval call binding the contract method 0x42151440.
//
// Solidity: function lastValidVWAP(bytes32 ) view returns(uint256)
func (f *FuturesSession) LastValidVWAP(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.LastValidVWAP(&f.CallOpts, arg0)
}

// LastValidVWAP is a free data retrieval call binding the contract method 0x42151440.
//
// Solidity: function lastValidVWAP(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) LastValidVWAP(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.LastValidVWAP(&f.CallOpts, arg0)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (f *FuturesCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "marketCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "marketSerialToID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

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

// MaxBaseCurrencyCap is a free data retrieval call binding the contract method 0xf8720a75.
//
// Solidity: function maxBaseCurrencyCap() view returns(uint256)
func (f *FuturesCaller) MaxBaseCurrencyCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "maxBaseCurrencyCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBaseCurrencyCap is a free data retrieval call binding the contract method 0xf8720a75.
//
// Solidity: function maxBaseCurrencyCap() view returns(uint256)
func (f *FuturesSession) MaxBaseCurrencyCap() (*big.Int, error) {
	return f.Contract.MaxBaseCurrencyCap(&f.CallOpts)
}

// MaxBaseCurrencyCap is a free data retrieval call binding the contract method 0xf8720a75.
//
// Solidity: function maxBaseCurrencyCap() view returns(uint256)
func (f *FuturesCallerSession) MaxBaseCurrencyCap() (*big.Int, error) {
	return f.Contract.MaxBaseCurrencyCap(&f.CallOpts)
}

// MostRecentEpochQuantity is a free data retrieval call binding the contract method 0xf7a28a1a.
//
// Solidity: function mostRecentEpochQuantity(bytes32 ) view returns(uint256)
func (f *FuturesCaller) MostRecentEpochQuantity(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "mostRecentEpochQuantity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// MostRecentEpochVolume is a free data retrieval call binding the contract method 0x13626422.
//
// Solidity: function mostRecentEpochVolume(bytes32 ) view returns(uint256)
func (f *FuturesCaller) MostRecentEpochVolume(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "mostRecentEpochVolume", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// MostRecentEpochWeightedAverageContractIndexDiff is a free data retrieval call binding the contract method 0x32d22e7d.
//
// Solidity: function mostRecentEpochWeightedAverageContractIndexDiff(bytes32 ) view returns(int256)
func (f *FuturesCaller) MostRecentEpochWeightedAverageContractIndexDiff(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "mostRecentEpochWeightedAverageContractIndexDiff", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MostRecentEpochWeightedAverageContractIndexDiff is a free data retrieval call binding the contract method 0x32d22e7d.
//
// Solidity: function mostRecentEpochWeightedAverageContractIndexDiff(bytes32 ) view returns(int256)
func (f *FuturesSession) MostRecentEpochWeightedAverageContractIndexDiff(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentEpochWeightedAverageContractIndexDiff(&f.CallOpts, arg0)
}

// MostRecentEpochWeightedAverageContractIndexDiff is a free data retrieval call binding the contract method 0x32d22e7d.
//
// Solidity: function mostRecentEpochWeightedAverageContractIndexDiff(bytes32 ) view returns(int256)
func (f *FuturesCallerSession) MostRecentEpochWeightedAverageContractIndexDiff(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentEpochWeightedAverageContractIndexDiff(&f.CallOpts, arg0)
}

// MostRecentmostRecentEpochVolumeEpochQuantity is a free data retrieval call binding the contract method 0x45194767.
//
// Solidity: function mostRecentmostRecentEpochVolumeEpochQuantity(bytes32 ) view returns(uint256)
func (f *FuturesCaller) MostRecentmostRecentEpochVolumeEpochQuantity(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "mostRecentmostRecentEpochVolumeEpochQuantity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MostRecentmostRecentEpochVolumeEpochQuantity is a free data retrieval call binding the contract method 0x45194767.
//
// Solidity: function mostRecentmostRecentEpochVolumeEpochQuantity(bytes32 ) view returns(uint256)
func (f *FuturesSession) MostRecentmostRecentEpochVolumeEpochQuantity(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentmostRecentEpochVolumeEpochQuantity(&f.CallOpts, arg0)
}

// MostRecentmostRecentEpochVolumeEpochQuantity is a free data retrieval call binding the contract method 0x45194767.
//
// Solidity: function mostRecentmostRecentEpochVolumeEpochQuantity(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) MostRecentmostRecentEpochVolumeEpochQuantity(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.MostRecentmostRecentEpochVolumeEpochQuantity(&f.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (f *FuturesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "positionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// PositionCountInTEF is a free data retrieval call binding the contract method 0x4ded3208.
//
// Solidity: function positionCountInTEF(bytes32 ) view returns(uint256)
func (f *FuturesCaller) PositionCountInTEF(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "positionCountInTEF", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PositionCountInTEF is a free data retrieval call binding the contract method 0x4ded3208.
//
// Solidity: function positionCountInTEF(bytes32 ) view returns(uint256)
func (f *FuturesSession) PositionCountInTEF(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.PositionCountInTEF(&f.CallOpts, arg0)
}

// PositionCountInTEF is a free data retrieval call binding the contract method 0x4ded3208.
//
// Solidity: function positionCountInTEF(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) PositionCountInTEF(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.PositionCountInTEF(&f.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 subAccountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 entryPrice, uint256 margin, int256 cumulativeFundingEntry)
func (f *FuturesCaller) Positions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SubAccountID           [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	EntryPrice             *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
}, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		SubAccountID           [32]byte
		MarketID               [32]byte
		Direction              uint8
		Quantity               *big.Int
		EntryPrice             *big.Int
		Margin                 *big.Int
		CumulativeFundingEntry *big.Int
	})

	outstruct.SubAccountID = out[0].([32]byte)
	outstruct.MarketID = out[1].([32]byte)
	outstruct.Direction = out[2].(uint8)
	outstruct.Quantity = out[3].(*big.Int)
	outstruct.EntryPrice = out[4].(*big.Int)
	outstruct.Margin = out[5].(*big.Int)
	outstruct.CumulativeFundingEntry = out[6].(*big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 subAccountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 entryPrice, uint256 margin, int256 cumulativeFundingEntry)
func (f *FuturesSession) Positions(arg0 *big.Int) (struct {
	SubAccountID           [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	EntryPrice             *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
}, error) {
	return f.Contract.Positions(&f.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 subAccountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 entryPrice, uint256 margin, int256 cumulativeFundingEntry)
func (f *FuturesCallerSession) Positions(arg0 *big.Int) (struct {
	SubAccountID           [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	EntryPrice             *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
}, error) {
	return f.Contract.Positions(&f.CallOpts, arg0)
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (f *FuturesCaller) PreSigned(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "preSigned", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "restrictedDeposits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// SpotMarkets is a free data retrieval call binding the contract method 0xf5aa035e.
//
// Solidity: function spotMarkets(bytes32 ) view returns(bytes32 marketID, address baseToken, address quoteToken, string ticker, ((uint256),(uint256),(uint256)) transactionFees)
func (f *FuturesCaller) SpotMarkets(opts *bind.CallOpts, arg0 [32]byte) (struct {
	MarketID        [32]byte
	BaseToken       common.Address
	QuoteToken      common.Address
	Ticker          string
	TransactionFees TypesTransactionFees
}, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "spotMarkets", arg0)

	outstruct := new(struct {
		MarketID        [32]byte
		BaseToken       common.Address
		QuoteToken      common.Address
		Ticker          string
		TransactionFees TypesTransactionFees
	})

	outstruct.MarketID = out[0].([32]byte)
	outstruct.BaseToken = out[1].(common.Address)
	outstruct.QuoteToken = out[2].(common.Address)
	outstruct.Ticker = out[3].(string)
	outstruct.TransactionFees = out[4].(TypesTransactionFees)

	return *outstruct, err

}

// SpotMarkets is a free data retrieval call binding the contract method 0xf5aa035e.
//
// Solidity: function spotMarkets(bytes32 ) view returns(bytes32 marketID, address baseToken, address quoteToken, string ticker, ((uint256),(uint256),(uint256)) transactionFees)
func (f *FuturesSession) SpotMarkets(arg0 [32]byte) (struct {
	MarketID        [32]byte
	BaseToken       common.Address
	QuoteToken      common.Address
	Ticker          string
	TransactionFees TypesTransactionFees
}, error) {
	return f.Contract.SpotMarkets(&f.CallOpts, arg0)
}

// SpotMarkets is a free data retrieval call binding the contract method 0xf5aa035e.
//
// Solidity: function spotMarkets(bytes32 ) view returns(bytes32 marketID, address baseToken, address quoteToken, string ticker, ((uint256),(uint256),(uint256)) transactionFees)
func (f *FuturesCallerSession) SpotMarkets(arg0 [32]byte) (struct {
	MarketID        [32]byte
	BaseToken       common.Address
	QuoteToken      common.Address
	Ticker          string
	TransactionFees TypesTransactionFees
}, error) {
	return f.Contract.SpotMarkets(&f.CallOpts, arg0)
}

// SubAccountDeposits is a free data retrieval call binding the contract method 0x666ffb9b.
//
// Solidity: function subAccountDeposits(bytes32 , address ) view returns(uint256)
func (f *FuturesCaller) SubAccountDeposits(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "subAccountDeposits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "subAccountIdToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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
	var out []interface{}
	err := f.contract.Call(opts, &out, "subAccountToMarketToPositionID", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// TransactionsExecuted is a free data retrieval call binding the contract method 0x0228e168.
//
// Solidity: function transactionsExecuted(bytes32 ) view returns(bool)
func (f *FuturesCaller) TransactionsExecuted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := f.contract.Call(opts, &out, "transactionsExecuted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransactionsExecuted is a free data retrieval call binding the contract method 0x0228e168.
//
// Solidity: function transactionsExecuted(bytes32 ) view returns(bool)
func (f *FuturesSession) TransactionsExecuted(arg0 [32]byte) (bool, error) {
	return f.Contract.TransactionsExecuted(&f.CallOpts, arg0)
}

// TransactionsExecuted is a free data retrieval call binding the contract method 0x0228e168.
//
// Solidity: function transactionsExecuted(bytes32 ) view returns(bool)
func (f *FuturesCallerSession) TransactionsExecuted(arg0 [32]byte) (bool, error) {
	return f.Contract.TransactionsExecuted(&f.CallOpts, arg0)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address approveTo, uint256 positionID) returns()
func (f *FuturesTransactor) Approve(opts *bind.TransactOpts, approveTo common.Address, positionID *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "approve", approveTo, positionID)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address approveTo, uint256 positionID) returns()
func (f *FuturesSession) Approve(approveTo common.Address, positionID *big.Int) (*types.Transaction, error) {
	return f.Contract.Approve(&f.TransactOpts, approveTo, positionID)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address approveTo, uint256 positionID) returns()
func (f *FuturesTransactorSession) Approve(approveTo common.Address, positionID *big.Int) (*types.Transaction, error) {
	return f.Contract.Approve(&f.TransactOpts, approveTo, positionID)
}

// ApproveForReceiving is a paid mutator transaction binding the contract method 0xf5a5f0b3.
//
// Solidity: function approveForReceiving(bytes32 receiverSubAccountID, uint256 positionID) returns()
func (f *FuturesTransactor) ApproveForReceiving(opts *bind.TransactOpts, receiverSubAccountID [32]byte, positionID *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "approveForReceiving", receiverSubAccountID, positionID)
}

// ApproveForReceiving is a paid mutator transaction binding the contract method 0xf5a5f0b3.
//
// Solidity: function approveForReceiving(bytes32 receiverSubAccountID, uint256 positionID) returns()
func (f *FuturesSession) ApproveForReceiving(receiverSubAccountID [32]byte, positionID *big.Int) (*types.Transaction, error) {
	return f.Contract.ApproveForReceiving(&f.TransactOpts, receiverSubAccountID, positionID)
}

// ApproveForReceiving is a paid mutator transaction binding the contract method 0xf5a5f0b3.
//
// Solidity: function approveForReceiving(bytes32 receiverSubAccountID, uint256 positionID) returns()
func (f *FuturesTransactorSession) ApproveForReceiving(receiverSubAccountID [32]byte, positionID *big.Int) (*types.Transaction, error) {
	return f.Contract.ApproveForReceiving(&f.TransactOpts, receiverSubAccountID, positionID)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xdedfc1f1.
//
// Solidity: function batchCancelOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders) returns()
func (f *FuturesTransactor) BatchCancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchCancelOrders", orders)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xdedfc1f1.
//
// Solidity: function batchCancelOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders) returns()
func (f *FuturesSession) BatchCancelOrders(orders []Order) (*types.Transaction, error) {
	return f.Contract.BatchCancelOrders(&f.TransactOpts, orders)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xdedfc1f1.
//
// Solidity: function batchCancelOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders) returns()
func (f *FuturesTransactorSession) BatchCancelOrders(orders []Order) (*types.Transaction, error) {
	return f.Contract.BatchCancelOrders(&f.TransactOpts, orders)
}

// BatchCheckFunding is a paid mutator transaction binding the contract method 0x2c7c2134.
//
// Solidity: function batchCheckFunding(bytes32[] marketIDs) returns(bool[])
func (f *FuturesTransactor) BatchCheckFunding(opts *bind.TransactOpts, marketIDs [][32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchCheckFunding", marketIDs)
}

// BatchCheckFunding is a paid mutator transaction binding the contract method 0x2c7c2134.
//
// Solidity: function batchCheckFunding(bytes32[] marketIDs) returns(bool[])
func (f *FuturesSession) BatchCheckFunding(marketIDs [][32]byte) (*types.Transaction, error) {
	return f.Contract.BatchCheckFunding(&f.TransactOpts, marketIDs)
}

// BatchCheckFunding is a paid mutator transaction binding the contract method 0x2c7c2134.
//
// Solidity: function batchCheckFunding(bytes32[] marketIDs) returns(bool[])
func (f *FuturesTransactorSession) BatchCheckFunding(marketIDs [][32]byte) (*types.Transaction, error) {
	return f.Contract.BatchCheckFunding(&f.TransactOpts, marketIDs)
}

// BatchDepositForSubAccounts is a paid mutator transaction binding the contract method 0xc5d3277b.
//
// Solidity: function batchDepositForSubAccounts(address[] baseCurrencies, bytes32[] subAccountIDs, uint256[] amounts) returns()
func (f *FuturesTransactor) BatchDepositForSubAccounts(opts *bind.TransactOpts, baseCurrencies []common.Address, subAccountIDs [][32]byte, amounts []*big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchDepositForSubAccounts", baseCurrencies, subAccountIDs, amounts)
}

// BatchDepositForSubAccounts is a paid mutator transaction binding the contract method 0xc5d3277b.
//
// Solidity: function batchDepositForSubAccounts(address[] baseCurrencies, bytes32[] subAccountIDs, uint256[] amounts) returns()
func (f *FuturesSession) BatchDepositForSubAccounts(baseCurrencies []common.Address, subAccountIDs [][32]byte, amounts []*big.Int) (*types.Transaction, error) {
	return f.Contract.BatchDepositForSubAccounts(&f.TransactOpts, baseCurrencies, subAccountIDs, amounts)
}

// BatchDepositForSubAccounts is a paid mutator transaction binding the contract method 0xc5d3277b.
//
// Solidity: function batchDepositForSubAccounts(address[] baseCurrencies, bytes32[] subAccountIDs, uint256[] amounts) returns()
func (f *FuturesTransactorSession) BatchDepositForSubAccounts(baseCurrencies []common.Address, subAccountIDs [][32]byte, amounts []*big.Int) (*types.Transaction, error) {
	return f.Contract.BatchDepositForSubAccounts(&f.TransactOpts, baseCurrencies, subAccountIDs, amounts)
}

// BatchDepositForTraders is a paid mutator transaction binding the contract method 0x487ebde1.
//
// Solidity: function batchDepositForTraders(address[] baseCurrencies, address[] traders, uint256[] amounts) returns()
func (f *FuturesTransactor) BatchDepositForTraders(opts *bind.TransactOpts, baseCurrencies []common.Address, traders []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchDepositForTraders", baseCurrencies, traders, amounts)
}

// BatchDepositForTraders is a paid mutator transaction binding the contract method 0x487ebde1.
//
// Solidity: function batchDepositForTraders(address[] baseCurrencies, address[] traders, uint256[] amounts) returns()
func (f *FuturesSession) BatchDepositForTraders(baseCurrencies []common.Address, traders []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return f.Contract.BatchDepositForTraders(&f.TransactOpts, baseCurrencies, traders, amounts)
}

// BatchDepositForTraders is a paid mutator transaction binding the contract method 0x487ebde1.
//
// Solidity: function batchDepositForTraders(address[] baseCurrencies, address[] traders, uint256[] amounts) returns()
func (f *FuturesTransactorSession) BatchDepositForTraders(baseCurrencies []common.Address, traders []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return f.Contract.BatchDepositForTraders(&f.TransactOpts, baseCurrencies, traders, amounts)
}

// BatchExecuteTransactions is a paid mutator transaction binding the contract method 0xfc74896d.
//
// Solidity: function batchExecuteTransactions((uint256,uint256,uint256,address,bytes)[] transactions, bytes[] signatures) payable returns(bytes[] returnData)
func (f *FuturesTransactor) BatchExecuteTransactions(opts *bind.TransactOpts, transactions []LibZeroExTransactionZeroExTransaction, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchExecuteTransactions", transactions, signatures)
}

// BatchExecuteTransactions is a paid mutator transaction binding the contract method 0xfc74896d.
//
// Solidity: function batchExecuteTransactions((uint256,uint256,uint256,address,bytes)[] transactions, bytes[] signatures) payable returns(bytes[] returnData)
func (f *FuturesSession) BatchExecuteTransactions(transactions []LibZeroExTransactionZeroExTransaction, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchExecuteTransactions(&f.TransactOpts, transactions, signatures)
}

// BatchExecuteTransactions is a paid mutator transaction binding the contract method 0xfc74896d.
//
// Solidity: function batchExecuteTransactions((uint256,uint256,uint256,address,bytes)[] transactions, bytes[] signatures) payable returns(bytes[] returnData)
func (f *FuturesTransactorSession) BatchExecuteTransactions(transactions []LibZeroExTransactionZeroExTransaction, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.BatchExecuteTransactions(&f.TransactOpts, transactions, signatures)
}

// BatchLiquidatePosition is a paid mutator transaction binding the contract method 0x1b934aea.
//
// Solidity: function batchLiquidatePosition(uint256[] positionIDs, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[][] orders, uint256[] quantities, bytes[][] signatures, address liquidationCaller) returns((uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) BatchLiquidatePosition(opts *bind.TransactOpts, positionIDs []*big.Int, orders [][]Order, quantities []*big.Int, signatures [][][]byte, liquidationCaller common.Address) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchLiquidatePosition", positionIDs, orders, quantities, signatures, liquidationCaller)
}

// BatchLiquidatePosition is a paid mutator transaction binding the contract method 0x1b934aea.
//
// Solidity: function batchLiquidatePosition(uint256[] positionIDs, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[][] orders, uint256[] quantities, bytes[][] signatures, address liquidationCaller) returns((uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) BatchLiquidatePosition(positionIDs []*big.Int, orders [][]Order, quantities []*big.Int, signatures [][][]byte, liquidationCaller common.Address) (*types.Transaction, error) {
	return f.Contract.BatchLiquidatePosition(&f.TransactOpts, positionIDs, orders, quantities, signatures, liquidationCaller)
}

// BatchLiquidatePosition is a paid mutator transaction binding the contract method 0x1b934aea.
//
// Solidity: function batchLiquidatePosition(uint256[] positionIDs, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[][] orders, uint256[] quantities, bytes[][] signatures, address liquidationCaller) returns((uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) BatchLiquidatePosition(positionIDs []*big.Int, orders [][]Order, quantities []*big.Int, signatures [][][]byte, liquidationCaller common.Address) (*types.Transaction, error) {
	return f.Contract.BatchLiquidatePosition(&f.TransactOpts, positionIDs, orders, quantities, signatures, liquidationCaller)
}

// BatchSettleExpiryFuturesPosition is a paid mutator transaction binding the contract method 0xb378df0b.
//
// Solidity: function batchSettleExpiryFuturesPosition(bytes32 marketID, bytes32[] subAccountIDs) returns((uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) BatchSettleExpiryFuturesPosition(opts *bind.TransactOpts, marketID [32]byte, subAccountIDs [][32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "batchSettleExpiryFuturesPosition", marketID, subAccountIDs)
}

// BatchSettleExpiryFuturesPosition is a paid mutator transaction binding the contract method 0xb378df0b.
//
// Solidity: function batchSettleExpiryFuturesPosition(bytes32 marketID, bytes32[] subAccountIDs) returns((uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) BatchSettleExpiryFuturesPosition(marketID [32]byte, subAccountIDs [][32]byte) (*types.Transaction, error) {
	return f.Contract.BatchSettleExpiryFuturesPosition(&f.TransactOpts, marketID, subAccountIDs)
}

// BatchSettleExpiryFuturesPosition is a paid mutator transaction binding the contract method 0xb378df0b.
//
// Solidity: function batchSettleExpiryFuturesPosition(bytes32 marketID, bytes32[] subAccountIDs) returns((uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) BatchSettleExpiryFuturesPosition(marketID [32]byte, subAccountIDs [][32]byte) (*types.Transaction, error) {
	return f.Contract.BatchSettleExpiryFuturesPosition(&f.TransactOpts, marketID, subAccountIDs)
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

// ChangeAuctionTimeInterval is a paid mutator transaction binding the contract method 0x5852ca61.
//
// Solidity: function changeAuctionTimeInterval(uint256 auctionTimeInterval_) returns()
func (f *FuturesTransactor) ChangeAuctionTimeInterval(opts *bind.TransactOpts, auctionTimeInterval_ *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeAuctionTimeInterval", auctionTimeInterval_)
}

// ChangeAuctionTimeInterval is a paid mutator transaction binding the contract method 0x5852ca61.
//
// Solidity: function changeAuctionTimeInterval(uint256 auctionTimeInterval_) returns()
func (f *FuturesSession) ChangeAuctionTimeInterval(auctionTimeInterval_ *big.Int) (*types.Transaction, error) {
	return f.Contract.ChangeAuctionTimeInterval(&f.TransactOpts, auctionTimeInterval_)
}

// ChangeAuctionTimeInterval is a paid mutator transaction binding the contract method 0x5852ca61.
//
// Solidity: function changeAuctionTimeInterval(uint256 auctionTimeInterval_) returns()
func (f *FuturesTransactorSession) ChangeAuctionTimeInterval(auctionTimeInterval_ *big.Int) (*types.Transaction, error) {
	return f.Contract.ChangeAuctionTimeInterval(&f.TransactOpts, auctionTimeInterval_)
}

// ChangeMarketFundingInterval is a paid mutator transaction binding the contract method 0x7a540e43.
//
// Solidity: function changeMarketFundingInterval(bytes32 marketID, uint256 fundingInterval) returns()
func (f *FuturesTransactor) ChangeMarketFundingInterval(opts *bind.TransactOpts, marketID [32]byte, fundingInterval *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMarketFundingInterval", marketID, fundingInterval)
}

// ChangeMarketFundingInterval is a paid mutator transaction binding the contract method 0x7a540e43.
//
// Solidity: function changeMarketFundingInterval(bytes32 marketID, uint256 fundingInterval) returns()
func (f *FuturesSession) ChangeMarketFundingInterval(marketID [32]byte, fundingInterval *big.Int) (*types.Transaction, error) {
	return f.Contract.ChangeMarketFundingInterval(&f.TransactOpts, marketID, fundingInterval)
}

// ChangeMarketFundingInterval is a paid mutator transaction binding the contract method 0x7a540e43.
//
// Solidity: function changeMarketFundingInterval(bytes32 marketID, uint256 fundingInterval) returns()
func (f *FuturesTransactorSession) ChangeMarketFundingInterval(marketID [32]byte, fundingInterval *big.Int) (*types.Transaction, error) {
	return f.Contract.ChangeMarketFundingInterval(&f.TransactOpts, marketID, fundingInterval)
}

// ChangeMarketInitialMarginRatio is a paid mutator transaction binding the contract method 0x50f2722b.
//
// Solidity: function changeMarketInitialMarginRatio(bytes32 marketID, (uint256) initialMarginRatio) returns()
func (f *FuturesTransactor) ChangeMarketInitialMarginRatio(opts *bind.TransactOpts, marketID [32]byte, initialMarginRatio PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMarketInitialMarginRatio", marketID, initialMarginRatio)
}

// ChangeMarketInitialMarginRatio is a paid mutator transaction binding the contract method 0x50f2722b.
//
// Solidity: function changeMarketInitialMarginRatio(bytes32 marketID, (uint256) initialMarginRatio) returns()
func (f *FuturesSession) ChangeMarketInitialMarginRatio(marketID [32]byte, initialMarginRatio PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketInitialMarginRatio(&f.TransactOpts, marketID, initialMarginRatio)
}

// ChangeMarketInitialMarginRatio is a paid mutator transaction binding the contract method 0x50f2722b.
//
// Solidity: function changeMarketInitialMarginRatio(bytes32 marketID, (uint256) initialMarginRatio) returns()
func (f *FuturesTransactorSession) ChangeMarketInitialMarginRatio(marketID [32]byte, initialMarginRatio PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketInitialMarginRatio(&f.TransactOpts, marketID, initialMarginRatio)
}

// ChangeMarketMaintenanceMarginRatio is a paid mutator transaction binding the contract method 0x267f8feb.
//
// Solidity: function changeMarketMaintenanceMarginRatio(bytes32 marketID, (uint256) maintenanceMarginRatio) returns()
func (f *FuturesTransactor) ChangeMarketMaintenanceMarginRatio(opts *bind.TransactOpts, marketID [32]byte, maintenanceMarginRatio PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMarketMaintenanceMarginRatio", marketID, maintenanceMarginRatio)
}

// ChangeMarketMaintenanceMarginRatio is a paid mutator transaction binding the contract method 0x267f8feb.
//
// Solidity: function changeMarketMaintenanceMarginRatio(bytes32 marketID, (uint256) maintenanceMarginRatio) returns()
func (f *FuturesSession) ChangeMarketMaintenanceMarginRatio(marketID [32]byte, maintenanceMarginRatio PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketMaintenanceMarginRatio(&f.TransactOpts, marketID, maintenanceMarginRatio)
}

// ChangeMarketMaintenanceMarginRatio is a paid mutator transaction binding the contract method 0x267f8feb.
//
// Solidity: function changeMarketMaintenanceMarginRatio(bytes32 marketID, (uint256) maintenanceMarginRatio) returns()
func (f *FuturesTransactorSession) ChangeMarketMaintenanceMarginRatio(marketID [32]byte, maintenanceMarginRatio PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketMaintenanceMarginRatio(&f.TransactOpts, marketID, maintenanceMarginRatio)
}

// ChangeMarketMakerTxFee is a paid mutator transaction binding the contract method 0x2c23608b.
//
// Solidity: function changeMarketMakerTxFee(bytes32 marketID, (uint256) makerTxFee) returns()
func (f *FuturesTransactor) ChangeMarketMakerTxFee(opts *bind.TransactOpts, marketID [32]byte, makerTxFee PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMarketMakerTxFee", marketID, makerTxFee)
}

// ChangeMarketMakerTxFee is a paid mutator transaction binding the contract method 0x2c23608b.
//
// Solidity: function changeMarketMakerTxFee(bytes32 marketID, (uint256) makerTxFee) returns()
func (f *FuturesSession) ChangeMarketMakerTxFee(marketID [32]byte, makerTxFee PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketMakerTxFee(&f.TransactOpts, marketID, makerTxFee)
}

// ChangeMarketMakerTxFee is a paid mutator transaction binding the contract method 0x2c23608b.
//
// Solidity: function changeMarketMakerTxFee(bytes32 marketID, (uint256) makerTxFee) returns()
func (f *FuturesTransactorSession) ChangeMarketMakerTxFee(marketID [32]byte, makerTxFee PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketMakerTxFee(&f.TransactOpts, marketID, makerTxFee)
}

// ChangeMarketOracle is a paid mutator transaction binding the contract method 0x555eb5ed.
//
// Solidity: function changeMarketOracle(bytes32 marketID, address oracle) returns()
func (f *FuturesTransactor) ChangeMarketOracle(opts *bind.TransactOpts, marketID [32]byte, oracle common.Address) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMarketOracle", marketID, oracle)
}

// ChangeMarketOracle is a paid mutator transaction binding the contract method 0x555eb5ed.
//
// Solidity: function changeMarketOracle(bytes32 marketID, address oracle) returns()
func (f *FuturesSession) ChangeMarketOracle(marketID [32]byte, oracle common.Address) (*types.Transaction, error) {
	return f.Contract.ChangeMarketOracle(&f.TransactOpts, marketID, oracle)
}

// ChangeMarketOracle is a paid mutator transaction binding the contract method 0x555eb5ed.
//
// Solidity: function changeMarketOracle(bytes32 marketID, address oracle) returns()
func (f *FuturesTransactorSession) ChangeMarketOracle(marketID [32]byte, oracle common.Address) (*types.Transaction, error) {
	return f.Contract.ChangeMarketOracle(&f.TransactOpts, marketID, oracle)
}

// ChangeMarketOracle0 is a paid mutator transaction binding the contract method 0xdb39261d.
//
// Solidity: function changeMarketOracle(bytes32 marketID, uint256 fundingInterval) returns()
func (f *FuturesTransactor) ChangeMarketOracle0(opts *bind.TransactOpts, marketID [32]byte, fundingInterval *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMarketOracle0", marketID, fundingInterval)
}

// ChangeMarketOracle0 is a paid mutator transaction binding the contract method 0xdb39261d.
//
// Solidity: function changeMarketOracle(bytes32 marketID, uint256 fundingInterval) returns()
func (f *FuturesSession) ChangeMarketOracle0(marketID [32]byte, fundingInterval *big.Int) (*types.Transaction, error) {
	return f.Contract.ChangeMarketOracle0(&f.TransactOpts, marketID, fundingInterval)
}

// ChangeMarketOracle0 is a paid mutator transaction binding the contract method 0xdb39261d.
//
// Solidity: function changeMarketOracle(bytes32 marketID, uint256 fundingInterval) returns()
func (f *FuturesTransactorSession) ChangeMarketOracle0(marketID [32]byte, fundingInterval *big.Int) (*types.Transaction, error) {
	return f.Contract.ChangeMarketOracle0(&f.TransactOpts, marketID, fundingInterval)
}

// ChangeMarketRelayerTxFee is a paid mutator transaction binding the contract method 0x6db346d3.
//
// Solidity: function changeMarketRelayerTxFee(bytes32 marketID, (uint256) relayerTxFee) returns()
func (f *FuturesTransactor) ChangeMarketRelayerTxFee(opts *bind.TransactOpts, marketID [32]byte, relayerTxFee PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMarketRelayerTxFee", marketID, relayerTxFee)
}

// ChangeMarketRelayerTxFee is a paid mutator transaction binding the contract method 0x6db346d3.
//
// Solidity: function changeMarketRelayerTxFee(bytes32 marketID, (uint256) relayerTxFee) returns()
func (f *FuturesSession) ChangeMarketRelayerTxFee(marketID [32]byte, relayerTxFee PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketRelayerTxFee(&f.TransactOpts, marketID, relayerTxFee)
}

// ChangeMarketRelayerTxFee is a paid mutator transaction binding the contract method 0x6db346d3.
//
// Solidity: function changeMarketRelayerTxFee(bytes32 marketID, (uint256) relayerTxFee) returns()
func (f *FuturesTransactorSession) ChangeMarketRelayerTxFee(marketID [32]byte, relayerTxFee PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketRelayerTxFee(&f.TransactOpts, marketID, relayerTxFee)
}

// ChangeMarketTakerTxFee is a paid mutator transaction binding the contract method 0xbba21c56.
//
// Solidity: function changeMarketTakerTxFee(bytes32 marketID, (uint256) takerTxFee) returns()
func (f *FuturesTransactor) ChangeMarketTakerTxFee(opts *bind.TransactOpts, marketID [32]byte, takerTxFee PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMarketTakerTxFee", marketID, takerTxFee)
}

// ChangeMarketTakerTxFee is a paid mutator transaction binding the contract method 0xbba21c56.
//
// Solidity: function changeMarketTakerTxFee(bytes32 marketID, (uint256) takerTxFee) returns()
func (f *FuturesSession) ChangeMarketTakerTxFee(marketID [32]byte, takerTxFee PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketTakerTxFee(&f.TransactOpts, marketID, takerTxFee)
}

// ChangeMarketTakerTxFee is a paid mutator transaction binding the contract method 0xbba21c56.
//
// Solidity: function changeMarketTakerTxFee(bytes32 marketID, (uint256) takerTxFee) returns()
func (f *FuturesTransactorSession) ChangeMarketTakerTxFee(marketID [32]byte, takerTxFee PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMarketTakerTxFee(&f.TransactOpts, marketID, takerTxFee)
}

// ChangeMaxBaseCurrencyCap is a paid mutator transaction binding the contract method 0xea7a4b3f.
//
// Solidity: function changeMaxBaseCurrencyCap(uint256 maxBaseCurrencyCap_) returns()
func (f *FuturesTransactor) ChangeMaxBaseCurrencyCap(opts *bind.TransactOpts, maxBaseCurrencyCap_ *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMaxBaseCurrencyCap", maxBaseCurrencyCap_)
}

// ChangeMaxBaseCurrencyCap is a paid mutator transaction binding the contract method 0xea7a4b3f.
//
// Solidity: function changeMaxBaseCurrencyCap(uint256 maxBaseCurrencyCap_) returns()
func (f *FuturesSession) ChangeMaxBaseCurrencyCap(maxBaseCurrencyCap_ *big.Int) (*types.Transaction, error) {
	return f.Contract.ChangeMaxBaseCurrencyCap(&f.TransactOpts, maxBaseCurrencyCap_)
}

// ChangeMaxBaseCurrencyCap is a paid mutator transaction binding the contract method 0xea7a4b3f.
//
// Solidity: function changeMaxBaseCurrencyCap(uint256 maxBaseCurrencyCap_) returns()
func (f *FuturesTransactorSession) ChangeMaxBaseCurrencyCap(maxBaseCurrencyCap_ *big.Int) (*types.Transaction, error) {
	return f.Contract.ChangeMaxBaseCurrencyCap(&f.TransactOpts, maxBaseCurrencyCap_)
}

// ChangeMinimumMarginRatio is a paid mutator transaction binding the contract method 0x5d19277f.
//
// Solidity: function changeMinimumMarginRatio((uint256) MINIMUM_MARGIN_RATIO_) returns()
func (f *FuturesTransactor) ChangeMinimumMarginRatio(opts *bind.TransactOpts, MINIMUM_MARGIN_RATIO_ PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.contract.Transact(opts, "changeMinimumMarginRatio", MINIMUM_MARGIN_RATIO_)
}

// ChangeMinimumMarginRatio is a paid mutator transaction binding the contract method 0x5d19277f.
//
// Solidity: function changeMinimumMarginRatio((uint256) MINIMUM_MARGIN_RATIO_) returns()
func (f *FuturesSession) ChangeMinimumMarginRatio(MINIMUM_MARGIN_RATIO_ PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMinimumMarginRatio(&f.TransactOpts, MINIMUM_MARGIN_RATIO_)
}

// ChangeMinimumMarginRatio is a paid mutator transaction binding the contract method 0x5d19277f.
//
// Solidity: function changeMinimumMarginRatio((uint256) MINIMUM_MARGIN_RATIO_) returns()
func (f *FuturesTransactorSession) ChangeMinimumMarginRatio(MINIMUM_MARGIN_RATIO_ PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.ChangeMinimumMarginRatio(&f.TransactOpts, MINIMUM_MARGIN_RATIO_)
}

// CheckFunding is a paid mutator transaction binding the contract method 0x70f1c88b.
//
// Solidity: function checkFunding(bytes32 marketID) returns(bool)
func (f *FuturesTransactor) CheckFunding(opts *bind.TransactOpts, marketID [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "checkFunding", marketID)
}

// CheckFunding is a paid mutator transaction binding the contract method 0x70f1c88b.
//
// Solidity: function checkFunding(bytes32 marketID) returns(bool)
func (f *FuturesSession) CheckFunding(marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.CheckFunding(&f.TransactOpts, marketID)
}

// CheckFunding is a paid mutator transaction binding the contract method 0x70f1c88b.
//
// Solidity: function checkFunding(bytes32 marketID) returns(bool)
func (f *FuturesTransactorSession) CheckFunding(marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.CheckFunding(&f.TransactOpts, marketID)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x73ea5056.
//
// Solidity: function closePosition(bytes32 subAccountID, bytes32 marketID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bool isRevertingOnPartialFills, bytes[] signatures) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactor) ClosePosition(opts *bind.TransactOpts, subAccountID [32]byte, marketID [32]byte, orders []Order, quantity *big.Int, isRevertingOnPartialFills bool, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "closePosition", subAccountID, marketID, orders, quantity, isRevertingOnPartialFills, signatures)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x73ea5056.
//
// Solidity: function closePosition(bytes32 subAccountID, bytes32 marketID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bool isRevertingOnPartialFills, bytes[] signatures) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesSession) ClosePosition(subAccountID [32]byte, marketID [32]byte, orders []Order, quantity *big.Int, isRevertingOnPartialFills bool, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.ClosePosition(&f.TransactOpts, subAccountID, marketID, orders, quantity, isRevertingOnPartialFills, signatures)
}

// ClosePosition is a paid mutator transaction binding the contract method 0x73ea5056.
//
// Solidity: function closePosition(bytes32 subAccountID, bytes32 marketID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bool isRevertingOnPartialFills, bytes[] signatures) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactorSession) ClosePosition(subAccountID [32]byte, marketID [32]byte, orders []Order, quantity *big.Int, isRevertingOnPartialFills bool, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.ClosePosition(&f.TransactOpts, subAccountID, marketID, orders, quantity, isRevertingOnPartialFills, signatures)
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

// CreateFuturesMarket is a paid mutator transaction binding the contract method 0xc47e094b.
//
// Solidity: function createFuturesMarket(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatio, (uint256) maintenanceMarginRatio, uint256 fundingInterval, uint256 expirationTime, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage) returns()
func (f *FuturesTransactor) CreateFuturesMarket(opts *bind.TransactOpts, ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatio PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, expirationTime *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createFuturesMarket", ticker, baseCurrency, oracle, initialMarginRatio, maintenanceMarginRatio, fundingInterval, expirationTime, makerTxFee, takerTxFee, relayerFeePercentage)
}

// CreateFuturesMarket is a paid mutator transaction binding the contract method 0xc47e094b.
//
// Solidity: function createFuturesMarket(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatio, (uint256) maintenanceMarginRatio, uint256 fundingInterval, uint256 expirationTime, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage) returns()
func (f *FuturesSession) CreateFuturesMarket(ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatio PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, expirationTime *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.CreateFuturesMarket(&f.TransactOpts, ticker, baseCurrency, oracle, initialMarginRatio, maintenanceMarginRatio, fundingInterval, expirationTime, makerTxFee, takerTxFee, relayerFeePercentage)
}

// CreateFuturesMarket is a paid mutator transaction binding the contract method 0xc47e094b.
//
// Solidity: function createFuturesMarket(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatio, (uint256) maintenanceMarginRatio, uint256 fundingInterval, uint256 expirationTime, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage) returns()
func (f *FuturesTransactorSession) CreateFuturesMarket(ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatio PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, expirationTime *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.CreateFuturesMarket(&f.TransactOpts, ticker, baseCurrency, oracle, initialMarginRatio, maintenanceMarginRatio, fundingInterval, expirationTime, makerTxFee, takerTxFee, relayerFeePercentage)
}

// CreateFuturesMarketWithFixedMarketId is a paid mutator transaction binding the contract method 0x28f9d7a5.
//
// Solidity: function createFuturesMarketWithFixedMarketId(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatio, (uint256) maintenanceMarginRatio, uint256 fundingInterval, uint256 expirationTime, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage, bytes32 marketID) returns()
func (f *FuturesTransactor) CreateFuturesMarketWithFixedMarketId(opts *bind.TransactOpts, ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatio PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, expirationTime *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad, marketID [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createFuturesMarketWithFixedMarketId", ticker, baseCurrency, oracle, initialMarginRatio, maintenanceMarginRatio, fundingInterval, expirationTime, makerTxFee, takerTxFee, relayerFeePercentage, marketID)
}

// CreateFuturesMarketWithFixedMarketId is a paid mutator transaction binding the contract method 0x28f9d7a5.
//
// Solidity: function createFuturesMarketWithFixedMarketId(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatio, (uint256) maintenanceMarginRatio, uint256 fundingInterval, uint256 expirationTime, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage, bytes32 marketID) returns()
func (f *FuturesSession) CreateFuturesMarketWithFixedMarketId(ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatio PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, expirationTime *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad, marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.CreateFuturesMarketWithFixedMarketId(&f.TransactOpts, ticker, baseCurrency, oracle, initialMarginRatio, maintenanceMarginRatio, fundingInterval, expirationTime, makerTxFee, takerTxFee, relayerFeePercentage, marketID)
}

// CreateFuturesMarketWithFixedMarketId is a paid mutator transaction binding the contract method 0x28f9d7a5.
//
// Solidity: function createFuturesMarketWithFixedMarketId(string ticker, address baseCurrency, address oracle, (uint256) initialMarginRatio, (uint256) maintenanceMarginRatio, uint256 fundingInterval, uint256 expirationTime, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage, bytes32 marketID) returns()
func (f *FuturesTransactorSession) CreateFuturesMarketWithFixedMarketId(ticker string, baseCurrency common.Address, oracle common.Address, initialMarginRatio PermyriadMathPermyriad, maintenanceMarginRatio PermyriadMathPermyriad, fundingInterval *big.Int, expirationTime *big.Int, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad, marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.CreateFuturesMarketWithFixedMarketId(&f.TransactOpts, ticker, baseCurrency, oracle, initialMarginRatio, maintenanceMarginRatio, fundingInterval, expirationTime, makerTxFee, takerTxFee, relayerFeePercentage, marketID)
}

// CreateSpotMarket is a paid mutator transaction binding the contract method 0xadb44540.
//
// Solidity: function createSpotMarket(string ticker, address baseToken, address quoteToken, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage) returns()
func (f *FuturesTransactor) CreateSpotMarket(opts *bind.TransactOpts, ticker string, baseToken common.Address, quoteToken common.Address, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createSpotMarket", ticker, baseToken, quoteToken, makerTxFee, takerTxFee, relayerFeePercentage)
}

// CreateSpotMarket is a paid mutator transaction binding the contract method 0xadb44540.
//
// Solidity: function createSpotMarket(string ticker, address baseToken, address quoteToken, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage) returns()
func (f *FuturesSession) CreateSpotMarket(ticker string, baseToken common.Address, quoteToken common.Address, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.CreateSpotMarket(&f.TransactOpts, ticker, baseToken, quoteToken, makerTxFee, takerTxFee, relayerFeePercentage)
}

// CreateSpotMarket is a paid mutator transaction binding the contract method 0xadb44540.
//
// Solidity: function createSpotMarket(string ticker, address baseToken, address quoteToken, (uint256) makerTxFee, (uint256) takerTxFee, (uint256) relayerFeePercentage) returns()
func (f *FuturesTransactorSession) CreateSpotMarket(ticker string, baseToken common.Address, quoteToken common.Address, makerTxFee PermyriadMathPermyriad, takerTxFee PermyriadMathPermyriad, relayerFeePercentage PermyriadMathPermyriad) (*types.Transaction, error) {
	return f.Contract.CreateSpotMarket(&f.TransactOpts, ticker, baseToken, quoteToken, makerTxFee, takerTxFee, relayerFeePercentage)
}

// CreateSubAccountAndDeposit is a paid mutator transaction binding the contract method 0x06bc2e72.
//
// Solidity: function createSubAccountAndDeposit(address baseCurrency, uint96 subAccountNonce, uint256 amount) returns(bytes32)
func (f *FuturesTransactor) CreateSubAccountAndDeposit(opts *bind.TransactOpts, baseCurrency common.Address, subAccountNonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createSubAccountAndDeposit", baseCurrency, subAccountNonce, amount)
}

// CreateSubAccountAndDeposit is a paid mutator transaction binding the contract method 0x06bc2e72.
//
// Solidity: function createSubAccountAndDeposit(address baseCurrency, uint96 subAccountNonce, uint256 amount) returns(bytes32)
func (f *FuturesSession) CreateSubAccountAndDeposit(baseCurrency common.Address, subAccountNonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateSubAccountAndDeposit(&f.TransactOpts, baseCurrency, subAccountNonce, amount)
}

// CreateSubAccountAndDeposit is a paid mutator transaction binding the contract method 0x06bc2e72.
//
// Solidity: function createSubAccountAndDeposit(address baseCurrency, uint96 subAccountNonce, uint256 amount) returns(bytes32)
func (f *FuturesTransactorSession) CreateSubAccountAndDeposit(baseCurrency common.Address, subAccountNonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateSubAccountAndDeposit(&f.TransactOpts, baseCurrency, subAccountNonce, amount)
}

// CreateSubAccountForTraderWithNonce is a paid mutator transaction binding the contract method 0x510d4093.
//
// Solidity: function createSubAccountForTraderWithNonce(address trader, uint96 subAccountNonce) returns(bytes32)
func (f *FuturesTransactor) CreateSubAccountForTraderWithNonce(opts *bind.TransactOpts, trader common.Address, subAccountNonce *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createSubAccountForTraderWithNonce", trader, subAccountNonce)
}

// CreateSubAccountForTraderWithNonce is a paid mutator transaction binding the contract method 0x510d4093.
//
// Solidity: function createSubAccountForTraderWithNonce(address trader, uint96 subAccountNonce) returns(bytes32)
func (f *FuturesSession) CreateSubAccountForTraderWithNonce(trader common.Address, subAccountNonce *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateSubAccountForTraderWithNonce(&f.TransactOpts, trader, subAccountNonce)
}

// CreateSubAccountForTraderWithNonce is a paid mutator transaction binding the contract method 0x510d4093.
//
// Solidity: function createSubAccountForTraderWithNonce(address trader, uint96 subAccountNonce) returns(bytes32)
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

// DepositForSubaccount is a paid mutator transaction binding the contract method 0x9d44c2b1.
//
// Solidity: function depositForSubaccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesTransactor) DepositForSubaccount(opts *bind.TransactOpts, baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "depositForSubaccount", baseCurrency, subAccountID, amount)
}

// DepositForSubaccount is a paid mutator transaction binding the contract method 0x9d44c2b1.
//
// Solidity: function depositForSubaccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesSession) DepositForSubaccount(baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositForSubaccount(&f.TransactOpts, baseCurrency, subAccountID, amount)
}

// DepositForSubaccount is a paid mutator transaction binding the contract method 0x9d44c2b1.
//
// Solidity: function depositForSubaccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesTransactorSession) DepositForSubaccount(baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositForSubaccount(&f.TransactOpts, baseCurrency, subAccountID, amount)
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

// EmergencyStopMarket is a paid mutator transaction binding the contract method 0xd135eeb3.
//
// Solidity: function emergencyStopMarket(bytes32 marketID) returns()
func (f *FuturesTransactor) EmergencyStopMarket(opts *bind.TransactOpts, marketID [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "emergencyStopMarket", marketID)
}

// EmergencyStopMarket is a paid mutator transaction binding the contract method 0xd135eeb3.
//
// Solidity: function emergencyStopMarket(bytes32 marketID) returns()
func (f *FuturesSession) EmergencyStopMarket(marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.EmergencyStopMarket(&f.TransactOpts, marketID)
}

// EmergencyStopMarket is a paid mutator transaction binding the contract method 0xd135eeb3.
//
// Solidity: function emergencyStopMarket(bytes32 marketID) returns()
func (f *FuturesTransactorSession) EmergencyStopMarket(marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.EmergencyStopMarket(&f.TransactOpts, marketID)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2280c910.
//
// Solidity: function executeTransaction((uint256,uint256,uint256,address,bytes) transaction, bytes signature) payable returns(bytes)
func (f *FuturesTransactor) ExecuteTransaction(opts *bind.TransactOpts, transaction LibZeroExTransactionZeroExTransaction, signature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "executeTransaction", transaction, signature)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2280c910.
//
// Solidity: function executeTransaction((uint256,uint256,uint256,address,bytes) transaction, bytes signature) payable returns(bytes)
func (f *FuturesSession) ExecuteTransaction(transaction LibZeroExTransactionZeroExTransaction, signature []byte) (*types.Transaction, error) {
	return f.Contract.ExecuteTransaction(&f.TransactOpts, transaction, signature)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2280c910.
//
// Solidity: function executeTransaction((uint256,uint256,uint256,address,bytes) transaction, bytes signature) payable returns(bytes)
func (f *FuturesTransactorSession) ExecuteTransaction(transaction LibZeroExTransactionZeroExTransaction, signature []byte) (*types.Transaction, error) {
	return f.Contract.ExecuteTransaction(&f.TransactOpts, transaction, signature)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xd6128693.
//
// Solidity: function liquidatePosition(bytes32 subAccountID, bytes32 marketID, address liquidationCaller, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactor) LiquidatePosition(opts *bind.TransactOpts, subAccountID [32]byte, marketID [32]byte, liquidationCaller common.Address, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "liquidatePosition", subAccountID, marketID, liquidationCaller, orders, quantity, signatures)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xd6128693.
//
// Solidity: function liquidatePosition(bytes32 subAccountID, bytes32 marketID, address liquidationCaller, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesSession) LiquidatePosition(subAccountID [32]byte, marketID [32]byte, liquidationCaller common.Address, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.LiquidatePosition(&f.TransactOpts, subAccountID, marketID, liquidationCaller, orders, quantity, signatures)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xd6128693.
//
// Solidity: function liquidatePosition(bytes32 subAccountID, bytes32 marketID, address liquidationCaller, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactorSession) LiquidatePosition(subAccountID [32]byte, marketID [32]byte, liquidationCaller common.Address, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.LiquidatePosition(&f.TransactOpts, subAccountID, marketID, liquidationCaller, orders, quantity, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xfd7d3b8a.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes32 subAccountID, bool isRevertingOnPartialFills, bytes[] signatures, address feeRecipientAddress) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactor) MarketOrders(opts *bind.TransactOpts, orders []Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, isRevertingOnPartialFills bool, signatures [][]byte, feeRecipientAddress common.Address) (*types.Transaction, error) {
	return f.contract.Transact(opts, "marketOrders", orders, quantity, margin, subAccountID, isRevertingOnPartialFills, signatures, feeRecipientAddress)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xfd7d3b8a.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes32 subAccountID, bool isRevertingOnPartialFills, bytes[] signatures, address feeRecipientAddress) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesSession) MarketOrders(orders []Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, isRevertingOnPartialFills bool, signatures [][]byte, feeRecipientAddress common.Address) (*types.Transaction, error) {
	return f.Contract.MarketOrders(&f.TransactOpts, orders, quantity, margin, subAccountID, isRevertingOnPartialFills, signatures, feeRecipientAddress)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xfd7d3b8a.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes32 subAccountID, bool isRevertingOnPartialFills, bytes[] signatures, address feeRecipientAddress) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] results)
func (f *FuturesTransactorSession) MarketOrders(orders []Order, quantity *big.Int, margin *big.Int, subAccountID [32]byte, isRevertingOnPartialFills bool, signatures [][]byte, feeRecipientAddress common.Address) (*types.Transaction, error) {
	return f.Contract.MarketOrders(&f.TransactOpts, orders, quantity, margin, subAccountID, isRevertingOnPartialFills, signatures, feeRecipientAddress)
}

// MatchSpotOrders is a paid mutator transaction binding the contract method 0xc7e34978.
//
// Solidity: function matchSpotOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns(((uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)) matchedFillResults)
func (f *FuturesTransactor) MatchSpotOrders(opts *bind.TransactOpts, leftOrder Order, rightOrder Order, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "matchSpotOrders", leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchSpotOrders is a paid mutator transaction binding the contract method 0xc7e34978.
//
// Solidity: function matchSpotOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns(((uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)) matchedFillResults)
func (f *FuturesSession) MatchSpotOrders(leftOrder Order, rightOrder Order, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MatchSpotOrders(&f.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchSpotOrders is a paid mutator transaction binding the contract method 0xc7e34978.
//
// Solidity: function matchSpotOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns(((uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256)) matchedFillResults)
func (f *FuturesTransactorSession) MatchSpotOrders(leftOrder Order, rightOrder Order, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MatchSpotOrders(&f.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
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

// MultiMatchSpotOrders is a paid mutator transaction binding the contract method 0xd8e2259c.
//
// Solidity: function multiMatchSpotOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) returns(((uint256,uint256,uint256,uint256)[],(uint256,uint256,uint256,uint256)[]) batchMatchedFillResults)
func (f *FuturesTransactor) MultiMatchSpotOrders(opts *bind.TransactOpts, leftOrders []Order, rightOrders []Order, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "multiMatchSpotOrders", leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// MultiMatchSpotOrders is a paid mutator transaction binding the contract method 0xd8e2259c.
//
// Solidity: function multiMatchSpotOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) returns(((uint256,uint256,uint256,uint256)[],(uint256,uint256,uint256,uint256)[]) batchMatchedFillResults)
func (f *FuturesSession) MultiMatchSpotOrders(leftOrders []Order, rightOrders []Order, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return f.Contract.MultiMatchSpotOrders(&f.TransactOpts, leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// MultiMatchSpotOrders is a paid mutator transaction binding the contract method 0xd8e2259c.
//
// Solidity: function multiMatchSpotOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) returns(((uint256,uint256,uint256,uint256)[],(uint256,uint256,uint256,uint256)[]) batchMatchedFillResults)
func (f *FuturesTransactorSession) MultiMatchSpotOrders(leftOrders []Order, rightOrders []Order, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return f.Contract.MultiMatchSpotOrders(&f.TransactOpts, leftOrders, rightOrders, leftSignatures, rightSignatures)
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

// PreSign is a paid mutator transaction binding the contract method 0x46c02d7a.
//
// Solidity: function preSign(bytes32 hash) payable returns()
func (f *FuturesTransactor) PreSign(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "preSign", hash)
}

// PreSign is a paid mutator transaction binding the contract method 0x46c02d7a.
//
// Solidity: function preSign(bytes32 hash) payable returns()
func (f *FuturesSession) PreSign(hash [32]byte) (*types.Transaction, error) {
	return f.Contract.PreSign(&f.TransactOpts, hash)
}

// PreSign is a paid mutator transaction binding the contract method 0x46c02d7a.
//
// Solidity: function preSign(bytes32 hash) payable returns()
func (f *FuturesTransactorSession) PreSign(hash [32]byte) (*types.Transaction, error) {
	return f.Contract.PreSign(&f.TransactOpts, hash)
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

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xbc053e73.
//
// Solidity: function setApprovalForAll(bytes32 fromSubAccountID, address operator, bool isApproved) returns()
func (f *FuturesTransactor) SetApprovalForAll(opts *bind.TransactOpts, fromSubAccountID [32]byte, operator common.Address, isApproved bool) (*types.Transaction, error) {
	return f.contract.Transact(opts, "setApprovalForAll", fromSubAccountID, operator, isApproved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xbc053e73.
//
// Solidity: function setApprovalForAll(bytes32 fromSubAccountID, address operator, bool isApproved) returns()
func (f *FuturesSession) SetApprovalForAll(fromSubAccountID [32]byte, operator common.Address, isApproved bool) (*types.Transaction, error) {
	return f.Contract.SetApprovalForAll(&f.TransactOpts, fromSubAccountID, operator, isApproved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xbc053e73.
//
// Solidity: function setApprovalForAll(bytes32 fromSubAccountID, address operator, bool isApproved) returns()
func (f *FuturesTransactorSession) SetApprovalForAll(fromSubAccountID [32]byte, operator common.Address, isApproved bool) (*types.Transaction, error) {
	return f.Contract.SetApprovalForAll(&f.TransactOpts, fromSubAccountID, operator, isApproved)
}

// SetApprovalForMarket is a paid mutator transaction binding the contract method 0xa01856c1.
//
// Solidity: function setApprovalForMarket(bytes32 fromSubAccountID, address operator, bytes32 marketID, bool isApproved) returns()
func (f *FuturesTransactor) SetApprovalForMarket(opts *bind.TransactOpts, fromSubAccountID [32]byte, operator common.Address, marketID [32]byte, isApproved bool) (*types.Transaction, error) {
	return f.contract.Transact(opts, "setApprovalForMarket", fromSubAccountID, operator, marketID, isApproved)
}

// SetApprovalForMarket is a paid mutator transaction binding the contract method 0xa01856c1.
//
// Solidity: function setApprovalForMarket(bytes32 fromSubAccountID, address operator, bytes32 marketID, bool isApproved) returns()
func (f *FuturesSession) SetApprovalForMarket(fromSubAccountID [32]byte, operator common.Address, marketID [32]byte, isApproved bool) (*types.Transaction, error) {
	return f.Contract.SetApprovalForMarket(&f.TransactOpts, fromSubAccountID, operator, marketID, isApproved)
}

// SetApprovalForMarket is a paid mutator transaction binding the contract method 0xa01856c1.
//
// Solidity: function setApprovalForMarket(bytes32 fromSubAccountID, address operator, bytes32 marketID, bool isApproved) returns()
func (f *FuturesTransactorSession) SetApprovalForMarket(fromSubAccountID [32]byte, operator common.Address, marketID [32]byte, isApproved bool) (*types.Transaction, error) {
	return f.Contract.SetApprovalForMarket(&f.TransactOpts, fromSubAccountID, operator, marketID, isApproved)
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

// SetReceiptApprovalForAll is a paid mutator transaction binding the contract method 0xaec3d88a.
//
// Solidity: function setReceiptApprovalForAll(bytes32 receiverSubAccountID, address operator, bool isApproved) returns()
func (f *FuturesTransactor) SetReceiptApprovalForAll(opts *bind.TransactOpts, receiverSubAccountID [32]byte, operator common.Address, isApproved bool) (*types.Transaction, error) {
	return f.contract.Transact(opts, "setReceiptApprovalForAll", receiverSubAccountID, operator, isApproved)
}

// SetReceiptApprovalForAll is a paid mutator transaction binding the contract method 0xaec3d88a.
//
// Solidity: function setReceiptApprovalForAll(bytes32 receiverSubAccountID, address operator, bool isApproved) returns()
func (f *FuturesSession) SetReceiptApprovalForAll(receiverSubAccountID [32]byte, operator common.Address, isApproved bool) (*types.Transaction, error) {
	return f.Contract.SetReceiptApprovalForAll(&f.TransactOpts, receiverSubAccountID, operator, isApproved)
}

// SetReceiptApprovalForAll is a paid mutator transaction binding the contract method 0xaec3d88a.
//
// Solidity: function setReceiptApprovalForAll(bytes32 receiverSubAccountID, address operator, bool isApproved) returns()
func (f *FuturesTransactorSession) SetReceiptApprovalForAll(receiverSubAccountID [32]byte, operator common.Address, isApproved bool) (*types.Transaction, error) {
	return f.Contract.SetReceiptApprovalForAll(&f.TransactOpts, receiverSubAccountID, operator, isApproved)
}

// SetReceiptApprovalForMarket is a paid mutator transaction binding the contract method 0x14229a3f.
//
// Solidity: function setReceiptApprovalForMarket(bytes32 receiverSubAccountID, address operator, bytes32 marketID, bool isApproved) returns()
func (f *FuturesTransactor) SetReceiptApprovalForMarket(opts *bind.TransactOpts, receiverSubAccountID [32]byte, operator common.Address, marketID [32]byte, isApproved bool) (*types.Transaction, error) {
	return f.contract.Transact(opts, "setReceiptApprovalForMarket", receiverSubAccountID, operator, marketID, isApproved)
}

// SetReceiptApprovalForMarket is a paid mutator transaction binding the contract method 0x14229a3f.
//
// Solidity: function setReceiptApprovalForMarket(bytes32 receiverSubAccountID, address operator, bytes32 marketID, bool isApproved) returns()
func (f *FuturesSession) SetReceiptApprovalForMarket(receiverSubAccountID [32]byte, operator common.Address, marketID [32]byte, isApproved bool) (*types.Transaction, error) {
	return f.Contract.SetReceiptApprovalForMarket(&f.TransactOpts, receiverSubAccountID, operator, marketID, isApproved)
}

// SetReceiptApprovalForMarket is a paid mutator transaction binding the contract method 0x14229a3f.
//
// Solidity: function setReceiptApprovalForMarket(bytes32 receiverSubAccountID, address operator, bytes32 marketID, bool isApproved) returns()
func (f *FuturesTransactorSession) SetReceiptApprovalForMarket(receiverSubAccountID [32]byte, operator common.Address, marketID [32]byte, isApproved bool) (*types.Transaction, error) {
	return f.Contract.SetReceiptApprovalForMarket(&f.TransactOpts, receiverSubAccountID, operator, marketID, isApproved)
}

// SetSignatureValidatorApproval is a paid mutator transaction binding the contract method 0x77fcce68.
//
// Solidity: function setSignatureValidatorApproval(address validatorAddress, bool approval) payable returns()
func (f *FuturesTransactor) SetSignatureValidatorApproval(opts *bind.TransactOpts, validatorAddress common.Address, approval bool) (*types.Transaction, error) {
	return f.contract.Transact(opts, "setSignatureValidatorApproval", validatorAddress, approval)
}

// SetSignatureValidatorApproval is a paid mutator transaction binding the contract method 0x77fcce68.
//
// Solidity: function setSignatureValidatorApproval(address validatorAddress, bool approval) payable returns()
func (f *FuturesSession) SetSignatureValidatorApproval(validatorAddress common.Address, approval bool) (*types.Transaction, error) {
	return f.Contract.SetSignatureValidatorApproval(&f.TransactOpts, validatorAddress, approval)
}

// SetSignatureValidatorApproval is a paid mutator transaction binding the contract method 0x77fcce68.
//
// Solidity: function setSignatureValidatorApproval(address validatorAddress, bool approval) payable returns()
func (f *FuturesTransactorSession) SetSignatureValidatorApproval(validatorAddress common.Address, approval bool) (*types.Transaction, error) {
	return f.Contract.SetSignatureValidatorApproval(&f.TransactOpts, validatorAddress, approval)
}

// SettleExpiryFuturesPosition is a paid mutator transaction binding the contract method 0xa30b87ad.
//
// Solidity: function settleExpiryFuturesPosition(bytes32 marketID, bytes32 subAccountID) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactor) SettleExpiryFuturesPosition(opts *bind.TransactOpts, marketID [32]byte, subAccountID [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "settleExpiryFuturesPosition", marketID, subAccountID)
}

// SettleExpiryFuturesPosition is a paid mutator transaction binding the contract method 0xa30b87ad.
//
// Solidity: function settleExpiryFuturesPosition(bytes32 marketID, bytes32 subAccountID) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesSession) SettleExpiryFuturesPosition(marketID [32]byte, subAccountID [32]byte) (*types.Transaction, error) {
	return f.Contract.SettleExpiryFuturesPosition(&f.TransactOpts, marketID, subAccountID)
}

// SettleExpiryFuturesPosition is a paid mutator transaction binding the contract method 0xa30b87ad.
//
// Solidity: function settleExpiryFuturesPosition(bytes32 marketID, bytes32 subAccountID) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactorSession) SettleExpiryFuturesPosition(marketID [32]byte, subAccountID [32]byte) (*types.Transaction, error) {
	return f.Contract.SettleExpiryFuturesPosition(&f.TransactOpts, marketID, subAccountID)
}

// SettleMarket is a paid mutator transaction binding the contract method 0xe039c00e.
//
// Solidity: function settleMarket(bytes32 marketID) returns(bool)
func (f *FuturesTransactor) SettleMarket(opts *bind.TransactOpts, marketID [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "settleMarket", marketID)
}

// SettleMarket is a paid mutator transaction binding the contract method 0xe039c00e.
//
// Solidity: function settleMarket(bytes32 marketID) returns(bool)
func (f *FuturesSession) SettleMarket(marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.SettleMarket(&f.TransactOpts, marketID)
}

// SettleMarket is a paid mutator transaction binding the contract method 0xe039c00e.
//
// Solidity: function settleMarket(bytes32 marketID) returns(bool)
func (f *FuturesTransactorSession) SettleMarket(marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.SettleMarket(&f.TransactOpts, marketID)
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

// TransferPosition is a paid mutator transaction binding the contract method 0x9380ba16.
//
// Solidity: function transferPosition(uint256 positionID, bytes32 receiverSubAccountID, uint256 transferQuantity) returns()
func (f *FuturesTransactor) TransferPosition(opts *bind.TransactOpts, positionID *big.Int, receiverSubAccountID [32]byte, transferQuantity *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "transferPosition", positionID, receiverSubAccountID, transferQuantity)
}

// TransferPosition is a paid mutator transaction binding the contract method 0x9380ba16.
//
// Solidity: function transferPosition(uint256 positionID, bytes32 receiverSubAccountID, uint256 transferQuantity) returns()
func (f *FuturesSession) TransferPosition(positionID *big.Int, receiverSubAccountID [32]byte, transferQuantity *big.Int) (*types.Transaction, error) {
	return f.Contract.TransferPosition(&f.TransactOpts, positionID, receiverSubAccountID, transferQuantity)
}

// TransferPosition is a paid mutator transaction binding the contract method 0x9380ba16.
//
// Solidity: function transferPosition(uint256 positionID, bytes32 receiverSubAccountID, uint256 transferQuantity) returns()
func (f *FuturesTransactorSession) TransferPosition(positionID *big.Int, receiverSubAccountID [32]byte, transferQuantity *big.Int) (*types.Transaction, error) {
	return f.Contract.TransferPosition(&f.TransactOpts, positionID, receiverSubAccountID, transferQuantity)
}

// TransferToSubAccount is a paid mutator transaction binding the contract method 0xddf84a44.
//
// Solidity: function transferToSubAccount(address baseCurrency, bytes32 fromSubAccountID, bytes32 toSubAccountID, uint256 amount) returns()
func (f *FuturesTransactor) TransferToSubAccount(opts *bind.TransactOpts, baseCurrency common.Address, fromSubAccountID [32]byte, toSubAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "transferToSubAccount", baseCurrency, fromSubAccountID, toSubAccountID, amount)
}

// TransferToSubAccount is a paid mutator transaction binding the contract method 0xddf84a44.
//
// Solidity: function transferToSubAccount(address baseCurrency, bytes32 fromSubAccountID, bytes32 toSubAccountID, uint256 amount) returns()
func (f *FuturesSession) TransferToSubAccount(baseCurrency common.Address, fromSubAccountID [32]byte, toSubAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.TransferToSubAccount(&f.TransactOpts, baseCurrency, fromSubAccountID, toSubAccountID, amount)
}

// TransferToSubAccount is a paid mutator transaction binding the contract method 0xddf84a44.
//
// Solidity: function transferToSubAccount(address baseCurrency, bytes32 fromSubAccountID, bytes32 toSubAccountID, uint256 amount) returns()
func (f *FuturesTransactorSession) TransferToSubAccount(baseCurrency common.Address, fromSubAccountID [32]byte, toSubAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.TransferToSubAccount(&f.TransactOpts, baseCurrency, fromSubAccountID, toSubAccountID, amount)
}

// VaporizePosition is a paid mutator transaction binding the contract method 0xcaa065b9.
//
// Solidity: function vaporizePosition(bytes32 subAccountID, bytes32 marketID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactor) VaporizePosition(opts *bind.TransactOpts, subAccountID [32]byte, marketID [32]byte, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "vaporizePosition", subAccountID, marketID, orders, quantity, signatures)
}

// VaporizePosition is a paid mutator transaction binding the contract method 0xcaa065b9.
//
// Solidity: function vaporizePosition(bytes32 subAccountID, bytes32 marketID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesSession) VaporizePosition(subAccountID [32]byte, marketID [32]byte, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.VaporizePosition(&f.TransactOpts, subAccountID, marketID, orders, quantity, signatures)
}

// VaporizePosition is a paid mutator transaction binding the contract method 0xcaa065b9.
//
// Solidity: function vaporizePosition(bytes32 subAccountID, bytes32 marketID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes[] signatures) returns((uint256,uint256,uint256,uint256) results)
func (f *FuturesTransactorSession) VaporizePosition(subAccountID [32]byte, marketID [32]byte, orders []Order, quantity *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.VaporizePosition(&f.TransactOpts, subAccountID, marketID, orders, quantity, signatures)
}

// WithdrawForSubAccount is a paid mutator transaction binding the contract method 0x803a90d8.
//
// Solidity: function withdrawForSubAccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesTransactor) WithdrawForSubAccount(opts *bind.TransactOpts, baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "withdrawForSubAccount", baseCurrency, subAccountID, amount)
}

// WithdrawForSubAccount is a paid mutator transaction binding the contract method 0x803a90d8.
//
// Solidity: function withdrawForSubAccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesSession) WithdrawForSubAccount(baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawForSubAccount(&f.TransactOpts, baseCurrency, subAccountID, amount)
}

// WithdrawForSubAccount is a paid mutator transaction binding the contract method 0x803a90d8.
//
// Solidity: function withdrawForSubAccount(address baseCurrency, bytes32 subAccountID, uint256 amount) returns()
func (f *FuturesTransactorSession) WithdrawForSubAccount(baseCurrency common.Address, subAccountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawForSubAccount(&f.TransactOpts, baseCurrency, subAccountID, amount)
}

// WithdrawFromInsurancePool is a paid mutator transaction binding the contract method 0xb7d4078d.
//
// Solidity: function withdrawFromInsurancePool(address receiver, bytes32 marketID, uint256 amount) returns()
func (f *FuturesTransactor) WithdrawFromInsurancePool(opts *bind.TransactOpts, receiver common.Address, marketID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "withdrawFromInsurancePool", receiver, marketID, amount)
}

// WithdrawFromInsurancePool is a paid mutator transaction binding the contract method 0xb7d4078d.
//
// Solidity: function withdrawFromInsurancePool(address receiver, bytes32 marketID, uint256 amount) returns()
func (f *FuturesSession) WithdrawFromInsurancePool(receiver common.Address, marketID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawFromInsurancePool(&f.TransactOpts, receiver, marketID, amount)
}

// WithdrawFromInsurancePool is a paid mutator transaction binding the contract method 0xb7d4078d.
//
// Solidity: function withdrawFromInsurancePool(address receiver, bytes32 marketID, uint256 amount) returns()
func (f *FuturesTransactorSession) WithdrawFromInsurancePool(receiver common.Address, marketID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawFromInsurancePool(&f.TransactOpts, receiver, marketID, amount)
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

// FilterAccountCreation is a free log retrieval operation binding the contract event 0x6da62924036c396387aa0340b6df622a60ea3f2dcd19bd4a7d77df13013f3043.
//
// Solidity: event AccountCreation(address indexed creator, bytes32 subAccountID, uint96 subAccountNonce)
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

// WatchAccountCreation is a free log subscription operation binding the contract event 0x6da62924036c396387aa0340b6df622a60ea3f2dcd19bd4a7d77df13013f3043.
//
// Solidity: event AccountCreation(address indexed creator, bytes32 subAccountID, uint96 subAccountNonce)
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

// ParseAccountCreation is a log parse operation binding the contract event 0x6da62924036c396387aa0340b6df622a60ea3f2dcd19bd4a7d77df13013f3043.
//
// Solidity: event AccountCreation(address indexed creator, bytes32 subAccountID, uint96 subAccountNonce)
func (f *FuturesFilterer) ParseAccountCreation(log types.Log) (*FuturesAccountCreation, error) {
	event := new(FuturesAccountCreation)
	if err := f.contract.UnpackLog(event, "AccountCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	SubAccountID   [32]byte
	OrderHash      [32]byte
	MarketID       [32]byte
	BaseCurrency   common.Address
	ContractPrice  *big.Int
	QuantityFilled *big.Int
	IsLong         bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFuturesCancel is a free log retrieval operation binding the contract event 0x7197f6513c22ccffd678f7d190c7c57116b046d6ea34e83c1ba7fd841f966ec4.
//
// Solidity: event FuturesCancel(bytes32 indexed subAccountID, bytes32 indexed orderHash, bytes32 indexed marketID, address baseCurrency, uint256 contractPrice, uint256 quantityFilled, bool isLong)
func (f *FuturesFilterer) FilterFuturesCancel(opts *bind.FilterOpts, subAccountID [][32]byte, orderHash [][32]byte, marketID [][32]byte) (*FuturesCancelIterator, error) {

	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesCancel", subAccountIDRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesCancelIterator{contract: f.contract, event: "FuturesCancel", logs: logs, sub: sub}, nil
}

// WatchFuturesCancel is a free log subscription operation binding the contract event 0x7197f6513c22ccffd678f7d190c7c57116b046d6ea34e83c1ba7fd841f966ec4.
//
// Solidity: event FuturesCancel(bytes32 indexed subAccountID, bytes32 indexed orderHash, bytes32 indexed marketID, address baseCurrency, uint256 contractPrice, uint256 quantityFilled, bool isLong)
func (f *FuturesFilterer) WatchFuturesCancel(opts *bind.WatchOpts, sink chan<- *FuturesCancel, subAccountID [][32]byte, orderHash [][32]byte, marketID [][32]byte) (event.Subscription, error) {

	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesCancel", subAccountIDRule, orderHashRule, marketIDRule)
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

// ParseFuturesCancel is a log parse operation binding the contract event 0x7197f6513c22ccffd678f7d190c7c57116b046d6ea34e83c1ba7fd841f966ec4.
//
// Solidity: event FuturesCancel(bytes32 indexed subAccountID, bytes32 indexed orderHash, bytes32 indexed marketID, address baseCurrency, uint256 contractPrice, uint256 quantityFilled, bool isLong)
func (f *FuturesFilterer) ParseFuturesCancel(log types.Log) (*FuturesCancel, error) {
	event := new(FuturesCancel)
	if err := f.contract.UnpackLog(event, "FuturesCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesMarketCreationIterator is returned from FilterFuturesMarketCreation and is used to iterate over the raw logs and unpacked data for FuturesMarketCreation events raised by the Futures contract.
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

// FuturesMarketCreation represents a FuturesMarketCreation event raised by the Futures contract.
type FuturesMarketCreation struct {
	MarketID               [32]byte
	Ticker                 common.Hash
	Oracle                 common.Address
	BaseCurrency           common.Address
	MaintenanceMarginRatio *big.Int
	InitialMarginRatio     *big.Int
	MakerTxFee             *big.Int
	TakerTxFee             *big.Int
	RelayerFeePercentage   *big.Int
	FundingInterval        *big.Int
	ExpirationTime         *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFuturesMarketCreation is a free log retrieval operation binding the contract event 0xec7604a06b6f1d53d9f93d29b6df434191db13fab896b33d7ac38c636a5923c2.
//
// Solidity: event FuturesMarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle, address baseCurrency, uint256 maintenanceMarginRatio, uint256 initialMarginRatio, uint256 makerTxFee, uint256 takerTxFee, uint256 relayerFeePercentage, uint256 fundingInterval, uint256 expirationTime)
func (f *FuturesFilterer) FilterFuturesMarketCreation(opts *bind.FilterOpts, marketID [][32]byte, ticker []string, oracle []common.Address) (*FuturesMarketCreationIterator, error) {

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

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesMarketCreation", marketIDRule, tickerRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &FuturesMarketCreationIterator{contract: f.contract, event: "FuturesMarketCreation", logs: logs, sub: sub}, nil
}

// WatchFuturesMarketCreation is a free log subscription operation binding the contract event 0xec7604a06b6f1d53d9f93d29b6df434191db13fab896b33d7ac38c636a5923c2.
//
// Solidity: event FuturesMarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle, address baseCurrency, uint256 maintenanceMarginRatio, uint256 initialMarginRatio, uint256 makerTxFee, uint256 takerTxFee, uint256 relayerFeePercentage, uint256 fundingInterval, uint256 expirationTime)
func (f *FuturesFilterer) WatchFuturesMarketCreation(opts *bind.WatchOpts, sink chan<- *FuturesMarketCreation, marketID [][32]byte, ticker []string, oracle []common.Address) (event.Subscription, error) {

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

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesMarketCreation", marketIDRule, tickerRule, oracleRule)
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
				if err := f.contract.UnpackLog(event, "FuturesMarketCreation", log); err != nil {
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

// ParseFuturesMarketCreation is a log parse operation binding the contract event 0xec7604a06b6f1d53d9f93d29b6df434191db13fab896b33d7ac38c636a5923c2.
//
// Solidity: event FuturesMarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle, address baseCurrency, uint256 maintenanceMarginRatio, uint256 initialMarginRatio, uint256 makerTxFee, uint256 takerTxFee, uint256 relayerFeePercentage, uint256 fundingInterval, uint256 expirationTime)
func (f *FuturesFilterer) ParseFuturesMarketCreation(log types.Log) (*FuturesMarketCreation, error) {
	event := new(FuturesMarketCreation)
	if err := f.contract.UnpackLog(event, "FuturesMarketCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesTradeIterator is returned from FilterFuturesTrade and is used to iterate over the raw logs and unpacked data for FuturesTrade events raised by the Futures contract.
type FuturesTradeIterator struct {
	Event *FuturesTrade // Event containing the contract specifics and raw log

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
func (it *FuturesTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesTrade)
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
		it.Event = new(FuturesTrade)
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
func (it *FuturesTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesTrade represents a FuturesTrade event raised by the Futures contract.
type FuturesTrade struct {
	IsResultingPositionLong bool
	PositionID              *big.Int
	MarketID                [32]byte
	SubAccountID            [32]byte
	ResultingMargin         *big.Int
	ResultingEntryPrice     *big.Int
	ResultingQuantity       *big.Int
	CumulativeFundingEntry  *big.Int
	IsTradeDirectionLong    bool
	OrderHash               [32]byte
	ExecutionPrice          *big.Int
	ExecutionQuantity       *big.Int
	TotalForOrderFilled     *big.Int
	TradeType               uint8
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterFuturesTrade is a free log retrieval operation binding the contract event 0x2ee27d89e8394df57c5c3e6c07c03c128cc28e35450ee92fcdb9990932ec564e.
//
// Solidity: event FuturesTrade(bool isResultingPositionLong, uint256 positionID, bytes32 indexed marketID, bytes32 indexed subAccountID, uint256 resultingMargin, uint256 resultingEntryPrice, uint256 resultingQuantity, int256 cumulativeFundingEntry, bool isTradeDirectionLong, bytes32 indexed orderHash, uint256 executionPrice, uint256 executionQuantity, uint256 totalForOrderFilled, uint8 tradeType)
func (f *FuturesFilterer) FilterFuturesTrade(opts *bind.FilterOpts, marketID [][32]byte, subAccountID [][32]byte, orderHash [][32]byte) (*FuturesTradeIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesTrade", marketIDRule, subAccountIDRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &FuturesTradeIterator{contract: f.contract, event: "FuturesTrade", logs: logs, sub: sub}, nil
}

// WatchFuturesTrade is a free log subscription operation binding the contract event 0x2ee27d89e8394df57c5c3e6c07c03c128cc28e35450ee92fcdb9990932ec564e.
//
// Solidity: event FuturesTrade(bool isResultingPositionLong, uint256 positionID, bytes32 indexed marketID, bytes32 indexed subAccountID, uint256 resultingMargin, uint256 resultingEntryPrice, uint256 resultingQuantity, int256 cumulativeFundingEntry, bool isTradeDirectionLong, bytes32 indexed orderHash, uint256 executionPrice, uint256 executionQuantity, uint256 totalForOrderFilled, uint8 tradeType)
func (f *FuturesFilterer) WatchFuturesTrade(opts *bind.WatchOpts, sink chan<- *FuturesTrade, marketID [][32]byte, subAccountID [][32]byte, orderHash [][32]byte) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesTrade", marketIDRule, subAccountIDRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesTrade)
				if err := f.contract.UnpackLog(event, "FuturesTrade", log); err != nil {
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

// ParseFuturesTrade is a log parse operation binding the contract event 0x2ee27d89e8394df57c5c3e6c07c03c128cc28e35450ee92fcdb9990932ec564e.
//
// Solidity: event FuturesTrade(bool isResultingPositionLong, uint256 positionID, bytes32 indexed marketID, bytes32 indexed subAccountID, uint256 resultingMargin, uint256 resultingEntryPrice, uint256 resultingQuantity, int256 cumulativeFundingEntry, bool isTradeDirectionLong, bytes32 indexed orderHash, uint256 executionPrice, uint256 executionQuantity, uint256 totalForOrderFilled, uint8 tradeType)
func (f *FuturesFilterer) ParseFuturesTrade(log types.Log) (*FuturesTrade, error) {
	event := new(FuturesTrade)
	if err := f.contract.UnpackLog(event, "FuturesTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesMarginChangedIterator is returned from FilterMarginChanged and is used to iterate over the raw logs and unpacked data for MarginChanged events raised by the Futures contract.
type FuturesMarginChangedIterator struct {
	Event *FuturesMarginChanged // Event containing the contract specifics and raw log

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
func (it *FuturesMarginChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesMarginChanged)
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
		it.Event = new(FuturesMarginChanged)
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
func (it *FuturesMarginChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesMarginChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesMarginChanged represents a MarginChanged event raised by the Futures contract.
type FuturesMarginChanged struct {
	PositionID    *big.Int
	ChangedMargin *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarginChanged is a free log retrieval operation binding the contract event 0x6bbd00a8b71e37b424550ee42f7d18af2790fc0a874cbd296c6705205b298d00.
//
// Solidity: event MarginChanged(uint256 indexed positionID, int256 changedMargin)
func (f *FuturesFilterer) FilterMarginChanged(opts *bind.FilterOpts, positionID []*big.Int) (*FuturesMarginChangedIterator, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "MarginChanged", positionIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesMarginChangedIterator{contract: f.contract, event: "MarginChanged", logs: logs, sub: sub}, nil
}

// WatchMarginChanged is a free log subscription operation binding the contract event 0x6bbd00a8b71e37b424550ee42f7d18af2790fc0a874cbd296c6705205b298d00.
//
// Solidity: event MarginChanged(uint256 indexed positionID, int256 changedMargin)
func (f *FuturesFilterer) WatchMarginChanged(opts *bind.WatchOpts, sink chan<- *FuturesMarginChanged, positionID []*big.Int) (event.Subscription, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "MarginChanged", positionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesMarginChanged)
				if err := f.contract.UnpackLog(event, "MarginChanged", log); err != nil {
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

// ParseMarginChanged is a log parse operation binding the contract event 0x6bbd00a8b71e37b424550ee42f7d18af2790fc0a874cbd296c6705205b298d00.
//
// Solidity: event MarginChanged(uint256 indexed positionID, int256 changedMargin)
func (f *FuturesFilterer) ParseMarginChanged(log types.Log) (*FuturesMarginChanged, error) {
	event := new(FuturesMarginChanged)
	if err := f.contract.UnpackLog(event, "MarginChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// FuturesPositionTransferIterator is returned from FilterPositionTransfer and is used to iterate over the raw logs and unpacked data for PositionTransfer events raised by the Futures contract.
type FuturesPositionTransferIterator struct {
	Event *FuturesPositionTransfer // Event containing the contract specifics and raw log

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
func (it *FuturesPositionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesPositionTransfer)
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
		it.Event = new(FuturesPositionTransfer)
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
func (it *FuturesPositionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesPositionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesPositionTransfer represents a PositionTransfer event raised by the Futures contract.
type FuturesPositionTransfer struct {
	SenderPositionID                            *big.Int
	ReceiverPositionID                          *big.Int
	FromSubAccountID                            [32]byte
	ToSubAccountID                              [32]byte
	MarketID                                    [32]byte
	DirectionalStatus                           uint8
	SrcResultingPositionMargin                  *big.Int
	SrcResultingPositionEntryPrice              *big.Int
	SrcResultingPositionQuantity                *big.Int
	SrcResultingPositionCumulativeFundingEntry  *big.Int
	DestResultingPositionMargin                 *big.Int
	DestResultingPositionEntryPrice             *big.Int
	DestResultingPositionQuantity               *big.Int
	DestResultingPositionCumulativeFundingEntry *big.Int
	Raw                                         types.Log // Blockchain specific contextual infos
}

// FilterPositionTransfer is a free log retrieval operation binding the contract event 0xdea221d5d06ee8a5f9f089295d3292f2f89c054d223c4a3c452d4d90df510402.
//
// Solidity: event PositionTransfer(uint256 senderPositionID, uint256 receiverPositionID, bytes32 indexed fromSubAccountID, bytes32 indexed toSubAccountID, bytes32 marketID, uint8 directionalStatus, uint256 srcResultingPositionMargin, uint256 srcResultingPositionEntryPrice, uint256 srcResultingPositionQuantity, int256 srcResultingPositionCumulativeFundingEntry, uint256 destResultingPositionMargin, uint256 destResultingPositionEntryPrice, uint256 destResultingPositionQuantity, int256 destResultingPositionCumulativeFundingEntry)
func (f *FuturesFilterer) FilterPositionTransfer(opts *bind.FilterOpts, fromSubAccountID [][32]byte, toSubAccountID [][32]byte) (*FuturesPositionTransferIterator, error) {

	var fromSubAccountIDRule []interface{}
	for _, fromSubAccountIDItem := range fromSubAccountID {
		fromSubAccountIDRule = append(fromSubAccountIDRule, fromSubAccountIDItem)
	}
	var toSubAccountIDRule []interface{}
	for _, toSubAccountIDItem := range toSubAccountID {
		toSubAccountIDRule = append(toSubAccountIDRule, toSubAccountIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "PositionTransfer", fromSubAccountIDRule, toSubAccountIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesPositionTransferIterator{contract: f.contract, event: "PositionTransfer", logs: logs, sub: sub}, nil
}

// WatchPositionTransfer is a free log subscription operation binding the contract event 0xdea221d5d06ee8a5f9f089295d3292f2f89c054d223c4a3c452d4d90df510402.
//
// Solidity: event PositionTransfer(uint256 senderPositionID, uint256 receiverPositionID, bytes32 indexed fromSubAccountID, bytes32 indexed toSubAccountID, bytes32 marketID, uint8 directionalStatus, uint256 srcResultingPositionMargin, uint256 srcResultingPositionEntryPrice, uint256 srcResultingPositionQuantity, int256 srcResultingPositionCumulativeFundingEntry, uint256 destResultingPositionMargin, uint256 destResultingPositionEntryPrice, uint256 destResultingPositionQuantity, int256 destResultingPositionCumulativeFundingEntry)
func (f *FuturesFilterer) WatchPositionTransfer(opts *bind.WatchOpts, sink chan<- *FuturesPositionTransfer, fromSubAccountID [][32]byte, toSubAccountID [][32]byte) (event.Subscription, error) {

	var fromSubAccountIDRule []interface{}
	for _, fromSubAccountIDItem := range fromSubAccountID {
		fromSubAccountIDRule = append(fromSubAccountIDRule, fromSubAccountIDItem)
	}
	var toSubAccountIDRule []interface{}
	for _, toSubAccountIDItem := range toSubAccountID {
		toSubAccountIDRule = append(toSubAccountIDRule, toSubAccountIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "PositionTransfer", fromSubAccountIDRule, toSubAccountIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesPositionTransfer)
				if err := f.contract.UnpackLog(event, "PositionTransfer", log); err != nil {
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

// ParsePositionTransfer is a log parse operation binding the contract event 0xdea221d5d06ee8a5f9f089295d3292f2f89c054d223c4a3c452d4d90df510402.
//
// Solidity: event PositionTransfer(uint256 senderPositionID, uint256 receiverPositionID, bytes32 indexed fromSubAccountID, bytes32 indexed toSubAccountID, bytes32 marketID, uint8 directionalStatus, uint256 srcResultingPositionMargin, uint256 srcResultingPositionEntryPrice, uint256 srcResultingPositionQuantity, int256 srcResultingPositionCumulativeFundingEntry, uint256 destResultingPositionMargin, uint256 destResultingPositionEntryPrice, uint256 destResultingPositionQuantity, int256 destResultingPositionCumulativeFundingEntry)
func (f *FuturesFilterer) ParsePositionTransfer(log types.Log) (*FuturesPositionTransfer, error) {
	event := new(FuturesPositionTransfer)
	if err := f.contract.UnpackLog(event, "PositionTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	MarketID                     [32]byte
	FundingFee                   *big.Int
	CumulativeFunding            *big.Int
	ExpiryOrNextFundingTimestamp *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterSetFunding is a free log retrieval operation binding the contract event 0x4f78e101b9f3e3fbc0cab373ac4f3bc7f68f50610c03844ec1293dd9cd78a35d.
//
// Solidity: event SetFunding(bytes32 indexed marketID, int256 fundingFee, int256 cumulativeFunding, uint256 expiryOrNextFundingTimestamp)
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

// WatchSetFunding is a free log subscription operation binding the contract event 0x4f78e101b9f3e3fbc0cab373ac4f3bc7f68f50610c03844ec1293dd9cd78a35d.
//
// Solidity: event SetFunding(bytes32 indexed marketID, int256 fundingFee, int256 cumulativeFunding, uint256 expiryOrNextFundingTimestamp)
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

// ParseSetFunding is a log parse operation binding the contract event 0x4f78e101b9f3e3fbc0cab373ac4f3bc7f68f50610c03844ec1293dd9cd78a35d.
//
// Solidity: event SetFunding(bytes32 indexed marketID, int256 fundingFee, int256 cumulativeFunding, uint256 expiryOrNextFundingTimestamp)
func (f *FuturesFilterer) ParseSetFunding(log types.Log) (*FuturesSetFunding, error) {
	event := new(FuturesSetFunding)
	if err := f.contract.UnpackLog(event, "SetFunding", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// FuturesSpotMarketCreationIterator is returned from FilterSpotMarketCreation and is used to iterate over the raw logs and unpacked data for SpotMarketCreation events raised by the Futures contract.
type FuturesSpotMarketCreationIterator struct {
	Event *FuturesSpotMarketCreation // Event containing the contract specifics and raw log

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
func (it *FuturesSpotMarketCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesSpotMarketCreation)
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
		it.Event = new(FuturesSpotMarketCreation)
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
func (it *FuturesSpotMarketCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesSpotMarketCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesSpotMarketCreation represents a SpotMarketCreation event raised by the Futures contract.
type FuturesSpotMarketCreation struct {
	MarketID             [32]byte
	Ticker               common.Hash
	BaseToken            common.Address
	QuoteToken           common.Address
	MakerTxFee           *big.Int
	TakerTxFee           *big.Int
	RelayerFeePercentage *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSpotMarketCreation is a free log retrieval operation binding the contract event 0x36e5492cb554023c13d9df48b0fa6b166ccd07ea459dd0098d89052bf0775f7a.
//
// Solidity: event SpotMarketCreation(bytes32 indexed marketID, string indexed ticker, address baseToken, address quoteToken, uint256 makerTxFee, uint256 takerTxFee, uint256 relayerFeePercentage)
func (f *FuturesFilterer) FilterSpotMarketCreation(opts *bind.FilterOpts, marketID [][32]byte, ticker []string) (*FuturesSpotMarketCreationIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "SpotMarketCreation", marketIDRule, tickerRule)
	if err != nil {
		return nil, err
	}
	return &FuturesSpotMarketCreationIterator{contract: f.contract, event: "SpotMarketCreation", logs: logs, sub: sub}, nil
}

// WatchSpotMarketCreation is a free log subscription operation binding the contract event 0x36e5492cb554023c13d9df48b0fa6b166ccd07ea459dd0098d89052bf0775f7a.
//
// Solidity: event SpotMarketCreation(bytes32 indexed marketID, string indexed ticker, address baseToken, address quoteToken, uint256 makerTxFee, uint256 takerTxFee, uint256 relayerFeePercentage)
func (f *FuturesFilterer) WatchSpotMarketCreation(opts *bind.WatchOpts, sink chan<- *FuturesSpotMarketCreation, marketID [][32]byte, ticker []string) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "SpotMarketCreation", marketIDRule, tickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesSpotMarketCreation)
				if err := f.contract.UnpackLog(event, "SpotMarketCreation", log); err != nil {
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

// ParseSpotMarketCreation is a log parse operation binding the contract event 0x36e5492cb554023c13d9df48b0fa6b166ccd07ea459dd0098d89052bf0775f7a.
//
// Solidity: event SpotMarketCreation(bytes32 indexed marketID, string indexed ticker, address baseToken, address quoteToken, uint256 makerTxFee, uint256 takerTxFee, uint256 relayerFeePercentage)
func (f *FuturesFilterer) ParseSpotMarketCreation(log types.Log) (*FuturesSpotMarketCreation, error) {
	event := new(FuturesSpotMarketCreation)
	if err := f.contract.UnpackLog(event, "SpotMarketCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesSpotTradeIterator is returned from FilterSpotTrade and is used to iterate over the raw logs and unpacked data for SpotTrade events raised by the Futures contract.
type FuturesSpotTradeIterator struct {
	Event *FuturesSpotTrade // Event containing the contract specifics and raw log

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
func (it *FuturesSpotTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesSpotTrade)
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
		it.Event = new(FuturesSpotTrade)
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
func (it *FuturesSpotTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesSpotTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesSpotTrade represents a SpotTrade event raised by the Futures contract.
type FuturesSpotTrade struct {
	MarketID                   [32]byte
	SubAccountID               [32]byte
	OrderHash                  [32]byte
	SentAsset                  common.Address
	GainedAsset                common.Address
	SentAmount                 *big.Int
	GainedAmount               *big.Int
	CurrentSentAssetDeposits   *big.Int
	CurrentGainedAssetDeposits *big.Int
	TotalForOrderFilled        *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSpotTrade is a free log retrieval operation binding the contract event 0xf0283823dfd5d2cf09704eb9d26d1f5098a512c6d7a7b4a74ca0927230c27699.
//
// Solidity: event SpotTrade(bytes32 indexed marketID, bytes32 indexed subAccountID, bytes32 indexed orderHash, address sentAsset, address gainedAsset, uint256 sentAmount, uint256 gainedAmount, uint256 currentSentAssetDeposits, uint256 currentGainedAssetDeposits, uint256 totalForOrderFilled)
func (f *FuturesFilterer) FilterSpotTrade(opts *bind.FilterOpts, marketID [][32]byte, subAccountID [][32]byte, orderHash [][32]byte) (*FuturesSpotTradeIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "SpotTrade", marketIDRule, subAccountIDRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &FuturesSpotTradeIterator{contract: f.contract, event: "SpotTrade", logs: logs, sub: sub}, nil
}

// WatchSpotTrade is a free log subscription operation binding the contract event 0xf0283823dfd5d2cf09704eb9d26d1f5098a512c6d7a7b4a74ca0927230c27699.
//
// Solidity: event SpotTrade(bytes32 indexed marketID, bytes32 indexed subAccountID, bytes32 indexed orderHash, address sentAsset, address gainedAsset, uint256 sentAmount, uint256 gainedAmount, uint256 currentSentAssetDeposits, uint256 currentGainedAssetDeposits, uint256 totalForOrderFilled)
func (f *FuturesFilterer) WatchSpotTrade(opts *bind.WatchOpts, sink chan<- *FuturesSpotTrade, marketID [][32]byte, subAccountID [][32]byte, orderHash [][32]byte) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "SpotTrade", marketIDRule, subAccountIDRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesSpotTrade)
				if err := f.contract.UnpackLog(event, "SpotTrade", log); err != nil {
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

// ParseSpotTrade is a log parse operation binding the contract event 0xf0283823dfd5d2cf09704eb9d26d1f5098a512c6d7a7b4a74ca0927230c27699.
//
// Solidity: event SpotTrade(bytes32 indexed marketID, bytes32 indexed subAccountID, bytes32 indexed orderHash, address sentAsset, address gainedAsset, uint256 sentAmount, uint256 gainedAmount, uint256 currentSentAssetDeposits, uint256 currentGainedAssetDeposits, uint256 totalForOrderFilled)
func (f *FuturesFilterer) ParseSpotTrade(log types.Log) (*FuturesSpotTrade, error) {
	event := new(FuturesSpotTrade)
	if err := f.contract.UnpackLog(event, "SpotTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesSubaccountDepositsChangeIterator is returned from FilterSubaccountDepositsChange and is used to iterate over the raw logs and unpacked data for SubaccountDepositsChange events raised by the Futures contract.
type FuturesSubaccountDepositsChangeIterator struct {
	Event *FuturesSubaccountDepositsChange // Event containing the contract specifics and raw log

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
func (it *FuturesSubaccountDepositsChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesSubaccountDepositsChange)
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
		it.Event = new(FuturesSubaccountDepositsChange)
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
func (it *FuturesSubaccountDepositsChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesSubaccountDepositsChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesSubaccountDepositsChange represents a SubaccountDepositsChange event raised by the Futures contract.
type FuturesSubaccountDepositsChange struct {
	DepositChangeType uint8
	SubAccountID      [32]byte
	BaseCurrency      common.Address
	ChangeAmount      *big.Int
	CurrentAmount     *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSubaccountDepositsChange is a free log retrieval operation binding the contract event 0xdf3bc6fae0ce8484ef20d7d7c446de8b6aca50d03505e4ca06629917d8d1bde6.
//
// Solidity: event SubaccountDepositsChange(uint8 indexed depositChangeType, bytes32 indexed subAccountID, address indexed baseCurrency, int256 changeAmount, uint256 currentAmount)
func (f *FuturesFilterer) FilterSubaccountDepositsChange(opts *bind.FilterOpts, depositChangeType []uint8, subAccountID [][32]byte, baseCurrency []common.Address) (*FuturesSubaccountDepositsChangeIterator, error) {

	var depositChangeTypeRule []interface{}
	for _, depositChangeTypeItem := range depositChangeType {
		depositChangeTypeRule = append(depositChangeTypeRule, depositChangeTypeItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}
	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "SubaccountDepositsChange", depositChangeTypeRule, subAccountIDRule, baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return &FuturesSubaccountDepositsChangeIterator{contract: f.contract, event: "SubaccountDepositsChange", logs: logs, sub: sub}, nil
}

// WatchSubaccountDepositsChange is a free log subscription operation binding the contract event 0xdf3bc6fae0ce8484ef20d7d7c446de8b6aca50d03505e4ca06629917d8d1bde6.
//
// Solidity: event SubaccountDepositsChange(uint8 indexed depositChangeType, bytes32 indexed subAccountID, address indexed baseCurrency, int256 changeAmount, uint256 currentAmount)
func (f *FuturesFilterer) WatchSubaccountDepositsChange(opts *bind.WatchOpts, sink chan<- *FuturesSubaccountDepositsChange, depositChangeType []uint8, subAccountID [][32]byte, baseCurrency []common.Address) (event.Subscription, error) {

	var depositChangeTypeRule []interface{}
	for _, depositChangeTypeItem := range depositChangeType {
		depositChangeTypeRule = append(depositChangeTypeRule, depositChangeTypeItem)
	}
	var subAccountIDRule []interface{}
	for _, subAccountIDItem := range subAccountID {
		subAccountIDRule = append(subAccountIDRule, subAccountIDItem)
	}
	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "SubaccountDepositsChange", depositChangeTypeRule, subAccountIDRule, baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesSubaccountDepositsChange)
				if err := f.contract.UnpackLog(event, "SubaccountDepositsChange", log); err != nil {
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

// ParseSubaccountDepositsChange is a log parse operation binding the contract event 0xdf3bc6fae0ce8484ef20d7d7c446de8b6aca50d03505e4ca06629917d8d1bde6.
//
// Solidity: event SubaccountDepositsChange(uint8 indexed depositChangeType, bytes32 indexed subAccountID, address indexed baseCurrency, int256 changeAmount, uint256 currentAmount)
func (f *FuturesFilterer) ParseSubaccountDepositsChange(log types.Log) (*FuturesSubaccountDepositsChange, error) {
	event := new(FuturesSubaccountDepositsChange)
	if err := f.contract.UnpackLog(event, "SubaccountDepositsChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesTransactionExecutionIterator is returned from FilterTransactionExecution and is used to iterate over the raw logs and unpacked data for TransactionExecution events raised by the Futures contract.
type FuturesTransactionExecutionIterator struct {
	Event *FuturesTransactionExecution // Event containing the contract specifics and raw log

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
func (it *FuturesTransactionExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesTransactionExecution)
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
		it.Event = new(FuturesTransactionExecution)
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
func (it *FuturesTransactionExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesTransactionExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesTransactionExecution represents a TransactionExecution event raised by the Futures contract.
type FuturesTransactionExecution struct {
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransactionExecution is a free log retrieval operation binding the contract event 0xa4a7329f1dd821363067e07d359e347b4af9b1efe4b6cccf13240228af3c800d.
//
// Solidity: event TransactionExecution(bytes32 indexed transactionHash)
func (f *FuturesFilterer) FilterTransactionExecution(opts *bind.FilterOpts, transactionHash [][32]byte) (*FuturesTransactionExecutionIterator, error) {

	var transactionHashRule []interface{}
	for _, transactionHashItem := range transactionHash {
		transactionHashRule = append(transactionHashRule, transactionHashItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "TransactionExecution", transactionHashRule)
	if err != nil {
		return nil, err
	}
	return &FuturesTransactionExecutionIterator{contract: f.contract, event: "TransactionExecution", logs: logs, sub: sub}, nil
}

// WatchTransactionExecution is a free log subscription operation binding the contract event 0xa4a7329f1dd821363067e07d359e347b4af9b1efe4b6cccf13240228af3c800d.
//
// Solidity: event TransactionExecution(bytes32 indexed transactionHash)
func (f *FuturesFilterer) WatchTransactionExecution(opts *bind.WatchOpts, sink chan<- *FuturesTransactionExecution, transactionHash [][32]byte) (event.Subscription, error) {

	var transactionHashRule []interface{}
	for _, transactionHashItem := range transactionHash {
		transactionHashRule = append(transactionHashRule, transactionHashItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "TransactionExecution", transactionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesTransactionExecution)
				if err := f.contract.UnpackLog(event, "TransactionExecution", log); err != nil {
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

// ParseTransactionExecution is a log parse operation binding the contract event 0xa4a7329f1dd821363067e07d359e347b4af9b1efe4b6cccf13240228af3c800d.
//
// Solidity: event TransactionExecution(bytes32 indexed transactionHash)
func (f *FuturesFilterer) ParseTransactionExecution(log types.Log) (*FuturesTransactionExecution, error) {
	event := new(FuturesTransactionExecution)
	if err := f.contract.UnpackLog(event, "TransactionExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// FuturesUpdateValuesForVWAPIterator is returned from FilterUpdateValuesForVWAP and is used to iterate over the raw logs and unpacked data for UpdateValuesForVWAP events raised by the Futures contract.
type FuturesUpdateValuesForVWAPIterator struct {
	Event *FuturesUpdateValuesForVWAP // Event containing the contract specifics and raw log

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
func (it *FuturesUpdateValuesForVWAPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesUpdateValuesForVWAP)
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
		it.Event = new(FuturesUpdateValuesForVWAP)
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
func (it *FuturesUpdateValuesForVWAPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesUpdateValuesForVWAPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesUpdateValuesForVWAP represents a UpdateValuesForVWAP event raised by the Futures contract.
type FuturesUpdateValuesForVWAP struct {
	MarketID                               [32]byte
	MostRecentEpochVolume                  *big.Int
	MostRecentEpochQuantity                *big.Int
	MostRecentEpochScaledContractIndexDiff *big.Int
	Raw                                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateValuesForVWAP is a free log retrieval operation binding the contract event 0x88459a5f4fa75a54716457659c83594066f178d0df407b402d59c5ed0d643351.
//
// Solidity: event UpdateValuesForVWAP(bytes32 indexed marketID, uint256 mostRecentEpochVolume, uint256 mostRecentEpochQuantity, int256 mostRecentEpochScaledContractIndexDiff)
func (f *FuturesFilterer) FilterUpdateValuesForVWAP(opts *bind.FilterOpts, marketID [][32]byte) (*FuturesUpdateValuesForVWAPIterator, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "UpdateValuesForVWAP", marketIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesUpdateValuesForVWAPIterator{contract: f.contract, event: "UpdateValuesForVWAP", logs: logs, sub: sub}, nil
}

// WatchUpdateValuesForVWAP is a free log subscription operation binding the contract event 0x88459a5f4fa75a54716457659c83594066f178d0df407b402d59c5ed0d643351.
//
// Solidity: event UpdateValuesForVWAP(bytes32 indexed marketID, uint256 mostRecentEpochVolume, uint256 mostRecentEpochQuantity, int256 mostRecentEpochScaledContractIndexDiff)
func (f *FuturesFilterer) WatchUpdateValuesForVWAP(opts *bind.WatchOpts, sink chan<- *FuturesUpdateValuesForVWAP, marketID [][32]byte) (event.Subscription, error) {

	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "UpdateValuesForVWAP", marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesUpdateValuesForVWAP)
				if err := f.contract.UnpackLog(event, "UpdateValuesForVWAP", log); err != nil {
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

// ParseUpdateValuesForVWAP is a log parse operation binding the contract event 0x88459a5f4fa75a54716457659c83594066f178d0df407b402d59c5ed0d643351.
//
// Solidity: event UpdateValuesForVWAP(bytes32 indexed marketID, uint256 mostRecentEpochVolume, uint256 mostRecentEpochQuantity, int256 mostRecentEpochScaledContractIndexDiff)
func (f *FuturesFilterer) ParseUpdateValuesForVWAP(log types.Log) (*FuturesUpdateValuesForVWAP, error) {
	event := new(FuturesUpdateValuesForVWAP)
	if err := f.contract.UnpackLog(event, "UpdateValuesForVWAP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
