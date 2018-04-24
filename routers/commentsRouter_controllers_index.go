package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
