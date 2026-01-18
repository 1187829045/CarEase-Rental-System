package api

import (
	"car.rental/internal/user/service"
	"errors"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	user := service.NewUser()
	err := user.Register(c)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  errors.New("注册失败"),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": errors.New("注册成功"),
	})
}

//redis加锁 key expiredtime
