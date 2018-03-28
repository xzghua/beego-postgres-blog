package routers

import (
	"bee-go-myBlog/controllers"
	"github.com/astaxie/beego"
	"bee-go-myBlog/controllers/auth"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/my",&controllers.MainController{},"get:MyTest")
    beego.Router("/console/home",&auth.HomeController{},"get:Index")

    beego.Router("/console/post",&auth.PostController{},"get:Index")
    beego.Router("/console/post/create",&auth.PostController{},"get:Create")
    beego.Router("/console/post",&auth.PostController{},"post:Store")
    beego.Router("/console/post",&auth.PostController{},"put:Update")
    beego.Router("/console/post",&auth.PostController{},"delete:Destroy")

	beego.Router("/console/cate",&auth.CateController{},"get:Index")
	beego.Router("/console/cate/create",&auth.CateController{},"get:Create")
	beego.Router("/console/cate",&auth.CateController{},"post:Store")
	beego.Router("/console/cate",&auth.CateController{},"put:Update")
	beego.Router("/console/cate",&auth.CateController{},"delete:Destroy")


    beego.Router("/console/tag/auto",&auth.TagController{},"get:GetTagByLike")

}
