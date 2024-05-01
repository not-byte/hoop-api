package model

import "time"

type Team struct {
	ID          int64
	CitiesID    int64
	Name        string
	Description string
}

type Player struct {
	ID         int64
	AccountsID int64
	TeamsID    int64
	FirstName  string
	LastName   string
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