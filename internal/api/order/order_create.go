package api

import (
	"time"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/dao/model"
	"car.rental/pkg/response"
	"car.rental/struct/order"
	"github.com/gin-gonic/gin"
)

// CreateOrder 创建订单
func CreateOrder(c *gin.Context) {
	var form order.OrderCreate
	if err := c.ShouldBindJSON(&form); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	// 获取当前用户ID
	userID, exists := c.Get("userId")
	if !exists {
		response.Unauthorized(c, "用户未登录")
		return
	}

	// 解析时间
	startTime, err := time.Parse("2006-01-02 15:04:05", form.StartTime)
	if err != nil {
		response.BadRequest(c, "开始时间格式错误")
		return
	}

	endTime, err := time.Parse("2006-01-02 15:04:05", form.EndTime)
	if err != nil {
		response.BadRequest(c, "结束时间格式错误")
		return
	}

	// 检查时间是否有效
	if endTime.Before(startTime) {
		response.BadRequest(c, "结束时间不能早于开始时间")
		return
	}

	// 检查车辆是否存在
	car, err := dao.GetCarByID(form.CarID)
	if err != nil {
		response.BadRequest(c, "车辆不存在")
		return
	}

	// 检查车辆是否可用
	if car.Status != 1 {
		response.BadRequest(c, "车辆不可用")
		return
	}

	// 创建订单
	newOrder := &model.RentalOrder{
		CarID:     form.CarID,
		UserID:    userID.(uint),
		StoreID:   form.StoreID,
		StartTime: startTime,
		EndTime:   endTime,
		OrderType: form.OrderType,
		Status:    0, // 初始状态为待取车
	}

	if err := dao.CreateOrder(newOrder); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// 更新车辆状态为已出租
	if err := dao.UpdateCarStatus(form.CarID, 2); err != nil {
		response.InternalError(c, "更新车辆状态失败: "+err.Error())
		return
	}

	response.Success(c, nil)
}
