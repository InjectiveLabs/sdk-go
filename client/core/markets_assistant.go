package core

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/InjectiveLabs/sdk-go/client/exchange"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/maps"
	"gopkg.in/ini.v1"
)

type MarketsAssistant struct {
	tokensBySymbol    map[string]Token
	tokensByDenom     map[string]Token
	spotMarkets       map[string]SpotMarket
	derivativeMarkets map[string]DerivativeMarket
}

func newMarketsAssistant() MarketsAssistant {
	return MarketsAssistant{
		tokensBySymbol:    make(map[string]Token),
		tokensByDenom:     make(map[string]Token),
		spotMarkets:       make(map[string]SpotMarket),
		derivativeMarkets: make(map[string]DerivativeMarket),
	}
}

// Deprecated: use NewMarketsAssistantUsingExchangeClient instead
func NewMarketsAssistant(networkName string) (MarketsAssistant, error) {
	assistant := newMarketsAssistant()
	fileName := getFileAbsPath(fmt.Sprintf("../metadata/assets/%s.ini", networkName))
	metadataFile, err := ini.Load(fileName)

	if err != nil {
		return assistant, err
	}

	for _, section := range metadataFile.Sections() {
		sectionName := section.Name()
		if strings.HasPrefix(sectionName, "0x") {
			description := section.Key("description").Value()

			decimals, _ := section.Key("quote").Int()
			quoteToken := Token{
				Name:     "",
				Symbol:   "",
				Denom:    "",
				Address:  "",
				Decimals: int32(decimals),
				Logo:     "",
				Updated:  -1,
			}

			minPriceTickSize := decimal.RequireFromString(section.Key("min_price_tick_size").String())
			minQuantityTickSize := decimal.RequireFromString(section.Key("min_quantity_tick_size").String())

			if strings.Contains(description, "Spot") {
				baseDecimals, _ := section.Key("quote").Int()
				baseToken := Token{
					Name:     "",
					Symbol:   "",
					Denom:    "",
					Address:  "",
					Decimals: int32(baseDecimals),
					Logo:     "",
					Updated:  -1,
				}

				market := SpotMarket{
					Id:                  sectionName,
					Status:              "",
					Ticker:              description,
					BaseToken:           baseToken,
					QuoteToken:          quoteToken,
					MakerFeeRate:        decimal.NewFromInt32(0),
					TakerFeeRate:        decimal.NewFromInt32(0),
					ServiceProviderFee:  decimal.NewFromInt32(0),
					MinPriceTickSize:    minPriceTickSize,
					MinQuantityTickSize: minQuantityTickSize,
				}

				assistant.spotMarkets[market.Id] = market
			} else {
				market := DerivativeMarket{
					Id:                     sectionName,
					Status:                 "",
					Ticker:                 description,
					OracleBase:             "",
					OracleQuote:            "",
					OracleType:             "",
					OracleScaleFactor:      1,
					InitialMarginRatio:     decimal.NewFromInt32(0),
					MaintenanceMarginRatio: decimal.NewFromInt32(0),
					QuoteToken:             quoteToken,
					MakerFeeRate:           decimal.NewFromInt32(0),
					TakerFeeRate:           decimal.NewFromInt32(0),
					ServiceProviderFee:     decimal.NewFromInt32(0),
					MinPriceTickSize:       minPriceTickSize,
					MinQuantityTickSize:    minQuantityTickSize,
				}

				assistant.derivativeMarkets[market.Id] = market
			}
		} else {
			if sectionName != "DEFAULT" {
				tokenDecimals, _ := section.Key("decimals").Int()
				newToken := Token{
					Name:     sectionName,
					Symbol:   sectionName,
					Denom:    section.Key("peggy_denom").String(),
					Address:  "",
					Decimals: int32(tokenDecimals),
					Logo:     "",
					Updated:  -1,
				}

				assistant.tokensByDenom[newToken.Denom] = newToken
				assistant.tokensBySymbol[newToken.Symbol] = newToken
			}
		}
	}

	return assistant, nil
}

