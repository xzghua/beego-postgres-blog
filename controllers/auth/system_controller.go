package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/models"
)

type SystemController struct {
	controllers.BaseController
}


func (s *SystemController) Index() {
	system,_ := models.AllLink()
	s.Data["system"] = system
	s.Layout = "auth/master.tpl"
	s.TplName = "auth/system/index.tpl"
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
