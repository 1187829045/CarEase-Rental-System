package router

import (
	api2 "car.rental/internal/user/api"
	"fmt"
	"net/http"
	"time"
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
	app := engine.Group("/car_rental")
	{
		app.POST("/login", api2.Login)
		app.POST("/register", api2.Register)
	}
	return engine
}
