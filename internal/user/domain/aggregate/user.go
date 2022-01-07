package aggregate

import "strconv"

// User  用户聚合根
type User struct {
	ID       int64  `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
}

func (u *User) GetUserInfoKey(id int64) string {
	return "user:info" + strconv.Itoa(int(id))
}
