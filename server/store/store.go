package store

import (
	"context"
	"math/big"
	"tournament_api/server/model"
	"tournament_api/server/types"
)

type Store interface {
	GetAccountByEmail(email string) (*model.Account, error)
	CreateAccount(ctx context.Context, email string, password string, mailToken int8) error
	UpdateAccount(account *model.Account) error
	LoginAccount(id big.Int) error
	VerifyAccount(id big.Int) error
	DeleteAccount(id big.Int) error
	GetAccounts() ([]model.Account, error)

	GetTeams() ([]model.TeamDTO, error)
	GetTeam(id big.Int) (*model.TeamDTO, error)
	CreateTeam(ctx context.Context, team *types.Team) error
	UpdateTeam(team *types.Team) error
	DeleteTeam(id big.Int) error

	GetPlayers() ([]model.PlayerDTO, error)
	GetPlayer(id big.Int) (*model.PlayerDTO, error)
	CreatePlayer(ctx context.Context, player *types.Player) error
	UpdatePlayer(player *types.Player) error
	DeletePlayer(id big.Int) error
}
