package dao

import (
	"car.rental/dao/model"
	"car.rental/global"
)

// CreateOrder 创建订单
func CreateOrder(order *model.RentalOrder) error {
	result := global.DB.Create(order)
	return result.Error
}

// GetOrderByID 根据ID获取订单
func GetOrderByID(orderID uint) (*model.RentalOrder, error) {
	var order model.RentalOrder
	result := global.DB.Where("order_id = ?", orderID).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// ListOrders 列出订单
func ListOrders(status *int8, userID *uint) ([]*model.RentalOrder, error) {
	var orders []*model.RentalOrder
	db := global.DB.Model(&model.RentalOrder{})

	// 筛选条件
	if status != nil {
		db = db.Where("status = ?", *status)
	}
	if userID != nil {
		db = db.Where("user_id = ?", *userID)
	}

	// 按创建时间倒序排序
	result := db.Order("created_at DESC").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

// CountOrdersByStatus 统计订单状态数量
func CountOrdersByStatus(status int8) (int64, error) {
	var count int64
	result := global.DB.Model(&model.RentalOrder{}).Where("status = ?", status).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// UpdateOrderStatus 更新订单状态
func UpdateOrderStatus(orderID uint, status int8) error {
	result := global.DB.Model(&model.RentalOrder{}).Where("order_id = ?", orderID).Update("status", status)
	return result.Error
}
