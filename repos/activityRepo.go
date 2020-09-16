package repos

import (
	"gin-demo/database"
)

// Activity struct
type Activity struct {
	ID   string `json: "id"`
	Name string `json: "name"`
}

func FirstActvity() Activity {

	var activity Activity

	//查询数据
	database.DB.First(&activity)

	return activity
}

func SelectActvity() []Activity {

	var activity []Activity

	//查询数据
	database.DB.Find(&activity)

	return activity
}
