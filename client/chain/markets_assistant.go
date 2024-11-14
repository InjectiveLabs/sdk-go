package chain

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/shopspring/decimal"

	"github.com/InjectiveLabs/sdk-go/client/core"
	"github.com/InjectiveLabs/sdk-go/client/exchange"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
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
	binaryOptionMarkets map[string]core.BinaryOptionMarket
}

func newMarketsAssistant() MarketsAssistant {
	return MarketsAssistant{
		tokensBySymbol:      make(map[string]core.Token),
		tokensByDenom:       make(map[string]core.Token),
		spotMarkets:         make(map[string]core.SpotMarket),
		derivativeMarkets:   make(map[string]core.DerivativeMarket),
		binaryOptionMarkets: make(map[string]core.BinaryOptionMarket),
	}
}

// Deprecated: Use NewMarketsAssistant instead
func NewMarketsAssistantInitializedFromChain(ctx context.Context, exchangeClient exchange.ExchangeClient) (MarketsAssistant, error) {
	assistant := newMarketsAssistant()

	officialTokens, err := core.LoadTokens(exchangeClient.GetNetwork().OfficialTokensListURL)
	if err == nil {
		for i := range officialTokens {
			tokenMetadata := officialTokens[i]
			if tokenMetadata.Denom != "" {
				// add tokens to the assistant ensuring all of them get assigned a unique symbol
				tokenRepresentation(tokenMetadata.GetSymbol(), tokenMetadata, tokenMetadata.Denom, &assistant)
			}
		}
	}

	spotMarketsRequest := spotExchangePB.MarketsRequest{
		MarketStatus: "active",
	}
	spotMarkets, err := exchangeClient.GetSpotMarkets(ctx, &spotMarketsRequest)

	if err != nil {
		return assistant, err
	}

	for _, marketInfo := range spotMarkets.GetMarkets() {
		if marketInfo.GetBaseTokenMeta().GetSymbol() == "" || marketInfo.GetQuoteTokenMeta().GetSymbol() == "" {
			continue
		}

		var baseTokenSymbol, quoteTokenSymbol string
		if strings.Contains(marketInfo.GetTicker(), "/") {
			baseAndQuote := strings.Split(marketInfo.GetTicker(), "/")
			baseTokenSymbol, quoteTokenSymbol = baseAndQuote[0], baseAndQuote[1]
		} else {
			baseTokenSymbol = marketInfo.GetBaseTokenMeta().GetSymbol()
			quoteTokenSymbol = marketInfo.GetQuoteTokenMeta().GetSymbol()
		}

		baseToken := tokenRepresentation(baseTokenSymbol, marketInfo.GetBaseTokenMeta(), marketInfo.GetBaseDenom(), &assistant)
		quoteToken := tokenRepresentation(quoteTokenSymbol, marketInfo.GetQuoteTokenMeta(), marketInfo.GetQuoteDenom(), &assistant)

		makerFeeRate := decimal.RequireFromString(marketInfo.GetMakerFeeRate())
		takerFeeRate := decimal.RequireFromString(marketInfo.GetTakerFeeRate())
		serviceProviderFee := decimal.RequireFromString(marketInfo.GetServiceProviderFee())
		minPriceTickSize := decimal.RequireFromString(marketInfo.GetMinPriceTickSize())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.GetMinQuantityTickSize())
		minNotional := decimal.RequireFromString(marketInfo.GetMinNotional())

		market := core.SpotMarket{
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
			MinNotional:         minNotional,
		}

		assistant.spotMarkets[market.Id] = market
	}

	derivativeMarketsRequest := derivativeExchangePB.MarketsRequest{
		MarketStatus: "active",
	}
	derivativeMarkets, err := exchangeClient.GetDerivativeMarkets(ctx, &derivativeMarketsRequest)

	if err != nil {
		return assistant, err
	}

	for _, marketInfo := range derivativeMarkets.GetMarkets() {
		if marketInfo.GetQuoteTokenMeta().GetSymbol() == "" {
			continue
		}

		quoteTokenSymbol := marketInfo.GetQuoteTokenMeta().GetSymbol()

		quoteToken := tokenRepresentation(quoteTokenSymbol, marketInfo.GetQuoteTokenMeta(), marketInfo.GetQuoteDenom(), &assistant)

		initialMarginRatio := decimal.RequireFromString(marketInfo.GetInitialMarginRatio())
		maintenanceMarginRatio := decimal.RequireFromString(marketInfo.GetMaintenanceMarginRatio())
		makerFeeRate := decimal.RequireFromString(marketInfo.GetMakerFeeRate())
		takerFeeRate := decimal.RequireFromString(marketInfo.GetTakerFeeRate())
		serviceProviderFee := decimal.RequireFromString(marketInfo.GetServiceProviderFee())
		minPriceTickSize := decimal.RequireFromString(marketInfo.GetMinPriceTickSize())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.GetMinQuantityTickSize())
		minNotional := decimal.RequireFromString(marketInfo.GetMinNotional())

		market := core.DerivativeMarket{
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
			MinNotional:            minNotional,
		}

		assistant.derivativeMarkets[market.Id] = market
	}

	return assistant, nil
}

