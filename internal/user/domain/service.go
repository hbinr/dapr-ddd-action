package domain

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/domain/aggregate"
	"github.com/dapr-ddd-action/internal/user/domain/data/po"
	"github.com/jinzhu/copier"
)

// 入参 do -> po
// 出参 po -> do

// UserService 用户领域服务
type UserService struct {
	repo UserRepository
}

func NewUserDomain(repo UserRepository) UserService {
	return UserService{repo}
}

// ListUsersPage 分页查询 user
func (u *UserService) ListUsersPage(ctx context.Context, pageNum int, pageSize int) ([]aggregate.User, error) {
	panic("not implemented")
}

func (u *UserService) GetUserById(ctx context.Context, id int64) (*aggregate.User, error) {
	userPO, err := u.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	userDO := &aggregate.User{}
	if err = copier.Copy(userDO, userPO); err != nil {
		return nil, err
	}

	return userDO, nil
}

// GetUserFromCache 获取 user(查询缓存)
// Dapr 底层调用 GET 请求: http://127.0.0.1:3500/v1.0/state/ddd-action-statestore/user:info6
// key: user:info6
func (u *UserService) GetUserFromCache(ctx context.Context, id int64) (*aggregate.User, error) {
	userDO := &aggregate.User{}

	userPO, err := u.repo.GetUserFromCache(ctx, userDO.GetUserInfoKey(id))
	if err != nil {
		return nil, err
	}
	if err = copier.Copy(userDO, userPO); err != nil {
		return nil, err
	}

	return userDO, nil
}

// UpdateUser 修改 user

func (u *UserService) UpdateUser(ctx context.Context, userDO *aggregate.User) error {
	userPO := &po.User{}
	if err := copier.Copy(userPO, userDO); err != nil {
		return err
	}

	return u.repo.SaveUser(ctx, userPO)
}

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
func (u *UserService) SaveUser(ctx context.Context, userDO *aggregate.User) error {
	userPO := &po.User{}
	if err := copier.Copy(userPO, userDO); err != nil {
		return err
	}

	return u.repo.SaveUserCache(ctx, userDO.GetUserInfoKey(userDO.ID), userPO)
}
