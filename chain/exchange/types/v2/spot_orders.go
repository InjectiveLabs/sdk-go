package v2

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

var _ types.IOrderInfo = &OrderInfo{}

func (m *SpotOrder) ToSpotMarketOrder(sender sdk.AccAddress, balanceHold math.LegacyDec, orderHash common.Hash) *SpotMarketOrder {
	if m.OrderInfo.FeeRecipient == "" {
		m.OrderInfo.FeeRecipient = sender.String()
	}
	return &SpotMarketOrder{
		OrderInfo:    m.OrderInfo,
		BalanceHold:  balanceHold,
		OrderHash:    orderHash.Bytes(),
		OrderType:    m.OrderType,
		TriggerPrice: m.TriggerPrice,
	}
}

func (m *SpotLimitOrder) ToStandardized() *TrimmedLimitOrder {
	return &TrimmedLimitOrder{
		Price:        m.OrderInfo.Price,
		Quantity:     m.OrderInfo.Quantity,
		OrderHash:    common.BytesToHash(m.OrderHash).Hex(),
		SubaccountId: m.OrderInfo.SubaccountId,
	}
}

func (m *SpotOrder) GetNewSpotLimitOrder(sender sdk.AccAddress, orderHash common.Hash) *SpotLimitOrder {
	if m.OrderInfo.FeeRecipient == "" {
		m.OrderInfo.FeeRecipient = sender.String()
	}
	return &SpotLimitOrder{
		OrderInfo:       m.OrderInfo,
		OrderType:       m.OrderType,
		Fillable:        m.OrderInfo.Quantity,
		TriggerPrice:    m.TriggerPrice,
		OrderHash:       orderHash.Bytes(),
		ExpirationBlock: m.ExpirationBlock,
	}
}

func (m *SpotOrder) SubaccountID() common.Hash {
	return m.OrderInfo.SubaccountID()
}

func (m *SpotOrder) IsFromDefaultSubaccount() bool {
	return m.OrderInfo.IsFromDefaultSubaccount()
}

func (m *SpotLimitOrder) IsFromDefaultSubaccount() bool {
	return m.OrderInfo.IsFromDefaultSubaccount()
}

func (m *SpotLimitOrder) Cid() string {
	return m.OrderInfo.GetCid()
}

func (m *SpotMarketOrder) Cid() string {
	return m.OrderInfo.GetCid()
}

func (m *SpotLimitOrder) SubaccountID() common.Hash {
	return m.OrderInfo.SubaccountID()
}

func (m *SpotMarketOrder) SubaccountID() common.Hash {
	return m.OrderInfo.SubaccountID()
}

func (m *SpotMarketOrder) IsFromDefaultSubaccount() bool {
	return m.OrderInfo.IsFromDefaultSubaccount()
}

func (m *SpotMarketOrder) SdkAccAddress() sdk.AccAddress {
	return m.SubaccountID().Bytes()[:common.AddressLength]
}

func (m *SpotLimitOrder) SdkAccAddress() sdk.AccAddress {
	return m.SubaccountID().Bytes()[:common.AddressLength]
}

func (m *SpotLimitOrder) FeeRecipient() common.Address {
	return m.OrderInfo.FeeRecipientAddress()
}

func (m *SpotMarketOrder) FeeRecipient() common.Address {
	return m.OrderInfo.FeeRecipientAddress()
}

func (m *SpotOrder) CheckTickSize(minPriceTickSize, minQuantityTickSize math.LegacyDec) error {
	if types.BreachesMinimumTickSize(m.OrderInfo.Price, minPriceTickSize) {
		return errors.Wrapf(
			types.ErrInvalidPrice,
			"price %s must be a multiple of the minimum price tick size %s",
			m.OrderInfo.Price.String(),
			minPriceTickSize.String(),
		)
	}
	if types.BreachesMinimumTickSize(m.OrderInfo.Quantity, minQuantityTickSize) {
		return errors.Wrapf(
			types.ErrInvalidQuantity,
			"quantity %s must be a multiple of the minimum quantity tick size %s",
			m.OrderInfo.Quantity.String(),
			minQuantityTickSize.String(),
		)
	}
	return nil
}

func (m *SpotOrder) CheckNotional(minNotional math.LegacyDec) error {
	orderNotional := m.GetQuantity().Mul(m.GetPrice())
	if !minNotional.IsNil() && orderNotional.LT(minNotional) {
		return errors.Wrapf(
			types.ErrInvalidNotional,
			"order notional (%s) is less than the minimum notional for the market (%s)",
			orderNotional.String(),
			minNotional.String(),
		)
	}
	return nil
}

func (m *SpotOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *SpotOrder) Cid() string {
	return m.OrderInfo.Cid
}

func (m *SpotLimitOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *SpotLimitOrder) Hash() common.Hash {
	return common.BytesToHash(m.OrderHash)
}

func (m *SpotMarketOrder) Hash() common.Hash {
	return common.BytesToHash(m.OrderHash)
}

func (m *SpotOrder) IsConditional() bool {
	return m.OrderType.IsConditional()
}

