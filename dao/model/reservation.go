package model

import "time"

// 预定订单表
type Reservation struct {
	BaseModel
	ReservationID uint      `gorm:"column:reservation_id;primaryKey;autoIncrement"`
	CarID         uint      `gorm:"column:car_id;not null"`
	UserID        uint      `gorm:"column:user_id;not null"`
	StoreID       uint      `gorm:"column:store_id;not null"` //门店id
	StartTime     time.Time `gorm:"column:start_time;type:datetime"`
	EndTime       time.Time `gorm:"column:end_time;type:datetime"`
	Status        int8      `gorm:"column:status;default:0"` //状态 0-待确认 1-已确认 2-已取消
}
