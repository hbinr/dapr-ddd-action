package query

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/app/assemble"
	"github.com/dapr-ddd-action/internal/user/domain"
)

// 入参 query 参数尽量是struct，提高扩展性，除了只根据 id 或者 uuid 查询的请求(只有单个参数)
// 出参 do -> dto
type UsersPageHandler struct {
	repo domain.UserRepository
}

func NewUsersPageHandler(repo domain.UserRepository) UsersPageHandler {
	return UsersPageHandler{repo}
}

func (u UsersPageHandler) Handler(ctx context.Context, query *UsersPageQuery) ([]assemble.UserDTO, error) {
	userDO, err := u.repo.ListUsersPage(ctx, query.GetCurrentPage(), query.GetPageSize())
	if err != nil {
		return nil, err
	}
	res := make([]assemble.UserDTO, 0, len(userDO))

	for _, user := range userDO {
		res = append(res, assemble.ToUserDTO(user))
	}

	return res, nil
}
