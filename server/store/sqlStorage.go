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
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = initialize(db)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to initialize tables: %w", err)
	}

	fmt.Println("Successfully connected to the database!")
	return &SQLStore{DB: db}, nil
}

func (s *SQLStore) Get() any {
	var value any = "mock"
	return value
}

func initialize(db *sql.DB) error {
	content, err := os.ReadFile("/home/pszymanski/git/tournament-api/storage/sql/tables.sql")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	db.Exec(string(content))
	return nil
}
