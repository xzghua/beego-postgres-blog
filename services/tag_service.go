package services

import (
	"bee-go-myBlog/models"
	"github.com/astaxie/beego/orm"
	"bee-go-myBlog/common"
	"time"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"github.com/astaxie/beego"
)


func IndexAllTag() ([]interface{}) {
	cache := common.Cache()
	key := "index:tag:list"
	res := cache.Get(key)

	if res == nil {
		tag,_ := models.IndexAllTag()
		if len(tag) == 0 {
			tagCreate := &models.Tags{
				Name	:	"默认分类",
				TagNum	:	0,
			}
			models.AddTags(tagCreate)
			tag,_ = models.IndexAllTag()
		}
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(tag)
		cache.Put(key,data,timeoutDuration)
		return tag
	}

	var err error
	var tagList  []interface{}
	string1,_ := redis.String(res,err)
	json.Unmarshal([]byte(string1),&tagList)
	return tagList

}

func GetAllMyTag(page int64) (interface{},error ){
	tag,err := models.AllTag(page)
	//if err == nil {
	//	return tag,nil
	//}
	if len(tag) == 0 {
		tagCreate := &models.Tags{
			Name	:	"默认分类",
			TagNum	:	0,
		}
		models.AddTags(tagCreate)
		tag,_ = models.AllTag(page)
	}
	return tag,err
}



func GetTagByLike(param string) (ml map[int64]string) {
	var query map[string]string
	var fields []string
	var sortby []string
	var order []string
	var offset int64
	var limit int64

	query = make(map[string]string)
	if param != "" {
		query["name__icontains"] = param
	}

	fields = []string{"Id", "Name", "TagNum"}
	cates,_ := models.GetAllTags(query, fields, sortby, order, offset, limit)
	cate := make(map[int64]string, len(cates))
	for _, v := range cates{
		cate[v.(map[string]interface{})["Id"].(int64)] = v.(map[string]interface{})["Name"].(string)
	}

	return cate
}

func DeleteTag(tagId int64) (res bool) {
	o := orm.NewOrm()
	o.Begin()
	_,delArtTagErr := o.Delete(&models.ArticleTag{TagId:tagId},"TagId")

	_,err :=o.Delete(&models.Tags{Id:tagId})
	if err ==nil && delArtTagErr == nil {
		o.Commit()
		return true
	} else {
		o.Rollback()
		return false
	}
}


func IndexGetTagByName(tagName string) (v models.Tags,err error) {
	cache := common.Cache()
	key := "index:tag:search:"+tagName
	res := cache.Get(key)
	if res == nil {
		o := orm.NewOrm()
		v = models.Tags{Name:tagName}
		err = o.Read(&v,"Name")
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(v)
		cache.Put(key,data,timeoutDuration)
		return v,err
	}
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&v)
	return v,err
}

func IndexGetTagPost(tagId int64,page2 int64) (tagPost []*models.ArticleTag,err error) {
	cache := common.Cache()
	tagIdString:=strconv.FormatInt(tagId,10)
	pageString := strconv.FormatInt(page2,10)
	key := "index:tag:posts:by:tagid:"+tagIdString+":page:"+pageString
	res := cache.Get(key)
	if res == nil {
		o := orm.NewOrm()
		limit, _ := beego.AppConfig.Int64("page_offset")
		offset := (page2 - 1) * limit
		_,err = o.QueryTable(new(models.ArticleTag)).Filter("TagId",tagId).OrderBy("Id").Limit(limit,offset).All(&tagPost)
		timeoutDuration := 24 * 30 * time.Hour
		data ,_ := json.Marshal(tagPost)
		cache.Put(key,data,timeoutDuration)
		return tagPost,err
	}
	string1,err := redis.String(res,err)
	json.Unmarshal([]byte(string1),&tagPost)
	return
}