package types

import "cosmossdk.io/errors"

const (
	TickerPerm                  = 1 << iota
	MinPriceTickSizePerm        = 1 << iota
	MinQuantityTickSizePerm     = 1 << iota
	MinNotionalPerm             = 1 << iota
	InitialMarginRationPerm     = 1 << iota
	MaintenanceMarginRationPerm = 1 << iota
)

const MaxPerm = TickerPerm | MinPriceTickSizePerm | MinQuantityTickSizePerm | MinNotionalPerm | InitialMarginRationPerm | MaintenanceMarginRationPerm

type MarketAdminPermissions int

func (p MarketAdminPermissions) HasPerm(pp MarketAdminPermissions) bool {
	return p&pp != 0
}

func (p MarketAdminPermissions) CheckSpotMarketPermissions(msg *MsgUpdateSpotMarket) error {
	if msg.HasTickerUpdate() && !p.HasPerm(TickerPerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update market ticker")
	}

	if msg.HasMinPriceTickSizeUpdate() && !p.HasPerm(MinPriceTickSizePerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update min_price_tick_size")
	}

	if msg.HasMinQuantityTickSizeUpdate() && !p.HasPerm(MinQuantityTickSizePerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update market min_quantity_tick_size")
	}

	if msg.HasMinNotionalUpdate() && !p.HasPerm(MinNotionalPerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update market min_notional")
	}

	return nil
}

func (p MarketAdminPermissions) CheckDerivativeMarketPermissions(msg *MsgUpdateDerivativeMarket) error {
	if msg.HasTickerUpdate() && !p.HasPerm(TickerPerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update ticker")
	}

	if msg.HasMinPriceTickSizeUpdate() && !p.HasPerm(MinPriceTickSizePerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update min_price_tick_size")
	}

	if msg.HasMinQuantityTickSizeUpdate() && !p.HasPerm(MinQuantityTickSizePerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update min_quantity_tick_size")
	}

	if msg.HasMinNotionalUpdate() && !p.HasPerm(MinNotionalPerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update market min_notional")
	}

	if msg.HasInitialMarginRatioUpdate() && !p.HasPerm(InitialMarginRationPerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update initial_margin_ratio")
	}

	if msg.HasMaintenanceMarginRatioUpdate() && !p.HasPerm(MaintenanceMarginRationPerm) {
		return errors.Wrap(ErrInvalidAccessLevel, "admin does not have permission to update maintenance_margin_ratio")
	}

	return nil
}

func EmptyAdminInfo() AdminInfo {
	return AdminInfo{
		Admin:            "",
		AdminPermissions: 0,
	}
}