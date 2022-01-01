package domain

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/domain/do"
)

type UserRepository interface {
	// ListUsersPage 分页查询 user
	ListUsersPage(ctx context.Context, pageNum, pageSize int) ([]do.User, error)
	GetUserById(context.Context, int64) (*do.User, error)
	// GetUserFromCache 获取 user(查询缓存)
	GetUserFromCache(context.Context, int64) (*do.User, error)

	// SaveUserCache 保存 user(缓存)
	SaveUserCache(context.Context, *do.User) error
	// UpdateUser 修改 user
	UpdateUser(context.Context, *do.User) error
}
