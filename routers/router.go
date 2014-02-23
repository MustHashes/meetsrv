package routers

import (
	"github.com/MustHashes/meetsrv/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// returns version
	beego.Router("/", &controllers.MainController{})

	// create controller
	beego.Router("/create", &controllers.CreateController{})

	// user controller
	beego.Router("/user", &controllers.UserController{})

	// list controller
	beego.Router("/list", &controllers.MeetController{}, "get:List")

	// get controller
	beego.Router("/:id", &controllers.MeetController{})

	// join controller
	// beego.Router("/:id/join", &controllers.MeetController{}, "post:Join")

	// // memberlist controller
	// beego.Router("/:id/members", &controllers.MeetController{}, "get:Members")

	// // trending "newsfeed" controller
	// beego.Router("/:id/trending", &controllers.MeetController{}, "get:Trending")

	// // config controller
	// beego.Router("/:id/config", &controllers.MeetController{}, "get:GetConfig;post:PostConfig")
}
