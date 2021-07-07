package v1

import (
	"gin-use/src/controller"
	"gin-use/src/model/request"
	"gin-use/src/util/validator"
	"github.com/gin-gonic/gin"
)



var (
	err error
	reqAccount request.Account
)


func AccountInfo(c *gin.Context) {

	//参数校验
	if err = validator.ParseRequest(c, &reqAccount); err != nil {
        return
    }

	controller.ResponseHttpOK("ok", "请求成功", nil, c)
}
