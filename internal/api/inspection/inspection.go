package api

import (
	// "net/http"
	"strconv"

	"car.rental/consts"
	"car.rental/dao"
	// "car.rental/dao/model"
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

// GetInspectionDetail 获取检测报告详情
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

// UpdateInspection 更新检测报告
func UpdateInspection(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	var updateData inspection.InspectionUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	// 获取现有报告
	report, err := dao.GetInspectionReportByID(uint(id))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// 更新字段
	if updateData.Mileage != nil {
		report.Mileage = *updateData.Mileage
	}
	if updateData.Exterior != nil {
		report.Exterior = *updateData.Exterior
	}
	if updateData.Interior != nil {
		report.Interior = *updateData.Interior
	}
	if updateData.Notes != nil {
		report.Notes = *updateData.Notes
	}
	if updateData.Photos != nil {
		report.Photos = *updateData.Photos
	}
	if updateData.InspectorName != nil {
		report.InspectorName = *updateData.InspectorName
	}
	if updateData.Status != nil {
		report.Status = *updateData.Status
	}

	if err := dao.UpdateInspectionReport(report); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, report)
}