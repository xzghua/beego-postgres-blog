package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/models"
	"html/template"
	"github.com/astaxie/beego"
	"bee-go-myBlog/common"
	"bee-go-myBlog/requests"
)

type SystemController struct {
	controllers.BaseController
}


func (s *SystemController)URLMapping()  {
	s.Mapping("Index",s.Index)
	s.Mapping("Update",s.Update)
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
	u := common.SystemUpdate{}
	if err := s.ParseForm(&u); err != nil {
		s.MyReminder("error","")
	}
	code ,message := requests.IphptValidate(s.Ctx,"System")
	if code != 0 {
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

