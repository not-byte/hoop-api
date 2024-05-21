package store

import (
	"context"
	"database/sql"
	"tournament_api/server/model"
	"tournament_api/server/types"
)

func (store *SQLStore) GetPlayers(teamID int64) ([]model.PlayerDTO, error) {
	var (
		stmt *sql.Stmt
		rows *sql.Rows
		err  error
	)

	if teamID == -1 {
		stmt, err = store.DB.Prepare("SELECT id, first_name, last_name, age FROM players")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err = stmt.Query()
	} else {
		stmt, err = store.DB.Prepare("SELECT id, first_name, last_name, age FROM players WHERE teams_id = $1")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		rows, err = stmt.Query(teamID)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []model.PlayerDTO

	for rows.Next() {
		var player model.PlayerDTO
		if err := rows.Scan(
			&player.ID,
			&player.FirstName,
			&player.LastName,
			&player.Age,
		); err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
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

func (store *SQLStore) GetPlayer(id int64) (*model.PlayerDTO, error) {
	fail := func(err error) error {
		return err
	}

	var player model.PlayerDTO

	err := store.DB.QueryRow("SELECT first_name, last_name, age FROM players WHERE id = $1", id).Scan(
		&player.FirstName,
		&player.LastName,
		&player.Age,
	)
	if err != nil {
		return nil, fail(err)
	}

	return &player, nil
}

func (s *SQLStore) CreatePlayer(ctx context.Context, player *types.Player) error {
	return nil
}

func (s *SQLStore) UpdatePlayer(player *types.Player) error {
	return nil
}

func (s *SQLStore) DeletePlayer(id int64) error {
	return nil
}
