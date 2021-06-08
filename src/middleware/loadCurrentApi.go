package middleware

import (
	"fmt"
	"gin-use/src/controller"
	"gin-use/src/global"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// 加载当前api
func LoadCurrentApi() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取请求的URI
		requestApi := c.Request.URL.RequestURI()
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
			controller.Response("not_exist", "不存在api!", nil, c)
			c.AbortWithStatus(404)
			return
		}
		//请求方法不对
		isExistApiMethod := gjson.Get(isExistApi.String(), requestMethod)
		if isExistApiMethod.String() == "" {
			controller.Response("not_exist", "不存在api!", nil, c)
			c.AbortWithStatus(404)
			return
		}

		fmt.Println("---currentApi-----", isExistApiMethod)
		
 		//gjson 传的值需要string 所以在locadCurrentApi的时候 设置currentApi类型为string json
		c.Set("currentApi", isExistApiMethod.String())
	}
}
