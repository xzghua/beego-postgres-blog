package services

import (
	"bee-go-myBlog/models"
	"github.com/astaxie/beego/orm"
	"fmt"

	"strings"
	"time"
	"bee-go-myBlog/common"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

func StorePost(u common.PostCreate) (res bool) {
	var post models.Articles
	post.Title = u.Title
	post.Content = u.Content
	post.Abstract = u.Abstract
	post.BodyOriginal = u.BodyOriginal
	post.UserId = 1

	o := orm.NewOrm()
	o.Begin()
	postId,postAddErr :=models.AddArticles(&post)
	if postAddErr != nil {
		o.Rollback()
		return false
	}
	_,postCateRelErr := AddPostCateRel(postId,u.Category)
	if postCateRelErr != nil {
		o.Rollback()
		return false
	}
	tag := strings.Split(u.Tag,",")
	for _,v := range tag {
		tagId,err := models.AddTagWithUnique(v)
		if err != nil {
			o.Rollback()
			return false
		}
		_,err = AddPostTagRel(postId,tagId)
		if err != nil {
			o.Rollback()
			return false
		}
	}
	o.Commit()
	return true
}



func GetMyAllPost(page int64) (ml []interface{}, err error){
	post,err := models.AllArticle(page)
	if post != nil {
		for key, val := range post {
			userId := val.(map[string]interface{})["UserId"].(int64)
			user,_ := models.GetUsersById(userId)
			postId := val.(map[string]interface{})["Id"].(int64)
			deletedAt := val.(map[string]interface{})["DeletedAt"].(time.Time)
			fmt.Println(postId,"帖子ID",deletedAt.Unix(),deletedAt.Unix() > 0)
			if deletedAt.Unix() > 0 {
				post[key].(map[string]interface{})["time_status"] = true
 			} else {
				post[key].(map[string]interface{})["time_status"] = false
			}
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

func DelPostCateRel(postId int64) ( err error) {
	o := orm.NewOrm()
	artCate := models.ArticleCate{ArtId:postId}
	err = o.Read(&artCate,"ArtId")
	if err == nil {
		_,err = o.Delete(&models.ArticleCate{ArtId:postId},"ArtId")
		return
	}
	return nil
}

func DelPostTagRel(postId int64) ( err error) {
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
	_,err =o.Delete(&models.ArticleTag{ArtId:postId},"ArtId")
	return err
}


func PostUpdate(postUpdate models.Articles,id64 int64,cate int64,tags string) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	updateArtErr := models.UpdateArticlesById(&postUpdate)
	if updateArtErr != nil {
		err = updateArtErr
	}
	delPostCateErr := DelPostCateRel(id64)
	if delPostCateErr != nil {
		err = delPostCateErr
	}
	_,addPostCateErr := AddPostCateRel(id64,cate)
	if addPostCateErr != nil {
		err = addPostCateErr
	}

	delPostTagErr := DelPostTagRel(id64)
	if delPostTagErr != nil {
		err = delPostTagErr
	}
	tag := strings.Split(tags,",")
	for _,v := range tag {
		tagId,addTagErr := models.AddTagWithUnique(v)
		if addTagErr != nil {
			err = addTagErr
		}
		_,addPostTagErr := AddPostTagRel(id64,tagId)
		if addPostTagErr != nil {
			err = addPostTagErr
		}
	}

	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}
	return
}

func PostDestroy(id int64) (err error) {
	o := orm.NewOrm()
	res,err :=models.GetArticlesById(id)
	fmt.Println(res.DeletedAt.Unix() > 0,res.DeletedAt.Unix())
	if res.DeletedAt.Unix() > 0 {
		var theTime = new(time.Time)
		_,err = o.Update(&models.Articles{Id:id,DeletedAt:*theTime},"DeletedAt")
	} else {
		nowTime := time.Now()
		_,err = o.Update(&models.Articles{Id:id,DeletedAt:nowTime},"DeletedAt")
	}

	return
}


func IndexPostList(page int64) (ml []interface{}, err error) {
	cache := common.Cache()
	pageString :=strconv.FormatInt(page,10)
	key := "index:post:list:"+pageString
	res := cache.Get(key)

	if res == nil {
		post,err := models.IndexAllPost(page)
		if post != nil {
			for key, val := range post {
				userId := val.(map[string]interface{})["UserId"].(int64)
				user,_ := models.GetUsersById(userId)
				postId := val.(map[string]interface{})["Id"].(int64)
				deletedAt := val.(map[string]interface{})["DeletedAt"].(time.Time)
				fmt.Println(postId,"帖子ID",deletedAt.Unix(),deletedAt.Unix() > 0)
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
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(post)
		cache.Put(key,data,timeoutDuration)
		return post,err
	}

	var post  []interface{}
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&post)

	return post,err

}