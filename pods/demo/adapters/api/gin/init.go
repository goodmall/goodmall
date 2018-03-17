package gin

import (
	_ "fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/asaskevich/EventBus"
	"github.com/goodmall/goodmall/app"
	"github.com/goodmall/goodmall/pods/demo/usecase"
)

// InitPod 集成入口  系统应用（SysApp）可用通过此方法把该模块的功能集成到系统总体版图去

func InitPod(engine *gin.Engine, env app.Env) {

	r := engine

	r.GET("/todo", func(c *gin.Context) {
		id := c.Query("id")

		id2, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(200, gin.H{
				"error": "not a number",
			})
			return
		}

		ii := usecase.NewTodoInteractor()
		response, _ := ii.GetTodo(id2)

		c.JSON(200, gin.H{
			"message": response,
		})

	})

	// TODO  我们可以在初始化方法中 触发一些事件 供内部钩子注册

	// evbus.
}
