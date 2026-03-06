package consts

const (
	JWTIssuer       = "3f1c8b46-7d29-4e5b-9b12-89f3c6a0d1e2"
	UserRoleID      = 0 // 普通用户角色ID
	AdminRoleID     = 1 // 管理员角色ID
	InspectorRoleID = 2 // 检测师角色ID

	// 短信相关常量
	SMSLimitKeyPrefix     = "sms:limit:"
	SMSCounterKeyPrefix   = "sms:counter:"
	SMSBlacklistKeyPrefix = "sms:blacklist:"
)

// 订单操作类型常量
const (
	ActionPickup = "pickup" // 取车
	ActionReturn = "return" // 还车
	ActionExtend = "extend" // 续租
)

// 订单状态常量
const (
	OrderStatusPending   = 0 // 待取车
	OrderStatusConfirmed = 1 // 已取车-进行中
	OrderStatusCanceled  = 2 // 已取消
	OrderStatusCompleted = 3 // 已完成
)

// 车辆状态常量
const (
	CarStatusPending   = 0 // 待检测
	CarStatusAvailable = 1 // 可用
	CarStatusRented    = 2 // 已出租
)
