package store

import (
	"context"
	"database/sql"
	"fmt"
	"tournament_api/server/model"
)

func (store *SQLStore) GetTeams() ([]model.TeamDTO, error) {
	fail := func(err error) error {
		return fmt.Errorf("GetTeams: %v", err)
	}

	stmt, err := store.DB.Prepare("SELECT teams.id, categories.id, cities.id, teams.name, teams.email, teams.phone FROM teams, categories, cities WHERE teams.categories_id = categories.id AND teams.cities_id = cities.id")
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
			&team.CategoryID,
			&team.CityID,
			&team.Name,
			&team.Email,
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

func (store *SQLStore) GetTeam(id uint64) (*model.TeamDTO, error) {
	fail := func(err error) error {
		return fmt.Errorf("GetTeam: %v", err)
	}

	var team model.TeamDTO

	err := store.DB.QueryRow("SELECT teams.id, categories.id, cities.id, teams.name, teams.email, teams.phone FROM teams, categories, cities WHERE teams.categories_id = categories.id AND teams.cities_id = cities.id AND teams.id = $1", id).Scan(
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

func (store *SQLStore) CreateTeam(ctx context.Context, team *model.Team) error {
	tx, err := store.DB.BeginTx(ctx, nil)
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

	//if err := insertPlayers(tx, team.Players, id); err != nil {
	//	return fmt.Errorf("CreateTeam: %v", err)
	//}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("CreateTeam: committing transaction: %v", err)
	}

	return nil
}

func (store *SQLStore) UpdateTeam(team *model.Team) error {
	//tx, err := store.DB.BeginTx(context.Background(), nil)
	//if err != nil {
	//	return fmt.Errorf("UpdateTeam: starting transaction: %v", err)
	//}
	//defer tx.Rollback()
	//
	//_, err = tx.Exec("UPSERT teams SET name =  email =$1, $2, $3, description = phone = $4 WHERE id = $5",
	//	team.Name, team.Email, team.Description, team.Phone, team.ID)
	//if err != nil {
	//	return fmt.Errorf("UpdateTeam: %v", err)
	//}
	//
	//if err := deletePlayers(tx, *team.ID); err != nil {
	//	return fmt.Errorf("UpdateTeam: %v", err)
	//}
	//
	//if err := insertPlayers(tx, team.Players, team.ID); err != nil {
	//	return fmt.Errorf("UpdateTeam: %v", err)
	//}
	//
	//if err := tx.Commit(); err != nil {
	//	return fmt.Errorf("UpdateTeam: committing transaction: %v", err)
	//}

	return nil
}

func (store *SQLStore) DeleteTeam(id uint64) error {
	_, err := store.DB.Exec("DELETE FROM teams WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("DeleteTeam: %v", err)
	}
	return nil
}

func insertTeam(tx *sql.Tx, team *model.Team) (*uint64, error) {
	var id uint64
	//if err := tx.QueryRow(
	//	"INSERT INTO teams (name, email, description, phone, categories_id, cities_id) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING RETURNING ID",
	//	team.Name, team.Email, team.Description, team.Phone, team.Gender, team.Category,
	//).Scan(&id); err != nil {
	//	return nil, fmt.Errorf("inserting team : %v", err)
	//}

	return &id, nil
}
