package controller

import (
	"context"
	"net/http"

	"github.com/dapr-ddd-action/app/pkg/constant/e"

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

	id, err := chix.QueryInt64("idx", r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		chix.RespError(w, r, e.CodeInvalidParams)
	}
	userDto, _ := u.service.GetUser(context.Background(), id)
	chix.RespSuccess(w, r, userDto)
}

type UpdateUserReq struct {
	Id       int64  `query:"id,required"`
	UserName string `json:"userName,required"  vd:"len($)>5; msg:sprintf('用户名必须大于5个字符')"`
}

func (u UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req := new(UpdateUserReq)
	if err := chix.BindAndValidate(req, r); err != nil {
		render.Respond(w, r, err)
		return
	}

	if err := u.service.UpdateUser(context.Background(), req.Id, req.UserName); err != nil {
		render.Respond(w, r, "failed")
	}
	render.Respond(w, r, "OK")
}
