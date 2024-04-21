package main

import (
	"flag"
	"fmt"
	"log"
	"tournament_api/server/api"
	"tournament_api/server/config"
	"tournament_api/server/store"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	listenAddr := flag.String("listenaddr", config.PORT, "the server address")

	store, err := store.NewSQLStore(config)
	if err != nil {
		log.Fatalf("Failed to create SQL Store: %v", err)
	}

	server := api.NewServer(*listenAddr, store, config)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start a HTTP Server: %v", err)
	}

	defer store.DB.Close()

	fmt.Println("server running on:", *listenAddr)
}
