package api

import (
	"strconv"
	"strings"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	_struct "car.rental/struct/user"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
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
	if !isAdmin && currentUserID.(uint) != uint(id) {
		response.Forbidden(c, "无权限访问该用户信息")
		return
	}

	user, err := dao.GetUserByID(uint(id))
	if err != nil {
		response.InternalError(c, err.Error())
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
	response.Success(c, resp)
}
