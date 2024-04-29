package types

import "time"

type Permission struct {
	ID    int64
	Type  string
	Flags int
}

type City struct {
	ID    int64
	Name  string
	State string
}


type Audit struct {
	ID      int64
	Time    time.Time
	Status  int
	Message string
}
