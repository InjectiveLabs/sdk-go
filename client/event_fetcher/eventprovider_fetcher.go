package event_fetcher

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"cosmossdk.io/errors"
	"github.com/InjectiveLabs/sdk-go/client/common"
	eventproviderPB "github.com/InjectiveLabs/sdk-go/exchange/event_provider_rpc/pb"
	"github.com/InjectiveLabs/suplog"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type eventproviderFetcher struct {
	ep          eventproviderPB.EventProviderAPIClient
	filters     FilterMap
	filterTypes []string
}

func dialConnection(eventproviderURL string) (*grpc.ClientConn, error) {
	// ignore cert verification if server has tls
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	conn, err := grpc.DialContext(
		ctx,
		eventproviderURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(common.DialerFunc),
		grpc.WithBlock(),
	)
	if err != nil {
		suplog.WithError(err).Warningln("cannot connect eventprovider with non-TLS mode. Trying TLS mode now...")

		ctx2, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFn()
		tlsConfig := tls.Config{InsecureSkipVerify: true}
		creds := credentials.NewTLS(&tlsConfig)
		conn, err = grpc.DialContext(
			ctx2,
			eventproviderURL,
			grpc.WithTransportCredentials(creds),
			grpc.WithContextDialer(common.DialerFunc),
			grpc.WithBlock(),
		)
	}

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewEventProviderFetcher(epAddr string, filters FilterMap) (*eventproviderFetcher, error) {
	// TODO: Support connect with both tls and non-tls
	// create grpc client
	conn, err := dialConnection(epAddr)
	if err != nil {
		err := fmt.Errorf("failed to connect to the gRPC '%s': %w", epAddr, err)
		return nil, err
	}

	return &eventproviderFetcher{
		ep:          eventproviderPB.NewEventProviderAPIClient(conn),
		filters:     filters,
		filterTypes: maps.Keys(filters),
	}, nil
}

func (ef *eventproviderFetcher) Fetch(ctx context.Context, height int64) ([]Event, error) {
	resp, err := ef.ep.GetABCIBlockEvents(ctx, &eventproviderPB.GetABCIBlockEventsRequest{
		Height:     int32(height),
		EventTypes: ef.filterTypes,
	})
	if err != nil {
		return nil, err
	}

	if resp.RawBlock == nil {
		return nil, nil
	}

	result := parseEventsToContainers(resp.RawBlock.BeginBlockEvents, ef.filters)
	for _, txEvents := range resp.RawBlock.TxsResults {
		result = append(result, parseEventsToContainers(txEvents.Events, ef.filters)...)
	}

	result = append(result, parseEventsToContainers(resp.RawBlock.EndBlockEvents, ef.filters)...)
	return result, nil
}

func parseEventsToContainers(evList []*eventproviderPB.ABCIEvent, filters FilterMap) (result []Event) {
	for _, e := range evList {
		someEvent, exist := filters[e.Type]
		if !exist {
			continue
		}

		container := reflect.New(reflect.TypeOf(someEvent).Elem()).Interface().(Event)
		err := eventProviderEventToContainer(e, container)
		if err != nil {
			suplog.WithError(err).Warningf("cannot parse event: %s", e.Type)
			continue
		}
		result = append(result, container)
	}
	return result
}

func eventProviderEventToContainer(raw *eventproviderPB.ABCIEvent, event Event) error {
	evFieldsMap := make(map[string]json.RawMessage)
	for _, attr := range raw.Attributes {
		// if use cometbft to query injective-core v1.10, it will results in base64-encoded string,
		// and we need to parse base64
		key, value := attr.Key, attr.Value
		if json.Valid([]byte(value)) {
			evFieldsMap[key] = json.RawMessage(value)
		} else {
			// convert to string for json marshal
			evFieldsMap[key] = json.RawMessage(fmt.Sprintf(`"%s"`, value))
		}
	}

	// TODO: (phuc): Optimize this marshal - unmarshal
	data, err := json.Marshal(&evFieldsMap)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal temporary fields map before unpacking event")
		return err
	}

	if err := event.Unmarshal("1", data); err != nil {
		err = errors.Wrap(err, fmt.Sprintf("failed to unpack event from comet log attributes: %s", string(data)))
		return err
	}
	return nil
}
