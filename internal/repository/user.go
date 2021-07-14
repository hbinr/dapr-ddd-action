package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/dapr-ddd-action/pkg/daprhelp"

	"github.com/dapr-ddd-action/internal/repository/po"
	dapr "github.com/dapr/go-sdk/client"
)

type UserRepository interface {
	QueryUserById(ctx context.Context, id int) (*po.UserPO, error)
}

type userRepo struct {
	client dapr.Client
}

func NewUserRepo() (UserRepository, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}
	return &userRepo{client: client}, nil
}

// QueryUserById 通过 dapr  InvokeBinding API 访问MySQL
func (u *userRepo) QueryUserById(ctx context.Context, id int) (*po.UserPO, error) {
	querySQL := fmt.Sprintf(`select * from user where id = %d`, id)

	in := daprhelp.BuildBindingRequest(
		"dapr-ddd-action-mysql",
		"query",
		"sql",
		querySQL,
		nil)

	out, err := u.client.InvokeBinding(ctx, in)

	if err != nil {
		return nil, errors.Wrapf(err, "repository: QueryUserById faild, sql:[%s]", querySQL)
	}

	var resPO []*po.UserPO

	if err = json.Unmarshal(out.Data, &resPO); err != nil {
		fmt.Println("err:", err)

		return nil, err
	}

	return resPO[0], nil
}
