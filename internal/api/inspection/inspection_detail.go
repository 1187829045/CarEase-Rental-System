package api

import (
	"strconv"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetInspectionDetail 检测单详情
func GetInspectionDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	report, err := dao.GetInspectionReportByID(uint(id))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, report)
}