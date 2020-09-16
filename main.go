package main

import (
	"gin-demo/database"

	// "os"
	// "strconv"

	"gin-demo/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var r = gin.Default()

func main() {

	//是否开启数据库调试
	// enableDBLogMode, _ := strconv.ParseBool(os.Getenv("ENABLE_DB_LOGMODE"))
	// database.DB.LogMode(enableDBLogMode)

	test()
	//初始化路由
	r := routes.InitRouter()

	r.Run(":" + "8081")

	defer database.DB.Close()

}
