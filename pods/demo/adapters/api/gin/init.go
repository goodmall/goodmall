package gin

import (
	_ "fmt"
	_ "strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/asaskevich/EventBus"
	"github.com/goodmall/goodmall/app"
	//"github.com/goodmall/goodmall/pods/demo/infra/repo/bolthold"
	// "github.com/goodmall/goodmall/pods/demo/infra/repo/ql"
	"github.com/goodmall/goodmall/pods/demo/infra/repo/mysql"
	//	"github.com/goodmall/goodmall/pods/demo/infra/repo/tiedot"
	"github.com/goodmall/goodmall/pods/demo/usecase"

	"github.com/jinzhu/gorm"
)

// InitPod 集成入口  系统应用（SysApp）可用通过此方法把该模块的功能集成到系统总体版图去

func InitPod(engine *gin.Engine, env app.Env) {

	r := engine

	//tr := ql.NewTodoRepo() // bolthold.NewTodoRepo() // tiedot.NewTodoRepo()
	db, err := gorm.Open("mysql", app.Config.DSNMysql)
	if err != nil {
		panic("failed to connect database" + app.Config.DSNMysql)
	}
	db.LogMode(true) // 开启日志 生产环境中可以关闭
	tr := mysql.NewTodoRepo(db)
	ti := usecase.NewTodoInteractor(tr, env.EventBus)

	th := TodoHandler{ts: ti}
	// var _ TodoHandler = th

	r.GET("/todo", th.Get)
	r.GET("/todo/count", th.Count)
	r.POST("/todo", th.Create)
	r.PUT("/todo", th.Update)
	r.GET("/todos", th.Query)
	r.DELETE("/todo", th.Delete)
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
