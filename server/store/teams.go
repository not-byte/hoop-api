package store

import (
	"context"
	"database/sql"
	"fmt"
	"tournament_api/server/api/model"
)

func (s *SQLStore) GetTeams() ([]model.Team, error) {
	fail := func(err error) error {
		return fmt.Errorf("GetTeams: %v", err)
	}

	stmt, err := s.DB.Prepare("SELECT * FROM teams")
	if err != nil {
		return nil, fail(fmt.Errorf("preparing statement: %v", err))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, fail(fmt.Errorf("executing query: %v", err))
	}
	defer rows.Close()

	var teams []model.Team

	for rows.Next() {
		var team model.Team
		if err := rows.Scan(
			&team.ID,
			&team.CitiesID,
			&team.Name,
			&team.Description,
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

func (s *SQLStore) GetTeam(id int) (*model.Team, error) {
	fail := func(err error) error {
		return fmt.Errorf("GetTeam: %v", err)
	}

	var team model.Team

	err := s.DB.QueryRow("SELECT * FROM teams WHERE id = $1", 1).Scan(&team.ID, &team.CitiesID, &team.Name, &team.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fail(err)
		}
		return nil, fail(err)
	}

	return &team, nil

}

func (s *SQLStore) CreateTeam(ctx context.Context, name *string, description *string) error {
	fail := func(err error) error {
		return fmt.Errorf("CreateTeam: %v", err)
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	if _, err := tx.Exec(
		"INSERT INTO cities (type) VALUES (DEFAULT) ON CONFLICT DO NOTHING",
	); err != nil {
		return fail(err)
	}

	result, err := tx.Exec(
		"INSERT INTO teams (name, description) VALUES ($1, $2) ON CONFLICT DO NOTHING",
		name,
		description,
	)
	if err != nil {
		return fail(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fail(err)
	}

	if rowsAffected == 0 {
		return fail(fmt.Errorf("0 rows were affected, something went wrong"))
	}

	if err := tx.Commit(); err != nil {
		return fail(err)
	}

	return nil
}

func (s *SQLStore) UpdateTeam(id int) error {
	return nil
}

func (s *SQLStore) DeleteTeam(id int) error {
	return nil
}
