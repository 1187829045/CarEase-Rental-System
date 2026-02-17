package api

import (
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
	if form.Role != nil {
		user.Role = int8(*form.Role)
	}
	if err := dao.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	resp := _struct.UserInfo{
		ID:       user.UserId,
		Mobile:   user.Mobile,
		UserName: user.UserName,
		Birthday: user.Birthday,
		Gender:   user.Gender,
		Role:     user.Role,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": resp,
	})
}
