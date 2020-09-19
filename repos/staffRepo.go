package repos

import (
	"gin-demo/database"
)

// Activity struct
type Staff struct {
	ID     string `json: "id"`
	Name   string `json: "name"`
	Mobile string `json: "mobile"`
}

func FirstStaff() Staff {

	var staff Staff

	//查询数据
	database.DB.First(&staff)

	return staff
}

func SelectStaff() []Staff {

	var staff []Staff

	//查询数据
	database.DB.Find(&staff)

	//条件查询
	// database.DB.Where("name = ?", "活动").First(&staff)

	return staff
}
