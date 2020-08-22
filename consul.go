package arctic

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

var _ Arctic = (*consulArctic)(nil)

type consulArctic struct {
	client *api.Client
	config *api.Config
	store  *Store
	prefix string
}

func NewConsulArctic(
	client *api.Client,
	config *api.Config,
	store *Store,
	prefix string,
) (Arctic, error) {
	if prefix == "" {
		return nil, fmt.Errorf("empty prefix")
	}

	a := &consulArctic{
		client: client,
		config: config,
		store:  store,
		prefix: prefix,
	}

	kvPairs, _, err := a.client.KV().List(prefix, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	for _, kvPair := range kvPairs {
		a.store.put(kvPair.Key, kvPair.Value)
	}

	if err := a.watch(); err != nil {
		return nil, fmt.Errorf("failed to watch: %w", err)
	}

	return a, nil
}

func (a *consulArctic) Get(key string) []byte {
	key = ComposeKey(a.prefix, key)

	return a.store.get(key)
}

func (a *consulArctic) watch() error {
	params := map[string]interface{}{
		"type":   "keyprefix",
		"prefix": a.prefix,
	}

	plan, err := watch.Parse(params)
	if err != nil {
		return fmt.Errorf("failed to parse params %+v: %w", params, err)
	}

	plan.Handler = func(_ uint64, raw interface{}) {
		a.store.clear()

		if raw == nil {
			return
		}

		kvPairs, ok := raw.(api.KVPairs)
		if !ok {
			return
		}

		for _, kvPair := range kvPairs {
			a.store.put(kvPair.Key, kvPair.Value)
		}
	}

	go func(plan *watch.Plan, config *api.Config) {
		if err := plan.Run(config.Address); err != nil {
			log.Fatalf("failed to run plan %s: %s\n", config.Address, err)
		}
	}(plan, a.config)

	return nil
}
