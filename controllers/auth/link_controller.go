package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/models"
)

type LinkController struct {
	controllers.BaseController
}


func (l *LinkController) Index() {
	link,_ := models.AllLink()
	l.Data["link"] = link
	l.Layout = "auth/master.tpl"
	l.TplName = "auth/link/index.tpl"
}

func (l *LinkController) Create() {

}

func (l *LinkController) Store() {

}

func (l *LinkController) Show() {

}

func (l *LinkController) Update() {

}
func (l *LinkController) Destroy() {

}
