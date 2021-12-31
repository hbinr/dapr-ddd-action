package command

type EditUserInfo struct {
	Id       int64
	UserID   string
	UserName string
	Age      uint
	//	如需要其他领域对象,则已嵌套方式定义在此处
	//	Vip do.Vip
}
