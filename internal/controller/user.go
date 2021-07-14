package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dapr-ddd-action/pkg/daprhelp"

	"github.com/dapr/go-sdk/service/common"

	"github.com/dapr-ddd-action/internal/service"
)

type UserController struct {
	service service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		service: userService,
	}
}

// GetUser http service
func (u UserController) GetUser(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	out = new(common.Content)
	if in.Verb == http.MethodGet {
		// in.QueryString 值为 name=value&name2=value2, 需要自己解析
		idStr := daprhelp.GetQuery(in.QueryString, "id")
		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			_ = id
			if err != nil {
				out.Data = []byte(err.Error())
			}

			userDto, _ := u.service.GetUser(ctx, id)

			data, _ := json.Marshal(userDto)
			out.Data = data
		}
	}
	return
}
