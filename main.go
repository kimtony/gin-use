package main

import (
	"gin-demo/database"
	"gin-demo/middleware"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"

	"gin-demo/routes"
)

var r = gin.Default()

func main() {
	//middleware
	r.Use(middleware.TestRead())
	//数据库
	database.Init()
	println(os.Getenv("DB_HOST"), os.Getenv("DB_USER"))
	//初始化路由
	r := routes.InitRouter()

	r.Run(":" + "8081")

	defer database.DB.Close()

}