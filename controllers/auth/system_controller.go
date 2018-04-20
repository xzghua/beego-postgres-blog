package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/models"
	"fmt"
	"html/template"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
)

type SystemController struct {
	controllers.BaseController
}

type SystemUpdate struct {
	Title			string 		`form:"title" valid:"Required;MaxSize(23)"`
	STitle			string 		`form:"s_title" `
	Description 	string 		`form:"description" `
	SeoKey			string 		`form:"seo_key" `
	SeoDes			string 		`form:"seo_des" `
	RecordNumber	string 		`form:"record_number" `
}

// @router /console/system [get]
func (s *SystemController) Index() {
	beego.ReadFromRequest(&s.Controller)
	system,_ := models.GetSystemsById(1)
	if system == nil {
		var systemCreate = models.Systems{
			Title	:	"叶落山城秋",
		}
		id,_ := models.AddSystems(&systemCreate)
		system,_ = models.GetSystemsById(id)
	}
	s.Data["xsrfdata"]= template.HTML(s.XSRFFormHTML())
	s.Data["system"] = system
	s.Layout = "auth/master.tpl"
	s.TplName = "auth/system/update.tpl"
}

// @router /console/system/1 [put]
func (s *SystemController) Update() {
	u := SystemUpdate{}
	valid := validation.Validation{}
	if err := s.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	if err != nil {
	}
	if !b {
		_,message :=s.RequestValidate(valid)
		s.MyReminder("error",message)
		s.Redirect("/console/system",302)
		return
	}

	system := models.Systems{
		Id				:	1,
		Title			:	u.Title,
		STitle			:	u.STitle,
		Description		:	u.Description,
		SeoKey			:	u.SeoKey,
		SeoDes			:	u.SeoDes,
		RecordNumber	:	u.RecordNumber,
	}

	models.UpdateSystemsById(&system)
	s.MyReminder("success","操作成功")
	s.Redirect("/console/system",302)
	return
}

