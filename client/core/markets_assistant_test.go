package core

import (
	"context"
	"strings"
	"testing"

	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"github.com/stretchr/testify/assert"
)

func TestMarketAssistantCreationUsingMarketsFromExchange(t *testing.T) {
	mockExchange := MockExchangeClient{}
	var spotMarketInfos []*spotExchangePB.SpotMarketInfo
	var derivativeMarketInfos []*derivativeExchangePB.DerivativeMarketInfo
	injUsdtSpotMarketInfo := createINJUSDTSpotMarketInfo()
	apeUsdtSpotMarketInfo := createAPEUSDTSpotMarketInfo()
	btcUsdtDerivativeMarketInfo := createBTCUSDTDerivativeMarketInfo()

	spotMarketInfos = append(spotMarketInfos, injUsdtSpotMarketInfo)
	spotMarketInfos = append(spotMarketInfos, apeUsdtSpotMarketInfo)
	derivativeMarketInfos = append(derivativeMarketInfos, btcUsdtDerivativeMarketInfo)

	mockExchange.SpotMarketsResponses = append(mockExchange.SpotMarketsResponses, &spotExchangePB.MarketsResponse{
		Markets: spotMarketInfos,
	})
	mockExchange.DerivativeMarketsResponses = append(mockExchange.DerivativeMarketsResponses, &derivativeExchangePB.MarketsResponse{
		Markets: derivativeMarketInfos,
	})

	ctx := context.Background()
	assistant, err := NewMarketsAssistantUsingExchangeClient(ctx, &mockExchange)

	assert.NoError(t, err)

	tokens := assistant.AllTokens()

	assert.Len(t, tokens, 5)

	symbols := strings.Split(injUsdtSpotMarketInfo.Ticker, "/")
	injSymbol, usdtSymbol := symbols[0], symbols[1]
	symbols = strings.Split(apeUsdtSpotMarketInfo.Ticker, "/")
	apeSymbol := symbols[0]
	alternativeUSDTName := apeUsdtSpotMarketInfo.QuoteTokenMeta.Name
	usdtPerpSymbol := btcUsdtDerivativeMarketInfo.QuoteTokenMeta.Symbol

	_, isPresent := tokens[injSymbol]
	assert.True(t, isPresent)
	_, isPresent = tokens[usdtSymbol]
	assert.True(t, isPresent)
	_, isPresent = tokens[alternativeUSDTName]
	assert.True(t, isPresent)
	_, isPresent = tokens[apeSymbol]
	assert.True(t, isPresent)
	_, isPresent = tokens[usdtPerpSymbol]
	assert.True(t, isPresent)

	spotMarkets := assistant.AllSpotMarkets()
	assert.Len(t, spotMarkets, 2)

	_, isPresent = spotMarkets[injUsdtSpotMarketInfo.MarketId]
	assert.True(t, isPresent)
	_, isPresent = spotMarkets[apeUsdtSpotMarketInfo.MarketId]
	assert.True(t, isPresent)

	derivativeMarkets := assistant.AllDerivativeMarkets()
	assert.Len(t, derivativeMarkets, 1)

	_, isPresent = derivativeMarkets[btcUsdtDerivativeMarketInfo.MarketId]
	assert.True(t, isPresent)
}
