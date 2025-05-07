package types

import (
	"fmt"
	"math"
	"math/big"
	"math/bits"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	maxWordLen = sdkmath.MaxBitLen / bits.UintSize
)

var MaxInt256 *big.Int

func init() {
	var tmp big.Int
	MaxInt256 = tmp.Lsh(big.NewInt(1), sdkmath.MaxBitLen).Sub(&tmp, big.NewInt(1))
}

// SafeInt64 checks for overflows while casting a uint64 to int64 value.
func SafeInt64(value uint64) (int64, error) {
	if value > uint64(math.MaxInt64) {
		return 0, errorsmod.Wrapf(errortypes.ErrInvalidHeight, "uint64 value %v cannot exceed %v", value, int64(math.MaxInt64))
	}

	return int64(value), nil
}

func SafeInt(value uint) (int, error) {
	if value > uint(math.MaxInt64) {
		return 0, errorsmod.Wrapf(errortypes.ErrInvalidHeight, "uint value %v cannot exceed %v", value, int(math.MaxInt64))
	}

	return int(value), nil
}

// SafeNewIntFromBigInt constructs Int from big.Int, return error if more than 256bits
func SafeNewIntFromBigInt(i *big.Int) (sdkmath.Int, error) {
	if !IsValidInt256(i) {
		return sdkmath.NewInt(0), fmt.Errorf("big int out of bound: %s", i)
	}
	return sdkmath.NewIntFromBigInt(i), nil
}

// SaturatedNewInt constructs Int from big.Int, truncate if more than 256bits
func SaturatedNewInt(i *big.Int) sdkmath.Int {
	if !IsValidInt256(i) {
		i = MaxInt256
	}
	return sdkmath.NewIntFromBigInt(i)
}

// IsValidInt256 check the bound of 256 bit number
func IsValidInt256(i *big.Int) bool {
	return i == nil || !bigIntOverflows(i)
}

// check if the big int overflows,
// NOTE: copied from cosmos-sdk.
func bigIntOverflows(i *big.Int) bool {
	// overflow is defined as i.BitLen() > MaxBitLen
	// however this check can be expensive when doing many operations.
	// So we first check if the word length is greater than maxWordLen.
	// However the most significant word could be zero, hence we still do the bitlen check.
	if len(i.Bits()) > maxWordLen {
		return i.BitLen() > sdkmath.MaxBitLen
	}
	return false
}
