package routers

import (
	"github.com/goodmall/goodmall/cmd/api/beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
