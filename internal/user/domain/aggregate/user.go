package aggregate

import (
	"strconv"
	"time"
)

// User  用户聚合根
type User struct {
	ID        int64     `json:"id,"`
	UserName  string    `json:"user_name,"`
	Age       int       `json:"age,"`
	CreatedAt time.Time `json:"create_time,"`
	UpdatedAt time.Time `json:"update_time,"`
}

func (u *User) GetUserInfoKey(id int64) string {
	return "user:info" + strconv.Itoa(int(id))
}
