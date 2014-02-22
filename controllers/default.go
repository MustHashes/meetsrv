package controllers

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	this.Data["json"] = bson.M{"Version": "0.0.1"}

	this.ServeJson()
}
