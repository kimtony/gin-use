package main

import (
	"fmt"
	"gin-use/bootstrap"
	"gin-use/configs"
	_ "gin-use/docs"
	"gin-use/src/global"
	"gin-use/src/routes"

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
	//系统初始化
	bootstrap.Init()

	defer global.DB.DbRClose()
	defer global.DB.DbWClose()

	// 初始化 HTTP 服务
	engine := routes.InitRouter()
	if err := engine.Run(fmt.Sprintf(":%s", configs.ProjectPort())); err != nil {
		global.Logger.Errorf("HTTP Server启动失败, err:%v", err)
	}

}
