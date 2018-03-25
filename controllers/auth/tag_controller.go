package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/services"
)

type TagController struct {
	controllers.BaseController
}


func (c *TagController) GetTagByLike() {
	param := c.GetString("term")
	res := services.GetTagByLike(param)
	c.Data["json"] = res
	c.ServeJSON()
}