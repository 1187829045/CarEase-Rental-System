package reservation

import "time"

type ReservationInfo struct {
	ID        uint      `json:"id"`
	CarID     uint      `json:"carId"`
	UserID    uint      `json:"userId"`
	StoreID   uint      `json:"storeId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Status    int8      `json:"status"`
}

type OrderInfo struct {
	ID               uint      `json:"id"`
	ReservationID    uint      `json:"reservationId"`
	CarID            uint      `json:"carId"`
	UserID           uint      `json:"userId"`
	StoreID          uint      `json:"storeId"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
	OrderType        int8      `json:"orderType"`
	Status           int8      `json:"status"`
	PickupInspection string    `json:"pickupInspection"`
	ReturnInspection string    `json:"returnInspection"`
	PickupMileage    int64     `json:"pickupMileage"`
	ReturnMileage    int64     `json:"returnMileage"`
	PickupFuel       float64   `json:"pickupFuel"`
	ReturnFuel       float64   `json:"returnFuel"`
	PickupPhotos     string    `json:"pickupPhotos"`
	ReturnPhotos     string    `json:"returnPhotos"`
	DamageReport     string    `json:"damageReport"`
}

type OrderCreateForm struct {
	ReservationID *uint     `json:"reservationId"`
	CarID         uint      `json:"carId"`
	UserID        uint      `json:"userId"`
	StoreID       uint      `json:"storeId"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
}

type OrderQuery struct {
	Status *int8 `form:"status"`
}

type OrderPickupForm struct {
	Inspection string  `json:"inspection" binding:"required"`
	Mileage    int64   `json:"mileage" binding:"required"`
	Fuel       float64 `json:"fuel" binding:"required"`
	Photos     string  `json:"photos"`
}

type OrderReturnForm struct {
	Inspection string  `json:"inspection" binding:"required"`
	Mileage    int64   `json:"mileage" binding:"required"`
	Fuel       float64 `json:"fuel" binding:"required"`
	Photos     string  `json:"photos"`
}

type OrderExtendForm struct {
	EndTime time.Time `json:"endTime" binding:"required"`
}

type OrderDamageReportForm struct {
	Report string `json:"report" binding:"required"`
}

type ReservationCreateForm struct {
	CarID     uint      `json:"carId" binding:"required"`
	UserID    uint      `json:"userId" binding:"required"`
	StoreID   uint      `json:"storeId" binding:"required"`
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
}

type ReservationUpdateForm struct {
	CarID     *uint     `json:"carId"`
	UserID    *uint     `json:"userId"`
	StoreID   *uint     `json:"storeId"`
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Status    *int8     `json:"status"`
}

type ReservationQuery struct {
	CarID   *uint `form:"carId"`
	UserID  *uint `form:"userId"`
	StoreID *uint `form:"storeId"`
	Status  *int8 `form:"status"`
}