package service

import (
	"context"

	"github.com/dapr-ddd-action/app/pkg/constant/e"

	"go.uber.org/zap"

	"github.com/dapr-ddd-action/app/user/internal/repository"
	"github.com/dapr-ddd-action/app/user/internal/repository/po"
	"github.com/dapr-ddd-action/app/user/internal/service/dto"

	"github.com/jinzhu/copier"
)

type UserService struct {
	repo   repository.UserRepository
	logger *zap.Logger
}

func NewUserService(repo repository.UserRepository,
	logger *zap.Logger) UserService {
	return UserService{repo: repo, logger: logger}
}

func (u UserService) GetUser(ctx context.Context, id int64) (resDTO *dto.UserDTO, err error) {
	var userPO *po.UserPO

	if userPO, err = u.repo.QueryUserById(ctx, id); err != nil {
		return
	}

	resDTO = new(dto.UserDTO)
	if err = copier.Copy(resDTO, userPO); err != nil {
		err = e.ErrConvDataErr
		return
	}

	return
}

func (u UserService) UpdateUser(ctx context.Context, id int64, userName string) error {

	return u.repo.UpdateUser(ctx, id, userName)

}
