package Routes

import (
	"gin-demo/Controllers/Api"
	"gin-demo/Middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func InitRouter() *gin.Engine {

	// r.Use(Middleware.Session("xiaojipu"))
	r.Use(Middleware.CORSMiddleware())

	api := r.Group("/api")
	{

		//健康检查
		api.GET("/health", Api.Health)

		api.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"hello": "kenny",
			})
		})
	}

	return r
}
