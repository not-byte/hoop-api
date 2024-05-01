package model

import (
	"fmt"
	"time"
	"tournament_api/server/utils"
)

type Account struct {
	Email         *string
	Password      *string
	ID            int64
	PermissionsID int64
	CreatedOn     time.Time
	LoggedOn      time.Time
	Verified      bool
	MailToken     int8
}

func (a Account) String() string {
	return fmt.Sprintf("ID: %d, PermissionsID: %d, Email: %s, Password: %s, CreatedOn: %v, LoggedOn: %v, Verified: %v, MailToken: %v",
		a.ID, a.PermissionsID, utils.DerefString(a.Email), utils.DerefString(a.Password), a.CreatedOn, a.LoggedOn, a.Verified, a.MailToken)
}

type AccountPermission struct {
	PermissionsID int64
	AccountsID    int64
}
