package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

	store, sqlErr := store.NewSQLStore(config)
	if sqlErr != nil {
		log.Fatalf("Failed to create SQL Store: %v", sqlErr)
	}

	server := api.NewServer(*listenAddr, store, config)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start a HTTP Server: %v", err)
	}

	defer store.DB.Close()

	if err = http.ListenAndServe(config.PORT, nil); err != nil {
		log.Fatalf("Failed to start a HTTP Server: %v", err)
	}

	fmt.Println("server running on:", *listenAddr)
}
