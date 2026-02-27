package api

import (
	"encoding/json"
	"time"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/dao/model"
	"car.rental/middlewares"
	"car.rental/pkg/response"
	_struct "car.rental/struct"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginForm := _struct.PassWordLoginForm{}
	if err := c.ShouldBind(&loginForm); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
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
				response.InternalError(c, err.Error())
				return
			}
			userInfo = newUser
		} else {
			response.InternalError(c, err.Error())
			return
		}
	}
	j := middlewares.NewJWT()
	authorityIds := make([]string, 0)
	if userInfo.Role != "" {
		if err := json.Unmarshal([]byte(userInfo.Role), &authorityIds); err != nil {
			response.InternalError(c, err.Error())
			return
		}
	}

	claims := middlewares.CustomClaims{
		ID:           userInfo.UserId,
		NickName:     userInfo.UserName,
		AuthorityIds: authorityIds,
		Mobile:       loginForm.Mobile,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*30,
			Issuer:    consts.JWTIssuer,
		},
	}
	token, err := j.CreateToken(claims)

	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"token": token,
		"msg":   "登陆成功",
	})
}
