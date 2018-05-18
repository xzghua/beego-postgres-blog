package index

import (
	"beego-postgres-blog/controllers"
	"strconv"
	"beego-postgres-blog/services"
	"beego-postgres-blog/models"
	"fmt"
)

type HomeController struct {
	controllers.BaseController
}

func (h *HomeController)URLMapping()  {
	h.Mapping("Index",h.Index)
	h.Mapping("Detail",h.Detail)
	h.Mapping("Cate",h.Cate)
	h.Mapping("Tag",h.Tag)
	h.Mapping("Archive",h.Archive)
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
	fmt.Println(system)
	link := services.IndexLinkList()
	totalPage,lastPage,currentPage,nextPage := models.PostPaginate(page2,"index")
	h.Data["totalPage"] = totalPage
	h.Data["lastPage"] = lastPage
	h.Data["currentPage"] = currentPage
	h.Data["nextPage"] = nextPage
	h.Data["post"] = post

	h.Data["system"] = system
	h.Data["cate"] = cate
	h.Data["link"] = link
	h.Data["tag"] = tag


	if system.CdnType == 1 {
		h.Layout = "home/local/master.tpl"
		h.TplName = "home/local/index.tpl"
	} else {
		h.Layout = "home/qi/master.tpl"
		h.TplName = "home/qi/index.tpl"
	}

}

//@router /detail/:id([0-9]+ [get]
func (h *HomeController) Detail() {
	id := h.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	post := services.IndexPostDetail(id64)
	lastPost := services.IndexPostLast(id64)
	lastBefore := services.IndexPostBefore(id64)
	cate := services.IndexAllCateBySort()
	tag := services.IndexAllTag()
	system := services.IndexSystem()
	link := services.IndexLinkList()
	if lastPost == nil {
		h.Data["lastPostCond"] = false
	} else {
		h.Data["lastPostCond"] = true
	}
	if lastBefore == nil {
		h.Data["lastBeforeCond"] = false
	} else {
		h.Data["lastBeforeCond"] = true
	}
	postTag := services.IndexPostTag(id64)

	go services.PostReadNumAdd(id64)
	//评论
	//阅读数+1
	h.Data["system"] = system
	h.Data["cate"] = cate
	h.Data["link"] = link
	h.Data["tag"] = tag
	h.Data["postTag"] = postTag
	h.Data["detail"] = post
	h.Data["lastPost"] = lastPost
	h.Data["lastBefore"] = lastBefore

	if system.CdnType == 1 {
		h.Layout = "home/local/master.tpl"
	} else {
		h.Layout = "home/qi/master.tpl"
	}
	h.TplName = "home/detail.tpl"

}

//@router /categories/:cate [get]
func (h *HomeController) Cate() {
	page := h.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	cate := h.Ctx.Input.Param(":cate")
	categories,err := services.IndexGetCateByName(cate)
	if err != nil {
		//走404页面
	}
	postCates,err := services.IndexGetCatePost(categories.Id,page2)
	if err != nil {
	}
	posts,err := services.IndexPostByCateIds(postCates,categories.Id,page2)
	if err != nil {
	}

	totalPage,lastPage,currentPage,nextPage := models.IndexCatePostPaginate(page2,categories.Id)
	h.Data["totalPage"] = totalPage
	h.Data["lastPage"] = lastPage
	h.Data["currentPage"] = currentPage
	h.Data["nextPage"] = nextPage

	cates := services.IndexAllCateBySort()
	tag := services.IndexAllTag()
	system := services.IndexSystem()
	link := services.IndexLinkList()
	h.Data["posts"] = posts
	h.Data["system"] = system
	h.Data["cate"] = cates
	h.Data["link"] = link
	h.Data["tag"] = tag

	if system.CdnType == 1 {
		h.Layout = "home/local/master.tpl"
		h.TplName = "home/local/cate.tpl"
	} else {
		h.Layout = "home/qi/master.tpl"
		h.TplName = "home/qi/cate.tpl"
	}
}

//@router /tags/:tag [get]
func (h *HomeController) Tag() {
	page := h.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	tag := h.Ctx.Input.Param(":tag")
	tagOne,err := services.IndexGetTagByName(tag)
	if err != nil {
	}
	postTags,err := services.IndexGetTagPost(tagOne.Id,page2)
	if err != nil {
	}
	posts,err := services.IndexPostByTagIds(postTags,tagOne.Id,page2)
	if err != nil {
	}

	totalPage,lastPage,currentPage,nextPage := models.IndexTagPostPaginate(page2,tagOne.Id)
	h.Data["totalPage"] = totalPage
	h.Data["lastPage"] = lastPage
	h.Data["currentPage"] = currentPage
	h.Data["nextPage"] = nextPage

	cate := services.IndexAllCateBySort()
	tags := services.IndexAllTag()
	system := services.IndexSystem()
	link := services.IndexLinkList()
	h.Data["posts"] = posts
	h.Data["system"] = system
	h.Data["cate"] = cate
	h.Data["link"] = link
	h.Data["tag"] = tags

	if system.CdnType == 1 {
		h.Layout = "home/local/master.tpl"
		h.TplName = "home/local/tag.tpl"
	} else {
		h.Layout = "home/qi/master.tpl"
		h.TplName = "home/qi/tag.tpl"
	}

}

//@router /archive [get]
func (h *HomeController) Archive() {

	postList,timeArr := services.IndexPostAllList()
	if postList == nil {
		//走404
	}

	cate := services.IndexAllCateBySort()
	tags := services.IndexAllTag()
	system := services.IndexSystem()
	link := services.IndexLinkList()
	h.Data["system"] = system
	h.Data["cate"] = cate
	h.Data["link"] = link
	h.Data["tag"] = tags
	h.Data["postList"] = postList
	h.Data["timeArr"] = timeArr
	if system.CdnType == 1 {
		h.Layout = "home/local/master.tpl"
	} else {
		h.Layout = "home/qi/master.tpl"
	}
	h.TplName = "home/archive.tpl"
}