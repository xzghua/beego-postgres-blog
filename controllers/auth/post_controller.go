package auth

import (
	"strconv"
	"bee-go-myBlog/services"
	"fmt"
	"html/template"
	"bee-go-myBlog/controllers"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	"strings"
	"bee-go-myBlog/models"
)

type PostController struct {
	controllers.BaseController
}

type PostCreate struct {
	Title 			string 	`form:"title" valid:"Required;MaxSize(23)"`
	Abstract 		string	`form:"abstract" valid:"Required;MaxSize(50)"`
	Category 		int64 	`form:"category" valid:"Required"`
	Tag 			string 	`form:"tag" valid:"Required"`
	Content 		string 	`form:"editormd-html-code" valid:"Required"`
	BodyOriginal 	string 	`form:"content" valid:"Required"`
}

func (p *PostController) Index()  {
	page := p.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	post,_ := services.GetMyAllPost(page2)
	totalPage,lastPage,currentPage,nextPage := models.PostPaginate(page2)
	p.Data["post"] = post
	p.Data["totalPage"] = totalPage
	p.Data["lastPage"] = lastPage
	p.Data["currentPage"] = currentPage
	p.Data["nextPage"] = nextPage
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/index.tpl"
}

func (p *PostController) Create() {

	beego.ReadFromRequest(&p.Controller)
	cate := services.GetAllCateBySort()
	p.Data["xsrfdata"]=template.HTML(p.XSRFFormHTML())
	p.Data["cate"] = cate
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/create.tpl"
}

func (p *PostController) Store()  {
	u := PostCreate{}
	valid := validation.Validation{}
	if err := p.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	if err != nil {
	}
	if !b {
		_,message :=p.RequestValidate(valid)
		p.MyReminder("error",message)
		p.Redirect("/console/post/create",302)
		return
	}

	fmt.Println(u)

	var post models.Articles

	post.Title = u.Title
	post.Content = u.Content
	post.Abstract = u.Abstract
	post.BodyOriginal = u.BodyOriginal
	post.UserId = 1

	postId,err :=models.AddArticles(&post)
	if err != nil {
		p.MyReminder("error","系统内部错误")
	}
	_,er := services.AddPostCateRel(postId,u.Category)
	if er != nil {
		p.MyReminder("error","系统内部错误")
	}

	tag := strings.Split(u.Tag,",")
	for _,v := range tag {
		tagId,_ := models.AddTagWithUnique(v)
		services.AddPostTagRel(postId,tagId)
	}

	p.Redirect("/console/post",302)

}

func (p *PostController) Edit() {
	id :=p.Ctx.Input.Param(":id")
	fmt.Println(id)
	id64, _ := strconv.ParseInt(id, 10, 64)
	post,_ := models.GetArticlesById(id64)
	cateList := services.GetAllCateBySort()
	_,maps := models.GetTagIdByPostId(id64)
	tags := models.GetTagByTagIds(maps)
	artCate,_ := models.GetCateIdByPostId(id64)
	postCate,_ := models.GetCategoriesById(artCate.CateId)
	//fmt.Println(artTag)

	p.Data["post"] = post
	p.Data["cate"] = cateList
	p.Data["tag"] = tags
	p.Data["PostCateId"] = postCate.Id
	fmt.Println(postCate.Id)
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/edit.tpl"
}

func (p *PostController) Update()  {

}

func (p *PostController) Destroy()  {

}