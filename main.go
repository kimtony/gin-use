package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/gin-gonic/gin"
	"gin-use/src/routes"
	"gin-use/configs"
	"gin-use/src/util/logger"
	"gin-use/src/util/env"
	_ "gin-use/docs" // 千万不要忘了导入把你上一步生成的docs
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

// @host 127.0.0.1:8081
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

	// 初始化 HTTP 服务
	r := routes.InitRouter()

    if err := r.Run(":8081"); err != nil {
        fmt.Println("HTTP Server启动失败")
    }

	//是否开启数据库调试
	// enableDBLogMode, _ := strconv.ParseBool(os.Getenv("ENABLE_DB_LOGMODE"))
	// db.DB.LogMode(enableDBLogMode)

	//初始化路由
	// r := routes.InitRouter()

    // if err := r.Run(":8081"); err != nil {
    //     zap.L().Fatal("HTTP Server启动失败", zap.Error(err))
    // }

	//sentry  util.Sentry()
	



}
