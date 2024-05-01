package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tournament_api/server/model"
	"tournament_api/server/types"
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

func (s *Server) handleTeamCreation(w http.ResponseWriter, r *http.Request) {
	var team types.Team
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, "Error while decoding: ", http.StatusInternalServerError)
		return
	}

	err = s.store.CreateTeam(r.Context(), team.Name, team.Description)
	if err != nil {
		http.Error(w, "Error while creating a team: ", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "creation  successful"})
}
