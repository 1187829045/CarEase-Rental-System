package model

import (
	"time"
)

type CarGoods struct {
	BaseModel
	CarID            uint      `gorm:"column:car_id;primaryKey;autoIncrement"`                      // 车辆ID
	UserID           uint      `gorm:"column:user_id;index"`                                        // 所属用户ID
	Brand            string    `gorm:"column:brand;type:varchar(20);not null"`                      // 品牌
	Model            string    `gorm:"column:model;type:varchar(30);not null"`                      // 型号
	Color            string    `gorm:"column:color;type:varchar(10);not null"`                      // 颜色
	LicensePlate     string    `gorm:"column:license_plate;type:varchar(15);unique;not null;index"` // 车牌号
	Displacement     float64   `gorm:"column:displacement;type:decimal(3,1)"`                       // 排量
	DriveType        string    `gorm:"column:drive_type;type:varchar(10);not null"`                 // 驱动方式 手动，自动
	Status           int8      `gorm:"column:status;default:0;comment '0:待检测, 1:可用, 2:已出租'"`        // 车辆状态 //0-待检测 1-可用 2-已出租
	DailyRent        float64   `gorm:"column:daily_rent;type:decimal(10,2);not null"`               // 日租金
	Mileage          int64     `gorm:"column:mileage;default:0"`                                    // 里程数
	Description      string    `gorm:"column:description;type:text"`                                // 车辆描述
	Image            string    `gorm:"column:image;type:varchar(255)"`                              // 车辆图片
	RegistrationDate time.Time `gorm:"column:registration_date;type:date"`                          // 上牌日期
}

func (CarGoods) TableName() string {
	return "car_goods"
}
