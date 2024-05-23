package model

import (
	"time"
)

type Permission struct {
	ID    uint64
	Type  string
	Flags int
}

type Audit struct {
	ID      uint64
	Time    time.Time
	Status  int
	Message string
}
