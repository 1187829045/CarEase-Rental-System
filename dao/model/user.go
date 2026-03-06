package model

import (
	"time"
)

type User struct {
	BaseModel
	UserId   uint       `gorm:"column:user_id;primaryKey;autoIncrement"`
	Mobile   string     `gorm:"column:mobile;index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"column:password;type:varchar(100);not null"`
	UserName string     `gorm:"column:user_name;type:varchar(20);index"`
	Birthday *time.Time `gorm:"column:birthday;type:datetime;index"`
	Gender   int8       `gorm:"column:gender;default:male;type:varchar(6) comment '0表示女, 1表示男';index"`
	Role     string     `gorm:"column:role;default:'user';index"`      // 角色标识
	BankCard string     `gorm:"column:bank_card;type:varchar(20);index"` // 银行卡号
	IDCard   string     `gorm:"column:id_card;type:varchar(18);index"`   // 身份证号
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
