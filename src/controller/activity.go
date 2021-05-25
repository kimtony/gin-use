package controller

import (
	"fmt"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

// Hello hanlde api status
func Activity(c *gin.Context) {
	result := make(chan string)
	go func() {
		//模拟网络访问
		time.Sleep(8 * time.Second)
		result <- "服务端结果"
	}()
	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(5 * time.Second):
		fmt.Println("网络访问超时了")
	}
	// data := repos.SelectActivity()
	c.JSON(http.StatusOK, ReturnSuccess(123))
}
