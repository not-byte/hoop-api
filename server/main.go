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
	port := ":8080"
	listenAddr := flag.String("listenaddr", port, "the server address")
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	store := store.NewSQLStore(config)
	server := api.NewServer(*listenAddr, store, config)

	fmt.Println("server running on:", *listenAddr)

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	http.ListenAndServe(port, nil)
}
