package model

type CategoryDTO struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type Category struct {
	ID        uint64
	TeamLimit uint64
	Name      string
	Gender    string
}
