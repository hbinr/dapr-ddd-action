package command

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/dapr-ddd-action/internal/user/domain/do"

	"github.com/dapr-ddd-action/internal/user/domain"
)

// EditUserInfoHandler 业务编排
type EditUserInfoHandler struct {
	repo domain.UserRepository
}

func NewEditUserInfoHandler(repo domain.UserRepository) EditUserInfoHandler {
	if repo == nil {
		panic("nil domain.UserRepository")
	}

	return EditUserInfoHandler{repo}
}

func (e EditUserInfoHandler) Handler(ctx context.Context, userDto EditUserInfo) error {
	userDO := new(do.User)
	if err := copier.Copy(userDO, userDto); err != nil {
		return err
	}

	if err := e.repo.UpdateUser(ctx, userDO); err != nil {
		return err
	}

	// if err := e.repo.SaveUserCache(ctx, userDO); err != nil {
	// 	return err
	// }

	return nil
}
