package auth

import (
	"github.com/astaxie/beego"
	"fmt"
	"bee-go-myBlog/models"
	"strconv"
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

	post,err := models.AllArticle(page2)

	var userIds []int64
	if post != nil {
		for key, val := range post {
			userId := val.(map[string]interface{})["Id"].(int64)
			userIds = append(userIds,userId)
			user,_ := models.GetUsersById(userId)
			post[key].(map[string]interface{})["user_name"] = user.Name
			post[key].(map[string]interface{})["user_id"] = user.Id
		}
	}

	fmt.Println(post,"打印结果")

	p.Data["post"] = post
	p.Layout = "auth/master.tpl"
	p.TplName = "auth/post/index.tpl"
}

func (p *PostController) Create() {



	var post models.Articles
	//post.Title = "测试标题"
	//post.UserId = 1
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