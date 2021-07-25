package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"

	"github.com/dapr-ddd-action/app/user/internal/service"
)

type UserController struct {
	service service.UserService
}

func RegisterUserRouter(mux *chi.Mux, us service.UserService) {
	ctl := UserController{service: us}
	mux.Get("/user/{id}", ctl.GetUser)
}

func (u UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt(intStr) failed", zap.Error(err), zap.String("param", idStr))
		return
	}
	userDto, err := u.service.GetUser(context.Background(), id)
	render.Respond(w, r, userDto)
}
