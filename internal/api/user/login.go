package api

import (
	"time"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/dao/model"
	"car.rental/middlewares"
	_struct "car.rental/struct"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	loginForm := _struct.PassWordLoginForm{}
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	userInfo, err := dao.GetUserByMobile(loginForm.Mobile)
	if err != nil {
		if err.Error() == consts.UserNotFound {
			newUser := &model.User{
				Mobile:   loginForm.Mobile,
				UserName: "",
			}
			if err = dao.CreateUser(newUser); err != nil {
				c.JSON(500, gin.H{
					"code": 500,
					"msg":  err,
				})
				return
			}
			userInfo = newUser
		} else {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  err,
			})
			return
		}
	}
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID:       userInfo.UserId,
		NickName: userInfo.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*30,
			Issuer:    consts.JWTIssuer,
		},
	}
	token, err := j.CreateToken(claims)

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
