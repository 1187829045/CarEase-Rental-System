package api

import (
	"encoding/json"
	"net/http"

	"car.rental/dao"
	"car.rental/pkg/response"
	_struct "car.rental/struct/user"
	"github.com/gin-gonic/gin"
	"strconv"
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
		authorityIds := make([]string, 0)
		if user.Role != "" {
			if err := json.Unmarshal([]byte(user.Role), &authorityIds); err != nil {
				response.InternalError(c, err.Error())
				return
			}
		}
		items = append(items, _struct.UserInfo{
			ID:       user.UserId,
			Mobile:   user.Mobile,
			UserName: user.UserName,
			Birthday: user.Birthday,
			Gender:   user.Gender,
			Role:     authorityIds,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": items,
	})
}
