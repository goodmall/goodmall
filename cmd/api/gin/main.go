package main

import (
	// "./engine"  // 不能用相路径导入 不然出现很诡异的错误！
	"github.com/goodmall/goodmall/cmd/api/gin/engine"
)

func main() {

	engine.GetMainEngine().Run() // listen and serve on 0.0.0.0:8080
}
