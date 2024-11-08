package v2

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	v1 "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

func NewMarketOrderForLiquidation(
	position *Position,
	positionSubaccountID common.Hash,
	liquidator sdk.AccAddress,
	worstPrice math.LegacyDec,
) *DerivativeMarketOrder {
	var (
		orderType OrderType
	)

	// if long position, market sell order
	// if short position, market buy order
	if position.IsLong {
		orderType = OrderType_SELL
	} else {
		orderType = OrderType_BUY
	}

	order := DerivativeMarketOrder{
		OrderInfo: OrderInfo{
			SubaccountId: positionSubaccountID.Hex(),
			FeeRecipient: liquidator.String(),
			Price:        worstPrice,
			Quantity:     position.Quantity,
		},
		OrderType:    orderType,
		Margin:       math.LegacyZeroDec(),
		MarginHold:   math.LegacyZeroDec(),
		TriggerPrice: nil,
	}

	return &order
}

func (m *DerivativeLimitOrder) ToTrimmed() *TrimmedDerivativeLimitOrder {
	return &TrimmedDerivativeLimitOrder{
		Price:     m.OrderInfo.Price,
		Quantity:  m.OrderInfo.Quantity,
		Margin:    m.Margin,
		Fillable:  m.Fillable,
		IsBuy:     m.IsBuy(),
		OrderHash: common.BytesToHash(m.OrderHash).Hex(),
		Cid:       m.Cid(),
	}
}

func (m *DerivativeLimitOrder) ToStandardized() *TrimmedLimitOrder {
	return &TrimmedLimitOrder{
		Price:        m.OrderInfo.Price,
		Quantity:     m.OrderInfo.Quantity,
		OrderHash:    common.BytesToHash(m.OrderHash).Hex(),
		SubaccountId: m.OrderInfo.SubaccountId,
	}
}

func (o *DerivativeMarketOrderCancel) GetCancelDepositDelta() *v1.DepositDelta {
	order := o.MarketOrder
	// no market order quantity was executed, so refund the entire margin hold
	if order.IsVanilla() && o.CancelQuantity.Equal(order.OrderInfo.Quantity) {
		return &v1.DepositDelta{
			AvailableBalanceDelta: order.MarginHold,
			TotalBalanceDelta:     math.LegacyZeroDec(),
		}
	}
	return nil
}

func (o *DerivativeMarketOrder) GetCancelRefundAmount() math.LegacyDec {
	if o.IsVanilla() {
		return o.MarginHold
	}
	return math.LegacyZeroDec()
}

func NewDerivativeMarketOrder(o *DerivativeOrder, sender sdk.AccAddress, orderHash common.Hash) *DerivativeMarketOrder {
	if o.OrderInfo.FeeRecipient == "" {
		o.OrderInfo.FeeRecipient = sender.String()
	}
	return &DerivativeMarketOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		MarginHold:   math.LegacyZeroDec(),
		TriggerPrice: o.TriggerPrice,
		OrderHash:    orderHash.Bytes(),
	}
}

func NewDerivativeLimitOrder(o *DerivativeOrder, sender sdk.AccAddress, orderHash common.Hash) *DerivativeLimitOrder {
	if o.OrderInfo.FeeRecipient == "" {
		o.OrderInfo.FeeRecipient = sender.String()
	}
	return &DerivativeLimitOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		Fillable:     o.OrderInfo.Quantity,
		TriggerPrice: o.TriggerPrice,
		OrderHash:    orderHash.Bytes(),
	}
}

func (m *DerivativeLimitOrder) ToDerivativeOrder(marketID string) *DerivativeOrder {
	return &DerivativeOrder{
		MarketId:     marketID,
		OrderInfo:    m.OrderInfo,
		OrderType:    m.OrderType,
		Margin:       m.Margin,
		TriggerPrice: m.TriggerPrice,
	}
}
func (o *DerivativeMarketOrder) ToDerivativeOrder(marketID string) *DerivativeOrder {
	return &DerivativeOrder{
		MarketId:     marketID,
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		TriggerPrice: o.TriggerPrice,
	}
}

func (m *DerivativeLimitOrder) HasEqualOrWorsePrice(price math.LegacyDec) bool {
	// the buy order has a worse price than the input price if it's less than
	if m.IsBuy() {
		return m.Price().LTE(price)
	}
	return m.Price().GTE(price)
}

func (o *DerivativeMarketOrder) HasEqualOrWorsePrice(price math.LegacyDec) bool {
	// the buy order has a worse price than the input price if it's less than
	if o.IsBuy() {
		return o.Price().LTE(price)
	}
	return o.Price().GTE(price)
}

