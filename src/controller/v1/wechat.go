package v1


import (
	"gin-use/src/controller"
	"github.com/gin-gonic/gin"


)

// Wechat 微信服务
// @Summary 微信服务
// @Description 微信服务
// @Tags 微信服务
// @Accept application/json
// @Produce application/json
// @Param name path string true "Name"
// role ["account"]
// @Success 200
// @Router /api/wechat [get]
func Wechat(c *gin.Context) {
	// data := 
	controller.Response("ok","请求成功",nil,c)
}
