package car

import (
	"car.rental/dao/model"
)

type CarListQuery struct {
	Status *int8 `form:"status"`
}

// CarResponse 车辆响应结构体
type CarResponse struct {
	CarID            uint    `json:"car_id"`            // 车辆ID
	Brand            string  `json:"brand"`             // 品牌
	Model            string  `json:"model"`             // 型号
	Color            string  `json:"color"`             // 颜色
	LicensePlate     string  `json:"license_plate"`     // 车牌号
	Status           int8    `json:"status"`            //0-可用 1- 维修 2-已出租
	DailyRent        float64 `json:"daily_rent"`        // 日租金
	Mileage          int64   `json:"mileage"`           // 里程数
	Image            string  `json:"image"`             // 车辆图片
	RegistrationDate string  `json:"registration_date"` // 上牌日期
}

type CarListResp struct {
	Items  []*CarResponse  `json:"items"`
	Counts map[int32]int64 `json:"counts"`
}

// ConvertToCarResponse 将模型转换为响应结构体
func ConvertToCarResponse(cars []*model.CarGoods) []*CarResponse {
	responses := make([]*CarResponse, len(cars))
	for i, car := range cars {
		responses[i] = &CarResponse{
			CarID:            car.CarID,
			Brand:            car.Brand,
			Model:            car.Model,
			Color:            car.Color,
			LicensePlate:     car.LicensePlate,
			Status:           car.Status,
			DailyRent:        car.DailyRent,
			Mileage:          car.Mileage,
			Image:            car.Image,
			RegistrationDate: car.RegistrationDate.Format("2006-01-02"),
		}
	}
	return responses
}
