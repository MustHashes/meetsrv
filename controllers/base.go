package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	this.Ctx.ResponseWriter.Header().Set("Pragma", "no-cache")
	this.Ctx.ResponseWriter.Header().Set("Expires", "0")
}
