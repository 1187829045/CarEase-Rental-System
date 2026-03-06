package api

import (
	"time"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/internal/api/common"
	"car.rental/pkg/response"
	"car.rental/struct/order"
	"github.com/gin-gonic/gin"
)

// OrderAction 订单操作（取车、还车、续租）
func OrderAction(c *gin.Context) {
	var form order.OrderAction
	if err := c.ShouldBindJSON(&form); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}
	if form.OrderID == 0 {
		response.BadRequest(c, "订单ID不能为空")
		return
	}
	// 获取订单
	orderModel, err := dao.GetOrderByID(form.OrderID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	// 检查用户权限
	if !common.CheckOrderPermission(orderModel.UserID, c) {
		return
	}
	// 根据操作类型处理
	switch form.Action {
	case consts.ActionPickup:
		// 取车操作
		if orderModel.Status != consts.OrderStatusPending {
			response.BadRequest(c, "订单状态不是待取车")
			return
		}
		if form.Mileage != nil {
			orderModel.PickupMileage = *form.Mileage
		}
		if form.Fuel != nil {
			orderModel.PickupFuel = *form.Fuel
		}
		if form.Photos != nil {
			orderModel.PickupPhotos = *form.Photos
		}
		orderModel.Status = consts.OrderStatusConfirmed
	case consts.ActionReturn:
		// 还车操作
		if orderModel.Status != consts.OrderStatusConfirmed {
			response.BadRequest(c, "订单状态不是进行中")
			return
		}
		if form.Mileage != nil {
			orderModel.ReturnMileage = *form.Mileage
		}
		if form.Fuel != nil {
			orderModel.ReturnFuel = *form.Fuel
		}
		if form.Photos != nil {
			orderModel.ReturnPhotos = *form.Photos
		}
		// 更新订单状态为已还车
		orderModel.Status = consts.OrderStatusCompleted

		// 更新车辆状态为可用
		if err := dao.UpdateCarStatus(orderModel.CarID, consts.CarStatusAvailable); err != nil {
			response.InternalError(c, "更新车辆状态失败: "+err.Error())
			return
		}

	case consts.ActionExtend:
		// 续租操作
		if orderModel.Status != consts.OrderStatusConfirmed {
			response.BadRequest(c, "订单状态不是进行中")
			return
		}
		// 解析续租结束时间
		if form.ExtendEndTime == nil {
			response.BadRequest(c, "续租结束时间不能为空")
			return
		}
		extendEndTime, err := time.Parse("2006-01-02 15:04:05", *form.ExtendEndTime)
		if err != nil {
			response.BadRequest(c, "续租结束时间格式错误")
			return
		}
		// 检查时间是否有效
		if extendEndTime.Before(orderModel.EndTime) {
			response.BadRequest(c, "续租结束时间不能早于原结束时间")
			return
		}
		orderModel.EndTime = extendEndTime
	default:
		response.BadRequest(c, "无效的操作类型")
		return
	}

	// 保存订单
	if err := dao.UpdateOrder(orderModel); err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, nil)
}
