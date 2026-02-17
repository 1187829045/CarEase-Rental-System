package api

import (
	"net/http"
	"sync"

	"car.rental/consts"
	"car.rental/dao"
	_struct "car.rental/struct/car"
	"github.com/gin-gonic/gin"
)

func GetListCars(c *gin.Context) {
	var q _struct.CarListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	resp := _struct.CarListResp{}
	resp.Counts = map[int32]int64{}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer func() {
			recover()
			wg.Done()
		}()
		avail, err := dao.CountCarsByStatus(0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
		resp.Counts[2] = rented
	}()
	wg.Wait()
	cars, err := dao.ListCars(q.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	resp.Items = cars
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": resp,
	})
}
