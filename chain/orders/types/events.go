package types

// Orders module event types
const (
	EventTypeNewOrder                  = "new_order"
	EventTypeNewDerivativeOrder        = "new_derivative_order"
	EventTypeSoftCancelOrder           = "soft_cancel"
	EventTypeHardCancelDerivativeOrder = "hard_cancel_derivative_order"
	EventTypeFillDerivativeOrder       = "fill_derivative_order"
	EventTypeHardCancelSpotOrder       = "hard_cancel_spot_order"
	EventTypeFillSpotOrder             = "fill_spot_order"

	AttributeKeyOrderHash     = "order_hash"
	AttributeKeyMarketID      = "market_id"
	AttributeKeyTradePairHash = "trade_pair_hash"
	AttributeKeySignedOrder   = "signed_order"
	AttributeKeyFilledAmount  = "filled_amount"
)
