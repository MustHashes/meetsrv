package controllers

import "github.com/MustHashes/meetsrv/models"

type MeetController struct {
	BaseController
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

func (this *MeetController) List() {
	this.Data["json"] = models.FindAllEvents()
	this.ServeJson()
}
