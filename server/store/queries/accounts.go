package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"tournament_api/server/config"
	"tournament_api/server/store"
	"tournament_api/server/types"
	"tournament_api/server/utils"

	"github.com/cockroachdb/cockroach-go/crdb"
	_ "github.com/lib/pq"
)

func main() {
	account1 := CreateAccount("test1@test.com", "Test1")
	account2 := CreateAccount("test2@test.com", "Test2")

	VerifyAccount(account1.ID)

	GetAccounts()

	VerifyAccount(account2.ID)

	GetAccounts()

	DeleteAccount(account1.ID)
	DeleteAccount(account2.ID)
}

func CreateAccount(email, password string) *types.Account {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	store, err := store.NewSQLStore(config)
	if err != nil {
		log.Fatalf("Failed to create SQL Store: %v", err)
	}

	var account types.Account = types.Account{
		Email:    email,
		Password: password,
	}

	err = crdb.ExecuteTx(context.Background(), store.DB, nil, func(tx *sql.Tx) error {
		if err := tx.QueryRow(
			"INSERT INTO permissions (type) VALUES (DEFAULT) ON CONFLICT DO NOTHING RETURNING id",
		).Scan(&account.PermissionsID); err != nil {
			return err
		}

		account.Token = utils.GenerateToken()

		if err := tx.QueryRow(
			"INSERT INTO accounts (permissions_id, email, password, token) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING RETURNING id",
			account.PermissionsID,
			account.Email,
			account.Password,
			account.Token,
		).Scan(&account.ID); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("ID: %d\n", account.ID)
	fmt.Printf("PermissionsID: %d\n", account.PermissionsID)
	fmt.Printf("Email: %s\n", account.Email)
	fmt.Printf("Password: %s\n", account.Password)
	fmt.Printf("CreatedOn: %s\n", account.CreatedOn)
	fmt.Printf("LoggedOn: %s\n", account.LoggedOn)
	fmt.Printf("Verified: %t\n", account.Verified)
	fmt.Printf("Token: %d\n", account.Token)

	return &account
}

func GetAccount(id int64) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	store, err := store.NewSQLStore(config)
	if err != nil {
		log.Fatalf("Failed to create SQL Store: %v", err)
	}

	store.DB.Query(
		"SELECT * FROM accounts WHERE id = $1",
		id,
	)
}

func UpdateAccount(account *types.Account) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	store, err := store.NewSQLStore(config)
	if err != nil {
		log.Fatalf("Failed to create SQL Store: %v", err)
	}

	store.DB.Query(
		"DELETE FROM accounts WHERE id = $1 ON CONFLICT DO NOTHING",
		account.ID,
	)
}

func LoginAccount(id int64) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	store, err := store.NewSQLStore(config)
	if err != nil {
		log.Fatalf("Failed to create SQL Store: %v", err)
	}

	store.DB.Query(
		"UPDATE accounts SET logged_on = now() WHERE id = $1 ",
		id,
	)
}

func VerifyAccount(id int64) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	store, err := store.NewSQLStore(config)
	if err != nil {
		log.Fatalf("Failed to create SQL Store: %v", err)
	}

	_, err = store.DB.Query(
		"UPDATE accounts SET verified = $1 WHERE id = $2",
		true,
		id,
	)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func DeleteAccount(id int64) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	store, err := store.NewSQLStore(config)
	if err != nil {
		log.Fatalf("Failed to create SQL Store: %v", err)
	}

	_, err = store.DB.Query(
		"DELETE FROM accounts WHERE id = $1",
		id,
	)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func GetAccounts() []types.Account {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	store, err := store.NewSQLStore(config)
	if err != nil {
		log.Fatalf("Failed to create SQL Store: %v", err)
	}

	rows, err := store.DB.Query(
		"SELECT * FROM accounts WHERE verified = $1",
		true,
	)
	if err != nil {
		log.Fatalf("Error: %v", err)
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
			&account.Token,
			&account.CreatedOn,
			&account.LoggedOn,
			&account.Verified,
		); err != nil {
			log.Fatalf("Error: %v", err)
		}

		accounts = append(accounts, account)

		fmt.Printf("ID: %d\n", account.ID)
		fmt.Printf("PermissionsID: %d\n", account.PermissionsID)
		fmt.Printf("Email: %s\n", account.Email)
		fmt.Printf("Password: %s\n", account.Password)
		fmt.Printf("CreatedOn: %s\n", account.CreatedOn)
		fmt.Printf("LoggedOn: %s\n", account.LoggedOn)
		fmt.Printf("Verified: %t\n", account.Verified)
		fmt.Printf("Token: %d\n", account.Token)
	}

	return accounts
}
