package ports

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/dapr-ddd-action/internal/user/app/command"

	"github.com/dapr-ddd-action/internal/user/app"

	"github.com/dapr-ddd-action/pkg/httpx"
)

type UserController struct {
	app app.Application
}

func RegisterUserRouter(r *mux.Router, app app.Application) {
	ctl := UserController{app}
	r.HandleFunc("/user/{id}", ctl.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/user", ctl.UpdateUser).Methods(http.MethodPut)
}

func (u UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := httpx.QueryInt64("id", r)
	if err != nil {
		httpx.BadRequest("invalid-id", "invalid-id", err, w)
		return
	}

	userDto, err := u.app.Queries.UserInfo.Handler(r.Context(), id)
	if err != nil {
		httpx.RespWithError(err, "服务器开小差了，请稍后再试", w)
		return
	}
	httpx.RespSuccess(userDto, w)
}

func (u UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req := new(UpdateUserReq)
	if err := httpx.BindAndValidate(req, r); err != nil {
		// 特殊处理 go-tagexpr  参数验证错误, 友好提示给用户
		msg := strings.Split(err.Error(), "cause=")[1]
		httpx.BadRequest("invalid-request", msg, err, w)
		return
	}
	if req.ID == 0 {
		httpx.BadRequest("invalid-request", "invalid id", nil, w)
		return
	}

	editUserInfo := command.EditUserInfo{
		ID:       req.ID,
		UserName: req.UserName,
	}

	if err := u.app.Commands.EditUserInfo.Handler(r.Context(), editUserInfo); err != nil {
		httpx.RespWithError(err, "服务器开小差了，请稍后再试", w)
		return
	}

	httpx.RespSuccess(nil, w)
}
