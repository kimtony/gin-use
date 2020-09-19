package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type error interface {
	Error() string
}

var DB *sql.DB

func Init() error {

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		println("failed to connect database", err.Error())
		panic(err)
	}
	// defer db.Close()
	err = DB.Ping()
	if err != nil {
		println("failed to connect database", err.Error())
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return nil

}
