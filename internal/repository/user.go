package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dapr-ddd-action/internal/pkg/constant/e"

	"github.com/dapr-ddd-action/internal/pkg/constant/daprc"
	"github.com/pkg/errors"

	"github.com/dapr-ddd-action/internal/repository/po"
	"github.com/dapr-ddd-action/pkg/daprhelp"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/kit/logger"
)

type UserRepository interface {
	QueryUserById(ctx context.Context, id int) (*po.UserPO, error)
}

type userRepo struct {
	client dapr.Client
	logger logger.Logger
}

func NewUserRepo(logger logger.Logger) (UserRepository, error) {
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
func (u *userRepo) QueryUserById(ctx context.Context, id int) (*po.UserPO, error) {
	querySQL := fmt.Sprintf(`select * from user where id = %d`, id)

	in := daprhelp.BuildBindingRequest(
		daprc.DaprMySQLBindName,
		daprc.DaprMySQLOperation,
		daprc.DaprMySQLMetaDataKey,
		querySQL,
		nil)

	out, err := u.client.InvokeBinding(ctx, in)

	u.logger.Infof("out :%s\n", string(out.Data))

	var resPO []*po.UserPO

	if err != nil {
		return nil, errors.Wrapf(err, "repository: QueryUserById faild, sql:[%s]", querySQL)
	}

	// ????? 必须这样判断返回的数据是否为空吗？ 太丑了~
	if string(out.Data) == "null" {
		return nil, e.ErrNotFound
	}

	if err = json.Unmarshal(out.Data, &resPO); err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	return resPO[0], nil

}
