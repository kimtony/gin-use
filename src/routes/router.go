package routes

import (
	"gin-use/src/middleware"
	"gin-use/src/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r = gin.Default()

func InitRouter() *gin.Engine {
	

	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api")
	{

		//健康检查
		api.GET("/health", controller.Health)

		//员工
		// api.GET("/staff", controller.Staff)

		//redis测试
		// api.GET("/redis/test", controller.RedisTest)

		//token
		// api.GET("/token/test", controller.Token)

		//文档
		url := ginSwagger.URL("http://192.168.1.163:8081/swagger/doc.json")
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	}

	return r
}
