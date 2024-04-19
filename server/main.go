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

	store, sql_err := store.NewSQLStore(config)
	if sql_err != nil {
		log.Fatalf("Failed to create store: %v", sql_err)
	}

	server := api.NewServer(*listenAddr, store, config)

	fmt.Println("server running on:", *listenAddr)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	defer store.DB.Close()

	http.ListenAndServe(config.PORT, nil)
}
