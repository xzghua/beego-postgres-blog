package controllers

import (
	"github.com/astaxie/beego"
	"bee-go-myBlog/models"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

type createUser struct {
	name string
	email string
	password string
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.m111e"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) MyTest() {

	var user users.Users
	user.Email = "211@11.com"
	user.Name = "叶落山城秋"
	user.Password = "mima1234"

	id,err := users.AddUsers(&user)

	if err == nil {
		logs.Info("新增的ID是",id)
		fmt.Print(id)
	}

	logs.Info("2新增的ID是",id)

	//c.Data["Website"] = "ylsc633"
	//c.Data["Email"] = "ylsc633@gamil.com"
	//c.TplName = "index.tpl"
}