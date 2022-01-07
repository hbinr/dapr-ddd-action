package repository

import (
	"context"
	"fmt"

	"github.com/dapr-ddd-action/internal/user/domain/data/entity"
	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/dapr-ddd-action/pkg/jsonx"
)

func (u userRepo) SaveUserCache(ctx context.Context, key string, user *entity.User) error {
	storeName := "ddd-action-statestore"

	data, err := jsonx.Marshal(user)
	if err != nil {
		return err
	}

	if err := u.client.SaveState(ctx, storeName, key, data); err != nil {
		return err
	}
	return nil
}

func (u userRepo) GetUserFromCache(ctx context.Context, key string) (userDO *entity.User, err error) {
	storeName := "ddd-action-statestore"

	userDO = new(entity.User)
	item, err := u.client.GetState(ctx, storeName, key)
	if err != nil {
		return
	}

	if item.Value == nil {
		err = errorx.NotFound("redis: GetUser key=%s", key)
		return
	}

	if err = jsonx.Unmarshal(item.Value, userDO); err != nil {
		return
	}
	// 输出示例: data retrieved [key:user:info6 etag:11]: &{ID:6 UserName:redis-test2}
	fmt.Printf("data retrieved [key:%s etag:%s]: %+v\n", item.Key, item.Etag, userDO)
	return
}
