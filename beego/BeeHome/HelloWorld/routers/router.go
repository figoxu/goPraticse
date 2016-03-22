package routers

import (
	"github.com/figoxu/goPraticse/beego/BeeHome/HelloWorld/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
