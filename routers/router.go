package routers

import (
	"bee-go-myBlog/controllers"
	"github.com/astaxie/beego"
	"bee-go-myBlog/controllers/auth"
	"github.com/astaxie/beego/context"

)

func init() {
	filter()

	beego.Router("/", &controllers.MainController{})
    beego.Router("/my",&controllers.MainController{},"get:MyTest")
    beego.Router("/console/home",&auth.HomeController{},"get:Index")

    beego.Router("/console/post",&auth.PostController{},"get:Index")
    beego.Router("/console/post/create",&auth.PostController{},"get:Create")
    beego.Router("/console/post",&auth.PostController{},"post:Store")
    beego.Router("/console/post/:id([0-9]+/edit",&auth.PostController{},"get:Edit")
    beego.Router("/console/post/:id([0-9]+",&auth.PostController{},"put:Update")
    beego.Router("/console/post",&auth.PostController{},"delete:Destroy")

	beego.Router("/console/cate",&auth.CateController{},"get:Index")
	beego.Router("/console/cate/create",&auth.CateController{},"get:Create")
	beego.Router("/console/cate",&auth.CateController{},"post:Store")
	beego.Router("/console/cate",&auth.CateController{},"put:Update")
	beego.Router("/console/cate",&auth.CateController{},"delete:Destroy")


	beego.Router("/console/tag",&auth.TagController{},"get:Index")
	beego.Router("/console/tag/create",&auth.TagController{},"get:Create")
	beego.Router("/console/tag",&auth.TagController{},"post:Store")
	beego.Router("/console/tag",&auth.TagController{},"put:Update")
	beego.Router("/console/tag",&auth.TagController{},"delete:Destroy")

	beego.Router("/console/link",&auth.LinkController{},"get:Index")
	beego.Router("/console/link/create",&auth.LinkController{},"get:Create")
	beego.Router("/console/link",&auth.LinkController{},"post:Store")
	beego.Router("/console/link",&auth.LinkController{},"put:Update")
	beego.Router("/console/link",&auth.LinkController{},"delete:Destroy")

	beego.Router("/console/system",&auth.SystemController{},"get:Index")
	beego.Router("/console/system/create",&auth.SystemController{},"get:Create")
	beego.Router("/console/system",&auth.SystemController{},"post:Store")
	beego.Router("/console/system",&auth.SystemController{},"put:Update")
	beego.Router("/console/system",&auth.SystemController{},"delete:Destroy")



    beego.Router("/console/tag/auto",&auth.TagController{},"get:GetTagByLike")

}

func filter() {
	// 拓展http方法
	extMethod()
	// 权限校验
}

func extMethod() {
	var filter = func(ctx *context.Context) {
		method := ctx.Input.Query("_method")
		println(method)
		if method != "" && ctx.Input.IsPost() {
			ctx.Request.Method = method
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, filter)
}