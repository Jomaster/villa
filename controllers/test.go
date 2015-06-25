package controllers

import (
	"github.com/astaxie/beego"
)

type Test struct {
	beego.Controller
}

func (c *Test) Get() {
	c.Data["Title"] = "Какашко!"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}
 
