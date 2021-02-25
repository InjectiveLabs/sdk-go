package types

import (
	"bytes"
	"errors"
	"math/big"
	"time"

	chainsdk "github.com/InjectiveLabs/sdk-go"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	log "github.com/xlab/suplog"
)

// SafeSignedOrder is a special signed order structure
// for including in Msgs, because it consists of primitive types.
// Avoid using raw *big.Int in Msgs.
type SafeSignedOrder struct {
	// ChainID is a network identifier of the order.
	ChainID int64 `json:"chainID,omitempty"`
	// Exchange v3 contract address.
	ExchangeAddress Address `json:"exchangeAddress,omitempty"`
	// Address that created the order.
	MakerAddress Address `json:"makerAddress,omitempty"`
	// Address that is allowed to fill the order. If set to "0x0", any address is
	// allowed to fill the order.
	TakerAddress Address `json:"takerAddress,omitempty"`
	// Address that will receive fees when order is filled.
	FeeRecipientAddress Address `json:"feeRecipientAddress,omitempty"`
	// Address that is allowed to call Exchange contract methods that affect this
	// order. If set to "0x0", any address is allowed to call these methods.
	SenderAddress Address `json:"senderAddress,omitempty"`
	// Amount of makerAsset being offered by maker. Must be greater than 0.
	MakerAssetAmount BigNum `json:"makerAssetAmount,omitempty"`
	// Amount of takerAsset being bid on by maker. Must be greater than 0.
	TakerAssetAmount BigNum `json:"takerAssetAmount,omitempty"`
	// Amount of Fee Asset paid to feeRecipientAddress by maker when order is filled. If set to
	// 0, no transfer of Fee Asset from maker to feeRecipientAddress will be attempted.
	MakerFee BigNum `json:"makerFee,omitempty"`
	// Amount of Fee Asset paid to feeRecipientAddress by taker when order is filled. If set to
	// 0, no transfer of Fee Asset from taker to feeRecipientAddress will be attempted.
	TakerFee BigNum `json:"takerFee,omitempty"`
	// Timestamp in seconds at which order expires.
	ExpirationTimeSeconds BigNum `json:"expirationTimeSeconds,omitempty"`
	// Arbitrary number to facilitate uniqueness of the order's hash.
	Salt BigNum `json:"salt,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerAsset.
	MakerAssetData HexBytes `json:"makerAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerAsset.
	TakerAssetData HexBytes `json:"takerAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerFee.
	MakerFeeAssetData HexBytes `json:"makerFeeAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerFee.
	TakerFeeAssetData HexBytes `json:"takerFeeAssetData,omitempty"`
	// Order signature.
	Signature HexBytes `json:"signature,omitempty"`
}

// PermyriadBase The scaling factor for Permyriads
var PermyriadBase = BigNum("10000").Int()

// NewSafeSignedOrder constructs a new SafeSignedOrder from given zeroex.SignedOrder.
func NewSafeSignedOrder(o *chainsdk.SignedOrder) *SafeSignedOrder {
	return zo2so(o)
}

// ToSignedOrder returns an appropriate zeroex.SignedOrder defined by SafeSignedOrder.
func (m *BaseOrder) ToSignedOrder() *chainsdk.SignedOrder {
	o, err := so2zo(m)
	if err != nil {
		panic(err)
	}
	return o
}

func (order *Order) IsReduceOnly() bool {
	return BigNum(order.GetOrder().GetMakerFee()).Int().Cmp(big.NewInt(0)) == 0
}

func (order *Order) DoesValidationPass(
	isLong bool,
	market *DerivativeMarket,
	currBlockTime time.Time,
) error {
	err := order.ComputeAndSetOrderType()
	if err != nil {
		log.Infoln("fail")
		return err
	}

	isOrderExpired := order.Order.IsExpired(currBlockTime)
	if isOrderExpired {
		return sdkerrors.Wrapf(ErrOrderExpired, "order expiration %s <= block time %d", order.GetOrder().GetExpirationTimeSeconds(), currBlockTime.Unix())
	}
	if order.OrderType == 0 {
		margin := BigNum(order.Order.GetMakerFee()).Int()
		contractPriceMarginRequirement := order.ComputeContractPriceMarginRequirement(market)
		if margin.Cmp(contractPriceMarginRequirement) < 0 {
			return sdkerrors.Wrapf(ErrOverLeveragedOrder, "margin %s < contractPriceMarginRequirement %s", margin.String(), contractPriceMarginRequirement.String())
		}

		indexPriceMarginRequirement := order.ComputeIndexPriceMarginRequirement(isLong, market)
		indexPrice := BigNum(market.GetIndexPrice()).Int()

		if isLong && indexPrice.Cmp(indexPriceMarginRequirement) < 0 {
			return sdkerrors.Wrapf(ErrOverLeveragedOrder, "indexPrice %s <= indexPriceReq %s", market.GetIndexPrice(), order.IndexPriceRequirement)
		} else if !isLong && indexPrice.Cmp(indexPriceMarginRequirement) > 0 {
			return sdkerrors.Wrapf(ErrOverLeveragedOrder, "indexPrice %s >= indexPriceReq %s", market.GetIndexPrice(), order.IndexPriceRequirement)
		}
	}

	return nil
}

