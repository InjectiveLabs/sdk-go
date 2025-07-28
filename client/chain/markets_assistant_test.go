package chain

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	exchangev1types "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/InjectiveLabs/sdk-go/client/core"
	"github.com/InjectiveLabs/sdk-go/client/exchange"
)

func TestMarketAssistantCreation(t *testing.T) {
	tokensList := []core.TokenMetadata{
		{
			Address:           "0x44C21afAaF20c270EBbF5914Cfc3b5022173FEB7",
			IsNative:          false,
			TokenVerification: "verified",
			Decimals:          6,
			CoinGeckoId:       "",
			Name:              "Ape",
			Symbol:            "APE",
			Logo:              "https://imagedelivery.net/lPzngbR8EltRfBOi_WYaXw/6f015260-c589-499f-b692-a57964af9900/public",
			Denom:             "peggy0x44C21afAaF20c270EBbF5914Cfc3b5022173FEB7",
			TokenType:         "erc20",
			ExternalLogo:      "unknown.png",
		},
		{
			Address:           "0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5",
			IsNative:          false,
			TokenVerification: "verified",
			Decimals:          6,
			Symbol:            "USDT",
			Name:              "Tether",
			Logo:              "https://imagedelivery.net/lPzngbR8EltRfBOi_WYaXw/a0bd252b-1005-47ef-d209-7c1c4a3cbf00/public",
			CoinGeckoId:       "tether",
			Denom:             "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5",
			TokenType:         "erc20",
			ExternalLogo:      "usdt.png",
		},
		{
			Decimals:          6,
			Symbol:            "USDT",
			Name:              "Other USDT",
			Logo:              "https://imagedelivery.net/lPzngbR8EltRfBOi_WYaXw/6f015260-c589-499f-b692-a57964af9900/public",
			CoinGeckoId:       "",
			Address:           "factory/inj10vkkttgxdeqcgeppu20x9qtyvuaxxev8qh0awq/usdt",
			Denom:             "factory/inj10vkkttgxdeqcgeppu20x9qtyvuaxxev8qh0awq/usdt",
			ExternalLogo:      "unknown.png",
			TokenType:         "tokenFactory",
			TokenVerification: "internal",
		},
		{
			Address:           "inj",
			IsNative:          true,
			TokenVerification: "verified",
			Decimals:          18,
			Symbol:            "INJ",
			Name:              "Injective",
			Logo:              "https://imagedelivery.net/lPzngbR8EltRfBOi_WYaXw/18984c0b-3e61-431d-241d-dfbb60b57600/public",
			CoinGeckoId:       "injective-protocol",
			Denom:             "inj",
			TokenType:         "native",
			ExternalLogo:      "injective-v3.png",
		},
		{
			Decimals:          6,
			Symbol:            "USDTPERP",
			Name:              "USDT PERP",
			Logo:              "https://static.alchemyapi.io/images/assets/825.png",
			CoinGeckoId:       "",
			Address:           "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			Denom:             "peggy0xdAC17F958D2ee523a2206206994597C13D831ec7",
			ExternalLogo:      "unknown.png",
			TokenType:         "tokenFactory",
			TokenVerification: "internal",
		},
	}

	marshalledTokensList, err := json.Marshal(tokensList)
	assert.NoError(t, err)

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(marshalledTokensList)
	}))
	defer httpServer.Close()

	network := common.NewNetwork()
	network.OfficialTokensListURL = httpServer.URL

	mockChain := MockChainClient{}
	mockChain.Network = network
	var spotMarketInfos []*exchangev1types.SpotMarket
	var fullDerivativeMarkets []*exchangev1types.FullDerivativeMarket
	var binaryOptionsMarkets []*exchangev1types.BinaryOptionsMarket
	injUsdtSpotMarketInfo := createINJUSDTChainSpotMarketV1()
	apeUsdtSpotMarketInfo := createAPEUSDTChainSpotMarketV1()
	btcUsdtDerivativeMarketInfo := createBTCUSDTChainDerivativeMarketV1()
	betBinaryOptionsMarket := createFirstMatchBetBinaryOptionsMarketV1()

	spotMarketInfos = append(spotMarketInfos, injUsdtSpotMarketInfo)
	spotMarketInfos = append(spotMarketInfos, apeUsdtSpotMarketInfo)
	fullDerivativeMarkets = append(fullDerivativeMarkets, &exchangev1types.FullDerivativeMarket{
		Market: btcUsdtDerivativeMarketInfo,
	})
	binaryOptionsMarkets = append(binaryOptionsMarkets, betBinaryOptionsMarket)

	mockChain.QuerySpotMarketsResponses = append(mockChain.QuerySpotMarketsResponses, &exchangev1types.QuerySpotMarketsResponse{
		Markets: spotMarketInfos,
	})
	mockChain.QueryDerivativeMarketsResponses = append(mockChain.QueryDerivativeMarketsResponses, &exchangev1types.QueryDerivativeMarketsResponse{
		Markets: fullDerivativeMarkets,
	})
	mockChain.QueryBinaryMarketsResponses = append(mockChain.QueryBinaryMarketsResponses, &exchangev1types.QueryBinaryMarketsResponse{
		Markets: binaryOptionsMarkets,
	})

	ctx := context.Background()
	assistant, err := NewMarketsAssistant(ctx, &mockChain)

	require.NoError(t, err)

	tokens := assistant.AllTokens()

	assert.Len(t, tokens, 5)

	symbols := strings.Split(injUsdtSpotMarketInfo.Ticker, "/")
	injSymbol, usdtSymbol := symbols[0], symbols[1]
	symbols = strings.Split(apeUsdtSpotMarketInfo.Ticker, "/")
	apeSymbol := symbols[0]
	alternativeUSDTName := "Other USDT"

	usdtPerpToken := tokens["USDTPERP"]
	usdtPerpSymbol := usdtPerpToken.Symbol

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

