package store

import (
	"context"
	"tournament_api/server/model"
	"tournament_api/server/types"
)

type Store interface {
	GetAccountByEmail(email string) (*model.Account, error)
	CreateAccount(ctx context.Context, email string, password string, mailToken int8) error
	UpdateAccount(account *model.Account) error
	LoginAccount(id int64) error
	VerifyAccount(id int64) error
	DeleteAccount(id int64) error
	GetAccounts() ([]model.Account, error)

	GetTeams() ([]model.TeamDTO, error)
	GetTeam(id uint64) (*model.Team, error)
	CreateTeam(ctx context.Context, team *types.Team) error
	UpdateTeam(team *types.Team) error
	DeleteTeam(id uint64) error

	GetPlayers(team_id uint64) ([]model.PlayerDTO, error)
	GetPlayer(id uint64) (*model.PlayerDTO, error)
	CreatePlayer(ctx context.Context, player *types.Player) error
	UpdatePlayer(player *types.Player) error
	DeletePlayer(id uint64) error
}
