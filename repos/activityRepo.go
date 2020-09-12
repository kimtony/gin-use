package repos

import (
	"gin-demo/database"
	"fmt"
)

// Activity struct
type Activity struct {
	ID   string `json: "id"`
	Name string `json: "name"`
}

func SelectActvity() Activity {
	activity := Activity{}
	//查询数据
	rows, err := database.DB.Query("SELECT id,name FROM activity;")
	if err != nil {
		println("---error--------", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&activity.ID, &activity.Name)
		if err != nil {
			println("---error-----22---", err.Error())
		}
		fmt.Println("id = ", activity.ID, "\nname = ", activity.Name, "-----------")

	}
	return activity
}
