package router

import (
	"fmt"
	"net/http"
	"time"

	carHandler "car.rental/internal/api/car"
	userHandler "car.rental/internal/api/user"
	"car.rental/middlewares"
	"github.com/gin-gonic/gin"
)

// NewHTTPRouter create http router.
func NewHTTPRouter() *gin.Engine {
	engine := gin.New()
	fmt.Println("sever lottery in ", time.Now().Format(time.DateTime))
	engine.Use(gin.Recovery())
	// 加载404错误页面
	engine.NoRoute(func(c *gin.Context) {
		// 实现内部重定向
		c.JSON(http.StatusNotFound, gin.H{
			"title": "404 not found",
		})
	})
	// router group.
	auth := engine.Group("/car_rental/v1/auth")
	{
		auth.POST("/login", userHandler.Login)
		auth.POST("/send_sms", userHandler.SendSMS)
	}
	api := engine.Group("/car_rental/v1")
	api.Use(middlewares.JWTAuth())
	{
		api.GET("/cars", carHandler.GetListCars)
		api.GET("/cars/:id", carHandler.GetCarDetail)
		api.PUT("/cars/:id", carHandler.UpdateCarInfo)
	}
	api.Use(middlewares.JWTAuth(), middlewares.AdminOnly())
	{
		api.GET("/user_list", userHandler.GetUserList)
		api.GET("/user/:id", userHandler.GetUserInfo)
		api.PUT("/user/:id", userHandler.UpdateUserInfo)
	}
	return engine
}
