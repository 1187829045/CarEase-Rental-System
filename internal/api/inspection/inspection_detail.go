package api

import (
	"strconv"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	"car.rental/struct/inspection"
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

	// 构建检测报告响应
	inspectionResp := &inspection.InspectionReportResponse{
		ReportID:       report.ReportID,
		CarID:          report.CarID,
		UserID:         report.InspectorID,
		Type:           report.Type,
		Mileage:        report.Mileage,
		Exterior:       report.Exterior,
		Interior:       report.Interior,
		Notes:          report.Notes,
		Photos:         report.Photos,
		InspectorName:  report.InspectorName,
		InspectionTime: report.InspectionTime.Format("2006-01-02 15:04:05"),
		Status:         report.Status,
	}
	// 如果有检测人ID，查询检测人信息
	if report.InspectorID > 0 {
		inspector, err := dao.GetUserByID(report.InspectorID)
		if err == nil {
			// 构建检测人信息，排除敏感字段
			inspectionResp.UserName = inspector.UserName
			inspectionResp.Mobile = inspector.Mobile
		}
	}

	response.Success(c, inspectionResp)
}
