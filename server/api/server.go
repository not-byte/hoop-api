package api

import (
	"fmt"
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

func (server *Server) Start() error {
	rootRoute := fmt.Sprintf("/v%d", server.config.VERSION)
	rootRouter := mux.NewRouter().PathPrefix(rootRoute).Subrouter()
	rootRouter.Use(server.HeadersMiddleware)

	if server.config.PRODUCTION {
		rootRouter.Use(server.APIKeyMiddleware)
	} else {
		rootRouter.Use(server.CORSMiddleware)
	}

	if server.config.COMPRESS {
		rootRouter.Use(GzipMiddleware)
	}

	rootRouter.HandleFunc("/", server.handleGetAll)

	authRouter := rootRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("", server.handleGetAll).Methods("GET")
	authRouter.HandleFunc("/login", server.handleLogin).Methods("POST")
	authRouter.HandleFunc("/register", server.handleRegister).Methods("POST")

	teamsRouter := rootRouter.PathPrefix("/teams").Subrouter()
	//teamsRouter.Use(s.TokenRefreshMiddleware, s.Authenticate)
	teamsRouter.HandleFunc("", server.handleGetAllTeams).Methods("GET")
	teamsRouter.HandleFunc("/{id}", server.handleGetTeam).Methods("GET")
	teamsRouter.HandleFunc("/{id}/players", server.handleGetTeamPlayers).Methods("GET")
	teamsRouter.HandleFunc("", server.handleTeamCreation).Methods("POST")

	playersRouter := rootRouter.PathPrefix("/players").Subrouter()
	//playersRouter.Use(s.TokenRefreshMiddleware, s.AuthenticateMiddleware)
	playersRouter.HandleFunc("", server.handleGetAllPlayers).Methods("GET")

	http.Handle("/", rootRouter)

	return http.ListenAndServe(server.listenAddr, nil)
}
