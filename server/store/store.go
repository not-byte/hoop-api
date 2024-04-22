package store

import (
	"context"
	"tournament_api/server/types"
)

type Store interface {
	Get() any
	GetAccount(id int64) (*types.Account, error)
	CreateAccount(ctx context.Context, email string, password string) error
	UpdateAccount(account *types.Account) error
	LoginAccount(id int64) error
	VerifyAccount(id int64) error
	DeleteAccount(id int64) error
	GetAccounts() ([]types.Account, error)
}
