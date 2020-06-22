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

// LibOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type LibOrderOrder struct {
	MakerAddress          common.Address
	TakerAddress          common.Address
	FeeRecipientAddress   common.Address
	SenderAddress         common.Address
	MakerAssetAmount      *big.Int
	TakerAssetAmount      *big.Int
	MakerFee              *big.Int
	TakerFee              *big.Int
	ExpirationTimeSeconds *big.Int
	Salt                  *big.Int
	MakerAssetData        []byte
	TakerAssetData        []byte
	MakerFeeAssetData     []byte
	TakerFeeAssetData     []byte
}

// LibOrderOrderInfo is an auto generated low-level Go binding around an user-defined struct.
type LibOrderOrderInfo struct {
	OrderStatus                 uint8
	OrderHash                   [32]byte
	OrderTakerAssetFilledAmount *big.Int
}

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

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (f *FuturesCaller) EIP1271MAGICVALUE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := f.contract.Call(opts, out, "EIP1271_MAGIC_VALUE")
	return *ret0, err
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

// AccountIdToAddress is a free data retrieval call binding the contract method 0x899706c0.
//
// Solidity: function accountIdToAddress(bytes32 ) view returns(address)
func (f *FuturesCaller) AccountIdToAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := f.contract.Call(opts, out, "accountIdToAddress", arg0)
	return *ret0, err
}

// AccountIdToAddress is a free data retrieval call binding the contract method 0x899706c0.
//
// Solidity: function accountIdToAddress(bytes32 ) view returns(address)
func (f *FuturesSession) AccountIdToAddress(arg0 [32]byte) (common.Address, error) {
	return f.Contract.AccountIdToAddress(&f.CallOpts, arg0)
}

// AccountIdToAddress is a free data retrieval call binding the contract method 0x899706c0.
//
// Solidity: function accountIdToAddress(bytes32 ) view returns(address)
func (f *FuturesCallerSession) AccountIdToAddress(arg0 [32]byte) (common.Address, error) {
	return f.Contract.AccountIdToAddress(&f.CallOpts, arg0)
}

// AccountNonce is a free data retrieval call binding the contract method 0x106cffe2.
//
// Solidity: function accountNonce(address ) view returns(uint256)
func (f *FuturesCaller) AccountNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "accountNonce", arg0)
	return *ret0, err
}

// AccountNonce is a free data retrieval call binding the contract method 0x106cffe2.
//
// Solidity: function accountNonce(address ) view returns(uint256)
func (f *FuturesSession) AccountNonce(arg0 common.Address) (*big.Int, error) {
	return f.Contract.AccountNonce(&f.CallOpts, arg0)
}

