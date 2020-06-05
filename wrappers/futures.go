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
	_ = abi.U256
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

// InjectiveFuturesABI is the input ABI used to generate the binding from.
const InjectiveFuturesABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumMarginRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContractAddressIfExists\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"}],\"name\":\"FuturesCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"FuturesClose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"FuturesFill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"MarketCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"SignatureValidatorApproval\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP1271_MAGIC_VALUE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_EXCHANGE_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accountIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"NAV\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"NPV\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressToAccountIDs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"name\":\"calcCumulativeFunding\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"calcLiquidationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"calcMinMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"closePosition\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"closePositionWithOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createAccountAndDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"}],\"name\":\"createMarket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositNewAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freeDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"fundPooledDeposits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"getBatchBalancesAndAssetProxyAllowances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"getOrderRelevantStates\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.OrderStatus\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAssetFilledAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLibOrder.OrderInfo[]\",\"name\":\"ordersInfo\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fillableTakerAssetAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"isValidSignature\",\"type\":\"bool[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getPositionIDsForTrader\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"positionIDs\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"}],\"name\":\"getPositionsForTrader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"NPV\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"internalType\":\"structTypes.Position[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidOrderSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"marketOrders\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketSerialToID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"marketToAccountToPositionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currFundingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFunding\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"leftOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"leftSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"matchOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"leftOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"leftSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"multiMatchOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pooledDeposits\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"positionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"NPV\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"preSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"restrictedDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unrestrictedDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifyClose\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"marketID\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Direction\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"NPV\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"minMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"cumulativeFundingEntry\",\"type\":\"int256\"}],\"internalType\":\"structTypes.Position\",\"name\":\"newPosition\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InjectiveFutures is an auto generated Go binding around an Ethereum contract.
type InjectiveFutures struct {
	InjectiveFuturesCaller     // Read-only binding to the contract
	InjectiveFuturesTransactor // Write-only binding to the contract
	InjectiveFuturesFilterer   // Log filterer for contract events
}

// InjectiveFuturesCaller is an auto generated read-only Go binding around an Ethereum contract.
type InjectiveFuturesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InjectiveFuturesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InjectiveFuturesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InjectiveFuturesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InjectiveFuturesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InjectiveFuturesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InjectiveFuturesSession struct {
	Contract     *InjectiveFutures // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InjectiveFuturesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InjectiveFuturesCallerSession struct {
	Contract *InjectiveFuturesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// InjectiveFuturesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InjectiveFuturesTransactorSession struct {
	Contract     *InjectiveFuturesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// InjectiveFuturesRaw is an auto generated low-level Go binding around an Ethereum contract.
type InjectiveFuturesRaw struct {
	Contract *InjectiveFutures // Generic contract binding to access the raw methods on
}

// InjectiveFuturesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InjectiveFuturesCallerRaw struct {
	Contract *InjectiveFuturesCaller // Generic read-only contract binding to access the raw methods on
}

// InjectiveFuturesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InjectiveFuturesTransactorRaw struct {
	Contract *InjectiveFuturesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInjectiveFutures creates a new instance of InjectiveFutures, bound to a specific deployed contract.
func NewInjectiveFutures(address common.Address, backend bind.ContractBackend) (*InjectiveFutures, error) {
	contract, err := bindInjectiveFutures(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InjectiveFutures{InjectiveFuturesCaller: InjectiveFuturesCaller{contract: contract}, InjectiveFuturesTransactor: InjectiveFuturesTransactor{contract: contract}, InjectiveFuturesFilterer: InjectiveFuturesFilterer{contract: contract}}, nil
}

// NewInjectiveFuturesCaller creates a new read-only instance of InjectiveFutures, bound to a specific deployed contract.
func NewInjectiveFuturesCaller(address common.Address, caller bind.ContractCaller) (*InjectiveFuturesCaller, error) {
	contract, err := bindInjectiveFutures(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InjectiveFuturesCaller{contract: contract}, nil
}

// NewInjectiveFuturesTransactor creates a new write-only instance of InjectiveFutures, bound to a specific deployed contract.
func NewInjectiveFuturesTransactor(address common.Address, transactor bind.ContractTransactor) (*InjectiveFuturesTransactor, error) {
	contract, err := bindInjectiveFutures(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InjectiveFuturesTransactor{contract: contract}, nil
}

// NewInjectiveFuturesFilterer creates a new log filterer instance of InjectiveFutures, bound to a specific deployed contract.
func NewInjectiveFuturesFilterer(address common.Address, filterer bind.ContractFilterer) (*InjectiveFuturesFilterer, error) {
	contract, err := bindInjectiveFutures(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InjectiveFuturesFilterer{contract: contract}, nil
}

// bindInjectiveFutures binds a generic wrapper to an already deployed contract.
func bindInjectiveFutures(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InjectiveFuturesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InjectiveFutures *InjectiveFuturesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InjectiveFutures.Contract.InjectiveFuturesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InjectiveFutures *InjectiveFuturesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.InjectiveFuturesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InjectiveFutures *InjectiveFuturesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.InjectiveFuturesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InjectiveFutures *InjectiveFuturesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InjectiveFutures.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InjectiveFutures *InjectiveFuturesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InjectiveFutures *InjectiveFuturesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.contract.Transact(opts, method, params...)
}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() constant returns(bytes4)
func (_InjectiveFutures *InjectiveFuturesCaller) EIP1271MAGICVALUE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "EIP1271_MAGIC_VALUE")
	return *ret0, err
}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() constant returns(bytes4)
func (_InjectiveFutures *InjectiveFuturesSession) EIP1271MAGICVALUE() ([4]byte, error) {
	return _InjectiveFutures.Contract.EIP1271MAGICVALUE(&_InjectiveFutures.CallOpts)
}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() constant returns(bytes4)
func (_InjectiveFutures *InjectiveFuturesCallerSession) EIP1271MAGICVALUE() ([4]byte, error) {
	return _InjectiveFutures.Contract.EIP1271MAGICVALUE(&_InjectiveFutures.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() constant returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesCaller) EIP712EXCHANGEDOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "EIP712_EXCHANGE_DOMAIN_HASH")
	return *ret0, err
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() constant returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return _InjectiveFutures.Contract.EIP712EXCHANGEDOMAINHASH(&_InjectiveFutures.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() constant returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesCallerSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return _InjectiveFutures.Contract.EIP712EXCHANGEDOMAINHASH(&_InjectiveFutures.CallOpts)
}

// AccountIdToAddress is a free data retrieval call binding the contract method 0x899706c0.
//
// Solidity: function accountIdToAddress(bytes32 ) constant returns(address)
func (_InjectiveFutures *InjectiveFuturesCaller) AccountIdToAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "accountIdToAddress", arg0)
	return *ret0, err
}

// AccountIdToAddress is a free data retrieval call binding the contract method 0x899706c0.
//
// Solidity: function accountIdToAddress(bytes32 ) constant returns(address)
func (_InjectiveFutures *InjectiveFuturesSession) AccountIdToAddress(arg0 [32]byte) (common.Address, error) {
	return _InjectiveFutures.Contract.AccountIdToAddress(&_InjectiveFutures.CallOpts, arg0)
}

// AccountIdToAddress is a free data retrieval call binding the contract method 0x899706c0.
//
// Solidity: function accountIdToAddress(bytes32 ) constant returns(address)
func (_InjectiveFutures *InjectiveFuturesCallerSession) AccountIdToAddress(arg0 [32]byte) (common.Address, error) {
	return _InjectiveFutures.Contract.AccountIdToAddress(&_InjectiveFutures.CallOpts, arg0)
}

