package main

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/gin-gonic/gin"

	"github.com/asaskevich/EventBus"
	"github.com/go-ozzo/ozzo-dbx"

	// "github.com/goodmall/goodmall/cmd/api/gin/app"
	"github.com/goodmall/goodmall/app"

	userpod "github.com/goodmall/goodmall/pods/user/adapters/api/gin"

	"github.com/goodmall/goodmall/pods/demo"
	demoPodHome "github.com/goodmall/goodmall/pods/demo/adapters/api/gin"
)

func main() {

	//<========================================================================|
	//                      加载系统配置文件  并实例化系统级组件

	// load application configurations
	if err := app.LoadConfig("./../../../config" /*, "./config"*/); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}
	// fmt.Println(app.Config)

	//     ##  实例化系统级组件

	// create the logger
	// logger := logrus.New()

	// connect to the database
	// db, err := dbx.MustOpen("mysql", app.Config.DSN)
	db, err := dbx.Open("mysql", "root:@/aheadmall")
	if err != nil {
		panic(err)
	}

	// create event bus
	bus := EventBus.New()
	// FIXME 注意到在app/env.go 中导入的依赖文件这里还要导入一遍 有点累赘 可以考虑吧系统组件的实例化功能 提取到env去这样避免冗余

	env := app.Env{
		Db:       db,
		EventBus: bus,
	}

	//========================================================================>|

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//========================================================================<|
	// 	##		集成各个Pod 并传递配置信息和环境变量到各个层去

	// 需要先配置必要信息 注入必要依赖: userpod.Configure(config).Init()?
	userpod.InitPod(r)

	demoPod := demo.NewDemoPod()
	// 将系统配置拷贝到模块配置中 模块配置只需要声明需要的字段 类似于模式匹配
	copier.Copy(&demoPod.Config, &app.Config)
	fmt.Print(demoPod)
	demoPodHome.InitPod(r, env)

	//========================================================================>|
	r.Run() // listen and serve on 0.0.0.0:8080
}
