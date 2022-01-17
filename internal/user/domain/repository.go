package domain

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/domain/aggregate"
)

// 入参、出参 都是 domain object
// 入参需要将 do 转为 po, 然后对po进行持久化
// 出参需要将 po 转为 do

// 依赖倒置的Repo接口 外部adapter负责实现

type UserRepository interface {
	// ListUsersPage 分页查询 user
	ListUsersPage(ctx context.Context, pageNum, pageSize int) ([]*aggregate.User, error)
	GetUserByID(context.Context, int64) (*aggregate.User, error)
	// SaveUser 修改 user
	SaveUser(context.Context, *aggregate.User) error

	// GetUserFromCache 获取 user(查询缓存)
	// Dapr 底层调用 GET 请求: http://127.0.0.1:3500/v1.0/state/ddd-action-statestore/user:info6
	// key: user:info6
	GetUserFromCache(context.Context, string) (*aggregate.User, error)

	// SaveUserCache 保存 user(缓存)
	// Dapr 底层调用 POST 请求 http://127.0.0.1:3500/v1.0/state/ddd-action-statestore
	// key: user:info6, data: 为代码中的业务逻辑组成的数据，是个数组，示例如下:
	// [{
	// 	"key":"user:info6",
	// 	"value": {
	// 		 "id": 6,
	// 		 "user_name": "redis-test333"
	// 	}
	// }]
	SaveUserCache(context.Context, string, *aggregate.User) error
}
