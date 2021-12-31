package domain

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/domain/do"
)

type UserRepository interface {
	// ListUsersPage 分页查询 user
	ListUsersPage(ctx context.Context, pageNum, pageSize int) ([]do.User, error)
	GetUserById(context.Context, int64) (*do.User, error)

	// UpdateUser 修改 user
	UpdateUser(context.Context, *do.User) error
}
