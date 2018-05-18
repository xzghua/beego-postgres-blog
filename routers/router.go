package routers

import (
	"github.com/astaxie/beego"
	"beego-postgres-blog/controllers/auth"
	"github.com/astaxie/beego/context"
	"beego-postgres-blog/controllers/index"
	"beego-postgres-blog/models"
)

func init() {
	filter()

	ns := beego.NewNamespace("/console",
		beego.NSInclude(
			&auth.PostController{},
			&auth.CateController{},
			&auth.LinkController{},
			&auth.TagController{},
			&auth.SystemController{},
		),
	)

	beego.AddNamespace(ns)

	beego.Include(&index.HomeController{})
	beego.Include(&auth.HomeController{})
}

func filter() {
	// 拓展http方法
	extMethod()
	// 权限校验
	checkAuth()

	forRegister()
}

func checkAuth() {
	var filter = func(ctx *context.Context) {
		sess := beego.GlobalSessions.SessionRegenerateID(ctx.ResponseWriter,ctx.Request)
		defer sess.SessionRelease(ctx.ResponseWriter)
		v := sess.Get("Id")
		if v == nil {
			ctx.Redirect(302,"/auth/login")
		}
	}
	beego.InsertFilter("/console/*",beego.BeforeRouter,filter)
}

func extMethod() {
	var filter = func(ctx *context.Context) {
		method := ctx.Input.Query("_method")
		//println(method)
		if method != "" && ctx.Input.IsPost() {
			ctx.Request.Method = method
		}
	}
	beego.InsertFilter("*", beego.BeforeRouter, filter)
}

func forRegister() {
	var filter = func(ctx *context.Context) {
		cnt,err := models.GetUserCount()
		if err != nil || cnt == -1 {
			ctx.Redirect(302,"/auth/login")
		}
		if cnt > 2 {
			ctx.Redirect(302,"/auth/login")
		}
	}
	beego.InsertFilter("/auth/register", beego.BeforeRouter, filter)

}