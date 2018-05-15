package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/cate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/cate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/cate/:id([0-9]+`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/cate/:id([0-9]+`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `/cate/:id([0-9]+/edit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/cate/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/auth/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"],
		beego.ControllerComments{
			Method: "PostLogin",
			Router: `/auth/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/auth/register`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"],
		beego.ControllerComments{
			Method: "PostRegister",
			Router: `/auth/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"],
		beego.ControllerComments{
			Method: "ClearCache",
			Router: `/console/clearCache`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"],
		beego.ControllerComments{
			Method: "Home",
			Router: `/console/home`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:HomeController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/console/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/link`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/link`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/link/:id([0-9]+`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/link/:id([0-9]+`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `/link/:id([0-9]+/edit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/link/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/post`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/post`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/post/:id([0-9]+`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/post/:id([0-9]+`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/post/:id([0-9]+/edit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:PostController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/post/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:SystemController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:SystemController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/system`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:SystemController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:SystemController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/system/1`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/tag`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/tag/:id([0-9]+`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/tag/:id([0-9]+`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `/tag/:id([0-9]+/edit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "GetTagByLike",
			Router: `/tag/auto`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/tag/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
