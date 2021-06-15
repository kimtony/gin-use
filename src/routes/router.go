package routes

import (
	"gin-use/configs"
	"gin-use/src/controller"
	v1 "gin-use/src/controller/v1"
	"gin-use/src/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r = gin.Default()

func InitRouter() *gin.Engine {

	api := r.Group("/api")
	{
		api.Use(
			middleware.CORSMiddleware(),  //cors跨域
			// middleware.LoadCurrentApi(),  //加载当前api
			// middleware.LoadCurrentUser(), //加载当前用户
			// middleware.ApiValidator(),    //api参数校验
		)

		//健康检查
		api.GET("/health", controller.Health)

		apiV1 := api.Group("/v1")
		{
			apiV1.POST("/account/info", v1.AccountInfo)
			apiV1.GET("/account/wechat", v1.Wechat)
		}
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
