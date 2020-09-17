package controller

import (
	"gin-demo/repos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello hanlde api status
func Staff(c *gin.Context) {
	data := repos.SelectStaff()
	c.JSON(http.StatusOK, ReturnSuccess(data))
}
