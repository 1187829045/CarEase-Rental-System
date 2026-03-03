package api

import (
	"strconv"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	_struct "car.rental/struct/user"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	role := c.Query("role")
	if role == "" {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 查询用户列表
	users, total, err := dao.ListUsersWithPagination(page, pageSize, role)
	if err != nil {
		response.InternalError(c, err.Error())
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
	resp := _struct.UserListResponse{
		List:     items,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	response.Success(c, resp)
}
