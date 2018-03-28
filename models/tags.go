package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"bee-go-myBlog/helper"
)

type Tags struct {
	Id        		int64  		`orm:"column(id);auto;unique" json:"id"`
	Name 			string 		`orm:"column(name);size(255);unique" json:"name"`
	TagNum      	int64     	`orm:"column(tag_num);default(0);null" json:"tag_num"`

}

func (t *Tags) TableName() string {
	return "y_tags"
}


func init() {
	orm.RegisterModel(new(Tags))
}

// AddTags insert a new Tags into database and returns
// last inserted Id on success.
func AddTags(m *Tags) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTagsById retrieves Tags by Id. Returns error if
// Id doesn't exist
func GetTagsById(id int64) (v *Tags, err error) {
	o := orm.NewOrm()
	v = &Tags{Id: id}
	if err = o.QueryTable(new(Tags)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTags retrieves all Tags matches certain condition. Returns empty list if
// no records exist
func GetAllTags(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Tags))
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

	var l []Tags
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

// UpdateTags updates Tags by Id and returns error if
// the record to be updated doesn't exist
func UpdateTagsById(m *Tags) (err error) {
	o := orm.NewOrm()
	v := Tags{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTags deletes Tags by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTags(id int64) (err error) {
	o := orm.NewOrm()
	v := Tags{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Tags{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func AddTagWithUnique(name string) (id int64, err error) {
	o := orm.NewOrm()
	v := Tags{Name: name}
	if _, id, err := o.ReadOrCreate(&v, "name",); err == nil {
		return id,err
	}
	return
}


func AllTag(page int64) (ml []interface{}, err error) {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64
	var offset int64
	limit, _ = beego.AppConfig.Int64("page_offset")
	fields = []string{"Id", "Name","TagNum"}
	offset = (page - 1 ) * limit
	res, err := GetAllTags(query, fields, sortby, order, offset, limit)
	return res, err
}


func TagPaginate(page int64) (totalPage int64,lastPage int64,currentPage int64,nextPage int64)  {
	var tag Tags
	tableName := tag.TableName()
	return helper.MyPaginate(page,tableName)
}