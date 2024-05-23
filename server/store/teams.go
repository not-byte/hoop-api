package store

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"tournament_api/server/model"
	"tournament_api/server/types"
)

func (s *SQLStore) GetTeams() ([]model.TeamDTO, error) {
	fail := func(err error) error {
		return fmt.Errorf("GetTeams: %v", err)
	}

	stmt, err := s.DB.Prepare("SELECT teams.id, categories.id, cities.id, teams.name, teams.email, teams.phone FROM teams, categories, cities WHERE teams.categories_id = categories.id AND teams.cities_id = cities.id")
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
		var idStr, categoryIDStr, cityIDStr string
		var team model.TeamDTO
		if err := rows.Scan(
			&idStr,
			&categoryIDStr,
			&cityIDStr,
			&team.Name,
			&team.Email,
			&team.Phone,
		); err != nil {
			return nil, fail(fmt.Errorf("scanning results: %v", err))
		}

		team.ID, _ = new(big.Int).SetString(idStr, 10)
		team.CategoryID, _ = new(big.Int).SetString(categoryIDStr, 10)
		team.CityID, _ = new(big.Int).SetString(cityIDStr, 10)

		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, fail(fmt.Errorf("iterating over rows: %v", err))
	}

	return teams, nil
}

func (s *SQLStore) GetTeam(id big.Int) (*model.TeamDTO, error) {
	fail := func(err error) error {
		return fmt.Errorf("GetTeam: %v", err)
	}

	var team model.TeamDTO

	err := s.DB.QueryRow("SELECT teams.id, categories.id, cities.id, teams.name, teams.email, teams.phone FROM teams, categories, cities WHERE teams.categories_id = categories.id AND teams.cities_id = cities.id AND teams.id = $1", id).Scan(
		&team.ID,
		&team.CategoryID,
		&team.CityID,
		&team.Name,
		&team.Email,
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

	if err := insertPlayers(tx, team.Players, id); err != nil {
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

	if err := deletePlayers(tx, *team.ID); err != nil {
		return fmt.Errorf("UpdateTeam: %v", err)
	}

	if err := insertPlayers(tx, team.Players, team.ID); err != nil {
		return fmt.Errorf("UpdateTeam: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("UpdateTeam: committing transaction: %v", err)
	}

	return nil
}

func (s *SQLStore) DeleteTeam(id big.Int) error {
	_, err := s.DB.Exec("DELETE FROM teams WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("DeleteTeam: %v", err)
	}
	return nil
}

func insertTeam(tx *sql.Tx, team *types.Team) (*big.Int, error) {
	var id big.Int
	if err := tx.QueryRow(
		"INSERT INTO teams (name, email, description, phone, gender, category) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING RETURNING ID",
		team.Name, team.Email, team.Description, team.Phone, team.Gender, team.Category,
	).Scan(&id); err != nil {
		return nil, fmt.Errorf("inserting team : %v", err)
	}

	return &id, nil
}
