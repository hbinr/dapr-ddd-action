package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"

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

func RegisterUserRouterChi(mux *chi.Mux, us service.UserService) {
	ctl := UserController{service: us}
	mux.Get("/user/{id}", ctl.GetUserChi)
}
func (u UserController) GetUserChi(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt(intStr) failed", zap.Error(err), zap.String("param", idStr))
		return
	}
	userDto, err := u.service.GetUser(context.Background(), id)
	render.Respond(w, r, userDto)
}
func (u UserController) GetUser(c *gin.Context) {
	id, err := ginx.QueryInt("id", c)

	if err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}

	userDto, err := u.service.GetUser(context.Background(), id)
	ginx.RespSuccess(c, userDto)

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
