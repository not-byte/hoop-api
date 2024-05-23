package model

import (
	"time"
)

type PlayerDTO struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int8      `json:"age"`
	Birthday  time.Time `json:"birthday"`
	Number    int       `json:"number"`
	Height    int       `json:"height"`
	Weight    int       `json:"weight"`
	Wingspan  int       `json:"wingspan"`
	Position  string    `json:"position"`
}

type Player struct {
	ID         uint64
	AccountsID uint64
	TeamsID    uint64
	FirstName  string
	LastName   string
	FullName   string
	Age        int8
	Birthday   time.Time
	Number     int
	Height     int
	Weight     int
	Wingspan   int
	Position   string
	CreatedOn  time.Time
}

type TeamPlayer struct {
	TeamsID   uint64
	PlayersID uint64
}
