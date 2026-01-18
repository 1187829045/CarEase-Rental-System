package dao

import (
	"car.rental/dao/model"
	"car.rental/global"
	"errors"
)

// CreateUserExtra 创建用户扩展信息
func CreateUserExtra(userExtra model.UserExtra) (err error) {
	result := global.DB.Create(&userExtra)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserExtraByUserID 根据用户ID获取扩展信息
func GetUserExtraByUserID(userID uint) (userExtra *model.UserExtra, err error) {
	result := global.DB.Where("user_id = ?", userID).First(&userExtra)
	if result.Error != nil {
		return nil, result.Error
	}
	return userExtra, nil
}

// UpdateUserExtra 更新用户扩展信息
func UpdateUserExtra(userExtra model.UserExtra) (err error) {
	result := global.DB.Save(&userExtra)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUserExtraByUserID 根据用户ID删除扩展信息
func DeleteUserExtraByUserID(userID uint) (err error) {
	result := global.DB.Where("user_id = ?", userID).Delete(&model.UserExtra{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("用户扩展信息不存在")
	}
	return nil
}
