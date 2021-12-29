package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/dapr-ddd-action/pkg/errorx/httperr"
	"github.com/gorilla/mux"

	"github.com/dapr-ddd-action/pkg/httpx"

	"github.com/dapr-ddd-action/app/user/internal/service"
)

type UserController struct {
	service service.UserService
}

func RegisterUserRouter(r *mux.Router, us service.UserService) {
	ctl := UserController{service: us}
	r.HandleFunc("/user/{id}", ctl.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/user", ctl.UpdateUser).Methods(http.MethodPut)
}

func (u UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// 	varsMap := mux.Vars(r)
	// actorType := varsMap["actorType"]
	id, err := httpx.QueryInt64("id", r)
	if err != nil {
		httperr.BadRequest("invalid-id", "", err, w, r)
		return
	}

	userDto, err := u.service.GetUser(context.Background(), id)
	if err != nil {
		httperr.RespondWithError(err, "", w, r)
		return
	}
	httpx.RespSuccess(userDto, w)
	// render.Respond(w, r, userDto)
}

type UpdateUserReq struct {
	Id       int64  `query:"id,required"`
	UserName string `json:"userName,required"  vd:"len($)>5; msg:sprintf('用户名[%v]必须大于5个字符',$)"`
}

func (u UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req := new(UpdateUserReq)
	if err := httpx.BindAndValidate(req, r); err != nil {
		// 特殊处理 go-tagexpr  参数验证错误, 友好提示给用户
		msg := strings.Split(err.Error(), "cause=")[1]

		httperr.BadRequest("invalid-request", msg, err, w, r)
		return
	}

	if err := u.service.UpdateUser(context.Background(), req.Id, req.UserName); err != nil {
		httperr.RespondWithError(err, "", w, r)
		return
	}

	httpx.RespSuccess("OK", w)
}
