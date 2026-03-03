package car

import (
	"strings"

	"car.rental/dao/model"
)

type CarListQuery struct {
	Status   *int8 `form:"status"`
	OnlyMine *bool `form:"only_mine"`
}

// CarResponse 车辆响应结构体
type CarResponse struct {
	CarID                uint    `json:"car_id"`                 // 车辆ID
	Brand                string  `json:"brand"`                  // 品牌
	Model                string  `json:"model"`                  // 型号
	Color                string  `json:"color"`                  // 颜色
	LicensePlate         string  `json:"license_plate"`          // 车牌号
	Status               int8    `json:"status"`                 //0-可用 1- 维修 2-已出租
	DailyRent            float64 `json:"daily_rent"`             // 日租金
	Mileage              int64   `json:"mileage"`                // 里程数
	Image                string  `json:"image"`                  // 车辆图片
	RegistrationDate     string  `json:"registration_date"`      // 上牌日期
	ShowInspectionButton bool    `json:"show_inspection_button"` // 是否展示检测按钮
}

type CarListResp struct {
	Items  []*CarResponse  `json:"items"`
	Counts map[int32]int64 `json:"counts"`
}

// ConvertToCarResponse 将模型转换为响应结构体
func ConvertToCarResponse(cars []*model.CarGoods, role string) []*CarResponse {
	responses := make([]*CarResponse, len(cars))
	isInspector := false
	if strings.Contains(role, "2") {
		isInspector = true
	}
	for i, car := range cars {
		// 当车辆状态为0（待检测）时，展示检测按钮
		showInspectionButton := car.Status == 0 && isInspector
		responses[i] = &CarResponse{
			CarID:                car.CarID,
			Brand:                car.Brand,
			Model:                car.Model,
			Color:                car.Color,
			LicensePlate:         car.LicensePlate,
			Status:               car.Status,
			DailyRent:            car.DailyRent,
			Mileage:              car.Mileage,
			Image:                car.Image,
			RegistrationDate:     car.RegistrationDate.Format("2006-01-02"),
			ShowInspectionButton: showInspectionButton,
		}
	}
	return responses
}
