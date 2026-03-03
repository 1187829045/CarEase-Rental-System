package api

import (
	"strconv"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/dao/model"
	"car.rental/pkg/response"
	_struct "car.rental/struct/car"
	"github.com/gin-gonic/gin"
)

func GetCarDetail(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}
	car, err := dao.GetCarByID(uint(id))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	// 获取该车辆的检测报告
	carID := uint(id)
	report, err := dao.GetLatestInspectionReportByCarID(carID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	// 构建响应数据
	responseData := buildCarDetailResponse(car, report)

	response.Success(c, responseData)
}

// buildCarDetailResponse 构建车辆详情响应数据
func buildCarDetailResponse(car *model.CarGoods, report *model.InspectionReport) _struct.CarDetailResponse {
	responseData := _struct.CarDetailResponse{
		CarID:            car.CarID,
		Brand:            car.Brand,
		Model:            car.Model,
		Color:            car.Color,
		LicensePlate:     car.LicensePlate,
		Displacement:     car.Displacement,
		DriveType:        car.DriveType,
		Status:           car.Status,
		DailyRent:        car.DailyRent,
		Mileage:          car.Mileage,
		Description:      car.Description,
		Image:            car.Image,
		RegistrationDate: car.RegistrationDate.Format("2006-01-02"),
		ShowRentButton:   false, // 默认不展示租借按钮
	}

	responseData.ReportID = report.ReportID
	responseData.InspectorName = report.InspectorName
	responseData.InspectionType = report.Type
	responseData.InspectionMileage = report.Mileage
	responseData.Exterior = report.Exterior
	responseData.Interior = report.Interior
	responseData.Notes = report.Notes
	responseData.Photos = report.Photos
	responseData.InspectionTime = report.InspectionTime.Format("2006-01-02 15:04:05")
	responseData.InspectionStatus = report.Status

	// 当车辆状态为可用且检测状态为已通过时，展示租借按钮
	if car.Status == 1 && report.Status == 1 {
		responseData.ShowRentButton = true
	}

	return responseData
}