func (m *SpotLimitOrder) IsConditional() bool {
	return m.OrderType.IsConditional()
}

func (m *SpotMarketOrder) IsConditional() bool {
	return m.OrderType.IsConditional()
}

func (m *SpotLimitOrder) GetUnfilledNotional() math.LegacyDec {
	return m.Fillable.Mul(m.OrderInfo.Price)
}
func (m *SpotLimitOrder) GetUnfilledFeeAmount(fee math.LegacyDec) math.LegacyDec {
	return m.GetUnfilledNotional().Mul(fee)
}

func (m *SpotOrder) GetBalanceHoldAndMarginDenom(market *SpotMarket) (math.LegacyDec, string) {
	var denom string
	var balanceHold math.LegacyDec
	if m.IsBuy() {
		denom = market.QuoteDenom
		if m.OrderType.IsPostOnly() {
			// for a PO limit buy in the ETH/USDT market, denom is USDT and balanceHold is (1 + makerFee)*(price * quantity)
			balanceHold = m.OrderInfo.GetNotional()
			if market.MakerFeeRate.IsPositive() {
				balanceHold = balanceHold.Add(m.OrderInfo.GetFeeAmount(market.MakerFeeRate))
			}
		} else {
			// for a normal limit buy in the ETH/USDT market, denom is USDT and balanceHold is (1 + takerFee)*(price * quantity)
			balanceHold = m.OrderInfo.GetNotional().Add(m.OrderInfo.GetFeeAmount(market.TakerFeeRate))
		}
	} else {
		// for a limit sell in the ETH/USDT market, denom is ETH and balanceHold is just quantity
		denom = market.BaseDenom
		balanceHold = m.OrderInfo.Quantity
	}

	return balanceHold, denom
}

func (m *SpotLimitOrder) GetUnfilledMarginHoldAndMarginDenom(market *SpotMarket, isTransient bool) (math.LegacyDec, string) {
	var denom string
	var balanceHold math.LegacyDec
	if m.IsBuy() {
		var tradeFeeRate math.LegacyDec

		if isTransient {
			tradeFeeRate = market.TakerFeeRate
		} else {
			tradeFeeRate = math.LegacyMaxDec(math.LegacyZeroDec(), market.MakerFeeRate)
		}

		// for a resting limit buy in the ETH/USDT market, denom is USDT and fillable amount is BalanceHold is
		// (1 + makerFee)*(price * quantity) since (takerFee - makerFee) is already refunded
		denom = market.QuoteDenom
		balanceHold = m.GetUnfilledNotional().Add(m.GetUnfilledFeeAmount(tradeFeeRate))
	} else {
		// for a limit sell in the ETH/USDT market, denom is ETH and balanceHold is just quantity
		denom = market.BaseDenom
		balanceHold = m.Fillable
	}

	return balanceHold, denom
}

func (m *SpotOrder) GetMarginDenom(market *SpotMarket) string {
	var denom string
	if m.IsBuy() {
		// for a market buy in the ETH/USDT market, margin denom is USDT
		denom = market.QuoteDenom
	} else {
		// for a market buy in the ETH/USDT market, margin denom is ETH
		denom = market.BaseDenom
	}
	return denom
}

// GetMarketOrderBalanceHold calculates the balance hold for the market order.
func (m *SpotOrder) GetMarketOrderBalanceHold(feeRate, bestPrice math.LegacyDec) math.LegacyDec {
	var balanceHold math.LegacyDec

	if m.IsBuy() {
		// required margin for best sell price = bestPrice * quantity * (1 + feeRate)
		requiredMarginForBestPrice := bestPrice.Mul(m.OrderInfo.Quantity).Mul(math.LegacyOneDec().Add(feeRate))
		requiredMarginForWorstPrice := m.OrderInfo.Price.Mul(m.OrderInfo.Quantity).Mul(math.LegacyOneDec().Add(feeRate))
		requiredMargin := math.LegacyMaxDec(requiredMarginForBestPrice, requiredMarginForWorstPrice)
		balanceHold = requiredMargin
	} else {
		// required margin for market sells just equals the quantity being sold
		balanceHold = m.OrderInfo.Quantity
	}
	return balanceHold
}

func (m *SpotLimitOrder) ToTrimmed() *TrimmedSpotLimitOrder {
	return &TrimmedSpotLimitOrder{
		Price:     m.OrderInfo.Price,
		Quantity:  m.OrderInfo.Quantity,
		Fillable:  m.Fillable,
		IsBuy:     m.IsBuy(),
		OrderHash: common.BytesToHash(m.OrderHash).Hex(),
		Cid:       m.Cid(),
	}
}

// ComputeOrderHash computes the order hash for given spot limit order
func (m *SpotOrder) ComputeOrderHash(nonce uint32) (common.Hash, error) {
	triggerPrice := ""
	if m.TriggerPrice != nil {
		triggerPrice = m.TriggerPrice.String()
	}

	return types.ComputeSpotOrderHash(m.MarketId, string(m.OrderType), triggerPrice, &m.OrderInfo, nonce)
}
