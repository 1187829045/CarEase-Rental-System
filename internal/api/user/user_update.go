package api

import (
	"net/http"
	"strings"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
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
	// 获取当前用户信息
	currentUserID, exists := c.Get("userId")
	if !exists {
		response.Unauthorized(c, "用户未登录")
		return
	}

	// 检查权限：只有管理员或用户本人可以访问
	authorityIds, _ := c.Get("authorityIds")
	isAdmin := false
	if strings.Contains(authorityIds.(string), "1") {
		isAdmin = true
	}

	// 如果不是管理员且不是用户本人，返回无权限
	if !isAdmin && currentUserID.(uint) != uint(form.ID) {
		response.Forbidden(c, "无权限访问该用户信息")
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
		user.Role = form.Role
	}
	if err := dao.UpdateUser(user); err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, nil)
}
