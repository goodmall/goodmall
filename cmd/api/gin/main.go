package main

import(
	"github.com/gin-gonic/gin"

	"github.com/goodmall/goodmall/pods/user/usecase"
)


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/userhelp", func(c *gin.Context) {

		// userInteractor := usecase.NewUserInteractor()
		userInteractor :=  usecase.NewUsecase(/* 依赖暂缺 */).NewUserInteractor() // usecase.NewUserInteractor()
		response := userInteractor.Help() 

		c.JSON(200, gin.H{
			"message": response,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}