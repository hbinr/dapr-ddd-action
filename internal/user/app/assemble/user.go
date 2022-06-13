package assemble

import (
	"github.com/dapr-ddd-action/internal/user/domain/aggregate"
	"github.com/dapr-ddd-action/pkg/util/timex"
)

// DTO Assembler是组装器，负责完成domain model对象到dto的转换，组装职责包括：

// 1. 完成类型转换、数据格式化；如日志格式化，状态enum装换为前端认识的string；
// 2. 将多个domain领域对象组装为需要的dto对象，比如查询帖子列表，需要从Post（帖子）领域对象中获取帖子的详情，还需要从User（用户）领域对象中获取用户的社交信息（昵称、简介、头像等）；
// 3. 将domain领域对象属性裁剪并组装为dto；某些场景下，可能并不需要所有domain领域对象的属性，比如User领域对象的password属性属于隐私相关属性，在“查询用户信息”case中不需要返回，需要裁剪掉。
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
