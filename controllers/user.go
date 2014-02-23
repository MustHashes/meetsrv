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
	User     *models.User
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
	response.User = user

	response.Leader = models.FindEventsByLeader(req.User)
	response.Attendee = models.FindEventsByAttendee(req.User)

	this.Data["json"] = &response
	this.ServeJson()
}

type NameRequestJSON struct {
	User string // phone number in format +[countrycode][number]
	Name string
}

func (this *UserController) Name() {
	var req NameRequestJSON
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Abort("400")
	}

	user := models.FindUser(req.User)
	user.Name = req.Name
	user.Update()

	this.Ctx.WriteString("OK")
}
