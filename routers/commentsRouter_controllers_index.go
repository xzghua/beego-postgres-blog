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

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"],
		beego.ControllerComments{
			Method: "Archive",
			Router: `/archive`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"],
		beego.ControllerComments{
			Method: "Cate",
			Router: `/categories/:cate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/detail/:id([0-9]+`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/index:HomeController"],
		beego.ControllerComments{
			Method: "Tag",
			Router: `/tags/:tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
