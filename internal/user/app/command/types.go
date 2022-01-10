package command

type EditUserInfoCmd struct {
	ID        int64  `json:"id" vd:"$>0"`
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"  vd:"len($)>=5 && len($)<=32; msg:sprintf('用户名长度必须大于5-32个字符之间')"`
	Age       int    `json:"age"  vd:"$>0; msg:sprintf('年龄必须大于0')"`
	CreatedAt string `json:"create_time"`
	UpdatedAt string `json:"update_time"`
}