package controller

import (
	"gin-use/src/global"
	"gin-use/src/util/snowflake"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type returnStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Health 健康检查
// @Summary 健康检查接口
// @Description 服务是否启动正常检查
// @Tags 监测服务
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200
// @Router /api/health [get]
func Health(c *gin.Context) {

	now := time.Now()
	id := snowflake.GenerateId()
	workerid, datacenterid := snowflake.GetDeviceID(id)
	global.Logger.Info("嘿，我能调用了")

	c.JSON(http.StatusOK, map[string]interface{}{
		"Time":         now,
		"id":           id,
		"datacenterid": datacenterid,
		"workerid":     workerid,
	})

}

/**
 * 定义成功返回的函数
 * param: interface{} data
 * return: SuccessStruct
 */
func ReturnSuccess(data interface{}) returnStruct {
	var successStruct returnStruct

	successStruct.Code = 200
	successStruct.Data = data
	successStruct.Message = "success"

	return successStruct
}

/**
 * 定义失败返回的函数
 * param: interface{} data
 * return: SuccessStruct
 */
func ReturnFaild(message string) returnStruct {
	var faildStruct returnStruct

	faildStruct.Code = 400
	faildStruct.Data = nil
	faildStruct.Message = message

	return faildStruct
}
