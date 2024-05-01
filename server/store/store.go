package store

import (
	"context"
	"tournament_api/server/api/model"
)

type Store interface {
	/*Accounts*/
	GetAccountByEmail(email string) (*model.Account, error)
	CreateAccount(ctx context.Context, email *string, password *string, mailToken int8) error
	UpdateAccount(account *model.Account) error
	LoginAccount(id int64) error
	VerifyAccount(id int64) error
	DeleteAccount(id int64) error
	GetAccounts() ([]model.Account, error)

	/*Teams*/
	GetTeams() ([]model.Team, error)
	GetTeam(id int) (*model.Team, error)
	CreateTeam(ctx context.Context, name *string, description *string) error
	UpdateTeam(id int) error
	DeleteTeam(id int) error
}
