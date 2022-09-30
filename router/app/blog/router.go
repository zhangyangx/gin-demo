package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(e *gin.Engine) {
	e.GET("/post", postHandler)
	e.GET("/comment", commentHandler)
}

func postHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello postHandler!",
	})
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello commentHandler!",
	})
}
