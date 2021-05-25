package repos

import (
	"gin-use/src/util/database"
)

// Activity struct
type Activity struct {
	ID   string `json: "id"`
	Name string `json: "name"`
}

func SelectActivity() []Activity {

	var activity []Activity

	//条件查询
	database.DB.Where("id = ?", "255969939325160448").First(&activity)

	return activity
}
