package auth

import (
	"github.com/astaxie/beego"
)

type PostController struct {
	beego.Controller
}

func (c *PostController) Index()  {
	c.TplName = "auth/index.tpl"
}