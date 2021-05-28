package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"

	"fmt"

	"gin-use/configs"
	_ "gin-use/docs"
	"gin-use/src/global"
	"gin-use/src/routes"
	"gin-use/src/util/cache"
	"gin-use/src/util/db"
	"gin-use/src/util/env"
	"gin-use/src/util/logger"
)

var r = gin.Default()

// @title swagger 接口文档
// @version 2.0
// @description
// @contact.name
// @contact.url
// @contact.email
// @license.name MIT
// @license.url https://github.com/xinliangnote/go-gin-api/blob/master/LICENSE
// @host 192.168.1.163:8081
// @BasePath
func main() {

	// 初始化 logger
	loggers, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName(), env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.ProjectLogFile()),
	)
	if err != nil {
		panic(err)
	}
	defer loggers.Sync()
	global.Logger = loggers

	// 初始化数据库
	dbRepo, err := db.New()
	if err != nil {
		loggers.Error("new db err", zap.Error(err))
	}
	global.DB = dbRepo

	//初始化缓存服务
	cacheRepo, err := cache.New()
	if err != nil {
		loggers.Error("new cahe err", zap.Error(err))
	}
	global.Cache = cacheRepo

	// 初始化 HTTP 服务
	r := routes.InitRouter()
	if err := r.Run(configs.ProjectPort()); err != nil {
		fmt.Println("HTTP Server启动失败")
	}

}
