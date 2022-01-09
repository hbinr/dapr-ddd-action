package query

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/app/assemble"
	"github.com/dapr-ddd-action/internal/user/domain"
)

// 入参 query
// 出参 do -> dto
type UsersPageHandler struct {
	service domain.UserService
}

func NewUsersPageHandler(service domain.UserService) UsersPageHandler {
	return UsersPageHandler{service}
}

func (u UsersPageHandler) Handler(ctx context.Context, query UsersPageQuery) ([]assemble.UserDTO, error) {
	usersDO, err := u.service.ListUsersPage(ctx, query.CurrentPage, query.PageSize)
	if err != nil {
		return nil, err
	}
	res := make([]assemble.UserDTO, 0, len(usersDO))

	for _, user := range usersDO {
		res = append(res, assemble.UserToDTO(&user))
	}

	return res, nil
}
