package _struct

type User struct {
	Username string `json:"username" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestPrize struct {
	UserName string `json:"username" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
}
