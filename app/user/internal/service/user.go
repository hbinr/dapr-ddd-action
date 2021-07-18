package service

import (
	"context"

	"github.com/dapr-ddd-action/app/user/internal/repository"
	"github.com/dapr-ddd-action/app/user/internal/repository/po"
	"github.com/dapr-ddd-action/app/user/internal/service/dto"

	"github.com/jinzhu/copier"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo: repo}
}

func (u UserService) GetUser(ctx context.Context, id int64) (resDTO *dto.UserDTO, err error) {
	var userPO *po.UserPO

	if userPO, err = u.repo.QueryUserById(ctx, id); err != nil {
		return
	}
	resDTO = new(dto.UserDTO)
	if err = copier.Copy(resDTO, userPO); err != nil {
		return
	}

	return
}
