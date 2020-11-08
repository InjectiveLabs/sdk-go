package ordervalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"

	sdk "github.com/InjectiveLabs/sdk-go"
)

const (
	// MaxOrderSizeInBytes is the maximum number of bytes allowed for encoded
	// orders. It allows for MultiAssetProxy orders with roughly 45 total ERC20
	// assets or roughly 36 total ERC721 assets (combined between both maker and
	// taker; depends on the other fields of the order).
	MaxOrderSizeInBytes = 16000
)

// RejectedOrderInfo encapsulates all the needed information to understand _why_ a 0x order
// was rejected (i.e. did not pass) order validation. Since there are many potential reasons, some
// Mesh-specific, others 0x-specific and others due to external factors (i.e., network
// disruptions, etc...), we categorize them into `Kind`s and uniquely identify the reasons for
// machines with a `Code`
type RejectedOrderInfo struct {
	OrderHash   common.Hash         `json:"orderHash"`
	SignedOrder *sdk.SignedOrder    `json:"signedOrder"`
	Kind        RejectedOrderKind   `json:"kind"`
	Status      RejectedOrderStatus `json:"status"`
}

// AcceptedOrderInfo represents an fillable order and how much it could be filled for
type AcceptedOrderInfo struct {
	OrderHash                common.Hash      `json:"orderHash"`
	SignedOrder              *sdk.SignedOrder `json:"signedOrder"`
	FillableTakerAssetAmount *big.Int         `json:"fillableTakerAssetAmount"`
	IsNew                    bool             `json:"isNew"`
}

type acceptedOrderInfoJSON struct {
	OrderHash                string           `json:"orderHash"`
	SignedOrder              *sdk.SignedOrder `json:"signedOrder"`
	FillableTakerAssetAmount string           `json:"fillableTakerAssetAmount"`
	IsNew                    bool             `json:"isNew"`
}

// MarshalJSON is a custom Marshaler for AcceptedOrderInfo
func (a AcceptedOrderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"orderHash":                a.OrderHash.Hex(),
		"signedOrder":              a.SignedOrder,
		"fillableTakerAssetAmount": a.FillableTakerAssetAmount.String(),
		"isNew":                    a.IsNew,
	})
}

// UnmarshalJSON implements a custom JSON unmarshaller for the OrderEvent type
func (a *AcceptedOrderInfo) UnmarshalJSON(data []byte) error {
	var acceptedOrderInfoJSON acceptedOrderInfoJSON
	err := json.Unmarshal(data, &acceptedOrderInfoJSON)
	if err != nil {
		return err
	}

	a.OrderHash = common.HexToHash(acceptedOrderInfoJSON.OrderHash)
	a.SignedOrder = acceptedOrderInfoJSON.SignedOrder
	a.IsNew = acceptedOrderInfoJSON.IsNew
	var ok bool
	a.FillableTakerAssetAmount, ok = math.ParseBig256(acceptedOrderInfoJSON.FillableTakerAssetAmount)
	if !ok {
		return errors.New("Invalid uint256 number encountered for FillableTakerAssetAmount")
	}
	return nil
}

