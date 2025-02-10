package types

import (
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
