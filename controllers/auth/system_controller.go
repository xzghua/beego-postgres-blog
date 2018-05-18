package auth

import (
	"beego-postgres-blog/controllers"
	"beego-postgres-blog/models"
	"html/template"
	"github.com/astaxie/beego"
	"beego-postgres-blog/common"
	"beego-postgres-blog/requests"
)

type SystemController struct {
	controllers.BaseController
}


func (s *SystemController)URLMapping()  {
	s.Mapping("Index",s.Index)
	s.Mapping("Update",s.Update)
}


// @router /system [get]
func (s *SystemController) Index() {
	beego.ReadFromRequest(&s.Controller)
	system,_ := models.GetSystemsById(1)
	if system == nil {
		var systemCreate = models.Systems{
			Title		:	"叶落山城秋",
			CommentType	:	1,
			CdnType		: 	1,
		}
		id,_ := models.AddSystems(&systemCreate)
		system,_ = models.GetSystemsById(id)
	}

	s.Data["xsrfdata"]= template.HTML(s.XSRFFormHTML())
	s.Data["system"] = system
	s.Data["comment"] = common.GetComment()
	s.Data["cdn"] = common.GetCdn()
	s.Layout = "auth/master.tpl"
	s.TplName = "auth/system/update.tpl"
}

// @router /system/1 [put]
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
		Id					:	1,
		Title				:	u.Title,
		STitle				:	u.STitle,
		Description			:	u.Description,
		SeoKey				:	u.SeoKey,
		SeoDes				:	u.SeoDes,
		RecordNumber		:	u.RecordNumber,
		CommentType			:	u.CommentType,
		GithubClientId		:	u.GithubClientId,
		GithubClientSecret	:	u.GithubClientSecret,
		CyAppId				:	u.CyAppId,
		CyAppKey			:	u.CyAppKey,
		CdnType				:	u.CdnType,
		CdnUrl				:	u.CdnUrl,
		GithubName			:	u.GithubName,
		GithubRepo			:	u.GithubRepo,
	}

	models.UpdateSystemsById(&system)
	s.MyReminder("success","操作成功")
	s.Redirect("/console/system",302)
	return
}

