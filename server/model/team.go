package model

import (
	"math/big"
	"time"
)

type TeamDTO struct {
	ID         big.Int `json:"id"`
	CategoryID big.Int `json:"category_id"`
	CityID     big.Int `json:"city_id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
}

type Team struct {
	ID           big.Int
	CitiesID     big.Int
	CategoriesID big.Int
	Name         string
	Category     string
	Email        string
	Phone        string
	City         string
	CreatedOn    time.Time
}
