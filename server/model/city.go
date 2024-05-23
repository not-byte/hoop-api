package model

type CityDTO struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type City struct {
	ID    uint64
	Name  string
	State string
}
