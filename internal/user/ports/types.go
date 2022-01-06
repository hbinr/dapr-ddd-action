package ports

type UpdateUserReq struct {
	ID       int64  `json:"id" vd:"$>0"`
	UserName string `json:"user_name"  vd:"len($)>5; msg:sprintf('用户名[%v]必须大于5个字符',$)"`
	Age      int    `json:"age"  vd:"$>0; msg:sprintf('年龄必须大于0')"`
}