// AccountNonce is a free data retrieval call binding the contract method 0x106cffe2.
//
// Solidity: function accountNonce(address ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) AccountNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "accountNonce", arg0)
	return *ret0, err
}

// AccountNonce is a free data retrieval call binding the contract method 0x106cffe2.
//
// Solidity: function accountNonce(address ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) AccountNonce(arg0 common.Address) (*big.Int, error) {
	return _InjectiveFutures.Contract.AccountNonce(&_InjectiveFutures.CallOpts, arg0)
}

// AccountNonce is a free data retrieval call binding the contract method 0x106cffe2.
//
// Solidity: function accountNonce(address ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) AccountNonce(arg0 common.Address) (*big.Int, error) {
	return _InjectiveFutures.Contract.AccountNonce(&_InjectiveFutures.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) constant returns(int256 NAV, int256 NPV, bytes32 accountID)
func (_InjectiveFutures *InjectiveFuturesCaller) Accounts(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _InjectiveFutures.contract.Call(opts, out, "accounts", arg0)
	return *ret, err
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) constant returns(int256 NAV, int256 NPV, bytes32 accountID)
func (_InjectiveFutures *InjectiveFuturesSession) Accounts(arg0 [32]byte) (struct {
	NAV       *big.Int
	NPV       *big.Int
	AccountID [32]byte
}, error) {
	return _InjectiveFutures.Contract.Accounts(&_InjectiveFutures.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) constant returns(int256 NAV, int256 NPV, bytes32 accountID)
func (_InjectiveFutures *InjectiveFuturesCallerSession) Accounts(arg0 [32]byte) (struct {
	NAV       *big.Int
	NPV       *big.Int
	AccountID [32]byte
}, error) {
	return _InjectiveFutures.Contract.Accounts(&_InjectiveFutures.CallOpts, arg0)
}

// AddressToAccountIDs is a free data retrieval call binding the contract method 0xfbd6d494.
//
// Solidity: function addressToAccountIDs(address , uint256 ) constant returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesCaller) AddressToAccountIDs(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "addressToAccountIDs", arg0, arg1)
	return *ret0, err
}

// AddressToAccountIDs is a free data retrieval call binding the contract method 0xfbd6d494.
//
// Solidity: function addressToAccountIDs(address , uint256 ) constant returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesSession) AddressToAccountIDs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _InjectiveFutures.Contract.AddressToAccountIDs(&_InjectiveFutures.CallOpts, arg0, arg1)
}

// AddressToAccountIDs is a free data retrieval call binding the contract method 0xfbd6d494.
//
// Solidity: function addressToAccountIDs(address , uint256 ) constant returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesCallerSession) AddressToAccountIDs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _InjectiveFutures.Contract.AddressToAccountIDs(&_InjectiveFutures.CallOpts, arg0, arg1)
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesCaller) AllowedValidators(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "allowedValidators", arg0, arg1)
	return *ret0, err
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesSession) AllowedValidators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _InjectiveFutures.Contract.AllowedValidators(&_InjectiveFutures.CallOpts, arg0, arg1)
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesCallerSession) AllowedValidators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _InjectiveFutures.Contract.AllowedValidators(&_InjectiveFutures.CallOpts, arg0, arg1)
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) constant returns(int256)
func (_InjectiveFutures *InjectiveFuturesCaller) CalcCumulativeFunding(opts *bind.CallOpts, marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "calcCumulativeFunding", marketID, cumulativeFundingEntry)
	return *ret0, err
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) constant returns(int256)
func (_InjectiveFutures *InjectiveFuturesSession) CalcCumulativeFunding(marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	return _InjectiveFutures.Contract.CalcCumulativeFunding(&_InjectiveFutures.CallOpts, marketID, cumulativeFundingEntry)
}

// CalcCumulativeFunding is a free data retrieval call binding the contract method 0xc5d135da.
//
// Solidity: function calcCumulativeFunding(bytes32 marketID, int256 cumulativeFundingEntry) constant returns(int256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) CalcCumulativeFunding(marketID [32]byte, cumulativeFundingEntry *big.Int) (*big.Int, error) {
	return _InjectiveFutures.Contract.CalcCumulativeFunding(&_InjectiveFutures.CallOpts, marketID, cumulativeFundingEntry)
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 _marketID, uint256 _quantity) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) CalcLiquidationFee(opts *bind.CallOpts, _marketID [32]byte, _quantity *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "calcLiquidationFee", _marketID, _quantity)
	return *ret0, err
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 _marketID, uint256 _quantity) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) CalcLiquidationFee(_marketID [32]byte, _quantity *big.Int) (*big.Int, error) {
	return _InjectiveFutures.Contract.CalcLiquidationFee(&_InjectiveFutures.CallOpts, _marketID, _quantity)
}

// CalcLiquidationFee is a free data retrieval call binding the contract method 0xbbcac0d3.
//
// Solidity: function calcLiquidationFee(bytes32 _marketID, uint256 _quantity) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) CalcLiquidationFee(_marketID [32]byte, _quantity *big.Int) (*big.Int, error) {
	return _InjectiveFutures.Contract.CalcLiquidationFee(&_InjectiveFutures.CallOpts, _marketID, _quantity)
}

// CalcMinMargin is a free data retrieval call binding the contract method 0x35c43c4e.
//
// Solidity: function calcMinMargin(bytes32 marketID, uint256 quantity, uint256 price) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) CalcMinMargin(opts *bind.CallOpts, marketID [32]byte, quantity *big.Int, price *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "calcMinMargin", marketID, quantity, price)
	return *ret0, err
}

// CalcMinMargin is a free data retrieval call binding the contract method 0x35c43c4e.
//
// Solidity: function calcMinMargin(bytes32 marketID, uint256 quantity, uint256 price) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) CalcMinMargin(marketID [32]byte, quantity *big.Int, price *big.Int) (*big.Int, error) {
	return _InjectiveFutures.Contract.CalcMinMargin(&_InjectiveFutures.CallOpts, marketID, quantity, price)
}

// CalcMinMargin is a free data retrieval call binding the contract method 0x35c43c4e.
//
// Solidity: function calcMinMargin(bytes32 marketID, uint256 quantity, uint256 price) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) CalcMinMargin(marketID [32]byte, quantity *big.Int, price *big.Int) (*big.Int, error) {
	return _InjectiveFutures.Contract.CalcMinMargin(&_InjectiveFutures.CallOpts, marketID, quantity, price)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesCaller) Cancelled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "cancelled", arg0)
	return *ret0, err
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _InjectiveFutures.Contract.Cancelled(&_InjectiveFutures.CallOpts, arg0)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesCallerSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _InjectiveFutures.Contract.Cancelled(&_InjectiveFutures.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "filled", arg0)
	return *ret0, err
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.Filled(&_InjectiveFutures.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.Filled(&_InjectiveFutures.CallOpts, arg0)
}

// FreeDeposits is a free data retrieval call binding the contract method 0x9a78538f.
//
// Solidity: function freeDeposits(address ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) FreeDeposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "freeDeposits", arg0)
	return *ret0, err
}

// FreeDeposits is a free data retrieval call binding the contract method 0x9a78538f.
//
// Solidity: function freeDeposits(address ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) FreeDeposits(arg0 common.Address) (*big.Int, error) {
	return _InjectiveFutures.Contract.FreeDeposits(&_InjectiveFutures.CallOpts, arg0)
}

// FreeDeposits is a free data retrieval call binding the contract method 0x9a78538f.
//
// Solidity: function freeDeposits(address ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) FreeDeposits(arg0 common.Address) (*big.Int, error) {
	return _InjectiveFutures.Contract.FreeDeposits(&_InjectiveFutures.CallOpts, arg0)
}

