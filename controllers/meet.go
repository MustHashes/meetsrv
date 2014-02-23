package controllers

import (
	"github.com/MustHashes/meetsrv/models"
	"github.com/astaxie/beego"
)

type MeetController struct {
	beego.Controller
}

func (this *MeetController) Get() {
	id := this.Ctx.Input.Param(":id")
	event := models.FindEvent(id)
	if event == nil {
		this.Abort("404")
	}

	this.Data["json"] = event
	this.ServeJson()
}
