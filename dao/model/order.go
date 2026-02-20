package model

import "time"

// 订单表
type Order struct {
	BaseModel
	OrderID          uint      `gorm:"column:order_id;primaryKey;autoIncrement"`
	ReservationID    uint      `gorm:"column:reservation_id;index"` //预定单id
	CarID            uint      `gorm:"column:car_id;not null"`       //车辆id
	UserID           uint      `gorm:"column:user_id;not null"`      //用户id
	StoreID          uint      `gorm:"column:store_id;not null"`     //门店id
	StartTime        time.Time `gorm:"column:start_time;type:datetime"` //开始时间
	EndTime          time.Time `gorm:"column:end_time;type:datetime"`   //结束时间
	OrderType        int8      `gorm:"column:order_type;default:0"`     //订单类型 0-普通订单 1-扩展订单
	Status           int8      `gorm:"column:status;default:0"`         //状态 0-待确认 1-已确认 2-已取消 3-已完成
	PickupInspection string    `gorm:"column:pickup_inspection;type:text"`
	ReturnInspection string    `gorm:"column:return_inspection;type:text"`
	PickupMileage    int64     `gorm:"column:pickup_mileage;default:0"`
	ReturnMileage    int64     `gorm:"column:return_mileage;default:0"`
	PickupFuel       float64   `gorm:"column:pickup_fuel;type:decimal(5,2)"`
	ReturnFuel       float64   `gorm:"column:return_fuel;type:decimal(5,2)"`
	PickupPhotos     string    `gorm:"column:pickup_photos;type:text"`
	ReturnPhotos     string    `gorm:"column:return_photos;type:text"`
	DamageReport     string    `gorm:"column:damage_report;type:text"`
}
