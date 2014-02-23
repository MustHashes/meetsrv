package controllers

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
)

type MainController struct {
	beego.Controller
}

type Description struct {
	Type        string
	Description string
}

func (this *MainController) Get() {

	this.Data["json"] = bson.M{
		"Version": "0.0.1",
		"Endpoints": bson.M{
			"/":       Description{Type: "GET", Description: "Get version, endpoints, etc."},
			"/create": Description{Type: "POST", Description: `Takes JSON for input (e.g. "{ "User" : "+100000000", "Name": "TestEvent" }") and returns the created event ID.`},
			"/user":   Description{Type: "POST", Description: `Takes JSON for input (e.g. "{ "User": "+100000000" }") and returns user statistics.`},
			"/:id":    Description{Type: "GET", Description: `Get a specific event. ':id' is the event ID.`},
		},
	}

	this.ServeJson()
}
