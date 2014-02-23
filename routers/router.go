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

	// user name controller
	beego.Router("/user/name", &controllers.UserController{}, "post:Name")

	// list controller
	beego.Router("/list", &controllers.MeetController{}, "get:List")

	// user socket controller
	beego.Router("/:id/socket", &controllers.MeetController{}, "get:Socket")

	// join (marking self here) controller
	beego.Router("/:id/join", &controllers.MeetController{}, "post:Join")

	// vote someone as here controller
	beego.Router("/:id/seen", &controllers.MeetController{}, "post:Seen")

	// // trending "newsfeed" controller
	// beego.Router("/:id/trending", &controllers.MeetController{}, "get:Trending")

	// // config controller
	// beego.Router("/:id/config", &controllers.MeetController{}, "get:GetConfig;post:PostConfig")

	// get controller
	beego.Router("/:id", &controllers.MeetController{})

}
