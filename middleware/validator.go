package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/xeipuuv/gojsonschema"
)

//
func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {

		//静态检查
		staticCheck()

		//读取service_define
		bytes, err := ioutil.ReadFile("service_define.json")
		if err != nil {
			panic("read service_define.json error")
		}

		requestUrl := c.Request.RequestURI
		requesrMethod := c.Request.Method
		fmt.Printf("******** 获取到请求的接口:%#v\n", requestUrl+"  "+requesrMethod)

		//校验
		var capi string
		for i := 0; true; i++ {

			//json读取service_define
			api := gjson.Get(string(bytes), "apis|"+strconv.Itoa(i))
			url := gjson.Get(api.String(), "url")
			method := gjson.Get(api.String(), "method")

			//校验通过
			if url.String() == requestUrl && method.String() == requesrMethod {
				println("******** api verfiy success ********")
				capi = api.String()
				break
			}
			if api.String() == "" {
				panic("service_define_not_connent")
			}
		}
		if capi == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"status": gin.H{
					"status_code": http.StatusNotFound,
					"status":      "fail",
				},
				"message": "api_not_found !!!",
			})
			panic("404")
		}
		c.Next()
	}

}

//静态检查
func staticCheck() {

	//读取json文件
	bytes, err := ioutil.ReadFile("service_define.json")
	if err != nil {
		panic("read service_define.json error")
	}
	bytes1, err1 := ioutil.ReadFile("static_check.json")
	if err1 != nil {
		panic("read static_check.json error")
	}

	//转为json格式
	schemaLoader := gojsonschema.NewStringLoader(string(bytes1))  // json格式
	documentLoader := gojsonschema.NewStringLoader(string(bytes)) // 待校验的json数据

	//校验
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("service_define static check is success\n")
	} else {
		fmt.Printf("service_define static check is fail :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
