package model

import (
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
