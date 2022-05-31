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
	// Keys for store prefixes
	BidsKey = []byte{0x01}

	// Keys for Smart contract execution
	ContractExecutionRequestIDKey   = []byte{0x02} // key to the smart contract execution request ID
	LatestSmartContractRequestIDKey = []byte{0x03} // key to the latest smart contract request ID

)

func GetContractExecutionRequestIDKey(requestID uint64) []byte {
	return append(ContractExecutionRequestIDKey, sdk.Uint64ToBigEndian(requestID)...)
}
