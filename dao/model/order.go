package model

import "time"

// 订单表
// RentalOrder 订单模型
type RentalOrder struct {
	BaseModel
	OrderID       uint      `gorm:"column:order_id;primaryKey;autoIncrement"`                  // 订单ID
	CarID         uint      `gorm:"column:car_id;not null;index"`                              // 车辆ID
	UserID        uint      `gorm:"column:user_id;not null;index"`                             // 用户ID
	StoreID       uint      `gorm:"column:store_id;not null;index"`                            // 门店ID
	StartTime     time.Time `gorm:"column:start_time;type:datetime;not null"`                  // 开始时间
	EndTime       time.Time `gorm:"column:end_time;type:datetime;not null"`                    // 结束时间
	OrderType     int8      `gorm:"column:order_type;default:0;comment:'0-预订单 1-租车订单'"`        // 订单类型
	Status        int8      `gorm:"column:status;default:0;comment:'0-待取车 1-已取车 2-已取消 3-已完成'"` // 状态
	PickupMileage int64     `gorm:"column:pickup_mileage;default:0"`                           // 取车里程
	ReturnMileage int64     `gorm:"column:return_mileage;default:0"`                           // 还车里程
	PickupFuel    float64   `gorm:"column:pickup_fuel;type:decimal(5,2)"`                      // 取车油量
	ReturnFuel    float64   `gorm:"column:return_fuel;type:decimal(5,2)"`                      // 还车油量
	PickupPhotos  string    `gorm:"column:pickup_photos;type:text"`                            // 取车照片
	ReturnPhotos  string    `gorm:"column:return_photos;type:text"`                            // 还车照片
	CancelReason  string    `gorm:"column:cancel_reason;type:text"`                             // 取消原因
}

// RentalOrder 指定表名
func (RentalOrder) TableName() string {
	return "rental_orders"
}
