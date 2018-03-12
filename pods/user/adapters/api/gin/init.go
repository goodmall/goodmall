package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/goodmall/goodmall/pods/user/usecase"
)

// InitPod 集成入口  系统应用（SysApp）可用通过此方法把该模块的功能集成到系统总体版图去

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

	r.GET("/username", func(c *gin.Context) {

		// userInteractor := usecase.NewUserInteractor()
		userInteractor := usecase.NewUsecase( /* 依赖暂缺 */ ).NewUserInteractor() // usecase.NewUserInteractor()

		un := c.Query("username")
		user, err := userInteractor.FindByUsername(un)
		if err != nil {

			c.JSON(200, gin.H{
				"message": "查不到呀!",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "" + un + " : " + fmt.Sprint(user),
		})

	})

	// TODO  我们可以在初始化方法中 触发一些事件 供内部钩子注册
}
