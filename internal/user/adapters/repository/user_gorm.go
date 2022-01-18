package repository

import (
	"context"
	"errors"
	"time"

	"github.com/dapr-ddd-action/internal/user/adapters/converter"
	"github.com/dapr-ddd-action/internal/user/domain/aggregate"
	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/dapr-ddd-action/pkg/util/pagination"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 使用 gorm 实现User 的 CURD

// ListUsersPage 分页查询 user
func (u userRepo) ListUsersPage(ctx context.Context, pageNum int, pageSize int) (userDOs []*aggregate.User, err error) {
	user := u.sqlClient.User
	userPOs, err := user.WithContext(ctx).Limit(pageSize).Offset(pagination.GetPageOffset(pageNum, pageSize)).Find()
	if err != nil {
		return
	}

	userDOs = make([]*aggregate.User, 0, len(userPOs))
	for _, item := range userPOs {
		userDOs = append(userDOs, converter.ToUserDO(item))
	}

	return
}

// GetUserByID 查询用户信息
func (u userRepo) GetUserByID(ctx context.Context, id int64) (userDO *aggregate.User, err error) {
	redisKey := userDO.GetUserInfoKey(id)
	// 1. 先查cache
	if userDO, err = u.GetUserFromCache(ctx, redisKey); err == nil {
		return
	}

	if !errorx.IsNotFound(err) {
		u.logger.Error("repo: get user from cache failed", zap.Error(err), zap.Int64("id", id))
	}

	// 2. 再查DB
	user := u.sqlClient.User
	userPO, err := user.WithContext(ctx).Where(user.ID.Eq(int32(id))).Take()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errorx.NotFound("user id = %d", id)
			return
		}
		return
	}
	userDO = converter.ToUserDO(userPO)

	// 3. DB查到后, 回写 redis
	err = u.SaveUserCache(ctx, redisKey, userDO)
	if err != nil {
		u.logger.Error("repository: SaveUserCache write redis failed",
			zap.Error(err), zap.String("redis key=%s", redisKey))
		return nil, err
	}
	return
}

// SaveUser 修改/保存 user
func (u userRepo) SaveUser(ctx context.Context, userDO *aggregate.User) error {
	user := u.sqlClient.User
	userPO := converter.FromUserDO(userDO)
	userPO.UpdateTime = time.Now()
	if userPO.ID == 0 {
		userPO.CreateTime = time.Now()
		if err := user.WithContext(ctx).Create(userPO); err != nil {
			return err
		}
	} else {
		_, err := user.WithContext(ctx).Omit(user.CreateTime).Where(user.ID.Eq(userPO.ID)).Updates(userPO)
		if err != nil {
			return err
		}
		return u.SaveUserCache(ctx, userDO.GetUserInfoKey(userDO.ID), userDO)
	}

	return nil
}
