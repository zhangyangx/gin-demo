package main

import (
	"fmt"
	"gin-demo/router/app/blog"
	shop "gin-demo/router/app/shop"
	"gin-demo/router/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(shop.Routers, blog.Routers)

	r := gin.New()
	// 注册一个全局令牌桶限流中间件
	//r.Use(middleware.RateLimitMiddleware(3*time.Second, 5))

	// 初始化路由
	r = routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
