package app

import (
	"github.com/dapr-ddd-action/internal/user/app/command"
	"github.com/dapr-ddd-action/internal/user/app/query"
)

//	req -> po
//	po -> res

type Application struct {
	Commands Commands
	Queries  Queries
}
type Commands struct {
	EditUserInfo command.EditUserInfoHandler
}

type Queries struct {
	UserInfo  query.UserInfoHandler
	UsersPage query.UsersPageHandler
}
