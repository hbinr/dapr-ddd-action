package dto

type UserDTO struct {
	Id       int64  `json:"id,omitempty"`
	UserName string `json:"userName,omitempty"`
}
