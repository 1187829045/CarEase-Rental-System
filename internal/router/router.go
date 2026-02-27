package router

import (
	"fmt"
	"net/http"
	"time"

	carHandler "car.rental/internal/api/car"
	// "car.rental/internal/api/inspection"
	reservationHandler "car.rental/internal/api/reservation"
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
	orders := engine.Group("/car_rental/v1/orders")
	orders.Use(middlewares.JWTAuth())
	{
		orders.GET("list", reservationHandler.ListOrders)                      //查询订单列表
		orders.POST("create", reservationHandler.CreateOrder)                  //创建订单
		orders.GET("detail/:id", reservationHandler.GetOrderDetail)            //查询订单详情
		orders.POST("pickup/:id", reservationHandler.PickupOrder)              //.car 取车
		orders.POST("return/:id", reservationHandler.ReturnOrder)              //.car 还车
		orders.POST("extend/:id", reservationHandler.ExtendOrder)              //.order 订单延长
		orders.POST("damage_report/:id", reservationHandler.DamageReportOrder) //.order 订单损坏上报
		orders.POST("cancel/:id", reservationHandler.CancelOrder)              //.order 订单取消
	}
	reservations := api.Group("/reservations")
	{
		reservations.POST("create", reservationHandler.CreateReservation)       //创建预定单
		reservations.GET("list", reservationHandler.ListReservations)           //查询预定单列表
		reservations.GET("detail/:id", reservationHandler.GetReservationDetail) //查询于订单详情
		reservations.PUT("update/:id", reservationHandler.UpdateReservation)    //更新预定单
		reservations.POST("cancel/:id", reservationHandler.CancelReservation)   //取消预定单
		reservations.POST("confirm/:id", reservationHandler.ConfirmReservation) //确认预定单
	}

	// // 检测报告路由
	// inspections := api.Group("/inspections")
	// {
	// 	inspections.GET("", inspection.GetInspectionList)           // 获取检测报告列表
	// 	inspections.GET(":id", inspection.GetInspectionDetail)     // 获取检测报告详情
	// 	inspections.PUT(":id", inspection.UpdateInspection)        // 更新检测报告
	// }
	api.Use(middlewares.JWTAuth(), middlewares.AdminOnly())
	{
		api.GET("/user_list", userHandler.GetUserList)
		api.GET("/user/:id", userHandler.GetUserInfo)
		api.PUT("/user/:id", userHandler.UpdateUserInfo)
	}
	return engine
}
