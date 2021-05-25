package controller

import (
	"gin-use/src/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type returnStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//健康检查
func Health(c *gin.Context) {

	// 月份 1,01,Jan,January
	// 日　 2,02,_2
	// 时　 3,03,15,PM,pm,AM,am
	// 分　 4,04
	// 秒　 5,05
	// 年　 06,2006
	// 时区 -07,-0700,Z0700,Z07:00,-07:00,MST
	// 周几 Mon,Monday

	now := time.Now()
	id := util.GenerateId()
	c.JSON(http.StatusOK, map[string]interface{}{
		"Time": now,
		"id":   id,
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
