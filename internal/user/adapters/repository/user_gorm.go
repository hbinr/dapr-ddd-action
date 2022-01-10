package repository

import (
	"context"
	"errors"
	"time"

	"github.com/dapr-ddd-action/internal/user/adapters/repository/data/po"

	"github.com/dapr-ddd-action/pkg/util/pagination"

	"github.com/dapr-ddd-action/pkg/errorx"
	"gorm.io/gorm"
)

// 使用 gorm 实现User 的 CURD

// ListUsersPage 分页查询 user
func (u userRepo) ListUsersPage(ctx context.Context, pageNum int, pageSize int) ([]*po.User, error) {
	user := u.sqlClient.User
	return user.WithContext(ctx).Limit(pageSize).Offset(pagination.GetPageOffset(pageNum, pageSize)).Find()
}

func (u userRepo) GetUserById(ctx context.Context, id int64) (*po.User, error) {
	user := u.sqlClient.User
	userPO, err := user.WithContext(ctx).Where(user.ID.Eq(id)).Take()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.NotFound("user id = %d", id)
		}
		return nil, err
	}

	return userPO, nil
}

// SaveUser 修改/保存 user
func (u userRepo) SaveUser(ctx context.Context, userPO *po.User) error {
	user := u.sqlClient.User
	userPO.UpdatedAt = time.Now()
	return user.WithContext(ctx).Where(user.ID.Eq(userPO.ID)).Save(userPO)
}