// AccountNonce is a free data retrieval call binding the contract method 0x106cffe2.
//
// Solidity: function accountNonce(address ) view returns(uint256)
func (f *FuturesCallerSession) AccountNonce(arg0 common.Address) (*big.Int, error) {
	return f.Contract.AccountNonce(&f.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(int256 NAV, int256 NPV, bytes32 accountID)
func (f *FuturesCaller) Accounts(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := f.contract.Call(opts, out, "accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(int256 NAV, int256 NPV, bytes32 accountID)
func (f *FuturesSession) Accounts(arg0 [32]byte) (struct {
	NAV       *big.Int
	NPV       *big.Int
	AccountID [32]byte
}, error) {
	return f.Contract.Accounts(&f.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(int256 NAV, int256 NPV, bytes32 accountID)
func (f *FuturesCallerSession) Accounts(arg0 [32]byte) (struct {
	NAV       *big.Int
	NPV       *big.Int
	AccountID [32]byte
}, error) {
	return f.Contract.Accounts(&f.CallOpts, arg0)
}

// AddressToAccountIDs is a free data retrieval call binding the contract method 0xfbd6d494.
//
// Solidity: function addressToAccountIDs(address , uint256 ) view returns(bytes32)
func (f *FuturesCaller) AddressToAccountIDs(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := f.contract.Call(opts, out, "addressToAccountIDs", arg0, arg1)
	return *ret0, err
}

// AddressToAccountIDs is a free data retrieval call binding the contract method 0xfbd6d494.
//
// Solidity: function addressToAccountIDs(address , uint256 ) view returns(bytes32)
func (f *FuturesSession) AddressToAccountIDs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return f.Contract.AddressToAccountIDs(&f.CallOpts, arg0, arg1)
}

// AddressToAccountIDs is a free data retrieval call binding the contract method 0xfbd6d494.
//
// Solidity: function addressToAccountIDs(address , uint256 ) view returns(bytes32)
func (f *FuturesCallerSession) AddressToAccountIDs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return f.Contract.AddressToAccountIDs(&f.CallOpts, arg0, arg1)
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
// Solidity: function calcLiquidationFee(bytes32 _marketID, uint256 _quantity) view returns(uint256)
func (f *FuturesCaller) CalcLiquidationFee(opts *bind.CallOpts, _marketID [32]byte, _quantity *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "calcLiquidationFee", _marketID, _quantity)
	return *ret0, err
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 _marketID, uint256 _quantity) view returns(uint256)
func (f *FuturesSession) CalcLiquidationFee(_marketID [32]byte, _quantity *big.Int) (*big.Int, error) {
	return f.Contract.CalcLiquidationFee(&f.CallOpts, _marketID, _quantity)
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 _marketID, uint256 _quantity) view returns(uint256)
func (f *FuturesCallerSession) CalcLiquidationFee(_marketID [32]byte, _quantity *big.Int) (*big.Int, error) {
	return f.Contract.CalcLiquidationFee(&f.CallOpts, _marketID, _quantity)
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

// FreeDeposits is a free data retrieval call binding the contract method 0x9a78538f.
//
// Solidity: function freeDeposits(address ) view returns(uint256)
func (f *FuturesCaller) FreeDeposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "freeDeposits", arg0)
	return *ret0, err
}

// FreeDeposits is a free data retrieval call binding the contract method 0x9a78538f.
//
// Solidity: function freeDeposits(address ) view returns(uint256)
func (f *FuturesSession) FreeDeposits(arg0 common.Address) (*big.Int, error) {
	return f.Contract.FreeDeposits(&f.CallOpts, arg0)
}

// FreeDeposits is a free data retrieval call binding the contract method 0x9a78538f.
//
// Solidity: function freeDeposits(address ) view returns(uint256)
func (f *FuturesCallerSession) FreeDeposits(arg0 common.Address) (*big.Int, error) {
	return f.Contract.FreeDeposits(&f.CallOpts, arg0)
}

// GetBatchBalancesAndAssetProxyAllowances is a free data retrieval call binding the contract method 0x5bf34fc7.
//
// Solidity: function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, address[] tokenAddresses) view returns(uint256[], uint256[])
func (f *FuturesCaller) GetBatchBalancesAndAssetProxyAllowances(opts *bind.CallOpts, ownerAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := f.contract.Call(opts, out, "getBatchBalancesAndAssetProxyAllowances", ownerAddress, tokenAddresses)
	return *ret0, *ret1, err
}

// GetBatchBalancesAndAssetProxyAllowances is a free data retrieval call binding the contract method 0x5bf34fc7.
//
// Solidity: function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, address[] tokenAddresses) view returns(uint256[], uint256[])
func (f *FuturesSession) GetBatchBalancesAndAssetProxyAllowances(ownerAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, []*big.Int, error) {
	return f.Contract.GetBatchBalancesAndAssetProxyAllowances(&f.CallOpts, ownerAddress, tokenAddresses)
}

// GetBatchBalancesAndAssetProxyAllowances is a free data retrieval call binding the contract method 0x5bf34fc7.
//
// Solidity: function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, address[] tokenAddresses) view returns(uint256[], uint256[])
func (f *FuturesCallerSession) GetBatchBalancesAndAssetProxyAllowances(ownerAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, []*big.Int, error) {
	return f.Contract.GetBatchBalancesAndAssetProxyAllowances(&f.CallOpts, ownerAddress, tokenAddresses)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, bytes[] signatures) view returns((uint8,bytes32,uint256)[] ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (f *FuturesCaller) GetOrderRelevantStates(opts *bind.CallOpts, orders []LibOrderOrder, signatures [][]byte) (struct {
	OrdersInfo                []LibOrderOrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	ret := new(struct {
		OrdersInfo                []LibOrderOrderInfo
		FillableTakerAssetAmounts []*big.Int
		IsValidSignature          []bool
	})
	out := ret
	err := f.contract.Call(opts, out, "getOrderRelevantStates", orders, signatures)
	return *ret, err
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, bytes[] signatures) view returns((uint8,bytes32,uint256)[] ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (f *FuturesSession) GetOrderRelevantStates(orders []LibOrderOrder, signatures [][]byte) (struct {
	OrdersInfo                []LibOrderOrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	return f.Contract.GetOrderRelevantStates(&f.CallOpts, orders, signatures)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, bytes[] signatures) view returns((uint8,bytes32,uint256)[] ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (f *FuturesCallerSession) GetOrderRelevantStates(orders []LibOrderOrder, signatures [][]byte) (struct {
	OrdersInfo                []LibOrderOrderInfo
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
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256)[])
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
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256)[])
func (f *FuturesSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return f.Contract.GetPositionsForTrader(&f.CallOpts, trader, marketID)
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) view returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256)[])
func (f *FuturesCallerSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return f.Contract.GetPositionsForTrader(&f.CallOpts, trader, marketID)
}

// GetTransferableAssetAmount is a free data retrieval call binding the contract method 0x478482d8.
//
// Solidity: function getTransferableAssetAmount(address ownerAddress) view returns(uint256 transferableAssetAmount)
func (f *FuturesCaller) GetTransferableAssetAmount(opts *bind.CallOpts, ownerAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "getTransferableAssetAmount", ownerAddress)
	return *ret0, err
}

// GetTransferableAssetAmount is a free data retrieval call binding the contract method 0x478482d8.
//
// Solidity: function getTransferableAssetAmount(address ownerAddress) view returns(uint256 transferableAssetAmount)
func (f *FuturesSession) GetTransferableAssetAmount(ownerAddress common.Address) (*big.Int, error) {
	return f.Contract.GetTransferableAssetAmount(&f.CallOpts, ownerAddress)
}

// GetTransferableAssetAmount is a free data retrieval call binding the contract method 0x478482d8.
//
// Solidity: function getTransferableAssetAmount(address ownerAddress) view returns(uint256 transferableAssetAmount)
func (f *FuturesCallerSession) GetTransferableAssetAmount(ownerAddress common.Address) (*big.Int, error) {
	return f.Contract.GetTransferableAssetAmount(&f.CallOpts, ownerAddress)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (f *FuturesCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := f.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (f *FuturesSession) IsOwner() (bool, error) {
	return f.Contract.IsOwner(&f.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (f *FuturesCallerSession) IsOwner() (bool, error) {
	return f.Contract.IsOwner(&f.CallOpts)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature) view returns(bool isValid)
func (f *FuturesCaller) IsValidOrderSignature(opts *bind.CallOpts, order LibOrderOrder, signature []byte) (bool, error) {
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
func (f *FuturesSession) IsValidOrderSignature(order LibOrderOrder, signature []byte) (bool, error) {
	return f.Contract.IsValidOrderSignature(&f.CallOpts, order, signature)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, bytes signature) view returns(bool isValid)
func (f *FuturesCallerSession) IsValidOrderSignature(order LibOrderOrder, signature []byte) (bool, error) {
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

// MarketToAccountToPositionID is a free data retrieval call binding the contract method 0x54971478.
//
// Solidity: function marketToAccountToPositionID(bytes32 , bytes32 ) view returns(uint256)
func (f *FuturesCaller) MarketToAccountToPositionID(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "marketToAccountToPositionID", arg0, arg1)
	return *ret0, err
}

// MarketToAccountToPositionID is a free data retrieval call binding the contract method 0x54971478.
//
// Solidity: function marketToAccountToPositionID(bytes32 , bytes32 ) view returns(uint256)
func (f *FuturesSession) MarketToAccountToPositionID(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return f.Contract.MarketToAccountToPositionID(&f.CallOpts, arg0, arg1)
}

// MarketToAccountToPositionID is a free data retrieval call binding the contract method 0x54971478.
//
// Solidity: function marketToAccountToPositionID(bytes32 , bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) MarketToAccountToPositionID(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return f.Contract.MarketToAccountToPositionID(&f.CallOpts, arg0, arg1)
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) view returns(bytes32 marketID, string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 indexPrice, uint256 currFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding)
func (f *FuturesCaller) Markets(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := f.contract.Call(opts, out, "markets", arg0)
	return *ret, err
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) view returns(bytes32 marketID, string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 indexPrice, uint256 currFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding)
func (f *FuturesSession) Markets(arg0 [32]byte) (struct {
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
	return f.Contract.Markets(&f.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) view returns(bytes32 marketID, string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 indexPrice, uint256 currFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding)
func (f *FuturesCallerSession) Markets(arg0 [32]byte) (struct {
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
	return f.Contract.Markets(&f.CallOpts, arg0)
}

// OrderPosition is a free data retrieval call binding the contract method 0x9a9137ce.
//
// Solidity: function orderPosition(bytes32 ) view returns(uint256)
func (f *FuturesCaller) OrderPosition(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "orderPosition", arg0)
	return *ret0, err
}

// OrderPosition is a free data retrieval call binding the contract method 0x9a9137ce.
//
// Solidity: function orderPosition(bytes32 ) view returns(uint256)
func (f *FuturesSession) OrderPosition(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.OrderPosition(&f.CallOpts, arg0)
}

// OrderPosition is a free data retrieval call binding the contract method 0x9a9137ce.
//
// Solidity: function orderPosition(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) OrderPosition(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.OrderPosition(&f.CallOpts, arg0)
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

// PooledDeposits is a free data retrieval call binding the contract method 0xef2bcc0f.
//
// Solidity: function pooledDeposits(bytes32 ) view returns(int256)
func (f *FuturesCaller) PooledDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "pooledDeposits", arg0)
	return *ret0, err
}

// PooledDeposits is a free data retrieval call binding the contract method 0xef2bcc0f.
//
// Solidity: function pooledDeposits(bytes32 ) view returns(int256)
func (f *FuturesSession) PooledDeposits(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.PooledDeposits(&f.CallOpts, arg0)
}

// PooledDeposits is a free data retrieval call binding the contract method 0xef2bcc0f.
//
// Solidity: function pooledDeposits(bytes32 ) view returns(int256)
func (f *FuturesCallerSession) PooledDeposits(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.PooledDeposits(&f.CallOpts, arg0)
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
// Solidity: function positions(uint256 ) view returns(bytes32 accountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, int256 NPV, uint256 minMargin, uint256 margin, int256 cumulativeFundingEntry)
func (f *FuturesCaller) Positions(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := f.contract.Call(opts, out, "positions", arg0)
	return *ret, err
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 accountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, int256 NPV, uint256 minMargin, uint256 margin, int256 cumulativeFundingEntry)
func (f *FuturesSession) Positions(arg0 *big.Int) (struct {
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
	return f.Contract.Positions(&f.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(bytes32 accountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, int256 NPV, uint256 minMargin, uint256 margin, int256 cumulativeFundingEntry)
func (f *FuturesCallerSession) Positions(arg0 *big.Int) (struct {
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

// RestrictedDeposits is a free data retrieval call binding the contract method 0x9a009c3f.
//
// Solidity: function restrictedDeposits(bytes32 ) view returns(uint256)
func (f *FuturesCaller) RestrictedDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "restrictedDeposits", arg0)
	return *ret0, err
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x9a009c3f.
//
// Solidity: function restrictedDeposits(bytes32 ) view returns(uint256)
func (f *FuturesSession) RestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.RestrictedDeposits(&f.CallOpts, arg0)
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x9a009c3f.
//
// Solidity: function restrictedDeposits(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) RestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.RestrictedDeposits(&f.CallOpts, arg0)
}

// UnrestrictedDeposits is a free data retrieval call binding the contract method 0xa890bb84.
//
// Solidity: function unrestrictedDeposits(bytes32 ) view returns(uint256)
func (f *FuturesCaller) UnrestrictedDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := f.contract.Call(opts, out, "unrestrictedDeposits", arg0)
	return *ret0, err
}

// UnrestrictedDeposits is a free data retrieval call binding the contract method 0xa890bb84.
//
// Solidity: function unrestrictedDeposits(bytes32 ) view returns(uint256)
func (f *FuturesSession) UnrestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.UnrestrictedDeposits(&f.CallOpts, arg0)
}

// UnrestrictedDeposits is a free data retrieval call binding the contract method 0xa890bb84.
//
// Solidity: function unrestrictedDeposits(bytes32 ) view returns(uint256)
func (f *FuturesCallerSession) UnrestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return f.Contract.UnrestrictedDeposits(&f.CallOpts, arg0)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) returns()
func (f *FuturesTransactor) CancelOrder(opts *bind.TransactOpts, order LibOrderOrder) (*types.Transaction, error) {
	return f.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) returns()
func (f *FuturesSession) CancelOrder(order LibOrderOrder) (*types.Transaction, error) {
	return f.Contract.CancelOrder(&f.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order) returns()
func (f *FuturesTransactorSession) CancelOrder(order LibOrderOrder) (*types.Transaction, error) {
	return f.Contract.CancelOrder(&f.TransactOpts, order)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xb012d424.
//
// Solidity: function closePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns(bytes32)
func (f *FuturesTransactor) ClosePosition(opts *bind.TransactOpts, positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "closePosition", positionID, order, quantity, signature)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xb012d424.
//
// Solidity: function closePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns(bytes32)
func (f *FuturesSession) ClosePosition(positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return f.Contract.ClosePosition(&f.TransactOpts, positionID, order, quantity, signature)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xb012d424.
//
// Solidity: function closePosition(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns(bytes32)
func (f *FuturesTransactorSession) ClosePosition(positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return f.Contract.ClosePosition(&f.TransactOpts, positionID, order, quantity, signature)
}

// ClosePositionWithOrders is a paid mutator transaction binding the contract method 0xeeb3da17.
//
// Solidity: function closePositionWithOrders(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes signature) returns()
func (f *FuturesTransactor) ClosePositionWithOrders(opts *bind.TransactOpts, positionID *big.Int, orders []LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "closePositionWithOrders", positionID, orders, quantity, signature)
}

// ClosePositionWithOrders is a paid mutator transaction binding the contract method 0xeeb3da17.
//
// Solidity: function closePositionWithOrders(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes signature) returns()
func (f *FuturesSession) ClosePositionWithOrders(positionID *big.Int, orders []LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return f.Contract.ClosePositionWithOrders(&f.TransactOpts, positionID, orders, quantity, signature)
}

// ClosePositionWithOrders is a paid mutator transaction binding the contract method 0xeeb3da17.
//
// Solidity: function closePositionWithOrders(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, bytes signature) returns()
func (f *FuturesTransactorSession) ClosePositionWithOrders(positionID *big.Int, orders []LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return f.Contract.ClosePositionWithOrders(&f.TransactOpts, positionID, orders, quantity, signature)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bytes32 accountID)
func (f *FuturesTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bytes32 accountID)
func (f *FuturesSession) CreateAccount() (*types.Transaction, error) {
	return f.Contract.CreateAccount(&f.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bytes32 accountID)
func (f *FuturesTransactorSession) CreateAccount() (*types.Transaction, error) {
	return f.Contract.CreateAccount(&f.TransactOpts)
}

// CreateAccountAndDeposit is a paid mutator transaction binding the contract method 0xa68a4c3d.
//
// Solidity: function createAccountAndDeposit(uint256 amount) returns(bytes32)
func (f *FuturesTransactor) CreateAccountAndDeposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createAccountAndDeposit", amount)
}

// CreateAccountAndDeposit is a paid mutator transaction binding the contract method 0xa68a4c3d.
//
// Solidity: function createAccountAndDeposit(uint256 amount) returns(bytes32)
func (f *FuturesSession) CreateAccountAndDeposit(amount *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateAccountAndDeposit(&f.TransactOpts, amount)
}

// CreateAccountAndDeposit is a paid mutator transaction binding the contract method 0xa68a4c3d.
//
// Solidity: function createAccountAndDeposit(uint256 amount) returns(bytes32)
func (f *FuturesTransactorSession) CreateAccountAndDeposit(amount *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateAccountAndDeposit(&f.TransactOpts, amount)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x1af69e5f.
//
// Solidity: function createMarket(string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 fundingInterval) returns()
func (f *FuturesTransactor) CreateMarket(opts *bind.TransactOpts, ticker string, oracle common.Address, initialMarginRatio *big.Int, liquidationPenalty *big.Int, fundingInterval *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createMarket", ticker, oracle, initialMarginRatio, liquidationPenalty, fundingInterval)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x1af69e5f.
//
// Solidity: function createMarket(string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 fundingInterval) returns()
func (f *FuturesSession) CreateMarket(ticker string, oracle common.Address, initialMarginRatio *big.Int, liquidationPenalty *big.Int, fundingInterval *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateMarket(&f.TransactOpts, ticker, oracle, initialMarginRatio, liquidationPenalty, fundingInterval)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x1af69e5f.
//
// Solidity: function createMarket(string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 fundingInterval) returns()
func (f *FuturesTransactorSession) CreateMarket(ticker string, oracle common.Address, initialMarginRatio *big.Int, liquidationPenalty *big.Int, fundingInterval *big.Int) (*types.Transaction, error) {
	return f.Contract.CreateMarket(&f.TransactOpts, ticker, oracle, initialMarginRatio, liquidationPenalty, fundingInterval)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (f *FuturesTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (f *FuturesSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Deposit(&f.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (f *FuturesTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Deposit(&f.TransactOpts, amount)
}

// DepositAccount is a paid mutator transaction binding the contract method 0x04c19cf0.
//
// Solidity: function depositAccount(bytes32 accountID, uint256 amount) returns()
func (f *FuturesTransactor) DepositAccount(opts *bind.TransactOpts, accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "depositAccount", accountID, amount)
}

// DepositAccount is a paid mutator transaction binding the contract method 0x04c19cf0.
//
// Solidity: function depositAccount(bytes32 accountID, uint256 amount) returns()
func (f *FuturesSession) DepositAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositAccount(&f.TransactOpts, accountID, amount)
}

// DepositAccount is a paid mutator transaction binding the contract method 0x04c19cf0.
//
// Solidity: function depositAccount(bytes32 accountID, uint256 amount) returns()
func (f *FuturesTransactorSession) DepositAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositAccount(&f.TransactOpts, accountID, amount)
}

// DepositNewAccount is a paid mutator transaction binding the contract method 0x4a692978.
//
// Solidity: function depositNewAccount(bytes32 accountID, uint256 amount) returns()
func (f *FuturesTransactor) DepositNewAccount(opts *bind.TransactOpts, accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "depositNewAccount", accountID, amount)
}

// DepositNewAccount is a paid mutator transaction binding the contract method 0x4a692978.
//
// Solidity: function depositNewAccount(bytes32 accountID, uint256 amount) returns()
func (f *FuturesSession) DepositNewAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositNewAccount(&f.TransactOpts, accountID, amount)
}

// DepositNewAccount is a paid mutator transaction binding the contract method 0x4a692978.
//
// Solidity: function depositNewAccount(bytes32 accountID, uint256 amount) returns()
func (f *FuturesTransactorSession) DepositNewAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.DepositNewAccount(&f.TransactOpts, accountID, amount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x833b2ea5.
//
// Solidity: function fillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes signature) returns(bytes32)
func (f *FuturesTransactor) FillOrder(opts *bind.TransactOpts, order LibOrderOrder, quantity *big.Int, margin *big.Int, signature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "fillOrder", order, quantity, margin, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x833b2ea5.
//
// Solidity: function fillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes signature) returns(bytes32)
func (f *FuturesSession) FillOrder(order LibOrderOrder, quantity *big.Int, margin *big.Int, signature []byte) (*types.Transaction, error) {
	return f.Contract.FillOrder(&f.TransactOpts, order, quantity, margin, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x833b2ea5.
//
// Solidity: function fillOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, uint256 margin, bytes signature) returns(bytes32)
func (f *FuturesTransactorSession) FillOrder(order LibOrderOrder, quantity *big.Int, margin *big.Int, signature []byte) (*types.Transaction, error) {
	return f.Contract.FillOrder(&f.TransactOpts, order, quantity, margin, signature)
}

// FundPooledDeposits is a paid mutator transaction binding the contract method 0x8f7e9459.
//
// Solidity: function fundPooledDeposits(uint256 amount, bytes32 marketID) returns()
func (f *FuturesTransactor) FundPooledDeposits(opts *bind.TransactOpts, amount *big.Int, marketID [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "fundPooledDeposits", amount, marketID)
}

// FundPooledDeposits is a paid mutator transaction binding the contract method 0x8f7e9459.
//
// Solidity: function fundPooledDeposits(uint256 amount, bytes32 marketID) returns()
func (f *FuturesSession) FundPooledDeposits(amount *big.Int, marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.FundPooledDeposits(&f.TransactOpts, amount, marketID)
}

// FundPooledDeposits is a paid mutator transaction binding the contract method 0x8f7e9459.
//
// Solidity: function fundPooledDeposits(uint256 amount, bytes32 marketID) returns()
func (f *FuturesTransactorSession) FundPooledDeposits(amount *big.Int, marketID [32]byte) (*types.Transaction, error) {
	return f.Contract.FundPooledDeposits(&f.TransactOpts, amount, marketID)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xd440e9b6.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes[] signatures) returns(bytes32)
func (f *FuturesTransactor) MarketOrders(opts *bind.TransactOpts, orders []LibOrderOrder, quantity *big.Int, margin *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "marketOrders", orders, quantity, margin, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xd440e9b6.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes[] signatures) returns(bytes32)
func (f *FuturesSession) MarketOrders(orders []LibOrderOrder, quantity *big.Int, margin *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.MarketOrders(&f.TransactOpts, orders, quantity, margin, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xd440e9b6.
//
// Solidity: function marketOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] orders, uint256 quantity, uint256 margin, bytes[] signatures) returns(bytes32)
func (f *FuturesTransactorSession) MarketOrders(orders []LibOrderOrder, quantity *big.Int, margin *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return f.Contract.MarketOrders(&f.TransactOpts, orders, quantity, margin, signatures)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns()
func (f *FuturesTransactor) MatchOrders(opts *bind.TransactOpts, leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "matchOrders", leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns()
func (f *FuturesSession) MatchOrders(leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MatchOrders(&f.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) leftOrder, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes leftSignature, bytes rightSignature) returns()
func (f *FuturesTransactorSession) MatchOrders(leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MatchOrders(&f.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes[] leftSignatures, bytes rightSignature) returns()
func (f *FuturesTransactor) MultiMatchOrders(opts *bind.TransactOpts, leftOrders []LibOrderOrder, rightOrder LibOrderOrder, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "multiMatchOrders", leftOrders, rightOrder, leftSignatures, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes[] leftSignatures, bytes rightSignature) returns()
func (f *FuturesSession) MultiMatchOrders(leftOrders []LibOrderOrder, rightOrder LibOrderOrder, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MultiMatchOrders(&f.TransactOpts, leftOrders, rightOrder, leftSignatures, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] leftOrders, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) rightOrder, bytes[] leftSignatures, bytes rightSignature) returns()
func (f *FuturesTransactorSession) MultiMatchOrders(leftOrders []LibOrderOrder, rightOrder LibOrderOrder, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return f.Contract.MultiMatchOrders(&f.TransactOpts, leftOrders, rightOrder, leftSignatures, rightSignature)
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

// VerifyClose is a paid mutator transaction binding the contract method 0xe8e73d08.
//
// Solidity: function verifyClose(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256) newPosition, bytes32 hash)
func (f *FuturesTransactor) VerifyClose(opts *bind.TransactOpts, positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "verifyClose", positionID, order, quantity, signature)
}

// VerifyClose is a paid mutator transaction binding the contract method 0xe8e73d08.
//
// Solidity: function verifyClose(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256) newPosition, bytes32 hash)
func (f *FuturesSession) VerifyClose(positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return f.Contract.VerifyClose(&f.TransactOpts, positionID, order, quantity, signature)
}

// VerifyClose is a paid mutator transaction binding the contract method 0xe8e73d08.
//
// Solidity: function verifyClose(uint256 positionID, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) order, uint256 quantity, bytes signature) returns((bytes32,bytes32,uint8,uint256,uint256,int256,uint256,uint256,int256) newPosition, bytes32 hash)
func (f *FuturesTransactorSession) VerifyClose(positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return f.Contract.VerifyClose(&f.TransactOpts, positionID, order, quantity, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (f *FuturesTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (f *FuturesSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Withdraw(&f.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (f *FuturesTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return f.Contract.Withdraw(&f.TransactOpts, amount)
}

// WithdrawAccount is a paid mutator transaction binding the contract method 0x4d6a5cbc.
//
// Solidity: function withdrawAccount(bytes32 accountID, uint256 amount) returns()
func (f *FuturesTransactor) WithdrawAccount(opts *bind.TransactOpts, accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.contract.Transact(opts, "withdrawAccount", accountID, amount)
}

// WithdrawAccount is a paid mutator transaction binding the contract method 0x4d6a5cbc.
//
// Solidity: function withdrawAccount(bytes32 accountID, uint256 amount) returns()
func (f *FuturesSession) WithdrawAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawAccount(&f.TransactOpts, accountID, amount)
}

// WithdrawAccount is a paid mutator transaction binding the contract method 0x4d6a5cbc.
//
// Solidity: function withdrawAccount(bytes32 accountID, uint256 amount) returns()
func (f *FuturesTransactorSession) WithdrawAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return f.Contract.WithdrawAccount(&f.TransactOpts, accountID, amount)
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
	PositionID *big.Int
	MarketID   [32]byte
	AccountID  [32]byte
	Quantity   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFuturesClose is a free log retrieval operation binding the contract event 0x792efce778d01fa86ba01f73d65f8d6cabf29a76ae688bb9c064d7ebbc4cb724.
//
// Solidity: event FuturesClose(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed accountID, uint256 quantity)
func (f *FuturesFilterer) FilterFuturesClose(opts *bind.FilterOpts, positionID []*big.Int, marketID [][32]byte, accountID [][32]byte) (*FuturesCloseIterator, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var accountIDRule []interface{}
	for _, accountIDItem := range accountID {
		accountIDRule = append(accountIDRule, accountIDItem)
	}

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesClose", positionIDRule, marketIDRule, accountIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesCloseIterator{contract: f.contract, event: "FuturesClose", logs: logs, sub: sub}, nil
}

// WatchFuturesClose is a free log subscription operation binding the contract event 0x792efce778d01fa86ba01f73d65f8d6cabf29a76ae688bb9c064d7ebbc4cb724.
//
// Solidity: event FuturesClose(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed accountID, uint256 quantity)
func (f *FuturesFilterer) WatchFuturesClose(opts *bind.WatchOpts, sink chan<- *FuturesClose, positionID []*big.Int, marketID [][32]byte, accountID [][32]byte) (event.Subscription, error) {

	var positionIDRule []interface{}
	for _, positionIDItem := range positionID {
		positionIDRule = append(positionIDRule, positionIDItem)
	}
	var marketIDRule []interface{}
	for _, marketIDItem := range marketID {
		marketIDRule = append(marketIDRule, marketIDItem)
	}
	var accountIDRule []interface{}
	for _, accountIDItem := range accountID {
		accountIDRule = append(accountIDRule, accountIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesClose", positionIDRule, marketIDRule, accountIDRule)
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

// ParseFuturesClose is a log parse operation binding the contract event 0x792efce778d01fa86ba01f73d65f8d6cabf29a76ae688bb9c064d7ebbc4cb724.
//
// Solidity: event FuturesClose(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed accountID, uint256 quantity)
func (f *FuturesFilterer) ParseFuturesClose(log types.Log) (*FuturesClose, error) {
	event := new(FuturesClose)
	if err := f.contract.UnpackLog(event, "FuturesClose", log); err != nil {
		return nil, err
	}
	return event, nil
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
func (f *FuturesFilterer) FilterFuturesFill(opts *bind.FilterOpts, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (*FuturesFillIterator, error) {

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

	logs, sub, err := f.contract.FilterLogs(opts, "FuturesFill", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return &FuturesFillIterator{contract: f.contract, event: "FuturesFill", logs: logs, sub: sub}, nil
}

// WatchFuturesFill is a free log subscription operation binding the contract event 0x6c3c91984205906d99b5675f1c6cec9054481daa142b00699944b6072a44b7e6.
//
// Solidity: event FuturesFill(address indexed makerAddress, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled, uint256 quantity)
func (f *FuturesFilterer) WatchFuturesFill(opts *bind.WatchOpts, sink chan<- *FuturesFill, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := f.contract.WatchLogs(opts, "FuturesFill", makerAddressRule, orderHashRule, marketIDRule)
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
				if err := f.contract.UnpackLog(event, "FuturesFill", log); err != nil {
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
func (f *FuturesFilterer) ParseFuturesFill(log types.Log) (*FuturesFill, error) {
	event := new(FuturesFill)
	if err := f.contract.UnpackLog(event, "FuturesFill", log); err != nil {
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

// WatchMarketCreation is a free log subscription operation binding the contract event 0xa89968288be926b3b7a2a5a0d70c9e55178a465c62e84ad1c3f0401d8c2f7fc1.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle)
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

// ParseMarketCreation is a log parse operation binding the contract event 0xa89968288be926b3b7a2a5a0d70c9e55178a465c62e84ad1c3f0401d8c2f7fc1.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle)
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
