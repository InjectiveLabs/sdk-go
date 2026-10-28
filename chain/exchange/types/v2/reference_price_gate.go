package v2

import (
	"cosmossdk.io/errors"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
)

const (
	// MaxReferenceWindowSeconds bounds the maximum age of a provider
	// observation used by an enabled reference-price gate.
	MaxReferenceWindowSeconds = 86400
	// MaxOpenDeviationBps bounds the maximum allowed deviation between a
	// position-opening fill price and the reference price (100%).
	MaxOpenDeviationBps = 10000
)

// IsEnabled reports whether the reference-price gate is turned on for the
// market. Note that an enabled gate whose source cannot be resolved blocks
// position-opening fills entirely (fail-closed).
func (c ReferencePriceGateConfig) IsEnabled() bool {
	return c.ReferencePriceSource != ReferencePriceSource_REFERENCE_PRICE_SOURCE_DISABLED
}

// isSupportedReferencePriceSource reports whether the chain implements the
// given reference price source. MULTI_ORACLE_TWAP and COMPOSITE_LONG remain
// reserved for the multi-source framework and are rejected until implemented.
func isSupportedReferencePriceSource(source ReferencePriceSource) bool {
	switch source {
	case ReferencePriceSource_REFERENCE_PRICE_SOURCE_DISABLED,
		ReferencePriceSource_REFERENCE_PRICE_SOURCE_PYTH_EMA_LONG,
		ReferencePriceSource_REFERENCE_PRICE_SOURCE_PYTH_PRO_EMA_LONG,
		ReferencePriceSource_REFERENCE_PRICE_SOURCE_CHAINLINK_DATA_STREAMS_SPOT,
		ReferencePriceSource_REFERENCE_PRICE_SOURCE_SEDA_FAST_SPOT:
		return true
	default:
		return false
	}
}

// Validate checks the shape of the config and that its source is implemented
// by the chain.
func (c ReferencePriceGateConfig) Validate() error {
	if c.ReferencePriceSource == ReferencePriceSource_REFERENCE_PRICE_SOURCE_DISABLED {
		if c.ReferenceWindowSeconds != 0 || c.MaxOpenDeviationBps != 0 {
			return errors.Wrap(
				types.ErrInvalidReferencePriceGateConfig,
				"reference_window_seconds and max_open_deviation_bps must be zero when the source is disabled",
			)
		}
		return nil
	}

	if !isSupportedReferencePriceSource(c.ReferencePriceSource) {
		switch c.ReferencePriceSource {
		case ReferencePriceSource_REFERENCE_PRICE_SOURCE_MULTI_ORACLE_TWAP,
			ReferencePriceSource_REFERENCE_PRICE_SOURCE_COMPOSITE_LONG:
			return errors.Wrapf(
				types.ErrUnsupportedReferencePriceSource,
				"reference_price_source %s is not implemented", c.ReferencePriceSource.String(),
			)
		default:
		}
		return errors.Wrapf(
			types.ErrInvalidReferencePriceGateConfig,
			"unknown reference_price_source value %d", c.ReferencePriceSource,
		)
	}
	return c.validateEnabledShape()
}

// ValidateForMarket is the COMPLETE check for a config about to be persisted:
// it runs Validate (shape and source support) and then the market-dependent
// requirement on top. Every provider-specific source reads raw state keyed by
// the market's OracleBase/OracleQuote, so the source must match the market's
// oracle type.
// Composing Validate here is deliberate: the stateful write boundaries (the
// market-update msg server, the params-update proposal path, genesis import)
// reach a config that may never have passed a msg's ValidateBasic — direct
// msg-server or internal execution would otherwise persist a malformed config
// such as DISABLED with nonzero fields, or an unimplemented enum source. A
// config that somehow still bypasses this fails closed at resolution.
func (c ReferencePriceGateConfig) ValidateForMarket(oracleType oracletypes.OracleType) error {
	if err := c.Validate(); err != nil {
		return err
	}
	requiredOracleType, hasRequirement := c.requiredOracleType()
	if hasRequirement && oracleType != requiredOracleType {
		if c.ReferencePriceSource == ReferencePriceSource_REFERENCE_PRICE_SOURCE_PYTH_EMA_LONG {
			return errors.Wrapf(
				types.ErrInvalidReferencePriceGateConfig,
				"reference_price_source %s requires a Pyth oracle market, got oracle type %s",
				c.ReferencePriceSource.String(), oracleType.String(),
			)
		}
		return errors.Wrapf(
			types.ErrInvalidReferencePriceGateConfig,
			"reference_price_source %s requires oracle type %s, got %s",
			c.ReferencePriceSource.String(), requiredOracleType.String(), oracleType.String(),
		)
	}
	return nil
}

func (c ReferencePriceGateConfig) requiredOracleType() (oracletypes.OracleType, bool) {
	switch c.ReferencePriceSource {
	case ReferencePriceSource_REFERENCE_PRICE_SOURCE_PYTH_EMA_LONG:
		return oracletypes.OracleType_Pyth, true
	case ReferencePriceSource_REFERENCE_PRICE_SOURCE_PYTH_PRO_EMA_LONG:
		return oracletypes.OracleType_PythPro, true
	case ReferencePriceSource_REFERENCE_PRICE_SOURCE_CHAINLINK_DATA_STREAMS_SPOT:
		return oracletypes.OracleType_ChainlinkDataStreams, true
	case ReferencePriceSource_REFERENCE_PRICE_SOURCE_SEDA_FAST_SPOT:
		return oracletypes.OracleType_SedaFast, true
	default:
		return oracletypes.OracleType_Unspecified, false
	}
}

func (c ReferencePriceGateConfig) validateEnabledShape() error {
	if c.ReferenceWindowSeconds == 0 || c.ReferenceWindowSeconds > MaxReferenceWindowSeconds {
		return errors.Wrapf(
			types.ErrInvalidReferencePriceGateConfig,
			"reference_window_seconds must be within [1, %d], got %d",
			MaxReferenceWindowSeconds, c.ReferenceWindowSeconds,
		)
	}
	if c.MaxOpenDeviationBps == 0 || c.MaxOpenDeviationBps > MaxOpenDeviationBps {
		return errors.Wrapf(
			types.ErrInvalidReferencePriceGateConfig,
			"max_open_deviation_bps must be within [1, %d], got %d",
			MaxOpenDeviationBps, c.MaxOpenDeviationBps,
		)
	}
	return nil
}
