package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"tournament_api/server/api"
	"tournament_api/server/store"
)

func main() {

	listenAddr := flag.String("listenaddr", ":3000", "the server address")

	store := store.NewSQLStore()
	server := api.NewServer(*listenAddr, store)
	fmt.Println("server running on:", *listenAddr)
	log.Fatal(server.Start())

	port := ":8080"
	http.ListenAndServe(port, nil)
}
