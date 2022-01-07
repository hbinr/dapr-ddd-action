package query

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/dapr-ddd-action/internal/user/domain"
)

// 入参 dto -> do
// 出参 do -> dto
type UsersPageHandler struct {
	service domain.UserService
}

func NewUsersPageHandler(service domain.UserService) UsersPageHandler {
	return UsersPageHandler{service}
}

func (u UsersPageHandler) Handler(ctx context.Context, pageNum, pageSize int) ([]User, error) {
	usersDO, err := u.service.ListUsersPage(ctx, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	res := make([]User, 0, len(usersDO))
	if err = copier.Copy(res, usersDO); err != nil {
		return nil, err
	}

	return res, nil
}
