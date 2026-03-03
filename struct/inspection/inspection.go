package inspection

// InspectionListQuery 检测报告列表查询参数
type InspectionListQuery struct {
	CarID       *uint `form:"car_id" query:"car_id"`             // 车辆ID
	InspectorID *uint `form:"inspector_id" query:"inspector_id"` // 检测人ID
	Type        *int8 `form:"type" query:"type"`                 // 检测类型 1:上架检测, 2:还车检测
	Status      *int8 `form:"status" query:"status"`             // 状态 0:待审核, 1:已通过, 2:已拒绝
}

// InspectionUpdate 检测报告更新参数
type InspectionUpdate struct {
	ReportID      uint    `json:"report_id" binding:"required"` // 报告ID
	Mileage       *int64  `json:"mileage"`                      // 里程数
	Exterior      *string `json:"exterior"`                     // 外观检测
	Interior      *string `json:"interior"`                     // 内饰检测
	Notes         *string `json:"notes"`                        // 备注
	Photos        *string `json:"photos"`                       // 照片URL
	InspectorName *string `json:"inspector_name"`               // 检测人姓名
	Status        *int8   `json:"status"`                       // 状态 0:待审核, 1:已通过, 2:已拒绝
}

// InspectionCreate 检测报告创建参数
type InspectionCreate struct {
	CarID         uint    `json:"car_id" binding:"required"`         // 车辆ID
	Type          int8    `json:"type" binding:"required,oneof=1 2"` // 检测类型 1:上架检测, 2:还车检测
	Mileage       int64   `json:"mileage"`                           // 里程数
	FuelLevel     float64 `json:"fuel_level"`                        // 油量
	Exterior      string  `json:"exterior"`                          // 外观检测
	Interior      string  `json:"interior"`                          // 内饰检测
	Mechanical    string  `json:"mechanical"`                        // 机械检测
	Electrical    string  `json:"electrical"`                        // 电气检测
	Notes         string  `json:"notes"`                             // 备注
	Photos        string  `json:"photos"`                            // 照片URL
	InspectorName string  `json:"inspector_name"`                    // 检测人姓名
}

// InspectionListResp 检测报告列表响应
type InspectionListResp struct {
	Items  []*InspectionReportResponse `json:"items"`  // 检测报告列表
	Counts map[int8]int64              `json:"counts"` // 各状态数量
}

// InspectionReportResponse 检测报告响应
type InspectionReportResponse struct {
	ReportID       uint   `json:"report_id"`       // 报告ID
	CarID          uint   `json:"car_id"`          // 车辆ID
	UserID         uint   `json:"user_id"`         // 用户ID
	Type           int8   `json:"type"`            // 检测类型
	Mileage        int64  `json:"mileage"`         // 里程数
	Exterior       string `json:"exterior"`        // 外观检测
	Interior       string `json:"interior"`        // 内饰检测
	Notes          string `json:"notes"`           // 备注
	Photos         string `json:"photos"`          // 照片URL
	InspectorName  string `json:"inspector_name"`  // 检测人姓名
	InspectionTime string `json:"inspection_time"` // 检测时间
	Status         int8   `json:"status"`          // 状态
	UserName       string `json:"user_name"`       // 用户名
	Mobile         string `json:"mobile"`          // 手机号
}

