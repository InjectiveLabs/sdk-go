package types

const (
	TickerPerm                  = 1 << iota
	MinPriceTickSizePerm        = 1 << iota
	MinQuantityTickSizePerm     = 1 << iota
	MinNotionalPerm             = 1 << iota
	InitialMarginRationPerm     = 1 << iota
	MaintenanceMarginRationPerm = 1 << iota

	MaxPerm = TickerPerm | MinPriceTickSizePerm | MinQuantityTickSizePerm | MinNotionalPerm | InitialMarginRationPerm | MaintenanceMarginRationPerm
)

type MarketAdminPermissions int

func (p MarketAdminPermissions) HasPerm(pp MarketAdminPermissions) bool {
	return p&pp != 0
}