// GetBatchBalancesAndAssetProxyAllowances is a free data retrieval call binding the contract method 0x5bf34fc7.
//
// Solidity: function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, address[] tokenAddresses) constant returns(uint256[], uint256[])
func (_InjectiveFutures *InjectiveFuturesCaller) GetBatchBalancesAndAssetProxyAllowances(opts *bind.CallOpts, ownerAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _InjectiveFutures.contract.Call(opts, out, "getBatchBalancesAndAssetProxyAllowances", ownerAddress, tokenAddresses)
	return *ret0, *ret1, err
}

// GetBatchBalancesAndAssetProxyAllowances is a free data retrieval call binding the contract method 0x5bf34fc7.
//
// Solidity: function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, address[] tokenAddresses) constant returns(uint256[], uint256[])
func (_InjectiveFutures *InjectiveFuturesSession) GetBatchBalancesAndAssetProxyAllowances(ownerAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, []*big.Int, error) {
	return _InjectiveFutures.Contract.GetBatchBalancesAndAssetProxyAllowances(&_InjectiveFutures.CallOpts, ownerAddress, tokenAddresses)
}

// GetBatchBalancesAndAssetProxyAllowances is a free data retrieval call binding the contract method 0x5bf34fc7.
//
// Solidity: function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, address[] tokenAddresses) constant returns(uint256[], uint256[])
func (_InjectiveFutures *InjectiveFuturesCallerSession) GetBatchBalancesAndAssetProxyAllowances(ownerAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, []*big.Int, error) {
	return _InjectiveFutures.Contract.GetBatchBalancesAndAssetProxyAllowances(&_InjectiveFutures.CallOpts, ownerAddress, tokenAddresses)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates([]LibOrderOrder orders, bytes[] signatures) constant returns([]LibOrderOrderInfo ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (_InjectiveFutures *InjectiveFuturesCaller) GetOrderRelevantStates(opts *bind.CallOpts, orders []LibOrderOrder, signatures [][]byte) (struct {
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
	err := _InjectiveFutures.contract.Call(opts, out, "getOrderRelevantStates", orders, signatures)
	return *ret, err
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates([]LibOrderOrder orders, bytes[] signatures) constant returns([]LibOrderOrderInfo ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (_InjectiveFutures *InjectiveFuturesSession) GetOrderRelevantStates(orders []LibOrderOrder, signatures [][]byte) (struct {
	OrdersInfo                []LibOrderOrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	return _InjectiveFutures.Contract.GetOrderRelevantStates(&_InjectiveFutures.CallOpts, orders, signatures)
}

// GetOrderRelevantStates is a free data retrieval call binding the contract method 0xe25cabf7.
//
// Solidity: function getOrderRelevantStates([]LibOrderOrder orders, bytes[] signatures) constant returns([]LibOrderOrderInfo ordersInfo, uint256[] fillableTakerAssetAmounts, bool[] isValidSignature)
func (_InjectiveFutures *InjectiveFuturesCallerSession) GetOrderRelevantStates(orders []LibOrderOrder, signatures [][]byte) (struct {
	OrdersInfo                []LibOrderOrderInfo
	FillableTakerAssetAmounts []*big.Int
	IsValidSignature          []bool
}, error) {
	return _InjectiveFutures.Contract.GetOrderRelevantStates(&_InjectiveFutures.CallOpts, orders, signatures)
}

// GetPositionIDsForTrader is a free data retrieval call binding the contract method 0xc3e49bb7.
//
// Solidity: function getPositionIDsForTrader(address trader, bytes32 marketID) constant returns(uint256[] positionIDs)
func (_InjectiveFutures *InjectiveFuturesCaller) GetPositionIDsForTrader(opts *bind.CallOpts, trader common.Address, marketID [32]byte) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "getPositionIDsForTrader", trader, marketID)
	return *ret0, err
}

// GetPositionIDsForTrader is a free data retrieval call binding the contract method 0xc3e49bb7.
//
// Solidity: function getPositionIDsForTrader(address trader, bytes32 marketID) constant returns(uint256[] positionIDs)
func (_InjectiveFutures *InjectiveFuturesSession) GetPositionIDsForTrader(trader common.Address, marketID [32]byte) ([]*big.Int, error) {
	return _InjectiveFutures.Contract.GetPositionIDsForTrader(&_InjectiveFutures.CallOpts, trader, marketID)
}

// GetPositionIDsForTrader is a free data retrieval call binding the contract method 0xc3e49bb7.
//
// Solidity: function getPositionIDsForTrader(address trader, bytes32 marketID) constant returns(uint256[] positionIDs)
func (_InjectiveFutures *InjectiveFuturesCallerSession) GetPositionIDsForTrader(trader common.Address, marketID [32]byte) ([]*big.Int, error) {
	return _InjectiveFutures.Contract.GetPositionIDsForTrader(&_InjectiveFutures.CallOpts, trader, marketID)
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) constant returns([]TypesPosition)
func (_InjectiveFutures *InjectiveFuturesCaller) GetPositionsForTrader(opts *bind.CallOpts, trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	var (
		ret0 = new([]TypesPosition)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "getPositionsForTrader", trader, marketID)
	return *ret0, err
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) constant returns([]TypesPosition)
func (_InjectiveFutures *InjectiveFuturesSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return _InjectiveFutures.Contract.GetPositionsForTrader(&_InjectiveFutures.CallOpts, trader, marketID)
}

// GetPositionsForTrader is a free data retrieval call binding the contract method 0x0088e8cc.
//
// Solidity: function getPositionsForTrader(address trader, bytes32 marketID) constant returns([]TypesPosition)
func (_InjectiveFutures *InjectiveFuturesCallerSession) GetPositionsForTrader(trader common.Address, marketID [32]byte) ([]TypesPosition, error) {
	return _InjectiveFutures.Contract.GetPositionsForTrader(&_InjectiveFutures.CallOpts, trader, marketID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesSession) IsOwner() (bool, error) {
	return _InjectiveFutures.Contract.IsOwner(&_InjectiveFutures.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesCallerSession) IsOwner() (bool, error) {
	return _InjectiveFutures.Contract.IsOwner(&_InjectiveFutures.CallOpts)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature(LibOrderOrder order, bytes signature) constant returns(bool isValid)
func (_InjectiveFutures *InjectiveFuturesCaller) IsValidOrderSignature(opts *bind.CallOpts, order LibOrderOrder, signature []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "isValidOrderSignature", order, signature)
	return *ret0, err
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature(LibOrderOrder order, bytes signature) constant returns(bool isValid)
func (_InjectiveFutures *InjectiveFuturesSession) IsValidOrderSignature(order LibOrderOrder, signature []byte) (bool, error) {
	return _InjectiveFutures.Contract.IsValidOrderSignature(&_InjectiveFutures.CallOpts, order, signature)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature(LibOrderOrder order, bytes signature) constant returns(bool isValid)
func (_InjectiveFutures *InjectiveFuturesCallerSession) IsValidOrderSignature(order LibOrderOrder, signature []byte) (bool, error) {
	return _InjectiveFutures.Contract.IsValidOrderSignature(&_InjectiveFutures.CallOpts, order, signature)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "marketCount")
	return *ret0, err
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) MarketCount() (*big.Int, error) {
	return _InjectiveFutures.Contract.MarketCount(&_InjectiveFutures.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) MarketCount() (*big.Int, error) {
	return _InjectiveFutures.Contract.MarketCount(&_InjectiveFutures.CallOpts)
}

// MarketSerialToID is a free data retrieval call binding the contract method 0xbae18473.
//
// Solidity: function marketSerialToID(uint256 ) constant returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesCaller) MarketSerialToID(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "marketSerialToID", arg0)
	return *ret0, err
}

// MarketSerialToID is a free data retrieval call binding the contract method 0xbae18473.
//
// Solidity: function marketSerialToID(uint256 ) constant returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesSession) MarketSerialToID(arg0 *big.Int) ([32]byte, error) {
	return _InjectiveFutures.Contract.MarketSerialToID(&_InjectiveFutures.CallOpts, arg0)
}

// MarketSerialToID is a free data retrieval call binding the contract method 0xbae18473.
//
// Solidity: function marketSerialToID(uint256 ) constant returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesCallerSession) MarketSerialToID(arg0 *big.Int) ([32]byte, error) {
	return _InjectiveFutures.Contract.MarketSerialToID(&_InjectiveFutures.CallOpts, arg0)
}

// MarketToAccountToPositionID is a free data retrieval call binding the contract method 0x54971478.
//
// Solidity: function marketToAccountToPositionID(bytes32 , bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) MarketToAccountToPositionID(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "marketToAccountToPositionID", arg0, arg1)
	return *ret0, err
}

// MarketToAccountToPositionID is a free data retrieval call binding the contract method 0x54971478.
//
// Solidity: function marketToAccountToPositionID(bytes32 , bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) MarketToAccountToPositionID(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.MarketToAccountToPositionID(&_InjectiveFutures.CallOpts, arg0, arg1)
}

// MarketToAccountToPositionID is a free data retrieval call binding the contract method 0x54971478.
//
// Solidity: function marketToAccountToPositionID(bytes32 , bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) MarketToAccountToPositionID(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.MarketToAccountToPositionID(&_InjectiveFutures.CallOpts, arg0, arg1)
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) constant returns(bytes32 marketID, string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 indexPrice, uint256 currFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding)
func (_InjectiveFutures *InjectiveFuturesCaller) Markets(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _InjectiveFutures.contract.Call(opts, out, "markets", arg0)
	return *ret, err
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) constant returns(bytes32 marketID, string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 indexPrice, uint256 currFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding)
func (_InjectiveFutures *InjectiveFuturesSession) Markets(arg0 [32]byte) (struct {
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
	return _InjectiveFutures.Contract.Markets(&_InjectiveFutures.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x7564912b.
//
// Solidity: function markets(bytes32 ) constant returns(bytes32 marketID, string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 indexPrice, uint256 currFundingTimestamp, uint256 fundingInterval, int256 cumulativeFunding)
func (_InjectiveFutures *InjectiveFuturesCallerSession) Markets(arg0 [32]byte) (struct {
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
	return _InjectiveFutures.Contract.Markets(&_InjectiveFutures.CallOpts, arg0)
}

// OrderPosition is a free data retrieval call binding the contract method 0x9a9137ce.
//
// Solidity: function orderPosition(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) OrderPosition(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "orderPosition", arg0)
	return *ret0, err
}

// OrderPosition is a free data retrieval call binding the contract method 0x9a9137ce.
//
// Solidity: function orderPosition(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) OrderPosition(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.OrderPosition(&_InjectiveFutures.CallOpts, arg0)
}

// OrderPosition is a free data retrieval call binding the contract method 0x9a9137ce.
//
// Solidity: function orderPosition(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) OrderPosition(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.OrderPosition(&_InjectiveFutures.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_InjectiveFutures *InjectiveFuturesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_InjectiveFutures *InjectiveFuturesSession) Owner() (common.Address, error) {
	return _InjectiveFutures.Contract.Owner(&_InjectiveFutures.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_InjectiveFutures *InjectiveFuturesCallerSession) Owner() (common.Address, error) {
	return _InjectiveFutures.Contract.Owner(&_InjectiveFutures.CallOpts)
}

// PooledDeposits is a free data retrieval call binding the contract method 0xef2bcc0f.
//
// Solidity: function pooledDeposits(bytes32 ) constant returns(int256)
func (_InjectiveFutures *InjectiveFuturesCaller) PooledDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "pooledDeposits", arg0)
	return *ret0, err
}

// PooledDeposits is a free data retrieval call binding the contract method 0xef2bcc0f.
//
// Solidity: function pooledDeposits(bytes32 ) constant returns(int256)
func (_InjectiveFutures *InjectiveFuturesSession) PooledDeposits(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.PooledDeposits(&_InjectiveFutures.CallOpts, arg0)
}

// PooledDeposits is a free data retrieval call binding the contract method 0xef2bcc0f.
//
// Solidity: function pooledDeposits(bytes32 ) constant returns(int256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) PooledDeposits(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.PooledDeposits(&_InjectiveFutures.CallOpts, arg0)
}

// PositionCount is a free data retrieval call binding the contract method 0xe7702d05.
//
// Solidity: function positionCount() constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) PositionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "positionCount")
	return *ret0, err
}