func (o *DerivativeMarketOrder) ResizeReduceOnlyOrder(
	newQuantity math.LegacyDec,
	oracleScaleFactor uint32,
	isBinaryOptionsOrder bool,
) {
	quantityDecrement := o.OrderInfo.Quantity.Sub(newQuantity)

	// No-op if increasing quantity or order is a vanilla order
	if !quantityDecrement.IsPositive() || o.IsVanilla() {
		return
	}

	if isBinaryOptionsOrder {
		o.OrderInfo.Quantity = newQuantity
		if o.IsVanilla() {
			o.Margin = o.GetRequiredBinaryOptionsMargin(oracleScaleFactor)
		}
	} else {
		o.Margin = o.Margin.Mul(newQuantity).Quo(o.OrderInfo.Quantity)
		o.OrderInfo.Quantity = newQuantity
	}
}

func (m *DerivativeLimitOrder) GetRequiredBinaryOptionsMargin(oracleScaleFactor uint32) math.LegacyDec {
	// Margin = Price * Quantity for buys
	if m.IsBuy() {
		notional := m.Price().Mul(m.OrderInfo.Quantity)
		return notional
	}
	// Margin = (scaled(1) - Price) * Quantity for sells
	return m.OrderInfo.Quantity.Mul(types.GetScaledPrice(math.LegacyOneDec(), oracleScaleFactor).Sub(m.Price()))
}

func (o *DerivativeMarketOrder) GetRequiredBinaryOptionsMargin(oracleScaleFactor uint32) math.LegacyDec {
	// Margin = Price * Quantity for buys
	if o.IsBuy() {
		notional := o.Price().Mul(o.OrderInfo.Quantity)
		return notional
	}
	// Margin = (scaled(1) - Price) * Quantity for sells
	return o.OrderInfo.Quantity.Mul(types.GetScaledPrice(math.LegacyOneDec(), oracleScaleFactor).Sub(o.Price()))
}

func (m *DerivativeLimitOrder) GetCancelDepositDelta(feeRate math.LegacyDec) *v1.DepositDelta {
	return &v1.DepositDelta{
		AvailableBalanceDelta: m.GetCancelRefundAmount(feeRate),
		TotalBalanceDelta:     math.LegacyZeroDec(),
	}
}

func (m *DerivativeLimitOrder) GetCancelRefundAmount(feeRate math.LegacyDec) math.LegacyDec {
	marginHoldRefund := math.LegacyZeroDec()
	if m.IsVanilla() {
		// negative fees are only accounted for upon matching
		positiveFeePart := math.LegacyMaxDec(math.LegacyZeroDec(), feeRate)
		//nolint:all
		// Refund = (FillableQuantity / Quantity) * (Margin + Price * Quantity * feeRate)
		notional := m.OrderInfo.Price.Mul(m.OrderInfo.Quantity)
		marginHoldRefund = m.Fillable.Mul(m.Margin.Add(notional.Mul(positiveFeePart))).Quo(m.OrderInfo.Quantity)
	}
	return marginHoldRefund
}

func (o *DerivativeOrder) CheckTickSize(minPriceTickSize, minQuantityTickSize math.LegacyDec) error {
	if types.BreachesMinimumTickSize(o.OrderInfo.Price, minPriceTickSize) {
		return errors.Wrapf(types.ErrInvalidPrice, "price %s must be a multiple of the minimum price tick size %s", o.OrderInfo.Price.String(), minPriceTickSize.String())
	}
	if types.BreachesMinimumTickSize(o.OrderInfo.Quantity, minQuantityTickSize) {
		return errors.Wrapf(types.ErrInvalidQuantity, "quantity %s must be a multiple of the minimum quantity tick size %s", o.OrderInfo.Quantity.String(), minQuantityTickSize.String())
	}
	return nil
}

func (o *DerivativeOrder) CheckNotional(minNotional math.LegacyDec) error {
	orderNotional := o.GetQuantity().Mul(o.GetPrice())
	if !minNotional.IsNil() && orderNotional.LT(minNotional) {
		return errors.Wrapf(types.ErrInvalidNotional, "order notional (%s) is less than the minimum notional for the market (%s)", orderNotional.String(), minNotional.String())
	}
	return nil
}

func (o *DerivativeOrder) GetRequiredBinaryOptionsMargin(oracleScaleFactor uint32) math.LegacyDec {
	// Margin = Price * Quantity for buys
	if o.IsBuy() {
		notional := o.Price().Mul(o.OrderInfo.Quantity)
		return notional
	}
	// Margin = (scaled(1) - Price) * Quantity for sells
	return o.OrderInfo.Quantity.Mul(types.GetScaledPrice(math.LegacyOneDec(), oracleScaleFactor).Sub(o.Price()))
}

