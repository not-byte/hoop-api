package model

import (
	"time"
)

type TeamDTO struct {
	ID         uint64 `json:"id"`
	CategoryID uint64 `json:"category_id"`
	CityID     uint64 `json:"city_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
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
