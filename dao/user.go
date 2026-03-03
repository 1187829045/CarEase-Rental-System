package dao

import (
	"errors"

	"car.rental/consts"
	"car.rental/dao/model"
	"car.rental/global"
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

// ListUsersWithPagination 分页获取用户列表，支持角色筛选
func ListUsersWithPagination(page, pageSize int, role string) (users []*model.User, total int64, err error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 构建查询
	query := global.DB.Model(&model.User{})

	// 角色筛选
	if role != "" {
		query = query.Where("role LIKE ?", "%"+role+"%")
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	result := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, total, nil
}
