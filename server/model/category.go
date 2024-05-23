package model

import "math/big"

type CategoryDTO struct {
	ID     big.Int `json:"id"`
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
}

type Category struct {
	ID        big.Int
	TeamLimit big.Int
	Name      string
	Gender    string
}
