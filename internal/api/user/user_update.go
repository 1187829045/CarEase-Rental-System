package api

import (
	"encoding/json"
	"net/http"

	"car.rental/consts"
	"car.rental/dao"
	_struct "car.rental/struct/user"
	"github.com/gin-gonic/gin"
)

func UpdateUserInfo(c *gin.Context) {
	var form _struct.UserUpdateForm
	if err := c.ShouldBindJSON(&form); err != nil || form.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	user, err := dao.GetUserByID(form.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if form.Mobile != nil && *form.Mobile != "" {
		user.Mobile = *form.Mobile
	}
	if form.UserName != nil && *form.UserName != "" {
		user.UserName = *form.UserName
	}
	if form.Birthday != nil {
		user.Birthday = form.Birthday
	}
	if form.Gender != nil {
		user.Gender = *form.Gender
	}
	if len(form.Role) > 0 {
		// 将Role转换为JSON字符串
		roleJSON, err := json.Marshal(form.Role)
		if err == nil {
			user.Role = string(roleJSON)
		} else {
			// 如果JSON转换失败，直接存储为字符串
			user.Role = ""
		}
	}
	if err := dao.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": nil,
	})
}
