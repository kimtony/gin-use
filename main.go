package main

import (
	"encoding/json"
	"fmt"
	"gin-use/bootstrap"
	"gin-use/configs"
	_ "gin-use/docs"
	"gin-use/src/global"
	"gin-use/src/model"
	"gin-use/src/routes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	r = gin.Default()
)

// @title swagger 接口文档
// @version 2.0
// @description
// @contact.name
// @contact.url
// @contact.email
// @license.name MIT
// @license.url https://www.baidu.com
// @host 192.168.1.163:8081
// @BasePath
func main() {
	test()
	//系统初始化
	bootstrap.Init()

	defer global.DB.DbRClose()
	defer global.DB.DbWClose()

	// 初始化 HTTP 服务
	engine := routes.InitRouter()
	if err := engine.Run(fmt.Sprintf(":%s", configs.ProjectPort())); err != nil {
		fmt.Println("HTTP Server启动失败")
	}

}

func test() {
	//读取service_define
	bytes, err := ioutil.ReadFile("./docs/swagger.json")
	if err != nil {
		panic(err)
	}
	var swagger *model.Swagger
	json.Unmarshal([]byte(bytes), &swagger)
	// fmt.Println("---swagger------", swagger)

	for key, value := range swagger.Paths {
		fmt.Println("--key-value------", key, value)

		fmt.Println("--key-value------", key)

	}
}
