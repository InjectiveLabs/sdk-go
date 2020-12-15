package types

import (
	"errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
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
		return "statusinvalidtriggerprice"
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
	case "statusinvalidtriggerprice":
		return StatusInvalidTriggerPrice
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
	hex := common.ToHex(h)
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
	return common.ToHex([]byte(h))
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

func (p TradePair) ComputeHash() (common.Hash, error) {
	if p.Hash != "" {
		return common.HexToHash(p.Hash), nil
	}

	if len(p.MakerAssetData) == 0 {
		return common.Hash{}, errors.New("hash error: no maker asset data specified")
	} else if len(p.TakerAssetData) == 0 {
		return common.Hash{}, errors.New("hash error: no taker asset data specified")
	}

	var hash common.Hash

	if strings.Compare(p.MakerAssetData, p.TakerAssetData) < 0 {
		hash = common.BytesToHash(keccak256([]byte(p.MakerAssetData), []byte(p.TakerAssetData)))
	} else {
		hash = common.BytesToHash(keccak256([]byte(p.TakerAssetData), []byte(p.MakerAssetData)))
	}

	return hash, nil
}

func ComputeMarketHash(exchangeAddress common.Address, marketID common.Hash) common.Hash {
	hash := common.BytesToHash(keccak256(
		marketID.Bytes(),
		exchangeAddress.Bytes(),
	))
	return hash
}

func (p DerivativeMarket) Hash() (common.Hash, error) {
	if len(p.GetMarketId()) == 0 {
		return common.Hash{}, errors.New("hash error: no MarketId specified")
	} else if len(p.GetExchangeAddress()) == 0 {
		return common.Hash{}, errors.New("hash error: no ExchangeAddress specified")
	}

	var hash common.Hash

	marketIdBytes := common.FromHex(p.MarketId)
	if len(marketIdBytes) > common.HashLength {
		marketIdBytes = marketIdBytes[:common.HashLength]
	}

	hash = common.BytesToHash(keccak256(
		marketIdBytes,
		common.HexToAddress(p.ExchangeAddress).Bytes(),
	))

	return hash, nil
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
