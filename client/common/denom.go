package common

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type Denom struct {
	Description         string
	Base                int
	Quote               int
	MinPriceTickSize    float64
	MinQuantityTickSize float64
}

func LoadMetadata(network Network, marketId string) Denom {
	fileName := getFileAbsPath(fmt.Sprintf("../metadata/assets/%s.ini", network.Name))
	cfg, err := ini.Load(fileName)
	if err != nil {
		panic(err)
	}
	return Denom{
		Description:         cfg.Section(marketId).Key("description").String(),
		Base:                cfg.Section(marketId).Key("base").MustInt(),
		Quote:               cfg.Section(marketId).Key("quote").MustInt(),
		MinPriceTickSize:    cfg.Section(marketId).Key("min_price_tick_size").MustFloat64(),
		MinQuantityTickSize: cfg.Section(marketId).Key("min_quantity_tick_size").MustFloat64(),
	}
}
