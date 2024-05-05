package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tournament_api/server/model"
	"tournament_api/server/types"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func (s *Server) handleGetAllTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := s.store.GetTeams()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string][]model.Team{"teams": teams})
}

func (s *Server) handleGetTeam(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid team ID"+err.Error(), http.StatusBadRequest)
		return
	}

	team, err := s.store.GetTeam(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid login credentials"+err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]model.Team{"team": *team})
}

func (s *Server) handleTeamCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	var team types.Team
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

	err = s.store.CreateTeam(r.Context(), &team)
	if err != nil {
		http.Error(w, "Error while creating a team: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "creation  successful"})
}