func NewMarketsAssistant(ctx context.Context, chainClient ChainClient) (MarketsAssistant, error) {
	assistant := newMarketsAssistant()
	err := assistant.initializeFromChain(ctx, chainClient)

	return assistant, err
}

func NewMarketsAssistantWithAllTokens(ctx context.Context, exchangeClient exchange.ExchangeClient, chainClient ChainClient) (MarketsAssistant, error) {
	assistant := newMarketsAssistant()
	assistant.initializeTokensFromChainDenoms(ctx, chainClient)
	err := assistant.initializeFromChain(ctx, chainClient)

	return assistant, err
}

func uniqueSymbol(symbol, denom, tokenMetaSymbol, tokenMetaName string, assistant MarketsAssistant) string {
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

func tokenRepresentation(symbol string, tokenMeta TokenMetadata, denom string, assistant *MarketsAssistant) core.Token {
	_, isPresent := assistant.tokensByDenom[denom]

	if !isPresent {
		uniqueSymbol := uniqueSymbol(symbol, denom, tokenMeta.GetSymbol(), tokenMeta.GetName(), *assistant)

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

func (assistant *MarketsAssistant) AllBinaryOptionMarkets() map[string]core.BinaryOptionMarket {
	return assistant.binaryOptionMarkets
}

func (assistant *MarketsAssistant) initializeTokensFromChainDenoms(ctx context.Context, chainClient ChainClient) {
	var denomsMetadata []banktypes.Metadata
	var nextKey []byte

	for readNextPage := true; readNextPage; readNextPage = len(nextKey) > 0 {
		pagination := query.PageRequest{Key: nextKey}
		result, err := chainClient.GetDenomsMetadata(ctx, &pagination)

		if err != nil {
			panic(err)
		}

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

			uniqueSymbol := uniqueSymbol(symbol, denom, symbol, name, *assistant)

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

func (assistant *MarketsAssistant) initializeFromChain(ctx context.Context, chainClient ChainClient) error {
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

		makerFeeRate := decimal.RequireFromString(marketInfo.GetMakerFeeRate().String())
		takerFeeRate := decimal.RequireFromString(marketInfo.GetTakerFeeRate().String())
		serviceProviderFee := decimal.RequireFromString(marketInfo.GetRelayerFeeShareRate().String())
		minPriceTickSize := decimal.RequireFromString(marketInfo.GetMinPriceTickSize().String())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.GetMinQuantityTickSize().String())
		minNotional := decimal.RequireFromString(marketInfo.GetMinNotional().String())

		market := core.SpotMarket{
			Id:                  marketInfo.GetMarketId(),
			Status:              marketInfo.GetMarketStatus().String(),
			Ticker:              marketInfo.GetTicker(),
			BaseToken:           baseToken,
			QuoteToken:          quoteToken,
			MakerFeeRate:        makerFeeRate,
			TakerFeeRate:        takerFeeRate,
			ServiceProviderFee:  serviceProviderFee,
			MinPriceTickSize:    minPriceTickSize,
			MinQuantityTickSize: minQuantityTickSize,
			MinNotional:         minNotional,
		}

		assistant.spotMarkets[market.Id] = market
	}

	derivativeMarkets, err := chainClient.FetchChainDerivativeMarkets(ctx, "Active", nil, false)

	if err != nil {
		return err
	}

	for _, fullMarket := range derivativeMarkets.GetMarkets() {
		marketInfo := fullMarket.GetMarket()

		quoteToken, quoteTokenFound := assistant.tokensByDenom[marketInfo.GetQuoteDenom()]
		if !quoteTokenFound {
			// Ignore the market because it references a token that is not in the token list
			continue
		}

		initialMarginRatio := decimal.RequireFromString(marketInfo.GetInitialMarginRatio().String())
		maintenanceMarginRatio := decimal.RequireFromString(marketInfo.MaintenanceMarginRatio.String())
		makerFeeRate := decimal.RequireFromString(marketInfo.GetMakerFeeRate().String())
		takerFeeRate := decimal.RequireFromString(marketInfo.GetTakerFeeRate().String())
		serviceProviderFee := decimal.RequireFromString(marketInfo.GetRelayerFeeShareRate().String())
		minPriceTickSize := decimal.RequireFromString(marketInfo.GetMinPriceTickSize().String())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.GetMinQuantityTickSize().String())
		minNotional := decimal.RequireFromString(marketInfo.GetMinNotional().String())

		market := core.DerivativeMarket{
			Id:                     marketInfo.MarketId,
			Status:                 marketInfo.GetMarketStatus().String(),
			Ticker:                 marketInfo.GetTicker(),
			OracleBase:             marketInfo.OracleBase,
			OracleQuote:            marketInfo.OracleQuote,
			OracleType:             marketInfo.OracleType.String(),
			OracleScaleFactor:      marketInfo.GetOracleScaleFactor(),
			InitialMarginRatio:     initialMarginRatio,
			MaintenanceMarginRatio: maintenanceMarginRatio,
			QuoteToken:             quoteToken,
			MakerFeeRate:           makerFeeRate,
			TakerFeeRate:           takerFeeRate,
			ServiceProviderFee:     serviceProviderFee,
			MinPriceTickSize:       minPriceTickSize,
			MinQuantityTickSize:    minQuantityTickSize,
			MinNotional:            minNotional,
		}

		assistant.derivativeMarkets[market.Id] = market
	}

	binaryOptionsMarkets, err := chainClient.FetchChainBinaryOptionsMarkets(ctx, "Active")

	if err != nil {
		return err
	}

	for _, marketInfo := range binaryOptionsMarkets.GetMarkets() {
		quoteToken, quoteTokenFound := assistant.tokensByDenom[marketInfo.GetQuoteDenom()]
		if !quoteTokenFound {
			// Ignore the market because it references a token that is not in the token list
			continue
		}

		makerFeeRate := decimal.RequireFromString(marketInfo.GetMakerFeeRate().String())
		takerFeeRate := decimal.RequireFromString(marketInfo.GetTakerFeeRate().String())
		serviceProviderFee := decimal.RequireFromString(marketInfo.GetRelayerFeeShareRate().String())
		minPriceTickSize := decimal.RequireFromString(marketInfo.GetMinPriceTickSize().String())
		minQuantityTickSize := decimal.RequireFromString(marketInfo.GetMinQuantityTickSize().String())
		minNotional := decimal.RequireFromString(marketInfo.GetMinNotional().String())

		market := core.BinaryOptionMarket{
			Id:                  marketInfo.MarketId,
			Status:              marketInfo.GetMarketStatus().String(),
			Ticker:              marketInfo.GetTicker(),
			OracleSymbol:        marketInfo.OracleSymbol,
			OracleProvider:      marketInfo.OracleProvider,
			OracleType:          marketInfo.OracleType.String(),
			OracleScaleFactor:   marketInfo.GetOracleScaleFactor(),
			ExpirationTimestamp: marketInfo.ExpirationTimestamp,
			SettlementTimestamp: marketInfo.SettlementTimestamp,
			QuoteToken:          quoteToken,
			MakerFeeRate:        makerFeeRate,
			TakerFeeRate:        takerFeeRate,
			ServiceProviderFee:  serviceProviderFee,
			MinPriceTickSize:    minPriceTickSize,
			MinQuantityTickSize: minQuantityTickSize,
			MinNotional:         minNotional,
		}

		if marketInfo.SettlementPrice != nil {
			settlementPrice := decimal.RequireFromString(marketInfo.SettlementPrice.String())
			market.SettlementPrice = &settlementPrice
		}

		assistant.binaryOptionMarkets[market.Id] = market
	}

	return nil
}
