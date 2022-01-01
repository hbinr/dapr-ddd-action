package ports

type UpdateUserReq struct {
	ID       int64  `query:"id,required"`
	UserName string `json:"userName,required"  vd:"len($)>5; msg:sprintf('用户名[%v]必须大于5个字符',$)"`
}
