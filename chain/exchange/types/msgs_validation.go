package types

import (
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func UpdateSpotMarketMessageValidateBasic(msg UpdateSpotMarketMessage) error {
	if err := ValidateAddress(msg.GetAdmin()); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.GetAdmin())
	}

	if !IsHexHash(msg.GetMarketId()) {
		return errors.Wrap(ErrMarketInvalid, msg.GetMarketId())
	}

	hasNoUpdate := !msg.HasTickerUpdate() &&
		!msg.HasMinPriceTickSizeUpdate() &&
		!msg.HasMinQuantityTickSizeUpdate() &&
		!msg.HasMinNotionalUpdate()

	if hasNoUpdate {
		return errors.Wrap(ErrBadField, "no update value present")
	}

	if len(msg.GetNewTicker()) > MaxTickerLength {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not exceed %d characters", MaxTickerLength)
	}

	if msg.HasMinPriceTickSizeUpdate() {
		if err := ValidateTickSize(msg.GetNewMinPriceTickSize()); err != nil {
			return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
		}
	}

	if msg.HasMinQuantityTickSizeUpdate() {
		if err := ValidateTickSize(msg.GetNewMinQuantityTickSize()); err != nil {
			return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
		}
	}

	if msg.HasMinNotionalUpdate() {
		if err := ValidateMinNotional(msg.GetNewMinNotional()); err != nil {
			return errors.Wrap(ErrInvalidNotional, err.Error())
		}
	}

	return nil
}
