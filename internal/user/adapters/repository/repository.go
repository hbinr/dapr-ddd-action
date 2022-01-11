package repository

import (
	"github.com/dapr-ddd-action/internal/user/adapters/repository/data/dao"
	"github.com/dapr-ddd-action/internal/user/domain"
	"go.uber.org/zap"

	dapr "github.com/dapr/go-sdk/client"
)

// 入参 po
// 响应 po

type userRepo struct {
	client    dapr.Client
	logger    *zap.Logger
	sqlClient *dao.Query
}

func NewUserRepo(
	client dapr.Client,
	logger *zap.Logger,
	sqlClient *dao.Query) domain.UserRepository {
	return &userRepo{client, logger, sqlClient}
}
