package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

type error interface {
	Error() string
}

var DB *gorm.DB

func init() {

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		println("failed to connect database", err.Error())
	} else {
		fmt.Println("connect success")
	}

	if DB.Error != nil {
		println("failed to connect database", err.Error())
	}
	//gorm禁用表复数
	DB.SingularTable(true)

	// 数据库自动迁移
	// db.Debug().AutoMigrate(&Account{}, &Contact{})
}

// 返回数据库对象的指针
func GetDB() *gorm.DB {
	return DB
}
