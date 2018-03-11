package main

import (
	"github.com/gin-gonic/gin"

	userpod "github.com/goodmall/goodmall/pods/user/adapters/api/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 需要先配置必要信息 注入必要依赖: userpod.Configure(config).Init()?
	userpod.InitPod(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
