package api

import (
	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	"car.rental/struct/inspection"
	"github.com/gin-gonic/gin"
)

// GetInspectionList 获取检测报告列表
func GetInspectionList(c *gin.Context) {
	var q inspection.InspectionListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	reports, err := dao.ListInspectionReports(q.OrderID, q.CarID, q.UserID, q.Type, q.Status)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, reports)
}