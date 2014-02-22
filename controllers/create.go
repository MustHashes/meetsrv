package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/MustHashes/meetsrv/models"
)

type CreateController struct {
	beego.Controller
}

type CreateJSON struct {
	User string // phone number in format +[countrycode][number]\
	Name string // name of the event
}

func (this *CreateController) Post() {
	var req CreateJSON
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Abort("400")
	}

	models.FindUser(req.User)
	event := models.CreateEvent(req.Name, req.User)

	this.Ctx.WriteString(event.Id.Hex())
}
