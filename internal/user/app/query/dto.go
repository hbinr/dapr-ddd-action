package query

type User struct {
	ID       int64  `json:"id,omitempty"`
	UserName string `json:"userName,omitempty"`
}
