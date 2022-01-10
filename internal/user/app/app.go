package app

import (
	"github.com/dapr-ddd-action/internal/user/app/command"
	"github.com/dapr-ddd-action/internal/user/app/query"
	"github.com/dapr-ddd-action/internal/user/domain"
)

// Application 使用CQRS架构实现，该层主要是进行服务的编排，调 domain 层的领域服务

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

func NewApplication(repo domain.UserRepository) Application {
	commands := Commands{
		EditUserInfo: command.NewEditUserInfoHandler(repo),
	}
	queries := Queries{
		UserInfo:  query.NewUsersInfoHandler(repo),
		UsersPage: query.NewUsersPageHandler(repo),
	}

	return Application{Commands: commands, Queries: queries}
}
