package client

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

type ContractSet struct {
	FuturesContract     common.Address
	PriceFeederContract common.Address
}

type ContractDiscoverer interface {
	GetContractSet(ctx context.Context) ContractSet
}

func NewContractDiscoverer(ordersQuerier exchangetypes.QueryClient) ContractDiscoverer {
	return &contractDiscoverer{
		queryClient: ordersQuerier,
		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "contractDiscoverer",
		}),
	}
}

type contractDiscoverer struct {
	queryClient exchangetypes.QueryClient
	logger      log.Logger
}

const defaultRetryDelay = 10 * time.Second

func (c *contractDiscoverer) GetContractSet(ctx context.Context) (set ContractSet) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			resp, err := c.queryClient.QueryDerivativeMarkets(ctx, &exchangetypes.QueryDerivativeMarketsRequest{})
			if err != nil {
				c.logger.WithError(err).Warningln("failed to query derivative markets, retry in 10s")
				time.Sleep(defaultRetryDelay)
				continue
			} else if resp == nil || len(resp.Markets) == 0 {
				c.logger.Warningln("no derivative markets returned, retry in 10s")
				time.Sleep(defaultRetryDelay)
				continue
			}

			for _, market := range resp.Markets {
				set.FuturesContract = common.HexToAddress(market.ExchangeAddress)
				set.PriceFeederContract = common.HexToAddress(market.Oracle)
				if set.FuturesContract != (common.Address{}) &&
					set.PriceFeederContract != (common.Address{}) {
					return
				}
			}

			if set.FuturesContract == (common.Address{}) ||
				set.PriceFeederContract == (common.Address{}) {
				c.logger.WithFields(log.Fields{
					"futures_contract":      set.FuturesContract.Hex(),
					"price_feeder_contract": set.PriceFeederContract.Hex(),
				}).Warningln("could not discover some market contracts, retry in 10s")
				continue
			}
		}
	}
}
