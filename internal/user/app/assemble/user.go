package assemble

import (
	"github.com/dapr-ddd-action/internal/user/domain/aggregate"
	"github.com/dapr-ddd-action/pkg/util/timex"
)

// do -> dto
// note: 但是不要 dto 转 do
type UserDTO struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	Role       uint   `json:"role"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func ToUserDTO(user *aggregate.User) UserDTO {
	return UserDTO{
		ID:         user.ID,
		Username:   user.Username,
		Password:   user.Password,
		Email:      user.Email,
		Phone:      user.Phone,
		Question:   user.Question,
		Answer:     user.Answer,
		Role:       user.Role,
		CreateTime: timex.DateToString(user.CreateTime),
		UpdateTime: timex.DateToString(user.UpdateTime),
	}
}
