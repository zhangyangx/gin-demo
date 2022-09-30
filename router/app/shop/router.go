package blog

import (
	"gin-demo/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Routers(e *gin.Engine) {
	// 为/goods注册一个令牌桶中间件
	e.GET("/goods", middleware.RateLimitMiddleware(1*time.Second, 3), goodsHandler)
	e.GET("/checkout", checkoutHandler)
}

func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello goodsHandler!",
	})
}

func checkoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello checkoutHandler!",
	})
}
