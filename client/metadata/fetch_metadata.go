package main

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	"os"

	//derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"math"
	"strconv"
)

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

var metadataTemplate = `[%s]
description = '%s %s %s'
base = %d
quote = %d
min_price_tick_size = %.18f
min_display_price_tick_size = %.4f
min_quantity_tick_size = %f
min_display_quantity_tick_size = %.4f

`
var symbolTemplate = `[%s]
peggy_denom = %s
decimals = %s

`

func FetchDenom(network common.Network) {
	metadataOutput := ""
	symbols := make(map[string][]string)
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	// fetch spot markets
	spotMarketsReq := spotExchangePB.MarketsRequest{MarketStatus: "active"}
	ctx := context.Background()
	spotRes, err := exchangeClient.GetSpotMarkets(ctx, spotMarketsReq)
	if err != nil {
		panic(err)
	}
	for _, m := range spotRes.Markets {
		// skip markets that don't have enough metadata
		if m.BaseTokenMeta == nil || m.QuoteTokenMeta == nil {
			continue
		}
		// append symbols to map
		symbols[m.BaseTokenMeta.Symbol] = []string{m.BaseDenom, fmt.Sprintf("%d", m.BaseTokenMeta.Decimals)}
		symbols[m.QuoteTokenMeta.Symbol] = []string{m.BaseDenom, fmt.Sprintf("%d", m.QuoteTokenMeta.Decimals)}

		// format market metadata into ini entry
		minPriceTickSize, err := strconv.ParseFloat(m.MinPriceTickSize, 64)
		if err != nil {
			panic(err)
		}
		minQuantityTickSize, err := strconv.ParseFloat(m.MinQuantityTickSize, 64)
		if err != nil {
			panic(err)
		}
		minDisplayPriceTickSize := minPriceTickSize / math.Pow(10, float64(m.QuoteTokenMeta.Decimals-m.BaseTokenMeta.Decimals))
		minDisplayQuantityTickSize := minQuantityTickSize / math.Pow(10, float64(m.BaseTokenMeta.Decimals))
		config := fmt.Sprintf(
			metadataTemplate,
			m.MarketId,
			network.Name, "Spot", m.Ticker,
			m.BaseTokenMeta.Decimals,
			m.QuoteTokenMeta.Decimals,
			minPriceTickSize,
			minDisplayPriceTickSize,
			minQuantityTickSize,
			minDisplayQuantityTickSize,
		)
		metadataOutput += config
	}

	//fetch derivative markets
	derivativeMarketsReq := derivativeExchangePB.MarketsRequest{MarketStatus: "active"}
	derivativeRes, err := exchangeClient.GetDerivativeMarkets(ctx, derivativeMarketsReq)
	if err != nil {
		panic(err)
	}
	for _, m := range derivativeRes.Markets {
		// skip markets that don't have quote metadata
		if m.QuoteTokenMeta == nil {
			continue
		}
		// append symbols to map
		symbols[m.QuoteTokenMeta.Symbol] = []string{m.QuoteDenom, string(m.QuoteTokenMeta.Decimals)}
		// format market metadata into ini entry
		minPriceTickSize, err := strconv.ParseFloat(m.MinPriceTickSize, 64)
		if err != nil {
			panic(err)
		}
		minQuantityTickSize, err := strconv.ParseFloat(m.MinQuantityTickSize, 64)
		if err != nil {
			panic(err)
		}
		minDisplayPriceTickSize := minPriceTickSize / math.Pow(10, float64(m.QuoteTokenMeta.Decimals))
		config := fmt.Sprintf(
			metadataTemplate,
			m.MarketId,
			network.Name, "Derivative", m.Ticker,
			0,
			m.QuoteTokenMeta.Decimals,
			minPriceTickSize,
			minDisplayPriceTickSize,
			minQuantityTickSize,
			minQuantityTickSize,
		)
		metadataOutput += config
	}

	// format into ini entry
	for k, v := range symbols {
		symbol := fmt.Sprintf(
			symbolTemplate,
			k, v[0], v[1],
		)
		metadataOutput += symbol
	}

	fileName := fmt.Sprintf("client/metadata/assets/%s.ini", network.Name)
	err = os.WriteFile(fileName, []byte(metadataOutput), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	devnet := common.LoadNetwork("devnet", "")
	testnet := common.LoadNetwork("testnet", "lb")
	mainnet := common.LoadNetwork("mainnet", "lb")
	FetchDenom(devnet)
	FetchDenom(testnet)
	FetchDenom(mainnet)
}
