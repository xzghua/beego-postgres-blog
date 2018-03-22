package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"github.com/astaxie/beego/validation"

)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) RequestValidate(valid validation.Validation ) {

	for _, err := range valid.Errors {
		log.Println(err.Key, err.Message)
		return
	}

}