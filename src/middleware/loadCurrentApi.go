package middleware

import (
	"gin-use/src/controller"
	"gin-use/src/global"
	"io/ioutil"
	"strings"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// 加载当前api
func LoadCurrentApi() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取请求的URI  有可能url传参,根据问号分割 拿到url
		// requestApi := c.Request.URL.RequestURI()
		requestApi := strings.Split(c.Request.URL.RequestURI(), "?")[0]
		fmt.Println("-----requestApi----",requestApi)
		//获取请求方法  拿到的method为大写需要转为小写
		requestMethod := strings.ToLower(c.Request.Method)
		global.Logger.Debugf("---获取到请求的接口--- : %v", requestApi+"  "+requestMethod)

		//读取swagger.json   获取swagger所有api
		bytes, err := ioutil.ReadFile("./docs/swagger.json")
		if err != nil {
			panic(err)
		}
		apis := gjson.Get(string(bytes), "paths")

		//获取该请求的url是否存在swagger文档里面
		isExistApi := gjson.Get(apis.String(), requestApi)

		if isExistApi.String() == "" {
			controller.Response(http.StatusNotFound,"not_exist", "不存在api!", nil, c)
			c.Abort()
			return
		}
		//请求方法不对
		isExistApiMethod := gjson.Get(isExistApi.String(), requestMethod)
		if isExistApiMethod.String() == "" {
			controller.Response(http.StatusNotFound,"not_exist", "不存在api!", nil, c)
			c.Abort()
			return
		}

		// fmt.Println("---currentApi-----", isExistApiMethod)

 		//gjson 传的值需要string 所以在locadCurrentApi的时候 设置currentApi类型为string json
		c.Set("currentApi", isExistApiMethod.String())
	}
}
