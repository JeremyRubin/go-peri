package controllers

import (
	"github.com/astaxie/beego"
)

type ChatController struct {
	beego.Controller
}

func (this *ChatController) Get() {
	this.TplNames = "test.tpl"
}
