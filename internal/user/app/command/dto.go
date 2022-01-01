package command

type EditUserInfo struct {
	ID       int64  `json:"id,omitempty"`
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Age      uint   `json:"age,omitempty"`
	//	如需要其他领域对象,则已嵌套方式定义在此处
	//	Vip do.Vip
}
