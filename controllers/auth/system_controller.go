package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/models"
)

type SystemController struct {
	controllers.BaseController
}


func (s *SystemController) Index() {
	system,_ := models.GetSystemsById(1)
	if system == nil {
		var systemCreate = models.Systems{
			Title	:	"叶落山城秋",
		}
		id,_ := models.AddSystems(&systemCreate)
		system,_ = models.GetSystemsById(id)
	}
	s.Data["system"] = system
	s.Layout = "auth/master.tpl"
	s.TplName = "auth/system/update.tpl"
}

func (s *SystemController) Create() {

}

func (s *SystemController) Store() {

}

func (s *SystemController) Show() {

}

func (s *SystemController) Update() {

}
func (s *SystemController) Destroy() {

}
