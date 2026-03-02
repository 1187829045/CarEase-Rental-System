package api

import (
	"encoding/json"
	"net/http"

	"car.rental/dao"
	_struct "car.rental/struct/user"
	"car.rental/tools"
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
		// 处理角色
		var roleStrings []string
		if user.Role != "" {
			if err := json.Unmarshal([]byte(user.Role), &roleStrings); err != nil {
				// 如果解析失败，说明是单个角色字符串
				roleStrings = []string{user.Role}
			}
		}
		
		// 使用工具函数转换角色
		authorityIds := tools.ConvertStringRolesToInt8(roleStrings)
		
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
