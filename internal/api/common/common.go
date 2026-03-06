package common

import (
	"strings"

	"car.rental/pkg/response"
	"github.com/gin-gonic/gin"
)

func CheckOrderPermission(orderUserID uint, c *gin.Context) bool {
	// 检查用户是否登录
	userID, exists := c.Get("userId")
	if !exists {
		response.Unauthorized(c, "用户未登录")
		return false
	}

	// 检查是否是管理员
	isAdmin := false
	authorityIds, authExists := c.Get("authorityIds")
	if authExists {
		if role, ok := authorityIds.(string); ok {
			if strings.Contains(role, "1") {
				isAdmin = true
			}
		}
	}

	// 如果不是管理员，检查是否是订单的所有者
	if !isAdmin && orderUserID != userID.(uint) {
		response.Forbidden(c, "无权操作此订单")
		return false
	}

	return true
}

// 如果是管理员，返回nil（表示查询所有订单）
// 如果不是管理员，返回当前用户ID（表示只查询自己的订单）
func GetOrderQueryUserID(c *gin.Context) (*uint, bool) {
	// 检查用户是否登录
	userID, exists := c.Get("userId")
	if !exists {
		response.Unauthorized(c, "用户未登录")
		return nil, false
	}

	// 检查是否是管理员
	isAdmin := false
	authorityIds, authExists := c.Get("authorityIds")
	if authExists {
		if role, ok := authorityIds.(string); ok {
			if strings.Contains(role, "1") {
				isAdmin = true
			}
		}
	}

	// 如果是管理员，返回nil
	if isAdmin {
		return nil, true
	}

	// 如果不是管理员，返回当前用户ID
	uid := userID.(uint)
	return &uid, true
}
