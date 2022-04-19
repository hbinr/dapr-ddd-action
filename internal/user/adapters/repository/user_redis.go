package repository

import (
	"context"

	"github.com/dapr-ddd-action/internal/pkg/constants"
	"github.com/dapr-ddd-action/internal/user/domain/aggregate"
	"go.uber.org/zap"

	"github.com/dapr-ddd-action/pkg/daprhelp"
	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/dapr-ddd-action/pkg/jsonx"
)

func (u userRepo) SaveUserCache(ctx context.Context, key string, userDO *aggregate.User) error {
	stateItem, err := daprhelp.BuildExpireStateItem(userDO.GetUserInfoKey(userDO.ID), userDO, constants.ExpireUserInfo)
	if err != nil {
		u.logger.Error("repository: GetUserByID write redis failed", zap.Error(err))
		return err
	}

	return u.client.SaveBulkState(ctx, constants.StateStoreName, stateItem)
}

func (u userRepo) GetUserFromCache(ctx context.Context, key string) (userDO *aggregate.User, err error) {
	userDO = new(aggregate.User)
	item, err := u.client.GetState(ctx, constants.StateStoreName, key, nil)
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
	return
}
