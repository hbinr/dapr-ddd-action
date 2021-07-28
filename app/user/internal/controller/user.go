package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/dapr-ddd-action/pkg/errorx/httperr"

	"github.com/dapr-ddd-action/pkg/chix"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5"

	"github.com/dapr-ddd-action/app/user/internal/service"
)

type UserController struct {
	service service.UserService
}

func RegisterUserRouter(mux *chi.Mux, us service.UserService) {
	ctl := UserController{service: us}
	mux.Get("/user/{id}", ctl.GetUser)
	mux.Put("/user", ctl.UpdateUser)
}

func (u UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// go-tagexpr 的  path 参数需要实现 PathParams 接口的 Get 方法, 比较复杂, 直接用 chi,URLParam()省事

	id, err := chix.QueryInt64("id", r)
	if err != nil {
		httperr.BadRequest("invalid-id", "", err, w, r)
		return
	}

	userDto, err := u.service.GetUser(context.Background(), id)
	if err != nil {
		httperr.RespondWithError(err, "", w, r)
		return
	}

	render.Respond(w, r, userDto)
}

type UpdateUserReq struct {
	Id       int64  `query:"id,required"`
	UserName string `json:"userName,required"  vd:"len($)>5; msg:sprintf('用户名[%v]必须大于5个字符',$)"`
}

func (u UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req := new(UpdateUserReq)
	if err := chix.BindAndValidate(req, r); err != nil {
		// 特殊处理 go-tagexpr  参数验证错误, 友好提示给用户
		msg := strings.Split(err.Error(), "cause=")[1]

		httperr.BadRequest("invalid-request", msg, err, w, r)
		return
	}

	if err := u.service.UpdateUser(context.Background(), req.Id, req.UserName); err != nil {
		httperr.RespondWithError(err, "", w, r)
		return
	}

	render.Respond(w, r, "OK")
}
