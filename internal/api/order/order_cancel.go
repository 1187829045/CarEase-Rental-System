package api

import (
	"car.rental/consts"
	"car.rental/dao"
	"car.rental/internal/api/common"
	"car.rental/pkg/response"
	"car.rental/struct/order"
	"github.com/gin-gonic/gin"
)

// CancelOrder 取消订单
func CancelOrder(c *gin.Context) {
	var form order.OrderCancel
	if err := c.ShouldBindJSON(&form); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
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
	if orderModel.Status == 2 || orderModel.Status == 3 {
		response.BadRequest(c, "订单已取消或已完成，无法再次取消")
		return
	}
	// 更新订单状态为已取消
	orderModel.Status = 2
	// 保存取消原因
	orderModel.CancelReason = form.Reason
	if err := dao.UpdateOrder(orderModel); err != nil {
		response.InternalError(c, err.Error())
		return
	}
	if orderModel.Status == 1 {
		if err := dao.UpdateCarStatus(orderModel.CarID, 1); err != nil {
			response.InternalError(c, "更新车辆状态失败: "+err.Error())
			return
		}
	}
	response.Success(c, nil)
}
