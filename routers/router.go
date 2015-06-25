package routers

import (
	"villa/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/test", &controllers.Test{})
}
