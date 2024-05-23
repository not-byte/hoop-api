package model

import (
	"math/big"
	"time"
)

type Account struct {
	Email         string
	Password      string
	ID            big.Int
	PermissionsID big.Int
	CreatedOn     time.Time
	LoggedOn      time.Time
	Verified      bool
	MailToken     int8
}

type AccountPermission struct {
	PermissionsID big.Int
	AccountsID    big.Int
}
