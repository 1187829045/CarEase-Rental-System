package user

import "time"

type UserInfo struct {
	ID       uint       `json:"id"`
	Mobile   string     `json:"mobile"`
	UserName string     `json:"userName"`
	Birthday *time.Time `json:"birthday"`
	Gender   int8       `json:"gender"`
	Role     string     `json:"role"`
}

// UserListResponse 用户列表响应结构体
type UserListResponse struct {
	List       []UserInfo       `json:"list"`        // 用户列表
	Total      int64            `json:"total"`       // 总用户数
	Page       int              `json:"page"`        // 当前页码
	PageSize   int              `json:"page_size"`   // 每页大小
	RoleCounts map[string]int64 `json:"role_counts"` // 各角色数量
}
type UserUpdateForm struct {
	ID       uint       `json:"id" binding:"required"`
	Mobile   *string    `json:"mobile"`
	UserName *string    `json:"userName"`
	Birthday *time.Time `json:"birthday"`
	Gender   *int8      `json:"gender"`
	Role     string     `json:"role"`
}
