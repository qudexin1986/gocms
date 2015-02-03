package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type GxMaster struct {
	Id         int    `orm:"column(id);auto"`
	User       string `orm:"column(user);size(50)"`
	Pwd        string `orm:"column(pwd);size(32)"`
	Usertype   string `orm:"column(usertype);size(100)"`
	Logincount int16  `orm:"column(logincount)"`
	Loginip    string `orm:"column(loginip);size(40)"`
	Logintime  int    `orm:"column(logintime)"`
}

func init() {
	orm.RegisterModel(new(GxMaster))
}

// AddGxMaster insert a new GxMaster into database and returns
// last inserted Id on success.
func AddGxMaster(m *GxMaster) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGxMasterById retrieves GxMaster by Id. Returns error if
// Id doesn't exist
func GetGxMasterById(id int) (v *GxMaster, err error) {
	o := orm.NewOrm()
	v = &GxMaster{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGxMaster retrieves all GxMaster matches certain condition. Returns empty list if
// no records exist
func GetAllGxMaster(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GxMaster))
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

	var l []GxMaster
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateGxMaster updates GxMaster by Id and returns error if
// the record to be updated doesn't exist
func UpdateGxMasterById(m *GxMaster) (err error) {
	o := orm.NewOrm()
	v := GxMaster{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGxMaster deletes GxMaster by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGxMaster(id int) (err error) {
	o := orm.NewOrm()
	v := GxMaster{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GxMaster{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
