package command

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/dapr-ddd-action/internal/user/domain/aggregate"

	"github.com/dapr-ddd-action/internal/user/domain"
)

// 入参 dto -> do
// 出参 do -> dto

// EditUserInfoHandler 业务编排
type EditUserInfoHandler struct {
	service domain.UserService
}

func NewEditUserInfoHandler(service domain.UserService) EditUserInfoHandler {

	return EditUserInfoHandler{service}
}

func (e EditUserInfoHandler) Handler(ctx context.Context, userDto EditUserInfo) error {
	userDO := new(aggregate.User)
	if err := copier.Copy(userDO, userDto); err != nil {
		return err
	}

	if err := e.service.UpdateUser(ctx, userDO); err != nil {
		return err
	}

	// if err := e.repo.SaveUserCache(ctx, userDO); err != nil {
	// 	return err
	// }

	return nil
}
