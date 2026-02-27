package api

import (
	"car.rental/consts"
	"car.rental/dao"
	"car.rental/dao/model"
	"car.rental/pkg/response"
	"car.rental/struct/car"
	"github.com/gin-gonic/gin"
)

// CreateCar 创建车辆
func CreateCar(c *gin.Context) {
	var form car.CarCreateForm
	if err := c.ShouldBindJSON(&form); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	// 获取当前用户ID
	userID, exists := c.Get("userId")
	if !exists {
		response.BadRequest(c, "用户未登录")
		return
	}

	// 创建车辆模型
	newCar := &model.CarGoods{
		UserID:           userID.(uint),
		Brand:            form.Brand,
		Model:            form.Model,
		Color:            form.Color,
		LicensePlate:     form.LicensePlate,
		Displacement:     form.Displacement,
		DriveType:        form.DriveType,
		Status:           0, //初始化默认为待检测-检测通过才会上架
		DailyRent:        form.DailyRent,
		Mileage:          form.Mileage,
		Description:      form.Description,
		Image:            form.Image,
		RegistrationDate: form.RegistrationDate,
	}

	if err := dao.CreateCar(newCar); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, newCar)
}
