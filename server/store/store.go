package store

import (
	"context"
	"tournament_api/server/types"
)

type Store interface {
	GetAccountByEmail(email string) (*types.Account, error)
	CreateAccount(ctx context.Context, email *string, password *string, mailToken int8) error
	UpdateAccount(account *types.Account) error
	LoginAccount(id int64) error
	VerifyAccount(id int64) error
	DeleteAccount(id int64) error
	GetAccounts() ([]types.Account, error)
}
