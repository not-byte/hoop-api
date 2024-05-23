package model

import (
	"time"
)

type Account struct {
	Email         string
	Password      string
	ID            uint64
	PermissionsID uint64
	CreatedOn     time.Time
	LoggedOn      time.Time
	Verified      bool
	MailToken     int8
}

type AccountPermission struct {
	PermissionsID uint64
	AccountsID    uint64
}
