package assemble

import (
	"github.com/dapr-ddd-action/internal/user/domain/aggregate"
	"github.com/dapr-ddd-action/pkg/util/timex"
)

type UserDTO struct {
	ID        int64  `json:"id"`
	Age       int    `json:"age"`
	UserName  string `json:"userName"`
	CreatedAt string `json:"createTime"`
	UpdatedAt string `json:"updateTime"`
}

func ToUserDTO(user *aggregate.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		UserName:  user.UserName,
		Age:       user.Age,
		CreatedAt: timex.DateToString(user.CreatedAt),
		UpdatedAt: timex.DateToString(user.UpdatedAt),
	}
}
