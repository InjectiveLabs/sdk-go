package types

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
)

var _ paramtypes.ParamSet = &Params{}

var (
	LargestDecPrice = math.LegacyMustNewDecFromStr("10000000")
)

const (
	// Each value below is the default value for each parameter when generating the default
	// genesis file.
	DefaultBandIBCEnabled         = false
	DefaultBandIbcRequestInterval = int64(7) // every 7 blocks
	DefaultBandIBCVersion         = "bandchain-1"
	DefaultBandIBCPortID          = "oracle"

	MaxPythExponent = 10
	MinPythExponent = -12
)

// Parameter keys
var (
	KeyPythContract = []byte("PythContract")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPythContract, &p.PythContract, validatePythContract),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		PythContract:                                "",
		ChainlinkVerifierProxyContract:              "",
		AcceptUnverifiedChainlinkDataStreamsReports: true,
		ChainlinkDataStreamsVerificationGasLimit:    500_000,
	}
}

// DefaultBandIBCParams returns a default set of band ibc parameters.
func DefaultBandIBCParams() BandIBCParams {
	return BandIBCParams{
		BandIbcEnabled:     DefaultBandIBCEnabled,
		IbcRequestInterval: DefaultBandIbcRequestInterval,
		IbcVersion:         DefaultBandIBCVersion,
		IbcPortId:          DefaultBandIBCPortID,
	}
}

// Validate performs basic validation on oracle parameters.
func (p Params) Validate() error {
	if err := validatePythContract(p.PythContract); err != nil {
		return fmt.Errorf("pyth_contract is incorrect: %w", err)
	}
	if err := ValidateChainlinkVerifierProxyContract(p.ChainlinkVerifierProxyContract); err != nil {
		return fmt.Errorf("chainlink_verifier_proxy_contract is incorrect: %w", err)
	}
	return nil
}

func DefaultTestBandIbcParams() *BandIBCParams {
	return &BandIBCParams{
		// true if Band IBC should be enabled
		BandIbcEnabled: true,
		// block request interval to send Band IBC prices
		IbcRequestInterval: 10,
		// band IBC source channel
		IbcSourceChannel: "channel-0",
		// band IBC version
		IbcVersion: "bandchain-1",
		// band IBC portID
		IbcPortId: "oracle",
	}
}

func validatePythContract(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == "" {
		return nil
	}

	_, err := sdk.AccAddressFromBech32(v)
	if err != nil {
		return fmt.Errorf("invalid PythContract value: %v", v)
	}

	return nil
}

// ValidateChainlinkVerifierProxyContract validates the Chainlink verifier proxy contract address.
func ValidateChainlinkVerifierProxyContract(i any) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == "" {
		return nil
	}

	if !common.IsHexAddress(v) {
		return fmt.Errorf("invalid Ethereum address: %s", v)
	}

	return nil
}
