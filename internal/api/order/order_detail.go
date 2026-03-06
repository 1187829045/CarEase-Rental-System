package api

import (
	"strconv"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	"car.rental/struct/order"
	"github.com/gin-gonic/gin"
	"car.rental/internal/api/common"
)

// GetOrderDetail 订单详情
func GetOrderDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	// 获取订单
	orderModel, err := dao.GetOrderByID(uint(id))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// 检查用户权限
	if !common.CheckOrderPermission(orderModel.UserID, c) {
		return
	}

	// 构建订单响应
	orderResp := &order.OrderResponse{
		OrderID:       orderModel.OrderID,
		CarID:         orderModel.CarID,
		UserID:        orderModel.UserID,
		StoreID:       orderModel.StoreID,
		StartTime:     orderModel.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:       orderModel.EndTime.Format("2006-01-02 15:04:05"),
		OrderType:     orderModel.OrderType,
		Status:        orderModel.Status,
		PickupMileage: orderModel.PickupMileage,
		ReturnMileage: orderModel.ReturnMileage,
		PickupFuel:    orderModel.PickupFuel,
		ReturnFuel:    orderModel.ReturnFuel,
		PickupPhotos:  orderModel.PickupPhotos,
		ReturnPhotos:  orderModel.ReturnPhotos,
	}

	// 构建响应数据
	responseData := &order.OrderDetailResponse{
		Order: orderResp,
		Car:   nil,
		User:  nil,
	}

	// 获取车辆信息
	car, err := dao.GetCarByID(orderModel.CarID)
	if err == nil {
		responseData.Car = &order.CarInfo{
			CarID:        car.CarID,
			Brand:        car.Brand,
			Model:        car.Model,
			LicensePlate: car.LicensePlate,
			DailyRent:    car.DailyRent,
		}
	}

	// 获取用户信息
	user, err := dao.GetUserByID(orderModel.UserID)
	if err == nil {
		responseData.User = &order.UserInfo{
			UserID:   user.UserId,
			UserName: user.UserName,
			Mobile:   user.Mobile,
		}
	}
	response.Success(c, responseData)
}
