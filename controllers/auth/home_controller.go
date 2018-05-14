package auth

import (
	"html/template"
	"bee-go-myBlog/requests"
	"bee-go-myBlog/common"
	"bee-go-myBlog/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"bee-go-myBlog/services"
)

type HomeController struct {
	controllers.BaseController
}

//@router /console/login [get]
func (h *HomeController) Login()  {
	beego.ReadFromRequest(&h.Controller)
	h.Data["xsrfdata"]= template.HTML(h.XSRFFormHTML())

	h.TplName = "auth/login.tpl"
}

//@router /console/login [post]
func (h *HomeController) PostLogin() {
	u := common.LoginRequest{}
	if err := h.ParseForm(&u); err != nil {
		h.MyReminder("error","校验内部出了错误")
	}

	code ,message := requests.IphptValidate(h.Ctx,"Login")
	fmt.Println(message)
	if code != 0 {
		h.MyReminder("error",message)
		h.Redirect("/console/register",302)
		return
	}
	user,res := services.UserPwdCheck(u)
	if !res {
		//出错了
		h.MyReminder("error","用户不存在或者账号密码错误")
	}
	h.SetSession("Id",user.Id)
	h.Redirect("/console/home",302)
}


//@router /console/register [get]
func (h *HomeController) Register() {
	beego.ReadFromRequest(&h.Controller)
	h.Data["xsrfdata"]= template.HTML(h.XSRFFormHTML())
	h.TplName = "auth/register.tpl"
}


//@router /console/register [post]
func (h *HomeController) PostRegister() {
	u := common.RegisterRequest{}
	if err := h.ParseForm(&u); err != nil {
		h.MyReminder("error","校验内部出了错误")
	}

	code ,message := requests.IphptValidate(h.Ctx,"Register")
	fmt.Println(message)
	if code != 0 {
		h.MyReminder("error",message)
		h.Redirect("/console/register",302)
		return
	}
	_,err := services.UserStore(u)
	if err != nil {
		h.MyReminder("error","注册失败")
	}
	h.Redirect("/console/login",302)
}