package gin

import (
	"github.com/gin-gonic/gin"

	"github.com/goodmall/goodmall/pods/user/usecase"
)

func InitPod(engine *gin.Engine) {

	r := engine

	r.GET("/userhelp", func(c *gin.Context) {

		// userInteractor := usecase.NewUserInteractor()
		userInteractor := usecase.NewUsecase( /* 依赖暂缺 */ ).NewUserInteractor() // usecase.NewUserInteractor()
		response := userInteractor.Help()

		c.JSON(200, gin.H{
			"message": response,
		})

	})

}
