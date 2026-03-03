package order

// OrderListQuery 订单列表查询参数
type OrderListQuery struct {
	Status *int8 `form:"status" query:"status"` // 订单状态
}

// OrderCreate 创建订单请求
type OrderCreate struct {
	CarID     uint   `json:"car_id" binding:"required"`     // 车辆ID
	StartTime string `json:"start_time" binding:"required"` // 开始时间
	EndTime   string `json:"end_time" binding:"required"`   // 结束时间
	StoreID   uint   `json:"store_id" binding:"required"`   // 门店ID
	OrderType int8   `json:"order_type" binding:"required"` // 订单类型 0-预订单 1-租车订单
}

// OrderResponse 订单响应
type OrderResponse struct {
	OrderID       uint    `json:"order_id"`       // 订单ID
	CarID         uint    `json:"car_id"`         // 车辆ID
	UserID        uint    `json:"user_id"`        // 用户ID
	StoreID       uint    `json:"store_id"`       // 门店ID
	StartTime     string  `json:"start_time"`     // 开始时间
	EndTime       string  `json:"end_time"`       // 结束时间
	OrderType     int8    `json:"order_type"`     // 订单类型
	Status        int8    `json:"status"`         // 订单状态
	PickupMileage int64   `json:"pickup_mileage"` // 取车里程
	ReturnMileage int64   `json:"return_mileage"` // 还车里程
	PickupFuel    float64 `json:"pickup_fuel"`    // 取车油量
	ReturnFuel    float64 `json:"return_fuel"`    // 还车油量
	PickupPhotos  string  `json:"pickup_photos"`  // 取车照片
	ReturnPhotos  string  `json:"return_photos"`  // 还车照片
}

// OrderDetailResponse 订单详情响应
type OrderDetailResponse struct {
	Order *OrderResponse `json:"order"` // 订单信息
	Car   *CarInfo       `json:"car"`   // 车辆信息
	Store *StoreInfo     `json:"store"` // 门店信息
	User  *UserInfo      `json:"user"`  // 用户信息
}

// CarInfo 车辆信息
type CarInfo struct {
	CarID        uint    `json:"car_id"`        // 车辆ID
	Brand        string  `json:"brand"`         // 品牌
	Model        string  `json:"model"`         // 型号
	LicensePlate string  `json:"license_plate"` // 车牌号
	DailyRent    float64 `json:"daily_rent"`    // 日租金
}

// StoreInfo 门店信息
type StoreInfo struct {
	StoreID   uint   `json:"store_id"`   // 门店ID
	StoreName string `json:"store_name"` // 门店名称
	Address   string `json:"address"`    // 地址
}

// UserInfo 用户信息
type UserInfo struct {
	UserID   uint   `json:"user_id"`   // 用户ID
	UserName string `json:"user_name"` // 用户名
	Mobile   string `json:"mobile"`    // 手机号
}

// OrderListResp 订单列表响应
type OrderListResp struct {
	Items  []*OrderResponse `json:"items"`  // 订单列表
	Counts map[int8]int64   `json:"counts"` // 各状态数量
}
