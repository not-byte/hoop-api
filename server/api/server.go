package api

import (
	"net/http"
	"tournament_api/server/store"
)

type Server struct {
	listenAddr string
	store      store.Store
}

func NewServer(listenAddr string, store store.Store) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.handleGetAll)
	return http.ListenAndServe(s.listenAddr, nil)
}
