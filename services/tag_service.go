package services

import (
	"bee-go-myBlog/models"
	"github.com/astaxie/beego/orm"
)

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