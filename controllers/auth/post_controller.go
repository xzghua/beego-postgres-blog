package auth

import (
	"github.com/astaxie/beego"
	"strconv"
	"bee-go-myBlog/services"
	"fmt"
	"html/template"
	"github.com/astaxie/beego/validation"
	"log"
)

type PostController struct {
	beego.Controller
}

type PostCreate struct {
	Title 		string 	`form:"title" valid:"Required;MaxSize(2)"`
	Category 	int64 	`form:"category" valid:"Required"`
	Tag 		string 	`form:"tag" valid:"Required"`
	Content 	string 	`form:"content" valid:"Required"`
}

//type PostCreate struct {
//	Title 		string
//	Category 	int64
//	Tag 		string
//	Content 	string
//}


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

	//fmt.Println(u)

	minAge := 12

	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	valid.MaxSize(u.Title, minAge, "age").Message("少儿不宜！")

	if err != nil {
		// handle error
	}
	flash:=beego.NewFlash()

	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			flash.Error(err.Message)
			flash.Store(&p.Controller)
			p.Redirect("/console/post/create",302)
			return
		}
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