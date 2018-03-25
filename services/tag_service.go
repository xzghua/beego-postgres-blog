package services

import (
	"bee-go-myBlog/models"

)

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
