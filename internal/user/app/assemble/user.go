package assemble

import "github.com/dapr-ddd-action/internal/user/domain/aggregate"

type UserDTO struct {
	ID       int64  `json:"id,omitempty"`
	UserName string `json:"userName,omitempty"`
}

func UserToDTO(user *aggregate.User) UserDTO {
	return UserDTO{
		ID:       user.ID,
		UserName: user.UserName,
	}
}
