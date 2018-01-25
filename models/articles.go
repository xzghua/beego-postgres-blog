package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)


type Articles struct {
	Id        		int64  		`orm:"column(id);auto;unique" json:"id"`
	Title     		string 		`orm:"column(title);size(255)" json:"title"`
	Content   		string    	`orm:"column(content);type(longtext);null" json:"content"`
	BodyOriginal   	string    	`orm:"column(body_original);type(longtext);null" json:"body_original"`
	UserId    		int64    	`orm:"column(user_id);null;default(0)" json:"user_id"`
	Password   		int64    	`orm:"column(password);null" json:"password"`
	Note   			string    	`orm:"column(note);null" json:"note"`
	ReadStatus  	int64    	`orm:"column(read_status);default(1)" json:"read_status"`
	Top   			bool    	`orm:"column(top);null;default(false)" json:"top"`
	Abstract   		string    	`orm:"column(abstract);size(500);null" json:"abstract"`
	ViewNum   		int64    	`orm:"column(view_num);default(0);null" json:"view_num"`
	CreatedAt      time.Time 	`orm:"column(created_at);default('0000-00-00 00:00:00');null;auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt      time.Time 	`orm:"column(updated_at);default('0000-00-00 00:00:00');null;auto_now;type(datetime)" json:"updated_at"`

}

func (a *Articles) TableName() string {
	return "y_articles"
}


func init() {
	orm.RegisterModel(new(Articles))
}

// AddArticles insert a new Articles into database and returns
// last inserted Id on success.
func AddArticles(m *Articles) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArticlesById retrieves Articles by Id. Returns error if
// Id doesn't exist
func GetArticlesById(id int64) (v *Articles, err error) {
	o := orm.NewOrm()
	v = &Articles{Id: id}
	if err = o.QueryTable(new(Articles)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArticles retrieves all Articles matches certain condition. Returns empty list if
// no records exist
func GetAllArticles(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Articles))
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

	var l []Articles
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

// UpdateArticles updates Articles by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticlesById(m *Articles) (err error) {
	o := orm.NewOrm()
	v := Articles{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArticles deletes Articles by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticles(id int64) (err error) {
	o := orm.NewOrm()
	v := Articles{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Articles{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func AllArticle() (ml []interface{}, err error) {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	return GetAllArticles(query,fields,sortby,order,offset,limit)
}
