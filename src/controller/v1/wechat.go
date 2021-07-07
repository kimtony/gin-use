package v1


import (
	"gin-use/src/controller"
	"github.com/gin-gonic/gin"


)


func Wechat(c *gin.Context) {
	
	controller.ResponseHttpOK("ok","请求成功",nil,c)
}
