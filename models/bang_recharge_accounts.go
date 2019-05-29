package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type BangRechargeAccounts struct {
	Id             int       `orm:"column(id);auto"`
	UserId         uint      `orm:"column(user_id)" description:"用户ID"`
	Amount         uint      `orm:"column(amount)" description:"总额"`
	RechargeMoney  uint      `orm:"column(recharge_money)" description:"累计充值的钱（单位分）"`
	RechargeAmount uint      `orm:"column(recharge_amount)" description:"累计充值总额"`
	BonusAmount    uint      `orm:"column(bonus_amount)" description:"累计赠送总额"`
	Balance        uint      `orm:"column(balance)" description:"可用余额"`
	Version        uint      `orm:"column(version)" description:"版本"`
	CreatedAt      time.Time `orm:"column(created_at);type(timestamp);auto_now_add" description:"创建时间"`
	UpdatedAt      time.Time `orm:"column(updated_at);type(timestamp);auto_now" description:"更新时间"`
}

func (t *BangRechargeAccounts) TableName() string {
	return "bang_recharge_accounts"
}

func init() {
	orm.RegisterModel(new(BangRechargeAccounts))
}

// AddBangRechargeAccounts insert a new BangRechargeAccounts into database and returns
// last inserted Id on success.
func AddBangRechargeAccounts(m *BangRechargeAccounts) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBangRechargeAccountsById retrieves BangRechargeAccounts by Id. Returns error if
// Id doesn't exist
func GetBangRechargeAccountsById(id int) (v *BangRechargeAccounts, err error) {
	o := orm.NewOrm()
	v = &BangRechargeAccounts{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBangRechargeAccounts retrieves all BangRechargeAccounts matches certain condition. Returns empty list if
// no records exist
func GetAllBangRechargeAccounts(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BangRechargeAccounts))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []BangRechargeAccounts
	qs = qs.OrderBy(sortFields...)
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

// UpdateBangRechargeAccounts updates BangRechargeAccounts by Id and returns error if
// the record to be updated doesn't exist
func UpdateBangRechargeAccountsById(m *BangRechargeAccounts) (err error) {
	o := orm.NewOrm()
	v := BangRechargeAccounts{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBangRechargeAccounts deletes BangRechargeAccounts by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBangRechargeAccounts(id int) (err error) {
	o := orm.NewOrm()
	v := BangRechargeAccounts{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BangRechargeAccounts{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
