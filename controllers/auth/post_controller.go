package auth

import (
	"github.com/astaxie/beego"
	"fmt"
	"bee-go-myBlog/models"
)

type PostController struct {
	beego.Controller
}

func (p *PostController) Index()  {


	l,err := models.AllArticle()

	//for _,v := range l {
	//	fmt.Printf("%#v\n",v.(models.Articles).Id)
	//}

	if err == nil {
		fmt.Print(l)
	}


	p.Data["zgh"] = l
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/index.tpl"
}

func (p *PostController) Create() {

	var post models.Articles
	//post.Title = "测试标题"
	post.UserId = 1
	//post.Content = "这是一个测试的内容,测试一下效果的问题"

	id,err := models.AddArticles(&post)

	if err == nil {
		fmt.Print("都是正常的")
	} else {
		fmt.Print(id,"有问题")
	}
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/create.tpl"
	//p.LayoutSections = make(map[string]string)
	//p.LayoutSections["Scripts"] = "auth/post/post_script.tpl"
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