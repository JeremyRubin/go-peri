package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me1"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
