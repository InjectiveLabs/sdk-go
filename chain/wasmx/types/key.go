package types

// ModuleName and StoreKey are retained for historical account decoding and
// deleting the retired store. They must never be reused for another module.
const (
	ModuleName = "xwasm"
	StoreKey   = ModuleName
)
