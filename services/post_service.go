package services

import (
	"bee-go-myBlog/models"
	"github.com/astaxie/beego/orm"
	"fmt"
	"reflect"
)

func GetMyAllPost(page int64) (ml []interface{}, err error){
	post,err := models.AllArticle(page)
	if post != nil {
		for key, val := range post {
			userId := val.(map[string]interface{})["UserId"].(int64)
			user,_ := models.GetUsersById(userId)
			postId := val.(map[string]interface{})["Id"].(int64)
			fmt.Println(postId,"帖子ID",reflect.TypeOf(postId))

			cateId,_ := GetCateByPostId(postId)
			cate,_ := models.GetCategoriesById(cateId)
			if cate == nil {
				post[key].(map[string]interface{})["cate_name"] = ""
			} else {
				post[key].(map[string]interface{})["cate_name"] = cate.DisplayName
			}
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
	if news, id, err := o.ReadOrCreate(&artCate, "tag_id","art_id"); err == nil {
		if !news {
			tag,_ :=models.GetTagsById(tagId)
			tagUpdate := &models.Tags{
				Id:tag.Id,
				TagNum:tag.TagNum + 1,
			}
			o.Update(tagUpdate,"TagNum")
		}
		return id,err
	}

	return
}

func DelPostCateRel(postId int64,cateId int64) ( err error) {
	o := orm.NewOrm()
	artCate := models.ArticleCate{CateId: cateId,ArtId:postId}
	err = o.Read(&artCate,"CateId","ArtId")
	if err == nil {
		err = models.DeleteArticleCate(artCate.Id)
		return
	}
	return nil
}

func DelPostTagRel(postId int64) ( err error) {
	//postId = 223
	o := orm.NewOrm()
	var maps []orm.Params
	_,err =o.QueryTable(new(models.ArticleTag)).Filter("ArtId",postId).Values(&maps)
	for _,v := range maps {
		tag,_ :=models.GetTagsById(v["TagId"].(int64))
		tagUpdate := &models.Tags{
			Id:tag.Id,
			TagNum:tag.TagNum - 1,
		}
		o.Update(tagUpdate,"TagNum")
	}

	return err
}