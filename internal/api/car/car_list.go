package api

import (
	"sync"

	"car.rental/consts"
	"car.rental/dao"
	"car.rental/pkg/response"
	_struct "car.rental/struct/car"
	"github.com/gin-gonic/gin"
)

func GetListCars(c *gin.Context) {
	var q _struct.CarListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}
	resp := _struct.CarListResp{}
	resp.Counts = map[int32]int64{}
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		avail, err := dao.CountCarsByStatus(0)
		if err != nil {
			response.InternalError(c, err.Error())
			return
		}
		resp.Counts[0] = avail
	}()
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		unavail, err := dao.CountCarsByStatus(1)
		if err != nil {
			response.InternalError(c, err.Error())
			return
		}
		resp.Counts[1] = unavail
	}()
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		rented, err := dao.CountCarsByStatus(2)
		if err != nil {
			response.InternalError(c, err.Error())
			return
		}
		resp.Counts[2] = rented
	}()
	// 用于存储列表查询结果和错误
	var cars []*_struct.CarResponse
	var listErr error

	// 启动协程查询车辆列表
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		
		// 处理OnlyMine参数
		var userID *uint
		if q.OnlyMine != nil && *q.OnlyMine {
			if id, exists := c.Get("userId"); exists {
				if uid, ok := id.(uint); ok {
					userID = &uid
				}
			}
		}
		
		result, err := dao.ListCars(q.Status, userID)
		if err != nil {
			listErr = err
			return
		}
		cars = _struct.ConvertToCarResponse(result)
	}()

	wg.Wait()

	// 检查列表查询是否出错
	if listErr != nil {
		response.InternalError(c, listErr.Error())
		return
	}
	resp.Items = cars
	response.Success(c, resp)
}
