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

// TypesPosition is an auto generated low-level Go binding around an user-defined struct.
type TypesPosition struct {
	AccountID              [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	ContractPrice          *big.Int
	NPV                    *big.Int
	MinMargin              *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
}

// FuturesABI is the input ABI used to generate the binding from.
const FuturesABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumMarginRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContractAddressIfExists\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"}],\"name\":\"FuturesCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"FuturesClose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"FuturesFill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"MarketCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"SignatureValidatorApproval\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP1271_MAGIC_VALUE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_EXCHANGE_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accountIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"NAV\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"NPV\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressToAccountIDs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"name\":\"calcCumulativeFunding\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"calcLiquidationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"calcMinMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"closePosition\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"closePositionWithOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createAccountAndDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"}],\"name\":\"createMarket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositNewAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freeDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"fundPooledDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"getBatchBalancesAndAssetProxyAllowances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"getOrderRelevantStates\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.OrderStatus\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAssetFilledAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLibOrder.OrderInfo[]\",\"name\":\"ordersInfo\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fillableTakerAssetAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"isValidSignature\",\"type\":\"bool[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getPositionIDsForTrader\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"positionIDs\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getPositionsForTrader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"NPV\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"internalType\":\"structTypes.Position[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"getTransferableAssetAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transferableAssetAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidOrderSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"marketOrders\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketSerialToID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"marketToAccountToPositionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currFundingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFunding\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"leftOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"leftSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"matchOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"leftOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"leftSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"multiMatchOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pooledDeposits\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"positionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"NPV\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"preSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"restrictedDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unrestrictedDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifyClose\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"NPV\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"internalType\":\"structTypes.Position\",\"name\":\"newPosition\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
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
func (_Futures *FuturesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Futures.Contract.FuturesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Futures *FuturesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Futures.Contract.FuturesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Futures *FuturesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Futures.Contract.FuturesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Futures *FuturesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Futures.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Futures *FuturesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Futures.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Futures *FuturesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Futures.Contract.contract.Transact(opts, method, params...)
}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (_Futures *FuturesCaller) EIP1271MAGICVALUE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "EIP1271_MAGIC_VALUE")
	return *ret0, err
}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (_Futures *FuturesSession) EIP1271MAGICVALUE() ([4]byte, error) {
	return _Futures.Contract.EIP1271MAGICVALUE(&_Futures.CallOpts)
}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (_Futures *FuturesCallerSession) EIP1271MAGICVALUE() ([4]byte, error) {
	return _Futures.Contract.EIP1271MAGICVALUE(&_Futures.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (_Futures *FuturesCaller) EIP712EXCHANGEDOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "EIP712_EXCHANGE_DOMAIN_HASH")
	return *ret0, err
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (_Futures *FuturesSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return _Futures.Contract.EIP712EXCHANGEDOMAINHASH(&_Futures.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (_Futures *FuturesCallerSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return _Futures.Contract.EIP712EXCHANGEDOMAINHASH(&_Futures.CallOpts)
}

// AccountIdToAddress is a free data retrieval call binding the contract method 0x899706c0.
//
// Solidity: function accountIdToAddress(bytes32 ) view returns(address)
func (_Futures *FuturesCaller) AccountIdToAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "accountIdToAddress", arg0)
	return *ret0, err
}

// AccountIdToAddress is a free data retrieval call binding the contract method 0x899706c0.
//
// Solidity: function accountIdToAddress(bytes32 ) view returns(address)
func (_Futures *FuturesSession) AccountIdToAddress(arg0 [32]byte) (common.Address, error) {
	return _Futures.Contract.AccountIdToAddress(&_Futures.CallOpts, arg0)
}

// AccountIdToAddress is a free data retrieval call binding the contract method 0x899706c0.
//
// Solidity: function accountIdToAddress(bytes32 ) view returns(address)
func (_Futures *FuturesCallerSession) AccountIdToAddress(arg0 [32]byte) (common.Address, error) {
	return _Futures.Contract.AccountIdToAddress(&_Futures.CallOpts, arg0)
}

// AccountNonce is a free data retrieval call binding the contract method 0x106cffe2.
//
// Solidity: function accountNonce(address ) view returns(uint256)
func (_Futures *FuturesCaller) AccountNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "accountNonce", arg0)
	return *ret0, err
}

// AccountNonce is a free data retrieval call binding the contract method 0x106cffe2.
//
// Solidity: function accountNonce(address ) view returns(uint256)
func (_Futures *FuturesSession) AccountNonce(arg0 common.Address) (*big.Int, error) {
	return _Futures.Contract.AccountNonce(&_Futures.CallOpts, arg0)
}

