package model

import (
	"time"
)

type Account struct {
	Email         string
	Password      string
	ID            int64
	PermissionsID int64
	CreatedOn     time.Time
	LoggedOn      time.Time
	Verified      bool
	MailToken     int8
}

type AccountPermission struct {
	PermissionsID int64
	AccountsID    int64
}
