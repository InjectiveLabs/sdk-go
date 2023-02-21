package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is set to xwasm and not wasmx to avoid Potential key collision between KVStores
	// (see assertNoPrefix in cosmos-sdk/types/store.go)
	ModuleName = "xwasm"
	StoreKey   = ModuleName
	TStoreKey  = "transient_xwasm"
)

var (
	ContractsByGasPricePrefix = []byte{0x01} // key to the smart contract execution request ID
	ContractsIndexPrefix      = []byte{0x02}
)

func GetContractsByGasPriceKey(gasPrice uint64, address sdk.AccAddress) []byte {
	return append(ContractsByGasPricePrefix, getGasPriceAddressInfix(gasPrice, address)...)
}

func getGasPriceAddressInfix(gasPrice uint64, address sdk.AccAddress) []byte {
	return append(sdk.Uint64ToBigEndian(gasPrice), address.Bytes()...)
}

// GetContractsIndexKey provides the key for the contract address => gasPrice
func GetContractsIndexKey(address sdk.AccAddress) []byte {
	return append(ContractsIndexPrefix, address.Bytes()...)
}
