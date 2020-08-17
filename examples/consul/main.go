package main

import (
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/haunt98/arctic"
	"github.com/haunt98/assert"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	address := viper.GetString("address")
	assert.Bool(true, address != "", "empty address")

	config := &api.Config{
		Address: address,
	}

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to new client: %s", err)
	}

	prefix := viper.GetString("prefix")
	assert.Bool(true, prefix != "", "empty prefix")

	arc, err := arctic.NewConsulArctic(client, config, arctic.NewStore(), prefix)
	if err != nil {
		log.Fatalf("Failed to new consul artic: %s", err)
	}

	key := viper.GetString("key")
	assert.Bool(true, key != "", "empty key")

	for {
		value := arc.Get(arctic.ComposeKey(prefix, key))
		log.Printf("Value %s\n", value)

		time.Sleep(time.Second)
	}
}
