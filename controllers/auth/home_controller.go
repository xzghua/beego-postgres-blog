package auth

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

//@router /console/login
func (h *HomeController) Login()  {

	h.TplName = "auth/login.tpl"
}




//@router /console/register
func (h *HomeController) Register() {
	h.TplName = "auth/register.tpl"
}