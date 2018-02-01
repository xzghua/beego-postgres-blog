package auth

import (
	"github.com/astaxie/beego"
	"bee-go-myBlog/services"
)

type CateController struct {
	beego.Controller
}


func (c *CateController) GetCateByLike() {
	param := c.GetString("term")
	res := services.GetCateByLike(param)
	c.Data["json"] = res
	c.ServeJSON()
}