package requests

import (
	"github.com/astaxie/beego/validation"
	validation2 "bee-go-myBlog/validation"
	"bee-go-myBlog/common"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)


func IphptValidate(c *context.Context,env string) (code int,message string) {
	var u interface{}
	switch env {
		case "Post":
			u = &common.PostCreate{}
		case "Cate":
			u = &common.CateRequest{}
		case "Link":
			u = &common.LinkCreate{}
		case "Tag":
			u = &common.TagRequest{}
		case "System":
			u = &common.SystemUpdate{}
		case "Register":
			u = &common.RegisterRequest{}
		default:
			//u = common.PostCreate{}
	}
	beego.ParseForm(c.Request.Form,u)
	valid := validation.Validation{}
	b, err := valid.Valid(u)
	if err != nil {
		return 4000000001,"内部表单校验问题,请检查后再试"
	}
	if !b {
		return RequestValidate(env,valid)
	} else {
		return 0,""
	}
}


func RequestValidate(env string,valid validation.Validation) (int,string) {
	code,message := 4000000000,"系统内部错误"
	for _, err := range valid.Errors {
		code,message = validation2.CustomErrCodeAndMessage(env,err.Field, err.Name)
		fmt.Println(code,message)
		return code,message
	}
	return code,message
}
