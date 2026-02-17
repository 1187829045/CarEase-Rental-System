package dao

import (
	"errors"

	"car.rental/dao/model"
	"car.rental/global"
)

// CreateCar 创建汽车信息
func CreateCar(car *model.Car) (err error) {
	result := global.DB.Create(car)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ListCars 按条件筛选车辆列表
func ListCars(status *int8) (cars []*model.Car, err error) {
	db := global.DB.Model(&model.Car{})
	if status != nil {
		db = db.Where("status = ?", *status)
	}
	result := db.Find(&cars)
	if result.Error != nil {
		return nil, result.Error
	}
	return cars, nil
}

func CountCarsByStatus(status int8) (int64, error) {
	var count int64
	result := global.DB.Model(&model.Car{}).Where("status = ?", status).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// GetCarByID 根据汽车ID获取汽车信息
func GetCarByID(carID uint) (car *model.Car, err error) {
	result := global.DB.Where("car_id = ?", carID).First(&car)
	if result.Error != nil {
		return nil, result.Error
	}
	return car, nil
}

// GetCarByLicensePlate 根据车牌号获取汽车信息
func GetCarByLicensePlate(licensePlate string) (car *model.Car, err error) {
	result := global.DB.Where("license_plate = ?", licensePlate).First(&car)
	if result.Error != nil {
		return nil, result.Error
	}
	return car, nil
}

// UpdateCar 更新汽车信息
func UpdateCar(car *model.Car) (err error) {
	result := global.DB.Save(car)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteCarByID 根据汽车ID删除汽车信息
func DeleteCarByID(carID uint) (err error) {
	result := global.DB.Where("car_id = ?", carID).Delete(&model.Car{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("汽车信息不存在")
	}
	return nil
}

// GetAvailableCars 获取可用的汽车列表
func GetAvailableCars() (cars []*model.Car, err error) {
	result := global.DB.Where("status = ?", 0).Find(&cars)
	if result.Error != nil {
		return nil, result.Error
	}
	return cars, nil
}
