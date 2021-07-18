package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"

	"github.com/dapr-ddd-action/app/user/service/internal/repository/po"

	"github.com/dapr-ddd-action/app/pkg/constant/daprc"
	"github.com/dapr-ddd-action/app/pkg/constant/e"

	"github.com/pkg/errors"

	"github.com/dapr-ddd-action/pkg/daprhelp"
	dapr "github.com/dapr/go-sdk/client"
)

type UserRepository interface {
	QueryUserById(ctx context.Context, id int64) (*po.UserPO, error)
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
func (u *userRepo) QueryUserById(ctx context.Context, id int64) (*po.UserPO, error) {
	querySQL := fmt.Sprintf(`select * from user where id = %d`, id)

	in := daprhelp.BuildBindingRequest(
		daprc.DaprMySQLBindName,
		daprc.DaprMySQLOperation,
		daprc.DaprMySQLMetaDataKey,
		querySQL,
		nil)

	out, err := u.client.InvokeBinding(ctx, in)

	var resPO []*po.UserPO

	if err != nil {
		return nil, errors.Wrapf(err, "repository: QueryUserById faild, sql:[%s]", querySQL)
	}

	// ????? 必须这样判断返回的数据是否为空吗？ 太丑了~
	if string(out.Data) == "null" {
		u.logger.Error("user not found", zap.Error(err))
		return nil, e.ErrUserNotExist
	}

	if err = json.Unmarshal(out.Data, &resPO); err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	return resPO[0], nil

}
