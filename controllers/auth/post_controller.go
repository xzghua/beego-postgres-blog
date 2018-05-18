package auth

import (
	"strconv"
	"beego-postgres-blog/services"
	"html/template"
	"beego-postgres-blog/controllers"
	"github.com/astaxie/beego"
	"beego-postgres-blog/models"
	"beego-postgres-blog/common"
	"beego-postgres-blog/requests"
)

type PostController struct {
	controllers.BaseController
}

func (p *PostController)URLMapping()  {
	p.Mapping("Index",p.Index)
	p.Mapping("Create",p.Create)
	p.Mapping("Store",p.Store)
	p.Mapping("Update",p.Update)
	p.Mapping("Destroy",p.Destroy)
}

// @router /post [get]
func (p *PostController) Index()  {
	beego.ReadFromRequest(&p.Controller)
	page := p.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	post,_ := services.GetMyAllPost(page2)
	totalPage,lastPage,currentPage,nextPage := models.PostPaginate(page2,"console")
	p.Data["xsrfdata"]=template.HTML(p.XSRFFormHTML())
	p.Data["post"] = post
	p.Data["totalPage"] = totalPage
	p.Data["lastPage"] = lastPage
	p.Data["currentPage"] = currentPage
	p.Data["nextPage"] = nextPage
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/index.tpl"
}

// @router /post/create [get]
func (p *PostController) Create() {

	beego.ReadFromRequest(&p.Controller)
	cate := services.GetAllCateBySort()
	p.Data["xsrfdata"]=template.HTML(p.XSRFFormHTML())
	p.Data["cate"] = cate
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/create.tpl"
}

// @router /post [post]
func (p *PostController) Store()  {
	u := common.PostCreate{}
	if err := p.ParseForm(&u); err != nil {
		p.MyReminder("error","")
	}

	code ,message := requests.IphptValidate(p.Ctx,"Post")
	if code != 0 {
		p.MyReminder("error",message)
		p.Redirect("/console/post/create",302)
		return
	}

	id := p.GetSession("Id")
	if id == nil {
		p.MyReminder("error","登录失效")
		p.Redirect("/auth/login",302)
		return
	}
	idInt64,ok := id.(int64)
	if !ok {
		p.MyReminder("error","登录失效哦")
		p.Redirect("/auth/login",302)
		return
	}
	res := services.StorePost(u,idInt64)
	if res {
		p.MyReminder("success","创建文章成功")
	} else {
		p.MyReminder("error","创建文章失败,请检查后再试")
	}
	p.Redirect("/console/post",302)

}

// @router /post/:id([0-9]+/edit [get]
func (p *PostController) Edit() {
	beego.ReadFromRequest(&p.Controller)
	id :=p.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	post,_ := models.GetArticlesById(id64)
	cateList := services.GetAllCateBySort()
	_,maps := models.GetTagIdByPostId(id64)
	tagNames := models.GetTagNameArrByTagIds(maps)
	artCate,_ := models.GetCateIdByPostId(id64)
	postCate,_ := models.GetCategoriesById(artCate.CateId)
	p.Data["xsrfdata"]=template.HTML(p.XSRFFormHTML())
	p.Data["post"] = post
	p.Data["cate"] = cateList
	p.Data["tag"] = tagNames
	p.Data["PostCateId"] = postCate.Id
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/edit.tpl"
}

// @router /post/:id([0-9]+ [put]
func (p *PostController) Update()  {
	id :=p.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)

	u := common.PostCreate{}
	if err := p.ParseForm(&u); err != nil {
		p.MyReminder("error","")
	}

	code ,message := requests.IphptValidate(p.Ctx,"Post")
	if code != 0 {
		p.MyReminder("error",message)
		p.Redirect("/console/post/create",302)
		return
	}

	postUpdate := models.Articles{
		Id				:	id64,
		Title			:	u.Title,
		UserId 			: 	1,
		Content 		: 	u.Content,
		Abstract 		: 	u.Abstract,
		BodyOriginal 	: 	u.BodyOriginal,
	}

	err := services.PostUpdate(postUpdate,id64,u.Category,u.Tag)
	if err != nil {
		p.MyReminder("error","修改失败,请检查后再试")
	} else {
		p.MyReminder("success","修改成功")
	}
	p.Redirect("/console/post",302)

}

// @router /post/:id([0-9]+ [delete]
func (p *PostController) Destroy()  {
	id :=p.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	err := services.PostDestroy(id64)
	if  err != nil {
		p.MyReminder("error","操作失败,请检查代码是啥问题")
	} else {
		p.MyReminder("success","操作成功")
	}
	p.Redirect("/console/post",302)
}
