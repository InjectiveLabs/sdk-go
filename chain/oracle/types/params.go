package types

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	KeyPythProVerifierContract     = []byte("PythProVerifierContract")
	KeyPythProVerificationGasLimit = []byte("PythProVerificationGasLimit")
	KeyPythProVerificationFee      = []byte("PythProVerificationFee")
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
		paramtypes.NewParamSetPair(KeyPythProVerifierContract, &p.PythProVerifierContract, validatePythProVerifierContract),
		paramtypes.NewParamSetPair(KeyPythProVerificationGasLimit, &p.PythProVerificationGasLimit, validateUint64Param),
		paramtypes.NewParamSetPair(KeyPythProVerificationFee, &p.PythProVerificationFee, validateUint64Param),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		PythContract:                             "",
		ChainlinkVerifierProxyContract:           "",
		ChainlinkDataStreamsVerificationGasLimit: 500_000,
		PythProVerifierContract:                  "",
		PythProVerificationGasLimit:              500_000,
		PythProVerificationFee:                   1,
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
	if err := validatePythProVerifierContract(p.PythProVerifierContract); err != nil {
		return fmt.Errorf("pyth_pro_verifier_contract is incorrect: %w", err)
	}
	if p.ChainlinkDataStreamsVerificationGasLimit == 0 {
		return errors.New("chainlink_data_streams_verification_gas_limit must be greater than zero")
	}
	if p.PythProVerificationGasLimit == 0 {
		return errors.New("pyth_pro_verification_gas_limit must be greater than zero")
	}
	if err := ValidateSedaFastParams(p.SedaFastParams); err != nil {
		return fmt.Errorf("seda_fast_params is incorrect: %w", err)
	}
	return nil
}

func validateSedaFastPublicKey(key []byte) error {
	switch len(key) {
	case 0:
		return nil
	case 33:
		if key[0] != 0x02 && key[0] != 0x03 {
			return fmt.Errorf("public_key: compressed SEC1 key must start with 0x02 or 0x03, got 0x%02x", key[0])
		}
		if _, err := crypto.DecompressPubkey(key); err != nil {
			return fmt.Errorf("public_key: not a valid secp256k1 compressed point: %w", err)
		}
	case 65:
		if key[0] != 0x04 {
			return fmt.Errorf("public_key: uncompressed SEC1 key must start with 0x04, got 0x%02x", key[0])
		}
		if _, err := crypto.UnmarshalPubkey(key); err != nil {
			return fmt.Errorf("public_key: not a valid secp256k1 uncompressed point: %w", err)
		}
	default:
		return fmt.Errorf("public_key: invalid SEC1 length %d (must be 33 or 65 bytes)", len(key))
	}
	return nil
}

// ValidateSedaFastParams validates the SedaFastParams configuration.
// An empty public_key is allowed and means the SEDA Fast oracle is disabled.
func ValidateSedaFastParams(p SedaFastParams) error {
	if err := validateSedaFastPublicKey(p.PublicKey); err != nil {
		return err
	}

	if err := validateSedaFastProgramIDs("simple_program_ids", p.SimpleProgramIds); err != nil {
		return err
	}
	if err := validateSedaFastProgramIDs("json_program_ids", p.JsonProgramIds); err != nil {
		return err
	}

	simpleSet := make(map[string]struct{}, len(p.SimpleProgramIds))
	for _, id := range p.SimpleProgramIds {
		simpleSet[id] = struct{}{}
	}
	for _, id := range p.JsonProgramIds {
		if _, dup := simpleSet[id]; dup {
			return fmt.Errorf("program id %q appears in both simple_program_ids and json_program_ids", id)
		}
	}

	return nil
}

func validateSedaFastProgramIDs(field string, ids []string) error {
	seen := make(map[string]struct{}, len(ids))
	for _, id := range ids {
		if len(id) != 64 {
			return fmt.Errorf("%s: program id %q must be 64 hex characters, got %d", field, id, len(id))
		}
		if _, err := hex.DecodeString(id); err != nil {
			return fmt.Errorf("%s: program id %q is not valid hex: %w", field, id, err)
		}
		if strings.ToLower(id) != id {
			return fmt.Errorf("%s: program id %q must be lowercase hex", field, id)
		}
		if _, dup := seen[id]; dup {
			return fmt.Errorf("%s: duplicate program id %q", field, id)
		}
		seen[id] = struct{}{}
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

func validatePythProVerifierContract(i any) error {
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

func validateUint64Param(i any) error {
	if _, ok := i.(uint64); !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
