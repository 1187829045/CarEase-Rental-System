package api

import (
	"time"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/dao/model"
	"car.rental/pkg/response"
	"car.rental/struct/inspection"
	"github.com/gin-gonic/gin"
)

// CreateInspection 发起检测
func CreateInspection(c *gin.Context) {
	var createData inspection.InspectionCreate
	if err := c.ShouldBindJSON(&createData); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	// 获取当前用户ID
	userID, exists := c.Get("userId")
	if !exists {
		response.Unauthorized(c, "用户未登录")
		return
	}

	// 创建检测报告
	report := &model.InspectionReport{
		CarID:          createData.CarID,
		InspectorID:    userID.(uint),
		Type:           createData.Type,
		Mileage:        createData.Mileage,
		Exterior:       createData.Exterior,
		Interior:       createData.Interior,
		Notes:          createData.Notes,
		Photos:         createData.Photos,
		InspectorName:  createData.InspectorName,
		InspectionTime: time.Now(),
		Status:         0,
	}

	if err := dao.CreateInspectionReport(report); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, report)
}
