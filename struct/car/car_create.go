package car

import "time"

// CarCreateForm 车辆创建请求结构体
type CarCreateForm struct {
	Brand            string    `json:"brand" binding:"required"`         // 品牌
	Model            string    `json:"model" binding:"required"`         // 型号
	Color            string    `json:"color" binding:"required"`         // 颜色
	LicensePlate     string    `json:"license_plate" binding:"required"` // 车牌号
	Displacement     float64   `json:"displacement"`                     // 排量
	DriveType        string    `json:"drive_type" binding:"required"`    // 驱动方式
	DailyRent        float64   `json:"daily_rent" binding:"required"`    // 日租金
	Mileage          int64     `json:"mileage"`                          // 里程数
	Description      string    `json:"description"`                      // 车辆描述
	Image            string    `json:"image"`                            // 车辆图片
	RegistrationDate time.Time `json:"registration_date"`                // 上牌日期
}