func (o *DerivativeOrder) CheckMarginAndGetMarginHold(initialMarginRatio, executionMarkPrice, feeRate math.LegacyDec, marketType v1.MarketType, oracleScaleFactor uint32) (marginHold math.LegacyDec, err error) {
	notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
	positiveFeeRatePart := math.LegacyMaxDec(feeRate, math.LegacyZeroDec())
	feeAmount := notional.Mul(positiveFeeRatePart)

	marginHold = o.Margin.Add(feeAmount)
	if marketType == v1.MarketType_BinaryOption {
		requiredMargin := o.GetRequiredBinaryOptionsMargin(oracleScaleFactor)
		if !o.Margin.Equal(requiredMargin) {
			return math.LegacyDec{}, errors.Wrapf(types.ErrInsufficientMargin, "margin check: need %s but got %s", requiredMargin.String(), o.Margin.String())
		}
		return marginHold, nil
	}

	// For perpetual and expiry futures margins
	// Enforce that Margin ≥ InitialMarginRatio * Price * Quantity
	if o.Margin.LT(initialMarginRatio.Mul(notional)) {
		return math.LegacyDec{}, errors.Wrapf(types.ErrInsufficientMargin, "InitialMarginRatio Check: need at least %s but got %s", initialMarginRatio.Mul(notional).String(), o.Margin.String())
	}

	if err := o.CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, executionMarkPrice); err != nil {
		return math.LegacyDec{}, err
	}

	return marginHold, nil
}

func (o *DerivativeOrder) CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice math.LegacyDec) (err error) {
	// For Buys: MarkPrice ≥ (Margin - Price * Quantity) / ((InitialMarginRatio - 1) * Quantity)
	// For Sells: MarkPrice ≤ (Margin + Price * Quantity) / ((1 + InitialMarginRatio) * Quantity)
	markPriceThreshold := o.ComputeInitialMarginRequirementMarkPriceThreshold(initialMarginRatio)
	return CheckInitialMarginMarkPriceRequirement(o.IsBuy(), markPriceThreshold, markPrice)
}

func CheckInitialMarginMarkPriceRequirement(isBuyOrLong bool, markPriceThreshold, markPrice math.LegacyDec) error {
	if isBuyOrLong && markPrice.LT(markPriceThreshold) {
		return errors.Wrapf(types.ErrInsufficientMargin, "Buy MarkPriceThreshold Check: mark/trigger price %s must be GTE %s", markPrice.String(), markPriceThreshold.String())
	} else if !isBuyOrLong && markPrice.GT(markPriceThreshold) {
		return errors.Wrapf(types.ErrInsufficientMargin, "Sell MarkPriceThreshold Check: mark/trigger price %s must be LTE %s", markPrice.String(), markPriceThreshold.String())
	}
	return nil
}

// CheckValidConditionalPrice checks that conditional order type (STOP or TAKE) actually valid for current relation between triggerPrice and markPrice
func (o *DerivativeOrder) CheckValidConditionalPrice(markPrice math.LegacyDec) (err error) {
	if !o.IsConditional() {
		return nil
	}

	ok := true
	switch o.OrderType {
	case OrderType_STOP_BUY, OrderType_TAKE_SELL: // higher
		ok = o.TriggerPrice.GT(markPrice)
	case OrderType_STOP_SELL, OrderType_TAKE_BUY: // lower
		ok = o.TriggerPrice.LT(markPrice)
	}
	if !ok {
		return errors.Wrapf(types.ErrInvalidTriggerPrice, "order type %s incompatible with trigger price %s and markPrice %s", o.OrderType.String(), o.TriggerPrice.String(), markPrice.String())
	}
	return nil
}

// CheckBinaryOptionsPricesWithinBounds checks that binary options order prices don't exceed 1 (scaled)
func (o *DerivativeOrder) CheckBinaryOptionsPricesWithinBounds(oracleScaleFactor uint32) (err error) {
	maxScaledPrice := types.GetScaledPrice(math.LegacyOneDec(), oracleScaleFactor)
	if o.Price().GTE(maxScaledPrice) {
		return errors.Wrapf(types.ErrInvalidPrice, "price must be less than %s", maxScaledPrice.String())
	}

	if o.IsConditional() && o.TriggerPrice.GTE(maxScaledPrice) {
		return errors.Wrapf(types.ErrInvalidTriggerPrice, "trigger price must be less than %s", maxScaledPrice.String())
	}
	return nil
}

