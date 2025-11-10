package types

const (
	TickerPerm = 1 << iota
	MinPriceTickSizePerm
	MinQuantityTickSizePerm
	MinNotionalPerm
	InitialMarginRatioPerm
	MaintenanceMarginRatioPerm
	ReduceMarginRatioPerm
	OpenNotionalCapPerm

	MaxPerm = TickerPerm | MinPriceTickSizePerm | MinQuantityTickSizePerm | MinNotionalPerm |
		InitialMarginRatioPerm | MaintenanceMarginRatioPerm | ReduceMarginRatioPerm | OpenNotionalCapPerm
)

type MarketAdminPermissions int

func (p MarketAdminPermissions) HasPerm(pp MarketAdminPermissions) bool {
	return p&pp != 0
}
