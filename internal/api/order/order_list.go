package api

import (
	"sync"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/internal/api/common"
	"car.rental/pkg/response"
	"car.rental/struct/order"
	"github.com/gin-gonic/gin"
)

// GetOrderList 订单列表
func GetOrderList(c *gin.Context) {
	var q order.OrderListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	// 构建响应
	resp := order.OrderListResp{}
	resp.Counts = make(map[int8]int64)

	// 获取订单查询的用户ID
	queryUserID, ok := common.GetOrderQueryUserID(c)
	if !ok {
		return
	}

	// 并发查询各种状态的数量
	wg := sync.WaitGroup{}
	wg.Add(5)
	// 待取车数量
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		count, err := dao.CountOrdersByStatus(0, queryUserID)
		if err == nil {
			resp.Counts[0] = count
		}
	}()

	// 进行中数量
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		count, err := dao.CountOrdersByStatus(1, queryUserID)
		if err == nil {
			resp.Counts[1] = count
		}
	}()

	// 已还车数量
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		count, err := dao.CountOrdersByStatus(2, queryUserID)
		if err == nil {
			resp.Counts[2] = count
		}
	}()

	// 逾期数量
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		resp.Counts[3] = 0
	}()

	// 已取消数量
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		count, err := dao.CountOrdersByStatus(3, queryUserID)
		if err == nil {
			resp.Counts[3] = count
		}
	}()
	// 等待计数完成
	wg.Wait()
	// 查询订单列表
	orders, err := dao.ListOrders(q.Status, queryUserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	// 转换为响应格式
	resp.Items = make([]*order.OrderResponse, len(orders))
	for i, o := range orders {
		resp.Items[i] = &order.OrderResponse{
			OrderID:       o.OrderID,
			CarID:         o.CarID,
			UserID:        o.UserID,
			StoreID:       o.StoreID,
			StartTime:     o.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:       o.EndTime.Format("2006-01-02 15:04:05"),
			OrderType:     o.OrderType,
			Status:        o.Status,
			PickupMileage: o.PickupMileage,
			ReturnMileage: o.ReturnMileage,
			PickupFuel:    o.PickupFuel,
			ReturnFuel:    o.ReturnFuel,
			PickupPhotos:  o.PickupPhotos,
			ReturnPhotos:  o.ReturnPhotos,
		}
	}

	response.Success(c, resp)
}
