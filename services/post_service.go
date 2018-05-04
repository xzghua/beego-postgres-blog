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
		post,_ := models.IndexAllPostWithPaginate(page)
		if post != nil {
			for key, val := range post {
				userId := val.(map[string]interface{})["UserId"].(int64)
				user,_ := models.GetUsersById(userId)
				postId := val.(map[string]interface{})["Id"].(int64)
				_,maps := models.GetTagIdByPostId(postId)
				tagNames := models.GetTagByTagIds(maps)
				post[key].(map[string]interface{})["tag_list"] = tagNames
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
		//为什么这地方, 我还要再麻烦的存进去又取出来呢..  因为 他喵的 Tag的类型 虽然都是数组,但是一开始里面是 interface{},
		//后来进了缓存,再拿出来又是 map类型的..  妈了个鸡..  新手,不知道怎么搞
		res = cache.Get(key)
		//return post,err
	}

	var post  []interface{}
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&post)

	return post,err

}

func IndexPostDetail (id int64) (post *models.Articles) {
	cache := common.Cache()
	IdString :=strconv.FormatInt(id,10)
	key := "index:post:detail:"+IdString
	res := cache.Get(key)

	if res == nil {
		post,_ := models.GetArticlesById(id)
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(post)
		cache.Put(key,data,timeoutDuration)
		return post
	}
	var err error
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&post)
	return post
}

func IndexPostTag(id int64) (tags []interface{}) {
	cache := common.Cache()
	IdString :=strconv.FormatInt(id,10)
	key := "index:post:tag:"+IdString
	res := cache.Get(key)

	if res == nil {
		_,maps := models.GetTagIdByPostId(id)
		tags := models.GetTagByTagIds(maps)
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(tags)
		cache.Put(key,data,timeoutDuration)
		res = cache.Get(key)
	}
	var err error
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&tags)
	return tags
}

func IndexPostLast(id int64) (post *models.Articles) {
	cache := common.Cache()
	IdString :=strconv.FormatInt(id,10)
	key := "index:post:last:"+IdString
	res := cache.Get(key)
	if res == nil {
		post,_ := models.GetIndexArticlesById(id,"Id__lt","-Id")
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(post)
		cache.Put(key,data,timeoutDuration)
		return post
	}
	var err error
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&post)
	return post
}

func IndexPostBefore(id int64) (post *models.Articles) {
	cache := common.Cache()
	IdString :=strconv.FormatInt(id,10)
	key := "index:post:before:"+IdString
	res := cache.Get(key)
	if res == nil {
		post,_ := models.GetIndexArticlesById(id,"Id__gt","Id")
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(post)
		cache.Put(key,data,timeoutDuration)
		return post
	}
	var err error
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&post)
	return post
}

func PostReadNumAdd(id64 int64) {
	o := orm.NewOrm()
	v := models.Articles{Id: id64}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		var num int64
		v.ViewNum = v.ViewNum + 1
		if num, err = o.Update(&v, "ViewNum"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func IndexPostByCateIds(ids []*models.ArticleCate,cateId int64,page int64) (posts []models.Articles,err error) {
	cache := common.Cache()
	pageString := strconv.FormatInt(page,10)
	cateIdString:=strconv.FormatInt(cateId,10)
	key := "index:posts:data:by:cate:"+cateIdString+"page:"+pageString
	res := cache.Get(key)
	if res == nil {
		o := orm.NewOrm()
		var post models.Articles
		for _,v := range ids {
			err := o.QueryTable(new(models.Articles)).Filter("Id",v.ArtId).One(&post)
			if err != nil {
				return posts,err
			}
			posts = append(posts,post)
		}
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(posts)
		cache.Put(key,data,timeoutDuration)
	}
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&posts)
	return posts,nil
}

func IndexPostByTagIds(ids []*models.ArticleTag,tagId int64,page int64) (posts []models.Articles,err error) {
	cache := common.Cache()
	pageString := strconv.FormatInt(page,10)
	tagIdString:=strconv.FormatInt(tagId,10)
	key := "index:posts:data:by:tag:"+tagIdString+"page:"+pageString
	res := cache.Get(key)
	if res == nil {
		o := orm.NewOrm()
		var post models.Articles
		for _,v := range ids {
			err := o.QueryTable(new(models.Articles)).Filter("Id",v.ArtId).One(&post)
			if err != nil {
				return posts,err
			}
			posts = append(posts,post)
		}
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(posts)
		cache.Put(key,data,timeoutDuration)
	}
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&posts)
	return posts,nil
}



func IndexPostAllList() (posts map[string]map[string]interface{},timeArr []interface{}) {
	cache := common.Cache()
	key1 := "index:post:archive:post"
	key2 := "index:post:archive:timekey"
	res1 := cache.Get(key1)
	res2 := cache.Get(key2)

	if res1 == nil || res2 == nil {
		post,err := models.IndexAllPost()
		var test map[string]interface{}
		test = make(map[string]interface{})
		posts = make(map[string]map[string]interface{})
		if err == nil {
			for key,value := range post {
				createdAt,ok := value.(map[string]interface{})["CreatedAt"].(time.Time)
				if !ok {
					//报错
					return nil,nil
				}
				createAtString := createdAt.Format("2006-01")
				if !common.InArray(createAtString,timeArr) {
					test = make(map[string]interface{})
					timeArr = append(timeArr,createAtString)
				}
				test[strconv.Itoa(key)] = value
				posts[createAtString] = test
			}
		}
		timeoutDuration := 24 * 30 * time.Hour
		data1 ,_ := json.Marshal(posts)
		data2 ,_ := json.Marshal(timeArr)
		cache.Put(key1,data1,timeoutDuration)
		cache.Put(key2,data2,timeoutDuration)
		res1 = cache.Get(key1)
		res2 = cache.Get(key2)
	}
	var err error
	string1,err := redis.String(res1,err)
	string2,err := redis.String(res2,err)
	json.Unmarshal([]byte(string1),&posts)
	json.Unmarshal([]byte(string2),&timeArr)
	return posts,timeArr
}