func NewMarketsAssistantUsingExchangeClient(ctx context.Context, exchangeClient exchange.ExchangeClient) (MarketsAssistant, error) {
	assistant := newMarketsAssistant()
	spotMarketsRequest := spotExchangePB.MarketsRequest{
		MarketStatus: "active",
	}
	spotMarkets, err := exchangeClient.GetSpotMarkets(ctx, &spotMarketsRequest)

	if err != nil {
		return assistant, err
	}

	for _, marketInfo := range spotMarkets.GetMarkets() {
		if len(marketInfo.GetBaseTokenMeta().GetSymbol()) > 0 && len(marketInfo.GetQuoteTokenMeta().GetSymbol()) > 0 {
			var baseTokenSymbol, quoteTokenSymbol string
			if strings.Contains(marketInfo.GetTicker(), "/") {
				baseAndQuote := strings.Split(marketInfo.GetTicker(), "/")
				baseTokenSymbol, quoteTokenSymbol = baseAndQuote[0], baseAndQuote[1]
			} else {
				baseTokenSymbol = marketInfo.GetBaseTokenMeta().GetSymbol()
				quoteTokenSymbol = marketInfo.GetQuoteTokenMeta().GetSymbol()
			}

			baseToken := spotTokenRepresentation(baseTokenSymbol, marketInfo.GetBaseTokenMeta(), marketInfo.GetBaseDenom(), &assistant)
			quoteToken := spotTokenRepresentation(quoteTokenSymbol, marketInfo.GetQuoteTokenMeta(), marketInfo.GetQuoteDenom(), &assistant)

			makerFeeRate := decimal.RequireFromString(marketInfo.GetMakerFeeRate())
			takerFeeRate := decimal.RequireFromString(marketInfo.GetTakerFeeRate())
			serviceProviderFee := decimal.RequireFromString(marketInfo.GetServiceProviderFee())
			minPriceTickSize := decimal.RequireFromString(marketInfo.GetMinPriceTickSize())
			minQuantityTickSize := decimal.RequireFromString(marketInfo.GetMinQuantityTickSize())

			market := SpotMarket{
				Id:                  marketInfo.GetMarketId(),
				Status:              marketInfo.GetMarketStatus(),
				Ticker:              marketInfo.GetTicker(),
				BaseToken:           baseToken,
				QuoteToken:          quoteToken,
				MakerFeeRate:        makerFeeRate,
				TakerFeeRate:        takerFeeRate,
				ServiceProviderFee:  serviceProviderFee,
				MinPriceTickSize:    minPriceTickSize,
				MinQuantityTickSize: minQuantityTickSize,
			}

			assistant.spotMarkets[market.Id] = market
		}
	}

	derivativeMarketsRequest := derivativeExchangePB.MarketsRequest{
		MarketStatus: "active",
	}
	derivativeMarkets, err := exchangeClient.GetDerivativeMarkets(ctx, &derivativeMarketsRequest)

	if err != nil {
		return assistant, err
	}

	for _, marketInfo := range derivativeMarkets.GetMarkets() {
		if len(marketInfo.GetQuoteTokenMeta().GetSymbol()) > 0 {
			quoteTokenSymbol := marketInfo.GetQuoteTokenMeta().GetSymbol()

			quoteToken := derivativeTokenRepresentation(quoteTokenSymbol, marketInfo.GetQuoteTokenMeta(), marketInfo.GetQuoteDenom(), &assistant)

			initialMarginRatio := decimal.RequireFromString(marketInfo.GetInitialMarginRatio())
			maintenanceMarginRatio := decimal.RequireFromString(marketInfo.GetMaintenanceMarginRatio())
			makerFeeRate := decimal.RequireFromString(marketInfo.GetMakerFeeRate())
			takerFeeRate := decimal.RequireFromString(marketInfo.GetTakerFeeRate())
			serviceProviderFee := decimal.RequireFromString(marketInfo.GetServiceProviderFee())
			minPriceTickSize := decimal.RequireFromString(marketInfo.GetMinPriceTickSize())
			minQuantityTickSize := decimal.RequireFromString(marketInfo.GetMinQuantityTickSize())

			market := DerivativeMarket{
				Id:                     marketInfo.GetMarketId(),
				Status:                 marketInfo.GetMarketStatus(),
				Ticker:                 marketInfo.GetTicker(),
				OracleBase:             marketInfo.GetOracleBase(),
				OracleQuote:            marketInfo.GetOracleQuote(),
				OracleType:             marketInfo.GetOracleType(),
				OracleScaleFactor:      marketInfo.GetOracleScaleFactor(),
				InitialMarginRatio:     initialMarginRatio,
				MaintenanceMarginRatio: maintenanceMarginRatio,
				QuoteToken:             quoteToken,
				MakerFeeRate:           makerFeeRate,
				TakerFeeRate:           takerFeeRate,
				ServiceProviderFee:     serviceProviderFee,
				MinPriceTickSize:       minPriceTickSize,
				MinQuantityTickSize:    minQuantityTickSize,
			}

			assistant.derivativeMarkets[market.Id] = market
		}
	}

	return assistant, nil
}