// AccountNonce is a free data retrieval call binding the contract method 0x106cffe2.
//
// Solidity: function accountNonce(address ) view returns(uint256)
func (_Futures *FuturesCallerSession) AccountNonce(arg0 common.Address) (*big.Int, error) {
	return _Futures.Contract.AccountNonce(&_Futures.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(int256 NAV, int256 NPV, bytes32 accountID)
func (_Futures *FuturesCaller) Accounts(opts *bind.CallOpts, arg0 [32]byte) (struct {
	NAV       *big.Int
	NPV       *big.Int
	AccountID [32]byte
}, error) {
	ret := new(struct {
		NAV       *big.Int
		NPV       *big.Int
		AccountID [32]byte
	})
	out := ret
	err := _Futures.contract.Call(opts, out, "accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(int256 NAV, int256 NPV, bytes32 accountID)
func (_Futures *FuturesSession) Accounts(arg0 [32]byte) (struct {
	NAV       *big.Int
	NPV       *big.Int
	AccountID [32]byte
}, error) {
	return _Futures.Contract.Accounts(&_Futures.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(int256 NAV, int256 NPV, bytes32 accountID)
func (_Futures *FuturesCallerSession) Accounts(arg0 [32]byte) (struct {
	NAV       *big.Int
	NPV       *big.Int
	AccountID [32]byte
}, error) {
	return _Futures.Contract.Accounts(&_Futures.CallOpts, arg0)
}

// AddressToAccountIDs is a free data retrieval call binding the contract method 0xfbd6d494.
//
// Solidity: function addressToAccountIDs(address , uint256 ) view returns(bytes32)
func (_Futures *FuturesCaller) AddressToAccountIDs(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "addressToAccountIDs", arg0, arg1)
	return *ret0, err
}

// AddressToAccountIDs is a free data retrieval call binding the contract method 0xfbd6d494.
//
// Solidity: function addressToAccountIDs(address , uint256 ) view returns(bytes32)
func (_Futures *FuturesSession) AddressToAccountIDs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Futures.Contract.AddressToAccountIDs(&_Futures.CallOpts, arg0, arg1)
}

// AddressToAccountIDs is a free data retrieval call binding the contract method 0xfbd6d494.
//
// Solidity: function addressToAccountIDs(address , uint256 ) view returns(bytes32)
func (_Futures *FuturesCallerSession) AddressToAccountIDs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Futures.Contract.AddressToAccountIDs(&_Futures.CallOpts, arg0, arg1)
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) view returns(bool)
func (_Futures *FuturesCaller) AllowedValidators(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "allowedValidators", arg0, arg1)
	return *ret0, err
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) view returns(bool)
func (_Futures *FuturesSession) AllowedValidators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Futures.Contract.AllowedValidators(&_Futures.CallOpts, arg0, arg1)
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) view returns(bool)
func (_Futures *FuturesCallerSession) AllowedValidators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Futures.Contract.AllowedValidators(&_Futures.CallOpts, arg0, arg1)
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) view returns(int256)
func (_Futures *FuturesCaller) CalcCumulativeFunding(opts *bind.CallOpts, marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "calcCumulativeFunding", marketID, cumulativeFundingEntry)
	return *ret0, err
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) view returns(int256)
func (_Futures *FuturesSession) CalcCumulativeFunding(marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	return _Futures.Contract.CalcCumulativeFunding(&_Futures.CallOpts, marketID, cumulativeFundingEntry)
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) view returns(int256)
func (_Futures *FuturesCallerSession) CalcCumulativeFunding(marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	return _Futures.Contract.CalcCumulativeFunding(&_Futures.CallOpts, marketID, cumulativeFundingEntry)
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 _marketID, uint256 _quantity) view returns(uint256)
func (_Futures *FuturesCaller) CalcLiquidationFee(opts *bind.CallOpts, _marketID [32]byte, _quantity *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "calcLiquidationFee", _marketID, _quantity)
	return *ret0, err
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 _marketID, uint256 _quantity) view returns(uint256)
func (_Futures *FuturesSession) CalcLiquidationFee(_marketID [32]byte, _quantity *big.Int) (*big.Int, error) {
	return _Futures.Contract.CalcLiquidationFee(&_Futures.CallOpts, _marketID, _quantity)
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 _marketID, uint256 _quantity) view returns(uint256)
func (_Futures *FuturesCallerSession) CalcLiquidationFee(_marketID [32]byte, _quantity *big.Int) (*big.Int, error) {
	return _Futures.Contract.CalcLiquidationFee(&_Futures.CallOpts, _marketID, _quantity)
}

// CalcMinMargin is a free data retrieval call binding the contract method 0x35c43c4e.
//
// Solidity: function calcMinMargin(bytes32 marketID, uint256 quantity, uint256 price) view returns(uint256)
func (_Futures *FuturesCaller) CalcMinMargin(opts *bind.CallOpts, marketID [32]byte, quantity *big.Int, price *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "calcMinMargin", marketID, quantity, price)
	return *ret0, err
}

// CalcMinMargin is a free data retrieval call binding the contract method 0x35c43c4e.
//
// Solidity: function calcMinMargin(bytes32 marketID, uint256 quantity, uint256 price) view returns(uint256)
func (_Futures *FuturesSession) CalcMinMargin(marketID [32]byte, quantity *big.Int, price *big.Int) (*big.Int, error) {
	return _Futures.Contract.CalcMinMargin(&_Futures.CallOpts, marketID, quantity, price)
}

// CalcMinMargin is a free data retrieval call binding the contract method 0x35c43c4e.
//
// Solidity: function calcMinMargin(bytes32 marketID, uint256 quantity, uint256 price) view returns(uint256)
func (_Futures *FuturesCallerSession) CalcMinMargin(marketID [32]byte, quantity *big.Int, price *big.Int) (*big.Int, error) {
	return _Futures.Contract.CalcMinMargin(&_Futures.CallOpts, marketID, quantity, price)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Futures *FuturesCaller) Cancelled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "cancelled", arg0)
	return *ret0, err
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Futures *FuturesSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Futures.Contract.Cancelled(&_Futures.CallOpts, arg0)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Futures *FuturesCallerSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Futures.Contract.Cancelled(&_Futures.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Futures *FuturesCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "filled", arg0)
	return *ret0, err
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Futures *FuturesSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.Filled(&_Futures.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Futures *FuturesCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.Filled(&_Futures.CallOpts, arg0)
}

// FreeDeposits is a free data retrieval call binding the contract method 0x9a78538f.
//
// Solidity: function freeDeposits(address ) view returns(uint256)
func (_Futures *FuturesCaller) FreeDeposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "freeDeposits", arg0)
	return *ret0, err
}

// FreeDeposits is a free data retrieval call binding the contract method 0x9a78538f.
//
// Solidity: function freeDeposits(address ) view returns(uint256)
func (_Futures *FuturesSession) FreeDeposits(arg0 common.Address) (*big.Int, error) {
	return _Futures.Contract.FreeDeposits(&_Futures.CallOpts, arg0)
}

// FreeDeposits is a free data retrieval call binding the contract method 0x9a78538f.
//
// Solidity: function freeDeposits(address ) view returns(uint256)
func (_Futures *FuturesCallerSession) FreeDeposits(arg0 common.Address) (*big.Int, error) {
	return _Futures.Contract.FreeDeposits(&_Futures.CallOpts, arg0)
}

