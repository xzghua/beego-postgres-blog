package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"

	validation2 "bee-go-myBlog/validation"
	"fmt"
)

type BaseController struct {
	beego.Controller
}


func (c *BaseController) RequestValidate(valid validation.Validation) (int,string) {

	code,message := 4000000000,"系统内部错误"
	for _, err := range valid.Errors {
		code,message = validation2.CustomErrCodeAndMessage("Post",err.Field, err.Name)
		fmt.Println(code,message)
		return code,message
	}
	return code,message
}