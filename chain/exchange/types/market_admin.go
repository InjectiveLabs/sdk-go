package types

const (
	TickerPerm                 = 1 << iota
	MinPriceTickSizePerm       = 1 << iota
	MinQuantityTickSizePerm    = 1 << iota
	MinNotionalPerm            = 1 << iota
	InitialMarginRatioPerm     = 1 << iota
	MaintenanceMarginRatioPerm = 1 << iota
	ReduceMarginRatioPerm      = 1 << iota

	MaxPerm = TickerPerm | MinPriceTickSizePerm | MinQuantityTickSizePerm | MinNotionalPerm |
		InitialMarginRatioPerm | MaintenanceMarginRatioPerm | ReduceMarginRatioPerm
)

type MarketAdminPermissions int

func (p MarketAdminPermissions) HasPerm(pp MarketAdminPermissions) bool {
	return p&pp != 0
}
