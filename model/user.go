package model

import (
	"car.rental/consts"
	"car.rental/global"
	"errors"
	"time"
)

type User struct {
	BaseModel
	UserId   uint       `gorm:"column:user_id;primaryKey;autoIncrement"`
	Mobile   string     `gorm:"column:mobile;index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"column:password;type:varchar(100);not null"`
	UserName string     `gorm:"column:user_name;type:varchar(20)"`
	Birthday *time.Time `gorm:"column:birthday;type:datetime"`
	Gender   int8       `gorm:"column:gender;default:male;type:varchar(6) comment '0表示女, 1表示男'"`
}

func GetUser(userID uint) (user User, err error) {
	result := global.DB.Where("user_id = ?", userID).Find(&user)
	if result.RowsAffected == 0 {
		return User{}, errors.New(consts.UserNotFound)
	}
	return user, nil
}
func GetUserByMobile(mobile string) (user User, err error) {
	result := global.DB.Where("mobile = ?", mobile).Find(&user)
	if result.RowsAffected == 0 {
		return User{}, errors.New(consts.UserNotFound)
	}
	return user, nil
}

func CreateUser(user User) (err error) {
	result := global.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
