package controllers

import "labix.org/v2/mgo/bson"

type MainController struct {
	BaseController
}

type Description struct {
	Type        string
	Description string
}

func (this *MainController) Get() {

	this.Data["json"] = bson.M{
		"Version": "0.0.1",
		"Endpoints": bson.M{
			"/":           Description{Type: "GET", Description: "Get version, endpoints, etc."},
			"/create":     Description{Type: "POST", Description: `Takes JSON for input (e.g. "{ "User" : "+100000000", "Name": "TestEvent" }") and returns the created event ID.`},
			"/user":       Description{Type: "POST", Description: `Takes JSON for input (e.g. "{ "User": "+100000000" }") and returns user statistics.`},
			"/user/name":  Description{Type: "POST", Description: `Modify the user name. Takes JSON for input (parameters "User" and "Name") and returns "OK" if good.`},
			"/list":       Description{Type: "GET", Description: `Get the list of all events. Returned as an array of events.`},
			"/:id":        Description{Type: "GET", Description: `Get a specific event. ':id' is the event ID.`},
			"/:id/join":   Description{Type: "POST", Description: `Mark a user at the event. Takes JSON for input (parameters "User") and returns "OK" if ok`},
			"/:id/socket": Description{Type: "GET", Description: `Connect to the JSON websocket, which monitors users' scores in the event.`},
		},
	}

	this.ServeJson()
}
