package auth

import (
	"github.com/astaxie/beego"
	"strconv"
	"bee-go-myBlog/services"
	"fmt"
	"html/template"
)

type PostController struct {
	beego.Controller
}

type PostCreate struct {
	Title 		string 	`form:"title"`
	Category 	int64 	`form:"category"`
	Tag 		string 	`form:"tag"`
	Content 	string 	`form:"content"`
}

func (p *PostController) Index()  {
	page := p.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	post,_ := services.GetMyAllPost(page2)
	p.Data["post"] = post
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/index.tpl"
}

func (p *PostController) Create() {

	cate := services.GetAllCateBySort()
	p.Data["xsrfdata"]=template.HTML(p.XSRFFormHTML())
	p.Data["cate"] = cate
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/create.tpl"
}

func (p *PostController) Store()  {
	u := PostCreate{}
	if err := p.ParseForm(&u); err != nil {
		fmt.Println(err)
	}

	fmt.Println(
		u.Title,
		u.Tag)


	p.ServeJSON(true)

}

func (p *PostController) Show() {
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/show.tpl"
}

func (p *PostController) Update()  {

}

func (p *PostController) Destroy()  {

}