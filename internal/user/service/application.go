package service

import (
	"github.com/dapr-ddd-action/internal/user/domain"

	"github.com/dapr-ddd-action/internal/user/app/query"

	"github.com/dapr-ddd-action/internal/user/app/command"

	"github.com/dapr-ddd-action/internal/user/app"
)

//func NewApplication(ctx context.Context) (app.Application, func()) {
//	commands := app.Commands{
//		EditUserInfo: command.NewEditUserInfoHandler(nil),
//	}
//	queries := app.Queries{
//		UserInfo:  query.NewUsersInfoHandler(nil),
//		UsersPage: query.NewUsersPageHandler(nil),
//	}
//
//	return app.Application{Commands: commands, Queries: queries},
//		func() {
//			fmt.Println("暂时没有需要关闭的资源, 等调用gRPC服务时，需要关闭 grpc 连接")
//		}
//}

func NewApplication(userRepo domain.UserRepository) app.Application {
	commands := app.Commands{
		EditUserInfo: command.NewEditUserInfoHandler(userRepo),
	}
	queries := app.Queries{
		UserInfo:  query.NewUsersInfoHandler(userRepo),
		UsersPage: query.NewUsersPageHandler(userRepo),
	}

	return app.Application{Commands: commands, Queries: queries}
}
