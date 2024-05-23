package model

import "math/big"

type CityDTO struct {
	ID    big.Int `json:"id"`
	Name  string  `json:"name"`
	State string  `json:"state"`
}

type City struct {
	ID    big.Int
	Name  string
	State string
}
