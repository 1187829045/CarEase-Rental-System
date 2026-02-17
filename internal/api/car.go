package api

import (
	"net/http"
	"strconv"

	"car.rental/consts"
	"car.rental/dao"
	_struct "car.rental/struct"
	"github.com/gin-gonic/gin"
)

func ListCars(c *gin.Context) {
	var q _struct.CarListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	cars, err := dao.ListCars(q.Brand, q.Model, q.Status, q.MinPrice, q.MaxPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": cars,
	})
}

func GetCar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	car, err := dao.GetCarByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": car,
	})
}

func UpdateCar(c *gin.Context) {
	var form _struct.CarUpdateForm
	if err := c.ShouldBindJSON(&form); err != nil || form.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  consts.ErrInvalidParameter,
		})
		return
	}
	car, err := dao.GetCarByID(form.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if form.Brand != nil && *form.Brand != "" {
		car.Brand = *form.Brand
	}
	if form.Model != nil && *form.Model != "" {
		car.Model = *form.Model
	}
	if form.Color != nil && *form.Color != "" {
		car.Color = *form.Color
	}
	if form.LicensePlate != nil && *form.LicensePlate != "" {
		car.LicensePlate = *form.LicensePlate
	}
	if form.Seats != nil {
		car.Seats = *form.Seats
	}
	if form.FuelType != nil && *form.FuelType != "" {
		car.FuelType = *form.FuelType
	}
	if form.Displacement != nil {
		car.Displacement = *form.Displacement
	}
	if form.DriveType != nil && *form.DriveType != "" {
		car.DriveType = *form.DriveType
	}
	if form.Status != nil {
		car.Status = *form.Status
	}
	if form.DailyRent != nil {
		car.DailyRent = *form.DailyRent
	}
	if form.Mileage != nil {
		car.Mileage = *form.Mileage
	}
	if form.Description != nil && *form.Description != "" {
		car.Description = *form.Description
	}
	if err := dao.UpdateCar(car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": car,
	})
}
