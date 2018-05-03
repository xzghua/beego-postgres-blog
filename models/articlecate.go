package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"errors"
	"reflect"
	"fmt"
	"bee-go-myBlog/common"
)

type ArticleCate struct {
	Id     	int64 `orm:"column(id);auto;unique" json:"id"`
	ArtId 	int64 `orm:"column(art_id);" json:"art_id"`
	CateId 	int64 `orm:"column(cate_id);" json:"cate_id"`
}

func (a *ArticleCate) TableName() string {
	return "y_article_cate"
}

func init() {
	orm.RegisterModel(new(ArticleCate))
}



//func AddArticleCate (postId int,cateId int) {
//	//先试试有没有
//}




 //AddArticleCate insert a new ArticleCate into database and returns
 //last inserted Id on success.
func AddArticleCate(m *ArticleCate) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}






// GetArticleCateById retrieves ArticleCate by Id. Returns error if
// Id doesn't exist
func GetArticleCateById(id int64) (v *ArticleCate, err error) {
	o := orm.NewOrm()
	v = &ArticleCate{Id: id}
	if err = o.QueryTable(new(ArticleCate)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArticleCate retrieves all ArticleCate matches certain condition. Returns empty list if
// no records exist
func GetAllArticleCate(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ArticleCate))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ArticleCate
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateArticleCate updates ArticleCate by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticleCateById(m *ArticleCate) (err error) {
	o := orm.NewOrm()
	v := ArticleCate{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArticleCate deletes ArticleCate by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticleCate(id int64) (err error) {
	o := orm.NewOrm()
	v := ArticleCate{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ArticleCate{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func GetCateIdByPostId(postId int64) (artCate ArticleCate,err error) {
	o := orm.NewOrm()
	artCate = ArticleCate{ArtId:postId}
	err = o.Read(&artCate,"ArtId")
	if err == nil {
		return artCate,nil
	}
	return artCate,err
}

func IndexCatePostPaginate(page int64,cateId int64) (totalPage int64,lastPage int64,currentPage int64,nextPage int64)  {
	var artCate ArticleCate
	tableName := artCate.TableName()
	return common.IndexCatePaginate(page,cateId,tableName)

}