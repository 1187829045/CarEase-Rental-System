package model

import "time"

// 支付记录表
// Payment 支付记录模型
type Payment struct {
	BaseModel
	PaymentID      uint      `gorm:"column:payment_id;primaryKey;autoIncrement"`     // 支付记录ID
	OrderID        uint      `gorm:"column:order_id;not null;index"`                  // 订单ID
	Amount         float64   `gorm:"column:amount;type:decimal(10,2);not null;index"` // 支付金额
	PaymentType    string    `gorm:"column:payment_type;type:varchar(20);not null;index"`   // 支付类型：租金/附加服务/罚款
	PaymentMethod  string    `gorm:"column:payment_method;type:varchar(20);not null;index"` // 支付方式：支付宝/微信/银行卡
	TransactionID  string    `gorm:"column:transaction_id;type:varchar(100);index"`        // 交易ID
	Status         string    `gorm:"column:status;type:varchar(20);not null;index"`         // 支付状态：待支付/已支付/失败
	Description    string    `gorm:"column:description;type:text"`                    // 支付描述
	PaymentTime    time.Time `gorm:"column:payment_time;type:datetime;index"`               // 支付时间
}

// Payment 指定表名
func (Payment) TableName() string {
	return "payments"
}
