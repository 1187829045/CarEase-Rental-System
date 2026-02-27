package inspection

// InspectionListQuery 检测报告列表查询参数
type InspectionListQuery struct {
	OrderID *uint  `form:"order_id" query:"order_id"`
	CarID   *uint  `form:"car_id" query:"car_id"`
	UserID  *uint  `form:"user_id" query:"user_id"`
	Type    *int8  `form:"type" query:"type"`
	Status  *int8  `form:"status" query:"status"`
}

// InspectionUpdate 检测报告更新参数
type InspectionUpdate struct {
	Mileage       *int64  `json:"mileage"`
	Exterior      *string `json:"exterior"`
	Interior      *string `json:"interior"`
	Notes         *string `json:"notes"`
	Photos        *string `json:"photos"`
	InspectorName *string `json:"inspector_name"`
	Status        *int8   `json:"status"`
}

// InspectionCreate 检测报告创建参数
type InspectionCreate struct {
	OrderID        uint    `json:"order_id" binding:"required"`
	CarID          uint    `json:"car_id" binding:"required"`
	UserID         uint    `json:"user_id" binding:"required"`
	Type           int8    `json:"type" binding:"required,oneof=1 2"`
	Mileage        int64   `json:"mileage"`
	FuelLevel      float64 `json:"fuel_level"`
	Exterior       string  `json:"exterior"`
	Interior       string  `json:"interior"`
	Mechanical     string  `json:"mechanical"`
	Electrical     string  `json:"electrical"`
	Notes          string  `json:"notes"`
	Photos         string  `json:"photos"`
	InspectorName  string  `json:"inspector_name"`
}