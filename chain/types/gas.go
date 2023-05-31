package types

import (
	"fmt"
	"math"
	"sync"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

type infiniteGasMeter struct {
	consumed storetypes.Gas
	mux      sync.RWMutex
}

// NewThreadsafeInfiniteGasMeter returns a reference to a new thread-safe infiniteGasMeter.
func NewThreadsafeInfiniteGasMeter() storetypes.GasMeter {
	return &infiniteGasMeter{
		consumed: 0,
	}
}

func (g *infiniteGasMeter) GasRemaining() storetypes.Gas {
	return math.MaxUint64
}

func (g *infiniteGasMeter) GasConsumed() storetypes.Gas {
	g.mux.RLock()
	defer g.mux.RUnlock()

	return g.consumed
}

func (g *infiniteGasMeter) GasConsumedToLimit() storetypes.Gas {
	g.mux.RLock()
	defer g.mux.RUnlock()

	return g.consumed
}

func (g *infiniteGasMeter) Limit() storetypes.Gas {
	return 0
}

func (g *infiniteGasMeter) ConsumeGas(amount storetypes.Gas, descriptor string) {
	g.mux.Lock()
	defer g.mux.Unlock()

	var overflow bool
	// TODO: Should we set the consumed field after overflow checking?
	g.consumed, overflow = addUint64Overflow(g.consumed, amount)
	if overflow {
		panic(storetypes.ErrorGasOverflow{Descriptor: descriptor})
	}
}

// RefundGas will deduct the given amount from the gas consumed. If the amount is greater than the
// gas consumed, the function will panic.
//
// Use case: This functionality enables refunding gas to the trasaction or block gas pools so that
// EVM-compatible chains can fully support the go-ethereum StateDb interface.
// See https://github.com/cosmos/cosmos-sdk/pull/9403 for reference.
func (g *infiniteGasMeter) RefundGas(amount storetypes.Gas, descriptor string) {
	g.mux.Lock()
	defer g.mux.Unlock()

	if g.consumed < amount {
		panic(storetypes.ErrorNegativeGasConsumed{Descriptor: descriptor})
	}

	g.consumed -= amount
}

func (g *infiniteGasMeter) IsPastLimit() bool {
	return false
}

func (g *infiniteGasMeter) IsOutOfGas() bool {
	return false
}

func (g *infiniteGasMeter) String() string {
	g.mux.RLock()
	defer g.mux.RUnlock()

	return fmt.Sprintf("InfiniteGasMeter:\n  consumed: %d", g.consumed)
}

// addUint64Overflow performs the addition operation on two uint64 integers and
// returns a boolean on whether or not the result overflows.
func addUint64Overflow(a, b uint64) (uint64, bool) {
	if math.MaxUint64-a < b {
		return 0, true
	}

	return a + b, false
}
