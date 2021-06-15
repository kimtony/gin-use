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

// AccountInfo 用户个人信息
// @Summary request.WechatAccount
// @Description 用户个人信息
// @Tags account服务
// @Accept application/json
// @Produce application/json
// @Param string query string false "string valid" minlength(5) maxlength(10)
// @Param data body model.Account  true "微信参数"
// @Success 200 {object} model.Account
// @Router /api/v1/account/info [post]
func AccountInfo(c *gin.Context) {

	//参数校验
	if err = validator.ParseRequest(c, &reqAccount); err != nil {
        return
    }

	controller.ResponseHttpOK("ok", "请求成功", nil, c)
}
