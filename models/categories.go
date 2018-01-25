package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type Categories struct {
	Id        		int64  		`orm:"column(id);auto;unique" json:"id"`
	Name 			string 		`orm:"column(name);size(255);unique" json:"name"`
	DisplayName 	string 		`orm:"column(display_name);size(255);" json:"display_name"`
	ParentId	 	int64 		`orm:"column(parent_id);default(0)" json:"parent_id"`
	Description	 	string 		`orm:"column(description);size(2048);null" json:"description"`
	CreatedAt      time.Time 	`orm:"column(created_at);default('0000-00-00 00:00:00');null;auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt      time.Time 	`orm:"column(updated_at);default('0000-00-00 00:00:00');null;auto_now;type(datetime)" json:"updated_at"`
}

func (c *Categories) TableName() string {
	return "y_categories"
}

func init() {
	orm.RegisterModel(new(Categories))
}

// AddCategories insert a new Categories into database and returns
// last inserted Id on success.
func AddCategories(m *Categories) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCategoriesById retrieves Categories by Id. Returns error if
// Id doesn't exist
func GetCategoriesById(id int64) (v *Categories, err error) {
	o := orm.NewOrm()
	v = &Categories{Id: id}
	if err = o.QueryTable(new(Categories)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCategories retrieves all Categories matches certain condition. Returns empty list if
// no records exist
func GetAllCategories(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Categories))
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

	var l []Categories
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

// UpdateCategories updates Categories by Id and returns error if
// the record to be updated doesn't exist
func UpdateCategoriesById(m *Categories) (err error) {
	o := orm.NewOrm()
	v := Categories{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCategories deletes Categories by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCategories(id int64) (err error) {
	o := orm.NewOrm()
	v := Categories{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Categories{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