// PositionCount is a free data retrieval call binding the contract method 0xe7702d05.
//
// Solidity: function positionCount() constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) PositionCount() (*big.Int, error) {
	return _InjectiveFutures.Contract.PositionCount(&_InjectiveFutures.CallOpts)
}

// PositionCount is a free data retrieval call binding the contract method 0xe7702d05.
//
// Solidity: function positionCount() constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) PositionCount() (*big.Int, error) {
	return _InjectiveFutures.Contract.PositionCount(&_InjectiveFutures.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) constant returns(bytes32 accountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, int256 NPV, uint256 minMargin, uint256 margin, int256 cumulativeFundingEntry)
func (_InjectiveFutures *InjectiveFuturesCaller) Positions(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _InjectiveFutures.contract.Call(opts, out, "positions", arg0)
	return *ret, err
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) constant returns(bytes32 accountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, int256 NPV, uint256 minMargin, uint256 margin, int256 cumulativeFundingEntry)
func (_InjectiveFutures *InjectiveFuturesSession) Positions(arg0 *big.Int) (struct {
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
	return _InjectiveFutures.Contract.Positions(&_InjectiveFutures.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) constant returns(bytes32 accountID, bytes32 marketID, uint8 direction, uint256 quantity, uint256 contractPrice, int256 NPV, uint256 minMargin, uint256 margin, int256 cumulativeFundingEntry)
func (_InjectiveFutures *InjectiveFuturesCallerSession) Positions(arg0 *big.Int) (struct {
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
	return _InjectiveFutures.Contract.Positions(&_InjectiveFutures.CallOpts, arg0)
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesCaller) PreSigned(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "preSigned", arg0, arg1)
	return *ret0, err
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesSession) PreSigned(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _InjectiveFutures.Contract.PreSigned(&_InjectiveFutures.CallOpts, arg0, arg1)
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) constant returns(bool)
func (_InjectiveFutures *InjectiveFuturesCallerSession) PreSigned(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _InjectiveFutures.Contract.PreSigned(&_InjectiveFutures.CallOpts, arg0, arg1)
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x9a009c3f.
//
// Solidity: function restrictedDeposits(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) RestrictedDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "restrictedDeposits", arg0)
	return *ret0, err
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x9a009c3f.
//
// Solidity: function restrictedDeposits(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) RestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.RestrictedDeposits(&_InjectiveFutures.CallOpts, arg0)
}

// RestrictedDeposits is a free data retrieval call binding the contract method 0x9a009c3f.
//
// Solidity: function restrictedDeposits(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) RestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.RestrictedDeposits(&_InjectiveFutures.CallOpts, arg0)
}

// UnrestrictedDeposits is a free data retrieval call binding the contract method 0xa890bb84.
//
// Solidity: function unrestrictedDeposits(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCaller) UnrestrictedDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InjectiveFutures.contract.Call(opts, out, "unrestrictedDeposits", arg0)
	return *ret0, err
}

