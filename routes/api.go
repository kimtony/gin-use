package routes

import (
	"gin-demo/controller"
	"gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func InitRouter() *gin.Engine {

	// r.Use(Middleware.Session("xiaojipu"))
	r.Use(middleware.Validator())

	api := r.Group("/api")
	{

		//健康检查
		api.GET("/health", controller.Health)

		//员工
		api.GET("/staff", controller.Staff)

		//redis测试
		api.GET("/redis/test", controller.RedisTest)

		//token
		api.GET("/token/test", controller.Token)

	}

	return r
}
