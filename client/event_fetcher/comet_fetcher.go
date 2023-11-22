package event_fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"cosmossdk.io/errors"
	tmclient "github.com/InjectiveLabs/sdk-go/client/tm"
	"github.com/InjectiveLabs/suplog"
	abci "github.com/cometbft/cometbft/abci/types"
)

type cometFetcher struct {
	cometClient tmclient.TendermintClient
	filters     FilterMap
}

func NewCometFetcher(url string, filterMap FilterMap) *cometFetcher {
	// TODO: replace this with GRPC when cometbft released
	c := tmclient.NewRPCClient(url)
	return &cometFetcher{
		cometClient: c,
		filters:     filterMap,
	}
}

func (cm *cometFetcher) Fetch(ctx context.Context, height int64) ([]Event, error) {
	// TODO: Also implement fetching for explorer tx event
	blockResults, err := cm.cometClient.GetBlockResults(ctx, height)
	if err != nil {
		return nil, err
	}

	// TODO(phuc): implement event filtering on cometbft side which enable us to to even faster fetch
	events := filterEvents(blockResults.BeginBlockEvents, cm.filters)
	for _, txResult := range blockResults.TxsResults {
		events = append(events, filterEvents(txResult.Events, cm.filters)...)
	}

	events = append(events, filterEvents(blockResults.EndBlockEvents, cm.filters)...)
	return nil, nil
}

func filterEvents(evList []abci.Event, filters FilterMap) (result []Event) {
	for _, e := range evList {
		someEvent, exist := filters[e.Type]
		if !exist {
			continue
		}

		container := reflect.New(reflect.TypeOf(someEvent).Elem()).Interface().(Event)
		err := toContainer(e, container)
		if err != nil {
			suplog.WithError(err).Warningf("cannot parse event: %s", e.Type)
			continue
		}
		result = append(result, container)
	}
	return result
}

func (cm *cometFetcher) LastHeight(ctx context.Context) (int64, error) {
	return cm.cometClient.GetLatestBlockHeight(ctx)
}

func toContainer(raw abci.Event, event Event) error {
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