// UnrestrictedDeposits is a free data retrieval call binding the contract method 0xa890bb84.
//
// Solidity: function unrestrictedDeposits(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesSession) UnrestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.UnrestrictedDeposits(&_InjectiveFutures.CallOpts, arg0)
}

// UnrestrictedDeposits is a free data retrieval call binding the contract method 0xa890bb84.
//
// Solidity: function unrestrictedDeposits(bytes32 ) constant returns(uint256)
func (_InjectiveFutures *InjectiveFuturesCallerSession) UnrestrictedDeposits(arg0 [32]byte) (*big.Int, error) {
	return _InjectiveFutures.Contract.UnrestrictedDeposits(&_InjectiveFutures.CallOpts, arg0)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder(LibOrderOrder order) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) CancelOrder(opts *bind.TransactOpts, order LibOrderOrder) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder(LibOrderOrder order) returns()
func (_InjectiveFutures *InjectiveFuturesSession) CancelOrder(order LibOrderOrder) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.CancelOrder(&_InjectiveFutures.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder(LibOrderOrder order) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) CancelOrder(order LibOrderOrder) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.CancelOrder(&_InjectiveFutures.TransactOpts, order)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xb012d424.
//
// Solidity: function closePosition(uint256 positionID, LibOrderOrder order, uint256 quantity, bytes signature) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesTransactor) ClosePosition(opts *bind.TransactOpts, positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "closePosition", positionID, order, quantity, signature)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xb012d424.
//
// Solidity: function closePosition(uint256 positionID, LibOrderOrder order, uint256 quantity, bytes signature) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesSession) ClosePosition(positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.ClosePosition(&_InjectiveFutures.TransactOpts, positionID, order, quantity, signature)
}

// ClosePosition is a paid mutator transaction binding the contract method 0xb012d424.
//
// Solidity: function closePosition(uint256 positionID, LibOrderOrder order, uint256 quantity, bytes signature) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesTransactorSession) ClosePosition(positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.ClosePosition(&_InjectiveFutures.TransactOpts, positionID, order, quantity, signature)
}

// ClosePositionWithOrders is a paid mutator transaction binding the contract method 0xeeb3da17.
//
// Solidity: function closePositionWithOrders(uint256 positionID, []LibOrderOrder orders, uint256 quantity, bytes signature) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) ClosePositionWithOrders(opts *bind.TransactOpts, positionID *big.Int, orders []LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "closePositionWithOrders", positionID, orders, quantity, signature)
}

// ClosePositionWithOrders is a paid mutator transaction binding the contract method 0xeeb3da17.
//
// Solidity: function closePositionWithOrders(uint256 positionID, []LibOrderOrder orders, uint256 quantity, bytes signature) returns()
func (_InjectiveFutures *InjectiveFuturesSession) ClosePositionWithOrders(positionID *big.Int, orders []LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.ClosePositionWithOrders(&_InjectiveFutures.TransactOpts, positionID, orders, quantity, signature)
}

// ClosePositionWithOrders is a paid mutator transaction binding the contract method 0xeeb3da17.
//
// Solidity: function closePositionWithOrders(uint256 positionID, []LibOrderOrder orders, uint256 quantity, bytes signature) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) ClosePositionWithOrders(positionID *big.Int, orders []LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.ClosePositionWithOrders(&_InjectiveFutures.TransactOpts, positionID, orders, quantity, signature)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bytes32 accountID)
func (_InjectiveFutures *InjectiveFuturesTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bytes32 accountID)
func (_InjectiveFutures *InjectiveFuturesSession) CreateAccount() (*types.Transaction, error) {
	return _InjectiveFutures.Contract.CreateAccount(&_InjectiveFutures.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bytes32 accountID)
func (_InjectiveFutures *InjectiveFuturesTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _InjectiveFutures.Contract.CreateAccount(&_InjectiveFutures.TransactOpts)
}

// CreateAccountAndDeposit is a paid mutator transaction binding the contract method 0xa68a4c3d.
//
// Solidity: function createAccountAndDeposit(uint256 amount) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesTransactor) CreateAccountAndDeposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "createAccountAndDeposit", amount)
}

// CreateAccountAndDeposit is a paid mutator transaction binding the contract method 0xa68a4c3d.
//
// Solidity: function createAccountAndDeposit(uint256 amount) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesSession) CreateAccountAndDeposit(amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.CreateAccountAndDeposit(&_InjectiveFutures.TransactOpts, amount)
}

// CreateAccountAndDeposit is a paid mutator transaction binding the contract method 0xa68a4c3d.
//
// Solidity: function createAccountAndDeposit(uint256 amount) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesTransactorSession) CreateAccountAndDeposit(amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.CreateAccountAndDeposit(&_InjectiveFutures.TransactOpts, amount)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x1af69e5f.
//
// Solidity: function createMarket(string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 fundingInterval) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) CreateMarket(opts *bind.TransactOpts, ticker string, oracle common.Address, initialMarginRatio *big.Int, liquidationPenalty *big.Int, fundingInterval *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "createMarket", ticker, oracle, initialMarginRatio, liquidationPenalty, fundingInterval)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x1af69e5f.
//
// Solidity: function createMarket(string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 fundingInterval) returns()
func (_InjectiveFutures *InjectiveFuturesSession) CreateMarket(ticker string, oracle common.Address, initialMarginRatio *big.Int, liquidationPenalty *big.Int, fundingInterval *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.CreateMarket(&_InjectiveFutures.TransactOpts, ticker, oracle, initialMarginRatio, liquidationPenalty, fundingInterval)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x1af69e5f.
//
// Solidity: function createMarket(string ticker, address oracle, uint256 initialMarginRatio, uint256 liquidationPenalty, uint256 fundingInterval) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) CreateMarket(ticker string, oracle common.Address, initialMarginRatio *big.Int, liquidationPenalty *big.Int, fundingInterval *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.CreateMarket(&_InjectiveFutures.TransactOpts, ticker, oracle, initialMarginRatio, liquidationPenalty, fundingInterval)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.Deposit(&_InjectiveFutures.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.Deposit(&_InjectiveFutures.TransactOpts, amount)
}

// DepositAccount is a paid mutator transaction binding the contract method 0x04c19cf0.
//
// Solidity: function depositAccount(bytes32 accountID, uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) DepositAccount(opts *bind.TransactOpts, accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "depositAccount", accountID, amount)
}

// DepositAccount is a paid mutator transaction binding the contract method 0x04c19cf0.
//
// Solidity: function depositAccount(bytes32 accountID, uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesSession) DepositAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.DepositAccount(&_InjectiveFutures.TransactOpts, accountID, amount)
}

// DepositAccount is a paid mutator transaction binding the contract method 0x04c19cf0.
//
// Solidity: function depositAccount(bytes32 accountID, uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) DepositAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.DepositAccount(&_InjectiveFutures.TransactOpts, accountID, amount)
}

// DepositNewAccount is a paid mutator transaction binding the contract method 0x4a692978.
//
// Solidity: function depositNewAccount(bytes32 accountID, uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) DepositNewAccount(opts *bind.TransactOpts, accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "depositNewAccount", accountID, amount)
}

// DepositNewAccount is a paid mutator transaction binding the contract method 0x4a692978.
//
// Solidity: function depositNewAccount(bytes32 accountID, uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesSession) DepositNewAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.DepositNewAccount(&_InjectiveFutures.TransactOpts, accountID, amount)
}

