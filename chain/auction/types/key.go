package types

const (
	ModuleName = "auction"
	StoreKey   = ModuleName
	TStoreKey  = "transient_auction"
)

var (
	// Keys for store prefixes
	BidsKey            = []byte{0x01}
	AuctionRoundKey    = []byte{0x03}
	KeyEndingTimeStamp = []byte{0x04}
)
