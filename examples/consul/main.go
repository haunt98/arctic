package main

import (
	"arctic"
	"log"
	"os"
	"time"

	"github.com/hashicorp/consul/api"
)

func main() {
	address := os.Getenv("ADDRESS")
	if address == "" {
		log.Fatalf("empty ADDRESS")
	}

	config := &api.Config{
		Address: address,
	}

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("failed to new client: %s", err)
	}

	prefix := os.Getenv("PREFIX")
	if prefix == "" {
		log.Fatalf("empty PREFIX")
	}

	store := arctic.NewStore()
	arctic, err := arctic.NewConsulArctic(client, config, store, prefix)
	if err != nil {
		log.Fatalf("failed to new consul artic: %s", err)
	}

	for {
		value := arctic.Get(prefix + "/name")
		log.Printf("value %s", value)

		time.Sleep(time.Second)
	}
}