func (order *Order) ComputeAndSetOrderType() error {
	orderTypeNumber := new(big.Int).SetBytes(common.FromHex(order.GetOrder().GetMakerFeeAssetData())[:common.HashLength]).Uint64()
	if orderTypeNumber == 0 || orderTypeNumber == 5 {
		order.OrderType = orderTypeNumber
	} else {
		return sdkerrors.Wrapf(ErrUnrecognizedOrderType, "Cannot recognize MakerFeeAssetData of %s", order.GetOrder().GetMakerFeeAssetData())
	}
	return nil
}

func (order *Order) ComputeIndexPriceMarginRequirement(isLong bool, market *DerivativeMarket) *big.Int {
	price := BigNum(order.Order.GetMakerAssetAmount()).Int()
	quantity := BigNum(order.Order.GetTakerAssetAmount()).Int()
	margin := BigNum(order.Order.GetMakerFee()).Int()
	pq := new(big.Int).Mul(price, quantity)
	alphaQuantity := ScalePermyriad(quantity, BigNum(market.InitialMarginRatio).Int())
	num := new(big.Int)
	denom := new(big.Int)

	if isLong {
		num = num.Sub(margin, pq)
		denom = denom.Sub(alphaQuantity, quantity)
	} else {
		num = num.Add(margin, pq)
		denom = denom.Add(alphaQuantity, quantity)
	}

	indexPriceReq := new(big.Int).Div(num, denom)
	order.IndexPriceRequirement = indexPriceReq.String()
	return indexPriceReq
}

// quantity * initialMarginRatio * price
func (order *Order) ComputeContractPriceMarginRequirement(market *DerivativeMarket) *big.Int {
	price := BigNum(order.Order.GetMakerAssetAmount()).Int()
	quantity := BigNum(order.Order.GetTakerAssetAmount()).Int()
	alphaQuantity := ScalePermyriad(quantity, BigNum(market.InitialMarginRatio).Int())
	return new(big.Int).Mul(alphaQuantity, price)
}

// orderMarginHold = (1 + txFeePermyriad / 10000) * assetAmount
func GetAssetAmountWithFeesApplied(assetAmount, txFeePermyriad *big.Int) (orderMarginHold *big.Int) {
	return IncrementByScaledPermyriad(assetAmount, txFeePermyriad)
}

// orderMarginHold = (1 + txFeePermyriad / 10000) * margin * (remainingQuantity) / order.quantity
func (o *BaseOrder) ComputeDerivativeOrderMarginHold(remainingQuantity, txFeePermyriad *big.Int) (orderMarginHold *big.Int) {
	margin := BigNum(o.GetMakerFee()).Int()
	scaledMargin := IncrementByScaledPermyriad(margin, txFeePermyriad)
	originalQuantity := BigNum(o.GetTakerAssetAmount()).Int()

	// TODO: filledAmount should always be zero with TEC since there will be no UnknownOrderHash
	numerator := new(big.Int).Mul(scaledMargin, remainingQuantity)

	// originalQuantity should never be zero, however
	if originalQuantity.Sign() == 0 {
		return scaledMargin
	}

	orderMarginHold = new(big.Int).Div(numerator, originalQuantity)
	return orderMarginHold
}

func (o *BaseOrder) IsExpired(currBlockTime time.Time) bool {
	blockTime := big.NewInt(currBlockTime.Unix())
	orderExpirationTime := BigNum(o.GetExpirationTimeSeconds()).Int()

	if orderExpirationTime.Cmp(blockTime) <= 0 {
		return true
	}
	return false
}

// return amount * (1 + permyriad/10000) = (amount + amount * permyriad/10000)
func IncrementByScaledPermyriad(amount, permyriad *big.Int) *big.Int {
	return new(big.Int).Add(amount, ScalePermyriad(amount, permyriad))
}

// return (amount * permyriad) / 10000
func ScalePermyriad(amount, permyriad *big.Int) *big.Int {
	scaleFactor := new(big.Int).Mul(amount, permyriad)
	return new(big.Int).Div(scaleFactor, PermyriadBase)
}

func ComputeSubaccountID(address string, takerFee string) common.Hash {
	return common.BytesToHash(append(common.HexToAddress(address).Bytes(), common.LeftPadBytes(BigNum(takerFee).Int().Bytes(), 12)...))
}

