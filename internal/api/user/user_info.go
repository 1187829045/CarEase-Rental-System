package api

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	user, err := dao.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	authorityIds := make([]string, 0)
	if user.Role != "" {
		if err := json.Unmarshal([]byte(user.Role), &authorityIds); err != nil {
			response.InternalError(c, err.Error())
			return
		}
	}
	resp := _struct.UserInfo{
		ID:       user.UserId,
		Mobile:   user.Mobile,
		UserName: user.UserName,
		Birthday: user.Birthday,
		Gender:   user.Gender,
		Role:     authorityIds,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": resp,
	})
}
