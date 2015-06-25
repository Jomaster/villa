package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.top-menu = "top-menu.tpl"
c.TplNames = "index.tpl"
//c.Layout = "index.tpl"
//c.LayoutSections = make(map[string]string)
//c.LayoutSections["Top-menu"] = "top-menu.tpl"
//c.LayoutSections["Footer"] = "footer.tpl"
}
