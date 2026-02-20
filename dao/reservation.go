package dao

import (
	"car.rental/dao/model"
	"car.rental/global"
	"errors"
)

func CreateReservation(reservation *model.Reservation) (err error) {
	result := global.DB.Create(reservation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateReservation(reservation *model.Reservation) (err error) {
	result := global.DB.Save(reservation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetReservationByID(reservationID uint) (reservation *model.Reservation, err error) {
	var r model.Reservation
	result := global.DB.Where("reservation_id = ?", reservationID).First(&r)
	if result.RowsAffected == 0 {
		return nil, errors.New("预定信息不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &r, nil
}

func ListReservations(carID *uint, userID *uint, storeID *uint, status *int8) (reservations []*model.Reservation, err error) {
	db := global.DB.Model(&model.Reservation{})
	if carID != nil {
		db = db.Where("car_id = ?", *carID)
	}
	if userID != nil {
		db = db.Where("user_id = ?", *userID)
	}
	if storeID != nil {
		db = db.Where("store_id = ?", *storeID)
	}
	if status != nil {
		db = db.Where("status = ?", *status)
	}
	result := db.Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}

func UpdateReservationStatus(reservationID uint, status int8) (err error) {
	result := global.DB.Model(&model.Reservation{}).Where("reservation_id = ?", reservationID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("预定信息不存在")
	}
	return nil
}

func CreateOrder(order *model.Order) (err error) {
	result := global.DB.Create(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateOrder(order *model.Order) (err error) {
	result := global.DB.Save(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetOrderByReservationID(reservationID uint) (order *model.Order, err error) {
	var o model.Order
	result := global.DB.Where("reservation_id = ?", reservationID).First(&o)
	if result.RowsAffected == 0 {
		return nil, errors.New("订单信息不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &o, nil
}

func GetOrderByID(orderID uint) (order *model.Order, err error) {
	var o model.Order
	result := global.DB.Where("order_id = ?", orderID).First(&o)
	if result.RowsAffected == 0 {
		return nil, errors.New("订单信息不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &o, nil
}

func ListOrders(status *int8) (orders []*model.Order, err error) {
	db := global.DB.Model(&model.Order{})
	if status != nil {
		db = db.Where("status = ?", *status)
	}
	result := db.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func UpdateOrderStatus(orderID uint, status int8) (err error) {
	result := global.DB.Model(&model.Order{}).Where("order_id = ?", orderID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("订单信息不存在")
	}
	return nil
}