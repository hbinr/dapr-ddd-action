package aggregate

import (
	"strconv"
	"time"
)

// User  用户聚合根
type User struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	Role       uint      `json:"role"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (u *User) GetUserInfoKey(id int64) string {
	return "user:info" + strconv.Itoa(int(id))
}
