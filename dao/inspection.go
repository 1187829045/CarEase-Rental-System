package dao

import (
	"car.rental/dao/model"
	"car.rental/global"
)

// CreateInspectionReport 创建检测报告
func CreateInspectionReport(report *model.InspectionReport) error {
	result := global.DB.Create(report)
	return result.Error
}

// GetInspectionReportByID 根据ID获取检测报告
func GetInspectionReportByID(reportID uint) (*model.InspectionReport, error) {
	var report model.InspectionReport
	result := global.DB.Where("report_id = ?", reportID).Find(&report)
	if result.Error != nil {
		return nil, result.Error
	}
	return &report, nil
}

// UpdateInspectionReport 更新检测报告
func UpdateInspectionReport(report *model.InspectionReport) error {
	result := global.DB.Save(report)
	return result.Error
}

// ListInspectionReports 检测报告列表
func ListInspectionReports(orderID, carID, userID *uint, reportType *int8, status *int8) ([]*model.InspectionReport, error) {
	var reports []*model.InspectionReport
	db := global.DB.Model(&model.InspectionReport{})

	// 筛选条件
	if orderID != nil {
		db = db.Where("order_id = ?", *orderID)
	}
	if carID != nil {
		db = db.Where("car_id = ?", *carID)
	}
	if userID != nil {
		db = db.Where("user_id = ?", *userID)
	}
	if reportType != nil {
		db = db.Where("type = ?", *reportType)
	}
	if status != nil {
		db = db.Where("status = ?", *status)
	}

	// 按检测时间倒序排序
	result := db.Order("inspection_time DESC").Find(&reports)
	if result.Error != nil {
		return nil, result.Error
	}
	return reports, nil
}

// GetInspectionReportByOrderAndType 根据订单ID和检测类型获取检测报告
func GetInspectionReportByOrderAndType(orderID uint, reportType int8) (*model.InspectionReport, error) {
	var report model.InspectionReport
	result := global.DB.Where("order_id = ? AND type = ?", orderID, reportType).First(&report)
	if result.Error != nil {
		return nil, result.Error
	}
	return &report, nil
}
