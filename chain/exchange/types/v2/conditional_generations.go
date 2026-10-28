package v2

var (
	// DerivativeConditionalMarketOrdersIndexPrefix indexes newly admitted conditional market orders.
	// The legacy 0x75 index remains read-only for creation, so its cleanup population cannot refill.
	DerivativeConditionalMarketOrdersIndexPrefix = []byte{0xa8}
	// DerivativeConditionalLimitOrdersIndexPrefix indexes newly admitted conditional limit orders.
	// Both indexes use marketID || direction || subaccountID || orderHash -> triggerPrice.
	DerivativeConditionalLimitOrdersIndexPrefix = []byte{0xa9}
)
