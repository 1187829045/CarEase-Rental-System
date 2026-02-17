package api

import (
	"net/http"
	"strconv"

	"car.rental/consts"
	"car.rental/dao"
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
