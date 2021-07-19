package controller

import (
	"context"

	"github.com/dapr-ddd-action/app/pkg/constant/e"

	"github.com/dapr-ddd-action/app/user/internal/service"

	"github.com/dapr-ddd-action/pkg/ginx"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func RegisterUserRouter(r *gin.Engine, us service.UserService) {
	ctl := UserController{service: us}
	r.GET("/user/:id", ctl.GetUser)
}

func (u UserController) GetUser(c *gin.Context) {
	id, err := ginx.QueryInt("id", c)

	if err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}

	userDto, err := u.service.GetUser(context.Background(), id)

	switch err {
	case nil:
		ginx.RespSuccess(c, userDto)
	case e.ErrUserNotExist:
		ginx.RespError(c, e.CodeUserNotExist)
	case e.ErrConvDataErr:
		ginx.RespError(c, e.CodeConvDataErr)
	default:
		ginx.RespError(c, e.CodeError)
	}
}
