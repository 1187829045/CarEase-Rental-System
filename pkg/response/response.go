package response

import (
	"github.com/gin-gonic/gin"
)

// Response 通用响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  "ok",
		Data: data,
	})
}

// SuccessMsg 成功响应（仅消息）
func SuccessMsg(c *gin.Context, msg string) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  msg,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{
		Code: code,
		Msg:  msg,
	})
}

// BadRequest 请求参数错误
func BadRequest(c *gin.Context, msg string) {
	Error(c, 400, msg)
}

// InternalError 内部服务器错误
func InternalError(c *gin.Context, msg string) {
	Error(c, 500, msg)
}

// TooManyRequests 请求过于频繁
func TooManyRequests(c *gin.Context, msg string) {
	Error(c, 429, msg)
}

// Unauthorized 未授权
func Unauthorized(c *gin.Context, msg string) {
	Error(c, 401, msg)
}

// Forbidden 禁止访问
func Forbidden(c *gin.Context, msg string) {
	Error(c, 403, msg)
}