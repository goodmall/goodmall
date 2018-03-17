package app

import (
	"github.com/asaskevich/EventBus"
	"github.com/go-ozzo/ozzo-dbx"
)

// Env 全局环境变量 维持各种全局性组件
// 和Context语义差不多 参考 http://www.alexedwards.net/blog/organising-database-access
// 注意和Config的区别
//     Config一般是从文件中加载 是静态配置
//     Env更像是动态环境变量 是运行期对象注册表 这点跟Request-scoped context 比较像 但后者更感觉是针对方法依赖参数的传递
// Request-scoped context: 这个用法 可以参考强哥的 https://github.com/qiangxue/golang-restful-starter-kit/blob/master/app/scope.go
type Env struct {
	// db models.Datastore
	Db *dbx.DB
	// 事件总线
	EventBus EventBus.Bus
}
