package _struct

type PassWordLoginForm struct {
	Mobile     string `form:"mobile" json:"mobile" binding:"required"`
	VerifyCode string `form:"code" json:"code" binding:"omitempty,len=6"`
}
