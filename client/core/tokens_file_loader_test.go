package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"
)

func TestLoadTokensFromUrl(t *testing.T) {
	tokensMetadata := make([]TokenMetadata, 2)
	tokensMetadata = append(tokensMetadata, TokenMetadata{
		Address:           "",
		IsNative:          true,
		Decimals:          9,
		Symbol:            "SOL",
		Name:              "Solana",
		Logo:              "https://imagedelivery.net/DYKOWp0iCc0sIkF-2e4dNw/2aa4deed-fa31-4d1a-ba0a-d698b84f3800/public",
		Creator:           "inj15jeczm4mqwtc9lk4c0cyynndud32mqd4m9xnmu",
		CoinGeckoId:       "solana",
		Denom:             "",
		TokenType:         "spl",
		TokenVerification: "verified",
		ExternalLogo:      "solana.png",
	},
	)
	tokensMetadata = append(tokensMetadata, TokenMetadata{
		Address:           "0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270",
		IsNative:          false,
		Decimals:          18,
		Symbol:            "WMATIC",
		Name:              "Wrapped Matic",
		Logo:              "https://imagedelivery.net/DYKOWp0iCc0sIkF-2e4dNw/0d061e1e-a746-4b19-1399-8187b8bb1700/public",
		Creator:           "inj169ed97mcnf8ay6rgvskn95n6tyt46uwvy5qgs0",
		CoinGeckoId:       "wmatic",
		Denom:             "0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270",
		TokenType:         "evm",
		TokenVerification: "verified",
		ExternalLogo:      "polygon.png",
	},
	)

	metadataString, err := json.Marshal(tokensMetadata)
	assert.NilError(t, err)

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(metadataString)
	}))
	defer httpServer.Close()

	loadedTokens, err := LoadTokens(httpServer.URL)
	assert.NilError(t, err)

	assert.Equal(t, len(loadedTokens), len(tokensMetadata))

	for i, metadata := range tokensMetadata {
		assert.Equal(t, loadedTokens[i], metadata)
	}
}

func TestLoadTokensFromUrlReturnsNoTokensWhenRequestErrorHappens(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer httpServer.Close()

	loadedTokens, err := LoadTokens(httpServer.URL)
	assert.Error(t, err, fmt.Sprintf("failed to load tokens from %s: %v %s", httpServer.URL, http.StatusNotFound, http.StatusText(http.StatusNotFound)))
	assert.Equal(t, len(loadedTokens), 0)
}
