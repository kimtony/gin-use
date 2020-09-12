package controller

import (
	"gin-demo/repos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello hanlde api status
func Activity(c *gin.Context) {
	data := repos.SelectActvity()
	c.JSON(http.StatusOK, ReturnSuccess(data))
}
