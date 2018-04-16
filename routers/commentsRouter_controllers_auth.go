package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/console/post`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/console/post`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/console/post/:id([0-9]+`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/console/post/:id([0-9]+`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/console/post/:id([0-9]+/edit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/console/post/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
