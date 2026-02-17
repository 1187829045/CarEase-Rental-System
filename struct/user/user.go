package user

import "time"

type UserInfo struct {
	ID       uint       `json:"id"`
	Mobile   string     `json:"mobile"`
	UserName string     `json:"userName"`
	Birthday *time.Time `json:"birthday"`
	Gender   int8       `json:"gender"`
	Role     int8       `json:"role"`
}
type UserUpdateForm struct {
	ID       uint       `json:"id" binding:"required"`
	Mobile   *string    `json:"mobile"`
	UserName *string    `json:"userName"`
	Birthday *time.Time `json:"birthday"`
	Gender   *int8      `json:"gender"`
	Role     *int8      `json:"role"`
}