func (o *DerivativeOrder) ComputeInitialMarginRequirementMarkPriceThreshold(initialMarginRatio math.LegacyDec) math.LegacyDec {
	return ComputeMarkPriceThreshold(o.IsBuy(), o.Price(), o.GetQuantity(), o.Margin, initialMarginRatio)
}

func ComputeMarkPriceThreshold(isBuyOrLong bool, price, quantity, margin, initialMarginRatio math.LegacyDec) math.LegacyDec {
	notional := price.Mul(quantity)
	var numerator, denominator math.LegacyDec
	if isBuyOrLong {
		numerator = margin.Sub(notional)
		denominator = initialMarginRatio.Sub(math.LegacyOneDec()).Mul(quantity)
	} else {
		numerator = margin.Add(notional)
		denominator = initialMarginRatio.Add(math.LegacyOneDec()).Mul(quantity)
	}
	return numerator.Quo(denominator)
}

func (m *DerivativeLimitOrder) CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice math.LegacyDec) (err error) {
	return m.ToDerivativeOrder("").CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice)
}

func (o *DerivativeMarketOrder) CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice math.LegacyDec) (err error) {
	return o.ToDerivativeOrder("").CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice)
}

func (o *DerivativeMarketOrder) ComputeOrderHash(nonce uint32, marketId string) (common.Hash, error) {
	triggerPrice := ""
	if o.TriggerPrice != nil {
		triggerPrice = o.TriggerPrice.String()
	}

	return types.ComputeOrderHash(marketId, o.OrderInfo.SubaccountId, o.OrderInfo.FeeRecipient, o.OrderInfo.Price.String(), o.OrderInfo.Quantity.String(), o.Margin.String(), triggerPrice, string(o.OrderType), nonce)
}

// ComputeOrderHash computes the order hash for given derivative limit order
func (o *DerivativeOrder) ComputeOrderHash(nonce uint32) (common.Hash, error) {
	triggerPrice := ""
	if o.TriggerPrice != nil {
		triggerPrice = o.TriggerPrice.String()
	}
	return types.ComputeOrderHash(o.MarketId, o.OrderInfo.SubaccountId, o.OrderInfo.FeeRecipient, o.OrderInfo.Price.String(), o.OrderInfo.Quantity.String(), o.Margin.String(), triggerPrice, string(o.OrderType), nonce)
}

func (o *DerivativeOrder) IsReduceOnly() bool {
	return o.Margin.IsZero()
}

func (o *DerivativeMarketOrder) IsReduceOnly() bool {
	return o.Margin.IsZero()
}

func (m *DerivativeLimitOrder) IsReduceOnly() bool {
	return m.Margin.IsZero()
}

func (m *DerivativeLimitOrder) Hash() common.Hash {
	return common.BytesToHash(m.OrderHash)
}

func (o *DerivativeMarketOrder) Hash() common.Hash {
	return common.BytesToHash(o.OrderHash)
}

func (m *DerivativeLimitOrder) FeeRecipient() common.Address {
	return m.OrderInfo.FeeRecipientAddress()
}

func (o *DerivativeMarketOrder) FeeRecipient() common.Address {
	return o.OrderInfo.FeeRecipientAddress()
}

func (o *DerivativeOrder) IsVanilla() bool {
	return !o.IsReduceOnly()
}

func NewV1DerivativeMarketOrderFromV2(market DerivativeMarket, order DerivativeMarketOrder) v1.DerivativeMarketOrder {
	v1OrderInfo := NewV1OrderInfoFromV2(&market, order.OrderInfo)
	v1Order := v1.DerivativeMarketOrder{
		OrderInfo:  v1OrderInfo,
		OrderType:  v1.OrderType(order.OrderType),
		Margin:     market.NotionalToChainFormat(order.Margin),
		MarginHold: market.NotionalToChainFormat(order.MarginHold),
		OrderHash:  order.OrderHash,
	}

	if order.TriggerPrice != nil {
		chainFormatTriggerPrice := market.PriceToChainFormat(*order.TriggerPrice)
		v1Order.TriggerPrice = &chainFormatTriggerPrice
	}

	return v1Order
}

func NewV1DerivativeLimitOrderFromV2(market MarketInterface, order DerivativeLimitOrder) v1.DerivativeLimitOrder {
	v1OrderInfo := NewV1OrderInfoFromV2(market, order.OrderInfo)
	v1Order := v1.DerivativeLimitOrder{
		OrderInfo: v1OrderInfo,
		OrderType: v1.OrderType(order.OrderType),
		Margin:    market.NotionalToChainFormat(order.Margin),
		Fillable:  order.Fillable,
		OrderHash: order.OrderHash,
	}

	if order.TriggerPrice != nil {
		chainFormatTriggerPrice := market.PriceToChainFormat(*order.TriggerPrice)
		v1Order.TriggerPrice = &chainFormatTriggerPrice
	}

	return v1Order
}

