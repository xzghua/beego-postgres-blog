package routers

import (
	"bee-go-myBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/my",&controllers.MainController{},"get:MyTest")
}
