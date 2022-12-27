package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"github.com/shopspring/decimal"
)

type MapOrderbook struct {
	Sequence uint64
	Levels   map[bool]map[string]*spotExchangePB.PriceLevel
}

func main() {
	network := common.LoadNetwork("devnet-1", "")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ctx := context.Background()
	marketIds := []string{"0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"}
	stream, err := exchangeClient.StreamSpotOrderbookUpdate(ctx, marketIds)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	updatesCh := make(chan *spotExchangePB.OrderbookLevelUpdates, 100000)
	receiving := make(chan struct{})
	var receivingClosed bool

	// stream orderbook price levels
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				res, err := stream.Recv()
				if err == io.EOF {
					fmt.Println("change stream needs to be restarted")
					os.Exit(0)
				}
				if err != nil {
					panic(err) // rester on errors
				}
				u := res.OrderbookLevelUpdates
				if !receivingClosed {
					fmt.Println("receiving updates from stream")
					close(receiving)
					receivingClosed = true
				}
				updatesCh <- u
			}
		}
	}()

	// ensure we are receiving updates before getting the orderbook snapshot,
	// we will skip past updates later.
	fmt.Println("waiting for streaming updates")
	<-receiving

	// prepare orderbooks map
	orderbooks := map[string]*MapOrderbook{}
	res, err := exchangeClient.GetSpotOrderbooksV2(ctx, marketIds)
	if err != nil {
		panic(err)
	}
	for _, ob := range res.Orderbooks {
		// init inner maps not ready
		_, ok := orderbooks[ob.MarketId]
		if !ok {
			orderbook := &MapOrderbook{
				Sequence: ob.Orderbook.Sequence,
				Levels:   map[bool]map[string]*spotExchangePB.PriceLevel{},
			}
			orderbook.Levels[true] = map[string]*spotExchangePB.PriceLevel{}
			orderbook.Levels[false] = map[string]*spotExchangePB.PriceLevel{}
			orderbooks[ob.MarketId] = orderbook
		}

		for _, level := range ob.Orderbook.Buys {
			orderbooks[ob.MarketId].Levels[true][level.Price] = level
		}
		for _, level := range ob.Orderbook.Sells {
			orderbooks[ob.MarketId].Levels[false][level.Price] = level
		}
	}

	// continuously consume level updates and maintain orderbook
	skippedPastEvents := false
	for {
		updates := <-updatesCh

		// validate orderbook
		orderbook, ok := orderbooks[updates.MarketId]
		if !ok {
			fmt.Println("updates channel closed, must restart")
			return // closed
		}

		// skip if update sequence <= orderbook sequence until it's ready to consume
		if !skippedPastEvents {
			if orderbook.Sequence >= updates.Sequence {
				continue
			}
			skippedPastEvents = true
		}

		// panic if update sequence > orderbook sequence + 1
		if updates.Sequence > orderbook.Sequence+1 {
			fmt.Printf("skipping levels: update sequence %d vs orderbook sequence %d\n", updates.Sequence, orderbook.Sequence)
			panic("missing orderbook update events from stream, must restart")
		}

		// update orderbook map
		orderbook.Sequence = updates.Sequence
		for isBuy, update := range map[bool][]*spotExchangePB.PriceLevelUpdate{
			true:  updates.Buys,
			false: updates.Sells,
		} {
			for _, level := range update {
				if level.IsActive {
					// upsert
					orderbook.Levels[isBuy][level.Price] = &spotExchangePB.PriceLevel{
						Price:     level.Price,
						Quantity:  level.Quantity,
						Timestamp: level.Timestamp,
					}
				} else {
					// remove inactive level
					delete(orderbook.Levels[isBuy], level.Price)
				}
			}
		}

		// construct orderbook arrays
		sells, buys := maintainOrderbook(orderbook.Levels)
		fmt.Println("after", orderbook.Sequence, len(sells), len(buys))

		if len(sells) > 0 && len(buys) > 0 {
			// assert orderbook
			topBuyPrice := decimal.RequireFromString(buys[0].Price)
			lowestSellPrice := decimal.RequireFromString(sells[0].Price)
			if topBuyPrice.GreaterThanOrEqual(lowestSellPrice) {
				panic(fmt.Errorf("crossed orderbook, must restart: buy %s >= %s sell", topBuyPrice, lowestSellPrice))
			}
		}

		res, _ = exchangeClient.GetSpotOrderbooksV2(ctx, marketIds)
		fmt.Println("query", res.Orderbooks[0].Orderbook.Sequence, len(res.Orderbooks[0].Orderbook.Sells), len(res.Orderbooks[0].Orderbook.Buys))

		// print orderbook
		fmt.Println(" [SELLS] ========")
		for _, s := range sells {
			fmt.Println(s)
		}
		fmt.Println(" [BUYS] ========")
		for _, b := range buys {
			fmt.Println(b)
		}
		fmt.Println("=======================================================")
	}
}

func maintainOrderbook(orderbook map[bool]map[string]*spotExchangePB.PriceLevel) (buys, sells []*spotExchangePB.PriceLevel) {
	for _, v := range orderbook[false] {
		sells = append(sells, v)
	}
	for _, v := range orderbook[true] {
		buys = append(buys, v)
	}

	sort.Slice(sells, func(i, j int) bool {
		return decimal.RequireFromString(sells[i].Price).LessThan(decimal.RequireFromString(sells[j].Price))
	})

	sort.Slice(buys, func(i, j int) bool {
		return decimal.RequireFromString(buys[i].Price).GreaterThan(decimal.RequireFromString(buys[j].Price))
	})

	return sells, buys
}
