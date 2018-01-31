package auth

import (
	"github.com/astaxie/beego"
	"strconv"
	"bee-go-myBlog/services"
)

type PostController struct {
	beego.Controller
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

	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/create.tpl"
}



func (p *PostController) Store()  {

}

func (p *PostController) Show() {
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/show.tpl"
}

func (p *PostController) Update()  {

}

func (p *PostController) Destroy()  {

}