package query

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/app/assemble"
	"github.com/dapr-ddd-action/internal/user/domain"
)

type UserInfoHandler struct {
	repo domain.UserRepository
}

func NewUsersInfoHandler(repo domain.UserRepository) UserInfoHandler {
	return UserInfoHandler{repo}
}

func (u UserInfoHandler) Handler(ctx context.Context, id int64) (res assemble.UserDTO, err error) {
	usersDO, err := u.repo.GetUserById(ctx, id)

	if err != nil {
		return
	}

	res = assemble.ToUserDTO(usersDO)
	return
}
