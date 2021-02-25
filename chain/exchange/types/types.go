package types

import (
	"errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

const (
	ADDRESS_LENGTH = 42
	BYTES32_LENGTH = 66
)

type Direction uint8

const (
	Long  Direction = 0
	Short Direction = 1
)

func DirectionFromString(direction string) Direction {
	switch direction {
	case "long":
		return Long
	case "short":
		return Short
	default:
		return Long
	}

}

func (d Direction) String() string {
	switch d {
	case Long:
		return "long"
	case Short:
		return "short"
	}
	return ""
}

// Order Type denotes an order type
type OrderType uint64

const (
	VANILLA_LIMIT     OrderType = 0
	STOP_LIMIT_PROFIT OrderType = 1
	STOP_LIMIT_LOSS   OrderType = 2
	STOP_LOSS         OrderType = 3
	TAKE_PROFIT       OrderType = 4
	REDUCE_ONLY       OrderType = 5
)

func ComputePositionBankruptcyPrice(
	entryPrice, margin, quantity, marketCumulativeFunding, positionCumFundingEntry *big.Int,
	isDirectionLong bool,
) *big.Int {
	// funding fee = position quantity * (market cumulative funding - position cumulative funding entry)
	fundingFee := new(big.Int).Mul(quantity, new(big.Int).Sub(marketCumulativeFunding, positionCumFundingEntry))

	isFundingPositive := fundingFee.Cmp(big.NewInt(0)) > 0

	// When the funding rate is positive, longs pay shorts. When negative, shorts pay longs.
	longPositionGetsPaid := isDirectionLong && !isFundingPositive
	longPositionPays := isDirectionLong && isFundingPositive

	shortPositionGetsPaid := !isDirectionLong && isFundingPositive
	shortPositionPays := !isDirectionLong && !isFundingPositive

	netMargin := BigNum(margin.String()).Int()

	if longPositionGetsPaid || shortPositionGetsPaid {
		netMargin = new(big.Int).Add(margin, fundingFee)
	} else if longPositionPays || shortPositionPays {
		netMargin = new(big.Int).Sub(margin, fundingFee)
	}

	var bankruptcyPrice *big.Int
	unitMargin := new(big.Int).Div(netMargin, quantity)
	if isDirectionLong {
		bankruptcyPrice = new(big.Int).Sub(entryPrice, unitMargin)
	} else {
		bankruptcyPrice = new(big.Int).Add(entryPrice, unitMargin)
	}
	return bankruptcyPrice
}

// OrderStatus encodes order status according to LibOrder.OrderStatus
type OrderStatus uint8

const (
	// StatusInvalid is the default value
	StatusInvalid OrderStatus = 0
	// StatusInsufficientMarginForContractPrice when Order does not have enough margin for contract price
	StatusInsufficientMarginForContractPrice OrderStatus = 1
	// StatusInsufficientMarginIndexPrice when Order does not have enough margin for index price
	StatusInsufficientMarginIndexPrice OrderStatus = 2
	// StatusFillable when order is fillable
	StatusFillable OrderStatus = 3
	// StatusExpired when order has already expired
	StatusExpired OrderStatus = 4
	// StatusFullyFilled when order is fully filled
	StatusFullyFilled OrderStatus = 5
	// StatusCancelled when order has been cancelled
	StatusCancelled OrderStatus = 6
	// Maker of the order does not have sufficient funds deposited to be filled.
	StatusUnfunded OrderStatus = 7
	// Index Price has not been triggered
	StatusUntriggered OrderStatus = 8
	// StatusInvalidTriggerPrice TakeProfit trigger price is lower than contract price or StopLoss trigger price is higher than contract price
	StatusInvalidTriggerPrice OrderStatus = 9
	// StatusReduceOnlyExpired when order has expired due to a reduce only condition invalidation
	StatusReduceOnlyExpired OrderStatus = 10
	// StatusSoftCancelled when order has been soft-cancelled
	StatusSoftCancelled OrderStatus = 11
)

func (o OrderStatus) String() string {
	switch o {
	case StatusInvalid:
		return "invalid"
	case StatusInsufficientMarginForContractPrice:
		return "insufficientMarginForContractPrice"
	case StatusInsufficientMarginIndexPrice:
		return "insufficientMarginIndexPrice"
	case StatusFillable:
		return "fillable"
	case StatusExpired:
		return "expired"
	case StatusFullyFilled:
		return "fullyFilled"
	case StatusCancelled:
		return "cancelled"
	case StatusUnfunded:
		return "unfunded"
	case StatusUntriggered:
		return "untriggered"
	case StatusInvalidTriggerPrice:
		return "invalidTriggerPrice"
	case StatusReduceOnlyExpired:
		return "reduceOnlyExpired"
	case StatusSoftCancelled:
		return "softCancelled"
	default:
		return ""
	}
}

func OrderStatusFromString(status string) OrderStatus {
	switch status {
	case "invalid":
		return StatusInvalid
	case "insufficientMarginForContractPrice":
		return StatusInsufficientMarginForContractPrice
	case "insufficientMarginIndexPrice":
		return StatusInsufficientMarginIndexPrice
	case "fillable":
		return StatusFillable
	case "expired":
		return StatusExpired
	case "fullyFilled":
		return StatusFullyFilled
	case "cancelled":
		return StatusCancelled
	case "unfunded":
		return StatusUnfunded
	case "untriggered":
		return StatusUntriggered
	case "invalidtriggerprice":
		return StatusInvalidTriggerPrice
	case "reduceOnlyExpired":
		return StatusReduceOnlyExpired
	case "softCancelled":
		return StatusSoftCancelled
	default:
		return StatusInvalid
	}
}

//type OrderSoftCancelRequest struct {
//	TxHash             common.Hash `json:"txHash"`
//	OrderHash          common.Hash `json:"orderHash"`
//	ApprovalSignatures [][]byte    `json:"approvalSignatures"`
//}

type Hash struct {
	common.Hash
}

func (h Hash) MarshalJSON() ([]byte, error) {
	hex := h.Hash.Hex()
	buf := make([]byte, 0, len(hex)+2)
	buf = append(buf, '"')
	buf = append(buf, hex...)
	buf = append(buf, '"')
	return buf, nil
}

type HexBytes []byte

func (h HexBytes) MarshalJSON() ([]byte, error) {
	hex := common.Bytes2Hex(h)
	buf := make([]byte, 0, len(hex)+2)
	buf = append(buf, '"')
	buf = append(buf, hex...)
	buf = append(buf, '"')
	return buf, nil
}

func (h *HexBytes) UnmarshalJSON(src []byte) error {
	if len(src) == 2 {
		return nil
	} else if len(src) < 2 {
		return errors.New("failed to parse: " + string(src))
	}

	*h = HexBytes(common.FromHex(string(src[1 : len(src)-1])))
	return nil
}

func (h HexBytes) String() string {
	return common.Bytes2Hex([]byte(h))
}

type Address struct {
	common.Address
}

func (a Address) MarshalJSON() ([]byte, error) {
	hex := a.Address.Hex()
	buf := make([]byte, 0, len(hex)+2)
	buf = append(buf, '"')
	buf = append(buf, hex...)
	buf = append(buf, '"')
	return buf, nil
}

const nullAddressHex = "0x0000000000000000000000000000000000000000"

func (a Address) IsEmpty() bool {
	if a.Hex() == nullAddressHex {
		return true
	}

	return false
}

type BigNum string

func (n BigNum) Int() *big.Int {
	i := new(big.Int)
	i.SetString(string(n), 10)
	return i
}

func NewBigNum(i *big.Int) BigNum {
	if i == nil {
		return "0"
	}
	return BigNum(i.String())
}

func (m MarginInfo) IsMarginHoldBreached() (availableMargin *big.Int, isBreached bool) {
	availableMargin = new(big.Int).Sub(BigNum(m.GetTotalDeposits()).Int(), BigNum(m.GetMarginHold()).Int())
	if availableMargin.Cmp(big.NewInt(0)) < 0 {
		return availableMargin, true
	}
	return availableMargin, false
}


func GetMarketIdFromAssetPair(baseAsset common.Address, quoteAsset common.Address, exchangeAddress common.Address) common.Hash {
	return common.BytesToHash(keccak256(baseAsset.Bytes(), quoteAsset.Bytes(), exchangeAddress.Bytes()))
}

func (bo *BaseOrder) GetMakerTakerAssets() (makerAsset common.Address, takerAsset common.Address) {
	makerAsset = common.BytesToAddress(common.FromHex(bo.GetMakerAssetData())[:common.AddressLength])
	takerAsset = common.BytesToAddress(common.FromHex(bo.GetTakerAssetData())[:common.AddressLength])
	return
}

func (m *DerivativeMarket) CheckExpiration(currBlockTime time.Time) error {
	nextFundingTimestamp, fundingInterval := BigNum(m.GetNextFundingTimestamp()).Int(), BigNum(m.GetFundingInterval()).Int()
	if fundingInterval.Cmp(big.NewInt(0)) == 0 {
		// expiration time must be greater than current block time
		if nextFundingTimestamp.Cmp(big.NewInt(currBlockTime.Unix())) <= 0 {
			return sdkerrors.Wrap(ErrMarketExpired, m.GetTicker())
		}
	}
	return nil
}

// keccak256 calculates and returns the Keccak256 hash of the input data.
func keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		_, _ = d.Write(b)
	}
	return d.Sum(nil)
}