// DepositNewAccount is a paid mutator transaction binding the contract method 0x4a692978.
//
// Solidity: function depositNewAccount(bytes32 accountID, uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) DepositNewAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.DepositNewAccount(&_InjectiveFutures.TransactOpts, accountID, amount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x833b2ea5.
//
// Solidity: function fillOrder(LibOrderOrder order, uint256 quantity, uint256 margin, bytes signature) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesTransactor) FillOrder(opts *bind.TransactOpts, order LibOrderOrder, quantity *big.Int, margin *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "fillOrder", order, quantity, margin, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x833b2ea5.
//
// Solidity: function fillOrder(LibOrderOrder order, uint256 quantity, uint256 margin, bytes signature) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesSession) FillOrder(order LibOrderOrder, quantity *big.Int, margin *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.FillOrder(&_InjectiveFutures.TransactOpts, order, quantity, margin, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x833b2ea5.
//
// Solidity: function fillOrder(LibOrderOrder order, uint256 quantity, uint256 margin, bytes signature) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesTransactorSession) FillOrder(order LibOrderOrder, quantity *big.Int, margin *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.FillOrder(&_InjectiveFutures.TransactOpts, order, quantity, margin, signature)
}

// FundPooledDeposits is a paid mutator transaction binding the contract method 0x8f7e9459.
//
// Solidity: function fundPooledDeposits(uint256 amount, bytes32 marketID) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) FundPooledDeposits(opts *bind.TransactOpts, amount *big.Int, marketID [32]byte) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "fundPooledDeposits", amount, marketID)
}

// FundPooledDeposits is a paid mutator transaction binding the contract method 0x8f7e9459.
//
// Solidity: function fundPooledDeposits(uint256 amount, bytes32 marketID) returns()
func (_InjectiveFutures *InjectiveFuturesSession) FundPooledDeposits(amount *big.Int, marketID [32]byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.FundPooledDeposits(&_InjectiveFutures.TransactOpts, amount, marketID)
}

// FundPooledDeposits is a paid mutator transaction binding the contract method 0x8f7e9459.
//
// Solidity: function fundPooledDeposits(uint256 amount, bytes32 marketID) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) FundPooledDeposits(amount *big.Int, marketID [32]byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.FundPooledDeposits(&_InjectiveFutures.TransactOpts, amount, marketID)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xd440e9b6.
//
// Solidity: function marketOrders([]LibOrderOrder orders, uint256 quantity, uint256 margin, bytes[] signatures) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesTransactor) MarketOrders(opts *bind.TransactOpts, orders []LibOrderOrder, quantity *big.Int, margin *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "marketOrders", orders, quantity, margin, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xd440e9b6.
//
// Solidity: function marketOrders([]LibOrderOrder orders, uint256 quantity, uint256 margin, bytes[] signatures) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesSession) MarketOrders(orders []LibOrderOrder, quantity *big.Int, margin *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.MarketOrders(&_InjectiveFutures.TransactOpts, orders, quantity, margin, signatures)
}

