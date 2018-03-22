package auth

import (
	"strconv"
	"bee-go-myBlog/services"
	"fmt"
	"html/template"
	"bee-go-myBlog/controllers"
	"github.com/astaxie/beego/validation"
)

type PostController struct {
	controllers.BaseController
}

type PostCreate struct {
	Title 		string 	`form:"title" valid:"Required;MaxSize(2)"`
	Category 	int64 	`form:"category" valid:"Required"`
	Tag 		string 	`form:"tag" valid:"Required"`
	Content 	string 	`form:"content" valid:"Required"`
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
	valid := validation.Validation{}
	if err := p.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	//valid.MaxSize(u.Title, 2, "Title").Message("少儿不宜！")
	b, err := valid.Valid(&u)
	if err != nil {
	}

	fmt.Println(u)

	if !b {
		fmt.Println("234234是是是是")
		p.RequestValidate(valid)
	}



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