package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dapr-ddd-action/internal/pkg/constant/e"

	"github.com/dapr-ddd-action/internal/service/dto"

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
			var (
				id      int
				data    []byte
				userDto *dto.UserDTO
			)
			if id, err = strconv.Atoi(idStr); err != nil {
				return nil, err
			}

			userDto, err = u.service.GetUser(ctx, id)

			switch err {
			case nil:
				if data, err = json.Marshal(userDto); err != nil {
					return nil, err
				}

				out.Data = data
				out.ContentType = "application/json"
			case e.ErrNotFound:
				return nil, e.ErrNotFound
			default:
				return nil, err
			}
		}
	}
	return
}
