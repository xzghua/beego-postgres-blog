package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.m111e"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) MyTest() {
	c.Data["Website"] = "ylsc633"
	c.Data["Email"] = "ylsc633@gamil.com"
	c.TplName = "index.tpl"
}