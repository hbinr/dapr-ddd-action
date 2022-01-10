package command

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/dapr-ddd-action/internal/user/domain/aggregate"

	"github.com/dapr-ddd-action/internal/user/domain"
)

// 入参 cmd ，参数尽量是struct，提高扩展性，除了只根据 id 或者 uuid 查询的请求(只有单个参数)
// 出参 do -> dto

// EditUserInfoHandler 业务编排
type EditUserInfoHandler struct {
	repo domain.UserRepository
}

func NewEditUserInfoHandler(repo domain.UserRepository) EditUserInfoHandler {

	return EditUserInfoHandler{repo}
}

func (e EditUserInfoHandler) Handler(ctx context.Context, cmd *EditUserInfoCmd) error {
	userDO := new(aggregate.User)
	if err := copier.Copy(userDO, cmd); err != nil {
		return err
	}

	if err := e.repo.SaveUser(ctx, userDO); err != nil {
		return err
	}

	// if err := e.repo.SaveUserCache(ctx, userDO); err != nil {
	// 	return err
	// }

	return nil
}
