package api

import (
	"net/http"

	"car.rental/dao"
	_struct "car.rental/struct/user"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {

	users, err := dao.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	items := make([]_struct.UserInfo, 0, len(users))
	for _, user := range users {
		items = append(items, _struct.UserInfo{
			ID:       user.UserId,
			Mobile:   user.Mobile,
			UserName: user.UserName,
			Birthday: user.Birthday,
			Gender:   user.Gender,
			Role:     user.Role,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": items,
	})
}
