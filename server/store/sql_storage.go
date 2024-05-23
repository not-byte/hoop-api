package store

import (
	"database/sql"
	"fmt"
	"os"
	"tournament_api/server/types"

	_ "github.com/lib/pq"
)

type SQLStore struct {
	DB *sql.DB
}

func NewSQLStore(config *types.AppConfig) (*SQLStore, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=defaultdb sslmode=disable",
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if !config.PRODUCTION {
		err = clear(db)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to clear tables: %w", err)
		}
	}

	err = initialize(db)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to initialize tables: %w", err)
	}

	if !config.PRODUCTION {
		seed(db, "categories")
		seed(db, "cities")
		seed(db, "players")
		seed(db, "teams")
	}

	return &SQLStore{DB: db}, nil
}

func clear(db *sql.DB) error {
	content, err := os.ReadFile("storage/sql/clear.sql")
	if err != nil {
		fmt.Println("error reading file:", err)
		return err
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return err
	}

	return nil
}

func initialize(db *sql.DB) error {
	content, err := os.ReadFile("storage/sql/create.sql")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	_, err = db.Exec(string(content))
	if err != nil {
		fmt.Println("error executing creation of tables:", err)
		return err
	}

	return nil
}

func seed(db *sql.DB, table string) error {
	content, err := os.ReadFile(fmt.Sprintf("storage/sql/seed/%s.sql", table))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	_, err = db.Exec(string(content))
	if err != nil {
		fmt.Println("error executing seeding of tables:", err)
		return err
	}

	return nil
}
