package main

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"sort"
	"time"
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
	}

	ctx := context.Background()
	marketIds := []string{"0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"}
	stream, err := exchangeClient.StreamSpotOrderbookUpdate(ctx, marketIds)
	if err != nil {
		fmt.Println(err)
	}

	levelCh := make(chan *spotExchangePB.OrderbookLevel, 100000)

	// stream orderbook price levels
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				res, err := stream.Recv()
				if err != nil {
					fmt.Println(err)
					return
				}
				levelCh <- res.OrderbookLevel
			}
		}
	}()

	time.Sleep(5 * time.Second)

	// prepare orderbooks map
	orderbooks := map[string]*MapOrderbook{}
	res, err := exchangeClient.GetSpotOrderbooks(ctx, marketIds)
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
		level := <-levelCh

		// validate orderbook
		orderbook, ok := orderbooks[level.MarketId]
		if !ok {
			panic("level update doesn't belong to any orderbooks")
		}

		// skip if update sequence <= orderbook sequence until it's ready to consume
		if !skippedPastEvents {
			if orderbook.Sequence >= level.Sequence {
				continue
			}
			skippedPastEvents = true
		}

		// panic if update sequence > orderbook sequence + 1
		if level.Sequence > orderbook.Sequence+1 {
			fmt.Println("skipping", level.Sequence, orderbook.Sequence)
			panic("missing orderbook update events from stream, must restart")
		}

		// update orderbook map
		orderbook.Sequence = level.Sequence
		if level.IsActive {
			// upsert
			orderbook.Levels[level.IsBuy][level.Price] = &spotExchangePB.PriceLevel{
				Price:     level.Price,
				Quantity:  level.Quantity,
				Timestamp: level.UpdatedAt,
			}
		} else {
			// remove inactive level
			delete(orderbook.Levels[level.IsBuy], level.Price)
		}

		// construct orderbook arrays
		sells, buys := maintainOrderbook(orderbook.Levels)
		fmt.Println("after", orderbook.Sequence, len(sells), len(buys))
		sells = sells[:1]
		buys = buys[0:1]

		// assert orderbook
		if sells[0].Price <= buys[0].Price {
			fmt.Println(sells[0], buys[0])
			panic("crossed orderbook, must restart")
		}

		res, _ = exchangeClient.GetSpotOrderbooks(ctx, marketIds)
		fmt.Println("query", res.Orderbooks[0].Orderbook.Sequence, len(res.Orderbooks[0].Orderbook.Sells), len(res.Orderbooks[0].Orderbook.Buys))

		// print orderbook
		for _, s := range sells {
			fmt.Println(s)
		}
		fmt.Println("========")
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
		return sells[i].Price > sells[j].Price
	})
	sort.Slice(buys, func(i, j int) bool {
		return buys[i].Price > buys[j].Price
	})

	return sells, buys
}
