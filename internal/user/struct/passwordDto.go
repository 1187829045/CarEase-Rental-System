package _struct

type PassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"` //手机号码格式有规范可寻， 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	UserName string `form:"username" json:"username" binding:"required"`
}
