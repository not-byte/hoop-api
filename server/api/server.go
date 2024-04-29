package api

import (
	"net/http"
	"tournament_api/server/store"
	"tournament_api/server/types"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	store      store.Store
	config     *types.AppConfig
}

func NewServer(listenAddr string, store store.Store, config *types.AppConfig) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
		config:     config,
	}
}

func (s *Server) Start() error {
	rootRouter := mux.NewRouter()
	if s.config.PRODUCTION {
		rootRouter.Use(s.APIKeyMiddleware)
	} else {
		rootRouter.Use(s.CORSmiddleware)
	}
	rootRouter.HandleFunc("/", s.handleGetAll)

	authRouter := rootRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("", s.handleGetAll)
	authRouter.HandleFunc("/login", s.handleLogin)
	authRouter.HandleFunc("/register", s.handleRegister)

	http.Handle("/", rootRouter)

	return http.ListenAndServe(s.listenAddr, nil)
}
