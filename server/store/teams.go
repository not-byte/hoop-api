package store

import (
	"context"
	"database/sql"
	"fmt"
	"tournament_api/server/model"
	"tournament_api/server/types"
	"tournament_api/server/utils"
)

func (s *SQLStore) GetTeams() ([]model.TeamDTO, error) {
	fail := func(err error) error {
		return fmt.Errorf("GetTeams: %v", err)
	}

	stmt, err := s.DB.Prepare("SELECT id, name, email, category, phone FROM teams")
	if err != nil {
		return nil, fail(fmt.Errorf("preparing statement: %v", err))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, fail(fmt.Errorf("executing query: %v", err))
	}
	defer rows.Close()

	var teams []model.TeamDTO

	for rows.Next() {
		var team model.TeamDTO
		if err := rows.Scan(
			&team.ID,
			&team.Name,
			&team.Email,
			&team.Category,
			&team.Phone,
		); err != nil {
			return nil, fail(fmt.Errorf("scanning results: %v", err))
		}
		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, fail(fmt.Errorf("iterating over rows: %v", err))
	}

	return teams, nil
}

func (s *SQLStore) GetTeam(id int64) (*model.Team, error) {
	fail := func(err error) error {
		return fmt.Errorf("GetTeam: %v", err)
	}

	var team model.Team

	err := s.DB.QueryRow("SELECT name, email, category, phone FROM teams WHERE id = $1", id).Scan(&team.ID,
		&team.Name,
		&team.Email,
		&team.Category,
		&team.Phone,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fail(err)
		}
		return nil, fail(err)
	}

	return &team, nil

}

func (s *SQLStore) CreateTeam(ctx context.Context, team *types.Team) error {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("CreateTeam: starting transaction: %v", err)
	}
	defer tx.Rollback()

	id, err := insertTeam(tx, team)
	if err != nil {
		return fmt.Errorf("CreateTeam: %v", err)
	}
	if id == nil {
		return fmt.Errorf("CreateTeam: %v", err)
	}

	if err := insertPlayers(tx, team.Players, int64(*id)); err != nil {
		return fmt.Errorf("CreateTeam: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("CreateTeam: committing transaction: %v", err)
	}

	return nil
}

func (s *SQLStore) UpdateTeam(team *types.Team) error {
	tx, err := s.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("UpdateTeam: starting transaction: %v", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE teams SET name = $1, email = $2, description = $3, phone = $4, gender = $5 WHERE id = $6",
		team.Name, team.Email, team.Description, team.Phone, team.Gender, team.ID)
	if err != nil {
		return fmt.Errorf("UpdateTeam: %v", err)
	}

	if err := deletePlayers(tx, int64(*team.ID)); err != nil {
		return fmt.Errorf("UpdateTeam: %v", err)
	}

	if err := insertPlayers(tx, team.Players, int64(*team.ID)); err != nil {
		return fmt.Errorf("UpdateTeam: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("UpdateTeam: committing transaction: %v", err)
	}

	return nil
}

func (s *SQLStore) DeleteTeam(id int64) error {
	_, err := s.DB.Exec("DELETE FROM teams WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("DeleteTeam: %v", err)
	}
	return nil
}

// helper functions
func insertTeam(tx *sql.Tx, team *types.Team) (*int, error) {
	var id int
	if err := tx.QueryRow(
		"INSERT INTO teams (name, email, description, phone, gender, category) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING RETURNING ID",
		team.Name, team.Email, team.Description, team.Phone, team.Gender, team.Category,
	).Scan(&id); err != nil {
		return nil, fmt.Errorf("inserting team : %v", err)
	}

	return &id, nil
}

func insertPlayers(tx *sql.Tx, players []*types.Player, team_id int64) error {
	for _, player := range players {
		player.TeamID = team_id
	}
	query, err := utils.BulkInsert(players, "players")
	if err != nil {
		return fmt.Errorf("failed to create bulk insert query: %v", err)
	}
	query += " ON CONFLICT DO NOTHING"
	fmt.Println(query)
	_, err = tx.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to execute the bulk insert statement: %v", err)
	}

	return nil
}

func deletePlayers(tx *sql.Tx, teamID int64) error {
	_, err := tx.Exec("DELETE FROM players WHERE team_id = $1", teamID)
	if err != nil {
		return fmt.Errorf("deletePlayers: %v", err)
	}
	return nil
}
