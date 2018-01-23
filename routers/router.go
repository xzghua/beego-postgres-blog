package routers

import (
	"bee-go-myBlog/controllers"
	"github.com/astaxie/beego"
	"bee-go-myBlog/controllers/auth"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/my",&controllers.MainController{},"get:MyTest")
    beego.Router("/console/home",&auth.PostController{},"get:Index")
}
