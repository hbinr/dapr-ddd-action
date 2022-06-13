package converter

import (
	"github.com/dapr-ddd-action/internal/user/adapters/repository/data/po"
	"github.com/dapr-ddd-action/internal/user/domain/aggregate"
)

// DO Converter 在Infrastructure层，Entity到DO的转化器没有一个标准名称，定位 Data Converter
// do -> po
// po -> do
// note: 待 判空 + 处理业务逻辑边界，如果有错误，需要增加 error 返回值

func FromUserDO(userDO *aggregate.User) *po.User {
	return &po.User{
		ID:         int32(userDO.ID),
		Username:   userDO.Username,
		Password:   userDO.Password,
		Email:      userDO.Email,
		Phone:      userDO.Phone,
		Question:   userDO.Question,
		Answer:     userDO.Answer,
		Role:       int32(userDO.Role),
		CreateTime: userDO.CreateTime,
		UpdateTime: userDO.UpdateTime,
	}
}

func ToUserDO(userPO *po.User) *aggregate.User {
	return &aggregate.User{
		ID:         int64(userPO.ID),
		Username:   userPO.Username,
		Password:   userPO.Password,
		Email:      userPO.Email,
		Phone:      userPO.Phone,
		Question:   userPO.Question,
		Answer:     userPO.Answer,
		Role:       uint(userPO.Role),
		CreateTime: userPO.CreateTime,
		UpdateTime: userPO.UpdateTime,
	}
}
