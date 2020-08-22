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
	assert.True(address != "", "empty address")

	config := &api.Config{
		Address: address,
	}

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}

	prefix := viper.GetString("prefix")
	assert.True(prefix != "", "empty prefix")

	arc, err := arctic.NewConsulArctic(client, config, arctic.NewStore(), prefix)
	if err != nil {
		log.Fatalf("failed to new consul artic: %s", err)
	}

	key := viper.GetString("key")
	assert.True(key != "", "empty key")

	for {
		value := arc.Get(key)
		log.Printf("value %s\n", value)

		time.Sleep(time.Second)
	}
}
