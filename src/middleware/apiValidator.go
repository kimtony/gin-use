package middleware

import (
	"fmt"
	// "gin-use/src/controller"
	// "io/ioutil"
	// "strconv"
	"reflect"
	// "gin-use/src/util/validator"
	"gin-use/src/model/request"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

var err error

// 接口校验
func ApiValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从loadCurrentApi 获取存在swagger.josn的 api
		currentApi, _ := c.Get("currentApi") //需将interface转为String类型
		var req request.WechatAccount

		fmt.Println("-----request.WechatAccount----", &req)
		requireParams := gjson.Get(currentApi.(string), "summary")
		req1 := requireParams
	
		// 获取结构体实例的反射类型对象
		typeOfCat := reflect.TypeOf(req1)
		fmt.Println("-----requireParams----", typeOfCat)
		// if err = validator.ParseRequest(c, &req1); err != nil {
		// 	return
		// }

	
	}
}
