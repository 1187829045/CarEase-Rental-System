package router

import (
	"fmt"
	"net/http"
	"time"

	api2 "car.rental/internal/api"

	// adminrouter "go-router/internal/app/admin/router
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
		auth.POST("/login", api2.Login)
		auth.POST("/send_sms", api2.SendSMS)
	}
	api := engine.Group("/car_rental/v1")
	{
		api.GET("/cars", api2.ListCars)
		api.GET("/cars/:id", api2.GetCar)
		api.PUT("/cars/:id", api2.UpdateCar)
	}
	return engine
}
