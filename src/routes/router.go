package routes

import (
	"gin-use/src/controller"
	"gin-use/src/middleware"
	
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/zap"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "time"
)

var r = gin.Default()

func InitRouter() *gin.Engine {
    engine := gin.New()
    // 使用zap日志库
    engine.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
    engine.Use(ginzap.RecoveryWithZap(zap.L(), true))
	// r.Use(Middleware.Session("xiaojipu"))
	// r.Use(middleware.Validator())
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