package query

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/domain"
	"github.com/jinzhu/copier"
)

type UserInfoHandler struct {
	repo domain.UserRepository
}

func NewUsersInfoHandler(repo domain.UserRepository) UserInfoHandler {
	if repo == nil {
		panic("nil domain.UserRepository")
	}

	return UserInfoHandler{repo}
}

func (u UserInfoHandler) Handler(ctx context.Context, id int64) (User, error) {
	// usersDO, err := u.repo.GetUserById(ctx, id)

	usersDO, err := u.repo.GetUserFromCache(ctx, id)

	if err != nil {
		return User{}, err
	}
	var res User
	if err = copier.Copy(&res, usersDO); err != nil {
		return User{}, err
	}

	return res, nil
}