func TestMarketAssistantCreationWithAllTokens(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("[]"))
	}))
	defer httpServer.Close()

	network := common.NewNetwork()
	network.OfficialTokensListURL = httpServer.URL

	mockExchange := exchange.MockExchangeClient{}
	mockExchange.Network = network
	mockChain := MockChainClient{}
	smartDenomMetadata := createSmartDenomMetadata()

	mockChain.QuerySpotMarketsResponses = append(mockChain.QuerySpotMarketsResponses, &exchangev1types.QuerySpotMarketsResponse{})
	mockChain.QueryDerivativeMarketsResponses = append(mockChain.QueryDerivativeMarketsResponses, &exchangev1types.QueryDerivativeMarketsResponse{})
	mockChain.QueryBinaryMarketsResponses = append(mockChain.QueryBinaryMarketsResponses, &exchangev1types.QueryBinaryMarketsResponse{})

	mockChain.DenomsMetadataResponses = append(mockChain.DenomsMetadataResponses, &banktypes.QueryDenomsMetadataResponse{
		Metadatas: []banktypes.Metadata{smartDenomMetadata},
	})

	ctx := context.Background()
	assistant, err := NewMarketsAssistantWithAllTokens(ctx, &mockExchange, &mockChain)

	assert.NoError(t, err)

	tokens := assistant.AllTokens()

	assert.Len(t, tokens, 1)

	_, isPresent := tokens[smartDenomMetadata.Symbol]
	assert.True(t, isPresent)
}

