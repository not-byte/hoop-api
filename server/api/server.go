package api

import (
	"go/types"
	"net/http"
	"tournament_api/server/store"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	store      store.Store
	config     *types.Config
}

func NewServer(listenAddr string, store store.Store, config *types.Config) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
		config:     config,
	}
}

func (s *Server) Start() error {
	rootRouter := mux.NewRouter()

	authRouter := rootRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("", s.handleGetAll)
	authRouter.HandleFunc("/login", s.handleLogin)

	http.Handle("/", rootRouter)

	return http.ListenAndServe(s.listenAddr, nil)
}
