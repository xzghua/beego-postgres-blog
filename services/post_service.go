package services

import (
	"bee-go-myBlog/models"
	"github.com/astaxie/beego/orm"

)

func GetMyAllPost(page int64) (ml []interface{}, err error){
	post,err := models.AllArticle(page)
	if post != nil {
		for key, val := range post {
			userId := val.(map[string]interface{})["UserId"].(int64)
			user,_ := models.GetUsersById(userId)
			post[key].(map[string]interface{})["user_name"] = user.Name
			post[key].(map[string]interface{})["user_id"] = user.Id
		}
	}
	return post,err
}

func AddPostCateRel(postId int64,cateId int64) (id int64, err error) {
	o := orm.NewOrm()
	artCate := models.ArticleCate{CateId: cateId,ArtId:postId}
	if _, id, err := o.ReadOrCreate(&artCate, "art_id","cate_id"); err == nil {
		return id,err
	}
	return
}

func AddPostTagRel(postId int64,tagId int64) (id int64, err error) {
	o := orm.NewOrm()
	artCate := models.ArticleTag{TagId: tagId,ArtId:postId}
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	if _, id, err := o.ReadOrCreate(&artCate, "tag_id","art_id"); err == nil {
		return id,err
	}
	return
}
