package services

import (
	"bee-go-myBlog/models"
	"time"
	"math/rand"
)

func GetMyAllCate() {

}

func GetAllCateBySort() ([]interface{}) {

	//var cateCreate models.Categories
	//cateCreate.Name = "名字-"+GetRandomString(4)
	//cateCreate.Description = "介绍-"+GetRandomString(16)
	//cateCreate.ParentId = rand.Int63n(10)
	//cateCreate.DisplayName = "别名-"+GetRandomString(8)
	//
	//id,err := models.AddCategories(&cateCreate)
	//fmt.Println(id,err)
	var query map[string]string
	var fields []string
	var sortby []string
	var order []string
	var offset int64
	var limit int64

	fields = []string{"Id", "Name", "DisplayName", "ParentId", "Description"}
	cate, _ := models.GetAllCategories(query, fields, sortby, order, offset, limit)

	//fmt.Println(cate)

	return  tree(cate, 0, 0,0)

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
				newHtml = GoRepeat("&nbsp;&nbsp;&nbsp;&nbsp;", level) + "|"
			}
			v.(map[string]interface{})["html"] = newHtml + GoRepeat(html, level)
			//fmt.Println(k,level,Id,ParentId,parent,"多少次")
			//fmt.Println(data)
			data[key] = v

			data = append(data, tree(cate, Id, level+1,key+1))

		}
	}
	return data
}

//func GoMerge(arr1 []interface{},arr2 []interface{}) ([]interface{}) {
//	var arr3 []interface{}
//	var i int64
//	i = 0
//	//for key,val := range arr1 {
//	//	for k,v := range arr2 {
//	//		if val == v {
//	//			arr3[i]
//	//		}
//	//	}
//	//}
//}


func GoRepeat(str string, num int64) (string) {
	var i int64
	newStr := ""
	if num != 0 {
		for i = 0; i < num; i++ {
			newStr += str
		}
	}
	return newStr
}


//生成随机字符串
func GetRandomString(l int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
