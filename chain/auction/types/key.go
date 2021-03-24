package types

const (
	ModuleName = "auction"
	StoreKey   = ModuleName
)

var (
	// Keys for store prefixes
	BidsKey            = []byte{0x01}
	KeyLastBid         = []byte{0x02}
	KeyAuctionRound    = []byte{0x03}
	KeyEndingTimeStamp = []byte{0x04}
)
