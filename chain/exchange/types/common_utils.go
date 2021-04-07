package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Subaccount struct {
	Address common.Address
	Nonce   *big.Int
}

func IsValidSubaccountID(subaccountID string) (*common.Address, bool) {
	if len(subaccountID) != 66 {
		return nil, false
	}
	subaccountIdBytes := common.FromHex(subaccountID)

	if len(subaccountIdBytes) != common.HashLength {
		return nil, false
	}
	addressBytes := subaccountIdBytes[:common.AddressLength]
	if !common.IsHexAddress(common.Bytes2Hex(addressBytes)) {
		return nil, false
	}
	address := common.BytesToAddress(addressBytes)
	return &address, true
}

func IsValidOrderHash(orderHash string) bool {
	if len(orderHash) != 66 {
		return false
	}
	orderHashBytes := common.FromHex(orderHash)

	if len(orderHashBytes) != common.HashLength {
		return false
	}
	return true
}
