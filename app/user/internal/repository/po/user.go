package po

type User struct {
	Id       int64  `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
}
