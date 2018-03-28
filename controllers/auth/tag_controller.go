package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/services"
	"strconv"
	"bee-go-myBlog/models"
)

type TagController struct {
	controllers.BaseController
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

}
func (t *TagController) Store() {

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