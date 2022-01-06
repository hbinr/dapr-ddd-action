package repository

import (
	"context"
	"fmt"

	"github.com/dapr-ddd-action/internal/user/domain/do"
	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/dapr-ddd-action/pkg/jsonx"
)

func (u userRepo) SaveUserCache(ctx context.Context, user *do.User) error {
	storeName := "ddd-action-statestore"

	data, err := jsonx.Marshal(user)
	if err != nil {
		return err
	}

	if err := u.client.SaveState(ctx, storeName, user.GetUserInfoKey(user.ID), data); err != nil {
		return err
	}
	return nil
}

func (u userRepo) GetUserFromCache(ctx context.Context, id int64) (userDO *do.User, err error) {
	storeName := "ddd-action-statestore"

	userDO = new(do.User)
	item, err := u.client.GetState(ctx, storeName, userDO.GetUserInfoKey(id))
	if err != nil {
		return
	}

	if item.Value == nil {
		err = errorx.NotFound("redis: id=%d", id)
		return
	}

	if err = jsonx.Unmarshal(item.Value, userDO); err != nil {
		return
	}
	// 输出示例: data retrieved [key:user:info6 etag:11]: &{ID:6 UserName:redis-test2}
	fmt.Printf("data retrieved [key:%s etag:%s]: %+v\n", item.Key, item.Etag, userDO)
	return
}
