package types

import (
	"cosmossdk.io/math"
)

type PositionPayout struct {
	Payout       math.LegacyDec
	PnlNotional  math.LegacyDec
	IsProfitable bool
}
