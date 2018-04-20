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

func (c *CateController)URLMapping()  {
	c.Mapping("Index",c.Index)
	c.Mapping("Create",c.Create)
	c.Mapping("Store",c.Store)
	c.Mapping("Show",c.Show)
	c.Mapping("Update",c.Update)
	c.Mapping("Destroy",c.Destroy)
}


// @router /console/cate [get]
func (c *CateController) Index() {
	beego.ReadFromRequest(&c.Controller)
	cate := services.GetAllCateBySort()
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Data["cate"] = cate
	c.Layout = "auth/master.tpl"
	c.TplName = "auth/cate/index.tpl"
}

// @router /console/cate/create [get]
func (c *CateController) Create() {
	beego.ReadFromRequest(&c.Controller)
	cate := services.GetAllCateBySort()
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Data["cate"] = cate
	c.Layout = "auth/master.tpl"
	c.TplName = "auth/cate/create.tpl"
}

// @router /console/cate [post]
func (c *CateController) Store() {
	u := cateRequest{}
	valid := validation.Validation{}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	if err != nil {
	}
	if !b {
		_,message :=c.RequestValidate(valid)
		c.MyReminder("error",message)
		c.Redirect("/console/cate/create",302)
		return
	}
	parentId,_:=strconv.ParseInt(u.ParentId, 10, 64)
	var cateCreate = &models.Categories{
		ParentId	:	parentId,
		Name		:	u.Name,
		DisplayName	:	u.DisplayName,
		Description	:	u.Description,
	}
	_,err = models.AddCategories(cateCreate)
	if err != nil {
		c.MyReminder("error","分类创建失败,请检查后再试")
	} else {
		c.MyReminder("success","分类创建成功")
	}

	c.Redirect("/console/cate",302)
}

// @router /console/cate/:id([0-9]+/edit [get]
func (c *CateController) Show() {
	beego.ReadFromRequest(&c.Controller)
	id := c.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	cate,_ := models.GetCategoriesById(id64)
	cateSort := services.GetAllCateBySort()
	c.Data["cate"] = cate
	c.Data["cateSort"] = cateSort
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Layout = "auth/master.tpl"
	c.TplName = "auth/cate/edit.tpl"
}

// @router /console/cate/:id([0-9]+ [put]
func (c *CateController) Update() {
	u := cateRequest{}
	valid := validation.Validation{}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	if err != nil {
	}
	if !b {
		_,message :=c.RequestValidate(valid)
		c.MyReminder("error",message)
		c.Redirect("/console/cate/create",302)
		return
	}
	id := c.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	parentId,_:=strconv.ParseInt(u.ParentId, 10, 64)
	//校验一下自己的上级不能是自己的下级
	ids := []int64{id64}
	resIds := []int64{0}
	_,res2,_ := services.GetSimilar(ids,resIds,0)
	for _,v := range res2 {
		if v == parentId {
			c.MyReminder("error","不能修改为自己的子类,请检查后再试")
			c.Redirect("/console/cate/"+id+"/edit",302)
			return
		}
	}
	var cateUpdate = &models.Categories{
		Id			:	id64,
		ParentId	:	parentId,
		Name		:	u.Name,
		DisplayName	:	u.DisplayName,
		Description	:	u.Description,
	}
	err = models.UpdateCategoriesById(cateUpdate)
	if err != nil {
		c.MyReminder("error","分类修改失败,请检查后再试")
	} else {
		c.MyReminder("success","分类修改成功")
	}

	c.Redirect("/console/cate",302)
}

// @router /console/cate/:id([0-9]+ [delete]
func (c *CateController) Destroy() {
	id := c.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	ids := []int64{id64}
	resIds := []int64{0}
	_,res2,_ := services.GetSimilar(ids,resIds,0)
	if len(res2) > 1 {
		c.MyReminder("error","还有子类,不允许删除,请检查后再试")
	} else {
		services.DeleteCateById(id64)
		c.MyReminder("success","操作成功")
	}
	c.Redirect("/console/cate",302)
	return
}




//作废
func (c *CateController) GetCateByLike() {
	param := c.GetString("term")
	res := services.GetCateByLike(param)
	c.Data["json"] = res
	c.ServeJSON()
}