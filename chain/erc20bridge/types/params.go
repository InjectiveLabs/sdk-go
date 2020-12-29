package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
)

// Parameter store key
var (
	ParamStoreKeyHubParams = []byte("hubparams")
	// DefaultGasLimit is gas limit we use for erc20bridge internal ethereum transactions
	DefaultGasLimit = uint64(100_000_000) // 100M
)

// ParamKeyTable - Key declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyHubParams, HubParams{}, validateHubParams),
	)
}

// NewParams creates a new Params object
func NewParams(hp HubParams) Params {
	return Params{
		HubParams: hp,
	}
}

func validateHubParams(i interface{}) error {
	v, ok := i.(HubParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v.HubAddress) != common.AddressLength {
		return fmt.Errorf("invalid hub address: %s", v.HubAddress)
	}

	return nil
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyHubParams, &p.HubParams, validateHubParams),
	}
}
