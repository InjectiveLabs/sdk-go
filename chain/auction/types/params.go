package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// Auction params default values
const (
	// DefaultAuctionPeriodDurationSeconds represents the number of seconds in two weeks
	DefaultAuctionPeriod int64 = 60 * 60 * 24 * 14
)

// Parameter keys
var (
	KeyAuctionPeriod = []byte("AuctionPeriod")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	auctionPeriod int64,
) Params {
	return Params{
		AuctionPeriod: auctionPeriod,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	// TODO: @albert, add the rest of the parameters
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAuctionPeriod, &p.AuctionPeriod, validateAuctionPeriodDuration),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		AuctionPeriod: DefaultAuctionPeriod,
	}
}

// Validate performs basic validation on auction parameters.
func (p Params) Validate() error {
	if err := validateAuctionPeriodDuration(p.AuctionPeriod); err != nil {
		return err
	}

	return nil
}

func validateAuctionPeriodDuration(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("AuctionPeriodDuration must be positive: %d", v)
	}

	return nil
}
