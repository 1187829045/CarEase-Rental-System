package api

import (
	"car.rental/internal/user/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	user := service.NewUser()
	token, err := user.Login(c)

	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":   "登陆成功",
		"token": token,
	})
}
