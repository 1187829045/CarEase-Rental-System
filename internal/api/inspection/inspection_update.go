package api

import (
	"time"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	"car.rental/struct/inspection"
	"github.com/gin-gonic/gin"
)

// UpdateInspection 检测更新
func UpdateInspection(c *gin.Context) {
	var updateData inspection.InspectionUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	// 获取现有报告
	report, err := dao.GetInspectionReportByID(updateData.ReportID)
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
	// 如果状态变为已通过或已拒绝，更新检测时间
	if updateData.Status != nil && (*updateData.Status == 1 || *updateData.Status == 2) {
		report.InspectionTime = time.Now()
	}

	if err := dao.UpdateInspectionReport(report); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// 如果检测通过，更新车辆状态为可用
	if updateData.Status != nil && *updateData.Status == 1 {
		if err := dao.UpdateCarStatus(report.CarID, 1); err != nil {
			response.InternalError(c, "更新车辆状态失败: "+err.Error())
			return
		}
	}
	response.Success(c, report)
}