func (o *DerivativeMarketOrder) IsVanilla() bool {
	return !o.IsReduceOnly()
}

func (m *DerivativeLimitOrder) IsVanilla() bool {
	return !m.IsReduceOnly()
}

func (m *DerivativeMarketOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *DerivativeLimitOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *DerivativeOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *DerivativeMarketOrder) Quantity() math.LegacyDec {
	return m.OrderInfo.Quantity
}

func (m *DerivativeMarketOrder) FillableQuantity() math.LegacyDec {
	return m.OrderInfo.Quantity
}

func (m *DerivativeMarketOrder) Price() math.LegacyDec {
	return m.OrderInfo.Price
}

func (m *DerivativeLimitOrder) Price() math.LegacyDec {
	return m.OrderInfo.Price
}

func (m *DerivativeOrder) Price() math.LegacyDec {
	return m.OrderInfo.Price
}

func (o *DerivativeOrder) IsConditional() bool {
	return o.OrderType.IsConditional()
}

func (o *DerivativeMarketOrder) IsConditional() bool {
	return o.OrderType.IsConditional()
}

func (m *DerivativeLimitOrder) IsConditional() bool {
	return m.OrderType.IsConditional()
}

func (m *DerivativeLimitOrder) Cid() string {
	return m.OrderInfo.GetCid()
}

func (o *DerivativeMarketOrder) Cid() string {
	return o.OrderInfo.GetCid()
}

func NewV2DerivativeOrderFromV1(market MarketInterface, order v1.DerivativeOrder) *DerivativeOrder {
	humanMargin := market.NotionalFromChainFormat(order.Margin)
	v2OrderInfo := NewV2OrderInfoFromV1(market, order.OrderInfo)
	v2Order := DerivativeOrder{
		MarketId:  order.MarketId,
		OrderInfo: *v2OrderInfo,
		OrderType: OrderType(order.OrderType),
		Margin:    humanMargin,
	}

	if order.TriggerPrice != nil && !order.TriggerPrice.IsNil() {
		humanPrice := market.PriceFromChainFormat(*order.TriggerPrice)
		v2Order.TriggerPrice = &humanPrice
	}

	return &v2Order
}

func (o *DerivativeOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeOrder) Cid() string {
	return o.OrderInfo.GetCid()
}

func (o *DerivativeOrder) IsFromDefaultSubaccount() bool {
	return o.OrderInfo.IsFromDefaultSubaccount()
}

func (o *DerivativeOrder) MarketID() common.Hash {
	return common.HexToHash(o.MarketId)
}

func (o *DerivativeMarketOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (m *DerivativeLimitOrder) SubaccountID() common.Hash {
	return m.OrderInfo.SubaccountID()
}

func (m *OrderInfo) SubaccountID() common.Hash {
	return common.HexToHash(m.SubaccountId)
}

func (m *OrderInfo) FeeRecipientAddress() common.Address {
	address, _ := sdk.AccAddressFromBech32(m.FeeRecipient)
	return common.BytesToAddress(address.Bytes())
}

func (m *DerivativeLimitOrder) SdkAccAddress() sdk.AccAddress {
	return sdk.AccAddress(m.SubaccountID().Bytes()[:common.AddressLength])
}

func (m *DerivativeLimitOrder) IsFromDefaultSubaccount() bool {
	return m.OrderInfo.IsFromDefaultSubaccount()
}

func (o *DerivativeMarketOrder) SdkAccAddress() sdk.AccAddress {
	return sdk.AccAddress(o.SubaccountID().Bytes()[:common.AddressLength])
}

func (o *DerivativeMarketOrder) IsFromDefaultSubaccount() bool {
	return o.OrderInfo.IsFromDefaultSubaccount()
}

func (o *TrimmedDerivativeLimitOrder) IsReduceOnly() bool {
	return o.Margin.IsZero()
}

func EmptyDerivativeMarketOrderResults() *DerivativeMarketOrderResults {
	return &DerivativeMarketOrderResults{
		Quantity:      math.LegacyZeroDec(),
		Price:         math.LegacyZeroDec(),
		Fee:           math.LegacyZeroDec(),
		PositionDelta: PositionDelta{},
		Payout:        math.LegacyZeroDec(),
	}
}
