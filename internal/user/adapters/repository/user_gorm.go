package repository

import (
	"context"
	"errors"
	"time"

	"github.com/dapr-ddd-action/internal/user/adapters/repository/data/po"
	"github.com/dapr-ddd-action/internal/user/domain/aggregate"
	"github.com/dapr-ddd-action/pkg/util/pagination"
	"github.com/jinzhu/copier"

	"github.com/dapr-ddd-action/pkg/errorx"
	"gorm.io/gorm"
)

// adapter 依赖 domain

// 使用 gorm 实现User 的 CURD

// ListUsersPage 分页查询 user
func (u userRepo) ListUsersPage(ctx context.Context, pageNum int, pageSize int) (userDOs []*aggregate.User, err error) {
	user := u.sqlClient.User
	userPOs, err := user.WithContext(ctx).Limit(pageSize).Offset(pagination.GetPageOffset(pageNum, pageSize)).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(userDOs, userPOs); err != nil {
		return
	}
	return
}

func (u userRepo) GetUserById(ctx context.Context, id int64) (userDO *aggregate.User, err error) {
	user := u.sqlClient.User
	userPO, err := user.WithContext(ctx).Where(user.ID.Eq(id)).Take()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errorx.NotFound("user id = %d", id)
			return
		}
		return
	}

	userDO = new(aggregate.User)
	if err = copier.Copy(userDO, userPO); err != nil {
		return
	}

	return
}

// SaveUser 修改/保存 user
func (u userRepo) SaveUser(ctx context.Context, userDO *aggregate.User) error {
	user := u.sqlClient.User
	userPO := new(po.User)

	if err := copier.Copy(userPO, userDO); err != nil {
		return err
	}

	userPO.UpdatedAt = time.Now()
	return user.WithContext(ctx).Where(user.ID.Eq(userPO.ID)).Save(userPO)
}
