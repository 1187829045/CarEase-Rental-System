package api

import (
	"strings"
	"sync"

	"car.rental/consts"
	"car.rental/dao"
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

	// 并发查询各种状态的数量
	wg := sync.WaitGroup{}
	wg.Add(5)

	// 待取车数量
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		count, err := dao.CountOrdersByStatus(0)
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
		count, err := dao.CountOrdersByStatus(1)
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
		count, err := dao.CountOrdersByStatus(2)
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
		count, err := dao.CountOrdersByStatus(3)
		if err == nil {
			resp.Counts[3] = count
		}
	}()

	// 检查用户权限
	var queryUserID *uint
	authorityIds, authExists := c.Get("authorityIds")
	isAdmin := false
	if authExists {
		if role, ok := authorityIds.(string); ok {
			if strings.Contains(role, "1") {
				isAdmin = true
			}
		}
	}

	// 如果不是管理员，只查询当前用户的订单
	if !isAdmin {
		if userID, exists := c.Get("userId"); exists {
			if uid, ok := userID.(uint); ok {
				queryUserID = &uid
			}
		}
	}

	// 查询订单列表
	orders, err := dao.ListOrders(q.Status, queryUserID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// 等待计数完成
	wg.Wait()

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
