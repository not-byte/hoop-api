package model

import "time"

type CategoryDTO struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type CityDTO struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type TeamDTO struct {
	ID         int64  `json:"id"`
	CategoryID int64  `json:"category_id"`
	CityID     int64  `json:"city_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type PlayerDTO struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int8   `json:"age"`
}

type Category struct {
	ID        int64
	TeamLimit int64
	Name      string
	Gender    string
}

type Team struct {
	ID           int64
	CitiesID     int64
	CategoriesID int64
	Name         string
	Category     string
	Email        string
	Phone        string
	City         string
	CreatedOn    time.Time
}

type Player struct {
	ID         int64
	AccountsID int64
	TeamsID    int64
	FirstName  string
	LastName   string
	Age        int8
	FullName   string
	Birthday   time.Time
	Number     int
	Height     int
	Weight     int
	Wingspan   int
	Position   string
}

type TeamPlayer struct {
	TeamsID   int64
	PlayersID int64
}
