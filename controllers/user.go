package controllers

import (
	"encoding/json"

	"github.com/MustHashes/meetsrv/models"
)

type UserController struct {
	BaseController
}

type UserRequestJSON struct {
	User string // phone number in format +[countrycode][number]
}

type UserResponseJSON struct {
	User     string
	Name     string
	Leader   []string
	Attendee []string
}

func (this *UserController) Post() {
	var req UserRequestJSON
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Abort("400")
	}

	response := UserResponseJSON{}

	user := models.FindUser(req.User)
	response.User = user.Number
	response.Name = "tobeimplemented"

	response.Leader = models.FindEventsByLeader(req.User)
	response.Attendee = models.FindEventsByAttendee(req.User)

	this.Data["json"] = &response
	this.ServeJson()
}
