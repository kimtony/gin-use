package v1

import (
	"gin-use/src/controller"

	"github.com/gin-gonic/gin"
)

// AccountInfo 用户个人信息
// @Summary 用户个人信息
// @Description 用户个人信息
// @Tags account服务
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Security token
// @param token header string true "Authorization"
// @Security OAuth2Application[account]
// @Success 200 {object} model.Account
// @Router /api/account/info [get]
func AccountInfo(c *gin.Context) {
	// data :=
	controller.Response("ok", "请求成功", nil, c)
}
