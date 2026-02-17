package model

type UserExtra struct {
	BaseModel
	UserID   uint   `gorm:"column:user_id;type:uint;not null;index"` // 用户ID，关联用户表
	BankCard string `gorm:"column:bank_card;type:varchar(20)"`       // 银行卡号
	IDCard   string `gorm:"column:id_card;type:varchar(18)"`         // 身份证号

	// 外键关联
	User User `gorm:"foreignKey:UserID;references:UserId"`
}
