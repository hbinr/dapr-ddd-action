package query

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/app/assemble"
	"github.com/dapr-ddd-action/internal/user/domain"
)

type UserInfoHandler struct {
	service domain.UserService
}

func NewUsersInfoHandler(service domain.UserService) UserInfoHandler {
	return UserInfoHandler{service}
}

func (u UserInfoHandler) Handler(ctx context.Context, id int64) (res assemble.UserDTO, err error) {
	usersDO, err := u.service.GetUserById(ctx, id)
	// usersDO, err := u.service.GetUserFromCache(ctx, id)

	if err != nil {
		return
	}

	res = assemble.UserToDTO(usersDO)
	return
}
