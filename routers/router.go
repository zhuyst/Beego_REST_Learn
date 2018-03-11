package routers

import (
	"beego_rest_learn/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.MainController{})
    //beego.Router("/list/", &controllers.MainController{},"get:List")
    //beego.Router("/:id",&controllers.MainController{},"get:GetOne")
	//
    //beego.Router("/",&controllers.MainController{},"post:Insert")
    //beego.Router("/:id",&controllers.MainController{},"put:Update")
    //beego.Router("/:id",&controllers.MainController{},"delete:Delete")
}