func uniqueSymbol(symbol string, denom string, tokenMetaSymbol string, tokenMetaName string, assistant MarketsAssistant) string {
	uniqueSymbol := denom
	_, isSymbolPresent := assistant.tokensBySymbol[symbol]
	if isSymbolPresent {
		_, isSymbolPresent = assistant.tokensBySymbol[tokenMetaSymbol]
		if isSymbolPresent {
			_, isSymbolPresent = assistant.tokensBySymbol[tokenMetaName]
			if !isSymbolPresent {
				uniqueSymbol = tokenMetaName
			}
		} else {
			uniqueSymbol = tokenMetaSymbol
		}
	} else {
		uniqueSymbol = symbol
	}

	return uniqueSymbol
}

func spotTokenRepresentation(symbol string, tokenMeta *spotExchangePB.TokenMeta, denom string, assistant *MarketsAssistant) Token {
	_, isPresent := assistant.tokensByDenom[denom]

	if !isPresent {
		uniqueSymbol := uniqueSymbol(symbol, denom, tokenMeta.GetSymbol(), tokenMeta.GetName(), *assistant)

		newToken := Token{
			Name:     tokenMeta.GetName(),
			Symbol:   symbol,
			Denom:    denom,
			Address:  tokenMeta.GetAddress(),
			Decimals: tokenMeta.GetDecimals(),
			Logo:     tokenMeta.GetLogo(),
			Updated:  tokenMeta.GetUpdatedAt(),
		}

		assistant.tokensByDenom[denom] = newToken
		assistant.tokensBySymbol[uniqueSymbol] = newToken
	}

	return assistant.tokensByDenom[denom]
}

func derivativeTokenRepresentation(symbol string, tokenMeta *derivativeExchangePB.TokenMeta, denom string, assistant *MarketsAssistant) Token {
	_, isPresent := assistant.tokensByDenom[denom]

	if !isPresent {
		uniqueSymbol := uniqueSymbol(symbol, denom, tokenMeta.GetSymbol(), tokenMeta.GetName(), *assistant)

		newToken := Token{
			Name:     tokenMeta.GetName(),
			Symbol:   symbol,
			Denom:    denom,
			Address:  tokenMeta.GetAddress(),
			Decimals: tokenMeta.GetDecimals(),
			Logo:     tokenMeta.GetLogo(),
			Updated:  tokenMeta.GetUpdatedAt(),
		}

		assistant.tokensByDenom[denom] = newToken
		assistant.tokensBySymbol[uniqueSymbol] = newToken
	}

	return assistant.tokensByDenom[denom]
}

func getFileAbsPath(relativePath string) string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), relativePath)
}

func (assistant MarketsAssistant) AllTokens() map[string]Token {
	return maps.Clone(assistant.tokensBySymbol)
}

func (assistant MarketsAssistant) AllSpotMarkets() map[string]SpotMarket {
	return maps.Clone(assistant.spotMarkets)
}

func (assistant MarketsAssistant) AllDerivativeMarkets() map[string]DerivativeMarket {
	return maps.Clone(assistant.derivativeMarkets)
}
