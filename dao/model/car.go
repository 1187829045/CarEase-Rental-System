package model

type Car struct {
	BaseModel
	CarID        uint    `gorm:"column:car_id;primaryKey;autoIncrement"`                      // 车辆ID
	Brand        string  `gorm:"column:brand;type:varchar(20);not null"`                      // 品牌
	Model        string  `gorm:"column:model;type:varchar(30);not null"`                      // 型号
	Color        string  `gorm:"column:color;type:varchar(10);not null"`                      // 颜色
	LicensePlate string  `gorm:"column:license_plate;type:varchar(15);unique;not null;index"` // 车牌号
	Seats        int8    `gorm:"column:seats;not null"`                                       // 座位数
	FuelType     string  `gorm:"column:fuel_type;type:varchar(10);not null"`                  // 燃料类型
	Displacement float64 `gorm:"column:displacement;type:decimal(3,1)"`                       // 排量
	DriveType    string  `gorm:"column:drive_type;type:varchar(10)"`                          // 驱动方式
	Status       int8    `gorm:"column:status;default:0;comment '0:可用, 1:不可用, 2:已出租'"`        // 车辆状态 可用，维修，已租出
	DailyRent    float64 `gorm:"column:daily_rent;type:decimal(10,2);not null"`               // 日租金
	Mileage      int64   `gorm:"column:mileage;default:0"`                                    // 里程数
	Description  string  `gorm:"column:description;type:text"`                                // 车辆描述
}
