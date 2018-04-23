package services

import (
	"bee-go-myBlog/models"
	"github.com/astaxie/beego/orm"
	"fmt"
	"bee-go-myBlog/common"
)

func GetAllCateBySort() ([]interface{}) {
	var query map[string]string
	var fields []string
	var sortby []string
	var order []string
	var offset int64
	var limit int64

	fields = []string{"Id", "Name", "DisplayName", "ParentId", "Description","CreatedAt"}
	cate, _ := models.GetAllCategories(query, fields, sortby, order, offset, limit)
	if len(cate) == 0 {
		cateCreate := &models.Categories{
			Name		:	"default",
			DisplayName	:	"默认分类",
			Description	:	"默认生成的顶级分类",
			ParentId	:	0,
		}
		models.AddCategories(cateCreate)
		cate, _ = models.GetAllCategories(query, fields, sortby, order, offset, limit)
	}
	return tree(cate, 0, 0,0)
}

func GetCateByLike(param string) (ml map[int64]string) {
	var query map[string]string
	var fields []string
	var sortby []string
	var order []string
	var offset int64
	var limit int64

	query = make(map[string]string)
	if param != "" {
		query["display_name__icontains"] = param
	}

	fields = []string{"Id", "Name", "DisplayName", "ParentId", "Description"}
	cates,_ := models.GetAllCategories(query, fields, sortby, order, offset, limit)
	cate := make(map[int64]string, len(cates))
	for _, v := range cates{
		cate[v.(map[string]interface{})["Id"].(int64)] = v.(map[string]interface{})["DisplayName"].(string)
	}

	return cate
}


func GetCateByPostId(postId int64) (int64,error) {
	o := orm.NewOrm()
	artCate := models.ArticleCate{ArtId:postId}
	err := o.Read(&artCate,"ArtId")
	if err == orm.ErrNoRows {
		return 0,err
	}
	return artCate.CateId,nil
}

func tree(cate []interface{}, parent int64, level int64,key int64) ([]interface{}) {
	html := "-"
	var data []interface{}
	for _, v := range cate {
		var ParentId = v.(map[string]interface{})["ParentId"].(int64)
		var Id = v.(map[string]interface{})["Id"].(int64)
		if ParentId == parent {
			var newHtml string
			if level != 0 {
				newHtml = common.GoRepeat("&nbsp;&nbsp;&nbsp;&nbsp;", level) + "|"
			}
			v.(map[string]interface{})["html"] = newHtml + common.GoRepeat(html, level)
			data = append(data,v)
			data = common.GoMerge(data,tree(cate, Id, level+1,key+1))
		}
	}
	return data
}

func GetSimilar(beginId []int64,resIds []int64,level int) (beginId2 []int64,resIds2 []int64,level2 int) {
	if len(beginId) != 0 {
		o := orm.NewOrm()
		var cates []*models.Categories
		num,err := o.QueryTable(new(models.Categories)).Filter("ParentId__in",beginId).All(&cates)
		fmt.Println(num,err)
		if err != nil {
		}
		if num != 0 {
			if level == 0 {
				resIds2 = beginId
			} else {
				resIds2 = resIds
			}
			for _,v := range cates {
				id := v.Id
				beginId2 = append(beginId2,id)
				resIds2 = append(resIds2,id)
			}
			level2 = level + 1
			return GetSimilar(beginId2,resIds2,level2)
		}
		if level == 0 && num == 0 {
			return beginId,beginId,level
		} else {
			return beginId,resIds,level
		}
	}
	return beginId,resIds,level
}

func DeleteCateById(cateId int64) {
	o := orm.NewOrm()
	o.Begin()
	var postCateList []*models.ArticleCate
	num,delPostCateErr := o.QueryTable(new(models.ArticleCate)).Filter("CateId",cateId).All(&postCateList)

	if num != 0 {
		for _,v := range postCateList {
			var postCateUpdate models.ArticleCate
			postCateUpdate.CateId = 1
			postCateUpdate.Id = v.Id
			_, delPostCateErr = o.Update(&postCateUpdate,"CateId")
			if delPostCateErr != nil {
				o.Rollback()
			}
		}
	}

	cate := models.Categories{Id:cateId}
	_,err := o.Delete(&cate)

	if err == nil && delPostCateErr == nil {
		o.Commit()
	} else {
		o.Rollback()
	}
}