// RejectedOrderStatus enumerates all the unique reasons for an orders rejection
type RejectedOrderStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RejectedOrderStatus values
var (
	ROEthRPCRequestFailed = RejectedOrderStatus{
		Code:    "EthRPCRequestFailed",
		Message: "network request to Ethereum RPC endpoint failed",
	}
	ROCoordinatorRequestFailed = RejectedOrderStatus{
		Code:    "CoordinatorRequestFailed",
		Message: "network request to coordinator server endpoint failed",
	}
	ROCoordinatorSoftCancelled = RejectedOrderStatus{
		Code:    "CoordinatorSoftCancelled",
		Message: "order was soft-cancelled via the coordinator server",
	}
	ROCoordinatorEndpointNotFound = RejectedOrderStatus{
		Code:    "CoordinatorEndpointNotFound",
		Message: "corresponding coordinator endpoint not found in CoordinatorRegistry contract",
	}
	ROInvalidMakerAssetAmount = RejectedOrderStatus{
		Code:    "OrderHasInvalidMakerAssetAmount",
		Message: "order makerAssetAmount cannot be 0",
	}
	ROInvalidTakerAssetAmount = RejectedOrderStatus{
		Code:    "OrderHasInvalidTakerAssetAmount",
		Message: "order takerAssetAmount cannot be 0",
	}
	ROExpired = RejectedOrderStatus{
		Code:    "OrderExpired",
		Message: "order expired according to latest block timestamp",
	}
	ROFullyFilled = RejectedOrderStatus{
		Code:    "OrderFullyFilled",
		Message: "order already fully filled",
	}
	ROCancelled = RejectedOrderStatus{
		Code:    "OrderCancelled",
		Message: "order cancelled",
	}
	ROUnfunded = RejectedOrderStatus{
		Code:    "OrderUnfunded",
		Message: "maker has insufficient balance or allowance for this order to be filled",
	}
	ROInvalidMakerAssetData = RejectedOrderStatus{
		Code:    "OrderHasInvalidMakerAssetData",
		Message: "order makerAssetData must encode a supported assetData type",
	}
	ROInvalidMakerFeeAssetData = RejectedOrderStatus{
		Code:    "OrderHasInvalidMakerFeeAssetData",
		Message: "order makerFeeAssetData must encode a supported assetData type",
	}
	ROInvalidTakerAssetData = RejectedOrderStatus{
		Code:    "OrderHasInvalidTakerAssetData",
		Message: "order takerAssetData must encode a supported assetData type",
	}
	ROInvalidTakerFeeAssetData = RejectedOrderStatus{
		Code:    "OrderHasInvalidTakerFeeAssetData",
		Message: "order takerFeeAssetData must encode a supported assetData type",
	}
	ROInvalidSignature = RejectedOrderStatus{
		Code:    "OrderHasInvalidSignature",
		Message: "order signature must be valid",
	}
	ROMaxExpirationExceeded = RejectedOrderStatus{
		Code:    "OrderMaxExpirationExceeded",
		Message: "order expiration too far in the future",
	}
	ROInternalError = RejectedOrderStatus{
		Code:    "InternalError",
		Message: "an unexpected internal error has occurred",
	}
	ROMaxOrderSizeExceeded = RejectedOrderStatus{
		Code:    "MaxOrderSizeExceeded",
		Message: fmt.Sprintf("order exceeds the maximum encoded size of %d bytes", MaxOrderSizeInBytes),
	}
	ROOrderAlreadyStoredAndUnfillable = RejectedOrderStatus{
		Code:    "OrderAlreadyStoredAndUnfillable",
		Message: "order is already stored and is unfillable. Mesh keeps unfillable orders in storage for a little while incase a block re-org makes them fillable again",
	}
	ROIncorrectChain = RejectedOrderStatus{
		Code:    "OrderForIncorrectChain",
		Message: "order was created for a different chain than the one this Mesh node is configured to support",
	}
	ROIncorrectExchangeAddress = RejectedOrderStatus{
		Code:    "IncorrectExchangeAddress",
		Message: "the exchange address for the order does not match the chain ID/network ID",
	}
	ROSenderAddressNotAllowed = RejectedOrderStatus{
		Code:    "SenderAddressNotAllowed",
		Message: "orders with a senderAddress are not currently supported",
	}
	RODatabaseFullOfOrders = RejectedOrderStatus{
		Code:    "DatabaseFullOfOrders",
		Message: "database is full of pinned orders and no orders can be deleted to make space (consider increasing MAX_ORDERS_IN_STORAGE)",
	}
)

// ROInvalidSchemaCode is the RejectedOrderStatus emitted if an order doesn't conform to the order schema
const ROInvalidSchemaCode = "InvalidSchema"

// RejectedOrderKind enumerates all kinds of reasons an order could be rejected by Mesh
type RejectedOrderKind string

// RejectedOrderKind values
const (
	ZeroExValidation = RejectedOrderKind("ZEROEX_VALIDATION")
	MeshError        = RejectedOrderKind("MESH_ERROR")
	MeshValidation   = RejectedOrderKind("MESH_VALIDATION")
	CoordinatorError = RejectedOrderKind("COORDINATOR_ERROR")
)

// ValidationResults defines the validation results returned from BatchValidate
// Within this context, an order is `Accepted` if it passes all the 0x schema tests
// and is fillable for a non-zero amount. An order is `Rejected` if it does not
// satisfy these conditions OR if we were unable to complete the validation process
// for whatever reason
type ValidationResults struct {
	Accepted []*AcceptedOrderInfo `json:"accepted"`
	Rejected []*RejectedOrderInfo `json:"rejected"`
}
