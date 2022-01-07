package query

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/domain"
	"github.com/jinzhu/copier"
)

type UserInfoHandler struct {
	service domain.UserService
}

func NewUsersInfoHandler(service domain.UserService) UserInfoHandler {
	return UserInfoHandler{service}
}

func (u UserInfoHandler) Handler(ctx context.Context, id int64) (User, error) {
	usersDO, err := u.service.GetUserById(ctx, id)

	// usersDO, err := u.service.GetUserFromCache(ctx, id)

	if err != nil {
		return User{}, err
	}
	var res User
	if err = copier.Copy(&res, usersDO); err != nil {
		return User{}, err
	}

	return res, nil
}
