package api

import (
	"context"
	"time"

	"car.rental/consts"
	"car.rental/global"
	"car.rental/pkg/response"
	_struct "car.rental/struct"
	"github.com/gin-gonic/gin"
)

func SendSMS(c *gin.Context) {
	form := _struct.SendSMSForm{}
	if err := c.ShouldBind(&form); err != nil {
		response.BadRequest(c, consts.ErrInvalidParameter)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 检查是否在黑名单中
	blacklistKey := consts.SMSBlacklistKeyPrefix + form.Mobile
	blacklisted, err := global.RedisClient.Exists(ctx, blacklistKey).Result()
	if err != nil {
		response.InternalError(c, "系统繁忙，请稍后再试")
		return
	}
	if blacklisted == 1 {
		response.Error(c, 403, "该手机号已被限制发送短信")
		return
	}

	// 检查1分钟内是否重复发送
	limitKey := consts.SMSLimitKeyPrefix + form.Mobile
	exists, err := global.RedisClient.Exists(ctx, limitKey).Result()
	if err != nil {
		response.InternalError(c, "系统繁忙，请稍后再试")
		return
	}
	if exists == 1 {
		response.TooManyRequests(c, "发送过于频繁,请1分钟后再试")
		return
	}

	// 检查10分钟内发送次数
	counterKey := consts.SMSCounterKeyPrefix + form.Mobile
	count, err := global.RedisClient.Incr(ctx, counterKey).Result()
	if err != nil {
		response.InternalError(c, "系统繁忙，请稍后再试")
		return
	}
	// 设置计数器过期时间为10分钟
	if count == 1 {
		if err := global.RedisClient.Expire(ctx, counterKey, 10*time.Minute).Err(); err != nil {
			// 记录错误但不影响响应
		}
	}
	// 如果10分钟内发送了5次，加入黑名单
	if count >= 5 {
		if err := global.RedisClient.Set(ctx, blacklistKey, "1", time.Hour).Err(); err != nil {
			// 记录错误但不影响响应
		}
		response.Error(c, 403, "发送次数过多，该手机号已被限制")
		return
	}

	// 设置1分钟内的发送限制
	if err := global.RedisClient.Set(ctx, limitKey, "1", time.Minute).Err(); err != nil {
		// 记录错误但不影响响应
	}
	// TODO: 实现具体的短信发送逻辑

	response.SuccessMsg(c, "短信发送成功")
}
