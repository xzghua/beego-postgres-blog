package index

import (
	"bee-go-myBlog/controllers"
	"strconv"
	"bee-go-myBlog/services"
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
	services.IndexPostList(page2)
	h.Layout = "home/master.tpl"
	h.TplName = "home/index.tpl"
}

func (h *HomeController) Detail() {

}

func (h *HomeController) Cate() {

}

func (h *HomeController) Tag() {

}

func (h *HomeController) Archive() {

}