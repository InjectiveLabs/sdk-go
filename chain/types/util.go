package types

import (
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func HasDuplicate[T comparable](elements []T) bool {
	seen := make(map[T]struct{})
	for idx := range elements {
		if _, ok := seen[elements[idx]]; ok {
			return true // Duplicate found
		}
		seen[elements[idx]] = struct{}{}
	}
	return false
}

func HasDuplicateCoins(slice []sdk.Coin) bool {
	seen := make(map[string]struct{})
	for _, item := range slice {
		if _, ok := seen[item.Denom]; ok {
			return true
		}
		seen[item.Denom] = struct{}{}
	}
	return false
}

// IterCb is a callback for IterateSafe that receives key and value bytes.
// Return true to stop iteration.
type IterCb func(k, v []byte) (stop bool)

// IterateSafe ensures the Iterator is closed even if the work done inside the callback panics.
func IterateSafe(iter storetypes.Iterator, callback IterCb) {
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		if callback(iter.Key(), iter.Value()) {
			return
		}
	}
}

// IterKeyCb is a callback for IterateKeysSafe that receives only the key bytes.
// Return true to stop iteration.
type IterKeyCb func(k []byte) (stop bool)

// IterateKeysSafe only iterates over keys and ensures the Iterator is closed even if the callback panics.
func IterateKeysSafe(iter storetypes.Iterator, callback IterKeyCb) {
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		if callback(iter.Key()) {
			return
		}
	}
}
