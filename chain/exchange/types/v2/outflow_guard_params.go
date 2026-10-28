package v2

import (
	"fmt"
	"strings"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"

	erc20types "github.com/InjectiveLabs/sdk-go/chain/erc20/types"
	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	peggytypes "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

// MaxOutflowGuardTrackedDenoms bounds the tracked denom list, which the guard
// prices once per block.
const MaxOutflowGuardTrackedDenoms = 16

// DefaultOutflowGuardMaxOutflowRatio returns the default net outflow cap (10%).
func DefaultOutflowGuardMaxOutflowRatio() math.LegacyDec {
	return math.LegacyNewDecWithPrec(1, 1)
}

// DefaultOutflowGuardParams returns the guard disabled with no tracked denoms.
func DefaultOutflowGuardParams() OutflowGuardParams {
	return OutflowGuardParams{
		Enabled:         false,
		MaxOutflowRatio: DefaultOutflowGuardMaxOutflowRatio(),
		TrackedDenoms:   []OutflowGuardTrackedDenom{},
	}
}

// IsMaxOutflowRatioOmitted reports whether the ratio carries no value. Absent
// bytes decode to nil and a re-marshalled absent value to zero; neither is a
// valid ratio.
func (p OutflowGuardParams) IsMaxOutflowRatioOmitted() bool {
	return p.MaxOutflowRatio.IsNil() || p.MaxOutflowRatio.IsZero()
}

// IsUnset reports whether the group carries no configuration at all, which is
// how params stored or submitted by clients predating the group decode.
func (p OutflowGuardParams) IsUnset() bool {
	return !p.Enabled && p.IsMaxOutflowRatioOmitted() && len(p.TrackedDenoms) == 0
}

// Validate performs basic validation on outflow guard parameters. An unset
// group is accepted: the params decode chokepoint and MsgUpdateParams map it
// to the defaults or the current value.
func (p OutflowGuardParams) Validate() error {
	if p.IsUnset() {
		return nil
	}
	if p.MaxOutflowRatio.IsNil() || !p.MaxOutflowRatio.IsPositive() || p.MaxOutflowRatio.GT(math.LegacyOneDec()) {
		return fmt.Errorf("max_outflow_ratio must be in (0, 1], got %s", p.MaxOutflowRatio)
	}
	if len(p.TrackedDenoms) > MaxOutflowGuardTrackedDenoms {
		return fmt.Errorf("tracked_denoms exceeds maximum of %d entries", MaxOutflowGuardTrackedDenoms)
	}
	seen := make(map[string]struct{}, len(p.TrackedDenoms))
	for i := range p.TrackedDenoms {
		tracked := p.TrackedDenoms[i]
		if err := tracked.Validate(); err != nil {
			return fmt.Errorf("tracked_denoms[%d]: %w", i, err)
		}
		canonical := CanonicalOutflowGuardDenom(tracked.Denom)
		if _, ok := seen[canonical]; ok {
			return fmt.Errorf("tracked_denoms entry %q is duplicated", tracked.Denom)
		}
		seen[canonical] = struct{}{}
	}
	return nil
}

// CanonicalOutflowGuardDenom returns the spelling the bank uses for denoms
// whose identifier is case-insensitive hex: IBC vouchers carry an uppercase
// hash, Peggy and ERC20 tokens a checksummed address. Any other denom is
// returned unchanged.
func CanonicalOutflowGuardDenom(denom string) string {
	peggyPrefix := peggytypes.PeggyDenomPrefix + peggytypes.PeggyDenomSeparator
	switch {
	case types.IsIBCDenom(denom):
		return "ibc/" + strings.ToUpper(strings.TrimPrefix(denom, "ibc/"))
	case strings.HasPrefix(denom, peggyPrefix) && ethcommon.IsHexAddress(strings.TrimPrefix(denom, peggyPrefix)):
		return peggytypes.PeggyDenomString(ethcommon.HexToAddress(strings.TrimPrefix(denom, peggyPrefix)))
	case strings.HasPrefix(denom, erc20types.DenomPrefix) && ethcommon.IsHexAddress(strings.TrimPrefix(denom, erc20types.DenomPrefix)):
		return erc20types.DenomPrefix + ethcommon.HexToAddress(strings.TrimPrefix(denom, erc20types.DenomPrefix)).Hex()
	}
	return denom
}

// Validate performs basic validation on a tracked denom.
func (d OutflowGuardTrackedDenom) Validate() error {
	if err := sdk.ValidateDenom(d.Denom); err != nil {
		return err
	}
	if d.Decimals > types.MaxDecimals {
		return fmt.Errorf("decimals must not exceed %d, got %d", types.MaxDecimals, d.Decimals)
	}
	// The notional and the window sum the values of all tracked denoms, so
	// every feed must price in USD: any other quote yields a base/quote ratio.
	if d.OracleQuote != oracletypes.QuoteUSD {
		return fmt.Errorf("oracle_quote must be %s, got %q", oracletypes.QuoteUSD, d.OracleQuote)
	}
	oracleParams := NewOracleParams(d.OracleBase, d.OracleQuote, 0, d.OracleType)
	return oracleParams.ValidateBasic()
}
