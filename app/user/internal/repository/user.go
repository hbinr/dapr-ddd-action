package repository

import (
	"context"
	"fmt"

	"github.com/dapr-ddd-action/app/pkg/constant"
	"github.com/pkg/errors"

	"github.com/dapr-ddd-action/pkg/errorx"

	"github.com/dapr-ddd-action/pkg/jsonx"

	"github.com/dapr-ddd-action/app/user/internal/repository/po"

	"go.uber.org/zap"

	"github.com/dapr-ddd-action/pkg/daprhelp"
	dapr "github.com/dapr/go-sdk/client"
)

var (
	ErrUserNotFound = errorx.NewConvertDataError("user is not found", "user-not-found")
)

type UserRepository interface {
	QueryUserById(ctx context.Context, id int64) (*po.User, error)
	UpdateUser(ctx context.Context, id int64, userName string) error
}

type userRepo struct {
	client dapr.Client
	logger *zap.Logger
}

func NewUserRepo(logger *zap.Logger) (UserRepository, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}
	return &userRepo{
		client: client,
		logger: logger,
	}, nil
}

// QueryUserById 通过 dapr  InvokeBinding API 访问MySQL
func (u *userRepo) QueryUserById(ctx context.Context, id int64) (*po.User, error) {
	// 1. 需要自己拼接SQL, 需要注意注入SQL的风险
	// 2. 另外, 多个SQL一起更新, 事务怎么处理？
	// 暂时没想到解决上述问题更好的方案
	selectSQL := fmt.Sprintf("select * from user where id = %d", id)

	in := daprhelp.BuildBindingRequest(
		constant.MySQLBindName,
		constant.MySQLOperationQuery,
		constant.MySQLMetaDataKey,
		selectSQL,
		nil)

	out, err := u.client.InvokeBinding(ctx, in)

	if err != nil {
		u.logger.Error("QueryUserById failed", zap.Error(err))
		return nil, err
	}
	// 必须这样判断返回的数据是否为空, 因为此处 err 返回为nil, 但是 data 为空. 太丑了~
	if string(out.Data) == "null" {
		u.logger.Error("user not found", zap.Error(err), zap.Int64("id", id))
		return nil, ErrUserNotFound
	}
	// out.Data 返回类型为数组
	var resPO []*po.User
	if err = jsonx.Unmarshal(out.Data, &resPO); err != nil {
		return nil, errors.Wrap(err, "unmarshal user failed")
	}

	return resPO[0], nil
}

func (u *userRepo) UpdateUser(ctx context.Context, id int64, userName string) error {
	updateSQL := fmt.Sprintf(`update user set user_name = '%s' where  id = %d`, userName, id)

	in := daprhelp.BuildBindingRequest(
		constant.MySQLBindName,
		constant.MySQLOperationExec,
		constant.MySQLMetaDataKey,
		updateSQL,
		nil)

	_, err := u.client.InvokeBinding(ctx, in)

	if err != nil {
		u.logger.Error("UpdateUser failed", zap.Error(err))
		return err
	}

	return nil
}
