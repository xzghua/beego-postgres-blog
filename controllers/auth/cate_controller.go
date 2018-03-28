package auth

import (
	"github.com/astaxie/beego"
	"bee-go-myBlog/services"
	"html/template"
	"github.com/astaxie/beego/validation"
	"fmt"
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/models"
	"strconv"
)

type CateController struct {
	controllers.BaseController
}

type cateRequest struct {
	ParentId		string 	`form:"parentId" valid:"Required"`
	Name 			string 	`form:"name" valid:"Required;MaxSize(30)"`
	DisplayName 	string 	`form:"displayName" valid:"Required;MaxSize(30)"`
	Description 	string 	`form:"description" valid:"MaxSize(150)"`
}

func (c *CateController) Index() {
	cate := services.GetAllCateBySort()
	c.Data["cate"] = cate
	c.Layout = "auth/master.tpl"
	c.TplName = "auth/cate/index.tpl"
}

func (c *CateController) Create() {
	beego.ReadFromRequest(&c.Controller)
	cate := services.GetAllCateBySort()
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Data["cate"] = cate
	c.Layout = "auth/master.tpl"
	c.TplName = "auth/cate/create.tpl"
}


func (c *CateController) Store() {
	u := cateRequest{}
	valid := validation.Validation{}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	fmt.Println(b,u,"开始======")
	if err != nil {
	}
	if !b {
		_,message :=c.RequestValidate(valid)
		c.MyReminder("error",message)
		c.Redirect("/console/cate/create",302)
		return
	}
	parentId,_:=strconv.ParseInt(u.ParentId, 10, 64)
	fmt.Println(u,parentId,"看看数据")
	var cateCreate = &models.Categories{
		ParentId	:	parentId,
		Name		:	u.Name,
		DisplayName	:	u.DisplayName,
		Description	:	u.Description,
	}
	_,err = models.AddCategories(cateCreate)
	if err != nil {
		c.MyReminder("error","系统内部错误")
	}

	c.Redirect("/console/cate",302)
}

func (c *CateController) Show() {

}

func (c *CateController) Update() {

}

func (c *CateController) Destroy() {

}

//作废
func (c *CateController) GetCateByLike() {
	param := c.GetString("term")
	res := services.GetCateByLike(param)
	c.Data["json"] = res
	c.ServeJSON()
}