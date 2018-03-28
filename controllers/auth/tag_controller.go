package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/services"
	"strconv"
	"bee-go-myBlog/models"
	"github.com/astaxie/beego"
	"html/template"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type TagController struct {
	controllers.BaseController
}

type TagRequest struct {
	Name 	string `form:"name" valid:"Required;MaxSize(30)"`
}


func (t *TagController) Index() {
	page := t.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	tag,_ := services.GetAllMyTag(page2)
	if tag == "" {
		t.MyReminder("error","系统内部错误")
	}
	totalPage,lastPage,currentPage,nextPage := models.TagPaginate(page2)
	t.Data["totalPage"] = totalPage
	t.Data["lastPage"] = lastPage
	t.Data["currentPage"] = currentPage
	t.Data["nextPage"] = nextPage
	t.Data["tag"] = tag
	t.Layout = "auth/master.tpl"
	t.TplName = "auth/tag/index.tpl"
}

func (t *TagController) Create() {
	beego.ReadFromRequest(&t.Controller)
	t.Data["xsrfdata"]=template.HTML(t.XSRFFormHTML())
	t.Layout = "auth/master.tpl"
	t.TplName = "auth/tag/create.tpl"
}

func (t *TagController) Store() {
	u := TagRequest{}
	valid := validation.Validation{}
	if err := t.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	if err != nil {
	}
	if !b {
		_,message := t.RequestValidate(valid)
		t.MyReminder("error",message)
		t.Redirect("/console/tag/create",302)
		return
	}
	var tagCreate = models.Tags{
		Name	:	u.Name,
	}
	_,err = models.AddTags(&tagCreate)
	if err != nil {
		t.MyReminder("error","系统内部错误")
	}
	t.Redirect("/console/tag",302)
}
func (t *TagController) Show() {

}
func (t *TagController) Update() {

}
func (t *TagController) Destroy() {

}


func (t *TagController) GetTagByLike() {
	param := t.GetString("term")
	res := services.GetTagByLike(param)
	t.Data["json"] = res
	t.ServeJSON()
}