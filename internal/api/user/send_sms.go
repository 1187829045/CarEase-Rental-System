package api

import (
	"car.rental/consts"
	_struct "car.rental/struct"
	"github.com/gin-gonic/gin"
)

func SendSMS(c *gin.Context) {
	form := _struct.SendSMSForm{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	//todo 调用service发送短信
	c.JSON(200, gin.H{
		"msg": "短信发送成功",
	})
}
