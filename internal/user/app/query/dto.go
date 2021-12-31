package query

type User struct {
	Id       int64  `json:"id,omitempty"`
	UserName string `json:"userName,omitempty"`
}
