package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tournament_api/server/model"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func (server *Server) handleGetAllTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := server.store.GetTeams()
	if err != nil {
		http.Error(w, "Invalid login credentials "+err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string][]model.TeamDTO{"teams": teams})
}

func (server *Server) handleGetTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id := new(uint64)
	_, err := fmt.Sscan(idStr, id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid team ID"+err.Error(), http.StatusBadRequest)
		return
	}

	team, err := server.store.GetTeam(*id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid login credentials "+err.Error(), http.StatusUnauthorized)
		return
	}
	if team == nil {
		http.Error(w, "Team not found "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]model.TeamDTO{"team": *team})
}

func (server *Server) handleTeamCreation(w http.ResponseWriter, r *http.Request) {

	var team model.Team
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, "Error while decoding: "+err.Error(), http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	err = validate.Struct(team)
	if err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = server.store.CreateTeam(r.Context(), &team)
	if err != nil {
		http.Error(w, "Error while creating a team: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "creation  successful"})
}