// GetBatchBalancesAndAssetProxyAllowances is a free data retrieval call binding the contract method 0x5bf34fc7.
//
// Solidity: function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, address[] tokenAddresses) view returns(uint256[], uint256[])
func (_Futures *FuturesCaller) GetBatchBalancesAndAssetProxyAllowances(opts *bind.CallOpts, ownerAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Futures.contract.Call(opts, out, "getBatchBalancesAndAssetProxyAllowances", ownerAddress, tokenAddresses)
	return *ret0, *ret1, err
}

// GetBatchBalancesAndAssetProxyAllowances is a free data retrieval call binding the contract method 0x5bf34fc7.
//
// Solidity: function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, address[] tokenAddresses) view returns(uint256[], uint256[])
func (_Futures *FuturesSession) GetBatchBalancesAndAssetProxyAllowances(ownerAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, []*big.Int, error) {
	return _Futures.Contract.GetBatchBalancesAndAssetProxyAllowances(&_Futures.CallOpts, ownerAddress, tokenAddresses)
}

// GetBatchBalancesAndAssetProxyAllowances is a free data retrieval call binding the contract method 0x5bf34fc7.
//
// Solidity: function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, address[] tokenAddresses) view returns(uint256[], uint256[])
func (_Futures *FuturesCallerSession) GetBatchBalancesAndAssetProxyAllowances(ownerAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, []*big.Int, error) {
	return _Futures.Contract.GetBatchBalancesAndAssetProxyAllowances(&_Futures.CallOpts, ownerAddress, tokenAddresses)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates([]Order orders, bytes[] signatures) constant returns([]OrderInfo ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (_Futures *FuturesCaller) GetOrderRelevantStates(opts *bind.CallOpts, orders []Order, signatures [][]byte) (struct {
	OrdersInfo                []OrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	ret := new(struct {
		OrdersInfo                []OrderInfo
		FillableTakerAssetAmounts []*big.Int
		IsValidSignature          []bool
	})
	out := ret
	err := _Futures.contract.Call(opts, out, "getOrderRelevantStates", orders, signatures)
	return *ret, err
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates([]Order orders, bytes[] signatures) constant returns([]OrderInfo ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (_Futures *FuturesSession) GetOrderRelevantStates(orders []Order, signatures [][]byte) (struct {
	OrdersInfo                []OrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	return _Futures.Contract.GetOrderRelevantStates(&_Futures.CallOpts, orders, signatures)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates([]Order orders, bytes[] signatures) constant returns([]OrderInfo ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (_Futures *FuturesCallerSession) GetOrderRelevantStates(orders []Order, signatures [][]byte) (struct {
	OrdersInfo                []OrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	return _Futures.Contract.GetOrderRelevantStates(&_Futures.CallOpts, orders, signatures)
}

// GetTransferableAssetAmount is a free data retrieval call binding the contract method 0x478482d8.
//
// Solidity: function getTransferableAssetAmount(address ownerAddress) constant returns(uint256 transferableAssetAmount)
func (_Futures *FuturesCaller) GetTransferableAssetAmount(opts *bind.CallOpts, ownerAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "getTransferableAssetAmount", ownerAddress)
	return *ret0, err
}

// GetTransferableAssetAmount is a free data retrieval call binding the contract method 0x478482d8.
//
// Solidity: function getTransferableAssetAmount(address ownerAddress) constant returns(uint256 transferableAssetAmount)
func (_Futures *FuturesSession) GetTransferableAssetAmount(ownerAddress common.Address) (*big.Int, error) {
	return _Futures.Contract.GetTransferableAssetAmount(&_Futures.CallOpts, ownerAddress)
}

// GetTransferableAssetAmount is a free data retrieval call binding the contract method 0x478482d8.
//
// Solidity: function getTransferableAssetAmount(address ownerAddress) constant returns(uint256 transferableAssetAmount)
func (_Futures *FuturesCallerSession) GetTransferableAssetAmount(ownerAddress common.Address) (*big.Int, error) {
	return _Futures.Contract.GetTransferableAssetAmount(&_Futures.CallOpts, ownerAddress)
}


// GetPositionIDsForTrader is a free data retrieval call binding the contract method 0xc3e49bb7.
//
// Solidity: function getPositionIDsForTrader(address trader, bytes32 marketID) constant returns(uint256[] positionIDs)
func (_Futures *FuturesCaller) GetPositionIDsForTrader(opts *bind.CallOpts, trader common.Address, marketID [32]byte) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "getPositionIDsForTrader", trader, marketID)
	return *ret0, err
}

// GetPositionIDsForTrader is a free data retrieval call binding the contract method 0xc3e49bb7.
//
// Solidity: function getPositionIDsForTrader(address trader, bytes32 marketID) constant returns(uint256[] positionIDs)
func (_Futures *FuturesSession) GetPositionIDsForTrader(trader common.Address, marketID [32]byte) ([]*big.Int, error) {
	return _Futures.Contract.GetPositionIDsForTrader(&_Futures.CallOpts, trader, marketID)
}

// GetPositionIDsForTrader is a free data retrieval call binding the contract method 0xc3e49bb7.
//
// Solidity: function getPositionIDsForTrader(address trader, bytes32 marketID) constant returns(uint256[] positionIDs)
func (_Futures *FuturesCallerSession) GetPositionIDsForTrader(trader common.Address, marketID [32]byte) ([]*big.Int, error) {
	return _Futures.Contract.GetPositionIDsForTrader(&_Futures.CallOpts, trader, marketID)
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256)[])
func (_Futures *FuturesCaller) GetPositionsForTrader(opts *bind.CallOpts, trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	var (
		ret0 = new([]TypesPosition)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "getPositionsForTrader", trader, marketID)
	return *ret0, err
}


// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256)[])
func (_Futures *FuturesSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return _Futures.Contract.GetPositionsForTrader(&_Futures.CallOpts, trader, marketID)
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256)[])
func (_Futures *FuturesCallerSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return _Futures.Contract.GetPositionsForTrader(&_Futures.CallOpts, trader, marketID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Futures *FuturesCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Futures *FuturesSession) IsOwner() (bool, error) {
	return _Futures.Contract.IsOwner(&_Futures.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Futures *FuturesCallerSession) IsOwner() (bool, error) {
	return _Futures.Contract.IsOwner(&_Futures.CallOpts)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature) view returns(bool isValid)
func (_Futures *FuturesCaller) IsValidOrderSignature(opts *bind.CallOpts, order Order, signature []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "isValidOrderSignature", order, signature)
	return *ret0, err
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature) view returns(bool isValid)
func (_Futures *FuturesSession) IsValidOrderSignature(order Order, signature []byte) (bool, error) {
	return _Futures.Contract.IsValidOrderSignature(&_Futures.CallOpts, order, signature)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature) view returns(bool isValid)
func (_Futures *FuturesCallerSession) IsValidOrderSignature(order Order, signature []byte) (bool, error) {
	return _Futures.Contract.IsValidOrderSignature(&_Futures.CallOpts, order, signature)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_Futures *FuturesCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "marketCount")
	return *ret0, err
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_Futures *FuturesSession) MarketCount() (*big.Int, error) {
	return _Futures.Contract.MarketCount(&_Futures.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_Futures *FuturesCallerSession) MarketCount() (*big.Int, error) {
	return _Futures.Contract.MarketCount(&_Futures.CallOpts)
}

// MarketSerialToID is a free data retrieval call binding the contract method 0xbae18473.
//
// Solidity: function marketSerialToID(uint256 ) view returns(bytes32)
func (_Futures *FuturesCaller) MarketSerialToID(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "marketSerialToID", arg0)
	return *ret0, err
}

// MarketSerialToID is a free data retrieval call binding the contract method 0xbae18473.
//
// Solidity: function marketSerialToID(uint256 ) view returns(bytes32)
func (_Futures *FuturesSession) MarketSerialToID(arg0 *big.Int) ([32]byte, error) {
	return _Futures.Contract.MarketSerialToID(&_Futures.CallOpts, arg0)
}

// MarketSerialToID is a free data retrieval call binding the contract method 0xbae18473.
//
// Solidity: function marketSerialToID(uint256 ) view returns(bytes32)
func (_Futures *FuturesCallerSession) MarketSerialToID(arg0 *big.Int) ([32]byte, error) {
	return _Futures.Contract.MarketSerialToID(&_Futures.CallOpts, arg0)
}

// MarketToAccountToPositionID is a free data retrieval call binding the contract method 0x54971478.
//
// Solidity: function marketToAccountToPositionID(bytes32 , bytes32 ) view returns(uint256)
func (_Futures *FuturesCaller) MarketToAccountToPositionID(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "marketToAccountToPositionID", arg0, arg1)
	return *ret0, err
}

// MarketToAccountToPositionID is a free data retrieval call binding the contract method 0x54971478.
//
// Solidity: function marketToAccountToPositionID(bytes32 , bytes32 ) view returns(uint256)
func (_Futures *FuturesSession) MarketToAccountToPositionID(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return _Futures.Contract.MarketToAccountToPositionID(&_Futures.CallOpts, arg0, arg1)
}

// MarketToAccountToPositionID is a free data retrieval call binding the contract method 0x54971478.
//
// Solidity: function marketToAccountToPositionID(bytes32 , bytes32 ) view returns(uint256)
func (_Futures *FuturesCallerSession) MarketToAccountToPositionID(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return _Futures.Contract.MarketToAccountToPositionID(&_Futures.CallOpts, arg0, arg1)
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) view returns(bytes32 marketID, string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 indexPrice, uint256 currFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding)
func (_Futures *FuturesCaller) Markets(opts *bind.CallOpts, arg0 [32]byte) (struct {
	MarketID             [32]byte
	Ticker               string
	Oracle               common.Address
	InitialMarginRatio   *big.Int
	LiquidationPenalty   *big.Int
	IndexPrice           *big.Int
	CurrFundingTimestamp *big.Int
	FundingInterval      *big.Int
	CumulativeFunding    *big.Int
}, error) {
	ret := new(struct {
		MarketID             [32]byte
		Ticker               string
		Oracle               common.Address
		InitialMarginRatio   *big.Int
		LiquidationPenalty   *big.Int
		IndexPrice           *big.Int
		CurrFundingTimestamp *big.Int
		FundingInterval      *big.Int
		CumulativeFunding    *big.Int
	})
	out := ret
	err := _Futures.contract.Call(opts, out, "markets", arg0)
	return *ret, err
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) view returns(bytes32 marketID, string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 indexPrice, uint256 currFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding)
func (_Futures *FuturesSession) Markets(arg0 [32]byte) (struct {
	MarketID             [32]byte
	Ticker               string
	Oracle               common.Address
	InitialMarginRatio   *big.Int
	LiquidationPenalty   *big.Int
	IndexPrice           *big.Int
	CurrFundingTimestamp *big.Int
	FundingInterval      *big.Int
	CumulativeFunding    *big.Int
}, error) {
	return _Futures.Contract.Markets(&_Futures.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) view returns(bytes32 marketID, string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 indexPrice, uint256 currFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding)
func (_Futures *FuturesCallerSession) Markets(arg0 [32]byte) (struct {
	MarketID             [32]byte
	Ticker               string
	Oracle               common.Address
	InitialMarginRatio   *big.Int
	LiquidationPenalty   *big.Int
	IndexPrice           *big.Int
	CurrFundingTimestamp *big.Int
	FundingInterval      *big.Int
	CumulativeFunding    *big.Int
}, error) {
	return _Futures.Contract.Markets(&_Futures.CallOpts, arg0)
}

// OrderPosition is a free data retrieval call binding the contract method 0x9a9137ce.
//
// Solidity: function orderPosition(bytes32 ) view returns(uint256)
func (_Futures *FuturesCaller) OrderPosition(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "orderPosition", arg0)
	return *ret0, err
}

// OrderPosition is a free data retrieval call binding the contract method 0x9a9137ce.
//
// Solidity: function orderPosition(bytes32 ) view returns(uint256)
func (_Futures *FuturesSession) OrderPosition(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.OrderPosition(&_Futures.CallOpts, arg0)
}

// OrderPosition is a free data retrieval call binding the contract method 0x9a9137ce.
//
// Solidity: function orderPosition(bytes32 ) view returns(uint256)
func (_Futures *FuturesCallerSession) OrderPosition(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.OrderPosition(&_Futures.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Futures *FuturesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Futures *FuturesSession) Owner() (common.Address, error) {
	return _Futures.Contract.Owner(&_Futures.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Futures *FuturesCallerSession) Owner() (common.Address, error) {
	return _Futures.Contract.Owner(&_Futures.CallOpts)
}

// PooledDeposits is a free data retrieval call binding the contract method 0xef2bcc0f.
//
// Solidity: function pooledDeposits(bytes32 ) view returns(int256)
func (_Futures *FuturesCaller) PooledDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "pooledDeposits", arg0)
	return *ret0, err
}

// PooledDeposits is a free data retrieval call binding the contract method 0xef2bcc0f.
//
// Solidity: function pooledDeposits(bytes32 ) view returns(int256)
func (_Futures *FuturesSession) PooledDeposits(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.PooledDeposits(&_Futures.CallOpts, arg0)
}

// PooledDeposits is a free data retrieval call binding the contract method 0xef2bcc0f.
//
// Solidity: function pooledDeposits(bytes32 ) view returns(int256)
func (_Futures *FuturesCallerSession) PooledDeposits(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.PooledDeposits(&_Futures.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 accountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, int256 NPV, uint256 minMargin, uint256 margin, int256 cumulativeFundingEntry)
func (_Futures *FuturesCaller) Positions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AccountID              [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	ContractPrice          *big.Int
	NPV                    *big.Int
	MinMargin              *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
}, error) {
	ret := new(struct {
		AccountID              [32]byte
		MarketID               [32]byte
		Direction              uint8
		Quantity               *big.Int
		ContractPrice          *big.Int
		NPV                    *big.Int
		MinMargin              *big.Int
		Margin                 *big.Int
		CumulativeFundingEntry *big.Int
	})
	out := ret
	err := _Futures.contract.Call(opts, out, "positions", arg0)
	return *ret, err
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 accountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, int256 NPV, uint256 minMargin, uint256 margin, int256 cumulativeFundingEntry)
func (_Futures *FuturesSession) Positions(arg0 *big.Int) (struct {
	AccountID              [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	ContractPrice          *big.Int
	NPV                    *big.Int
	MinMargin              *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
}, error) {
	return _Futures.Contract.Positions(&_Futures.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 accountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, int256 NPV, uint256 minMargin, uint256 margin, int256 cumulativeFundingEntry)
func (_Futures *FuturesCallerSession) Positions(arg0 *big.Int) (struct {
	AccountID              [32]byte
	MarketID               [32]byte
	Direction              uint8
	Quantity               *big.Int
	ContractPrice          *big.Int
	NPV                    *big.Int
	MinMargin              *big.Int
	Margin                 *big.Int
	CumulativeFundingEntry *big.Int
}, error) {
	return _Futures.Contract.Positions(&_Futures.CallOpts, arg0)
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (_Futures *FuturesCaller) PreSigned(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "preSigned", arg0, arg1)
	return *ret0, err
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (_Futures *FuturesSession) PreSigned(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Futures.Contract.PreSigned(&_Futures.CallOpts, arg0, arg1)
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (_Futures *FuturesCallerSession) PreSigned(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Futures.Contract.PreSigned(&_Futures.CallOpts, arg0, arg1)
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x9a009c3f.
//
// Solidity: function restrictedDeposits(bytes32 ) view returns(uint256)
func (_Futures *FuturesCaller) RestrictedDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "restrictedDeposits", arg0)
	return *ret0, err
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x9a009c3f.
//
// Solidity: function restrictedDeposits(bytes32 ) view returns(uint256)
func (_Futures *FuturesSession) RestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.RestrictedDeposits(&_Futures.CallOpts, arg0)
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x9a009c3f.
//
// Solidity: function restrictedDeposits(bytes32 ) view returns(uint256)
func (_Futures *FuturesCallerSession) RestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.RestrictedDeposits(&_Futures.CallOpts, arg0)
}

// UnrestrictedDeposits is a free data retrieval call binding the contract method 0xa890bb84.
//
// Solidity: function unrestrictedDeposits(bytes32 ) view returns(uint256)
func (_Futures *FuturesCaller) UnrestrictedDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Futures.contract.Call(opts, out, "unrestrictedDeposits", arg0)
	return *ret0, err
}

// UnrestrictedDeposits is a free data retrieval call binding the contract method 0xa890bb84.
//
// Solidity: function unrestrictedDeposits(bytes32 ) view returns(uint256)
func (_Futures *FuturesSession) UnrestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.UnrestrictedDeposits(&_Futures.CallOpts, arg0)
}

// UnrestrictedDeposits is a free data retrieval call binding the contract method 0xa890bb84.
//
// Solidity: function unrestrictedDeposits(bytes32 ) view returns(uint256)
func (_Futures *FuturesCallerSession) UnrestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return _Futures.Contract.UnrestrictedDeposits(&_Futures.CallOpts, arg0)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) returns()
func (_Futures *FuturesTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) returns()
func (_Futures *FuturesSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _Futures.Contract.CancelOrder(&_Futures.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) returns()
func (_Futures *FuturesTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _Futures.Contract.CancelOrder(&_Futures.TransactOpts, order)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xb012d424.
//
// Solidity: function closePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns(bytes32)
func (_Futures *FuturesTransactor) ClosePosition(opts *bind.TransactOpts, positionID *big.Int, order Order, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "closePosition", positionID, order, quantity, signature)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xb012d424.
//
// Solidity: function closePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns(bytes32)
func (_Futures *FuturesSession) ClosePosition(positionID *big.Int, order Order, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _Futures.Contract.ClosePosition(&_Futures.TransactOpts, positionID, order, quantity, signature)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xb012d424.
//
// Solidity: function closePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns(bytes32)
func (_Futures *FuturesTransactorSession) ClosePosition(positionID *big.Int, order Order, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _Futures.Contract.ClosePosition(&_Futures.TransactOpts, positionID, order, quantity, signature)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bytes32 accountID)
func (_Futures *FuturesTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bytes32 accountID)
func (_Futures *FuturesSession) CreateAccount() (*types.Transaction, error) {
	return _Futures.Contract.CreateAccount(&_Futures.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bytes32 accountID)
func (_Futures *FuturesTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _Futures.Contract.CreateAccount(&_Futures.TransactOpts)
}

// CreateAccountAndDeposit is a paid mutator transaction binding the contract method 0xa68a4c3d.
//
// Solidity: function createAccountAndDeposit(uint256 amount) returns(bytes32)
func (_Futures *FuturesTransactor) CreateAccountAndDeposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "createAccountAndDeposit", amount)
}

// CreateAccountAndDeposit is a paid mutator transaction binding the contract method 0xa68a4c3d.
//
// Solidity: function createAccountAndDeposit(uint256 amount) returns(bytes32)
func (_Futures *FuturesSession) CreateAccountAndDeposit(amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.CreateAccountAndDeposit(&_Futures.TransactOpts, amount)
}

// CreateAccountAndDeposit is a paid mutator transaction binding the contract method 0xa68a4c3d.
//
// Solidity: function createAccountAndDeposit(uint256 amount) returns(bytes32)
func (_Futures *FuturesTransactorSession) CreateAccountAndDeposit(amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.CreateAccountAndDeposit(&_Futures.TransactOpts, amount)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x1af69e5f.
//
// Solidity: function createMarket(string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 fundingInterval) returns()
func (_Futures *FuturesTransactor) CreateMarket(opts *bind.TransactOpts, ticker string, oracle common.Address, initialMarginRatio *big.Int, liquidationPenalty *big.Int, fundingInterval *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "createMarket", ticker, oracle, initialMarginRatio, liquidationPenalty, fundingInterval)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x1af69e5f.
//
// Solidity: function createMarket(string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 fundingInterval) returns()
func (_Futures *FuturesSession) CreateMarket(ticker string, oracle common.Address, initialMarginRatio *big.Int, liquidationPenalty *big.Int, fundingInterval *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.CreateMarket(&_Futures.TransactOpts, ticker, oracle, initialMarginRatio, liquidationPenalty, fundingInterval)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x1af69e5f.
//
// Solidity: function createMarket(string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 fundingInterval) returns()
func (_Futures *FuturesTransactorSession) CreateMarket(ticker string, oracle common.Address, initialMarginRatio *big.Int, liquidationPenalty *big.Int, fundingInterval *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.CreateMarket(&_Futures.TransactOpts, ticker, oracle, initialMarginRatio, liquidationPenalty, fundingInterval)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Futures *FuturesTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Futures *FuturesSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Deposit(&_Futures.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Futures *FuturesTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Deposit(&_Futures.TransactOpts, amount)
}

// DepositAccount is a paid mutator transaction binding the contract method 0x04c19cf0.
//
// Solidity: function depositAccount(bytes32 accountID, uint256 amount) returns()
func (_Futures *FuturesTransactor) DepositAccount(opts *bind.TransactOpts, accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "depositAccount", accountID, amount)
}

// DepositAccount is a paid mutator transaction binding the contract method 0x04c19cf0.
//
// Solidity: function depositAccount(bytes32 accountID, uint256 amount) returns()
func (_Futures *FuturesSession) DepositAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.DepositAccount(&_Futures.TransactOpts, accountID, amount)
}

// DepositAccount is a paid mutator transaction binding the contract method 0x04c19cf0.
//
// Solidity: function depositAccount(bytes32 accountID, uint256 amount) returns()
func (_Futures *FuturesTransactorSession) DepositAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.DepositAccount(&_Futures.TransactOpts, accountID, amount)
}

// DepositNewAccount is a paid mutator transaction binding the contract method 0x4a692978.
//
// Solidity: function depositNewAccount(bytes32 accountID, uint256 amount) returns()
func (_Futures *FuturesTransactor) DepositNewAccount(opts *bind.TransactOpts, accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "depositNewAccount", accountID, amount)
}

// DepositNewAccount is a paid mutator transaction binding the contract method 0x4a692978.
//
// Solidity: function depositNewAccount(bytes32 accountID, uint256 amount) returns()
func (_Futures *FuturesSession) DepositNewAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.DepositNewAccount(&_Futures.TransactOpts, accountID, amount)
}

// DepositNewAccount is a paid mutator transaction binding the contract method 0x4a692978.
//
// Solidity: function depositNewAccount(bytes32 accountID, uint256 amount) returns()
func (_Futures *FuturesTransactorSession) DepositNewAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.DepositNewAccount(&_Futures.TransactOpts, accountID, amount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x833b2ea5.
//
// Solidity: function fillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes signature) returns(bytes32)
func (_Futures *FuturesTransactor) FillOrder(opts *bind.TransactOpts, order Order, quantity *big.Int, margin *big.Int, signature []byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "fillOrder", order, quantity, margin, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x833b2ea5.
//
// Solidity: function fillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes signature) returns(bytes32)
func (_Futures *FuturesSession) FillOrder(order Order, quantity *big.Int, margin *big.Int, signature []byte) (*types.Transaction, error) {
	return _Futures.Contract.FillOrder(&_Futures.TransactOpts, order, quantity, margin, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x833b2ea5.
//
// Solidity: function fillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes signature) returns(bytes32)
func (_Futures *FuturesTransactorSession) FillOrder(order Order, quantity *big.Int, margin *big.Int, signature []byte) (*types.Transaction, error) {
	return _Futures.Contract.FillOrder(&_Futures.TransactOpts, order, quantity, margin, signature)
}

// FundPooledDeposits is a paid mutator transaction binding the contract method 0x8f7e9459.
//
// Solidity: function fundPooledDeposits(uint256 amount, bytes32 marketID) returns()
func (_Futures *FuturesTransactor) FundPooledDeposits(opts *bind.TransactOpts, amount *big.Int, marketID [32]byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "fundPooledDeposits", amount, marketID)
}

// FundPooledDeposits is a paid mutator transaction binding the contract method 0x8f7e9459.
//
// Solidity: function fundPooledDeposits(uint256 amount, bytes32 marketID) returns()
func (_Futures *FuturesSession) FundPooledDeposits(amount *big.Int, marketID [32]byte) (*types.Transaction, error) {
	return _Futures.Contract.FundPooledDeposits(&_Futures.TransactOpts, amount, marketID)
}

// FundPooledDeposits is a paid mutator transaction binding the contract method 0x8f7e9459.
//
// Solidity: function fundPooledDeposits(uint256 amount, bytes32 marketID) returns()
func (_Futures *FuturesTransactorSession) FundPooledDeposits(amount *big.Int, marketID [32]byte) (*types.Transaction, error) {
	return _Futures.Contract.FundPooledDeposits(&_Futures.TransactOpts, amount, marketID)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xd440e9b6.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes[] signatures) returns(bytes32)
func (_Futures *FuturesTransactor) MarketOrders(opts *bind.TransactOpts, orders []Order, quantity *big.Int, margin *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "marketOrders", orders, quantity, margin, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xd440e9b6.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes[] signatures) returns(bytes32)
func (_Futures *FuturesSession) MarketOrders(orders []Order, quantity *big.Int, margin *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Futures.Contract.MarketOrders(&_Futures.TransactOpts, orders, quantity, margin, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xd440e9b6.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes[] signatures) returns(bytes32)
func (_Futures *FuturesTransactorSession) MarketOrders(orders []Order, quantity *big.Int, margin *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Futures.Contract.MarketOrders(&_Futures.TransactOpts, orders, quantity, margin, signatures)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns()
func (_Futures *FuturesTransactor) MatchOrders(opts *bind.TransactOpts, leftOrder Order, rightOrder Order, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "matchOrders", leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns()
func (_Futures *FuturesSession) MatchOrders(leftOrder Order, rightOrder Order, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _Futures.Contract.MatchOrders(&_Futures.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns()
func (_Futures *FuturesTransactorSession) MatchOrders(leftOrder Order, rightOrder Order, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _Futures.Contract.MatchOrders(&_Futures.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes[] leftSignatures, bytes rightSignature) returns()
func (_Futures *FuturesTransactor) MultiMatchOrders(opts *bind.TransactOpts, leftOrders []Order, rightOrder Order, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "multiMatchOrders", leftOrders, rightOrder, leftSignatures, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes[] leftSignatures, bytes rightSignature) returns()
func (_Futures *FuturesSession) MultiMatchOrders(leftOrders []Order, rightOrder Order, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return _Futures.Contract.MultiMatchOrders(&_Futures.TransactOpts, leftOrders, rightOrder, leftSignatures, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes[] leftSignatures, bytes rightSignature) returns()
func (_Futures *FuturesTransactorSession) MultiMatchOrders(leftOrders []Order, rightOrder Order, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return _Futures.Contract.MultiMatchOrders(&_Futures.TransactOpts, leftOrders, rightOrder, leftSignatures, rightSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Futures *FuturesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Futures *FuturesSession) RenounceOwnership() (*types.Transaction, error) {
	return _Futures.Contract.RenounceOwnership(&_Futures.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Futures *FuturesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Futures.Contract.RenounceOwnership(&_Futures.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Futures *FuturesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Futures *FuturesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Futures.Contract.TransferOwnership(&_Futures.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Futures *FuturesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Futures.Contract.TransferOwnership(&_Futures.TransactOpts, newOwner)
}

// VerifyClose is a paid mutator transaction binding the contract method 0xe8e73d08.
//
// Solidity: function verifyClose(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256) newPosition, bytes32 hash)
func (_Futures *FuturesTransactor) VerifyClose(opts *bind.TransactOpts, positionID *big.Int, order Order, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "verifyClose", positionID, order, quantity, signature)
}

// VerifyClose is a paid mutator transaction binding the contract method 0xe8e73d08.
//
// Solidity: function verifyClose(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256) newPosition, bytes32 hash)
func (_Futures *FuturesSession) VerifyClose(positionID *big.Int, order Order, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _Futures.Contract.VerifyClose(&_Futures.TransactOpts, positionID, order, quantity, signature)
}

// VerifyClose is a paid mutator transaction binding the contract method 0xe8e73d08.
//
// Solidity: function verifyClose(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256) newPosition, bytes32 hash)
func (_Futures *FuturesTransactorSession) VerifyClose(positionID *big.Int, order Order, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _Futures.Contract.VerifyClose(&_Futures.TransactOpts, positionID, order, quantity, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Futures *FuturesTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Futures *FuturesSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Withdraw(&_Futures.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Futures *FuturesTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Withdraw(&_Futures.TransactOpts, amount)
}

// WithdrawAccount is a paid mutator transaction binding the contract method 0x4d6a5cbc.
//
// Solidity: function withdrawAccount(bytes32 accountID, uint256 amount) returns()
func (_Futures *FuturesTransactor) WithdrawAccount(opts *bind.TransactOpts, accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "withdrawAccount", accountID, amount)
}

// WithdrawAccount is a paid mutator transaction binding the contract method 0x4d6a5cbc.
//
// Solidity: function withdrawAccount(bytes32 accountID, uint256 amount) returns()
func (_Futures *FuturesSession) WithdrawAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.WithdrawAccount(&_Futures.TransactOpts, accountID, amount)
}

// WithdrawAccount is a paid mutator transaction binding the contract method 0x4d6a5cbc.
//
// Solidity: function withdrawAccount(bytes32 accountID, uint256 amount) returns()
func (_Futures *FuturesTransactorSession) WithdrawAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.WithdrawAccount(&_Futures.TransactOpts, accountID, amount)
}

// FuturesFillIterator is returned from FilterFuturesFill and is used to iterate over the raw logs and unpacked data for FuturesFill events raised by the Futures contract.
type FuturesFillIterator struct {
	Event *FuturesFill // Event containing the contract specifics and raw log

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
func (it *FuturesFillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesFill)
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
		it.Event = new(FuturesFill)
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
func (it *FuturesFillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesFill represents a FuturesFill event raised by the Futures contract.
type FuturesFill struct {
	MakerAddress   common.Address
	OrderHash      [32]byte
	MarketID       [32]byte
	ContractPrice  *big.Int
	QuantityFilled *big.Int
	Quantity       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFuturesFill is a free log retrieval operation binding the contract event 0x6c3c91984205906d99b5675f1c6cec9054481daa142b00699944b6072a44b7e6.
//
// Solidity: event FuturesFill(address indexed makerAddress, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled, uint256 quantity)
func (_Futures *FuturesFilterer) FilterFuturesFill(opts *bind.FilterOpts, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (*FuturesFillIterator, error) {

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

	logs, sub, err := _Futures.contract.FilterLogs(opts, "FuturesFill", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesFillIterator{contract: _Futures.contract, event: "FuturesFill", logs: logs, sub: sub}, nil
}

// WatchFuturesFill is a free log subscription operation binding the contract event 0x6c3c91984205906d99b5675f1c6cec9054481daa142b00699944b6072a44b7e6.
//
// Solidity: event FuturesFill(address indexed makerAddress, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled, uint256 quantity)
func (_Futures *FuturesFilterer) WatchFuturesFill(opts *bind.WatchOpts, sink chan<- *FuturesFill, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Futures.contract.WatchLogs(opts, "FuturesFill", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesFill)
				if err := _Futures.contract.UnpackLog(event, "FuturesFill", log); err != nil {
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

// ParseFuturesFill is a log parse operation binding the contract event 0x6c3c91984205906d99b5675f1c6cec9054481daa142b00699944b6072a44b7e6.
//
// Solidity: event FuturesFill(address indexed makerAddress, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled, uint256 quantity)
func (_Futures *FuturesFilterer) ParseFuturesFill(log types.Log) (*FuturesFill, error) {
	event := new(FuturesFill)
	if err := _Futures.contract.UnpackLog(event, "FuturesFill", log); err != nil {
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
	MarketID [32]byte
	Ticker   common.Hash
	Oracle   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketCreation is a free log retrieval operation binding the contract event 0xa89968288be926b3b7a2a5a0d70c9e55178a465c62e84ad1c3f0401d8c2f7fc1.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle)
func (_Futures *FuturesFilterer) FilterMarketCreation(opts *bind.FilterOpts, marketID [][32]byte, ticker []string, oracle []common.Address) (*FuturesMarketCreationIterator, error) {

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

	logs, sub, err := _Futures.contract.FilterLogs(opts, "MarketCreation", marketIDRule, tickerRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &FuturesMarketCreationIterator{contract: _Futures.contract, event: "MarketCreation", logs: logs, sub: sub}, nil
}

// WatchMarketCreation is a free log subscription operation binding the contract event 0xa89968288be926b3b7a2a5a0d70c9e55178a465c62e84ad1c3f0401d8c2f7fc1.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle)
func (_Futures *FuturesFilterer) WatchMarketCreation(opts *bind.WatchOpts, sink chan<- *FuturesMarketCreation, marketID [][32]byte, ticker []string, oracle []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Futures.contract.WatchLogs(opts, "MarketCreation", marketIDRule, tickerRule, oracleRule)
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
				if err := _Futures.contract.UnpackLog(event, "MarketCreation", log); err != nil {
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

// ParseMarketCreation is a log parse operation binding the contract event 0xa89968288be926b3b7a2a5a0d70c9e55178a465c62e84ad1c3f0401d8c2f7fc1.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle)
func (_Futures *FuturesFilterer) ParseMarketCreation(log types.Log) (*FuturesMarketCreation, error) {
	event := new(FuturesMarketCreation)
	if err := _Futures.contract.UnpackLog(event, "MarketCreation", log); err != nil {
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
func (_Futures *FuturesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FuturesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FuturesOwnershipTransferredIterator{contract: _Futures.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Futures *FuturesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FuturesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
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
				if err := _Futures.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Futures *FuturesFilterer) ParseOwnershipTransferred(log types.Log) (*FuturesOwnershipTransferred, error) {
	event := new(FuturesOwnershipTransferred)
	if err := _Futures.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Futures *FuturesFilterer) FilterSignatureValidatorApproval(opts *bind.FilterOpts, signerAddress []common.Address, validatorAddress []common.Address) (*FuturesSignatureValidatorApprovalIterator, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "SignatureValidatorApproval", signerAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &FuturesSignatureValidatorApprovalIterator{contract: _Futures.contract, event: "SignatureValidatorApproval", logs: logs, sub: sub}, nil
}

// WatchSignatureValidatorApproval is a free log subscription operation binding the contract event 0xa8656e308026eeabce8f0bc18048433252318ab80ac79da0b3d3d8697dfba891.
//
// Solidity: event SignatureValidatorApproval(address indexed signerAddress, address indexed validatorAddress, bool isApproved)
func (_Futures *FuturesFilterer) WatchSignatureValidatorApproval(opts *bind.WatchOpts, sink chan<- *FuturesSignatureValidatorApproval, signerAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "SignatureValidatorApproval", signerAddressRule, validatorAddressRule)
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
				if err := _Futures.contract.UnpackLog(event, "SignatureValidatorApproval", log); err != nil {
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
func (_Futures *FuturesFilterer) ParseSignatureValidatorApproval(log types.Log) (*FuturesSignatureValidatorApproval, error) {
	event := new(FuturesSignatureValidatorApproval)
	if err := _Futures.contract.UnpackLog(event, "SignatureValidatorApproval", log); err != nil {
		return nil, err
	}
	return event, nil
}
