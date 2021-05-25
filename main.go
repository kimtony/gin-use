package main

import (
	// "fmt"
	"gin-use/src/util/database"
	// "gin-use/src/util"
	"os"
	"strconv"
	// "encoding/json"

	// "os"
	// "strconv"
	"gin-use/src/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"gin-use/src/util/logger"
	"go.uber.org/zap"
	_ "gin-use/docs" // 千万不要忘了导入把你上一步生成的docs
)

var r = gin.Default()

func init() {
    // 初始化日志库
    log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, "./logs/gin-use.log")
}


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

  
	//是否开启数据库调试
	enableDBLogMode, _ := strconv.ParseBool(os.Getenv("ENABLE_DB_LOGMODE"))
	database.DB.LogMode(enableDBLogMode)

	//初始化路由
	r := routes.InitRouter()

    if err := r.Run(":8081"); err != nil {
        zap.L().Fatal("HTTP Server启动失败", zap.Error(err))
    }

	//sentry  util.Sentry()
	

	defer database.DB.Close()

}
