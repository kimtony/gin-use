package middleware

import (
	"fmt"
	"gin-use/src/controller"
	"gin-use/src/util"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// 加载当前用户
func LoadCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从loadCurrentApi 获取存在swagger.josn的 api
		currentApi, _ := c.Get("currentApi") //需将interface转为String类型

		//检查接口是否需要做token校验   security
		isNeedToken := gjson.Get(currentApi.(string), "security")
		if isNeedToken.String() != "" {
			fmt.Println("要做token校验")
			token := c.Request.Header.Get("Authorization")
			if token == "" {
				controller.Response("not_auth", "未登录或非法访问", nil, c)
                c.AbortWithStatus(401)
				return
			}
			//解析token
			verfiy, err := util.ParseToken(token)
			if err != nil {
				println("token无效!", fmt.Sprintf("%s", err))
				controller.Response("not_auth", fmt.Sprintf("%s", err), nil, c)
                c.AbortWithStatus(403)
				return
			}

			tokenPraseContent := fmt.Sprintf("%v", verfiy)
			println("---tokenPraseContent-----", tokenPraseContent)

			accountId := verfiy["account"]
			fmt.Printf("-------accountId------", accountId)
            //TODO:还需要对角色解析

		}

	}
}
