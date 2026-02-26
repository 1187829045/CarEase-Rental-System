package _struct

type SendSMSForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,len=11,regexp=^1[3-9]\\d{9}$"`
}
