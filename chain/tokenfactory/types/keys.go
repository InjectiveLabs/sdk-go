package types

import (
	"strings"
)

const (
	// ModuleName defines the module name
	ModuleName = "tokenfactory"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tokenfactory"
)

// KeySeparator is used to combine parts of the keys in the store
const KeySeparator = "|"

var (
	DenomAuthorityMetadataKey = []byte{0x01}
	DenomsPrefixKey           = []byte{0x02}
	CreatorPrefixKey          = []byte{0x03}
	AdminPrefixKey            = []byte{0x04}
)

// GetDenomPrefixStore returns the store prefix where all the data associated with a specific denom
// is stored
func GetDenomPrefixStore(denom string) []byte {
	return []byte(strings.Join([]string{string(DenomsPrefixKey), denom, ""}, KeySeparator))
}

// GetCreatorPrefix returns the store prefix where the list of the denoms created by a specific
// creator are stored.
func GetCreatorPrefix(creator string) []byte {
	return []byte(strings.Join([]string{string(CreatorPrefixKey), creator, ""}, KeySeparator))
}

// GetCreatorsPrefix returns the store prefix where a list of all creator addresses are stored
func GetCreatorsPrefix() []byte {
	return append(CreatorPrefixKey, []byte(KeySeparator)...)
}
