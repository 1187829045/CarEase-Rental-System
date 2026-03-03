package model

import (
	"time"
)

// InspectionReport 检测报告表
type InspectionReport struct {
	BaseModel
	ReportID       uint      `gorm:"column:report_id;primaryKey;autoIncrement"`             // 报告ID
	CarID          uint      `gorm:"column:car_id;index"`                                   // 车辆ID
	InspectorID    uint      `gorm:"column:inspector_id;index"`                             // 检测人ID
	Type           int8      `gorm:"column:type;comment '1:上架检测, 2:还车检测'"`                  // 检测类型
	Mileage        int64     `gorm:"column:mileage"`                                        // 里程数
	Exterior       string    `gorm:"column:exterior;type:text"`                             // 外观检测
	Interior       string    `gorm:"column:interior;type:text"`                             // 内饰检测
	Notes          string    `gorm:"column:notes;type:text"`                                // 备注
	Photos         string    `gorm:"column:photos;type:text"`                               // 照片URL
	InspectorName  string    `gorm:"column:inspector_name;type:varchar(50)"`                // 检测人姓名
	InspectionTime time.Time `gorm:"column:inspection_time;type:datetime"`                  // 检测时间
	Status         int8      `gorm:"column:status;default:0;comment '0:待审核, 1:已通过, 2:已拒绝'"` // 状态
}

// TableName 指定表名
func (InspectionReport) TableName() string {
	return "inspection_report"
}
