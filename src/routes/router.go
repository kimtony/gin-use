package routes

import (
	"gin-use/configs"
	"gin-use/src/controller"
	"gin-use/src/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	}

	//pprof
	pprof.Register(r, "/sys/pprof")

	//prometheus
	r.GET("/sys/metrics", gin.WrapH(promhttp.Handler()))

	//swagger接口文档
	url := ginSwagger.URL(configs.SwaggerUrl())
	r.GET("/sys/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
