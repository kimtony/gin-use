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
// @Param data body request.WechatAccount  true "微信参数"
// @param token header string true "Authorization"
// @Security OAuth2Application[account]
// role ["account"]
// @Success 200
// @Router /api/wechat [post]
func Wechat(c *gin.Context) {
	
	controller.ResponseHttpOK("ok","请求成功",nil,c)
}
