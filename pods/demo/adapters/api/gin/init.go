package gin

import (
	_ "fmt"
	_ "strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/asaskevich/EventBus"
	"github.com/goodmall/goodmall/app"
	"github.com/goodmall/goodmall/pods/demo/infra/repo/tiedot"
	"github.com/goodmall/goodmall/pods/demo/usecase"
)

// InitPod 集成入口  系统应用（SysApp）可用通过此方法把该模块的功能集成到系统总体版图去

func InitPod(engine *gin.Engine, env app.Env) {

	r := engine

	tr := tiedot.NewTodoRepo()
	ti := usecase.NewTodoInteractor(tr, env.EventBus)

	th := TodoHandler{ts: ti}
	// var _ TodoHandler = th

	r.GET("/todo", th.GetTodo)
	r.POST("/todo", th.CreateTodo)
	r.POST("/todos", th.Todos)
	/*
		r.GET("/todo", func(c *gin.Context) {
			id := c.Query("id")

			id2, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(200, gin.H{
					"error": "not a number",
				})
				return
			}

			ii := usecase.NewTodoInteractor(tr, env.EventBus)
			response, _ := ii.Todo(id2)

			c.JSON(200, gin.H{
				"message": response,
			})

		}) */

	// TODO  我们可以在初始化方法中 触发一些事件 供内部钩子注册

	// evbus.
}
