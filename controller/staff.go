package controller

import (
	"fmt"
	"gin-demo/database"
	"gin-demo/helpers"
	"gin-demo/repos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello hanlde api status
func Staff(c *gin.Context) {
	data := repos.SelectStaff()
	c.JSON(http.StatusOK, ReturnSuccess(data))
}

//token
func Token(c *gin.Context) {
	//生成token
	inpayload := map[string]interface{}{
		"account": "3345472145214",
		"name":    "lk",
		"age":     18,
	}

	token, err := helpers.GenToken(inpayload)
	println("------token-------", token)
	if err != nil {
		println("生成token报错")
	}

	verfiy, err := helpers.VerifyAction(token)

	println("--------verfiy-----", verfiy)

	if err != nil {
		println("解析token报错")
	}
}

//redis
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
