package _struct

type RegisterForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
