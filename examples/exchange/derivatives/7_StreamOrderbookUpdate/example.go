package main

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	"sort"
)

type MapOrderbook struct {
	Sequence uint64
	Levels   map[bool]map[string]*derivativeExchangePB.PriceLevel
}

func main() {
	network := common.LoadNetwork("devnet-1", "")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ctx := context.Background()
	marketIds := []string{"0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"}
	stream, err := exchangeClient.StreamDerivativeOrderbookUpdate(ctx, marketIds)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	levelCh := make(chan *derivativeExchangePB.OrderbookLevel, 100000)
	receiving := make(chan struct{})
	var receivingClosed bool

	// stream orderbook price levels
	go func() {
		defer close(levelCh)
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
				if !receivingClosed {
					close(receiving)
					receivingClosed = true
				}
			}
		}
	}()

	// ensure we are receiving updates before getting orderbook
	fmt.Println("waiting for streaming updates")
	<-receiving

	// prepare orderbooks map
	orderbooks := map[string]*MapOrderbook{}
	res, err := exchangeClient.GetDerivativeOrderbooks(ctx, marketIds)
	if err != nil {
		panic(err)
	}
	for _, ob := range res.Orderbooks {
		// init inner maps not ready
		_, ok := orderbooks[ob.MarketId]
		if !ok {
			orderbook := &MapOrderbook{
				Sequence: ob.Orderbook.Sequence,
				Levels:   make(map[bool]map[string]*derivativeExchangePB.PriceLevel),
			}
			orderbook.Levels[true] = make(map[string]*derivativeExchangePB.PriceLevel)
			orderbook.Levels[false] = make(map[string]*derivativeExchangePB.PriceLevel)
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
		level, ok := <-levelCh
		if !ok {
			fmt.Println("updates channel closed")
			return // closed
		}

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
			orderbook.Levels[level.IsBuy][level.Price] = &derivativeExchangePB.PriceLevel{
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

		if len(sells) > 0 && len(buys) > 0 {
			// assert orderbook
			if sellPrice, buyPrice := sells[0].Price, buys[0].Price; sellPrice <= buyPrice {
				panic(fmt.Errorf("crossed orderbook, must restart: buy %s > %s sell", buyPrice, sellPrice))
			}
		}

		res, _ = exchangeClient.GetDerivativeOrderbooks(ctx, marketIds)
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

func maintainOrderbook(orderbook map[bool]map[string]*derivativeExchangePB.PriceLevel) (buys, sells []*derivativeExchangePB.PriceLevel) {
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
