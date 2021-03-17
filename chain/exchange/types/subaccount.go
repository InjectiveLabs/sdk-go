package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Subaccount struct {
	Address common.Address
	Nonce   *big.Int
}

func IsValidSubaccountID(subaccountID string) (*common.Address, bool) {
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
