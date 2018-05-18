package auth

import (
	"beego-postgres-blog/controllers"
	"beego-postgres-blog/models"
	"github.com/astaxie/beego"
	"strconv"
	"html/template"
	"beego-postgres-blog/common"
	"beego-postgres-blog/requests"
)

type LinkController struct {
	controllers.BaseController
}

func (l *LinkController)URLMapping()  {
	l.Mapping("Index",l.Index)
	l.Mapping("Create",l.Create)
	l.Mapping("Store",l.Store)
	l.Mapping("Show",l.Show)
	l.Mapping("Update",l.Update)
	l.Mapping("Destroy",l.Destroy)
}

// @router /link [get]
func (l *LinkController) Index() {
	beego.ReadFromRequest(&l.Controller)
	page := l.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	link,_ := models.AllLink(page2)
	totalPage,lastPage,currentPage,nextPage := models.LinkPaginate(page2,"console")
	l.Data["xsrfdata"]= template.HTML(l.XSRFFormHTML())
	l.Data["link"] = link
	l.Data["totalPage"] = totalPage
	l.Data["lastPage"] = lastPage
	l.Data["currentPage"] = currentPage
	l.Data["nextPage"] = nextPage
	l.Layout = "auth/master.tpl"
	l.TplName = "auth/post/index.tpl"

	l.Layout = "auth/master.tpl"
	l.TplName = "auth/link/index.tpl"
}

// @router /link/create [get]
func (l *LinkController) Create() {
	beego.ReadFromRequest(&l.Controller)
	l.Data["xsrfdata"]=template.HTML(l.XSRFFormHTML())
	l.Layout = "auth/master.tpl"
	l.TplName = "auth/link/create.tpl"
}

// @router /link [post]
func (l *LinkController) Store() {
	u := common.LinkCreate{}
	if err := l.ParseForm(&u); err != nil {
		l.MyReminder("error","")
	}
	code ,message := requests.IphptValidate(l.Ctx,"Link")
	if code != 0 {
		l.MyReminder("error",message)
		l.Redirect("/console/link/create",302)
		return
	}
	link := models.Links{
		Name	:	u.Name,
		Link	:	u.Link,
		Sort	:	u.Sort,
	}
	models.AddLinks(&link)
	l.MyReminder("success","操作成功")
	l.Redirect("/console/link",302)
	return
}

// @router /link/:id([0-9]+/edit [get]
func (l *LinkController) Show() {
	beego.ReadFromRequest(&l.Controller)
	id := l.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	link,_ := models.GetLinksById(id64)
	l.Data["link"] = link
	l.Data["xsrfdata"]=template.HTML(l.XSRFFormHTML())
	l.Layout = "auth/master.tpl"
	l.TplName = "auth/link/edit.tpl"
}

// @router /link/:id([0-9]+ [put]
func (l *LinkController) Update() {
	u := common.LinkCreate{}
	if err := l.ParseForm(&u); err != nil {
		l.MyReminder("error","")
	}
	code ,message := requests.IphptValidate(l.Ctx,"Link")
	if code != 0 {
		l.MyReminder("error",message)
		l.Redirect("/console/link/create",302)
		return
	}
	id := l.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	link := models.Links{
		Id		:	id64,
		Name	:	u.Name,
		Sort	:	u.Sort,
		Link	:	u.Link,
	}
	models.UpdateLinksById(&link)
	l.MyReminder("success","操作成功")
	l.Redirect("/console/link",302)
	return
}

// @router /link/:id([0-9]+ [delete]
func (l *LinkController) Destroy() {
	id := l.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	models.DeleteLinks(id64)
	l.MyReminder("success","操作成功")
	l.Redirect("/console/link",302)
	return
}
