package api

import (
	"net/http"
	"fmt"
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
	rootRoute := fmt.Sprintf("/v%d", s.config.VERSION)
	rootRouter := mux.NewRouter().PathPrefix(rootRoute).Subrouter()
	rootRouter.Use(s.HeadersMiddleware)

	if s.config.PRODUCTION {
		rootRouter.Use(s.APIKeyMiddleware)
	} else {
		rootRouter.Use(s.CORSMiddleware)
	}

	rootRouter.HandleFunc("/", s.handleGetAll)

	authRouter := rootRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("", s.handleGetAll).Methods("GET")
	authRouter.HandleFunc("/login", s.handleLogin).Methods("POST")
	authRouter.HandleFunc("/register", s.handleRegister).Methods("POST")

	teamsRouter := rootRouter.PathPrefix("/teams").Subrouter()
	//teamsRouter.Use(s.TokenRefreshMiddleware, s.Authenticate)
	teamsRouter.HandleFunc("", s.handleGetAllTeams).Methods("GET")
	teamsRouter.HandleFunc("/{id}", s.handleGetTeam).Methods("GET")
	teamsRouter.HandleFunc("/{id}/players", s.handleGetTeamPlayers).Methods("GET")
	teamsRouter.HandleFunc("", s.handleTeamCreation).Methods("POST")

	playersRouter := rootRouter.PathPrefix("/players").Subrouter()
	//playersRouter.Use(s.TokenRefreshMiddleware, s.AuthenticateMiddleware)
	playersRouter.HandleFunc("", s.handleGetAllPlayers).Methods("GET")

	http.Handle("/", rootRouter)

	return http.ListenAndServe(s.listenAddr, nil)
}
