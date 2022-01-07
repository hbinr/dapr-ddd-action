package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/dapr-ddd-action/internal/pkg/constant"
	"github.com/dapr-ddd-action/internal/user/domain"
	"github.com/dapr-ddd-action/internal/user/domain/data/entity"

	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/dapr-ddd-action/pkg/jsonx"

	"go.uber.org/zap"

	"github.com/dapr-ddd-action/pkg/daprhelp"
	dapr "github.com/dapr/go-sdk/client"
)

// 入参 po
// 响应 po

type userRepo struct {
	client dapr.Client
	logger *zap.Logger
}

func NewUserRepo(client dapr.Client, logger *zap.Logger) domain.UserRepository {
	return &userRepo{client, logger}
}

// GetUserById 通过 dapr  InvokeBinding API 访问MySQL
func (u userRepo) GetUserById(ctx context.Context, id int64) (*entity.User, error) {
	// 1. 需要自己拼接SQL, 需要注意注入SQL的风险
	// 2. 另外, 多个SQL一起更新, 事务怎么处理？
	// 暂时没想到解决上述问题更好的方案
	selectSQL := fmt.Sprintf("select * from user where id = %d", id)

	in := daprhelp.BuildMySQLQueryBinding(constant.MySQLBindName, selectSQL)

	out, err := u.client.InvokeBinding(ctx, in)

	if err != nil {
		u.logger.Error("QueryUserById failed", zap.Error(err))
		return nil, err
	}
	// 必须这样判断返回的数据是否为空, 因为此处 err 返回为nil, 但是 data 为空. 太丑了~
	if string(out.Data) == "null" {
		u.logger.Error("user not found", zap.Error(err), zap.Int64("id", id))
		return nil, errorx.NotFound("mysql: id=%d", id)
	}
	// out.Data 返回类型为数组
	var resPO []*entity.User
	if err = jsonx.Unmarshal(out.Data, &resPO); err != nil {
		return nil, errors.Wrap(err, "unmarshal user failed")
	}

	return resPO[0], nil
}

func (u userRepo) UpdateUser(ctx context.Context, user *entity.User) error {
	updateSQL := fmt.Sprintf(`update user set user_name = '%s' where  id = %d`, user.UserName, user.ID)

	in := daprhelp.BuildMySQLExecBinding(constant.MySQLBindName, updateSQL)

	_, err := u.client.InvokeBinding(ctx, in)

	if err != nil {
		u.logger.Error("UpdateUser failed", zap.Error(err), zap.String("sql", updateSQL))
		return errorx.Internal(err, "UpdateUser failed")
	}

	return nil
}

func (u userRepo) ListUsersPage(ctx context.Context, pageNum, pageSize int) ([]entity.User, error) {
	panic("implement me")
}