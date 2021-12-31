package query

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/dapr-ddd-action/internal/user/domain"
)

type UsersPageHandler struct {
	repo domain.UserRepository
}

func NewUsersPageHandler(repo domain.UserRepository) UsersPageHandler {
	if repo == nil {
		panic("nil domain.UserRepository")
	}

	return UsersPageHandler{repo}
}

func (u UsersPageHandler) Handler(ctx context.Context, pageNum, pageSize int) ([]User, error) {
	usersDO, err := u.repo.ListUsersPage(ctx, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	res := make([]User, 0, len(usersDO))
	if err = copier.Copy(res, usersDO); err != nil {
		return nil, err
	}

	return res, nil
}
