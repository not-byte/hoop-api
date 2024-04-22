package store

import (
	"context"
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

func (s *SQLStore) GetAccount(id int64) (*types.Account, error) {
	stmt, err := s.DB.Prepare("SELECT id, name FROM accounts WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var account types.Account
	err = stmt.QueryRow(id).Scan(&account.ID, &account.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &account, nil
}

func (s *SQLStore) CreateAccount(ctx context.Context, email string, password string) error {

	fail := func(err error) error {
		return fmt.Errorf("CreateAccount: %v", err)
	}

	var account types.Account = types.Account{
		Email:    email,
		Password: password,
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	if err := tx.QueryRow(
		"INSERT INTO permissions (type) VALUES (DEFAULT) ON CONFLICT DO NOTHING RETURNING id",
	).Scan(&account.PermissionsID); err != nil {
		return fail(err)
	}

	if err := tx.QueryRow(
		"INSERT INTO accounts (permissions_id, email, password, token) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING RETURNING id",
		account.PermissionsID,
		account.Email,
		account.Password,
		account.MailToken,
	).Scan(&account.ID); err != nil {
		return fail(err)
	}

	if err := tx.Commit(); err != nil {
		return fail(err)
	}

	return nil
}

func (s *SQLStore) UpdateAccount(account *types.Account) error {

	stmt, err := s.DB.Prepare("DELETE FROM accounts WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(account.ID)
	if err != nil {
		return fmt.Errorf("UpdateAccount: executing delete: %v", err)
	}

	//probably not needed right now
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdateAccount: failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("UpdateAccount: no rows deleted, possibly no account with ID %d", account.ID)
	}

	return nil
}

func (s *SQLStore) LoginAccount(id int64) error {
	stmt, err := s.DB.Prepare("UPDATE accounts SET logged_on = now() WHERE id = $1")
	if err != nil {
		return fmt.Errorf("LoginAccount: preparing update statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("LoginAccount: executing update: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("LoginAccount: failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("LoginAccount: no accounts updated, possibly no account with ID %d exists", id)
	}

	return nil
}

func (s *SQLStore) VerifyAccount(id int64) error {
	stmt, err := s.DB.Prepare("UPDATE accounts SET verified = $1 WHERE id = $2")
	if err != nil {
		return fmt.Errorf("VerifyAccount: preparing statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(true, id)
	if err != nil {
		return fmt.Errorf("VerifyAccount: executing update: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("VerifyAccount: failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("VerifyAccount: no accounts updated, possibly no account with ID %d exists", id)
	}

	return nil
}

func (s *SQLStore) DeleteAccount(id int64) error {
	stmt, err := s.DB.Prepare("DELETE FROM accounts WHERE id = $1")
	if err != nil {
		return fmt.Errorf("DeleteAccount: preparing statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("DeleteAccount: executing delete: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteAccount: error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("DeleteAccount: no account found with ID %d", id)
	}

	return nil
}

func (s *SQLStore) GetAccounts() ([]types.Account, error) {
	stmt, err := s.DB.Prepare("SELECT * FROM accounts WHERE verified = $1")
	if err != nil {
		return nil, fmt.Errorf("GetAccounts: preparing statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(true)
	if err != nil {
		return nil, fmt.Errorf("GetAccounts: executing query: %v", err)
	}
	defer rows.Close()

	var accounts []types.Account

	for rows.Next() {
		var account types.Account
		if err := rows.Scan(
			&account.ID,
			&account.PermissionsID,
			&account.Email,
			&account.Password,
			&account.MailToken,
			&account.CreatedOn,
			&account.LoggedOn,
			&account.Verified,
		); err != nil {
			return nil, fmt.Errorf("GetAccounts: scanning results: %v", err)
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAccounts: iterating over rows: %v", err)
	}

	return accounts, nil
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
