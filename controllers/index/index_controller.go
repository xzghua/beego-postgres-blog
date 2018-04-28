package index

import (
	"bee-go-myBlog/controllers"
	"strconv"
	"bee-go-myBlog/services"
	"bee-go-myBlog/models"
	"github.com/astaxie/beego"
	"fmt"
)

type HomeController struct {
	controllers.BaseController
}

//@router / [get]
func (h *HomeController) Index() {
	page := h.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}

	post,err := services.IndexPostList(page2)
	cate := services.IndexAllCateBySort()
	tag := services.IndexAllTag()
	system := services.IndexSystem()
	link := services.IndexLinkList()
	totalPage,lastPage,currentPage,nextPage := models.PostPaginate(page2,"index")
	h.Data["totalPage"] = totalPage
	h.Data["lastPage"] = lastPage
	h.Data["currentPage"] = currentPage
	h.Data["nextPage"] = nextPage
	fmt.Println(totalPage,lastPage,currentPage,nextPage,"//////")
	h.Data["system"] = system
	h.Data["post"] = post
	for _,v := range tag {
		fmt.Println(v)
	}
	fmt.Println(system)
	h.Data["cate"] = cate
	h.Data["link"] = link
	h.Data["tag"] = tag
	staticMode := beego.AppConfig.String("staticmode")
	if staticMode == "local" {
		h.Layout = "home/local/master.tpl"
		h.TplName = "home/local/index.tpl"
	} else {
		h.Layout = "home/master.tpl"
		h.TplName = "home/index.tpl"
	}

}

func (h *HomeController) Detail() {

}

func (h *HomeController) Cate() {

}

func (h *HomeController) Tag() {

}

func (h *HomeController) Archive() {

}