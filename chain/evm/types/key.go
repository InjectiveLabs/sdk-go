package types

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName string name of module
	ModuleName = "evm"

	// StoreKey key for ethereum storage data, account code (StateDB) or block
	// related data for Web3.
	// The EVM module should use a prefix store.
	StoreKey = ModuleName

	// ObjectStoreKey is the key to access the EVM object store, that is reset
	// during the Commit phase.
	ObjectStoreKey = "object:" + ModuleName

	// RouterKey uses module name for routing
	RouterKey = ModuleName
)

// prefix bytes for the EVM persistent store
const (
	prefixCode = iota + 1
	prefixStorage
	prefixParams
)

// prefix bytes for the EVM object store
const (
	prefixObjectBloom = iota + 1
	prefixObjectGasUsed
	prefixObjectParams
	prefixObjectAnteVerifiedSender
)

// KVStore key prefixes
var (
	KeyPrefixCode    = []byte{prefixCode}
	KeyPrefixStorage = []byte{prefixStorage}
	KeyPrefixParams  = []byte{prefixParams}
)

// Object Store key prefixes
var (
	KeyPrefixObjectBloom   = []byte{prefixObjectBloom}
	KeyPrefixObjectGasUsed = []byte{prefixObjectGasUsed}
	// cache the `EVMBlockConfig` during the whole block execution
	KeyPrefixObjectParams = []byte{prefixObjectParams}
	// senders verified by the EVM ante handler, consumed by the MsgEthereumTx handler
	KeyPrefixObjectAnteVerifiedSender = []byte{prefixObjectAnteVerifiedSender}
)

// AddressStoragePrefix returns a prefix to iterate over a given account storage.
func AddressStoragePrefix(address common.Address) []byte {
	return append(KeyPrefixStorage, address.Bytes()...)
}

// StateKey defines the full key under which an account state is stored.
func StateKey(address common.Address, key []byte) []byte {
	return append(AddressStoragePrefix(address), key...)
}

func ObjectBloomKey(txIndex, msgIndex int) []byte {
	var key [1 + 8 + 8]byte
	key[0] = prefixObjectBloom
	binary.BigEndian.PutUint64(key[1:], uint64(txIndex))
	binary.BigEndian.PutUint64(key[9:], uint64(msgIndex))
	return key[:]
}

// ObjectAnteVerifiedSenderKey returns the object store key that records the sender the EVM ante handler verified
// for the Ethereum tx with the given hash, delivered at the given index in the block.
func ObjectAnteVerifiedSenderKey(txIndex int, ethTxHash common.Hash) []byte {
	var key [1 + 8 + common.HashLength]byte
	key[0] = prefixObjectAnteVerifiedSender
	binary.BigEndian.PutUint64(key[1:], uint64(txIndex))
	copy(key[9:], ethTxHash.Bytes())
	return key[:]
}
