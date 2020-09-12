package main

import (
	"gin-demo/Databases"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"

	"gin-demo/Routes"
)

func main() {

	//数据库
	Databases.Init()

	//初始化路由
	r := Routes.InitRouter()

	r.Run(":" + "8081")

	defer Databases.DB.Close()

}
