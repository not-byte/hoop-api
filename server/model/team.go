package model

import "time"

type Team struct {
	ID           int64
	CitiesID     *int64
	CategoriesID *int64
	Category     int32
	Name         string
	Email        string
	Phone        string
	Description  string
	Gender       string
	CreatedOn    time.Time
}

// TeamDTO is a data transfer object for teams
type TeamDTO struct {
	ID       int64
	Name     string
	Category string
	Email    string
	Phone    string
	City     string
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

type PlayerDTO struct {
	ID        int64
	FirstName string
	LastName  string
	Age       int8
}

type TeamPlayer struct {
	TeamsID   int64
	PlayersID int64
}