// MarketOrders is a paid mutator transaction binding the contract method 0xd440e9b6.
//
// Solidity: function marketOrders([]LibOrderOrder orders, uint256 quantity, uint256 margin, bytes[] signatures) returns(bytes32)
func (_InjectiveFutures *InjectiveFuturesTransactorSession) MarketOrders(orders []LibOrderOrder, quantity *big.Int, margin *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.MarketOrders(&_InjectiveFutures.TransactOpts, orders, quantity, margin, signatures)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders(LibOrderOrder leftOrder, LibOrderOrder rightOrder, bytes leftSignature, bytes rightSignature) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) MatchOrders(opts *bind.TransactOpts, leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "matchOrders", leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders(LibOrderOrder leftOrder, LibOrderOrder rightOrder, bytes leftSignature, bytes rightSignature) returns()
func (_InjectiveFutures *InjectiveFuturesSession) MatchOrders(leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.MatchOrders(&_InjectiveFutures.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders(LibOrderOrder leftOrder, LibOrderOrder rightOrder, bytes leftSignature, bytes rightSignature) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) MatchOrders(leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.MatchOrders(&_InjectiveFutures.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders([]LibOrderOrder leftOrders, LibOrderOrder rightOrder, bytes[] leftSignatures, bytes rightSignature) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) MultiMatchOrders(opts *bind.TransactOpts, leftOrders []LibOrderOrder, rightOrder LibOrderOrder, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "multiMatchOrders", leftOrders, rightOrder, leftSignatures, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders([]LibOrderOrder leftOrders, LibOrderOrder rightOrder, bytes[] leftSignatures, bytes rightSignature) returns()
func (_InjectiveFutures *InjectiveFuturesSession) MultiMatchOrders(leftOrders []LibOrderOrder, rightOrder LibOrderOrder, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.MultiMatchOrders(&_InjectiveFutures.TransactOpts, leftOrders, rightOrder, leftSignatures, rightSignature)
}

// MultiMatchOrders is a paid mutator transaction binding the contract method 0x86d7729c.
//
// Solidity: function multiMatchOrders([]LibOrderOrder leftOrders, LibOrderOrder rightOrder, bytes[] leftSignatures, bytes rightSignature) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) MultiMatchOrders(leftOrders []LibOrderOrder, rightOrder LibOrderOrder, leftSignatures [][]byte, rightSignature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.MultiMatchOrders(&_InjectiveFutures.TransactOpts, leftOrders, rightOrder, leftSignatures, rightSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InjectiveFutures *InjectiveFuturesSession) RenounceOwnership() (*types.Transaction, error) {
	return _InjectiveFutures.Contract.RenounceOwnership(&_InjectiveFutures.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InjectiveFutures.Contract.RenounceOwnership(&_InjectiveFutures.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InjectiveFutures *InjectiveFuturesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.TransferOwnership(&_InjectiveFutures.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.TransferOwnership(&_InjectiveFutures.TransactOpts, newOwner)
}

// VerifyClose is a paid mutator transaction binding the contract method 0xe8e73d08.
//
// Solidity: function verifyClose(uint256 positionID, LibOrderOrder order, uint256 quantity, bytes signature) returns(TypesPosition newPosition, bytes32 hash)
func (_InjectiveFutures *InjectiveFuturesTransactor) VerifyClose(opts *bind.TransactOpts, positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "verifyClose", positionID, order, quantity, signature)
}

// VerifyClose is a paid mutator transaction binding the contract method 0xe8e73d08.
//
// Solidity: function verifyClose(uint256 positionID, LibOrderOrder order, uint256 quantity, bytes signature) returns(TypesPosition newPosition, bytes32 hash)
func (_InjectiveFutures *InjectiveFuturesSession) VerifyClose(positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.VerifyClose(&_InjectiveFutures.TransactOpts, positionID, order, quantity, signature)
}

// VerifyClose is a paid mutator transaction binding the contract method 0xe8e73d08.
//
// Solidity: function verifyClose(uint256 positionID, LibOrderOrder order, uint256 quantity, bytes signature) returns(TypesPosition newPosition, bytes32 hash)
func (_InjectiveFutures *InjectiveFuturesTransactorSession) VerifyClose(positionID *big.Int, order LibOrderOrder, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.VerifyClose(&_InjectiveFutures.TransactOpts, positionID, order, quantity, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.Withdraw(&_InjectiveFutures.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.Withdraw(&_InjectiveFutures.TransactOpts, amount)
}

// WithdrawAccount is a paid mutator transaction binding the contract method 0x4d6a5cbc.
//
// Solidity: function withdrawAccount(bytes32 accountID, uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactor) WithdrawAccount(opts *bind.TransactOpts, accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.contract.Transact(opts, "withdrawAccount", accountID, amount)
}

// WithdrawAccount is a paid mutator transaction binding the contract method 0x4d6a5cbc.
//
// Solidity: function withdrawAccount(bytes32 accountID, uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesSession) WithdrawAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.WithdrawAccount(&_InjectiveFutures.TransactOpts, accountID, amount)
}

// WithdrawAccount is a paid mutator transaction binding the contract method 0x4d6a5cbc.
//
// Solidity: function withdrawAccount(bytes32 accountID, uint256 amount) returns()
func (_InjectiveFutures *InjectiveFuturesTransactorSession) WithdrawAccount(accountID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _InjectiveFutures.Contract.WithdrawAccount(&_InjectiveFutures.TransactOpts, accountID, amount)
}

// InjectiveFuturesFuturesCancelIterator is returned from FilterFuturesCancel and is used to iterate over the raw logs and unpacked data for FuturesCancel events raised by the InjectiveFutures contract.
type InjectiveFuturesFuturesCancelIterator struct {
	Event *InjectiveFuturesFuturesCancel // Event containing the contract specifics and raw log

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
func (it *InjectiveFuturesFuturesCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InjectiveFuturesFuturesCancel)
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
		it.Event = new(InjectiveFuturesFuturesCancel)
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
func (it *InjectiveFuturesFuturesCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InjectiveFuturesFuturesCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InjectiveFuturesFuturesCancel represents a FuturesCancel event raised by the InjectiveFutures contract.
type InjectiveFuturesFuturesCancel struct {
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
func (_InjectiveFutures *InjectiveFuturesFilterer) FilterFuturesCancel(opts *bind.FilterOpts, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (*InjectiveFuturesFuturesCancelIterator, error) {

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

	logs, sub, err := _InjectiveFutures.contract.FilterLogs(opts, "FuturesCancel", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return &InjectiveFuturesFuturesCancelIterator{contract: _InjectiveFutures.contract, event: "FuturesCancel", logs: logs, sub: sub}, nil
}

// WatchFuturesCancel is a free log subscription operation binding the contract event 0x414118d90fd71dbfe3eebc508a8edaebe20d4e43ac23c65ba56fe87edb7c61ca.
//
// Solidity: event FuturesCancel(address indexed makerAddress, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled)
func (_InjectiveFutures *InjectiveFuturesFilterer) WatchFuturesCancel(opts *bind.WatchOpts, sink chan<- *InjectiveFuturesFuturesCancel, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _InjectiveFutures.contract.WatchLogs(opts, "FuturesCancel", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InjectiveFuturesFuturesCancel)
				if err := _InjectiveFutures.contract.UnpackLog(event, "FuturesCancel", log); err != nil {
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
func (_InjectiveFutures *InjectiveFuturesFilterer) ParseFuturesCancel(log types.Log) (*InjectiveFuturesFuturesCancel, error) {
	event := new(InjectiveFuturesFuturesCancel)
	if err := _InjectiveFutures.contract.UnpackLog(event, "FuturesCancel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InjectiveFuturesFuturesCloseIterator is returned from FilterFuturesClose and is used to iterate over the raw logs and unpacked data for FuturesClose events raised by the InjectiveFutures contract.
type InjectiveFuturesFuturesCloseIterator struct {
	Event *InjectiveFuturesFuturesClose // Event containing the contract specifics and raw log

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
func (it *InjectiveFuturesFuturesCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InjectiveFuturesFuturesClose)
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
		it.Event = new(InjectiveFuturesFuturesClose)
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
func (it *InjectiveFuturesFuturesCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InjectiveFuturesFuturesCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InjectiveFuturesFuturesClose represents a FuturesClose event raised by the InjectiveFutures contract.
type InjectiveFuturesFuturesClose struct {
	PositionID *big.Int
	MarketID   [32]byte
	AccountID  [32]byte
	Quantity   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFuturesClose is a free log retrieval operation binding the contract event 0x792efce778d01fa86ba01f73d65f8d6cabf29a76ae688bb9c064d7ebbc4cb724.
//
// Solidity: event FuturesClose(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed accountID, uint256 quantity)
func (_InjectiveFutures *InjectiveFuturesFilterer) FilterFuturesClose(opts *bind.FilterOpts, positionID []*big.Int, marketID [][32]byte, accountID [][32]byte) (*InjectiveFuturesFuturesCloseIterator, error) {

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

	logs, sub, err := _InjectiveFutures.contract.FilterLogs(opts, "FuturesClose", positionIDRule, marketIDRule, accountIDRule)
	if err != nil {
		return nil, err
	}
	return &InjectiveFuturesFuturesCloseIterator{contract: _InjectiveFutures.contract, event: "FuturesClose", logs: logs, sub: sub}, nil
}

// WatchFuturesClose is a free log subscription operation binding the contract event 0x792efce778d01fa86ba01f73d65f8d6cabf29a76ae688bb9c064d7ebbc4cb724.
//
// Solidity: event FuturesClose(uint256 indexed positionID, bytes32 indexed marketID, bytes32 indexed accountID, uint256 quantity)
func (_InjectiveFutures *InjectiveFuturesFilterer) WatchFuturesClose(opts *bind.WatchOpts, sink chan<- *InjectiveFuturesFuturesClose, positionID []*big.Int, marketID [][32]byte, accountID [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _InjectiveFutures.contract.WatchLogs(opts, "FuturesClose", positionIDRule, marketIDRule, accountIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InjectiveFuturesFuturesClose)
				if err := _InjectiveFutures.contract.UnpackLog(event, "FuturesClose", log); err != nil {
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
func (_InjectiveFutures *InjectiveFuturesFilterer) ParseFuturesClose(log types.Log) (*InjectiveFuturesFuturesClose, error) {
	event := new(InjectiveFuturesFuturesClose)
	if err := _InjectiveFutures.contract.UnpackLog(event, "FuturesClose", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InjectiveFuturesFuturesFillIterator is returned from FilterFuturesFill and is used to iterate over the raw logs and unpacked data for FuturesFill events raised by the InjectiveFutures contract.
type InjectiveFuturesFuturesFillIterator struct {
	Event *InjectiveFuturesFuturesFill // Event containing the contract specifics and raw log

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
func (it *InjectiveFuturesFuturesFillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InjectiveFuturesFuturesFill)
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
		it.Event = new(InjectiveFuturesFuturesFill)
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
func (it *InjectiveFuturesFuturesFillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InjectiveFuturesFuturesFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InjectiveFuturesFuturesFill represents a FuturesFill event raised by the InjectiveFutures contract.
type InjectiveFuturesFuturesFill struct {
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
func (_InjectiveFutures *InjectiveFuturesFilterer) FilterFuturesFill(opts *bind.FilterOpts, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (*InjectiveFuturesFuturesFillIterator, error) {

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

	logs, sub, err := _InjectiveFutures.contract.FilterLogs(opts, "FuturesFill", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return &InjectiveFuturesFuturesFillIterator{contract: _InjectiveFutures.contract, event: "FuturesFill", logs: logs, sub: sub}, nil
}

// WatchFuturesFill is a free log subscription operation binding the contract event 0x6c3c91984205906d99b5675f1c6cec9054481daa142b00699944b6072a44b7e6.
//
// Solidity: event FuturesFill(address indexed makerAddress, bytes32 indexed orderHash, bytes32 indexed marketID, uint256 contractPrice, uint256 quantityFilled, uint256 quantity)
func (_InjectiveFutures *InjectiveFuturesFilterer) WatchFuturesFill(opts *bind.WatchOpts, sink chan<- *InjectiveFuturesFuturesFill, makerAddress []common.Address, orderHash [][32]byte, marketID [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _InjectiveFutures.contract.WatchLogs(opts, "FuturesFill", makerAddressRule, orderHashRule, marketIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InjectiveFuturesFuturesFill)
				if err := _InjectiveFutures.contract.UnpackLog(event, "FuturesFill", log); err != nil {
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
func (_InjectiveFutures *InjectiveFuturesFilterer) ParseFuturesFill(log types.Log) (*InjectiveFuturesFuturesFill, error) {
	event := new(InjectiveFuturesFuturesFill)
	if err := _InjectiveFutures.contract.UnpackLog(event, "FuturesFill", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InjectiveFuturesMarketCreationIterator is returned from FilterMarketCreation and is used to iterate over the raw logs and unpacked data for MarketCreation events raised by the InjectiveFutures contract.
type InjectiveFuturesMarketCreationIterator struct {
	Event *InjectiveFuturesMarketCreation // Event containing the contract specifics and raw log

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
func (it *InjectiveFuturesMarketCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InjectiveFuturesMarketCreation)
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
		it.Event = new(InjectiveFuturesMarketCreation)
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
func (it *InjectiveFuturesMarketCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InjectiveFuturesMarketCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InjectiveFuturesMarketCreation represents a MarketCreation event raised by the InjectiveFutures contract.
type InjectiveFuturesMarketCreation struct {
	MarketID [32]byte
	Ticker   common.Hash
	Oracle   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketCreation is a free log retrieval operation binding the contract event 0xa89968288be926b3b7a2a5a0d70c9e55178a465c62e84ad1c3f0401d8c2f7fc1.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle)
func (_InjectiveFutures *InjectiveFuturesFilterer) FilterMarketCreation(opts *bind.FilterOpts, marketID [][32]byte, ticker []string, oracle []common.Address) (*InjectiveFuturesMarketCreationIterator, error) {

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

	logs, sub, err := _InjectiveFutures.contract.FilterLogs(opts, "MarketCreation", marketIDRule, tickerRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &InjectiveFuturesMarketCreationIterator{contract: _InjectiveFutures.contract, event: "MarketCreation", logs: logs, sub: sub}, nil
}

// WatchMarketCreation is a free log subscription operation binding the contract event 0xa89968288be926b3b7a2a5a0d70c9e55178a465c62e84ad1c3f0401d8c2f7fc1.
//
// Solidity: event MarketCreation(bytes32 indexed marketID, string indexed ticker, address indexed oracle)
func (_InjectiveFutures *InjectiveFuturesFilterer) WatchMarketCreation(opts *bind.WatchOpts, sink chan<- *InjectiveFuturesMarketCreation, marketID [][32]byte, ticker []string, oracle []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _InjectiveFutures.contract.WatchLogs(opts, "MarketCreation", marketIDRule, tickerRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InjectiveFuturesMarketCreation)
				if err := _InjectiveFutures.contract.UnpackLog(event, "MarketCreation", log); err != nil {
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
func (_InjectiveFutures *InjectiveFuturesFilterer) ParseMarketCreation(log types.Log) (*InjectiveFuturesMarketCreation, error) {
	event := new(InjectiveFuturesMarketCreation)
	if err := _InjectiveFutures.contract.UnpackLog(event, "MarketCreation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InjectiveFuturesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InjectiveFutures contract.
type InjectiveFuturesOwnershipTransferredIterator struct {
	Event *InjectiveFuturesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InjectiveFuturesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InjectiveFuturesOwnershipTransferred)
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
		it.Event = new(InjectiveFuturesOwnershipTransferred)
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
func (it *InjectiveFuturesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InjectiveFuturesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InjectiveFuturesOwnershipTransferred represents a OwnershipTransferred event raised by the InjectiveFutures contract.
type InjectiveFuturesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InjectiveFutures *InjectiveFuturesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InjectiveFuturesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InjectiveFutures.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InjectiveFuturesOwnershipTransferredIterator{contract: _InjectiveFutures.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InjectiveFutures *InjectiveFuturesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InjectiveFuturesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InjectiveFutures.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InjectiveFuturesOwnershipTransferred)
				if err := _InjectiveFutures.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InjectiveFutures *InjectiveFuturesFilterer) ParseOwnershipTransferred(log types.Log) (*InjectiveFuturesOwnershipTransferred, error) {
	event := new(InjectiveFuturesOwnershipTransferred)
	if err := _InjectiveFutures.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InjectiveFuturesSignatureValidatorApprovalIterator is returned from FilterSignatureValidatorApproval and is used to iterate over the raw logs and unpacked data for SignatureValidatorApproval events raised by the InjectiveFutures contract.
type InjectiveFuturesSignatureValidatorApprovalIterator struct {
	Event *InjectiveFuturesSignatureValidatorApproval // Event containing the contract specifics and raw log

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
func (it *InjectiveFuturesSignatureValidatorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InjectiveFuturesSignatureValidatorApproval)
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
		it.Event = new(InjectiveFuturesSignatureValidatorApproval)
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
func (it *InjectiveFuturesSignatureValidatorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InjectiveFuturesSignatureValidatorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InjectiveFuturesSignatureValidatorApproval represents a SignatureValidatorApproval event raised by the InjectiveFutures contract.
type InjectiveFuturesSignatureValidatorApproval struct {
	SignerAddress    common.Address
	ValidatorAddress common.Address
	IsApproved       bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSignatureValidatorApproval is a free log retrieval operation binding the contract event 0xa8656e308026eeabce8f0bc18048433252318ab80ac79da0b3d3d8697dfba891.
//
// Solidity: event SignatureValidatorApproval(address indexed signerAddress, address indexed validatorAddress, bool isApproved)
func (_InjectiveFutures *InjectiveFuturesFilterer) FilterSignatureValidatorApproval(opts *bind.FilterOpts, signerAddress []common.Address, validatorAddress []common.Address) (*InjectiveFuturesSignatureValidatorApprovalIterator, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _InjectiveFutures.contract.FilterLogs(opts, "SignatureValidatorApproval", signerAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &InjectiveFuturesSignatureValidatorApprovalIterator{contract: _InjectiveFutures.contract, event: "SignatureValidatorApproval", logs: logs, sub: sub}, nil
}

// WatchSignatureValidatorApproval is a free log subscription operation binding the contract event 0xa8656e308026eeabce8f0bc18048433252318ab80ac79da0b3d3d8697dfba891.
//
// Solidity: event SignatureValidatorApproval(address indexed signerAddress, address indexed validatorAddress, bool isApproved)
func (_InjectiveFutures *InjectiveFuturesFilterer) WatchSignatureValidatorApproval(opts *bind.WatchOpts, sink chan<- *InjectiveFuturesSignatureValidatorApproval, signerAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _InjectiveFutures.contract.WatchLogs(opts, "SignatureValidatorApproval", signerAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InjectiveFuturesSignatureValidatorApproval)
				if err := _InjectiveFutures.contract.UnpackLog(event, "SignatureValidatorApproval", log); err != nil {
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
func (_InjectiveFutures *InjectiveFuturesFilterer) ParseSignatureValidatorApproval(log types.Log) (*InjectiveFuturesSignatureValidatorApproval, error) {
	event := new(InjectiveFuturesSignatureValidatorApproval)
	if err := _InjectiveFutures.contract.UnpackLog(event, "SignatureValidatorApproval", log); err != nil {
		return nil, err
	}
	return event, nil
}
