package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"] = append(beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"],
		beego.ControllerComments{
			Method: "Insert",
			Router: `/user/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"] = append(beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"] = append(beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"] = append(beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"] = append(beego.GlobalControllerRouter["beego_rest_learn/controllers:MainController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/user/list/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
