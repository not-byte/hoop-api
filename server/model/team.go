package model

import (
	"time"
)

type TeamDTO struct {
	ID       uint64      `json:"id"`
	Category CategoryDTO `json:"category"`
	City     CityDTO     `json:"city"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Phone    string      `json:"phone"`
}

type Team struct {
	ID           uint64
	CitiesID     uint64
	CategoriesID uint64
	Name         string
	Category     string
	Email        string
	Phone        string
	City         string
	CreatedOn    time.Time
}
