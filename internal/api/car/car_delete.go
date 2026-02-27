package api

import (
	"strconv"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	"github.com/gin-gonic/gin"
)

// DeleteCar 删除车辆
func DeleteCar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	// 获取当前用户信息
	userID, userExists := c.Get("userId")
	authorityID, authExists := c.Get("authorityId")
	if !userExists || !authExists {
		response.BadRequest(c, "用户未登录")
		return
	}

	// 获取车辆信息
	car, err := dao.GetCarByID(uint(id))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// 检查权限：只有是自己的车并且是管理员才可以删除
	isAdmin := false
	if authID, ok := authorityID.(uint); ok && authID == consts.AdminRoleID {
		isAdmin = true
	}

	isOwner := false
	if uid, ok := userID.(uint); ok && uid == car.UserID {
		isOwner = true
	}

	if !isAdmin && !isOwner {
		response.Error(c, 403, "无权限删除该车辆")
		return
	}

	if err := dao.DeleteCarByID(uint(id)); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, nil)
}