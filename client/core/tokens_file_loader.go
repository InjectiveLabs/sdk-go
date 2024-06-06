package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TokenMetadata struct {
	Address           string `json:"address"`
	IsNative          bool   `json:"isNative"`
	TokenVerification string `json:"tokenVerification"`
	Decimals          int32  `json:"decimals"`
	CoinGeckoId       string `json:"coinGeckoId"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	Logo              string `json:"logo"`
	Creator           string `json:"creator"`
	Denom             string `json:"denom"`
	TokenType         string `json:"tokenType"`
	ExternalLogo      string `json:"externalLogo"`
}

func (tm TokenMetadata) GetName() string {
	return tm.Name
}

func (tm TokenMetadata) GetAddress() string {
	return tm.Address
}

func (tm TokenMetadata) GetSymbol() string {
	return tm.Symbol
}

func (tm TokenMetadata) GetLogo() string {
	return tm.Logo
}

func (tm TokenMetadata) GetDecimals() int32 {
	return tm.Decimals
}

func (tm TokenMetadata) GetUpdatedAt() int64 {
	return -1
}

// LoadTokens loads tokens from the given file URL
func LoadTokens(tokensFileUrl string) ([]TokenMetadata, error) {
	var tokensMetadata []TokenMetadata
	response, err := http.Get(tokensFileUrl)
	if err != nil {
		return tokensMetadata, err
	}
	if 400 <= response.StatusCode {
		return tokensMetadata, fmt.Errorf("failed to load tokens from %s: %s", tokensFileUrl, response.Status)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&tokensMetadata)
	if err != nil {
		return make([]TokenMetadata, 0), err
	}

	return tokensMetadata, nil
}
