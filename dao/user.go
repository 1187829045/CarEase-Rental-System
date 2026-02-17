package dao

import (
	"car.rental/consts"
	"car.rental/dao/model"
	"car.rental/global"
	"errors"
)

func GetUserByMobile(mobile string) (user *model.User, err error) {
	var u model.User
	result := global.DB.Where("mobile = ?", mobile).First(&u)
	if result.RowsAffected == 0 {
		return nil, errors.New(consts.UserNotFound)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}

func GetUserByID(userID uint) (user *model.User, err error) {
	var u model.User
	result := global.DB.Where("user_id = ?", userID).First(&u)
	if result.RowsAffected == 0 {
		return nil, errors.New(consts.UserNotFound)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}

func CreateUser(user *model.User) (err error) {
	result := global.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUser(user *model.User) (err error) {
	result := global.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ListUsers() (users []*model.User, err error) {
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
