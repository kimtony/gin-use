package routes

import (
	"gin-demo/controller"
	"gin-demo/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func InitRouter() *gin.Engine {

	// r.Use(Middleware.Session("xiaojipu"))
	r.Use(middleware.CORSMiddleware())
	
	api := r.Group("/api")
	{

		//健康检查
		api.GET("/health", controller.Health)

		api.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"hello": "gin",
			})
		})

		api.GET("/activity", controller.Activity)
	}

	return r
}
