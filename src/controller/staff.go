package controller

import (
	"fmt"
	"gin-use/src/util/cache"
	"gin-use/src/util"
	"gin-use/src/repos"
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

	token, err := util.GenToken(inpayload)
	println("------token-------", token)
	if err != nil {
		println("生成token报错")
	}

	verfiy, err := util.VerifyAction(token)

	println("--------verfiy-----", verfiy)

	if err != nil {
		println("解析token报错")
	}
}

//redis
func RedisTest(c *gin.Context) {

	//redis init
	cache.RedisInit()

	const id = 218910024463480831
	redisKey := fmt.Sprintf("xicheng:account:%d", id)

	//set
	err := cache.RedisClient.Set(redisKey, 1008611, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	//get
	val, err := cache.RedisClient.Get(redisKey).Result()
	if err != nil {
		fmt.Printf("get redisKey failed, err:%v\n", err)
		return
	}

	fmt.Println("redisKey", val)
}
