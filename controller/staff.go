package controller

import (
	"fmt"
	"gin-demo/database"
	"gin-demo/repos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello hanlde api status
func Staff(c *gin.Context) {
	data := repos.SelectStaff()
	c.JSON(http.StatusOK, ReturnSuccess(data))
}

func RedisTest(c *gin.Context) {

	//redis init
	database.RedisInit()

	const id = 218910024463480831
	redisKey := fmt.Sprintf("xicheng:account:%d", id)

	//set
	err := database.RedisClient.Set(redisKey, 1008611, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	//get
	val, err := database.RedisClient.Get(redisKey).Result()
	if err != nil {
		fmt.Printf("get redisKey failed, err:%v\n", err)
		return
	}
	
	fmt.Println("redisKey", val)
}
