package model

import (
	"math/big"
	"time"
)

type Permission struct {
	ID    big.Int
	Type  string
	Flags int
}

type Audit struct {
	ID      big.Int
	Time    time.Time
	Status  int
	Message string
}
