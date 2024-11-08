package v2

import v1 "github.com/InjectiveLabs/sdk-go/chain/exchange/types"

func NewV1FullSpotMarketFromV2(fullSpotMarket FullSpotMarket) v1.FullSpotMarket {
	v1SpotMarket := NewV1SpotMarketFromV2(*fullSpotMarket.Market)
	newFullSpotMarket := v1.FullSpotMarket{
		Market: &v1SpotMarket,
	}

	if fullSpotMarket.MidPriceAndTob != nil {
		v1MidPriceAndTOB := NewV1MidPriceAndTOBFromV2(fullSpotMarket.Market, *fullSpotMarket.MidPriceAndTob)
		newFullSpotMarket.MidPriceAndTob = &v1MidPriceAndTOB
	}

	return newFullSpotMarket
}

func NewV1FullDerivativeMarketFromV2(fullDerivativeMarket FullDerivativeMarket) v1.FullDerivativeMarket {
	v1FullMarket := v1.FullDerivativeMarket{}

	switch info := fullDerivativeMarket.Info.(type) {
	case *FullDerivativeMarket_FuturesInfo:
		v1FuturesInfo := NewV1FuturesInfoFromV2(fullDerivativeMarket.Market, *info)
		v1FullMarket.Info = &v1FuturesInfo
	case *FullDerivativeMarket_PerpetualInfo:
		v1PerpetualInfo := NewV1PerpetualInfoFromV2(fullDerivativeMarket.Market, *info)
		v1FullMarket.Info = &v1PerpetualInfo
	}

	v1FullMarket.MarkPrice = fullDerivativeMarket.Market.PriceToChainFormat(fullDerivativeMarket.MarkPrice)

	v1DerivativeMarket := NewV1DerivativeMarketFromV2(*fullDerivativeMarket.Market)
	v1FullMarket.Market = &v1DerivativeMarket

	if fullDerivativeMarket.MidPriceAndTob != nil {
		v1MidPriceAndTOB := NewV1MidPriceAndTOBFromV2(fullDerivativeMarket.Market, *fullDerivativeMarket.MidPriceAndTob)
		v1FullMarket.MidPriceAndTob = &v1MidPriceAndTOB
	}

	return v1FullMarket
}

func NewV1MidPriceAndTOBFromV2(market MarketInterface, midPriceAndTOB MidPriceAndTOB) v1.MidPriceAndTOB {
	chainFormatMidPrice := market.PriceToChainFormat(*midPriceAndTOB.MidPrice)
	chainFormatBestBuyPrice := market.PriceToChainFormat(*midPriceAndTOB.BestBuyPrice)
	chainFormatBestSellPrice := market.PriceToChainFormat(*midPriceAndTOB.BestSellPrice)
	return v1.MidPriceAndTOB{
		MidPrice:      &chainFormatMidPrice,
		BestBuyPrice:  &chainFormatBestBuyPrice,
		BestSellPrice: &chainFormatBestSellPrice,
	}
}

func NewV1FuturesInfoFromV2(market MarketInterface, info FullDerivativeMarket_FuturesInfo) v1.FullDerivativeMarket_FuturesInfo {
	v1FuturesInfo := v1.ExpiryFuturesMarketInfo{
		MarketId:                           info.FuturesInfo.MarketId,
		ExpirationTimestamp:                info.FuturesInfo.ExpirationTimestamp,
		TwapStartTimestamp:                 info.FuturesInfo.TwapStartTimestamp,
		ExpirationTwapStartPriceCumulative: market.PriceToChainFormat(info.FuturesInfo.ExpirationTwapStartPriceCumulative),
		SettlementPrice:                    market.PriceToChainFormat(info.FuturesInfo.SettlementPrice),
	}
	return v1.FullDerivativeMarket_FuturesInfo{
		FuturesInfo: &v1FuturesInfo,
	}
}

func NewV1PerpetualInfoFromV2(market MarketInterface, perpetualInfo FullDerivativeMarket_PerpetualInfo) v1.FullDerivativeMarket_PerpetualInfo {
	v1PerpetualMarketInfo := NewV1PerpetualMarketInfoFromV2(*perpetualInfo.PerpetualInfo.MarketInfo)
	v1FundingInfo := NewV1FundingInfoFromV2(market, *perpetualInfo.PerpetualInfo.FundingInfo)
	return v1.FullDerivativeMarket_PerpetualInfo{
		PerpetualInfo: &v1.PerpetualMarketState{
			MarketInfo:  &v1PerpetualMarketInfo,
			FundingInfo: &v1FundingInfo,
		},
	}
}

func NewV1FundingInfoFromV2(market MarketInterface, fundingInfo PerpetualMarketFunding) v1.PerpetualMarketFunding {
	return v1.PerpetualMarketFunding{
		CumulativeFunding: market.NotionalToChainFormat(fundingInfo.CumulativeFunding),
		CumulativePrice:   market.PriceToChainFormat(fundingInfo.CumulativePrice),
		LastTimestamp:     fundingInfo.LastTimestamp,
	}
}

func NewV1PerpetualMarketInfoFromV2(perpetualMarketInfo PerpetualMarketInfo) v1.PerpetualMarketInfo {
	return v1.PerpetualMarketInfo{
		MarketId:             perpetualMarketInfo.MarketId,
		HourlyFundingRateCap: perpetualMarketInfo.HourlyFundingRateCap,
		HourlyInterestRate:   perpetualMarketInfo.HourlyInterestRate,
		NextFundingTimestamp: perpetualMarketInfo.NextFundingTimestamp,
		FundingInterval:      perpetualMarketInfo.FundingInterval,
	}
}

func NewV1TrimmedDerivativeLimitOrderFromV2(market MarketInterface, trimmedOrder TrimmedDerivativeLimitOrder) v1.TrimmedDerivativeLimitOrder {
	chainFormatPrice := market.PriceToChainFormat(trimmedOrder.Price)
	chainFormatQuantity := market.QuantityToChainFormat(trimmedOrder.Quantity)
	chainFormatMargin := market.NotionalToChainFormat(trimmedOrder.Margin)
	chainFormatFillable := market.QuantityToChainFormat(trimmedOrder.Fillable)
	return v1.TrimmedDerivativeLimitOrder{
		Price:     chainFormatPrice,
		Quantity:  chainFormatQuantity,
		Margin:    chainFormatMargin,
		Fillable:  chainFormatFillable,
		IsBuy:     trimmedOrder.IsBuy,
		OrderHash: trimmedOrder.OrderHash,
		Cid:       trimmedOrder.Cid,
	}
}

func NewV1TrimmedSpotLimitOrderFromV2(market MarketInterface, trimmedOrder *TrimmedSpotLimitOrder) *v1.TrimmedSpotLimitOrder {
	return &v1.TrimmedSpotLimitOrder{
		Price:     market.PriceToChainFormat(trimmedOrder.Price),
		Quantity:  market.QuantityToChainFormat(trimmedOrder.Quantity),
		Fillable:  market.QuantityToChainFormat(trimmedOrder.Fillable),
		IsBuy:     trimmedOrder.IsBuy,
		OrderHash: trimmedOrder.OrderHash,
		Cid:       trimmedOrder.Cid,
	}
}
