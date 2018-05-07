package routers

import (
	"github.com/astaxie/beego"
	"bee-go-myBlog/controllers/auth"
	"github.com/astaxie/beego/context"

	"bee-go-myBlog/controllers/index"
)

func init() {
	filter()


	beego.Include(&auth.PostController{})
	beego.Include(&auth.CateController{})
	beego.Include(&auth.LinkController{})
	beego.Include(&auth.TagController{})
	beego.Include(&auth.SystemController{})
	beego.Include(&index.HomeController{})
	beego.Include(&auth.HomeController{})
}

func filter() {
	// 拓展http方法
	extMethod()
	// 权限校验
}

func extMethod() {
	var filter = func(ctx *context.Context) {
		method := ctx.Input.Query("_method")
		println(method,ctx.Input.IsPost())
		if method != "" && ctx.Input.IsPost() {
			ctx.Request.Method = method
		}
	}
	beego.InsertFilter("*", beego.BeforeRouter, filter)
}