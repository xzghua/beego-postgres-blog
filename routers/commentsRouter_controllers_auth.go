package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/console/cate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/console/cate/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:CateController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/console/post`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:LinkController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/console/link`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

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

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:SystemController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:SystemController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/console/system`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/console/tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/console/tag`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "GetTagByLike",
			Router: `/console/tag/auto`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"] = append(beego.GlobalControllerRouter["bee-go-myBlog/controllers/auth:TagController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/console/tag/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