func TestHumanReadableMarketAssistantCreation(t *testing.T) {
	tokensList := []core.TokenMetadata{
		{
			Address:           "0x44C21afAaF20c270EBbF5914Cfc3b5022173FEB7",
			IsNative:          false,
			TokenVerification: "verified",
			Decimals:          6,
			CoinGeckoId:       "",
			Name:              "Ape",
			Symbol:            "APE",
			Logo:              "https://imagedelivery.net/lPzngbR8EltRfBOi_WYaXw/6f015260-c589-499f-b692-a57964af9900/public",
			Denom:             "peggy0x44C21afAaF20c270EBbF5914Cfc3b5022173FEB7",
			TokenType:         "erc20",
			ExternalLogo:      "unknown.png",
		},
		{
			Address:           "0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5",
			IsNative:          false,
			TokenVerification: "verified",
			Decimals:          6,
			Symbol:            "USDT",
			Name:              "Tether",
			Logo:              "https://imagedelivery.net/lPzngbR8EltRfBOi_WYaXw/a0bd252b-1005-47ef-d209-7c1c4a3cbf00/public",
			CoinGeckoId:       "tether",
			Denom:             "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5",
			TokenType:         "erc20",
			ExternalLogo:      "usdt.png",
		},
		{
			Decimals:          6,
			Symbol:            "USDT",
			Name:              "Other USDT",
			Logo:              "https://imagedelivery.net/lPzngbR8EltRfBOi_WYaXw/6f015260-c589-499f-b692-a57964af9900/public",
			CoinGeckoId:       "",
			Address:           "factory/inj10vkkttgxdeqcgeppu20x9qtyvuaxxev8qh0awq/usdt",
			Denom:             "factory/inj10vkkttgxdeqcgeppu20x9qtyvuaxxev8qh0awq/usdt",
			ExternalLogo:      "unknown.png",
			TokenType:         "tokenFactory",
			TokenVerification: "internal",
		},
		{
			Address:           "inj",
			IsNative:          true,
			TokenVerification: "verified",
			Decimals:          18,
			Symbol:            "INJ",
			Name:              "Injective",
			Logo:              "https://imagedelivery.net/lPzngbR8EltRfBOi_WYaXw/18984c0b-3e61-431d-241d-dfbb60b57600/public",
			CoinGeckoId:       "injective-protocol",
			Denom:             "inj",
			TokenType:         "native",
			ExternalLogo:      "injective-v3.png",
		},
		{
			Decimals:          6,
			Symbol:            "USDTPERP",
			Name:              "USDT PERP",
			Logo:              "https://static.alchemyapi.io/images/assets/825.png",
			CoinGeckoId:       "",
			Address:           "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			Denom:             "peggy0xdAC17F958D2ee523a2206206994597C13D831ec7",
			ExternalLogo:      "unknown.png",
			TokenType:         "tokenFactory",
			TokenVerification: "internal",
		},
	}

	marshalledTokensList, err := json.Marshal(tokensList)
	assert.NoError(t, err)

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(marshalledTokensList)
	}))
	defer httpServer.Close()

	network := common.NewNetwork()
	network.OfficialTokensListURL = httpServer.URL

	mockChain := MockChainClientV2{}
	mockChain.Network = network
	var spotMarketInfos []*exchangev2types.SpotMarket
	var fullDerivativeMarkets []*exchangev2types.FullDerivativeMarket
	var binaryOptionsMarkets []*exchangev2types.BinaryOptionsMarket
	injUsdtSpotMarketInfo := createINJUSDTChainSpotMarketV2()
	apeUsdtSpotMarketInfo := createAPEUSDTChainSpotMarketV2()
	btcUsdtDerivativeMarketInfo := createBTCUSDTChainDerivativeMarketV2()
	betBinaryOptionsMarket := createFirstMatchBetBinaryOptionsMarketV2()

	spotMarketInfos = append(spotMarketInfos, injUsdtSpotMarketInfo)
	spotMarketInfos = append(spotMarketInfos, apeUsdtSpotMarketInfo)
	fullDerivativeMarkets = append(fullDerivativeMarkets, &exchangev2types.FullDerivativeMarket{
		Market: btcUsdtDerivativeMarketInfo,
	})
	binaryOptionsMarkets = append(binaryOptionsMarkets, betBinaryOptionsMarket)

	mockChain.QuerySpotMarketsV2Responses = append(mockChain.QuerySpotMarketsV2Responses, &exchangev2types.QuerySpotMarketsResponse{
		Markets: spotMarketInfos,
	})
	mockChain.QueryDerivativeMarketsV2Responses = append(mockChain.QueryDerivativeMarketsV2Responses, &exchangev2types.QueryDerivativeMarketsResponse{
		Markets: fullDerivativeMarkets,
	})
	mockChain.QueryBinaryMarketsV2Responses = append(mockChain.QueryBinaryMarketsV2Responses, &exchangev2types.QueryBinaryMarketsResponse{
		Markets: binaryOptionsMarkets,
	})

	ctx := context.Background()
	assistant, err := NewHumanReadableMarketsAssistant(ctx, &mockChain)

	require.NoError(t, err)

	tokens := assistant.AllTokens()

	assert.Len(t, tokens, 5)

	symbols := strings.Split(injUsdtSpotMarketInfo.Ticker, "/")
	injSymbol, usdtSymbol := symbols[0], symbols[1]
	symbols = strings.Split(apeUsdtSpotMarketInfo.Ticker, "/")
	apeSymbol := symbols[0]
	alternativeUSDTName := "Other USDT"

	usdtPerpToken := tokens["USDTPERP"]
	usdtPerpSymbol := usdtPerpToken.Symbol

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

func TestHumanReadableMarketAssistantCreationWithAllTokens(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("[]"))
	}))
	defer httpServer.Close()

	network := common.NewNetwork()
	network.OfficialTokensListURL = httpServer.URL

	mockExchange := exchange.MockExchangeClient{}
	mockExchange.Network = network
	mockChain := MockChainClientV2{}
	smartDenomMetadata := createSmartDenomMetadata()

	mockChain.QuerySpotMarketsV2Responses = append(mockChain.QuerySpotMarketsV2Responses, &exchangev2types.QuerySpotMarketsResponse{})
	mockChain.QueryDerivativeMarketsV2Responses = append(mockChain.QueryDerivativeMarketsV2Responses, &exchangev2types.QueryDerivativeMarketsResponse{})
	mockChain.QueryBinaryMarketsV2Responses = append(mockChain.QueryBinaryMarketsV2Responses, &exchangev2types.QueryBinaryMarketsResponse{})

	mockChain.DenomsMetadataResponses = append(mockChain.DenomsMetadataResponses, &banktypes.QueryDenomsMetadataResponse{
		Metadatas: []banktypes.Metadata{smartDenomMetadata},
	})

	ctx := context.Background()
	assistant, err := NewHumanReadableMarketsAssistantWithAllTokens(ctx, &mockExchange, &mockChain)

	assert.NoError(t, err)

	tokens := assistant.AllTokens()

	assert.Len(t, tokens, 1)

	_, isPresent := tokens[smartDenomMetadata.Symbol]
	assert.True(t, isPresent)
}
