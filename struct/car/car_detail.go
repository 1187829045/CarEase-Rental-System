package car

// CarDetailResponse 车辆详情响应结构体（扁平化）
type CarDetailResponse struct {
	// 车辆基本信息
	CarID            uint    `json:"car_id"`            // 车辆ID
	Brand            string  `json:"brand"`             // 品牌
	Model            string  `json:"model"`             // 型号
	Color            string  `json:"color"`             // 颜色
	LicensePlate     string  `json:"license_plate"`     // 车牌号
	Displacement     float64 `json:"displacement"`     // 排量
	DriveType        string  `json:"drive_type"`        // 驱动方式
	Status           int8    `json:"status"`            // 车辆状态（0:可用, 1:不可用, 2:已出租）
	DailyRent        float64 `json:"daily_rent"`        // 日租金
	Mileage          int64   `json:"mileage"`           // 里程数
	Description      string  `json:"description"`      // 车辆描述
	Image            string  `json:"image"`             // 车辆图片
	RegistrationDate string  `json:"registration_date"` // 上牌日期

	// 最新检测报告信息
	ReportID         uint    `json:"report_id,omitempty"`         // 检测报告ID
	InspectorName    string  `json:"inspector_name,omitempty"`    // 检测人姓名
	InspectionType   int8    `json:"inspection_type,omitempty"`   // 检测类型（1:取车检测, 2:还车检测）
	InspectionMileage int64  `json:"inspection_mileage,omitempty"` // 检测时里程数
	Exterior         string  `json:"exterior,omitempty"`         // 外观检测
	Interior         string  `json:"interior,omitempty"`         // 内饰检测
	Notes            string  `json:"notes,omitempty"`            // 备注
	Photos           string  `json:"photos,omitempty"`           // 照片URL
	InspectionTime   string  `json:"inspection_time,omitempty"`   // 检测时间
	InspectionStatus int8    `json:"inspection_status,omitempty"` // 检测状态（0:待审核, 1:已通过, 2:已拒绝）

	// 功能按钮
	ShowRentButton   bool    `json:"show_rent_button"`   // 是否展示租借按钮
}