package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// Auction params default values
var (
	// DefaultAuctionPeriod represents the number of seconds in 1 week
	DefaultAuctionPeriod int64 = 60 * 60 * 24 * 7
	// DefaultMinNextBidIncrementRate represents default min increment rate 0.25%
	DefaultMinNextBidIncrementRate = sdk.NewDecWithPrec(25, 4)
)

// Parameter keys
var (
	KeyAuctionPeriod           = []byte("AuctionPeriod")
	KeyMinNextBidIncrementRate = []byte("MinNextBidIncrementRate")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	auctionPeriod int64,
	minNextBidIncrementRate sdk.Dec,
) Params {
	return Params{
		AuctionPeriod:           auctionPeriod,
		MinNextBidIncrementRate: minNextBidIncrementRate,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAuctionPeriod, &p.AuctionPeriod, validateAuctionPeriodDuration),
		paramtypes.NewParamSetPair(KeyMinNextBidIncrementRate, &p.MinNextBidIncrementRate, validateMinNextBidIncrementRate),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		AuctionPeriod:           DefaultAuctionPeriod,
		MinNextBidIncrementRate: DefaultMinNextBidIncrementRate,
	}
}

// Validate performs basic validation on auction parameters.
func (p Params) Validate() error {
	if err := validateAuctionPeriodDuration(p.AuctionPeriod); err != nil {
		return err
	}

	if err := validateMinNextBidIncrementRate(p.MinNextBidIncrementRate); err != nil {
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

func validateMinNextBidIncrementRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("MinNextBidIncrementRate cannot be nil")
	}

	if v.Equal(sdk.ZeroDec()) {
		return fmt.Errorf("MinNextBidIncrementRate must be positive: %s", v.String())
	}

	if v.GT(sdk.NewDecWithPrec(2, 1)) { // > 20%
		return fmt.Errorf("MinNextBidIncrementRate must be equal or less than 20 percent: %s", v.String())
	}

	return nil
}
