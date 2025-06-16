package chain

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/shopspring/decimal"

	"github.com/InjectiveLabs/sdk-go/client/core"
	"github.com/InjectiveLabs/sdk-go/client/exchange"
)

type TokenMetadata interface {
	GetName() string
	GetAddress() string
	GetSymbol() string
	GetLogo() string
	GetDecimals() int32
	GetUpdatedAt() int64
}

type MarketsAssistant struct {
	tokensBySymbol      map[string]core.Token
	tokensByDenom       map[string]core.Token
	spotMarkets         map[string]core.SpotMarket
	derivativeMarkets   map[string]core.DerivativeMarket
	binaryOptionMarkets map[string]core.DerivativeMarket
}

func newMarketsAssistant() MarketsAssistant {
	return MarketsAssistant{
		tokensBySymbol:      make(map[string]core.Token),
		tokensByDenom:       make(map[string]core.Token),
		spotMarkets:         make(map[string]core.SpotMarket),
		derivativeMarkets:   make(map[string]core.DerivativeMarket),
		binaryOptionMarkets: make(map[string]core.DerivativeMarket),
	}
}

type DenomsMetadataProvider interface {
	GetDenomsMetadata(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryDenomsMetadataResponse, error)
}

func NewMarketsAssistant(ctx context.Context, chainClient ChainClient) (MarketsAssistant, error) {
	assistant := newMarketsAssistant()
	err := assistant.initializeFromChainV1Markets(ctx, chainClient)

	return assistant, err
}

func NewHumanReadableMarketsAssistant(ctx context.Context, chainClient ChainClientV2) (MarketsAssistant, error) {
	assistant := newMarketsAssistant()
	err := assistant.initializeFromChainV2Markets(ctx, chainClient)

	return assistant, err
}

func NewMarketsAssistantWithAllTokens(ctx context.Context, exchangeClient exchange.ExchangeClient, chainClient ChainClient) (MarketsAssistant, error) {
	assistant := newMarketsAssistant()
	assistant.initializeTokensFromChainDenoms(ctx, chainClient)
	err := assistant.initializeFromChainV1Markets(ctx, chainClient)

	return assistant, err
}

func NewHumanReadableMarketsAssistantWithAllTokens(ctx context.Context, exchangeClient exchange.ExchangeClient, chainClient ChainClientV2) (MarketsAssistant, error) {
	assistant := newMarketsAssistant()
	assistant.initializeTokensFromChainDenoms(ctx, chainClient)
	err := assistant.initializeFromChainV2Markets(ctx, chainClient)

	return assistant, err
}

