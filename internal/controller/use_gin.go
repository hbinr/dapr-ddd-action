package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u UserController) SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}
