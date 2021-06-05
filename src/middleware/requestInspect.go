package middleware

import (
	"fmt"
	// "io/ioutil"

	"github.com/gin-gonic/gin"
)

// 拦截器
func RequestInspect() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取请求的URI
		requestApi := c.Request.URL.RequestURI()
		// 获取请求方法
		requestMethod := c.Request.Method
		fmt.Printf("----- 获取到请求的接口:%#v\n", requestApi+"  "+requestMethod)

	}
}