func uniqueSymbol(symbol, denom, tokenMetaSymbol, tokenMetaName string, tokensBySymbol map[string]core.Token) string {
	uniqueSymbol := denom
	_, isSymbolPresent := tokensBySymbol[symbol]
	if isSymbolPresent {
		_, isSymbolPresent = tokensBySymbol[tokenMetaSymbol]
		if isSymbolPresent {
			_, isSymbolPresent = tokensBySymbol[tokenMetaName]
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

func tokenRepresentation(symbol string, tokenMeta TokenMetadata, denom string, assistant *MarketsAssistant) core.Token {
	_, isPresent := assistant.tokensByDenom[denom]

	if !isPresent {
		uniqueSymbol := uniqueSymbol(symbol, denom, tokenMeta.GetSymbol(), tokenMeta.GetName(), assistant.tokensBySymbol)

		newToken := core.Token{
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

func (assistant *MarketsAssistant) AllTokens() map[string]core.Token {
	return assistant.tokensBySymbol
}

func (assistant *MarketsAssistant) AllTokensByDenom() map[string]core.Token {
	return assistant.tokensByDenom
}

func (assistant *MarketsAssistant) AllSpotMarkets() map[string]core.SpotMarket {
	return assistant.spotMarkets
}

func (assistant *MarketsAssistant) AllDerivativeMarkets() map[string]core.DerivativeMarket {
	return assistant.derivativeMarkets
}

func (assistant *MarketsAssistant) AllBinaryOptionMarkets() map[string]core.DerivativeMarket {
	return assistant.binaryOptionMarkets
}

func (assistant *MarketsAssistant) initializeTokensFromChainDenoms(ctx context.Context, denomsProvider DenomsMetadataProvider) {
	var denomsMetadata []banktypes.Metadata
	var nextKey []byte

	for readNextPage := true; readNextPage; readNextPage = len(nextKey) > 0 {
		pagination := query.PageRequest{Key: nextKey}
		result, err := denomsProvider.GetDenomsMetadata(ctx, &pagination)

		if err != nil {
			panic(err)
		}

		nextKey = result.GetPagination().GetNextKey()
		denomsMetadata = append(denomsMetadata, result.GetMetadatas()...)
	}

	for i := range denomsMetadata {
		denomMetadata := denomsMetadata[i]
		symbol := denomMetadata.GetSymbol()
		denom := denomMetadata.GetBase()

		_, isDenomPresent := assistant.tokensByDenom[denom]

		if symbol != "" && denom != "" && !isDenomPresent {
			name := denomMetadata.GetName()
			if name == "" {
				name = symbol
			}

			var decimals int32 = -1
			for _, denomUnit := range denomMetadata.GetDenomUnits() {
				exponent := int32(denomUnit.GetExponent())
				if exponent > decimals {
					decimals = exponent
				}
			}

			uniqueSymbol := uniqueSymbol(symbol, denom, symbol, name, assistant.tokensBySymbol)

			newToken := core.Token{
				Name:     name,
				Symbol:   symbol,
				Denom:    denom,
				Address:  "",
				Decimals: decimals,
				Logo:     denomMetadata.GetURI(),
				Updated:  -1,
			}

			assistant.tokensByDenom[denom] = newToken
			assistant.tokensBySymbol[uniqueSymbol] = newToken
		}
	}
}

func (assistant *MarketsAssistant) initializeFromChainV1Markets(ctx context.Context, chainClient ChainClient) error {
	officialTokens, err := core.LoadTokens(chainClient.GetNetwork().OfficialTokensListURL)
	if err == nil {
		for i := range officialTokens {
			tokenMetadata := officialTokens[i]
			if tokenMetadata.Denom != "" {
				// add tokens to the assistant ensuring all of them get assigned a unique symbol
				tokenRepresentation(tokenMetadata.GetSymbol(), tokenMetadata, tokenMetadata.Denom, assistant)
			}
		}
	}

	spotMarkets, err := chainClient.FetchChainSpotMarkets(ctx, "Active", nil)

	if err != nil {
		return err
	}

	for _, marketInfo := range spotMarkets.GetMarkets() {
		baseToken, baseTokenFound := assistant.tokensByDenom[marketInfo.GetBaseDenom()]
		quoteToken, quoteTokenFound := assistant.tokensByDenom[marketInfo.GetQuoteDenom()]

		if !baseTokenFound || !quoteTokenFound {
			// Ignore the market because it references tokens that are not in the token list
			continue
		}

		makerFeeRate := decimal.RequireFromString(marketInfo.MakerFeeRate.String())
		takerFeeRate := decimal.RequireFromString(marketInfo.TakerFeeRate.String())
		serviceProviderFee := decimal.RequireFromString(marketInfo.RelayerFeeShareRate.String())
		minPriceTickSize := decimal.RequireFromString(marketInfo.MinPriceTickSize.String())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.MinQuantityTickSize.String())
		minNotional := decimal.RequireFromString(marketInfo.MinNotional.String())

		market := core.SpotMarketV1{
			Id:                  marketInfo.GetMarketId(),
			Status:              marketInfo.Status.String(),
			Ticker:              marketInfo.GetTicker(),
			BaseToken:           baseToken,
			QuoteToken:          quoteToken,
			MakerFeeRate:        makerFeeRate,
			TakerFeeRate:        takerFeeRate,
			ServiceProviderFee:  serviceProviderFee,
			MinPriceTickSize:    minPriceTickSize,
			MinQuantityTickSize: minQuantityTickSize,
			MinNotional:         minNotional,
			BaseDecimals:        marketInfo.BaseDecimals,
			QuoteDecimals:       marketInfo.QuoteDecimals,
		}

		assistant.spotMarkets[market.Id] = market
	}

	derivativeMarkets, err := chainClient.FetchChainDerivativeMarkets(ctx, "Active", nil, false)

	if err != nil {
		return err
	}

	for _, fullMarket := range derivativeMarkets.GetMarkets() {
		marketInfo := fullMarket.GetMarket()

		quoteToken, quoteTokenFound := assistant.tokensByDenom[marketInfo.QuoteDenom]
		if !quoteTokenFound {
			// Ignore the market because it references a token that is not in the token list
			continue
		}

		initialMarginRatio := decimal.RequireFromString(marketInfo.InitialMarginRatio.String())
		maintenanceMarginRatio := decimal.RequireFromString(marketInfo.MaintenanceMarginRatio.String())
		makerFeeRate := decimal.RequireFromString(marketInfo.MakerFeeRate.String())
		takerFeeRate := decimal.RequireFromString(marketInfo.TakerFeeRate.String())
		serviceProviderFee := decimal.RequireFromString(marketInfo.RelayerFeeShareRate.String())
		minPriceTickSize := decimal.RequireFromString(marketInfo.MinPriceTickSize.String())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.MinQuantityTickSize.String())
		minNotional := decimal.RequireFromString(marketInfo.MinNotional.String())

		market := core.DerivativeMarketV1{
			Id:                     marketInfo.MarketId,
			Status:                 marketInfo.Status.String(),
			Ticker:                 marketInfo.Ticker,
			OracleBase:             marketInfo.OracleBase,
			OracleQuote:            marketInfo.OracleQuote,
			OracleType:             marketInfo.OracleType.String(),
			OracleScaleFactor:      marketInfo.OracleScaleFactor,
			InitialMarginRatio:     initialMarginRatio,
			MaintenanceMarginRatio: maintenanceMarginRatio,
			QuoteToken:             quoteToken,
			MakerFeeRate:           makerFeeRate,
			TakerFeeRate:           takerFeeRate,
			ServiceProviderFee:     serviceProviderFee,
			MinPriceTickSize:       minPriceTickSize,
			MinQuantityTickSize:    minQuantityTickSize,
			MinNotional:            minNotional,
			QuoteDecimals:          marketInfo.QuoteDecimals,
		}

		assistant.derivativeMarkets[market.Id] = market
	}

	binaryOptionsMarkets, err := chainClient.FetchChainBinaryOptionsMarkets(ctx, "Active")

	if err != nil {
		return err
	}

	for _, marketInfo := range binaryOptionsMarkets.GetMarkets() {
		quoteToken, quoteTokenFound := assistant.tokensByDenom[marketInfo.QuoteDenom]
		if !quoteTokenFound {
			// Ignore the market because it references a token that is not in the token list
			continue
		}

		makerFeeRate := decimal.RequireFromString(marketInfo.MakerFeeRate.String())
		takerFeeRate := decimal.RequireFromString(marketInfo.TakerFeeRate.String())
		serviceProviderFee := decimal.RequireFromString(marketInfo.RelayerFeeShareRate.String())
		minPriceTickSize := decimal.RequireFromString(marketInfo.MinPriceTickSize.String())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.MinQuantityTickSize.String())
		minNotional := decimal.RequireFromString(marketInfo.MinNotional.String())

		market := core.BinaryOptionMarketV1{
			Id:                  marketInfo.MarketId,
			Status:              marketInfo.Status.String(),
			Ticker:              marketInfo.Ticker,
			OracleSymbol:        marketInfo.OracleSymbol,
			OracleProvider:      marketInfo.OracleProvider,
			OracleType:          marketInfo.OracleType.String(),
			OracleScaleFactor:   marketInfo.OracleScaleFactor,
			ExpirationTimestamp: marketInfo.ExpirationTimestamp,
			SettlementTimestamp: marketInfo.SettlementTimestamp,
			QuoteToken:          quoteToken,
			MakerFeeRate:        makerFeeRate,
			TakerFeeRate:        takerFeeRate,
			ServiceProviderFee:  serviceProviderFee,
			MinPriceTickSize:    minPriceTickSize,
			MinQuantityTickSize: minQuantityTickSize,
			MinNotional:         minNotional,
			QuoteDecimals:       marketInfo.QuoteDecimals,
		}

		if marketInfo.SettlementPrice != nil {
			settlementPrice := decimal.RequireFromString(marketInfo.SettlementPrice.String())
			market.SettlementPrice = &settlementPrice
		}

		assistant.binaryOptionMarkets[market.Id] = market
	}

	return nil
}

func (assistant *MarketsAssistant) initializeFromChainV2Markets(ctx context.Context, chainClient ChainClientV2) error {
	officialTokens, err := core.LoadTokens(chainClient.GetNetwork().OfficialTokensListURL)
	if err == nil {
		for i := range officialTokens {
			tokenMetadata := officialTokens[i]
			if tokenMetadata.Denom != "" {
				// add tokens to the assistant ensuring all of them get assigned a unique symbol
				tokenRepresentation(tokenMetadata.GetSymbol(), tokenMetadata, tokenMetadata.Denom, assistant)
			}
		}
	}

	spotMarkets, err := chainClient.FetchChainSpotMarkets(ctx, "Active", nil)

	if err != nil {
		return err
	}

	for _, marketInfo := range spotMarkets.GetMarkets() {
		baseToken, baseTokenFound := assistant.tokensByDenom[marketInfo.GetBaseDenom()]
		quoteToken, quoteTokenFound := assistant.tokensByDenom[marketInfo.GetQuoteDenom()]

		if !baseTokenFound || !quoteTokenFound {
			// Ignore the market because it references tokens that are not in the token list
			continue
		}

		makerFeeRate := decimal.RequireFromString(marketInfo.MakerFeeRate.String())
		takerFeeRate := decimal.RequireFromString(marketInfo.TakerFeeRate.String())
		serviceProviderFee := decimal.RequireFromString(marketInfo.RelayerFeeShareRate.String())
		minPriceTickSize := decimal.RequireFromString(marketInfo.MinPriceTickSize.String())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.MinQuantityTickSize.String())
		minNotional := decimal.RequireFromString(marketInfo.MinNotional.String())

		market := core.SpotMarketV2{
			Id:                  marketInfo.GetMarketId(),
			Status:              marketInfo.Status.String(),
			Ticker:              marketInfo.GetTicker(),
			BaseToken:           baseToken,
			QuoteToken:          quoteToken,
			MakerFeeRate:        makerFeeRate,
			TakerFeeRate:        takerFeeRate,
			ServiceProviderFee:  serviceProviderFee,
			MinPriceTickSize:    minPriceTickSize,
			MinQuantityTickSize: minQuantityTickSize,
			MinNotional:         minNotional,
			BaseDecimals:        marketInfo.BaseDecimals,
			QuoteDecimals:       marketInfo.QuoteDecimals,
		}

		assistant.spotMarkets[market.Id] = market
	}

	derivativeMarkets, err := chainClient.FetchChainDerivativeMarkets(ctx, "Active", nil, false)

	if err != nil {
		return err
	}

	for _, fullMarket := range derivativeMarkets.GetMarkets() {
		marketInfo := fullMarket.GetMarket()

		quoteToken, quoteTokenFound := assistant.tokensByDenom[marketInfo.QuoteDenom]
		if !quoteTokenFound {
			// Ignore the market because it references a token that is not in the token list
			continue
		}

		initialMarginRatio := decimal.RequireFromString(marketInfo.InitialMarginRatio.String())
		maintenanceMarginRatio := decimal.RequireFromString(marketInfo.MaintenanceMarginRatio.String())
		reduceMarginRatio := decimal.RequireFromString(marketInfo.ReduceMarginRatio.String())
		makerFeeRate := decimal.RequireFromString(marketInfo.MakerFeeRate.String())
		takerFeeRate := decimal.RequireFromString(marketInfo.TakerFeeRate.String())
		serviceProviderFee := decimal.RequireFromString(marketInfo.RelayerFeeShareRate.String())
		minPriceTickSize := decimal.RequireFromString(marketInfo.MinPriceTickSize.String())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.MinQuantityTickSize.String())
		minNotional := decimal.RequireFromString(marketInfo.MinNotional.String())

		market := core.DerivativeMarketV2{
			Id:                     marketInfo.MarketId,
			Status:                 marketInfo.Status.String(),
			Ticker:                 marketInfo.Ticker,
			OracleBase:             marketInfo.OracleBase,
			OracleQuote:            marketInfo.OracleQuote,
			OracleType:             marketInfo.OracleType.String(),
			OracleScaleFactor:      marketInfo.OracleScaleFactor,
			InitialMarginRatio:     initialMarginRatio,
			MaintenanceMarginRatio: maintenanceMarginRatio,
			ReduceMarginRatio:      reduceMarginRatio,
			QuoteToken:             quoteToken,
			MakerFeeRate:           makerFeeRate,
			TakerFeeRate:           takerFeeRate,
			ServiceProviderFee:     serviceProviderFee,
			MinPriceTickSize:       minPriceTickSize,
			MinQuantityTickSize:    minQuantityTickSize,
			MinNotional:            minNotional,
			QuoteDecimals:          marketInfo.QuoteDecimals,
		}

		assistant.derivativeMarkets[market.Id] = market
	}

	binaryOptionsMarkets, err := chainClient.FetchChainBinaryOptionsMarkets(ctx, "Active")

	if err != nil {
		return err
	}

	for _, marketInfo := range binaryOptionsMarkets.GetMarkets() {
		quoteToken, quoteTokenFound := assistant.tokensByDenom[marketInfo.QuoteDenom]
		if !quoteTokenFound {
			// Ignore the market because it references a token that is not in the token list
			continue
		}

		makerFeeRate := decimal.RequireFromString(marketInfo.MakerFeeRate.String())
		takerFeeRate := decimal.RequireFromString(marketInfo.TakerFeeRate.String())
		serviceProviderFee := decimal.RequireFromString(marketInfo.RelayerFeeShareRate.String())
		minPriceTickSize := decimal.RequireFromString(marketInfo.MinPriceTickSize.String())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.MinQuantityTickSize.String())
		minNotional := decimal.RequireFromString(marketInfo.MinNotional.String())

		market := core.BinaryOptionMarketV2{
			Id:                  marketInfo.MarketId,
			Status:              marketInfo.Status.String(),
			Ticker:              marketInfo.Ticker,
			OracleSymbol:        marketInfo.OracleSymbol,
			OracleProvider:      marketInfo.OracleProvider,
			OracleType:          marketInfo.OracleType.String(),
			OracleScaleFactor:   marketInfo.OracleScaleFactor,
			ExpirationTimestamp: marketInfo.ExpirationTimestamp,
			SettlementTimestamp: marketInfo.SettlementTimestamp,
			QuoteToken:          quoteToken,
			MakerFeeRate:        makerFeeRate,
			TakerFeeRate:        takerFeeRate,
			ServiceProviderFee:  serviceProviderFee,
			MinPriceTickSize:    minPriceTickSize,
			MinQuantityTickSize: minQuantityTickSize,
			MinNotional:         minNotional,
			QuoteDecimals:       marketInfo.QuoteDecimals,
		}

		if marketInfo.SettlementPrice != nil {
			settlementPrice := decimal.RequireFromString(marketInfo.SettlementPrice.String())
			market.SettlementPrice = &settlementPrice
		}

		assistant.binaryOptionMarkets[market.Id] = market
	}

	return nil
}
