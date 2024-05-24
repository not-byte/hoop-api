package store

import (
	"context"
	"tournament_api/server/model"
)

type Store interface {
	GetAccountByEmail(email string) (*model.Account, error)
	CreateAccount(ctx context.Context, email string, password string, mailToken int8) error
	UpdateAccount(account *model.Account) error
	LoginAccount(id uint64) error
	VerifyAccount(id uint64) error
	DeleteAccount(id uint64) error
	GetAccounts() ([]model.Account, error)

	// GetCategories() ([]model.CategoryDTO, error)
	// GetCategory(id uint64) (*model.CategoryDTO, error)
	// CreateCategory(ctx context.Context, category *model.Category) error
	// UpdateCategory(category []model.CategoryDTO) error
	// DeleteCategory(id uint64) error

	// GetCities() ([]model.CityDTO, error)
	// GetCity(id uint64) (*model.CityDTO, error)
	// CreateCity(ctx context.Context, category *model.City) error
	// UpdateCity(category []model.CityDTO) error
	// DeleteCity(id uint64) error

	GetPlayers() ([]model.PlayerDTO, error)
	// GetPlayersByTeam(teamID uint64) ([]model.PlayerDTO, error)
	GetPlayer(id uint64) (*model.PlayerDTO, error)
	CreatePlayer(ctx context.Context, player *model.Player) error
	UpdatePlayer(player *model.Player) error
	DeletePlayer(id uint64) error

	GetTeams() ([]model.TeamDTO, error)
	GetTeam(id uint64) (*model.TeamDTO, error)
	CreateTeam(ctx context.Context, team *model.Team) error
	UpdateTeam(team *model.Team) error
	DeleteTeam(id uint64) error
}
