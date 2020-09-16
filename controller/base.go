package controller

import (
	Helpers "gin-demo/helpers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type returnStruct struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
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
	id := Helpers.GenerateId()
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
