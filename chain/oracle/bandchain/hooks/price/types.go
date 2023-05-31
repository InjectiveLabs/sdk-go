package price

import "github.com/InjectiveLabs/sdk-go/chain/oracle/bandchain/oracle/types"

type SymbolInput struct {
	Symbols            []string `json:"symbols"`
	MinimumSourceCount uint8    `json:"minimum_source_count"`
}

type SymbolOutput struct {
	Responses []Response `json:"responses"`
}

type Response struct {
	Symbol       string `json:"symbol"`
	ResponseCode uint8  `json:"response_code"`
	Rate         uint64 `json:"rate"`
}

type Input struct {
	Symbols    []string `json:"symbols"`
	Multiplier uint64   `json:"multiplier"`
}

type Output struct {
	Pxs []uint64 `json:"pxs"`
}

type Price struct {
	Symbol      string          `json:"symbol"`
	Multiplier  uint64          `json:"multiplier"`
	Px          uint64          `json:"px"`
	RequestID   types.RequestID `json:"request_id"`
	ResolveTime int64           `json:"resolve_time"`
}

func NewPrice(symbol string, multiplier, px uint64, reqID types.RequestID, resolveTime int64) Price {
	return Price{
		Symbol:      symbol,
		Multiplier:  multiplier,
		Px:          px,
		RequestID:   reqID,
		ResolveTime: resolveTime,
	}
}
