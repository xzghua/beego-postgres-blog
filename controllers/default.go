package controllers

import (
	"github.com/astaxie/beego"
	"beego-postgres-blog/models"
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


func (cc *MainController) Get() {
	flash := beego.NewFlash()
	flash.Error("jkasdjfklasdf..")
	flash.Store(&cc.Controller)
	//cc.Redirect("/",302)
	//cc.Data["Website"] = "beego.m111e"
	//cc.Data["Email"] = "astaxie@gmail.com"
	//cc.TplName = "index.tpl"
}

func (cc *MainController) MyTest() {

	var user models.Users
	user.Email = "211@11.com"
	user.Name = "叶落山城秋"
	user.Password = "mima1234"

	id,err := models.AddUsers(&user)

	if err == nil {
		logs.Info("新增的ID是",id)
		fmt.Print(id)
	}

	logs.Info("2新增的ID是",id)
	var updateUser models.Users
	updateUser.Id = 1
	updateUser.Email = "206@g9zz.com"
	updateUser.Name = "叶落"
	updateUser.Password = "222222222"
	err = models.UpdateUsersById(&updateUser)

	if err == nil {
		logs.Info("正确")
	} else {
		logs.Info("错误")
	}

	//c.Data["Website"] = "ylsc633"
	//c.Data["Email"] = "ylsc633@gamil.com"
	//c.TplName = "index.tpl"
}