package routers

import (
	"github.com/mikitu/go-file-manager/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/list", &controllers.MainController{}, "get:ListFiles")
    beego.Router("/delete", &controllers.MainController{}, "post:Delete")
    beego.Router("/download", &controllers.MainController{}, "post:Download")
    beego.Router("/upload", &controllers.MainController{}, "post:Upload")
    beego.Router("/create", &controllers.MainController{}, "post:NewFolder")
    beego.Router("/rename", &controllers.MainController{}, "post:Rename")
    beego.Router("/view", &controllers.MainController{}, "get:View")
}
