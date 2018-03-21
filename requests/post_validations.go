package requests

import (
	"github.com/astaxie/beego/validation"
	"fmt"
	"log"
)

func PostCreateValidate (u interface{}) {
	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	fmt.Println("结果1")

	if err != nil {
		// handle error
	}
	//flash:=beego.NewFlash()
	fmt.Println("结果2")

	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println( err.Message)
			//flash.Error(err.Message)
			//flash.Store(&p.Controller)
			//p.Redirect("/console/post/create",302)
			return
		}
	}

}