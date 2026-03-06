package router

import (
	"fmt"
	"net/http"
	"time"

	carHandler "car.rental/internal/api/car"
	inspectionHandler "car.rental/internal/api/inspection"
	orderHandler "car.rental/internal/api/order"
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
	// 登陆发送短信
	auth := engine.Group("/car_rental/v1/auth")
	{
		auth.POST("/login", userHandler.Login)
		auth.POST("/send_sms", userHandler.SendSMS)
	}
	api := engine.Group("/car_rental/v1/cars")
	api.Use(middlewares.JWTAuth())
	{
		api.GET("/list", carHandler.GetListCars)
		api.GET("/detail/:id", carHandler.GetCarDetail)
		api.POST("/create", carHandler.CreateCar)
		api.POST("/update", carHandler.UpdateCarInfo)
		api.DELETE("/delete/:id", carHandler.DeleteCar)
	}
	// 检测报告路由
	inspections := api.Group("/inspections").Use(middlewares.JWTAuth())
	{
		inspections.POST("create", inspectionHandler.CreateInspection)       // 发起检测
		inspections.GET("list", inspectionHandler.GetInspectionList)         // 获取检测报告列表
		inspections.GET("detail/:id", inspectionHandler.GetInspectionDetail) // 检测单详情
		inspections.POST("update", inspectionHandler.UpdateInspection)       // 检测更新
	}

	api.Use(middlewares.JWTAuth(), middlewares.AdminOnly())
	{
		api.GET("/user_list", userHandler.GetUserList)
	}
	api.Use(middlewares.JWTAuth())
	{
		api.GET("/user/deatil/:id", userHandler.GetUserInfo)
		api.POST("/user/update", userHandler.UpdateUserInfo)
	}

	// 订单路由
	orders := api.Group("/orders").Use(middlewares.JWTAuth())
	{
		orders.GET("/list", orderHandler.GetOrderList)         // 订单列表
		orders.POST("/create", orderHandler.CreateOrder)       // 创建订单
		orders.GET("/detail/:id", orderHandler.GetOrderDetail) // 订单详情
		orders.POST("/operate", orderHandler.OrderAction)      // 车辆操作（取车、还车、续租）
		orders.POST("/cancel", orderHandler.CancelOrder)       // 取消订单
	}

	return engine
}
