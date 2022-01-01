package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dapr-ddd-action/internal/user/domain/do"
	"github.com/dapr-ddd-action/pkg/jsonx"
)

func (u userRepo) SaveUserCache(ctx context.Context, user *do.User) error {
	storeName := "ddd-action-statestore"
	key := "user:info" + strconv.Itoa(int(user.ID))

	data, err := jsonx.Marshal(user)
	if err != nil {
		return err
	}

	if err := u.client.SaveState(ctx, storeName, key, data); err != nil {
		return err
	}
	return nil
}

func (u userRepo) GetUserFromCache(ctx context.Context, id int64) (*do.User, error) {
	storeName := "ddd-action-statestore"
	key := "user:info" + strconv.Itoa(int(id))

	item, err := u.client.GetState(ctx, storeName, key)
	if err != nil {
		return nil, err
	}
	if len(item.Value) == 0 {
		return nil, ErrUserNotFound
	}

	user := new(do.User)
	if err = jsonx.Unmarshal(item.Value, user); err != nil {
		return nil, err
	}
	// 输出示例: data retrieved [key:user:info6 etag:11]: &{ID:6 UserName:redis-test2}
	fmt.Printf("data retrieved [key:%s etag:%s]: %+v\n", item.Key, item.Etag, user)
	return user, nil
}