// GetDirectionMarketAndSubaccountID
func (o *BaseOrder) GetDirectionMarketAndSubaccountID(shouldGetMakerSubaccount bool) (isLong bool, marketID common.Hash, subaccountID common.Hash) {
	mData, tData := common.FromHex(o.GetMakerAssetData()), common.FromHex(o.GetTakerAssetData())

	if len(mData) > common.HashLength {
		mData = mData[:common.HashLength]
	}

	if len(tData) > common.HashLength {
		tData = tData[:common.HashLength]
	}

	if bytes.Equal(tData, common.Hash{}.Bytes()) {
		isLong = true
		marketID = common.BytesToHash(mData)
	} else {
		isLong = false
		marketID = common.BytesToHash(tData)
	}

	var address string

	if shouldGetMakerSubaccount {
		address = o.GetMakerAddress()
	} else {
		address = o.GetTakerAddress()
	}

	subaccountID = ComputeSubaccountID(address, o.GetTakerFee())

	return isLong, marketID, subaccountID
}

// zo2so internal function converts model from *zeroex.SignedOrder to *SafeSignedOrder.
func zo2so(o *chainsdk.SignedOrder) *SafeSignedOrder {
	if o == nil {
		return nil
	}
	return &SafeSignedOrder{
		ChainID:               o.ChainID.Int64(),
		ExchangeAddress:       Address{o.ExchangeAddress},
		MakerAddress:          Address{o.MakerAddress},
		TakerAddress:          Address{o.TakerAddress},
		FeeRecipientAddress:   Address{o.FeeRecipientAddress},
		SenderAddress:         Address{o.SenderAddress},
		MakerAssetAmount:      BigNum(o.MakerAssetAmount.String()),
		TakerAssetAmount:      BigNum(o.TakerAssetAmount.String()),
		MakerFee:              BigNum(o.MakerFee.String()),
		TakerFee:              BigNum(o.TakerFee.String()),
		ExpirationTimeSeconds: BigNum(o.ExpirationTimeSeconds.String()),
		Salt:                  BigNum(o.Salt.String()),
		MakerAssetData:        o.MakerAssetData,
		TakerAssetData:        o.TakerAssetData,
		MakerFeeAssetData:     o.MakerFeeAssetData,
		TakerFeeAssetData:     o.TakerFeeAssetData,
		Signature:             o.Signature,
	}
}

// so2zo internal function converts model from *SafeSignedOrder to *zeroex.SignedOrder.
func so2zo(o *BaseOrder) (*chainsdk.SignedOrder, error) {
	if o == nil {
		return nil, nil
	}
	order := chainsdk.Order{
		ChainID:             big.NewInt(o.ChainId),
		ExchangeAddress:     common.HexToAddress(o.ExchangeAddress),
		MakerAddress:        common.HexToAddress(o.MakerAddress),
		TakerAddress:        common.HexToAddress(o.TakerAddress),
		SenderAddress:       common.HexToAddress(o.SenderAddress),
		FeeRecipientAddress: common.HexToAddress(o.FeeRecipientAddress),
		MakerAssetData:      common.FromHex(o.MakerAssetData),
		MakerFeeAssetData:   common.FromHex(o.MakerFeeAssetData),
		TakerAssetData:      common.FromHex(o.TakerAssetData),
		TakerFeeAssetData:   common.FromHex(o.TakerFeeAssetData),
	}

	if v, ok := math.ParseBig256(string(o.MakerAssetAmount)); !ok {
		return nil, errors.New("makerAssetAmount parse failed")
	} else {
		order.MakerAssetAmount = v
	}
	if v, ok := math.ParseBig256(string(o.MakerFee)); !ok {
		return nil, errors.New("makerFee parse failed")
	} else {
		order.MakerFee = v
	}
	if v, ok := math.ParseBig256(string(o.TakerAssetAmount)); !ok {
		return nil, errors.New("takerAssetAmount parse failed")
	} else {
		order.TakerAssetAmount = v
	}
	if v, ok := math.ParseBig256(string(o.TakerFee)); !ok {
		return nil, errors.New("takerFee parse failed")
	} else {
		order.TakerFee = v
	}
	if v, ok := math.ParseBig256(string(o.ExpirationTimeSeconds)); !ok {
		return nil, errors.New("expirationTimeSeconds parse failed")
	} else {
		order.ExpirationTimeSeconds = v
	}
	if v, ok := math.ParseBig256(string(o.Salt)); !ok {
		return nil, errors.New("salt parse failed")
	} else {
		order.Salt = v
	}
	signedOrder := &chainsdk.SignedOrder{
		Order:     order,
		Signature: common.FromHex(o.Signature),
	}
	return signedOrder, nil
}
