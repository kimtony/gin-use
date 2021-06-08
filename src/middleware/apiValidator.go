package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// 接口校验
func ApiValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从loadCurrentApi 获取存在swagger.josn的 api
		currentApi, _ := c.Get("currentApi") //需将interface转为String类型

		//校验请求参数
		requireParams := gjson.Get(currentApi.(string), "parameters")

		var reqData = map[string]interface{}{}
		//遍历接口需要的参数个数
		for _, requireParam := range requireParams.Array() {
			// println("-----name-------", requireParam.String())
			fromType := gjson.Get(requireParam.String(), "in").String()
			requireParamName := gjson.Get(requireParam.String(), "name").String()
			requireParamType := gjson.Get(requireParam.String(), "type").String()
			var tmpData = "123"
			switch fromType {
			case "query":
				tmpData = c.Query(requireParamName)
				break
			case "body":
				tmpData = c.PostForm(requireParamName)
				break
			case "header":
				tmpData = c.GetHeader(requireParamName)
				break
			case "cookie":
				tmpData, _ = c.Cookie(requireParamName)
				break
			case "path":
				tmpData = c.Param(requireParamName)
				break
			default:
				tmpData = ""
				break
			}
			if requireParamName != "" && tmpData == "" {
				println("请求参数不能为空 ")
			}

			reqData[requireParamName] = tmpData
			println("-----tmpData-------", tmpData)
			println("-----fromType-------", fromType)
			println("-----requireParamName-------", requireParamName)
			println("-----requireParamType-------", requireParamType)
		}

		fmt.Println("reqData----------", reqData)
	}
}
